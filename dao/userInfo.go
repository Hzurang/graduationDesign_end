package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateUserInfo(userInfo *model.UserInfo) (err error) {
	err = global.Db.Create(userInfo).Error
	if err != nil {
		zap.L().Error("用户信息"+utils.GetCodeMsg(10010), zap.Error(err))
		return utils.GetStrAndError("用户信息", 10010)
	}
	return
}

func GetUserInfoByUserId(userId uint64) (userInfo *model.UserInfo, err error) {
	userInfo = &model.UserInfo{UserId: userId}
	err = global.Db.Where("user_id = ?", userId).First(userInfo).Error
	if err != nil {
		zap.L().Error("GetUserInfoByUserId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30015)
		}
		return nil, utils.GetStrAndError("用户信息", 10009)
	}
	return userInfo, nil
}

func UpdateUserInfoIntegralByUserId(userInfo *model.UserInfo) (err error) {
	res := global.Db.Model(userInfo).Where("user_id = ?", userInfo.UserId).Update("integral", userInfo.Integral)
	if res.Error != nil {
		zap.L().Error("UpdateUserInfoIntegralByUserId", zap.Error(err))
		return utils.GetError(30016)
	}
	return nil
}

func UpdateUserInfoByUserId(userInfo *model.UserInfo) (err error) {
	res := global.Db.Model(userInfo).Where("user_id = ?", userInfo.UserId).Updates(map[string]interface{}{
		"le_xue_app_id": userInfo.LeXueAppId,
		"gender":        userInfo.Gender,
		"school":        userInfo.School,
		"birthday":      userInfo.Birthday,
		"area":          userInfo.Area,
		"nickname":      userInfo.NickName,
		"signature":     userInfo.Signature,
	})
	if res.Error != nil {
		zap.L().Error("UpdateUserInfoByUserId", zap.Error(err))
		return utils.GetError(30020)
	}
	return nil
}

func GetUserInfoByLeXueAppIdAndUserId(leXueAppId string, userId uint64) (userInfo *model.UserInfo, err error) {
	userInfo = &model.UserInfo{LeXueAppId: leXueAppId, UserId: userId}
	err = global.Db.Where("le_xue_app_id = ? AND user_id != ?", leXueAppId, userId).First(userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(30015)
		}
		return nil, utils.GetStrAndError("用户信息", 10009)
	}
	return userInfo, nil
}

func GetUserInfoByInvitationCode(codeStr string) (userInfo *model.UserInfo, err error) {
	userInfo = &model.UserInfo{InvitationCode: codeStr}
	err = global.Db.Where("invitation_code = ?", codeStr).First(userInfo).Error
	if err == nil {
		return userInfo, utils.GetError(30023)
	}
	return nil, nil
}

func UpdateUserInfoEngLevelByUserId(userInfo *model.UserInfo) (err error) {
	res := global.Db.Model(userInfo).Where("user_id = ?", userInfo.UserId).Update("eng_level", userInfo.EngLevel)
	if res.Error != nil {
		zap.L().Error("UpdateUserInfoEngLevelByUserId", zap.Error(err))
		return utils.GetError(30028)
	}
	return nil
}

func UpdateUserInfoWordNeedReciteNumByUserId(userInfo *model.UserInfo) (err error) {
	res := global.Db.Model(userInfo).Where("user_id = ?", userInfo.UserId).Update("word_need_recite_num", userInfo.WordNeedReciteNum)
	if res.Error != nil {
		zap.L().Error("UpdateUserInfoWordNeedReciteNumByUserId", zap.Error(err))
		return utils.GetError(30029)
	}
	return nil
}

func GetEngLevelByUserId(userId uint64) (engLevel *int8, err error) {
	userInfo := &model.UserInfo{UserId: userId}
	err = global.Db.Table("user_info").Select("eng_level").Where("user_id = ?", userId).First(userInfo).Error
	if err != nil {
		zap.L().Error("GetEngLevelByUserId", zap.Error(err))
		return nil, utils.GetError(30031)
	}
	engLevel = &userInfo.EngLevel
	return engLevel, nil
}
