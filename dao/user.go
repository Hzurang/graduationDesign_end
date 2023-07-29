package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) (err error) {
	err = global.Db.Create(user).Error
	if err != nil {
		zap.L().Error("CreateUser", zap.Error(err))
		return utils.GetStrAndError("用户", 10010)
	}
	return
}

func GetUserByMobile(mobile string) (user *model.User, err error) {
	user = &model.User{Mobile: mobile}
	err = global.Db.Where("mobile = ?", mobile).First(user).Error
	if err != nil {
		zap.L().Error("GetUserByMobile", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30001)
		}
		return nil, utils.GetStrAndError("用户", 10009)
	}
	return user, nil
}

func GetUserByUserId(userId uint64) (user *model.User, err error) {
	user = &model.User{UserId: userId}
	err = global.Db.Where("user_id = ?", userId).First(user).Error
	if err != nil {
		zap.L().Error("GetUserByUserId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30001)
		}
		return nil, utils.GetStrAndError("用户", 10009)
	}
	return user, nil
}

func UpdateUserPasswordByUserId(user *model.User) (err error) {
	res := global.Db.Model(&user).Where("user_id = ?", user.UserId).Update("password", user.Password)
	if res.Error != nil {
		zap.L().Error("UpdateUserPasswordByUserId", zap.Error(err))
		return utils.GetError(30013)
	}
	return nil
}

func UpdateUserLoginTimeByUserId(user *model.User) (err error) {
	res := global.Db.Model(user).Where("user_id = ?", user.UserId).Update("last_login_time", user.LastLoginTime)
	if res.Error != nil {
		zap.L().Error("UpdateUserLoginTimeByUserId", zap.Error(err))
		return utils.GetError(30030)
	}
	return nil
}

func UpdateUserEmailByUserId(user *model.User) (err error) {
	res := global.Db.Model(&user).Where("user_id = ?", user.UserId).Update("email", user.Email)
	if res.Error != nil {
		zap.L().Error("UpdateUserEmailByUserId", zap.Error(err))
		return utils.GetError(30014)
	}
	return nil
}

func SoftDeleteUserByUserId(userId uint64) (err error) {
	user := &model.User{UserId: userId}
	global.Db.Model(user).Where("user_id = ?", userId).Update("status", 1)
	num := global.Db.Where("user_id = ?", userId).Delete(user).RowsAffected
	if num == 0 {
		zap.L().Error("SoftDeleteUserByUserId", zap.Error(err))
		return utils.GetError(30026)
	}
	return nil
}

func UpdateSoftDeleteUser(user *model.User) (err error) {
	err = global.Db.Unscoped().Model(user).Where("user_id = ?", user.UserId).Updates(map[string]interface{}{
		"deleted_at": nil,
		"status":     0,
	}).Error
	if err != nil {
		zap.L().Error("UpdateSoftDeleteUser", zap.Error(err))
		return utils.GetError(30027)
	}
	return nil
}

func GetSoftDeleteUserByUserId(userId uint64) (user *model.User, err error) {
	user = &model.User{UserId: userId}
	err = global.Db.Unscoped().Where("user_id = ?", userId).First(user).Error
	if err != nil {
		zap.L().Error("GetSoftDeleteUserByUserId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30001)
		}
		return nil, utils.GetStrAndError("用户", 10009)
	}
	return user, nil
}

func GetUserPage(pageNum int, pageSize int) (userList []model.User, total int64, err error) {
	userList = make([]model.User, 0, 50)
	offset := (pageNum - 1) * pageSize
	err = global.Db.Unscoped().Select("created_at, user_id, email, mobile, status").Limit(pageSize).Offset(offset).Order("created_at desc").Find(&userList).Error
	global.Db.Model(&userList).Count(&total)
	if err != nil {
		return nil, 0, utils.GetError(30037)
	}
	return userList, total, nil
}
