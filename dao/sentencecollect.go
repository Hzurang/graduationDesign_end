package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateCollectSentence(sentenceCollect *model.SentenceCollect) (err error) {
	err = global.Db.Create(sentenceCollect).Error
	if err != nil {
		zap.L().Error("CreateCollectSentence", zap.Error(err))
		return utils.GetError(70007)
	}
	return nil
}

func GetSoftDeleteCollectSentenceBySentenceIdAndUserId(sentenceId uint64, userId uint64) (sentenceCollect *model.SentenceCollect, err error) {
	sentenceCollect = &model.SentenceCollect{SentenceId: sentenceId, UserId: userId}
	err = global.Db.Unscoped().Where("sentence_id = ? AND user_id = ?", sentenceId, userId).First(sentenceCollect).Error
	if err == nil {
		return sentenceCollect, nil
	}
	return nil, err
}

func UpdateSoftDeleteCollectSentence(sentenceCollect *model.SentenceCollect) (err error) {
	err = global.Db.Unscoped().Model(sentenceCollect).Where("sentence_id = ? AND user_id = ?", sentenceCollect.SentenceId, sentenceCollect.UserId).Updates(map[string]interface{}{
		"deleted_at":  nil,
		"delete_isok": 0,
	}).Error
	if err != nil {
		zap.L().Error("UpdateSoftDeleteCollectSentence", zap.Error(err))
		return utils.GetError(70007)
	}
	return nil
}

func DeleteCollectSentenceBySentenceIdAndUserId(sentenceId uint64, userId uint64) (err error) {
	sentenceCollect := &model.SentenceCollect{SentenceId: sentenceId, UserId: userId}
	global.Db.Model(sentenceCollect).Where("sentence_id = ? AND user_id = ?", sentenceCollect.SentenceId, sentenceCollect.UserId).Update("delete_isok", 1)
	num := global.Db.Where("sentence_id = ? AND user_id = ?", sentenceCollect.SentenceId, sentenceCollect.UserId).Delete(sentenceCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectSentenceBySentenceIdAndUserId", zap.Error(err))
		return utils.GetError(70009)
	}
	return nil
}

func GetSentenceCollectListByUserId(userId uint64) (sentenceList []model.Sentence, err error) {
	sentenceList = make([]model.Sentence, 0, 60)
	err = global.Db.Select("sentence_id, publish_at, sentence_content, sentence_collect_num").
		Joins("INNER JOIN sentence_collect ON sentence.sentence_id = sentence_collect.sentence_id").
		Where("sentence_collect.user_id = ? AND sentence_collect.delete_isok = ? AND sentence.delete_isok = ?", userId, 0, 0).
		Find(&sentenceList).Error
	if err != nil {
		return nil, utils.GetError(70011)
	}
	return sentenceList, nil
}
