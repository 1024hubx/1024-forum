package config

import (
	"log"
	"os"
	"time"

	"github.com/casbin/casbin"
	"github.com/facebookgo/inject"
	"github.com/go-ini/ini"
)

var (
	cfg      *ini.File
	Gra      *inject.Graph
	Enforcer *casbin.Enforcer

	configPath     string
	configAuthPath string
	GOENV          string

	AppConfig     = &App{}
	ServerConfig  = &Server{}
	DbMysqlConfig = &DbMysql{}
	RedisConfig   = &Redis{}
	Mes           = &Message{}
	Log           = &Logs{}
	WxConfig      = &Wx{}
	GameConfig    = &Game{}
)

// 加载命令行文件
func init() {

	GOENV = os.Getenv("GOENV")

	if GOENV == "master" {
		configPath = "config/master.ini"
		configAuthPath = "config/master_auth_model.conf"
	} else {
		configPath = "config/dev.ini"
		configAuthPath = "config/dev_auth_model.conf"
	}
}

// 启动配置项
func InitConf() {

	var err error

	cfg, err = ini.Load(configPath) //加载配置文件
	if err != nil {
		log.Fatalf("Fail to parse 'config/*.ini': %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	ServerConfig.ServerReadTimeOut = ServerConfig.ServerReadTimeOut * time.Second
	ServerConfig.ServerWriteTimeOut = ServerConfig.ServerWriteTimeOut * time.Second

	mapTo("db.mysql", DbMysqlConfig)
	DbMysqlConfig.SetConnMaxLifetime = DbMysqlConfig.SetConnMaxLifetime * time.Second

	mapTo("redis", RedisConfig)
	RedisConfig.RedisDialTimeout = RedisConfig.RedisDialTimeout * time.Second
	RedisConfig.RedisReadTimeout = RedisConfig.RedisReadTimeout * time.Second
	RedisConfig.RedisWriteTimeout = RedisConfig.RedisWriteTimeout * time.Second
	mapTo("Message", Mes)
	mapTo("LogUrl", Log)

	mapTo("Wx", WxConfig)
	mapTo("Game", GameConfig)

	Gra = new(inject.Graph) //加载auth配置文件

}

// 配置文件结构转结构体
func mapTo(section string, v interface{}) {

	err := cfg.Section(section).MapTo(v)

	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
