package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateCollectEssay(essayCollect *model.EssayCollect) (err error) {
	err = global.Db.Create(essayCollect).Error
	if err != nil {
		zap.L().Error("CreateCollectEssay", zap.Error(err))
		return utils.GetError(50007)
	}
	return nil
}

func GetCollectEssayByEssayIdAndUserId(essayId uint64, userId uint64) (essayCollect *model.EssayCollect, err error) {
	essayCollect = &model.EssayCollect{EssayId: essayId, UserId: userId}
	err = global.Db.Where("essay_id = ? AND user_id = ?", essayId, userId).First(essayCollect).Error
	if err == nil {
		return essayCollect, utils.GetError(50006)
	}
	return nil, nil
}

func GetSoftDeleteCollectEssayByEssayIdAndUserId(essayId uint64, userId uint64) (essayCollect *model.EssayCollect, err error) {
	essayCollect = &model.EssayCollect{EssayId: essayId, UserId: userId}
	err = global.Db.Unscoped().Where("essay_id = ? AND user_id = ?", essayId, userId).First(essayCollect).Error
	if err == nil {
		return essayCollect, nil
	}
	return nil, err
}

func UpdateSoftDeleteCollectEssay(essayCollect *model.EssayCollect) (err error) {
	err = global.Db.Unscoped().Model(essayCollect).Where("essay_id = ? AND user_id = ?", essayCollect.EssayId, essayCollect.UserId).Updates(map[string]interface{}{
		"deleted_at":  nil,
		"delete_isok": 0,
	}).Error
	if err != nil {
		zap.L().Error("UpdateSoftDeleteCollectEssay", zap.Error(err))
		return utils.GetError(50007)
	}
	return nil
}

func DeleteCollectEssayByEssayIdAndUserId(essayId uint64, userId uint64) (err error) {
	essayCollect := &model.EssayCollect{EssayId: essayId, UserId: userId}
	global.Db.Model(essayCollect).Where("essay_id = ? AND user_id = ?", essayCollect.EssayId, essayCollect.UserId).Update("delete_isok", 1)
	num := global.Db.Where("essay_id = ? AND user_id = ?", essayCollect.EssayId, essayCollect.UserId).Delete(essayCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectEssayByEssayIdAndUserId", zap.Error(err))
		return utils.GetError(50008)
	}
	return nil
}

func DeleteCollectEssayByEssayId(essayId uint64) (err error) {
	essayCollect := &model.EssayCollect{EssayId: essayId}
	global.Db.Model(essayCollect).Where("essay_id = ?", essayCollect.EssayId).Update("delete_isok", 1)
	num := global.Db.Where("essay_id = ?", essayCollect.EssayId).Delete(essayCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectEssayByEssayIdAndUserId", zap.Error(err))
		return utils.GetError(50008)
	}
	return nil
}

func GetEssayCollectListByUserId(userId uint64) (essayList []model.Essay, err error) {
	essayList = make([]model.Essay, 0, 60)
	err = global.Db.Select("essay_id, essay_title, essay_author, publish_at, essay_type, essay_collect_num").
		Joins("INNER JOIN essay_collect ON essay.essay_id = essay_collect.essay_id").
		Where("essay_collect.user_id = ? AND essay_collect.delete_isok = ? AND essay.delete_isok = ?", userId, 0, 0).
		Find(&essayList).Error
	if err != nil {
		return nil, utils.GetError(50015)
	}
	return essayList, nil
}
