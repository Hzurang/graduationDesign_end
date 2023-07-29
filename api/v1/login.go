package v1

import (
	"errors"
	"ginStudy/config"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/service"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/*
MobileSignUpHandler
@author: LJR
@Description: 手机号注册
@param ctx
@Router /api/v1/user/signup_mobile [post]
*/
func MobileSignUpHandler(ctx *gin.Context) {
	p := new(model.ParamMobilePasswordSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("MobileSignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	loginUser, err := service.MobileSignUpService(p, ctx.ClientIP())
	if err != nil {
		if errors.Is(err, utils.GetStrAndError("用户信息", 10010)) || errors.Is(err, utils.GetStrAndError("用户", 10010)) {
			utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
			return
		}
		zap.L().Error("MobileSignUpService failed", zap.Error(err))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "注册成功，马上登录",
		"data": *loginUser,
	})
}

/*
MobileLoginHandler
@author: LJR
@Description: 手机号登录
@param ctx
@Router /api/v1/user/login_mobile [post]
*/
func MobileLoginHandler(ctx *gin.Context) {
	p := new(model.ParamMobilePasswordLogin)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("MobileLogin with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	//fmt.Println(ctx.ClientIP())
	loginUser, err := service.MobileLoginService(p, ctx.ClientIP())
	if err != nil {
		if errors.Is(err, utils.GetStrAndError("用户", 10009)) {
			utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
			return
		}
		zap.L().Error("MobileLoginService failed", zap.Error(err))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	/* 改一下试试
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg": "登陆成功*_*",
		"data": *loginUser,
	})
	*/
	// 泛型写法
	utils.Response(ctx, http.StatusOK, utils.SUCCESS, *loginUser, "登陆成功*_*")
}

/*
SMSCodeLoginHandler
@author: LJR
@Description: 手机验证码登录
@param ctx
@Router /api/v1/user/login_sms_code [post]
*/
func SMSCodeLoginHandler(ctx *gin.Context) {
	p := new(model.ParamSMSCodeLogin)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("SMSCodeLogin with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	loginUser, err := service.SMSCodeLoginService(p, ctx.ClientIP())
	if err != nil {
		if errors.Is(err, utils.GetStrAndError("用户", 10009)) {
			utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
			return
		}
		zap.L().Error("SMSCodeLoginService failed", zap.Error(err))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "登陆成功*_*",
		"data": *loginUser,
	})
}

/*
LogoutHandler
@author: LJR
@Description: 退出登录
@param ctx
@Router /api/v1/user/logout [post]
*/
func LogoutHandler(ctx *gin.Context) {
	ctx.Set("user_id", nil)
	token := ctx.GetString("token")
	global.RD.ZAdd("JWT_AUTH_USER:Baned", redis.Z{Score: float64(time.Now().Unix()), Member: token})
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "用户登出成功")
}

/*
AdminLoginHandler
@author: LJR
@Description: 管理员登录
@param ctx
@Router /admin/v1/backend/login [post]
*/
func AdminLoginHandler(ctx *gin.Context) {
	username := ctx.Query("user_name")
	password := ctx.Query("password")
	if username == "" || password == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30034))
		return
	}
	adminUser, err := service.AdminLoginService(username, password, ctx.ClientIP())
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	/* 改一下试试
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg": "登陆成功*_*",
		"data": *loginUser,
	})
	*/
	// 泛型写法
	utils.Response(ctx, http.StatusOK, utils.SUCCESS, *adminUser, "登陆成功*_*")
}

/*
AdminLogoutHandler
@author: LJR
@Description: 退出登录
@param ctx
@Router /admin/v1/backend/logout [post]
*/
func AdminLogoutHandler(ctx *gin.Context) {
	ctx.Set("admin_id", nil)
	ctx.Set("user_name", nil)
	token := ctx.GetString("admin_token")
	global.RD.ZAdd("JWT_AUTH_ADMIN:Baned", redis.Z{Score: float64(time.Now().Unix()), Member: token})
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "管理员登出成功")
}

/*
FindPwdBySMSCodeHandler
@author: LJR
@Description: 用户通过验证码 + 密码重置密码
@param ctx
@Router /api/v1/user/neglect/password [put]
*/
func FindPwdBySMSCodeHandler(ctx *gin.Context) {
	p := new(model.ParamResetPwdBySMSCode)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("FindPwdBySMSCode with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.FindPwdBySMSCodeService(p)
	if err != nil {
		zap.L().Error("FindPwdBySMSCodeService failed", zap.Error(err))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "找回密码成功*_*",
		"data": nil,
	})
}
