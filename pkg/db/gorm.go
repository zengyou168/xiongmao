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
	err = db.AutoMigrate(
		&model.Admin{},
		&model.Role{},
	)

	if err != nil {
		log.SugarVar.Errorf("自动迁移数据库：%v", err)
	}

	addTableComments(
		db,
		&model.Admin{},
		&model.Role{},
	)

	Gorm = db
}

func addTableComments(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		addTableComment(db, model)
	}
}

// addTableComment 为指定表添加备注
func addTableComment(db *gorm.DB, model interface{}) {

	tableName, ok := model.(interface {
		TableName() string
	})

	if !ok {
		return
	}

	tableComment, ok := model.(interface {
		TableComment() string
	})

	if !ok {
		return
	}

	database := config.DatabaseVar
	name := tableName.TableName()
	comment := tableComment.TableComment()
	sql := ""

	if database.Driver == "postgres" {
		sql = fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", name, comment)
	} else {
		sql = fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", name, comment)
	}

	db.Exec(sql)
}
