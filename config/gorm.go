package config

import (
	"fmt"
	"ginStudy/global"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/*
InitGorm
@author: LJR
@Description: 初始化Gorm的配置
*/
func InitGorm() (err error) {
	mysqlConfig := Config.MySQLConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Dbname, mysqlConfig.Charset,
		mysqlConfig.ParseTime, mysqlConfig.TimeZone)
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: mysqlConfig.Gorm.SkipDefaultTx,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConfig.Gorm.TablePrefix,
			SingularTable: mysqlConfig.Gorm.SingularTable,
		},
		PrepareStmt:                              mysqlConfig.Gorm.PreparedStmt,
		DisableForeignKeyConstraintWhenMigrating: mysqlConfig.Gorm.CloseForeignKey,
	}
	// 是否覆盖默认sql配置
	if mysqlConfig.Gorm.CoverLogger {
		setNewLogger(gormConfig)
	}
	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DisableDatetimePrecision:  false,
		SkipInitializeWithVersion: false,
	}), gormConfig)

	if err != nil {
		zap.L().DPanic(utils.GetCodeMsg(90005), zap.Error(err))
		return
	}

	mysqlDb, err := client.DB()
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(90006), zap.Error(err))
		return
	}

	mysqlDb.SetMaxIdleConns(Config.MySQLConfig.MaxIdleCons)
	mysqlDb.SetMaxOpenConns(Config.MySQLConfig.MaxOpenCons)

	// 赋值给全局变量
	global.Db = client

	//// 是否调用数据迁移
	//if mysqlConfig.AutoMigrate {
	//	core.AutoMigrate()
	//}

	//defer func() {
	//	if global.MysqlClient != nil {
	//		fmt.Println("数据库关闭连接！！！")
	//		db, _ := global.MysqlClient.DB()
	//		_ = db.Close()
	//	}
	//}()
	return
}

/*
CloseMySQL
@author: LJR
@Description: 封装 db.Close 方法
*/
func CloseMySQL() {
	db, _ := global.Db.DB()
	err := db.Close()
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(90007), zap.Error(err))
		return
	}
}

/*
setNewLogger
@author: LJR
@Description: 覆盖原先gorm的logger配置
@param gConfig
*/
func setNewLogger(gConfig *gorm.Config) {
	file, _ := os.OpenFile("../logs/sql/sql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	// 日志级别映射 error、info、warn
	logLevelMap := map[string]logger.LogLevel{
		"error": logger.Error,
		"info":  logger.Info,
		"warn":  logger.Warn,
	}
	var logLevel logger.LogLevel
	var ok bool
	if logLevel, ok = logLevelMap[Config.MySQLConfig.LogLevel]; !ok {
		logLevel = logger.Error
	}
	newLogger := logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             50 * time.Microsecond, //慢SQL时间
		LogLevel:                  logLevel,              // 记录日志级别
		IgnoreRecordNotFoundError: true,                  // 是否忽略ErrRecordNotFound(未查到记录错误)
		Colorful:                  false,                 // 开关颜色
	})
	gConfig.Logger = newLogger
}
