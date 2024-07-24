// Package initialize 数据库驱动
package initialize

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"panda/application"
	"strings"
	"time"
)

func CustomGorm(c *CustomLogger) *gorm.DB {

	data, err := os.ReadFile("application/application.yaml")

	if err != nil {
		c.Error(context.TODO(), "Error reading YAML file: %s\n", err)
	}

	var config application.Application

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		c.Error(context.TODO(), "Error parsing YAML file: %s\n", err)
	}

	database := config.Database
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

	if database.Log {
		option.Logger = c
	}

	// 连接数据库
	db, err := gorm.Open(dialector, option)

	if err != nil {
		c.Error(context.TODO(), "failed to connect database: ", err)
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
		c.Error(context.TODO(), "failed to create user: ", zap.Error(err))
	}

	fmt.Printf("User created: %+v\n", user)

	return GormInit(db, err)
}

func GormInit(db *gorm.DB, err error) *gorm.DB {
	if err != nil {
		log.Error("database start error", zap.Error(err))
		os.Exit(0)
		return nil
	}
	sqlDB, _ := db.DB()

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
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
