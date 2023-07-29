package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateListen(listen *model.Listen) (err error) {
	err = global.Db.Create(listen).Error
	if err != nil {
		zap.L().Error("CreateListen", zap.Error(err))
		return utils.GetStrAndError("听力", 10010)
	}
	return nil
}

func GetListenByTitle(listenTitle string) (listen *model.Listen, err error) {
	listen = &model.Listen{ListenTitle: listenTitle}
	if err := global.Db.Where("listen_title = ?", listenTitle).First(listen).Error; err == nil {
		return listen, utils.GetError(60001)
	}
	return nil, nil
}

func DeleteListenByListenId(listenId uint64) (err error) {
	listen := &model.Listen{ListenId: listenId}
	global.Db.Model(listen).Where("listen_id = ?", listenId).Update("delete_isok", 1)
	num := global.Db.Where("listen_id = ?", listenId).Delete(listen).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteListenByListenId", zap.Error(err))
		return utils.GetError(60003)
	}
	return nil
}

func GetListenByListenId(listenId uint64) (listen *model.Listen, err error) {
	listen = &model.Listen{ListenId: listenId}
	if err := global.Db.Where("listen_id = ?", listen.ListenId).First(listen).Error; err != nil {
		zap.L().Error("GetListenByListenId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(60005)
		}
		return nil, utils.GetStrAndError("听力", 10009)
	}
	return listen, nil
}

func UpdateListenByListenId(listen *model.Listen) (err error) {
	res := global.Db.Model(listen).Where("listen_id = ?", listen.ListenId).Updates(map[string]interface{}{
		"listen_title":       listen.ListenTitle,
		"listen_source":      listen.ListenSource,
		"listen_editor":      listen.ListenEditor,
		"publish_at":         listen.PublishAt,
		"listen_content":     listen.ListenContent,
		"listen_media_path":  listen.ListenMediaPath,
		"listen_mp_3_path":   listen.ListenMp3Path,
		"listen_type":        listen.ListenType,
		"listen_second_type": listen.ListenSecondType,
		"listen_collect_num": listen.ListenCollectNum,
	})
	if res.Error != nil {
		zap.L().Error("UpdateListenByListenId", zap.Error(err))
		return utils.GetError(60006)
	}
	return nil
}

func GetListenToListenInfoByListenType(listenType int8) ([]model.ListenCacheInfo, error) {
	var listenList []model.Listen
	err := global.Db.Where("listen_type = ?", listenType).Find(&listenList).Error
	if err != nil {
		return nil, utils.GetError(60008)
	}
	listenInfoList := make([]model.ListenCacheInfo, 0, len(listenList))
	for _, listen := range listenList {
		listenInfo := model.ListenCacheInfo{
			ListenId:         listen.ListenId,
			ListenTitle:      listen.ListenTitle,
			ListenSource:     listen.ListenSource,
			ListenEditor:     listen.ListenEditor,
			PublishAt:        listen.PublishAt,
			ListenSecondType: listen.ListenSecondType,
			ListenCollectNum: listen.ListenCollectNum,
		}
		listenInfoList = append(listenInfoList, listenInfo)
	}
	return listenInfoList, nil
}

func GetListenPage(listenType int8, pageNum int, pageSize int) (listenList []model.Listen, total int64, err error) {
	listenList = make([]model.Listen, 0, 50)
	offset := (pageNum - 1) * pageSize
	if listenType != 0 {
		err = global.Db.Select("created_at, listen_id, listen_title, listen_collect_num, listen_type").Limit(pageSize).Offset(offset).Where("listen_type = ?", listenType).Order("created_at desc").Find(&listenList).Error
		global.Db.Model(&listenList).Where("listen_type = ?", listenType).Count(&total)
		if err != nil {
			return nil, 0, utils.GetError(60016)
		}
		return listenList, total, nil
	}
	err = global.Db.Select("created_at, listen_id, listen_title, listen_collect_num, listen_type").Limit(pageSize).Offset(offset).Order("created_at desc").Find(&listenList).Error
	global.Db.Model(&listenList).Count(&total)
	if err != nil {
		return nil, 0, utils.GetError(60016)
	}
	return listenList, total, nil
}
