package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Response
@author: LJR
@Description: 统一返回格式 (泛型数据)
@param ctx
@param httpStatus
@param code
@param data 泛型数据
@param msg
*/
func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

/*
ResponseGin
@author: LJR
@Description: 统一返回格式 (传统数据)
@param ctx
@param httpStatus
@param code
@param data 数据
@param msg
*/
func ResponseGin(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

/*
ResponseParamError
@author: LJR
@Description: 一般用于请求参数缺漏的返回
@param ctx
@param code
@param data
@param err
*/
func ResponseParamError(ctx *gin.Context, code int, data gin.H, err error) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": code,
		"msg":  err.Error(),
		"data": data,
	})
}

/*
ResponseParamTypeError
@author: LJR
@Description: 一般用于请求参数类型错误的返回
@param ctx
@param code
@param data
@param err
*/
func ResponseParamTypeError(ctx *gin.Context, code int, data gin.H, err map[string]string) {
	var errs string
	for _, v := range err {
		errs = errs + v + " "
	}
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": code,
		"msg":  errs,
		"data": data,
	})
}

/*
ResponseError
@author: LJR
@Description: 封装通用返回错误
@param ctx
@param code
@param data
@param err
*/
func ResponseError(ctx *gin.Context, code int, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResponseSuccess(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func ResponseFail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
