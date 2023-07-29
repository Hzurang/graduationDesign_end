package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateAdmin(admin *model.Admin) (err error) {
	err = global.Db.Create(admin).Error
	if err != nil {
		zap.L().Error("CreateAdmin", zap.Error(err))
		return utils.GetStrAndError("管理员", 10010)
	}
	return
}

func GetAdminUserByUserName(userName string) (admin *model.Admin, err error) {
	admin = &model.Admin{UserName: userName}
	err = global.Db.Where("user_name = ?", userName).First(admin).Error
	if err != nil {
		zap.L().Error("GetAdminUserByUserName", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30035)
		}
		return nil, utils.GetStrAndError("管理员", 10009)
	}
	return admin, nil
}
