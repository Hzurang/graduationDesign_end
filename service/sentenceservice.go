package service

import (
	"encoding/json"
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const url = "https://sentence.iciba.com/index.php?c=dailysentence&m=getdetail&title="
const layout = "2006-01-02"

/*
DailySentenceApiTotalService
@author: LJR
@Description: 每日一句全爬取业务
@return err
*/
func DailySentenceApiTotalService() (err error) {
	now := time.Now()
	end, _ := time.Parse(layout, "2019-11-26")
	// 遍历2019-11-26到现在的日子
	for end.Before(now) || end.Equal(now) {
		sentence := new(model.Sentence)
		title := end.Format(layout)
		resp, err := http.Get(url + title)
		if err != nil {
			zap.L().Error(fmt.Sprintf("failed to get sentence detail for title %s: ", title), zap.Error(err))
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			var sentenceJSON model.ParamSentenceInfo
			err = json.NewDecoder(resp.Body).Decode(&sentenceJSON)
			if err != nil {
				zap.L().Error(fmt.Sprintf("failed to decode sentence detail for title %s: ", title), zap.Error(err))
				continue
			}
			if sentenceJSON.Errmsg == "success" {
				_, err = dao.GetSentenceByContent(sentenceJSON.Content)
				if err != nil {
					zap.L().Error("DailySentenceApiTotalService", zap.Error(err))
					continue
				}
				sentence.SentenceId, _ = config.GenID()
				sentence.PublishAt = sentenceJSON.Title
				sentence.SentenceContent = sentenceJSON.Content
				sentence.SentenceNote = sentenceJSON.Note
				sentence.SentenceTranslation = sentenceJSON.Translation
				sentence.SentencePicture = sentenceJSON.Picture
				sentence.SentenceAudioPath = sentenceJSON.AudioPath
				sentence.SentenceCollectNum = 0
				err = dao.CreateSentence(sentence)
				if err != nil {
					zap.L().Error("DailySentenceApiTotalService", zap.Error(err))
					continue
				}
			}
		}
		end = end.AddDate(0, 0, 1)
		time.Sleep(200 * time.Millisecond) // 加入 200 毫秒的延迟，防止请求过快被封 IP
	}
	return nil
}

/*
DailySentenceEveryDayService
@author: LJR
@Description: 每日一句按每天爬取业务
@return sentenceInfo
@return err
*/
func DailySentenceEveryDayService() (sentenceInfo *model.SentenceInfo, err error) {
	sentence := new(model.Sentence)
	sentenceInfo = new(model.SentenceInfo)
	now := time.Now()
	now = now.AddDate(0, 0, 1)
	title := now.Format(layout)
	resp, err := http.Get(url + title)
	if err != nil {
		zap.L().Error(fmt.Sprintf("failed to get sentence detail for title %s: ", title), zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var sentenceJSON model.ParamSentenceInfo
		err = json.NewDecoder(resp.Body).Decode(&sentenceJSON)
		if err != nil {
			zap.L().Error(fmt.Sprintf("failed to decode sentence detail for title %s: ", title), zap.Error(err))
			return nil, err
		}
		if sentenceJSON.Errmsg == "success" {
			_, err = dao.GetSentenceByContent(sentenceJSON.Content)
			if err != nil {
				zap.L().Error("DailySentenceApiTotalService", zap.Error(err))
				return nil, err
			}
			sentence.SentenceId, _ = config.GenID()
			sentence.PublishAt = sentenceJSON.Title
			sentence.SentenceContent = sentenceJSON.Content
			sentence.SentenceNote = sentenceJSON.Note
			sentence.SentenceTranslation = sentenceJSON.Translation
			sentence.SentencePicture = sentenceJSON.Picture
			sentence.SentenceAudioPath = sentenceJSON.AudioPath
			sentence.SentenceCollectNum = 0
			err = dao.CreateSentence(sentence)
			if err != nil {
				zap.L().Error("DailySentenceApiTotalService", zap.Error(err))
				return nil, err
			}
			sentenceInfo.SentenceId = sentence.SentenceId
			sentenceInfo.PublishAt = sentence.PublishAt
			sentenceInfo.SentenceContent = sentence.SentenceContent
			sentenceInfo.SentenceNote = sentence.SentenceNote
			sentenceInfo.SentenceTranslation = sentence.SentenceTranslation
			sentenceInfo.SentencePicture = sentence.SentencePicture
			sentenceInfo.SentenceAudioPath = sentence.SentenceAudioPath
			sentenceInfo.SentenceCollectNum = sentence.SentenceCollectNum
			return sentenceInfo, nil
		}
	}
	return sentenceInfo, nil
}

/*
CollectSentenceService
@author: LJR
@Description: 用户收藏句子业务逻辑
@param sentenceId
@param userId
@return err
*/
func CollectSentenceService(sentenceId uint64, userId uint64) (err error) {
	_, err = dao.GetSentenceBySentenceId(sentenceId)
	if err != nil {
		return utils.GetError(70005)
	}
	sentenceCollect, err := dao.GetSoftDeleteCollectSentenceBySentenceIdAndUserId(sentenceId, userId)
	sentence, _ := dao.GetSentenceBySentenceId(sentenceId)
	if err == nil {
		// 软删除更新
		if sentenceCollect.DeleteIsOk == 1 {
			sentence.SentenceCollectNum = sentence.SentenceCollectNum + 1
			_ = dao.UpdateSentenceCollectNumBySentenceId(sentence)
			err = dao.UpdateSoftDeleteCollectSentence(sentenceCollect)
			if err != nil {
				return err
			}
			return nil
		} else if sentenceCollect.DeleteIsOk == 0 {
			return utils.GetError(70008)
		}
	}
	sentenceCollect = &model.SentenceCollect{
		SentenceId: sentenceId,
		UserId:     userId,
	}
	err = dao.CreateCollectSentence(sentenceCollect)
	if err != nil {
		return err
	}
	sentence.SentenceCollectNum = sentence.SentenceCollectNum + 1
	_ = dao.UpdateSentenceCollectNumBySentenceId(sentence)
	return nil
}

/*
CancelCollectSentenceService
@author: LJR
@Description: 用户取消句子收藏业务逻辑
@param sentenceId
@param userId
@return err
*/
func CancelCollectSentenceService(sentenceId uint64, userId uint64) (err error) {
	_, err = dao.GetSentenceBySentenceId(sentenceId)
	if err != nil {
		return utils.GetError(50009)
	}
	err = dao.DeleteCollectSentenceBySentenceIdAndUserId(sentenceId, userId)
	if err != nil {
		return err
	}
	sentence, _ := dao.GetSentenceBySentenceId(sentenceId)
	sentence.SentenceCollectNum = sentence.SentenceCollectNum - 1
	_ = dao.UpdateSentenceCollectNumBySentenceId(sentence)
	return nil
}

/*
DailySentenceService
@author: LJR
@Description: 用户获取每日一句业务逻辑
@return sentenceInfo
@return err
*/
func DailySentenceService() (sentenceInfo *model.SentenceInfo, err error) {
	now := time.Now()
	publishAt := now.Format(layout)
	sentence, err := dao.GetSentenceByPublishAt(publishAt)
	if err != nil {
		// 没有就直接爬虫
		sentenceInfo, err = DailySentenceEveryDayService()
		if err != nil {
			return nil, err
		}
		return sentenceInfo, nil
	}
	sentenceInfo = new(model.SentenceInfo)
	sentenceInfo.SentenceId = sentence.SentenceId
	sentenceInfo.PublishAt = sentence.PublishAt
	sentenceInfo.SentenceContent = sentence.SentenceContent
	sentenceInfo.SentenceNote = sentence.SentenceNote
	sentenceInfo.SentenceTranslation = sentence.SentenceTranslation
	sentenceInfo.SentencePicture = sentence.SentencePicture
	sentenceInfo.SentenceAudioPath = sentence.SentenceAudioPath
	sentenceInfo.SentenceCollectNum = sentence.SentenceCollectNum
	return sentenceInfo, nil
}

/*
GetCollectSentenceService
@author: LJR
@Description: 用户获取句子收藏列表业务逻辑
@param userId
@return sentenceCollectInfoList
@return err
*/
func GetCollectSentenceService(userId uint64) (sentenceCollectInfoList []model.SentenceCollectInfo, err error) {
	sentenceList, err := dao.GetSentenceCollectListByUserId(userId)
	if err != nil {
		return nil, err
	}
	sentenceCollectInfoList = make([]model.SentenceCollectInfo, 0, len(sentenceList))
	for _, sentence := range sentenceList {
		sentenceCollectInfo := model.SentenceCollectInfo{
			SentenceId:         sentence.SentenceId,
			PublishAt:          sentence.PublishAt,
			SentenceContent:    sentence.SentenceContent,
			SentenceCollectNum: sentence.SentenceCollectNum,
		}
		sentenceCollectInfoList = append(sentenceCollectInfoList, sentenceCollectInfo)
	}
	return sentenceCollectInfoList, nil
}

/*
GetSentenceBySentenceIdService
@author: LJR
@Description: 用户点进去具体句子获取听力信息业务逻辑
@param sentenceId
@return sentence
@return err
*/
func GetSentenceBySentenceIdService(sentenceId uint64) (sentence *model.Sentence, err error) {
	sentence, err = dao.GetSentenceBySentenceId(sentenceId)
	if err != nil {
		return nil, err
	}
	return sentence, nil
}
