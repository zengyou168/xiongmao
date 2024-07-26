// Package config 配置参数
package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type configData struct {
	Server   serverData   `yaml:"server"`
	Database databaseData `yaml:"database"`
	Log      logData      `yaml:"log"`
	Casbin   casbinData   `yaml:"casbin"`
}

type serverData struct {
	Port int `yaml:"port"`
}

type databaseData struct {
	Driver string `yaml:"driver"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Addr   string `yaml:"addr"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbName"`
}

type logData struct {
	Path  string `yaml:"path"`
	Panda bool   `yaml:"panda"`
	Sql   bool   `yaml:"sql"`
}

type casbinData struct {
	SecretKey string `yaml:"secretKey"`
}

var configDataVar *configData
var ServerVar *serverData
var DatabaseVar *databaseData
var LogVar *logData
var CasbinVar *casbinData

func Init() {

	data, err := os.ReadFile("config/config.yaml")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(data, &configDataVar)

	if err != nil {
		fmt.Println(err)
		return
	}

	ServerVar = &configDataVar.Server
	DatabaseVar = &configDataVar.Database
	LogVar = &configDataVar.Log
	CasbinVar = &configDataVar.Casbin
}
