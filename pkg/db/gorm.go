// Package initialize 数据库驱动
package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"panda/config"
	"panda/pkg/log"
	"strings"
)

func Init() *gorm.DB {

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
	db.AutoMigrate(&User{})

	///addTableComment(db, "users", "用户表")

	// 创建 UserService 实例
	userService := &UserService{DB: db}

	// 创建新用户
	userCreate := UserCreate{
		Username: "john_doe",
		Password: "securepassword",
	}

	user, err := userService.CreateUser(userCreate)

	if err != nil {
		// log.Sugar.Errorf("failed to create user：%v", err)
	}

	fmt.Printf("User created: %+v\n", user)

	if err != nil {

		//  log.Sugar.Errorf("database start error：%v", err)

		os.Exit(0)

		return nil
	}

	return db
}

// BeforeCreate 在创建记录之前生成没有破折号的 UUID
func (user *User) BeforeCreate(*gorm.DB) (err error) {

	// 生成没有破折号的 UUID
	uuidWithHyphens := uuid.NewString()

	user.ID = strings.ReplaceAll(uuidWithHyphens, "-", "")

	return
}

// addTableComment 为指定表添加备注
func addTableComment(db *gorm.DB, tableName, comment string) {
	sql := fmt.Sprintf("COMMENT ON TABLE %s IS '%s';", tableName, comment)
	db.Exec(sql)
	/*if err != nil {
		log.Fatalf("添加表备注失败: %v", err)
	}*/
}

// User 模型
type User struct {
	ID       string `gorm:"type:char(32);primaryKey"`
	Username string `gorm:"type:varchar(100);comment:用户姓名"`
	Password string
}

// UserCreate 用于接收创建用户请求的数据
type UserCreate struct {
	Username string `gorm:"type:varchar(100);comment:用户姓名"`
	Password string
}

// UserService 处理用户相关操作
type UserService struct {
	DB *gorm.DB
}

// CreateUser 创建新用户
func (service *UserService) CreateUser(userCreate UserCreate) (*User, error) {
	user := User{
		Username: userCreate.Username,
		Password: userCreate.Password,
	}
	if err := service.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
