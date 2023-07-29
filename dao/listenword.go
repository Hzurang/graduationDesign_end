package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateListenWord(listenWord *model.ListenWord) (err error) {
	err = global.Db.Create(listenWord).Error
	if err != nil {
		zap.L().Error("CreateListenWord", zap.Error(err))
		return utils.GetStrAndError("单词", 10010)
	}
	return nil
}

func DeleteListenWordByListenId(listenId uint64) (err error) {
	listenWord := &model.ListenWord{ListenId: listenId}
	global.Db.Model(listenWord).Where("listen_id = ?", listenWord.ListenId).Update("delete_isok", 1)
	num := global.Db.Where("listen_id = ?", listenWord.ListenId).Delete(listenWord).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteListenWordByListenId", zap.Error(err))
		return utils.GetError(40003)
	}
	return nil
}

func GetListenWordListByListenId(listenId uint64) ([]model.ListenWordInfo, error) {
	var listenWordList []model.ListenWord
	err := global.Db.Where("listen_id = ?", listenId).Find(&listenWordList).Error
	if err != nil {
		return nil, utils.GetError(40006)
	}
	listenWordInfoList := make([]model.ListenWordInfo, 0, len(listenWordList))
	for _, listenWord := range listenWordList {
		listenWordInfo := model.ListenWordInfo{
			WordId:       listenWord.WordId,
			Word:         listenWord.Word,
			WordPhonetic: listenWord.WordPhonetic,
			WordMeaning:  listenWord.WordMeaning,
			WordMusic:    listenWord.WordMusic,
			WordNum:      listenWord.WordNum,
			ListenId:     listenWord.ListenId,
		}
		listenWordInfoList = append(listenWordInfoList, listenWordInfo)
	}
	return listenWordInfoList, nil
}
