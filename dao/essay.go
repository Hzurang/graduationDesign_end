package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateEssay(essay *model.Essay) (err error) {
	err = global.Db.Create(essay).Error
	if err != nil {
		zap.L().Error("CreateEssay", zap.Error(err))
		return utils.GetStrAndError("文章", 10010)
	}
	return
}

func GetEssayByTitleAndContent(essayTitle string, essayContent string) (essay *model.Essay, err error) {
	essay = &model.Essay{EssayTitle: essayTitle, EssayContent: essayContent}
	if err := global.Db.Where("essay_title = ? AND essay_content = ?", essayTitle, essayContent).First(essay).Error; err == nil {
		return essay, utils.GetError(50005)
	}
	return nil, nil
}

func DeleteEssayByEssayId(essayId uint64) (err error) {
	essay := &model.Essay{EssayId: essayId}
	global.Db.Model(essay).Where("essay_id = ?", essayId).Update("delete_isok", 1)
	num := global.Db.Where("essay_id = ?", essayId).Delete(essay).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteEssayByEssayId", zap.Error(err))
		return utils.GetError(50002)
	}
	return nil
}

func GetEssayByEssayId(essayId uint64) (essay *model.Essay, err error) {
	essay = &model.Essay{EssayId: essayId}
	if err := global.Db.Where("essay_id = ?", essay.EssayId).First(essay).Error; err != nil {
		zap.L().Error("GetEssayByEssayId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(50003)
		}
		return nil, utils.GetStrAndError("文章", 10009)
	}
	return essay, nil
}

func UpdateEssayByEssayId(essay *model.Essay) (err error) {
	res := global.Db.Model(essay).Where("essay_id = ?", essay.EssayId).Updates(map[string]interface{}{
		"essay_title":       essay.EssayTitle,
		"essay_author":      essay.EssayAuthor,
		"publish_at":        essay.PublishAt,
		"essay_content":     essay.EssayContent,
		"essay_isok":        essay.EssayIsOk,
		"essay_type":        essay.EssayType,
		"essay_collect_num": essay.EssayCollectNum,
	})
	if res.Error != nil {
		zap.L().Error("UpdateEssayByEssayId", zap.Error(err))
		return utils.GetError(50004)
	}
	return nil
}

func GetAllEssayToEssayInfo() ([]model.EssayInfo, error) {
	var essayList []model.Essay
	err := global.Db.Find(&essayList).Error
	if err != nil {
		return nil, utils.GetError(50010)
	}
	essayInfoList := make([]model.EssayInfo, 0, len(essayList))
	for _, essay := range essayList {
		essayInfo := model.EssayInfo{
			EssayId:      essay.EssayId,
			EssayTitle:   essay.EssayTitle,
			EssayAuthor:  essay.EssayAuthor,
			PublishAt:    essay.PublishAt,
			EssayContent: essay.EssayContent,
			EssayIsOk:    essay.EssayIsOk,
			EssayType:    essay.EssayType,
		}
		essayInfoList = append(essayInfoList, essayInfo)
	}
	return essayInfoList, nil
}

func GetEssayToEssayInfoByEssayType(essayType int8) ([]model.EssayInfo, error) {
	var essayList []model.Essay
	err := global.Db.Where("essay_type = ?", essayType).Find(&essayList).Error
	if err != nil {
		return nil, utils.GetError(50010)
	}
	essayInfoList := make([]model.EssayInfo, 0, len(essayList))
	for _, essay := range essayList {
		essayInfo := model.EssayInfo{
			EssayId:      essay.EssayId,
			EssayTitle:   essay.EssayTitle,
			EssayAuthor:  essay.EssayAuthor,
			PublishAt:    essay.PublishAt,
			EssayContent: essay.EssayContent,
			EssayIsOk:    essay.EssayIsOk,
			EssayType:    essay.EssayType,
		}
		essayInfoList = append(essayInfoList, essayInfo)
	}
	return essayInfoList, nil
}

func GetEssayPage(essayType int8, pageNum int, pageSize int) (essayList []model.Essay, total int64, err error) {
	essayList = make([]model.Essay, 0, 50)
	offset := (pageNum - 1) * pageSize
	if essayType != 0 {
		err = global.Db.Select("created_at, essay_id, essay_title, essay_type, essay_collect_num").Limit(pageSize).Offset(offset).Where("essay_type = ?", essayType).Order("publish_at desc").Find(&essayList).Error
		global.Db.Model(&essayList).Where("essay_type = ?", essayType).Count(&total)
		if err != nil {
			return nil, 0, utils.GetError(50016)
		}
		return essayList, total, nil
	}
	err = global.Db.Select("created_at, essay_id, essay_title, essay_type, essay_collect_num").Limit(pageSize).Offset(offset).Order("publish_at desc").Find(&essayList).Error
	global.Db.Model(&essayList).Count(&total)
	if err != nil {
		return nil, 0, utils.GetError(50016)
	}
	return essayList, total, nil
}

func GetAllEssayPage(essayType int8, pageNum int, pageSize int) (essayList []model.Essay, err error) {
	essayList = make([]model.Essay, 0, 20)
	offset := (pageNum - 1) * pageSize
	err = global.Db.Select("essay_id, essay_title, essay_author, publish_at, essay_collect_num").Limit(pageSize).Offset(offset).Where("essay_type = ?", essayType).Order("publish_at desc").Find(&essayList).Error
	if err != nil {
		return nil, utils.GetError(50016)
	}
	return essayList, nil
}
