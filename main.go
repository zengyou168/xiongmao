/**
 *
 * @Author: ZengYou
 * @Date: 2024/7/23
 */
package main

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"panda/initialize"
)

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

func main() {

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	// Create custom GORM logger
	customLogger := &initialize.CustomLogger{zapLogger: zapLogger}

	// Initialize GORM with custom logger

	// 连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/panda?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
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
		log.Fatal("failed to create user: ", err)
	}

	fmt.Printf("User created: %+v\n", user)
}
