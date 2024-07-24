// Package log SQL日志记录器
package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

// Sugar 定义全局 logger 变量
var Sugar *zap.SugaredLogger

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

    // 创建文件输出
    logFile, err := os.Create("logger/app.log")

    if err != nil {
        panic(err)
    }

    fileSyncer := zapcore.AddSync(logFile)

    // 创建 Core
    core := zapcore.NewCore(encoder, fileSyncer, zap.DebugLevel)

    // 创建 Logger
    logger := zap.New(core)

    // 获取 Sugar Logger 用于更方便的日志打印
    Sugar = logger.Sugar()

    defer Sugar.Sync()
}
