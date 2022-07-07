package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetDanmakuRoutes(route *gin.RouterGroup) {
	danmaku := route.Group("/danmaku")
	{

		danmaku.GET("/list", controller.GetDanmaku)

		//需要用户登录
		userAuth := danmaku.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.POST("/send", controller.SendDanmaku)
		}
	}
}
