package main

import (
	"context"
	"fmt"
	"ginStudy/config"
	"ginStudy/global"
	"ginStudy/initialize"
	"ginStudy/model"
	"ginStudy/routes"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initialize.InitAll()
	defer config.CloseMySQL()
	defer config.CloseRedis()
	defer func(log *zap.Logger) {
		_ = log.Sync()
	}(zap.L())
	//id, _ := config.GenID()
	//admin := &model.Admin{
	//	AdminId:  id,
	//	UserName: "backend",
	//	Password: encrypt.GetSHA256HashCode("LJRljr109109"),
	//}
	//dao.CreateAdmin(admin)

	//c := cron.InitCron()
	//c.Start()

	//c := colly.NewContext()
	//service.WordSpiderService(c, "CET4")

	//pattern := `\[.*\]`
	//re := regexp.MustCompile(pattern)
	//wordList, _ := dao.GetWordByWordType(3)
	//for _, word := range wordList {
	//	match1 := re.FindString(word.PhoneticTransEng)
	//	match2 := re.FindString(word.PhoneticTransAme)
	//	if match2 == "" {
	//		word.PhoneticTransEng = match1
	//		word.PhoneticTransAme = match1
	//		dao.UpdateWordType(word)
	//		continue
	//	} else if match2 == "" && match1 == "" {
	//		continue
	//	}
	//	word.PhoneticTransEng = match1
	//	word.PhoneticTransAme = match2
	//	dao.UpdateWordType(word)
	//	fmt.Println(word.PhoneticTransEng, "      ", word.PhoneticTransAme)
	//}
	//c := colly.NewContext()
	//service.ListenSpiderService(c, "17698")
	global.Db.Migrator().DropTable(&model.UserFeedback{})
	global.Db.AutoMigrate(&model.UserFeedback{})
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	route := routes.InitRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.Port),
		Handler: route,
	}
	//_ = route.Run(fmt.Sprintf("%s:%d", config.Config.HostIP, config.Config.Port))
	// 在单独的goroutine中启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()
	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 程序结束时优雅地关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown:", zap.Error(err))
	}
}
