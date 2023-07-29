package v1

import (
	"ginStudy/service"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
CollectSentenceHandler
@author: LJR
@Description: 用户收藏句子
@param ctx
@Router /api/v1/sentence/collection [post]
*/
func CollectSentenceHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	sentence_id := ctx.Query("sentence_id")
	if sentence_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(70003))
		return
	}
	num, _ := strconv.Atoi(sentence_id)
	sentenceId := uint64(num)
	err := service.CollectSentenceService(sentenceId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "句子收藏成功")
}

/*
CancelCollectSentenceHandler
@author: LJR
@Description: 用户取消句子文章
@param ctx
@Router /api/v1/sentence/cancellation/collection [delete]
*/
func CancelCollectSentenceHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	sentence_id := ctx.Query("sentence_id")
	if sentence_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(70003))
		return
	}
	num, _ := strconv.Atoi(sentence_id)
	sentenceId := uint64(num)
	err := service.CancelCollectSentenceService(sentenceId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "文章取消收藏成功")
}

/*
DailySentenceHandler
@author: LJR
@Description: 获取每日一句
@param ctx
@Router /api/v1/sentence/daily [get]
*/
func DailySentenceHandler(ctx *gin.Context) {
	sentence, err := service.DailySentenceService()
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取每日一句成功",
		"data": *sentence,
	})
}

/*
GetCollectSentenceListHandler
@author: LJR
@Description: 用户获取句子收藏列表
@param ctx
@Router /api/v1/sentence/collection/list [get]
*/
func GetCollectSentenceListHandler(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	if user_id == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(70003))
		return
	}
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	sentenceCollectInfo, err := service.GetCollectSentenceService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取句子收藏列表成功",
		"data": sentenceCollectInfo,
	})
}

/*
GetSentenceBySentenceIdHandler
@author: LJR
@Description: 用户点进去具体句子获取听力信息
@param ctx
@Router /api/v1/sentence [get]
*/
func GetSentenceBySentenceIdHandler(ctx *gin.Context) {
	sentence_id := ctx.Query("sentence_id")
	if sentence_id == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(70010))
		return
	}
	num, _ := strconv.Atoi(sentence_id)
	sentenceId := uint64(num)
	sentence, err := service.GetSentenceBySentenceIdService(sentenceId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取句子成功",
		"data": *sentence,
	})
}
