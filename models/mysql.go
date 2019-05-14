package models

import (
	"fmt"
	"forum/util/config"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// engine package db

var (
	engine *xorm.Engine
	err    error
)

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
func InitMysql() {

	engine, err = xorm.NewEngine(config.DbMysqlConfig.DbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		config.DbMysqlConfig.DbUser,
		config.DbMysqlConfig.DbPassword,
		config.DbMysqlConfig.DbHost,
		config.DbMysqlConfig.DbName,
	))
	// f, err := os.Create(config.Log.Sqllog)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") //上海时区

	if err = engine.Ping(); err != nil { //数据库 ping
		log.Fatalf("database ping err: %s", err.Error())
	}

	if config.ServerConfig.ServerRunEnv == "master" { //根据运行环境判断日志模式
		engine.ShowSQL(false)
	} else {
		engine.ShowSQL(true)
	}

	engine.SetMaxIdleConns(config.DbMysqlConfig.DbMaxIdleConn) //连接池的空闲数大小
	engine.SetMaxOpenConns(config.DbMysqlConfig.DbMaxOpenConn) //最大打开连接数

	timer := time.NewTicker(time.Minute * 30) //定时器 ping
	go func(x *xorm.Engine) {
		for range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("数据库连接错误: %#v\n", err.Error())
			}
		}
	}(engine)

}
