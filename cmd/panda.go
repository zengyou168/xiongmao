// 主程序入口
package main

import (
	"go.uber.org/zap"
	"panda/pkg/db"
	"panda/pkg/log"
)

func main() {

	zapLogger, _ := zap.NewProduction()

	defer func(zapLogger *zap.Logger) {
		_ = zapLogger.Sync()
	}(zapLogger)

	// 创建自定义 GORM 记录器
	customLogger := &log.CustomLogger{ZapLogger: zapLogger}

	// 使用自定义记录器初始化 GORM
	db.CustomGorm(customLogger)

}
