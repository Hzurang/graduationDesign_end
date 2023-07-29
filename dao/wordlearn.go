package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
)

func GetWordLearnCountByUserIdAndWordType(userId uint64, wordType int8) (count int64, err error) {
	wordLearn := &model.WordLearn{UserId: userId}
	err = global.Db.Model(wordLearn).Where("user_id = ? AND word_type = ? AND deep_master_times <> ?", userId, wordType, 3).Count(&count).Error
	if err != nil {
		return 0, utils.GetError(40019)
	}
	return count, nil
}

func UpdatesResetWordLearnByUserIdAndWordType(userId uint64, wordType int8) (err error) {
	wordLearn := &model.WordLearn{}
	err = global.Db.Model(wordLearn).Where("user_id = ? AND word_type = ?", userId, wordType).Updates(map[string]interface{}{
		"just_learned":     0,
		"is_need_learned":  0,
		"is_learned":       0,
		"last_master_time": 0,
		"last_review_time": 0,
		"master_degree":    0,
		"DeepMasterTimes":  0,
		"need_learn_date":  0,
		"need_review_date": 0,
	}).Error
	if err != nil {
		return utils.GetError(40020)
	}
	return nil
}
