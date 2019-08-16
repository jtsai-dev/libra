/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 13:15:17
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-13 18:49:38
 */
package pkg

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

type Config struct {
	App      App      `yaml:"App"`
	Wechat   Wechat   `yaml:"Wechat"`
	Server   Server   `yaml:"Server"`
	Database Database `yaml:"Database"`
	Redis    Redis    `yaml:"Redis"`
	Log      Log      `yaml:"Log"`
}

var Configs = &Config{}

type App struct {
	TokenExpiredSeconds int      `yaml:"TokenExpiredSeconds"`
	PageSize            int      `yaml:"PageSize"`
	JwtSecret           string   `yaml:"JwtSecret"`
	ImageSavePath       string   `yaml:"ImageSavePath"`
	ImageMaxSize        int      `yaml:"ImageMaxSize"`
	ImageAllowExts      []string `yaml:"ImageAllowExts"`
	ExportPath          string   `yaml:"ExportPath"`
}

type Wechat struct {
	AppId         string `yaml:"AppId"`
	AppSecret     string `yaml:"AppSecret"`
	DefaultAvatar string `yaml:"DefaultAvatar"`
}

type Server struct {
	RunMode string `yaml:"RunMode"`
	Port    int    `yaml:"Port"`
}

type Database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
	Charset  string `yaml:"Charset"`
}

type Redis struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	DBIndex  int    `yaml:"DBIndex"`
}

type Log struct {
	Path            string `yaml:"Path"`
	TimestampFormat string `yaml:"TimestampFormat"`
}

// Setup initialize the configuration instance
func Setup() {
	var err error
	configFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to load 'conf/app.yaml': %v", err)
	}
	err = yaml.Unmarshal(configFile, &Configs)
	if err != nil {
		log.Fatalf("yamlFile.Unmarshal err %v ", err)
	}
}
