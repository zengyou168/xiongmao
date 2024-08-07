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
	Jwt      jwtData      `yaml:"jwt"`
	Mqtt     mqttData     `yaml:"mqtt"`
	Log      logData      `yaml:"log"`
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
	Zone   string `yaml:"zone"`
}

type jwtData struct {
	Key string `yaml:"key"`
}

type mqttData struct {
	Server   string `yaml:"server"`
	ClientID string `yaml:"clientID"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Topic    string `yaml:"topic"`
}

type logData struct {
	Path     string `yaml:"path"`
	Xiongmao bool   `yaml:"xiongmao"`
	Sql      bool   `yaml:"sql"`
}

var configDataVar *configData
var ServerVar *serverData
var DatabaseVar *databaseData
var JwtKeyVar []byte
var MqttVar *mqttData
var LogVar *logData

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
	JwtKeyVar = []byte(configDataVar.Jwt.Key)
	MqttVar = &configDataVar.Mqtt
	LogVar = &configDataVar.Log
}
