/**
 * 数据库驱动
 * @Author: ZengYou
 * @Date: 2024/7/23
 */
package initialize

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func CustomGorm(c *CustomLogger) *gorm.DB {

	// 连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/panda?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: c,
	})

	if err != nil {
		c.Error(context.TODO(), "failed to connect database: ", err)
	}

	// 自动迁移数据库
	db.AutoMigrate(&User{})

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

// User 模型
type User struct {
	gorm.Model
	Username string
	Password string
}

// UserCreate 用于接收创建用户请求的数据
type UserCreate struct {
	Username string
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
