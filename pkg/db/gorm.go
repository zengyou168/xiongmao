// Package initialize 数据库驱动
package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "panda/config"
    "panda/internal/model"
    "panda/pkg/log"
)

var Gorm *gorm.DB

func Init() {

    database := config.DatabaseVar
    option := &gorm.Config{}

    var dialector gorm.Dialector

    if database.Driver == "postgres" {
        dialector = postgres.Open(fmt.Sprintf(
            "user=%s password=%s host=%s port=%d dbname=%s",
            database.User,
            database.Passwd,
            database.Addr,
            database.Port,
            database.DBName,
        ))
    } else {
        dialector = mysql.Open(fmt.Sprintf(
            "%s:%s@tcp(%s:%d)/%s",
            database.User,
            database.Passwd,
            database.Addr,
            database.Port,
            database.DBName,
        ))
    }

    if config.LogVar.Panda {
        option.Logger = log.ZapSqlLog()
    }

    // 连接数据库
    db, err := gorm.Open(dialector, option)

    if err != nil {
        //  SugarGlobalVar.Errorf("failed to connect database：%v", err)
    }

    // 自动迁移数据库
    db.AutoMigrate(&model.User{})

    addTableComment(db, "users", "用户表")

    if err != nil {
        log.SugarVar.Errorf("database start error：%v", err)
    }

    Gorm = db
}

// addTableComment 为指定表添加备注
func addTableComment(db *gorm.DB, tableName, comment string) {

    sql := fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", tableName, comment)

    db.Exec(sql)
}
