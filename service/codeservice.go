package service

import (
	"ginStudy/config"
	"ginStudy/global"
	"ginStudy/utils"
	"time"
)

/*
SendSMSCodeService
@author: LJR
@Description: 发送手机验证码业务
@param mobile
@return code
@return err
*/
func SendSMSCodeService(mobile string) (code string, err error) {
	code, err = config.SMSCode(mobile)
	if code != "" {
		err := global.RD.Set(mobile, code, 120*time.Second).Err()
		if err != nil {
			return "", utils.GetError(30008)
		}
		return code, nil
	}
	return
}

/*
SendEmailCodeService
@author: LJR
@Description: 发送邮箱验证码业务
@param email
@return err
*/
func SendEmailCodeService(email string) (err error) {
	code, err := config.EmailCode(email)
	if code != "" {
		err := global.RD.Set(email, code, 120*time.Second).Err()
		if err != nil {
			return utils.GetError(30008)
		}
		return nil
	}
	return
}
