package v1

import (
	"ginStudy/service"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
SendSMSCodeHandler
@author: LJR
@Description: 发送短信验证码
@param ctx
@Router /api/v1/user/sms_code [get]
*/
func SendSMSCodeHandler(ctx *gin.Context) {
	mobile := ctx.Query("mobile")
	if mobile == "" {
		utils.ResponseGin(ctx, http.StatusUnprocessableEntity, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(30004))
		return
	}
	flag := utils.VerifyMobileFormat(mobile)
	if flag != true {
		utils.ResponseGin(ctx, http.StatusUnprocessableEntity, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(30005))
		return
	}
	_, err := service.SendSMSCodeService(mobile)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "发送验证码成功*_*")
	//utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, gin.H{
	//	"code": code,
	//}, "发送验证码成功*_*")
}

/*
SendEmailCodeHandler
@author: LJR
@Description: 发送邮箱验证码
@param ctx
@Router /api/v1/user/email_code [get]
*/
func SendEmailCodeHandler(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		utils.ResponseGin(ctx, http.StatusUnprocessableEntity, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(30011))
		return
	}
	flag := utils.VerifyEmailFormat(email)
	if flag != true {
		utils.ResponseGin(ctx, http.StatusUnprocessableEntity, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(30012))
		return
	}
	err := service.SendEmailCodeService(email)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "发送验证码成功*_*")
}
