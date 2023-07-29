package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func CreateCollectListen(listenCollect *model.ListenCollect) (err error) {
	err = global.Db.Create(listenCollect).Error
	if err != nil {
		zap.L().Error("CreateCollectListen", zap.Error(err))
		return utils.GetError(60012)
	}
	return nil
}

func DeleteCollectListenByListenId(listenId uint64) (err error) {
	listenCollect := &model.ListenCollect{ListenId: listenId}
	global.Db.Model(listenCollect).Where("listen_id = ?", listenCollect.ListenId).Update("delete_isok", 1)
	num := global.Db.Where("listen_id = ?", listenCollect.ListenId).Delete(listenCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectListenByListenId", zap.Error(err))
		return utils.GetError(60004)
	}
	return nil
}

func GetSoftDeleteCollectListenByListenIdAndUserId(listenId uint64, userId uint64) (listenCollect *model.ListenCollect, err error) {
	listenCollect = &model.ListenCollect{ListenId: listenId, UserId: userId}
	err = global.Db.Unscoped().Where("listen_id = ? AND user_id = ?", listenId, userId).First(listenCollect).Error
	if err == nil {
		return listenCollect, nil
	}
	return nil, err
}

func UpdateSoftDeleteCollectListen(listenCollect *model.ListenCollect) (err error) {
	err = global.Db.Unscoped().Model(listenCollect).Where("listen_id = ? AND user_id = ?", listenCollect.ListenId, listenCollect.UserId).Updates(map[string]interface{}{
		"deleted_at":  nil,
		"delete_isok": 0,
	}).Error
	if err != nil {
		zap.L().Error("UpdateSoftDeleteCollectListen", zap.Error(err))
		return utils.GetError(60012)
	}
	return nil
}

func DeleteCollectListenByListenIdAndUserId(listenId uint64, userId uint64) (err error) {
	listenCollect := &model.ListenCollect{ListenId: listenId, UserId: userId}
	global.Db.Model(listenCollect).Where("listen_id = ? AND user_id = ?", listenCollect.ListenId, listenCollect.UserId).Update("delete_isok", 1)
	num := global.Db.Where("listen_id = ? AND user_id = ?", listenCollect.ListenId, listenCollect.UserId).Delete(listenCollect).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteCollectListenByListenIdAndUserId", zap.Error(err))
		return utils.GetError(60004)
	}
	return nil
}

func GetListenCollectListByUserId(userId uint64) (listenList []model.Listen, err error) {
	listenList = make([]model.Listen, 0, 60)
	err = global.Db.Select("listen_id, listen_title, listen_source, listen_editor, publish_at, listen_second_type, listen_collect_num").
		Joins("INNER JOIN listen_collect ON listen.listen_id = listen_collect.listen_id").
		Where("listen_collect.user_id = ? AND listen_collect.delete_isok = ? AND listen.delete_isok = ?", userId, 0, 0).
		Find(&listenList).Error
	if err != nil {
		return nil, utils.GetError(60015)
	}
	return listenList, nil
}
