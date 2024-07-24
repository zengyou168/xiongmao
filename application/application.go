// Package application 配置参数
package application

type Application struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Driver string `yaml:"driver"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Addr   string `yaml:"addr"`
		Port   int    `yaml:"port"`
		DBName string `yaml:"dbName"`
		Log    bool   `yaml:"log"`
	} `yaml:"database"`
}
