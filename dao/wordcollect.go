package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateCollectWord(wordCollect *model.WordCollect) (err error) {
	err = global.Db.Create(wordCollect).Error
	if err != nil {
		zap.L().Error("CreateCollectWord", zap.Error(err))
		return utils.GetError(40011)
	}
	return nil
}

func GetSoftDeleteCollectWordByWordIdAndUserId(wordId uint64, userId uint64) (wordCollect *model.WordCollect, err error) {
	wordCollect = &model.WordCollect{WordId: wordId, UserId: userId}
	err = global.Db.Unscoped().Where("word_id = ? AND user_id = ?", wordId, userId).First(wordCollect).Error
	if err == nil {
		return wordCollect, nil
	}
	return nil, err
}

func UpdateSoftDeleteCollectWord(wordCollect *model.WordCollect) (err error) {
	err = global.Db.Unscoped().Model(wordCollect).Where("word_id = ? AND user_id = ?", wordCollect.WordId, wordCollect.UserId).Updates(map[string]interface{}{
		"deleted_at":  nil,
		"delete_isok": 0,
	}).Error
	if err != nil {
		zap.L().Error("UpdateSoftDeleteCollectWord", zap.Error(err))
		return utils.GetError(40011)
	}
	return nil
}

func DeleteCollectWordByWordIdAndUserId(wordId uint64, userId uint64) (err error) {
	wordCollect := &model.WordCollect{WordId: wordId, UserId: userId}
	global.Db.Model(wordCollect).Where("word_id = ? AND user_id = ?", wordCollect.WordId, wordCollect.UserId).Update("delete_isok", 1)
	num := global.Db.Where("word_id = ? AND user_id = ?", wordCollect.WordId, wordCollect.UserId).Delete(wordCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectWordByWordIdAndUserId", zap.Error(err))
		return utils.GetError(40013)
	}
	return nil
}

func GetWordCollectListByUserId(userId uint64) (wordList []model.Word, err error) {
	wordList = make([]model.Word, 0, 60)
	err = global.Db.Select("word_id, word, word_meaning, word_type").
		Joins("INNER JOIN word_collect ON word.word_id = word_collect.word_id").
		Where("word_collect.user_id = ? AND word_collect.delete_isok = ? AND word.delete_isok = ?", userId, 0, 0).
		Find(&wordList).Error
	if err != nil {
		return nil, utils.GetError(40014)
	}
	return wordList, nil
}

func DeleteCollectWordByWordId(wordId uint64) (err error) {
	wordCollect := &model.WordCollect{WordId: wordId}
	global.Db.Model(wordCollect).Where("word_id = ?", wordCollect.WordId).Update("delete_isok", 1)
	num := global.Db.Where("word_id = ?", wordCollect.WordId).Delete(wordCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectWordByWordId", zap.Error(err))
		return utils.GetError(40013)
	}
	return nil
}
