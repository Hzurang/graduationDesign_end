package v1

import (
	"fmt"
	"ginStudy/service"
	"ginStudy/utils"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"net/http"
)

/*
EssaySpiderTotalHandler
@author: LJR
@Description: 文章通过爬虫进行数据更新 novel 1 love 2 essays 3
@param ctx
@Router /admin/v1/essay/spider/all [post]
*/
func EssaySpiderTotalHandler(ctx *gin.Context) {
	essayType := ctx.Query("essay_type")
	fmt.Println(essayType)
	if essayType == "" {
		zap.L().Error("ListenSpiderTotal with invalid param", zap.Error(utils.GetError(50012)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(50012))
		return
	}
	c := colly.NewContext()
	service.EssayPageNumSpiderService(c, essayType)
	service.EssayFirstSpiderService(c, essayType)
	service.EssaySpiderService(c, essayType)
	//_ = dao.DeleteEssayCache(essayType)
	//err := dao.SetEssayCache(essayType)
	//if err != nil {
	//	utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "文章数据更新成功*_*",
		"data": nil,
	})
}

/*
ListenSpiderTotalHandler
@author: LJR
@Description: 听力通过爬虫进行数据更新
@param ctx
@Router /admin/v1/listen/spider/all [post]
*/
func ListenSpiderTotalHandler(ctx *gin.Context) {
	listenType := ctx.Query("listen_type")
	if listenType == "" {
		zap.L().Error("ListenSpiderTotal with invalid param", zap.Error(utils.GetError(60007)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(60007))
		return
	}
	c := colly.NewContext()
	totalPageNum := service.ListenPageNumSpiderService(listenType)
	service.ListenUrlSpiderService(c, listenType, totalPageNum, false)
	cnt1, cnt2 := service.ListenSpiderService(c, listenType)
	m := map[string]interface{}{
		"listenNumMsg":     fmt.Sprintf("听力新增%d条", cnt1),
		"listenWordNumMsg": fmt.Sprintf("听力单词新增%d条", cnt2),
	}
	//_ = dao.DeleteListenCache(listenType)
	//err := dao.SetListenCache(listenType)
	//if err != nil {
	//	utils.ResponseError(ctx, utils.FAIL_ENTITY, nil, err.Error())
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  "听力数据更新成功*_*",
		"data": m,
	})
}

/*
WordSpiderTotalHandler
@author: LJR
@Description: 单词通过爬虫进行数据更新 CET4 CET6 TEM4 TEM8 KAOYAN GRE TOEFL IELTS
@param ctx
@Router /admin/v1/word/spider/all [post]
*/
func WordSpiderTotalHandler(ctx *gin.Context) {
	wordType := ctx.Query("word_type")
	if wordType == "" {
		zap.L().Error("WordSpiderTotal with invalid param", zap.Error(utils.GetError(40007)))
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(40007))
		return
	}
	c := colly.NewContext()
	totalPageNum := service.WordPageNumSpiderService(wordType)
	service.WordUrlSpiderService(c, wordType, totalPageNum)
	wordList := service.WordSpiderService(c, wordType)
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.SUCCESS,
		"msg":  fmt.Sprintf("单词数据更新成功*_*，共%d条", len(wordList)),
		"data": wordList,
	})
}

/*
DailySentenceTotalSpiderHandler
@author: LJR
@Description: 管理员获取从2019-11-26以来所有的每日一句
@param ctx
@Router /admin/v1/sentence/spider/all [post]
*/
func DailySentenceTotalSpiderHandler(ctx *gin.Context) {
	err := service.DailySentenceApiTotalService()
	if err != nil {
		utils.ResponseError(ctx, utils.FAIL_BUSINESS, nil, utils.GetCodeMsg(70002))
		return
	}
	utils.ResponseGin(ctx, http.StatusOK, utils.SUCCESS, nil, "操作成功")
}
