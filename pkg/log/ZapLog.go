// Package log SQL日志记录器
package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
	"xiongmao/config"
)

var SugarVar *zap.SugaredLogger

func Init() {

	// 创建 JSON 编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	path := config.LogVar.Path + "/panda"

	err := os.MkdirAll(path, 0755)

	// 创建日志目录
	if err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// 动态生成文件名
	currentTime := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf(path+"/%s.json", currentTime)

	// 创建文件输出
	file, _ := os.Create(fileName)
	fileSyncer := zapcore.AddSync(file)

	// 创建 Core
	core := zapcore.NewCore(encoder, fileSyncer, zap.DebugLevel)

	// 创建 Logger
	zapLogger := zap.New(core)

	// 获取 Sugar Logger 用于更方便的日志打印
	SugarVar = zapLogger.Sugar()

	defer SugarVar.Sync()
}
