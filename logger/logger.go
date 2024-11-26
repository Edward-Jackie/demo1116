package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// GLog 全局日志
var GLog *zap.Logger

// InitLogger
// 初始化日志服务
func InitLogger(filepath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEncodeer := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var allcode []zapcore.Core
	allcode = append(allcode, core)
	allcode = append(allcode, zapcore.NewCore(consoleEncodeer, consoleDebug, p))
	c := zapcore.NewTee(allcode...)
	//zap.AddCaller() //添加将调用函数信息记录到日志中的功能。
	GLog = zap.New(c, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:  filepath,
		MaxSize:   10240,
		MaxAge:    7,
		Compress:  true,
		LocalTime: true,
	}

	c := cron.New()
	c.AddFunc("0 0 0 1/1 * ?", func() {
		lumberJackLogger.Rotate()
	})
	c.Start()
	return zapcore.AddSync(lumberJackLogger)
}

// GormLog Orm框架Log
var GormLog *GormLogger

type GormLogger struct{}

func (GormLogger) Println(v ...interface{}) {
	GLog.Info(fmt.Sprintf("ORM： \n%v", v))
}

// GinLog Web服务日志
var GinLog GinLogger

type GinLogger struct{}

func (GinLogger) Write(p []byte) (n int, err error) {
	GLog.Info(fmt.Sprintf("%s\n", p))
	return n, err
}
