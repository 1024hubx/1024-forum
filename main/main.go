package main

import (
	"context"
	"fmt"
	"forum/models"
	"forum/routers"

	"forum/util/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//初始化定时任务
	//util.InitCron()
	config.InitConf()

	if config.ServerConfig.ServerRunEnv == "master" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	models.InitMysql()
	// models.InitRedis()
	router := routers.InitRouter()
	//初始化日志

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerConfig.ServerHttpPort),
		Handler:        router,
		ReadTimeout:    config.ServerConfig.ServerReadTimeOut,
		WriteTimeout:   config.ServerConfig.ServerWriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
