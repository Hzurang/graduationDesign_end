package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"ginStudy/utils/encrypt"
	"strconv"
	"time"
)

/*
MobileSignUpService
@author: LJR
@Description: 处理手机号注册业务
@param p
@param clientIP
@return loginUser
@return err
*/
func MobileSignUpService(p *model.ParamMobilePasswordSignUp, clientIP string) (loginUser *model.LoginUser, err error) {
	_, err = dao.GetUserByMobile(p.Mobile)
	if err == nil {
		return nil, utils.GetError(30007)
	}
	user := new(model.User)
	smsCode, _ := global.RD.Get(p.Mobile).Result()
	fmt.Println(smsCode)
	if smsCode == "" {
		return nil, utils.GetError(30009)
	}
	if smsCode != p.Code {
		return nil, utils.GetError(30010)
	} else {
		if p.InvitationCode != "" {
			userInfo, err := dao.GetUserInfoByInvitationCode(p.InvitationCode)
			if err != nil {
				// 邀请一个好友加两百积分
				userInfo.Integral = userInfo.Integral + 200
				_ = dao.UpdateUserInfoIntegralByUserId(userInfo)
			} else {
				return nil, utils.GetError(30024)
			}
		}
		// 密码加密成密文，准备写入数据库
		user.Password = encrypt.GetSHA256HashCode(p.Password)
		user.UserId, _ = config.GenID()
		user.Mobile = p.Mobile
		user.CreatedAt, user.UpdatedAt, user.LastLoginTime = time.Now(), time.Now(), time.Now()
		err = dao.CreateUser(user)
		if err != nil {
			return nil, err
		}
		id := strconv.FormatUint(user.UserId, 10)
		id = id[1:3] + id[len(id)-6:]
		var codeStr string
		for {
			codeStr = utils.GetInvCodeByUID(user.UserId, 16)
			_, err = dao.GetUserInfoByInvitationCode(codeStr)
			if err == nil {
				break
			}
		}
		userInfo := &model.UserInfo{
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserId:         user.UserId,
			LeXueAppId:     user.Mobile,
			NickName:       fmt.Sprintf("乐学英语_%s", id),
			HeadSculpture:  "https://img0.baidu.com/it/u=2556841078,3045963903&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
			InvitationCode: codeStr,
		}
		err = dao.CreateUserInfo(userInfo)
		if err != nil {
			return nil, err
		}
		userID := strconv.FormatUint(user.UserId, 10)
		loginUser = &model.LoginUser{
			UserId: userID,
			Email:  user.Email,
			Mobile: user.Mobile,
			IP:     clientIP,
		}
		loginUser.AcToken, loginUser.ReToken, err = config.GenToken(clientIP, user.UserId, user.Mobile, user.Email, "")
		global.RD.Del(user.Mobile)
		return loginUser, nil
	}
}

/*
MobileLoginService
@author: LJR
@Description: 处理手机号登录逻辑
@param p
@param clientIP
@return loginUser
@return err
*/
func MobileLoginService(p *model.ParamMobilePasswordLogin, clientIP string) (loginUser *model.LoginUser, err error) {
	user, err := dao.GetUserByMobile(p.Mobile)
	if err != nil {
		return nil, err
	}
	if user.Status == 1 {
		return nil, utils.GetError(30018)
	}
	var passwordStr = encrypt.GetSHA256HashCode(p.Password)
	if passwordStr != user.Password {
		return nil, utils.GetError(30002)
	}
	englevel, _ := dao.GetEngLevelByUserId(user.UserId)
	userID := strconv.FormatUint(user.UserId, 10)
	loginUser = &model.LoginUser{
		UserId:   userID,
		Email:    user.Email,
		Mobile:   user.Mobile,
		IP:       clientIP,
		EngLevel: strconv.Itoa(int(*englevel)),
	}
	user.LastLoginTime = time.Now()
	err = dao.UpdateUserLoginTimeByUserId(user)
	if err != nil {
		return nil, err
	}
	loginUser.AcToken, loginUser.ReToken, err = config.GenToken(clientIP, user.UserId, user.Mobile, user.Email, "")
	if err != nil {
		return nil, err
	}
	return loginUser, nil
}

/*
SMSCodeLoginService
@author: LJR
@Description: 处理短信验证码登录逻辑
@param p
@param clientIP
@return loginUser
@return err
*/
func SMSCodeLoginService(p *model.ParamSMSCodeLogin, clientIP string) (loginUser *model.LoginUser, err error) {
	user, err := dao.GetUserByMobile(p.Mobile)
	if err != nil {
		return nil, err
	}
	if user.Status == 1 {
		return nil, utils.GetError(30018)
	}
	smsCode, _ := global.RD.Get(user.Mobile).Result()
	if smsCode == "" {
		return nil, utils.GetError(30009)
	}
	if smsCode != p.Code {
		return nil, utils.GetError(30010)
	} else {
		englevel, _ := dao.GetEngLevelByUserId(user.UserId)
		userID := strconv.FormatUint(user.UserId, 10)
		loginUser = &model.LoginUser{
			UserId:   userID,
			Email:    user.Email,
			Mobile:   user.Mobile,
			IP:       clientIP,
			EngLevel: strconv.Itoa(int(*englevel)),
		}
		user.LastLoginTime = time.Now()
		err = dao.UpdateUserLoginTimeByUserId(user)
		if err != nil {
			return nil, err
		}
		loginUser.AcToken, loginUser.ReToken, err = config.GenToken(clientIP, user.UserId, user.Mobile, user.Email, "")
		if err != nil {
			return nil, err
		}
		global.RD.Del(user.Mobile)
		return loginUser, nil
	}
}

/*
AdminLoginService
@author: LJR
@Description: 管理员登录业务逻辑
@param username
@param password
@param clientIP
@return adminUser
@return err
*/
func AdminLoginService(username string, password string, clientIP string) (adminUser *model.AdminUser, err error) {
	admin, err := dao.GetAdminUserByUserName(username)
	if err != nil {
		return nil, err
	}
	password = encrypt.GetSHA256HashCode(password)
	if admin.Password != password {
		return nil, utils.GetError(30036)
	}
	adminUser = &model.AdminUser{
		AdminId:  admin.AdminId,
		UserName: admin.UserName,
	}
	adminUser.AcToken, adminUser.ReToken, err = config.GenToken(clientIP, admin.AdminId, "", "", admin.UserName)
	if err != nil {
		return nil, err
	}
	return adminUser, nil
}

/*
FindPwdBySMSCodeService
@author: LJR
@Description: 用户通过验证码 + 密码重置密码业务逻辑
@param p
@param token
@return err
*/
func FindPwdBySMSCodeService(p *model.ParamResetPwdBySMSCode) (err error) {
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
		return nil
	}
	return nil
}
