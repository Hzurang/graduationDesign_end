package service

import (
	"ginStudy/dao"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"ginStudy/utils/encrypt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

/*
ModifyPwdByOldPwdService
@author: LJR
@Description: 根据旧密码修改新密码业务逻辑(重新登录)
@param p
@param userId
@param token
@return err
*/
func ModifyPwdByOldPwdService(p *model.ParamModifyPwdByOldPwd, userId uint64, token string) (err error) {
	password := encrypt.GetSHA256HashCode(p.OldPassword)
	user, err := dao.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	if password != user.Password {
		return utils.GetError(30021)
	}
	user.Password = encrypt.GetSHA256HashCode(p.Password)
	if err = dao.UpdateUserPasswordByUserId(user); err != nil {
		return err
	}
	global.RD.ZAdd("JWT_AUTH_USER:Baned", redis.Z{Score: float64(time.Now().Unix()), Member: token})
	return nil
}

/*
ModifyPwdBySMSCodeService
@author: LJR
@Description: 根据验证码修改新密码业务逻辑
@param p
@param token
@return err
*/
func ModifyPwdBySMSCodeService(p *model.ParamModifyPwdBySMSCode, token string) (err error) {
	user, err := dao.GetUserByMobile(p.Mobile)
	if err != nil {
		return err
	}
	smsCode, _ := global.RD.Get(user.Mobile).Result()
	if smsCode == "" {
		return utils.GetError(30009)
	}
	if smsCode != p.Code {
		return utils.GetError(30010)
	} else {
		user.Password = encrypt.GetSHA256HashCode(p.Password)
		if err := dao.UpdateUserPasswordByUserId(user); err != nil {
			return err
		}
		global.RD.ZAdd("JWT_AUTH_USER:Baned", redis.Z{Score: float64(time.Now().Unix()), Member: token})
		return nil
	}
	return
}

/*
BindEmailService
@author: LJR
@Description: 绑定邮箱业务逻辑
@param p
@param userId
@return err
*/
func BindEmailService(p *model.ParamBindEmail, userId uint64) (err error) {
	user, err := dao.GetUserByUserId(userId)
	flag := false
	if err != nil {
		return err
	}
	// 第一次绑定加积分，换绑不加
	if user.Email == "" {
		flag = true
	}
	user.Email = p.Email
	emailCode, _ := global.RD.Get(p.Email).Result()
	if emailCode == "" {
		return utils.GetError(30009)
	}
	if emailCode != p.Code {
		return utils.GetError(30010)
	} else {
		if err := dao.UpdateUserEmailByUserId(user); err != nil {
			return err
		}
		userInfo, err := dao.GetUserInfoByUserId(userId)
		if err != nil {
			return err
		}
		if flag == true {
			// 认证邮箱加一百积分
			userInfo.Integral = userInfo.Integral + 100
			_ = dao.UpdateUserInfoIntegralByUserId(userInfo)
			return nil
		}
		return nil
	}
	return
}

/*
GetUserInfoService
@author: LJR
@Description: 获取用户信息业务逻辑
@param userID
@return p
@return err
*/
func GetUserInfoService(userID uint64) (p *model.ParamUserInfo, err error) {
	userInfo, err := dao.GetUserInfoByUserId(userID)
	if err != nil {
		return nil, err
	}
	p = new(model.ParamUserInfo)
	p.UserId = userInfo.UserId
	p.LeXueAppId = userInfo.LeXueAppId
	p.Gender = userInfo.Gender
	p.School = userInfo.School
	p.Birthday = userInfo.Birthday
	p.Area = userInfo.Area
	p.NickName = userInfo.NickName
	p.HeadSculpture = userInfo.HeadSculpture
	p.Integral = userInfo.Integral
	p.EngLevel = userInfo.EngLevel
	p.WordNeedReciteNum = userInfo.WordNeedReciteNum
	p.LastStartTime = userInfo.LastStartTime
	p.Role = userInfo.Role
	p.InvitationCode = userInfo.InvitationCode
	p.Signature = userInfo.Signature
	return p, nil
}

/*
ModifyUserInfoService
@author: LJR
@Description: 修改用户信息业务逻辑
@param p
@param userID
@return err
*/
func ModifyUserInfoService(p *model.ParamModifyUserInfo, userId uint64) (err error) {
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return err
	}
	_, err = dao.GetUserInfoByLeXueAppIdAndUserId(p.LeXueAppId, userId)
	if err == nil {
		return utils.GetError(30022)
	}
	userInfo.LeXueAppId = p.LeXueAppId
	userInfo.Gender = p.Gender
	userInfo.School = p.School
	userInfo.Birthday = p.Birthday
	userInfo.Area = p.Area
	userInfo.NickName = p.NickName
	userInfo.Signature = p.Signature
	err = dao.UpdateUserInfoByUserId(userInfo)
	if err != nil {
		return err
	}
	return nil
}

/*
DisableUserService
@author: LJR
@Description: 禁用用户业务逻辑
@param user_id
@return err
*/
func DisableUserService(user_id string) (err error) {
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	_, err = dao.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	err = dao.SoftDeleteUserByUserId(userId)
	if err != nil {
		return err
	}
	return nil
}

/*
RecoverUserService
@author: LJR
@Description: 恢复用户权限业务逻辑
@param user_id
@return err
*/
func RecoverUserService(user_id string) (err error) {
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	user, err := dao.GetSoftDeleteUserByUserId(userId)
	if err != nil {
		return err
	}
	err = dao.UpdateSoftDeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

/*
ResetEngLevelService
@author: LJR
@Description: 用户重置词书业务逻辑
@param userId
@return err
*/
func ResetEngLevelService(userId uint64) (err error) {
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return err
	}
	if userInfo.EngLevel != 0 {
		_ = dao.UpdatesResetWordLearnByUserIdAndWordType(userId, userInfo.EngLevel)
	}
	userInfo.EngLevel = 0
	err = dao.UpdateUserInfoEngLevelByUserId(userInfo)
	if err != nil {
		return err
	}
	return nil
}

/*
ModifyEngLevelService
@author: LJR
@Description: 用户设置词书业务逻辑
@param userId
@return err
*/
func ModifyEngLevelService(userId uint64, engLevel string) (err error) {
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return err
	}
	num, _ := strconv.Atoi(engLevel)
	englevel := int8(num)
	if englevel == 0 {
		return utils.GetError(30019)
	}
	userInfo.EngLevel = englevel
	err = dao.UpdateUserInfoEngLevelByUserId(userInfo)
	if err != nil {
		return err
	}
	return nil
}

/*
ModifyDailyWordService
@author: LJR
@Description: 用户更新每日单词量业务逻辑
@param userId
@param engLevel
@return err
*/
func ModifyDailyWordService(userId uint64, wordDailyNum int) (err error) {
	_, err = dao.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	userInfo := &model.UserInfo{UserId: userId, WordNeedReciteNum: wordDailyNum}
	err = dao.UpdateUserInfoWordNeedReciteNumByUserId(userInfo)
	if err != nil {
		return err
	}
	return nil
}

/*
GetVocabularyService
@author: LJR
@Description: 设置词书业务逻辑
@param userId
@param engLevel
@return err
*/
func GetVocabularyService(userId uint64) (paramUserEngLevel *model.ParamUserEnglevel, err error) {
	_, err = dao.GetUserByUserId(userId)
	if err != nil {
		return nil, err
	}
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return nil, err
	}
	paramUserEngLevel.EngLevel = userInfo.EngLevel
	paramUserEngLevel.WordNeedReciteNum = userInfo.WordNeedReciteNum
	return paramUserEngLevel, nil
}

/*
GetUpToDateLearnService
@author: LJR
@Description: 获取最新学习记录业务逻辑
@param userId
@param engLevel
@return err
*/
func GetUpToDateLearnService(userId uint64) (userDate *model.UserDate, err error) {
	_, err = dao.GetUserByUserId(userId)
	if err != nil {
		return nil, err
	}
	userDate, err = dao.GetUserDateByUserId(userId)
	if err != nil {
		return nil, err
	}
	return userDate, nil
}

/*
FinishLearnService
@author: LJR
@Description: 打卡业务逻辑
@param userId
@param engLevel
@return err
*/
func FinishLearnService(userId uint64, p *model.ParamUserDate) (err error) {
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return err
	}
	userDate := &model.UserDate{
		UserId:           userId,
		Date:             p.Date,
		WordLearnNumber:  p.WordLearnNumber,
		WordReviewNumber: p.WordReviewNumber,
		Remark:           p.Remark,
	}
	err = dao.CreateUserDate(userDate)
	if err != nil {
		return err
	}
	userInfo.Integral = userInfo.Integral + 10
	err = dao.UpdateUserInfoIntegralByUserId(userInfo)
	if err != nil {
		return err
	}
	return nil
}

/*
UpdateLearnService
@author: LJR
@Description: 补打卡业务逻辑
@param userId
@param engLevel
@return err
*/
func UpdateLearnService(userId uint64, date time.Time) (err error) {
	userInfo, err := dao.GetUserInfoByUserId(userId)
	if err != nil {
		return err
	}
	err = dao.GetUserDateByUserIdAndDate(userId, date)
	if err != nil {
		return err
	}
	if userInfo.Integral < 10 {
		return utils.GetError(30039)
	}
	userInfo.Integral = userInfo.Integral - 10
	_ = dao.UpdateUserInfoIntegralByUserId(userInfo)
	return nil
}

/*
GetUserAllService
@author: LJR
@Description: 根据页尺寸和页数等加载用户列表
@param page_num
@param page_size
@return WordList, total, err
*/
func GetUserAllService(page_num string, page_size string) (userList []model.User, total int64, err error) {
	pageNum, _ := strconv.Atoi(page_num)
	pageSize, _ := strconv.Atoi(page_size)
	userList, total, err = dao.GetUserPage(pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return userList, total, nil
}

/*
GetUserByMobileService
@author: LJR
@Description: 输入用户手机号进行查找业务逻辑
@param mobile
@return user, total, err
*/
func GetUserByMobileService(mobile string) (user *model.User, total int, err error) {
	user, err = dao.GetUserByMobile(mobile)
	if err != nil {
		return nil, 0, err
	}
	return user, 1, nil
}

/*
MineInfoService
@author: LJR
@Description: 用户进入我的页面时获取数据业务逻辑
@param id
@return essayInfo
@return essayWordList
@return err
*/
func MineInfoService(userId uint64) (paramMine *model.ParamMine, err error) {
	// 获取用户信息
	// 获取打卡信息的信息
	var wg sync.WaitGroup
	var userInfo *model.UserInfo
	var total int64
	var word int
	var err1 error
	var err2 error
	var err3 error
	paramMine = new(model.ParamMine)
	wg.Add(3)
	go func() {
		defer wg.Done()
		userInfo, err1 = dao.GetUserInfoByUserId(userId)
		paramMine.NickName = userInfo.NickName
		paramMine.Integral = strconv.FormatUint(userInfo.Integral, 10)
	}()
	go func() {
		defer wg.Done()
		total, err2 = dao.GetDaysTotalByUserId(userId)
		paramMine.Days = strconv.Itoa(int(total))
	}()
	go func() {
		defer wg.Done()
		word, err3 = dao.CalculateTotalWordLearnNumber(userId)
		paramMine.Words = strconv.Itoa(word)
	}()
	wg.Wait()
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}
	return paramMine, nil
}

func CheckLearnListService(userId uint64) (paramDateList []model.ParamDate, err error) {
	userDateList, err := dao.GetUserDateListByUserId(userId)
	if err != nil {
		return nil, err
	}
	paramDateList = make([]model.ParamDate, 0, len(userDateList))
	for _, userDate := range userDateList {
		dateStr := userDate.Date.Format("2006-01-02")
		parsedTime, _ := time.Parse("2006-01-02", dateStr)
		param := model.ParamDate{
			Year:             parsedTime.Year(),
			Month:            int(parsedTime.Month()),
			Date:             parsedTime.Day(),
			WordLearnNumber:  userDate.WordLearnNumber,
			WordReviewNumber: userDate.WordReviewNumber,
			Remark:           userDate.Remark,
		}
		paramDateList = append(paramDateList, param)
	}
	return paramDateList, nil
}

func CalendarUserInfoService(userId uint64, date string, month string, year string) (paramDate *model.ParamDate, err error) {
	Date, _ := strconv.Atoi(date)
	Month, _ := strconv.Atoi(month)
	Year, _ := strconv.Atoi(year)
	dayDate := time.Date(Year, time.Month(Month), Date, 0, 0, 0, 0, time.UTC)
	userDate, err := dao.GetUserDateByUserIdAndOther(userId, dayDate)
	if err != nil {
		return nil, err
	}
	dateStr := userDate.Date.Format("2006-01-02")
	parsedTime, _ := time.Parse("2006-01-02", dateStr)
	paramDate.Date = parsedTime.Day()
	paramDate.Month = int(parsedTime.Month())
	paramDate.Year = parsedTime.Year()
	paramDate.WordReviewNumber = userDate.WordReviewNumber
	paramDate.WordLearnNumber = userDate.WordLearnNumber
	paramDate.Remark = userDate.Remark
	return paramDate, nil
}
