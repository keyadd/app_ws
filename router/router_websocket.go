package router

import (
	v1 "app_ws/api/v1"
	"app_ws/global"
	"app_ws/initialize/wsmanage"
	"github.com/gin-gonic/gin"
)

func InitWebSocketRouter(Router *gin.RouterGroup) {
	newServer := wsmanage.NewServer()
	UserRouter := Router.Group("")
	{

		global.GVA_LOG.Info("register wsmanage handler")

		//配置路由
		newServer.AddRouter("ping", v1.PingRouter{}) //ping保持连接
		newServer.AddRouter("hello", v1.HelloRouter{})
		newServer.AddRouter("echo", v1.EchoRouter{})
		UserRouter.GET("/ws", newServer.Serve)
	}
}
