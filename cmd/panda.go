// 主程序入口
package main

import (
    "panda/pkg/db"
    "panda/pkg/log"
)

func main() {

    // 实例化 zap 日志
    log.Init()

    // 使用自定义记录器初始化 GORM
    db.CustomGorm()

}
