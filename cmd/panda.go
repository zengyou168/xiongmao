// 主程序入口
package main

import (
	"panda/config"
	"panda/pkg/db"
	"panda/pkg/log"
)

func main() {

	// 实例化项目配置
	config.Init()

	// 实例化 zap 日志
	log.Init()

	// 实例化 GORM
	db.Init()

}
