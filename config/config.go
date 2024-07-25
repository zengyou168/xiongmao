// Package config 配置参数
package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigData struct {
	Server   ServerData   `yaml:"server"`
	Database DatabaseData `yaml:"database"`
	Log      LogData      `yaml:"log"`
}

type ServerData struct {
	Port int `yaml:"port"`
}

type DatabaseData struct {
	Driver string `yaml:"driver"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Addr   string `yaml:"addr"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbName"`
}

type LogData struct {
	Path  string `yaml:"path"`
	Panda bool   `yaml:"panda"`
	Sql   bool   `yaml:"sql"`
}

var configDataVar *ConfigData
var ServerVar *ServerData
var DatabaseVar *DatabaseData
var LogVar *LogData

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
}
