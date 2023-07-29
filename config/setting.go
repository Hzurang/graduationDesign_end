package config

import (
	"fmt"
	"ginStudy/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config 全局变量 保存所有的配置信息
var Config = new(AppConfig)

/*
AppConfig
@author: LJR
@Description: 全局配置信息
@variable Name 'desc':名称
@variable AppMode 'desc':gin的具体生产环境
@variable Mode 'desc':log的环境
@variable MachineId 'desc':用于雪花算法生成ID
@variable Port 'desc':后端应用端口
*/
type AppConfig struct {
	Name              string `mapstructure:"name"`
	AppMode           string `mapstructure:"app_mode"`
	Mode              string `mapstructure:"mode"`
	StartTime         string `mapstructure:"start_time"`
	MachineId         uint16 `mapstructure:"machine_id"`
	HostIP            string `mapstructure:"host_ip"`
	Port              int    `mapstructure:"port"`
	*LogConfig        `mapstructure:"log"`
	*MySQLConfig      `mapstructure:"mysql"`
	*RedisConfig      `mapstructure:"redis"`
	*JWTConfig        `mapstructure:"jwt"`
	*APISMSCodeConfig `mapstructure:"api_sms"`
	*QQEmailConfig    `mapstructure:"qq_email"`
	*RefreshDbConfig  `mapstructure:"refresh_db"`
}

/*
LogConfig
@author: LJR
@Description: 日志配置信息
@variable Level 'desc':日志记录级别
@variable FileFormat 'desc':日志文件名格式
@variable Path 'desc':日志文件目录
@variable FileName 'desc':日志文件名
@variable MaxSize 'desc':单文件最大容量(单位MB)
@variable MaxAge 'desc':旧文件最多保存几天
@variable MaxBackups 'desc':保留旧文件的最大数量
*/
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileFormat string `mapstructure:"file_format"`
	Path       string `mapstructure:"path"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

/*
MySQLConfig
@author: LJR
@Description: mysql 配置信息
@variable Host 'desc':数据库主机地址
@variable Port 'desc':数据库连接端口
@variable User 'desc':数据库连接用户
@variable Password 'desc':密码
@variable Dbname 'desc':数据库名称
@variable Charset 'desc':数据库字符集
@variable ParseTime 'desc':解析time.Time类型
@variable AutoMigrate 'desc':开启时，每次服务启动都会根据实体创建/更新表结构
@variable TimeZone 'desc':时区, yaml 若设置 Asia/Shanghai, 需写成: Asia%2fShanghai
@variable LogLevel 'desc':数据库日志级别
@variable MaxOpenCons 'desc':数据库的最大连接数量
@variable MaxIdleCons 'desc':连接池中的最大闲置连接数
*/
type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	Charset     string `mapstructure:"charset"`
	ParseTime   bool   `mapstructure:"parseTime"`
	AutoMigrate bool   `mapstructure:"autoMigrate"`
	TimeZone    string `mapstructure:"timeZone"`
	LogLevel    string `mapstructure:"logLevel"`
	MaxOpenCons int    `mapstructure:"max_open_cons"`
	MaxIdleCons int    `mapstructure:"max_idle_cons"`
	Gorm        Gorm   `mapstructure:"gorm"`
}

/*
Gorm
@author: LJR
@Description: gorm 配置信息
@variable SkipDefaultTx 'desc':是否跳过默认事务
@variable TablePrefix 'desc':表前缀
@variable SingularTable 'desc':是否使用单数表名(默认复数)，启用后，User结构体表将是user
@variable CoverLogger 'desc':是否覆盖默认logger
@variable PreparedStmt 'desc':设置SQL缓存
@variable CloseForeignKey 'desc':禁用外键约束
*/
type Gorm struct {
	SkipDefaultTx   bool   `mapstructure:"skipDefaultTx"`
	TablePrefix     string `mapstructure:"tablePrefix"`
	SingularTable   bool   `mapstructure:"singularTable"`
	CoverLogger     bool   `mapstructure:"coverLogger"`
	PreparedStmt    bool   `mapstructure:"prepareStmt"`
	CloseForeignKey bool   `mapstructure:"disableForeignKeyConstraintWhenMigrating"`
}

/*
RedisConfig
@author: LJR
@Description: Redis 配置信息
@variable Host 'desc':Redis的主机地址
@variable Password 'desc':密码
@variable Port 'desc':Redis连接端口
@variable DB 'desc':
@variable PoolSize 'desc':连接池大小
*/
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

/*
JWTConfig
@author: LJR
@Description: JWT 配置信息
@variable Secret 'desc':密钥
@variable Issuer 'desc':签发人
*/
type JWTConfig struct {
	Secret  string `mapstructure:"secret"`
	Issuer  string `mapstructure:"issuer"`
	Subject string `mapstructure:"subject"`
}

/*
APISMSCodeConfig
@author: LJR
@Description: 短信配置信息
@variable APPCode 'desc':信息配置
*/
type APISMSCodeConfig struct {
	APPCode string `mapstructure:"app_code"`
}

/*
QQEmailConfig
@author: LJR
@Description: QQ邮箱配置信息
@variable UserName 'desc':QQ邮箱用户名
@variable Secret 'desc':QQ邮箱服务密钥
*/
type QQEmailConfig struct {
	UserName string `mapstructure:"username"`
	Secret   string `mapstructure:"Secret"`
}

/*
RefreshDbConfig
@author: LJR
@Description: 数据库清除的配置信息
@variable UserName 'desc':用户名
@variable Password 'desc':密码
*/
type RefreshDbConfig struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func InitConfig() (err error) {
	viper.SetConfigFile("yun_config.yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(90001), zap.Error(err))
		return
	}

	// 将读取的配置文件反序列化到Config结构中
	if err := viper.Unmarshal(Config); err != nil {
		zap.L().Error(utils.GetCodeMsg(90002), zap.Error(err))
	}

	// 热更新配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		// 同步修改Config结构体的值
		if err := viper.Unmarshal(Config); err != nil {
			zap.L().Error(utils.GetCodeMsg(90002), zap.Error(err))
		}
	})
	return
}
