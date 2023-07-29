package v1

import (
	"ginStudy/config"
	"ginStudy/model"
	"ginStudy/service"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

/*
ModifyPwdByOldPwdHandler
@author: LJR
@Description: 根据旧密码修改新密码
@param ctx
@Router /api/v1/user/modify_pwd_oldpwd [put]
*/
func ModifyPwdByOldPwdHandler(ctx *gin.Context) {
	p := new(model.ParamModifyPwdByOldPwd)
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("ModifyPwdByOldPwd with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	str, _ := ctx.Get("token")
	token := str.(string)
	err := service.ModifyPwdByOldPwdService(p, userId, token)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "修改成功，返回登录")
}

/*
ModifyPwdBySMSCodeHandler
@author: LJR
@Description: 根据验证码修改新密码
@param ctx
@Router /api/v1/user/modify_pwd_smscode [put]
*/
func ModifyPwdBySMSCodeHandler(ctx *gin.Context) {
	p := new(model.ParamModifyPwdBySMSCode)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("ModifyPwdBySMSCode with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	str, _ := ctx.Get("token")
	token := str.(string)
	err := service.ModifyPwdBySMSCodeService(p, token)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "修改成功，返回登录")
}

/*
BindEmailHandler
@author: LJR
@Description: 绑定邮箱
@param ctx
@Router /api/v1/user/improvement/email [put]
*/
func BindEmailHandler(ctx *gin.Context) {
	p := new(model.ParamBindEmail)
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("BindEmail with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.BindEmailService(p, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "绑定邮箱成功")
}

/*
GetUserInfoHandler
@author: LJR
@Description: 获取用户信息
@param ctx
@Router /api/v1/user/userInfo [get]
*/
func GetUserInfoHandler(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")
	if userID == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30017))
		return
	}
	userInfo, err := service.GetUserInfoService(userID)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取用户信息成功",
		"data": *userInfo,
	})
}

/*
ModifyUserInfoHandler
@author: LJR
@Description: 修改用户信息
@param ctx
@Router /api/v1/user/modify/userInfo [put]
*/
func ModifyUserInfoHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	p := new(model.ParamModifyUserInfo)
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("ModifyUserInfo with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.ModifyUserInfoService(p, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "修改用户信息成功")
}

/*
DisableUserHandler
@author: LJR
@Description: 禁用用户
@param ctx
@Router /admin/v1/user/prohibit [delete]
*/
func DisableUserHandler(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetStrAndError("，禁用失败", 30025))
		return
	}
	err := service.DisableUserService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "用户禁用成功")
}

/*
RecoverUserHandler
@author: LJR
@Description: 恢复用户权限
@param ctx
@Router /admin/v1/user/recovery [put]
*/
func RecoverUserHandler(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetStrAndError("，禁用失败", 30025))
		return
	}
	err := service.RecoverUserService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "用户权限恢复成功")
}

/*
ResetEngLevelHandler
@author: LJR
@Description: 重置词书
@param ctx
@Router /api/v1/user/reset/vocabulary [put]
*/
func ResetEngLevelHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	err := service.ResetEngLevelService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "用户重置词书成功")
}

/*
ModifyEngLevelHandler
@author: LJR
@Description: 设置词书
@param ctx
@Router /api/v1/user/modify/vocabulary [put]
*/
func ModifyEngLevelHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	engLevel := ctx.Query("eng_level")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	err := service.ModifyEngLevelService(userId, engLevel)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "用户设置词书成功")
}

/*
ModifyDailyWordHandler
@author: LJR
@Description: 设置每日单词要背的量
@param ctx
@Router /api/v1/user/modify/daily/word [put]
*/
func ModifyDailyWordHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	word_daily_num := ctx.Query("word_daily_num")
	if userId == 0 || word_daily_num == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	wordDailyNum, _ := strconv.Atoi(word_daily_num)
	err := service.ModifyDailyWordService(userId, wordDailyNum)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, nil, "用户更新每日单词量成功")
}

/*
GetVocabularyHandler
@author: LJR
@Description: 设置词书
@param ctx
@Router /api/v1/user/vocabulary [get]
*/
func GetVocabularyHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	engLevel, err := service.GetVocabularyService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取成功",
		"data": *engLevel,
	})
}

/*
GetUpToDateLearnHandler
@author: LJR
@Description: 获取最新的学习记录
@param ctx
@Router /api/v1/user/learn/up_to_date [get]
*/
func GetUpToDateLearnHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	userDate, err := service.GetUpToDateLearnService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取最新学习记录成功",
		"data": *userDate,
	})
}

/*
FinishLearnHandler
@author: LJR
@Description: 打卡
@param ctx
@Router /api/v1/user/learn/finish [post]
*/
func FinishLearnHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	p := new(model.ParamUserDate)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("FinishLearn with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.FinishLearnService(userId, p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "打卡成功",
		"data": nil,
	})
}

/*
MineInfoHandler
@author: LJR
@Description: 用户进入我的页面时获取数据
@param ctx
@Router /api/v1/user/mine [get]
*/
func MineInfoHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	mine, err := service.MineInfoService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取我的页面的数据成功",
		"data": mine,
	})
}

func CalendarUserInfoHandler(ctx *gin.Context) {
	date := ctx.Query("date")
	month := ctx.Query("month")
	year := ctx.Query("year")
	userId := ctx.GetUint64("user_id")
	if date == "" || month == "" || year == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(10001))
		return
	}
	param, err := service.CalendarUserInfoService(userId, date, month, year)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取成功",
		"data": *param,
	})
}

/*
GetUserAllHandler
@author: LJR
@Description: 根据页尺寸和页数等加载用户列表
@param ctx
@Router /admin/v1/users [get]
*/
func GetUserAllHandler(ctx *gin.Context) {
	page_num := ctx.Query("pagenum")
	page_size := ctx.Query("pagesize")
	if page_num == "" || page_size == "" {
		zap.L().Error("GetWordAll with invalid param", zap.Error(utils.GetError(10014)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(10014))
		return
	}
	userList, total, err := service.GetUserAllService(page_num, page_size)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"userList": userList,
		"total":    total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "用户列表获取成功",
		"data": data,
	})
}

/*
GetUserByMobileHandler
@author: LJR
@Description: 输入用户手机号进行查找
@param ctx
@Router /admin/v1/user/signal [get]
*/
func GetUserByMobileHandler(ctx *gin.Context) {
	mobile := ctx.Query("mobile")
	if mobile == "" {
		zap.L().Error("GetUserByMobile with invalid param", zap.Error(utils.GetError(30004)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(30004))
		return
	}
	user, total, err := service.GetUserByMobileService(mobile)
	userList := make([]*model.User, 0, 1)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	userList = append(userList, user)
	data := map[string]interface{}{
		"userList": userList,
		"total":    total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "用户获取成功",
		"data": data,
	})
}

/*
CheckInLearnHandler
@author: LJR
@Description: 补打卡
@param ctx
@Router /api/v1/user/learn/check [post]
*/
func CheckInLearnHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	date := ctx.Query("selected_date")
	selectedDate, err := time.Parse("2006-01-02", date)
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	err = service.UpdateLearnService(userId, selectedDate)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "补打卡成功",
		"data": nil,
	})
}

/*
CheckLearnListHandler
@author: LJR
@Description: 用户获取打卡日期列表
@param ctx
@Router /api/v1/user/check/list [get]
*/
func CheckLearnListHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	if userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	list, err := service.CheckLearnListService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取成功",
		"data": list,
	})
}
