package logmiddleware

import (
	"cmsserver/util"
	"time"

	"github.com/gin-gonic/gin"
)

//打印gin系统日志
func GinLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		end := time.Now()
		time := end.Sub(start)
		msg := end.Format("2006/01/02 - 15:04:05") + "-" + c.Request.Method + "-" + c.ClientIP() + c.Request.URL.Path + "-" + util.ShortDur(time)
		util.FmtSyslog(msg)
	}
}

//打印客户登录系统日志
func ClientLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg string
		start := time.Now()
		end := time.Now()
		time := end.Sub(start)
		//手机+孩子昵称+ip+操作+登陆/注册
		// 手机 孩子昵称 ip 操作  登录了易编学App
		mobile, _ := c.Get("mobile")
		child_name, _ := c.Get("ChildName")
		operating, _ := c.Get("operating")

		t, _ := c.Get("type")
		if t.(int64) == 1 {
			msg = mobile.(string) + "-" + child_name.(string) + "-" + string(c.ClientIP()) + "-" + operating.(string) + "app" + util.ShortDur(time)
			if t.(int64) == 2 {
				msg = mobile.(string) + "-" + child_name.(string) + "-" + string(c.ClientIP()) + "-" + operating.(string) + "官网" + util.ShortDur(time)
			}
			util.FmtClientLog(msg)
		}
	}
}

//打印员工登录系统日志
func StaffLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		end := time.Now()
		time := end.Sub(start)
		mobile, _ := c.Get("mobile")
		name, _ := c.Get("name")
		message, _ := c.Get("message")
		//手机
		//用户名
		//ip
		//干了什么。。（模块+动作+对象）
		msg := mobile.(string) + "-" + name.(string) + "-" + string(c.ClientIP()) + "-" + message.(string) + "-" + util.ShortDur(time)
		util.FmtStaffLog(msg)
	}
}
