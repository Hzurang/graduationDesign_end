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
)

/*
InsertListenHandler
@author: LJR
@Description: 手动添加听力
@param ctx
@Router /admin/v1/listen/insertion [post]
*/
func InsertListenHandler(ctx *gin.Context) {
	p := new(model.ParamInsertListen)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("InsertListen with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.InsertListenService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "添加听力成功")
}

/*
DeleteListenHandler
@author: LJR
@Description: 根据听力ID删除听力
@param ctx
@Router /admin/v1/listen/delete [delete]
*/
func DeleteListenHandler(ctx *gin.Context) {
	listenId := ctx.Query("listen_id")
	if listenId == "" {
		zap.L().Error("DeleteListen with invalid param", zap.Error(utils.GetError(60002)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(60002))
		return
	}
	if err := service.DeleteListenService(listenId); err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "删除听力成功")
}

/*
GetListenHandler
@author: LJR
@Description: 根据听力ID获取听力信息
@param ctx
@Router /admin/v1/listen [get]
*/
func GetListenHandler(ctx *gin.Context) {
	listenId := ctx.Query("listen_id")
	if listenId == "" {
		zap.L().Error("GetListen with invalid param", zap.Error(utils.GetError(60002)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(60002))
		return
	}
	listenInfo, err := service.GetListenService(listenId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取听力成功",
		"data": *listenInfo,
	})
}

/*
UpdateListenHandler
@author: LJR
@Description: 根据听力ID修改听力
@param ctx
@Router /admin/v1/listen/modify [put]
*/
func UpdateListenHandler(ctx *gin.Context) {
	p := new(model.ParamListenInfo)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateListen with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.UpdateListenService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "修改听力成功")
}

/*
GetListenListHandler
@author: LJR
@Description: 用户获取听力列表（传入类型）
@param ctx
@Router /admin/v1/listen/list [get]
*/
func GetListenListHandler(ctx *gin.Context) {
	listenType := ctx.Query("listen_type")
	if listenType == "" {
		zap.L().Error("GetListenList with invalid param", zap.Error(utils.GetError(60007)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(60007))
		return
	}
	listenInfoList, err := service.GetListenListService(listenType)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "听力列表加载成功",
		"data": listenInfoList,
	})
	return
}

/*
GetListenByListenIdHandler
@author: LJR
@Description: 用户根据听力ID获取听力信息
@param ctx
@Router /api/v1/listen [get]
*/
func GetListenByListenIdHandler(ctx *gin.Context) {
	listenId := ctx.Query("listen_id")
	if listenId == "" {
		zap.L().Error("GetListen with invalid param", zap.Error(utils.GetError(60002)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(60002))
		return
	}
	listenInfo, listenWordList, err := service.GetListenByListenIdService(listenId)
	m := map[string]interface{}{
		"listenInfo":     listenInfo,
		"listenWordList": listenWordList,
	}
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "听力内容获取成功",
		"data": m,
	})
}

/*
CollectListenHandler
@author: LJR
@Description: 用户收藏听力
@param ctx
@Router /api/v1/listen/collection [post]
*/
func CollectListenHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	listen_id := ctx.Query("listen_id")
	if listen_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(listen_id)
	listenId := uint64(num)
	err := service.CollectListenService(listenId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "听力收藏成功")
}

/*
CancelCollectListenHandler
@author: LJR
@Description: 用户取消收藏听力
@param ctx
@Router /api/v1/listen/cancellation/collection [delete]
*/
func CancelCollectListenHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	listen_id := ctx.Query("listen_id")
	if listen_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(listen_id)
	listenId := uint64(num)
	err := service.CancelCollectListenService(listenId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "听力取消收藏成功")
}

/*
GetCollectListenListHandler
@author: LJR
@Description: 用户获取听力收藏列表
@param ctx
@Router /api/v1/listen/collection/list [get]
*/
func GetCollectListenListHandler(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	if user_id == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	listenCollectInfo, err := service.GetCollectListenService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取听力收藏列表成功",
		"data": listenCollectInfo,
	})
}

/*
GetListenAllHandler
@author: LJR
@Description: 根据听力类型和页数等加载听力列表
@param ctx
@Router /admin/v1/listens [get]
*/
func GetListenAllHandler(ctx *gin.Context) {
	listen_type := ctx.Query("listentype")
	page_num := ctx.Query("pagenum")
	page_size := ctx.Query("pagesize")
	if page_num == "" || page_size == "" {
		zap.L().Error("GetListenAll with invalid param", zap.Error(utils.GetError(10014)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(10014))
		return
	}
	litenList, total, err := service.GetListenAllService(listen_type, page_num, page_size)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"listenList": litenList,
		"total":      total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "听力列表获取成功",
		"data": data,
	})
}
