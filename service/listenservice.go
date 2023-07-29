package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	jsonIter "github.com/json-iterator/go"
	"strconv"
	"sync"
	"time"
)

/*
InsertListenService
@author: LJR
@Description: 手动添加听力业务逻辑
@param p
@return err
*/
func InsertListenService(p *model.ParamInsertListen) (err error) {
	_, err = dao.GetListenByTitle(p.ListenTitle)
	if err != nil {
		return err
	}
	listenId, _ := config.GenID()
	listen := &model.Listen{
		ListenId:         listenId,
		ListenTitle:      p.ListenTitle,
		ListenSource:     p.ListenSource,
		ListenEditor:     p.ListenEditor,
		PublishAt:        time.Now(),
		ListenContent:    p.ListenContent,
		ListenMediaPath:  p.ListenMediaPath,
		ListenMp3Path:    p.ListenMp3Path,
		ListenType:       p.ListenType,
		ListenSecondType: p.ListenSecondType,
		ListenCollectNum: 0,
	}
	if err := dao.CreateListen(listen); err != nil {
		return err
	}
	return nil
}

/*
DeleteListenService
@author: LJR
@Description: 删除听力业务逻辑
@param id
@return err
*/
func DeleteListenService(id string) (err error) {
	num, _ := strconv.Atoi(id)
	listenId := uint64(num)
	if err := dao.DeleteListenByListenId(listenId); err != nil {
		return err
	}
	// 补充用户收藏文章取消，因为文章软删除了
	_ = dao.DeleteCollectListenByListenId(listenId)
	_ = dao.DeleteListenWordByListenId(listenId)
	return nil
}

/*
GetListenService
@author: LJR
@Description: 获取具体听力信息业务逻辑
@param id
@return listenInfo, err
*/
func GetListenService(id string) (listenInfo *model.ListenInfo, err error) {
	num, _ := strconv.Atoi(id)
	listenId := uint64(num)
	listen, err := dao.GetListenByListenId(listenId)
	if err != nil {
		return nil, err
	}
	listenInfo = new(model.ListenInfo)
	listenInfo.ListenId = listen.ListenId
	listenInfo.ListenTitle = listen.ListenTitle
	listenInfo.ListenSource = listen.ListenSource
	listenInfo.ListenEditor = listen.ListenEditor
	listenInfo.PublishAt = listen.PublishAt
	listenInfo.ListenContent = listen.ListenContent
	listenInfo.ListenMediaPath = listen.ListenMediaPath
	listenInfo.ListenMp3Path = listen.ListenMp3Path
	listenInfo.ListenType = listen.ListenType
	listenInfo.ListenSecondType = listen.ListenSecondType
	listenInfo.ListenCollectNum = listen.ListenCollectNum
	return listenInfo, nil
}

/*
UpdateListenService
@author: LJR
@Description: 修改具体听力信息业务逻辑
@param p
@return err
*/
func UpdateListenService(p *model.ParamListenInfo) (err error) {
	_, err = dao.GetListenByListenId(p.ListenId)
	if err != nil {
		return err
	}
	listen := &model.Listen{
		ListenId:         p.ListenId,
		ListenTitle:      p.ListenTitle,
		ListenSource:     p.ListenSource,
		ListenEditor:     p.ListenEditor,
		ListenContent:    p.ListenContent,
		ListenMediaPath:  p.ListenMediaPath,
		ListenMp3Path:    p.ListenMp3Path,
		ListenType:       p.ListenType,
		ListenSecondType: p.ListenSecondType,
	}
	if err = dao.UpdateListenByListenId(listen); err != nil {
		return err
	}
	return nil
}

/*
GetListenByListenIdService
@author: LJR
@Description: 用户点击听力获取听力信息业务逻辑
@param id
@return listenInfo
@return listenWordList
@return err
*/
func GetListenByListenIdService(id string) (listenInfo *model.ListenInfo, listenWordList []model.ListenWordInfo, err error) {
	num, _ := strconv.Atoi(id)
	listenId := uint64(num)
	var wg sync.WaitGroup
	var listen *model.Listen
	var err1 error
	var err2 error
	listenInfo = new(model.ListenInfo)
	wg.Add(2)
	go func() {
		defer wg.Done()
		listen, err1 = dao.GetListenByListenId(listenId)
		listenInfo.ListenId = listen.ListenId
		listenInfo.ListenTitle = listen.ListenTitle
		listenInfo.ListenSource = listen.ListenSource
		listenInfo.ListenEditor = listen.ListenEditor
		listenInfo.PublishAt = listen.PublishAt
		listenInfo.ListenContent = listen.ListenContent
		listenInfo.ListenMediaPath = listen.ListenMediaPath
		listenInfo.ListenMp3Path = listen.ListenMp3Path
		listenInfo.ListenType = listen.ListenType
		listenInfo.ListenSecondType = listen.ListenSecondType
		listenInfo.ListenCollectNum = listen.ListenCollectNum
	}()
	go func() {
		defer wg.Done()
		listenWordList, err2 = dao.GetListenWordListByListenId(listenId)
	}()
	wg.Wait()
	if err1 != nil {
		return nil, nil, err1
	}
	if err2 != nil {
		return nil, nil, err2
	}
	return listenInfo, listenWordList, nil
}

/*
CollectListenService
@author: LJR
@Description: 用户收藏听力业务逻辑
@param listenId
@param userId
@return err
*/
func CollectListenService(listenId uint64, userId uint64) (err error) {
	listen, err := dao.GetListenByListenId(listenId)
	if err != nil {
		return utils.GetError(60011)
	}
	listenCollect, err := dao.GetSoftDeleteCollectListenByListenIdAndUserId(listenId, userId)
	if err == nil {
		// 软删除更新
		if listenCollect.DeleteIsOk == 1 {
			listen.ListenCollectNum = listen.ListenCollectNum + 1
			_ = dao.UpdateListenByListenId(listen)
			err = dao.UpdateSoftDeleteCollectListen(listenCollect)
			if err != nil {
				return err
			}
			return nil
		} else if listenCollect.DeleteIsOk == 0 {
			return utils.GetError(60013)
		}
	}
	listenCollect = &model.ListenCollect{
		ListenId: listenId,
		UserId:   userId,
	}
	err = dao.CreateCollectListen(listenCollect)
	if err != nil {
		return err
	}
	listen.ListenCollectNum = listen.ListenCollectNum + 1
	err = dao.UpdateListenByListenId(listen)
	if err != nil {
		return err
	}
	return nil
}

/*
CancelCollectListenService
@author: LJR
@Description: 用户取消文章听力业务逻辑
@param listenId
@param userId
@return err
*/
func CancelCollectListenService(listenId uint64, userId uint64) (err error) {
	listen, err := dao.GetListenByListenId(listenId)
	if err != nil {
		return utils.GetError(60014)
	}
	err = dao.DeleteCollectListenByListenIdAndUserId(listenId, userId)
	if err != nil {
		return err
	}
	listen.ListenCollectNum = listen.ListenCollectNum - 1
	err = dao.UpdateListenByListenId(listen)
	if err != nil {
		return err
	}
	return nil
}

/*
GetListenListService
@author: LJR
@Description: 用户获取听力列表业务逻辑
@param listenType
@return listenCacheInfoList
@return err
*/
func GetListenListService(listenType string) (listenCacheInfoList []model.ListenCacheInfo, err error) {
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
	p, _ := global.RD.LIndex(listName, 0).Result()
	listenCacheInfoList = make([]model.ListenCacheInfo, 0, 100)
	jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(p), &listenCacheInfoList)
	if listenCacheInfoList == nil {
		//0为热点资讯传送门，1为国外媒体资讯，2为英语听力入门，3为可可之声，4为品牌英语听力
		listenCacheInfoList, err = dao.GetListenToListenInfoByListenType(listenInt)
		if err != nil {
			return nil, err
		}
		return listenCacheInfoList, nil
	}
	return listenCacheInfoList, nil
}

/*
GetCollectListenService
@author: LJR
@Description: 用户获取听力收藏列表业务逻辑
@param userId
@return listenCacheInfoList
@return err
*/
func GetCollectListenService(userId uint64) (listenCacheInfoList []model.ListenCacheInfo, err error) {
	listenList, err := dao.GetListenCollectListByUserId(userId)
	if err != nil {
		return nil, err
	}
	listenCacheInfoList = make([]model.ListenCacheInfo, 0, len(listenList))
	for _, listen := range listenList {
		listenCacheInfo := model.ListenCacheInfo{
			ListenId:         listen.ListenId,
			ListenTitle:      listen.ListenTitle,
			ListenSource:     listen.ListenSource,
			ListenEditor:     listen.ListenEditor,
			PublishAt:        listen.PublishAt,
			ListenSecondType: listen.ListenSecondType,
			ListenCollectNum: listen.ListenCollectNum,
		}
		listenCacheInfoList = append(listenCacheInfoList, listenCacheInfo)
	}
	return listenCacheInfoList, nil
}

/*
GetListenAllService
@author: LJR
@Description: 根据听力类型和页数等加载听力列表业务逻辑
@param listen_type
@param page_num
@param page_size
@return WordList, total, err
*/
func GetListenAllService(listen_type string, page_num string, page_size string) (listenList []model.Listen, total int64, err error) {
	pageNum, _ := strconv.Atoi(page_num)
	pageSize, _ := strconv.Atoi(page_size)
	var listenType int8
	switch listen_type {
	case "17698":
		listenType = 1
	case "media":
		listenType = 2
	case "chuji":
		listenType = 3
	case "jiaoxue":
		listenType = 4
	case "brand":
		listenType = 5
	}
	listenList, total, err = dao.GetListenPage(listenType, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return listenList, total, nil
}
