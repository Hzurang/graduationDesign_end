package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateEssayWord(essayWord *model.EssayWord) (err error) {
	err = global.Db.Create(essayWord).Error
	if err != nil {
		zap.L().Error("CreateEssayWord", zap.Error(err))
		return utils.GetStrAndError("单词", 10010)
	}
	return
}

func DeleteEssayWordByEssayId(essayId uint64) (err error) {
	essayWord := &model.EssayWord{EssayId: essayId}
	global.Db.Model(essayWord).Where("essay_id = ?", essayWord.EssayId).Update("delete_isok", 1)
	num := global.Db.Where("essay_id = ?", essayWord.EssayId).Delete(essayWord).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteEssayWordByEssayId", zap.Error(err))
		return utils.GetError(40001)
	}
	return nil
}

func GetEssayWordListByEssayId(essayId uint64) ([]model.EssayWordInfo, error) {
	var essayWordList []model.EssayWord
	err := global.Db.Where("essay_id = ?", essayId).Find(&essayWordList).Error
	if err != nil {
		return nil, utils.GetError(40002)
	}
	essayWordInfoList := make([]model.EssayWordInfo, 0, len(essayWordList))
	for _, essayWord := range essayWordList {
		essayWordInfo := model.EssayWordInfo{
			WordId:       essayWord.WordId,
			Word:         essayWord.Word,
			WordNum:      essayWord.WordNum,
			WordMusic:    essayWord.WordMusic,
			EssayId:      essayWord.EssayId,
			WordMeaning:  essayWord.WordMeaning,
			WordSentence: essayWord.WordSentence,
		}
		essayWordInfoList = append(essayWordInfoList, essayWordInfo)
	}
	return essayWordInfoList, nil
}
