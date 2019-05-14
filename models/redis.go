package models

import (
	"encoding/json"
	"fmt"
	"forum/util"

	"forum/util/config"

	"github.com/go-redis/redis"
	"github.com/lunny/log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.RedisConfig.RedisHost,         //主机
		Password:     config.RedisConfig.RedisPassword,     //密码
		DB:           0,                                    //默认库
		PoolSize:     config.RedisConfig.RedisPoolSize,     //最大连接数
		MinIdleConns: config.RedisConfig.RedisMinIdleConns, //默认连接数
		DialTimeout:  config.RedisConfig.RedisDialTimeout,  //拨号超时
		WriteTimeout: config.RedisConfig.RedisWriteTimeout, //写超时
		ReadTimeout:  config.RedisConfig.RedisReadTimeout,  //读超时
	})

	pong, err := RedisClient.Ping().Result() //ping连接
	if err != nil {
		log.Fatalf("redis setup faile: %s\n", err.Error())
	}

	fmt.Printf("redis setup success: %s\n", pong)
}

//获取redis中的对象
func GetRedis(key string, v interface{}) { //v 返回的数据
	val, err := RedisClient.Get(key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), v)
	}
}

//吧对象放入redis
func SetRedis(key string, v interface{}) {
	b, err := json.Marshal(v)
	if err == nil {
		err = RedisClient.Set(key, b, 0).Err()
	}
}

//删除redis中的对象
func DelRedis(key string, v interface{}) {

	err = RedisClient.Del(key).Err()
	if err != nil {
		util.CheckErr(err)
	}
}
