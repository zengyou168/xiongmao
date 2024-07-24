// Package log SQL日志记录器
package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

// ZapGormLogger 是使用 zap 的 GORM 的自定义记录器
type ZapGormLogger struct {
	logger *zap.Logger
}

func ZapSqlLog() *ZapGormLogger {

	// 自定义时间编码器
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// 创建自定义的 JSON EncoderConfig
	jsonEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder, // 自定义时间编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建 JSON Encoder
	jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig)

	path := "logger/sql"

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
	core := zapcore.NewCore(jsonEncoder, fileSyncer, zap.DebugLevel)

	// 创建 Logger
	loggerNew := zap.New(core)

	// 配置 GORM 日志记录器
	gormLogger := &ZapGormLogger{
		logger: loggerNew,
	}

	return gormLogger
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Sugar().Infof(msg, data...)
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Sugar().Warnf(msg, data...)
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Sugar().Errorf(msg, data...)
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.logger.Sugar().Errorf("err=%v elapsed=%v sql=%s rows=%d", err, elapsed, sql, rows)
	} else {
		l.logger.Sugar().Infof("elapsed=%v sql=%s rows=%d", elapsed, sql, rows)
	}
}
