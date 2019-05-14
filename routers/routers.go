package routers

import (
	"cmsserver/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(errorMiddleware()) //全局错误处理

	//增加对NoMethod和NoRoute的处理，注意放在路由底部
	router.NoMethod(func(c *gin.Context) {
		c.JSON(
			http.StatusMethodNotAllowed,
			gin.H{
				"error_code": 405,
				"message":    "方法不允许",
			})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,
			gin.H{
				"error_code": 404,
				"message":    "资料不存在",
			})
	})

	return router
}

//全局错误处理
func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if nerr := recover(); nerr != nil {
				err := nerr.(error)
				util.FmtSyslog(err.Error())
			}
		}()
		c.Next()
	}
}
