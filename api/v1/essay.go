package v1

import (
	"fmt"
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
InsertEssayHandler
@author: LJR
@Description: 手动添加文章
@param ctx
@Router /admin/v1/essay/insertion [post]
*/
func InsertEssayHandler(ctx *gin.Context) {
	p := new(model.ParamInsertEssay)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("InsertEssay with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.InsertEssayService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "添加文章成功")
}

/*
DeleteEssayHandler
@author: LJR
@Description: 根据文章ID删除文章
@param ctx
@Router /admin/v1/essay/delete [delete]
*/
func DeleteEssayHandler(ctx *gin.Context) {
	essayId := ctx.Query("essay_id")
	if essayId == "" {
		zap.L().Error("DeleteEssay with invalid param", zap.Error(utils.GetError(50001)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(50001))
		return
	}
	if err := service.DeleteEssayService(essayId); err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "删除文章成功")
}

/*
GetEssayHandler
@author: LJR
@Description: 根据文章ID获取文章信息
@param ctx
@Router /admin/v1/essay [get]
*/
func GetEssayHandler(ctx *gin.Context) {
	essayId := ctx.Query("essay_id")
	if essayId == "" {
		zap.L().Error("GetEssay with invalid param", zap.Error(utils.GetError(50001)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(50001))
		return
	}
	essayInfo, err := service.GetEssayService(essayId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取文章成功",
		"data": *essayInfo,
	})
}

/*
UpdateEssayHandler
@author: LJR
@Description: 根据文章ID修改文章
@param ctx
@Router /admin/v1/essay/modify [put]
*/
func UpdateEssayHandler(ctx *gin.Context) {
	p := new(model.ParamEssayInfo)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateEssay with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.UpdateEssayService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "修改文章成功")
}

/*
CollectEssayHandler
@author: LJR
@Description: 用户收藏文章
@param ctx
@Router /api/v1/essay/collection [post]
*/
func CollectEssayHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	essay_id := ctx.Query("essay_id")
	if essay_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(essay_id)
	essayId := uint64(num)
	err := service.CollectEssayService(essayId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "文章收藏成功")
}

/*
CancelCollectEssayHandler
@author: LJR
@Description: 用户取消收藏文章
@param ctx
@Router /api/v1/essay/cancellation/collection [delete]
*/
func CancelCollectEssayHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	essay_id := ctx.Query("essay_id")
	if essay_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(essay_id)
	essayId := uint64(num)
	err := service.CancelCollectEssayService(essayId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "文章取消收藏成功")
}

/*
GetEssayListHandler
@author: LJR
@Description: 用户获取文章列表
@param ctx
@Router /admin/v1/essay/list [get]
*/
func GetEssayListHandler(ctx *gin.Context) {
	essayType := ctx.Query("essay_type")
	if essayType == "" {
		zap.L().Error("GetEssayList with invalid param", zap.Error(utils.GetError(50012)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(50012))
		return
	}
	essayInfoList, err := service.GetEssayListService(essayType)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "文章列表加载成功",
		"data": essayInfoList,
	})
}

/*
GetEssayByEssayIdHandler
@author: LJR
@Description: 用户根据文章ID获取文章信息
@param ctx
@Router /api/v1/essay [get]
*/
func GetEssayByEssayIdHandler(ctx *gin.Context) {
	essayId := ctx.Query("essay_id")
	userId := ctx.GetUint64("user_id")
	if essayId == "" {
		zap.L().Error("GetEssay with invalid param", zap.Error(utils.GetError(50001)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(50001))
		return
	}
	fmt.Println(userId)
	//essayInfo, essayWordList, err := service.GetEssayByEssayIdService(essayId)
	//m := map[string]interface{}{
	//	"essayInfo":     essayInfo,
	//	"essayWordList": essayWordList,
	//}
	essayDetail, err := service.GetEssayByEssayIdService(essayId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "文章内容获取成功",
		"data": essayDetail,
		//"data": m, 布局问题
	})
}

/*
GetCollectEssayListHandler
@author: LJR
@Description: 用户获取文章收藏列表
@param ctx
@Router /api/v1/essay/collection/list [get]
*/
func GetCollectEssayListHandler(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	if user_id == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	essayCollectInfo, err := service.GetCollectEssayService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取文章收藏列表成功",
		"data": essayCollectInfo,
	})
}

/*
GetEssayAllHandler
@author: LJR
@Description: 根据文章类型和页数等加载文章列表
@param ctx
@Router /admin/v1/essays [get]
*/
func GetEssayAllHandler(ctx *gin.Context) {
	essay_type := ctx.Query("essaytype")
	page_num := ctx.Query("pagenum")
	page_size := ctx.Query("pagesize")
	if page_num == "" || page_size == "" {
		zap.L().Error("GetEssayAll with invalid param", zap.Error(utils.GetError(10014)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(10014))
		return
	}
	essayList, total, err := service.GetEssayAllService(essay_type, page_num, page_size)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"essayList": essayList,
		"total":     total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "文章列表获取成功",
		"data": data,
	})
}

/*
GetAllEssayHandler
@author: LJR
@Description: 根据文章类型和页数等加载文章列表
@param ctx
@Router /api/v1/essays [get]
*/
func GetAllEssayHandler(ctx *gin.Context) {
	essay_type := ctx.Query("essaytype")
	page_num := ctx.Query("pagenum")
	page_size := ctx.Query("pagesize")
	if page_num == "" || page_size == "" {
		zap.L().Error("GetEssayAll with invalid param", zap.Error(utils.GetError(10014)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(10014))
		return
	}
	essayList, err := service.GetAllEssayService(essay_type, page_num, page_size)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "文章列表获取成功",
		"data": essayList,
	})
}
