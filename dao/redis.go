package dao

import (
	"fmt"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"time"
)

func RemoveToken() {
	values, _ := global.RD.ZRangeWithScores("JWT_AUTH_USER:Baned", 0, -1).Result()
	now := time.Now().Unix()
	// 一个月的时间
	oneMonth := int64(30 * 24 * 60 * 60)
	for _, v := range values {
		diff := now - int64(v.Score)
		if diff > oneMonth {
			global.RD.ZRem("JWT_AUTH_USER:Baned", v.Member)
		}
	}
}

func SetEssayCache(essayType string) (err error) {
	var essayInt int8
	switch essayType {
	case "novel":
		essayInt = 0
	case "love":
		essayInt = 1
	case "essays":
		essayInt = 2
	}
	essayInfoList, err := GetEssayToEssayInfoByEssayType(essayInt)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(50010), zap.Error(err))
		return
	}
	listName := fmt.Sprintf("%sList", essayType)
	_, err = utils.RPushWithMarshal(listName, essayInfoList)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(50013)+" 类型:"+essayType, zap.Error(err))
		return utils.GetError(50013)
	}
	return nil
}

func DeleteEssayCache(essayType string) (err error) {
	listName := fmt.Sprintf("%sList", essayType)
	var essayCacheInfoList []*model.EssayInfo
	_, err = utils.RPopWithUnMarshal(listName, essayCacheInfoList)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(50014)+" 类型:"+essayType, zap.Error(err))
		return utils.GetError(50014)
	}
	return nil
}

func SetListenCache(listenType string) (err error) {
	var listenInt int8
	switch listenType {
	case "热点资讯传送门":
		listenInt = 0
	case "国外媒体资讯":
		listenInt = 1
	case "英语听力入门":
		listenInt = 2
	case "可可之声":
		listenInt = 3
	case "品牌英语听力":
		listenInt = 4
	}
	listenCacheInfoList, err := GetListenToListenInfoByListenType(listenInt)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(60008), zap.Error(err))
		return err
	}
	listName := fmt.Sprintf("%dList", listenInt)
	_, err = utils.RPushWithMarshal(listName, listenCacheInfoList)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(60009)+" 类型:"+listenType, zap.Error(err))
		return utils.GetError(60009)
	}
	return nil
}

func DeleteListenCache(listenType string) (err error) {
	var listenInt int8
	switch listenType {
	case "热点资讯传送门":
		listenInt = 0
	case "国外媒体资讯":
		listenInt = 1
	case "英语听力入门":
		listenInt = 2
	case "可可之声":
		listenInt = 3
	case "品牌英语听力":
		listenInt = 4
	}
	listName := fmt.Sprintf("%dList", listenInt)
	var listenCacheInfoList []*model.ListenCacheInfo
	_, err = utils.RPopWithUnMarshal(listName, listenCacheInfoList)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(60010)+" 类型:"+listenType, zap.Error(err))
		return utils.GetError(60010)
	}
	return nil
}
