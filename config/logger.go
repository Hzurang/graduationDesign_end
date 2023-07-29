package config

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"time"
)

var zapLog *zap.Logger

func InitLogger() (err error) {
	writeSyncer := getLogWrite(Config.LogConfig.MaxSize, Config.LogConfig.MaxAge, Config.LogConfig.MaxBackups)
	encoder := getEncoder()
	var level = new(zapcore.Level)
	err = level.UnmarshalText([]byte(Config.LogConfig.Level))
	if err != nil {
		return
	}

	var core zapcore.Core
	if Config.Mode == "dev" {
		// 进入开发模式
		consoleEncode := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),                                // 文件输出core
			zapcore.NewCore(consoleEncode, zapcore.Lock(os.Stdout), zapcore.DebugLevel)) // 终端输出core
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	zapLog = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(zapLog) // 替换zap包中全局的logger实例 后续在其他包中只需要使用zap.l()即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = getEncodeTime
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderCfg)
}

func getLogWrite(maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFile(),
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 获取日志文件名
func getLogFile() string {
	fileFormat := time.Now().Format(Config.LogConfig.FileFormat)
	fileName := strings.Join([]string{
		Config.LogConfig.FileName + "_" + fileFormat,
		"log"}, ".")
	return path.Join(Config.LogConfig.Path, fileName)
}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zapLog.Info(
			path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost))
	}
}

// GinRecovery 中间件
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zapLog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					zapLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zapLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
