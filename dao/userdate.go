package dao

import (
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"time"
)

func CreateUserDate(userDate *model.UserDate) (err error) {
	err = global.Db.Create(userDate).Error
	if err != nil {
		zap.L().Error("CreateUserDate", zap.Error(err))
		return utils.GetStrAndError("用户学习记录", 10010)
	}
	return
}

func GetUserDateByUserId(userId uint64) (userDate *model.UserDate, err error) {
	userDate = &model.UserDate{}
	err = global.Db.Where("user_id = ?", userId).Order("date desc").First(userDate).Error
	if err != nil {
		return nil, utils.GetError(30032)
	}
	return userDate, nil
}

func UpdateUserDateByUserId(userDate *model.UserDate) (err error) {
	res := global.Db.Model(userDate).Where("user_id = ?", userDate.UserId).Updates(map[string]interface{}{
		"date":               userDate.Date,
		"word_learn_number":  userDate.WordLearnNumber,
		"word_review_number": userDate.WordReviewNumber,
		"remark":             userDate.Remark,
	})
	if res.Error != nil {
		zap.L().Error("UpdateUserDateByUserId", zap.Error(err))
		return utils.GetError(30033)
	}
	return nil
}

func GetDaysTotalByUserId(userId uint64) (total int64, err error) {
	userDate := &model.UserDate{}
	err = global.Db.Model(userDate).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		zap.L().Error("GetDaysTotalByUserId", zap.Error(err))
		return 0, nil
	}
	return total, nil
}

func CalculateTotalWordLearnNumber(userId uint64) (int, error) {
	var totalWordLearnNumber int
	userDate := &model.UserDate{}
	err := global.Db.Model(userDate).Select("COALESCE(SUM(word_learn_number), 0)").Where("user_id = ?", userId).Scan(&totalWordLearnNumber).Error
	if err != nil {
		return 0, err
	}
	return totalWordLearnNumber, nil
}

func GetUserDateByUserIdAndDate(userId uint64, date time.Time) (err error) {
	userDate := &model.UserDate{}
	err = global.Db.Where("date = ? AND user_id = ?", date, userId).Find(userDate).Error
	if err == nil {
		return utils.GetError(30038)
	}
	return nil
}

func GetUserDateListByUserId(userId uint64) (userDateList []model.UserDate, err error) {
	userDateList = make([]model.UserDate, 0, 30)
	err = global.Db.Where("user_id = ?", userId).Find(&userDateList).Error
	if err != nil {
		return nil, utils.GetError(30040)
	}
	return userDateList, nil
}

func GetUserDateByUserIdAndOther(userId uint64, date time.Time) (userDate *model.UserDate, err error) {
	userDate = &model.UserDate{}
	err = global.Db.Where("user_id = ? AND date = ?", userId, date).First(userDate).Error
	if err != nil {
		return nil, utils.GetError(30040)
	}
	return userDate, nil
}
