package cron

import (
	"fmt"
	"ginStudy/dao"
	"ginStudy/service"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"
)

/*
InitCron
@author: LJR
@Description: 初始化 cron 定时任务
@return err
*/
func InitCron() *cron.Cron {
	c := cron.New()

	// 每周一的 00:00 更新将 15 天以来没有使用过的 Token 移除 Token 黑名单
	c.AddFunc("0 0 0 * * 0", func() {
		dao.RemoveToken()
	})

	// 每天晚上的 22:30 更新次日的每日一句
	c.AddFunc("30 22 * * *", func() {
		_, _ = service.DailySentenceEveryDayService()
	})

	// 每天 14:00 更新听力资源（缓存）（并发）
	c.AddFunc("0 14 * * *", func() {
		wg := sync.WaitGroup{}
		wg.Add(5)
		go func() {
			defer wg.Done()
			err := SpiderListenCron("17698")
			if err != nil {
				zap.L().Error("17698 error", zap.Error(err))
			}
		}()
		go func() {
			defer wg.Done()
			err := SpiderListenCron("media")
			if err != nil {
				zap.L().Error("media error", zap.Error(err))
			}
		}()
		go func() {
			defer wg.Done()
			err := SpiderListenCron("chuji")
			if err != nil {
				zap.L().Error("chuji error", zap.Error(err))
			}
		}()
		go func() {
			defer wg.Done()
			err := SpiderListenCron("jiaoxue")
			if err != nil {
				zap.L().Error("jiaoxue error", zap.Error(err))
			}
		}()
		go func() {
			defer wg.Done()
			err := SpiderListenCron("brand")
			if err != nil {
				zap.L().Error("brand error", zap.Error(err))
			}
		}()
		wg.Wait()
	})
	return c
}

func SpiderListenCron(listenType string) (err error) {
	c := colly.NewContext()
	totalPageNum := service.ListenPageNumSpiderService(listenType)
	service.ListenUrlSpiderService(c, listenType, totalPageNum, true)
	cnt1, cnt2 := service.ListenSpiderService(c, listenType)
	var listen string
	switch listenType {
	case "17698":
		listen = "热点资讯传送门"
	case "media":
		listen = "媒体资讯"
	case "chuji":
		listen = "英语听力入门"
	case "jiaoxue":
		listen = "可可之声"
	case "brand":
		listen = "品牌英语听力"
	}
	zap.L().Info(fmt.Sprintf("更新%s类型听力共%d条，所对应的单词共%d条", listen, cnt1, cnt2))
	err = dao.DeleteListenCache(listenType)
	if err != nil {
		return err
	}
	err = dao.SetListenCache(listenType)
	if err != nil {
		return err
	}
	return nil
}
