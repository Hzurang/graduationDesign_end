package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateSentence(sentence *model.Sentence) (err error) {
	err = global.Db.Create(sentence).Error
	if err != nil {
		zap.L().Error("CreateSentence", zap.Error(err))
		return utils.GetStrAndError("句子", 10010)
	}
	return
}

func GetSentenceByContent(sentenceContent string) (sentence *model.Sentence, err error) {
	sentence = &model.Sentence{SentenceContent: sentenceContent}
	if err := global.Db.Where("sentence_content = ?", sentenceContent).First(sentence).Error; err == nil {
		return sentence, utils.GetError(70001)
	}
	return nil, nil
}

func GetSentenceBySentenceId(sentenceId uint64) (sentence *model.Sentence, err error) {
	sentence = &model.Sentence{SentenceId: sentenceId}
	if err := global.Db.Where("sentence_id = ?", sentenceId).First(sentence).Error; err != nil {
		zap.L().Error("GetSentenceBySentenceId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(70004)
		}
		return nil, utils.GetStrAndError("句子", 10009)
	}
	return sentence, nil
}

func UpdateSentenceCollectNumBySentenceId(sentence *model.Sentence) (err error) {
	res := global.Db.Model(sentence).Where("sentence_id = ?", sentence.SentenceId).Update("sentence_collect_num", sentence.SentenceCollectNum)
	if res.Error != nil {
		zap.L().Error("UpdateSentenceBySentenceId", zap.Error(err))
		return utils.GetError(70006)
	}
	return nil
}

func GetSentenceByPublishAt(publishAt string) (sentence *model.Sentence, err error) {
	sentence = &model.Sentence{PublishAt: publishAt}
	if err := global.Db.Where("publish_at = ?", publishAt).First(sentence).Error; err != nil {
		zap.L().Error("GetSentenceByPublishAt", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(70004)
		}
		return nil, utils.GetStrAndError("句子", 10009)
	}
	return sentence, nil
}
