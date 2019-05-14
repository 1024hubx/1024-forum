package config

import "time"

type (
	// app设置
	App struct {
		AppBussines string
	}
	// 服务
	Server struct {
		ServerRunEnv       string
		ServerHttpPort     int
		ServerReadTimeOut  time.Duration
		ServerWriteTimeOut time.Duration
		ServerJwtSecret    string
	}
	// mysql数据库
	DbMysql struct {
		DbType             string
		DbUser             string
		DbPassword         string
		DbHost             string
		DbName             string
		DbTablePrefix      string
		SetConnMaxLifetime time.Duration
		DbMaxIdleConn      int
		DbMaxOpenConn      int
	}

	// redis
	Redis struct {
		RedisHost         string
		RedisPassword     string
		RedisDialTimeout  time.Duration
		RedisReadTimeout  time.Duration
		RedisWriteTimeout time.Duration
		RedisPoolSize     int
		RedisMinIdleConns int
	}
	// 短信
	Message struct {
		Name string
		Pwd  string
		Url  string
	}
	// 日志
	Logs struct {
		Syslog    string
		Clientlog string
		Stafflog  string
		Sqllog    string
	}
)
