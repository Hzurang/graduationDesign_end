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
GetWordMeaningHandler
@author: LJR
@Description: 获取单词内容API
@param ctx
@Router /api/v1/word/meaning [get]
*/
func GetWordMeaningHandler(ctx *gin.Context) {
	word := ctx.Query("word")
	if word == "" {
		zap.L().Error("GetWordMeaning with invalid param", zap.Error(utils.GetError(40005)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40005))
		return
	}
	wordTranslationInfo, err := service.GetWordMeaningService(word)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取单词内容成功",
		"data": *wordTranslationInfo,
	})
}

/*
GetWordDetailHandler
@author: LJR
@Description: 记单词时获取单词详情
@param ctx
@Router /api/v1/word/memory/meaning [get]
*/
func GetWordDetailHandler(ctx *gin.Context) {
	word_id := ctx.Query("word_id")
	if word_id == "" {
		zap.L().Error("GetWordMeaning with invalid param", zap.Error(utils.GetError(40008)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40008))
		return
	}
	num, _ := strconv.Atoi(word_id)
	wordId := uint64(num)
	wordDetail, err := service.GetWordDetailService(wordId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取单词内容成功",
		"data": *wordDetail,
	})
}

/*
CollectWordHandler
@author: LJR
@Description: 用户收藏单词
@param ctx
@Router /api/v1/word/collection [post]
*/
func CollectWordHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	word_id := ctx.Query("word_id")
	if word_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(word_id)
	wordId := uint64(num)
	err := service.CollectWordService(wordId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "单词收藏成功")
}

/*
CancelCollectWordHandler
@author: LJR
@Description: 用户取消收藏单词
@param ctx
@Router /api/v1/word/cancellation/collection [delete]
*/
func CancelCollectWordHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	word_id := ctx.Query("word_id")
	if word_id == "" || userId == 0 {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(word_id)
	wordId := uint64(num)
	err := service.CancelCollectWordService(wordId, userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "单词取消收藏成功")
}

/*
GetCollectWordListHandler
@author: LJR
@Description: 用户获取单词收藏列表
@param ctx
@Router /api/v1/word/collection/list [get]
*/
func GetCollectWordListHandler(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	if user_id == "" {
		utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, utils.GetError(30019))
		return
	}
	num, _ := strconv.Atoi(user_id)
	userId := uint64(num)
	wordCollectInfo, err := service.GetCollectWordService(userId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取单词收藏列表成功",
		"data": wordCollectInfo,
	})
}

/*
InsertWordHandler
@author: LJR
@Description: 手动添加单词
@param ctx
@Router /admin/v1/word/insertion [post]
*/
func InsertWordHandler(ctx *gin.Context) {
	p := new(model.ParamInsertWord)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("InsertWord with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	fmt.Println(p)
	err := service.InsertWordService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "添加单词成功")
}

/*
UpdateWordHandler
@author: LJR
@Description: 根据单词ID修改单词
@param ctx
@Router /admin/v1/word/modify [put]
*/
func UpdateWordHandler(ctx *gin.Context) {
	p := new(model.ParamWordInfo)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateWord with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 参数缺漏
			utils.ResponseParamError(ctx, utils.FAIL_BUSINESS, nil, err)
			return
		}
		utils.ResponseParamTypeError(ctx, utils.FAIL_BUSINESS, nil, utils.RemoveTopStruct(errs.Translate(config.Trans)))
		return
	}
	err := service.UpdateWordService(p)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "修改单词成功")
}

/*
GetWordHandler
@author: LJR
@Description: 根据单词ID获取单词信息
@param ctx
@Router /admin/v1/word/?word_id [get]
*/
func GetWordHandler(ctx *gin.Context) {
	wordId := ctx.Query("word_id")
	fmt.Println(wordId)
	if wordId == "" {
		zap.L().Error("GetListen with invalid param", zap.Error(utils.GetError(40008)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40008))
		return
	}
	wordInfo, err := service.GetWordService(wordId)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "获取单词成功",
		"data": *wordInfo,
	})
}

/*
DeleteWordHandler
@author: LJR
@Description: 根据单词ID删除单词
@param ctx
@Router /admin/v1/word/delete [delete]
*/
func DeleteWordHandler(ctx *gin.Context) {
	wordId := ctx.Query("word_id")
	if wordId == "" {
		zap.L().Error("DeleteWord with invalid param", zap.Error(utils.GetError(40008)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40008))
		return
	}
	if err := service.DeleteWordService(wordId); err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "删除单词成功")
}

func WordHandler(ctx *gin.Context) {
	justLearnedWords := make([]int64, 0, 3)
	justLearnedWords = append(justLearnedWords, 1, 2, 3)
	ctx.Set("1", 1)
	ctx.Set("justLearnedWords", justLearnedWords)
	c, _ := ctx.Get("justLearnedWords")
	fmt.Println(c)
}

func WordGetHandler(ctx *gin.Context) {
	justLearnedWords, _ := ctx.Get("justLearnedWords")
	anySlice, _ := justLearnedWords.([]interface{}) // 先将 any 转换为 interface{} 切片
	uint64Slice := make([]uint64, len(anySlice))
	for i, x := range anySlice {
		u64, _ := x.(uint64) // 将 interface{} 转换为 uint64
		uint64Slice[i] = u64
	}
	c, _ := ctx.Get("1")
	fmt.Println(uint64Slice)
	fmt.Println(c)
}

/*
DetermineVocabularyHandler
@author: LJR
@Description: 判断用户是否背完这本书 传 word_type 背完返回1，没有背完返回0
@param ctx
@Router /api/v1/word/vocabulary/termination [get]
*/
func DetermineVocabularyHandler(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	word_type := ctx.Query("word_type")
	if word_type == "" {
		zap.L().Error("DetermineVocabulary with invalid param", zap.Error(utils.GetError(40018)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40018))
		return
	}
	isEnd, err := service.DetermineVocabularyService(userId, word_type)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	switch isEnd {
	case 1:
		ctx.JSON(http.StatusOK, gin.H{
			"code": utils.SUCCESS,
			"msg":  "您的词书背完啦",
			"data": 1,
		})
	case 0:
		ctx.JSON(http.StatusOK, gin.H{
			"code": utils.SUCCESS,
			"msg":  "您的词书还未背完",
			"data": 0,
		})
	}
}

/*
GetWordAllHandler
@author: LJR
@Description: 根据单词类型和页数等加载单词列表
@param ctx
@Router /admin/v1/words [get]
*/
func GetWordAllHandler(ctx *gin.Context) {
	word_type := ctx.Query("wordtype")
	page_num := ctx.Query("pagenum")
	page_size := ctx.Query("pagesize")
	if page_num == "" || page_size == "" {
		zap.L().Error("GetWordAll with invalid param", zap.Error(utils.GetError(10014)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(10014))
		return
	}
	wordList, total, err := service.GetWordAllService(word_type, page_num, page_size)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"wordList": wordList,
		"total":    total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "单词列表获取成功",
		"data": data,
	})
}

/*
GetWordByWordHandler
@author: LJR
@Description: 输入单词进行查找
@param ctx
@Router /admin/v1/word/signal [get]
*/
func GetWordByWordHandler(ctx *gin.Context) {
	word := ctx.Query("word")
	if word == "" {
		zap.L().Error("GetWordByWord with invalid param", zap.Error(utils.GetError(40005)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40005))
		return
	}
	wordList, total, err := service.GetWordByWordService(word)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	data := map[string]interface{}{
		"wordList": wordList,
		"total":    total,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "单词获取成功",
		"data": data,
	})
}

func GetWordListHandler(ctx *gin.Context) {
	wordType := ctx.Query("word_type")
	if wordType == "" {
		zap.L().Error("GetWordList with invalid param", zap.Error(utils.GetError(40022)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40022))
		return
	}
	wordList, err := service.GetWordListService(wordType)
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "单词数据加载成功",
		"data": wordList,
	})
}
