package initialize

import (
	"ginStudy/config"
	"ginStudy/utils"
	"go.uber.org/zap"
)

func InitAll() {
	// 1.初始化配置信息
	if err := config.InitConfig(); err != nil {
		zap.L().Error(utils.GetCodeMsg(90003), zap.Error(err))
	}
	// 2.初始化翻译器
	if err := config.InitTrans("zh"); err != nil {
		zap.L().Error(utils.GetCodeMsg(90012), zap.Error(err))
	}

	// 3.初始化日志
	if err := config.InitLogger(); err != nil {
		zap.L().Error(utils.GetCodeMsg(90004), zap.Error(err))
	}
	zap.L().Debug("logger init success---------------")

	// 4.数据库初始化
	if err := config.InitGorm(); err != nil {
		zap.L().Error(utils.GetCodeMsg(10009), zap.Error(err))
	}

	// 5.雪花算法初始化
	if err := config.InitSnowFlake(config.Config.StartTime, config.Config.MachineId); err != nil {
		zap.L().Error(utils.GetCodeMsg(90009), zap.Error(err))
		return
	}

	// 6.初始化redis连接
	if err := config.InitRedis(); err != nil {
		zap.L().Error(utils.GetCodeMsg(90010), zap.Error(err))
	}
}
