/**
 * 日志记录器
 * @Author: ZengYou
 * @Date: 2024/7/23
 */
package initialize

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

// CustomLogger 是一个使用zap的自定义GORM记录器
type CustomLogger struct {
	ZapLogger *zap.Logger
}

// LogMode 设置日志级别
func (c *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	return c // 这可以扩展以支持不同的日志级别
}

// Info 记录信息级别消息
func (c *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	c.ZapLogger.Sugar().Infof(msg, data...)
}

// Warn 日志警告级别消息
func (c *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	c.ZapLogger.Sugar().Warnf(msg, data...)
}

// Error 记录错误级别消息
func (c *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	c.ZapLogger.Sugar().Errorf(msg, data...)
}

// Trace 记录 SQL 查询
func (c *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		c.ZapLogger.Sugar().Errorw("SQL Error",
			"duration", elapsed,
			"rows", rows,
			"sql", sql,
			"error", err,
		)
	} else {
		c.ZapLogger.Sugar().Infow("SQL Query",
			"duration", elapsed,
			"rows", rows,
			"sql", sql,
		)
	}
}
