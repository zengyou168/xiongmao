// Package db 数据库驱动
package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"xiongmao/config"
	"xiongmao/internal/model"
	"xiongmao/pkg/log"
)

var Gorm *gorm.DB

func Init() {

	database := config.DatabaseVar
	option := &gorm.Config{}

	var dialector gorm.Dialector

	if database.Driver == "postgres" {
		dialector = postgres.Open(fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s TimeZone=%s",
			database.User,
			database.Passwd,
			database.Addr,
			database.Port,
			database.DBName,
			database.Zone,
		))
	} else {
		dialector = mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?loc=%s",
			database.User,
			database.Passwd,
			database.Addr,
			database.Port,
			database.DBName,
			url.QueryEscape(database.Zone),
		))
	}

	if config.LogVar.Xiongmao {
		option.Logger = log.ZapSqlLog()
	}

	// 连接数据库
	db, err := gorm.Open(dialector, option)

	if err != nil {
		log.SugarVar.Errorf("连接数据库：%v", err)
	}

	// 自动迁移数据库
	err = db.AutoMigrate(&model.Admin{})

	if err != nil {
		log.SugarVar.Errorf("自动迁移数据库：%v", err)
	}

	addTableComment(db, "admin", "管理员")

	Gorm = db
}

// addTableComment 为指定表添加备注
func addTableComment(db *gorm.DB, tableName, comment string) {

	sql := fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", tableName, comment)

	db.Exec(sql)
}
