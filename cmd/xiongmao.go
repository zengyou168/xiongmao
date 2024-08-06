// 主程序入口
package main

import (
	"xiongmao/config"
	"xiongmao/internal/router"
	"xiongmao/pkg/db"
	"xiongmao/pkg/log"
	"xiongmao/pkg/mqtt"
)

func main() {

	// 实例化项目配置
	config.Init()

	// 实例化 zap 日志
	log.Init()

	// 实例化 GORM
	db.Init()

	// 实例化 mqtt，项目用不到 MQTT 就注释掉
	mqtt.Init()

	// 实例化 router，这个一定要放在最后
	router.Init()
}
