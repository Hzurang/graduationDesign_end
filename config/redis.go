package config

import (
	"fmt"
	"ginStudy/global"
	"ginStudy/utils"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

/*
InitRedis
@author: LJR
@Description: 初始化 redis
@return err
*/
func InitRedis() (err error) {
	global.RD = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			Config.RedisConfig.Host,
			Config.RedisConfig.Port),
		Password: Config.RedisConfig.Password,
		DB:       Config.RedisConfig.DB,
		PoolSize: Config.RedisConfig.PoolSize,
	})
	_, err = global.RD.Ping().Result()
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(90010), zap.Error(err))
		return err
	}
	return
}

/*
CloseRedis
@author: LJR
@Description: 封装关闭 redis 的方法
*/
func CloseRedis() {
	err := global.RD.Close()
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(90011), zap.Error(err))
		return
	}
}
