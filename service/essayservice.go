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
	"time"
)

/*
InsertEssayService
@author: LJR
@Description: 手动添加文章业务逻辑
@param p
@return err
*/
func InsertEssayService(p *model.ParamInsertEssay) (err error) {
	_, err = dao.GetEssayByTitleAndContent(p.EssayTitle, p.EssayContent)
	if err != nil {
		return err
	}
	var author string
	id, _ := config.GenID()
	if p.EssayAuthor == "" {
		author = "乐学"
	} else {
		author = p.EssayAuthor
	}
	essay := &model.Essay{
		EssayId:         id,
		EssayTitle:      p.EssayTitle,
		PublishAt:       time.Now(),
		EssayContent:    p.EssayContent,
		EssayIsOk:       p.EssayIsOk,
		EssayType:       p.EssayType,
		EssayAuthor:     author,
		EssayCollectNum: 0,
	}
	if err := dao.CreateEssay(essay); err != nil {
		return err
	}
	return nil
}

/*
DeleteEssayService
@author: LJR
@Description: 删除文章业务逻辑
@param id
@return err
*/
func DeleteEssayService(id string) (err error) {
	num, _ := strconv.Atoi(id)
	essayId := uint64(num)
	if err := dao.DeleteEssayByEssayId(essayId); err != nil {
		return err
	}
	// 补充用户收藏文章取消，因为文章软删除了
	_ = dao.DeleteCollectEssayByEssayId(essayId)
	_ = dao.DeleteEssayWordByEssayId(essayId)
	return nil
}

/*
GetEssayService
@author: LJR
@Description: 获取具体文章信息业务逻辑
@param id
@return essayInfo, err
*/
func GetEssayService(id string) (essayInfo *model.EssayInfo, err error) {
	num, _ := strconv.Atoi(id)
	essayId := uint64(num)
	essay, err := dao.GetEssayByEssayId(essayId)
	if err != nil {
		return nil, err
	}
	essayInfo = new(model.EssayInfo)
	essayInfo.EssayId = essay.EssayId
	essayInfo.EssayTitle = essay.EssayTitle
	essayInfo.EssayAuthor = essay.EssayAuthor
	essayInfo.PublishAt = essay.PublishAt
	essayInfo.EssayContent = essay.EssayContent
	essayInfo.EssayIsOk = essay.EssayIsOk
	essayInfo.EssayType = essay.EssayType
	essayInfo.EssayCollectNum = essay.EssayCollectNum
	return essayInfo, nil
}

/*
UpdateEssayService
@author: LJR
@Description: 修改具体文章信息业务逻辑
@param p
@return err
*/
func UpdateEssayService(p *model.ParamEssayInfo) (err error) {
	_, err = dao.GetEssayByEssayId(p.EssayId)
	if err != nil {
		return err
	}
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, p.PublishAt)
	essay := &model.Essay{
		EssayId:      p.EssayId,
		EssayTitle:   p.EssayTitle,
		EssayAuthor:  p.EssayAuthor,
		PublishAt:    t,
		EssayContent: p.EssayContent,
		EssayIsOk:    p.EssayIsOk,
		EssayType:    p.EssayType,
	}
	if err = dao.UpdateEssayByEssayId(essay); err != nil {
		return err
	}
	return nil
}

/*
CollectEssayService
@author: LJR
@Description: 用户收藏文章业务逻辑
@param essayId
@param userId
@return err
*/
func CollectEssayService(essayId uint64, userId uint64) (err error) {
	essay, err := dao.GetEssayByEssayId(essayId)
	if err != nil {
		return utils.GetError(50009)
	}
	essayCollect, err := dao.GetSoftDeleteCollectEssayByEssayIdAndUserId(essayId, userId)
	if err == nil {
		// 软删除更新
		if essayCollect.DeleteIsOk == 1 {
			essay.EssayCollectNum = essay.EssayCollectNum + 1
			_ = dao.UpdateEssayByEssayId(essay)
			err = dao.UpdateSoftDeleteCollectEssay(essayCollect)
			if err != nil {
				return err
			}
			return nil
		} else if essayCollect.DeleteIsOk == 0 {
			return utils.GetError(50006)
		}
	}
	essayCollect = &model.EssayCollect{
		EssayId: essayId,
		UserId:  userId,
	}
	err = dao.CreateCollectEssay(essayCollect)
	if err != nil {
		return err
	}
	essay.EssayCollectNum = essay.EssayCollectNum + 1
	err = dao.UpdateEssayByEssayId(essay)
	if err != nil {
		return err
	}
	return nil
}

/*
CancelCollectEssayService
@author: LJR
@Description: 用户取消文章收藏业务逻辑
@param essayId
@param userId
@return err
*/
func CancelCollectEssayService(essayId uint64, userId uint64) (err error) {
	essay, err := dao.GetEssayByEssayId(essayId)
	if err != nil {
		return utils.GetError(50011)
	}
	err = dao.DeleteCollectEssayByEssayIdAndUserId(essayId, userId)
	if err != nil {
		return err
	}
	essay.EssayCollectNum = essay.EssayCollectNum - 1
	err = dao.UpdateEssayByEssayId(essay)
	if err != nil {
		return err
	}
	return nil
}

///*
//GetEssayByEssayIdService
//@author: LJR
//@Description: 用户点击文章业务逻辑
//@param id
//@return essayInfo
//@return essayWordList
//@return err
//*/
//func GetEssayByEssayIdService(id string) (essayInfo *model.EssayInfo, essayWordList []model.EssayWordInfo, err error) {
//	num, _ := strconv.Atoi(id)
//	essayId := uint64(num)
//	// 获取文章信息
//	// 获取word的信息
//	var wg sync.WaitGroup
//	var essay *model.Essay
//	var err1 error
//	var err2 error
//	essayInfo = new(model.EssayInfo)
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		essay, err1 = dao.GetEssayByEssayId(essayId)
//		essayInfo.EssayId = essay.EssayId
//		essayInfo.EssayTitle = essay.EssayTitle
//		essayInfo.EssayAuthor = essay.EssayAuthor
//		essayInfo.PublishAt = essay.PublishAt
//		essayInfo.EssayContent = essay.EssayContent
//		essayInfo.EssayIsOk = essay.EssayIsOk
//		essayInfo.EssayType = essay.EssayType
//	}()
//	go func() {
//		defer wg.Done()
//		essayWordList, err2 = dao.GetEssayWordListByEssayId(essayId)
//	}()
//	wg.Wait()
//	if err1 != nil {
//		return nil, nil, err1
//	}
//	if err2 != nil {
//		return nil, nil, err2
//	}
//	return essayInfo, essayWordList, nil
//}

/*
GetEssayByEssayIdService
@author: LJR
@Description: 用户点击文章业务逻辑
@param id
@return essayDetail
@return err
*/
func GetEssayByEssayIdService(id string, userId uint64) (essayDetail *model.EssayDetail, err error) {
	num, _ := strconv.Atoi(id)
	essayId := uint64(num)
	var essay *model.Essay
	essayDetail = new(model.EssayDetail)
	essay, err = dao.GetEssayByEssayId(essayId)
	if err != nil {
		return nil, err
	}
	_, err1 := dao.GetCollectEssayByEssayIdAndUserId(essayId, userId)
	if err1 != nil {
		essayDetail.IsCollect = 1
	} else {
		essayDetail.IsCollect = 0
	}
	essayDetail.EssayId = strconv.FormatUint(essay.EssayId, 10)
	essayDetail.EssayTitle = essay.EssayTitle
	essayDetail.EssayAuthor = essay.EssayAuthor
	essayDetail.PublishAt = essay.PublishAt.Format("2006-01-02")
	essayDetail.EssayContent = essay.EssayContent
	essayDetail.EssayType = essay.EssayType
	return essayDetail, nil
}

/*
GetCollectEssayService
@author: LJR
@Description: 用户获取文章收藏列表业务逻辑
@param userId
@return sentenceCollectInfoList
@return err
*/
func GetCollectEssayService(userId uint64) (essayCollectInfoList []model.EssayCollectInfo, err error) {
	essayList, err := dao.GetEssayCollectListByUserId(userId)
	if err != nil {
		return nil, err
	}
	essayCollectInfoList = make([]model.EssayCollectInfo, 0, len(essayCollectInfoList))
	for _, essay := range essayList {
		essayCollectInfo := model.EssayCollectInfo{
			EssayId:         essay.EssayId,
			EssayTitle:      essay.EssayTitle,
			EssayAuthor:     essay.EssayAuthor,
			PublishAt:       essay.PublishAt,
			EssayType:       essay.EssayType,
			EssayCollectNum: essay.EssayCollectNum,
		}
		essayCollectInfoList = append(essayCollectInfoList, essayCollectInfo)
	}
	return essayCollectInfoList, nil
}

/*
GetEssayListService
@author: LJR
@Description: 用户获取文章列表业务逻辑
@param listenType
@return listenCacheInfoList
@return err
*/
func GetEssayListService(essayType string) (essayInfoList []model.EssayInfo, err error) {
	var essayInt int8
	listName := fmt.Sprintf("%sList", essayType)
	p, _ := global.RD.LIndex(listName, 0).Result()
	essayInfoList = make([]model.EssayInfo, 0, 100)
	jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(p), &essayInfoList)
	if essayInfoList == nil {
		//0为热点资讯传送门，1为国外媒体资讯，2为英语听力入门，3为可可之声，4为品牌英语听力
		switch essayType {
		case "novel":
			essayInt = 0
		case "love":
			essayInt = 1
		case "essays":
			essayInt = 2
		}
		essayInfoList, err = dao.GetEssayToEssayInfoByEssayType(essayInt)
		if err != nil {
			return nil, err
		}
		return essayInfoList, nil
	}
	return essayInfoList, nil
}

/*
GetEssayAllService
@author: LJR
@Description: 根据文章类型和页数等加载文章列表
@param essay_type
@param page_num
@param page_size
@return WordList, total, err
*/
func GetEssayAllService(essay_type string, page_num string, page_size string) (essayList []model.Essay, total int64, err error) {
	pageNum, _ := strconv.Atoi(page_num)
	pageSize, _ := strconv.Atoi(page_size)
	var essayType int8
	switch essay_type {
	case "novel":
		essayType = 1
	case "love":
		essayType = 2
	case "essays":
		essayType = 3
	case "shuangyu":
		essayType = 4
	}
	essayList, total, err = dao.GetEssayPage(essayType, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return essayList, total, nil
}

/*
GetAllEssayService
@author: LJR
@Description: 根据文章类型和页数等加载文章列表
@param essay_type
@param page_num
@param page_size
@return essayReturnList, err
*/
func GetAllEssayService(essay_type string, page_num string, page_size string) (essayReturnList []model.EssayList, err error) {
	pageNum, _ := strconv.Atoi(page_num)
	pageSize, _ := strconv.Atoi(page_size)
	var essayType int8
	switch essay_type {
	case "novel":
		essayType = 1
	case "love":
		essayType = 2
	case "essays":
		essayType = 3
	case "shuangyu":
		essayType = 4
	}
	essayList, err := dao.GetAllEssayPage(essayType, pageNum, pageSize)
	essayReturnList = make([]model.EssayList, 0, len(essayList))
	for _, essay := range essayList {
		essayReturn := model.EssayList{
			EssayId:         strconv.FormatUint(essay.EssayId, 10),
			EssayTitle:      essay.EssayTitle,
			EssayAuthor:     essay.EssayAuthor,
			EssayCollectNum: strconv.FormatUint(essay.EssayCollectNum, 10),
			PublishAt:       essay.PublishAt.Format("2006-01-02"),
		}
		essayReturnList = append(essayReturnList, essayReturn)
	}
	if err != nil {
		return nil, err
	}
	return essayReturnList, nil
}
