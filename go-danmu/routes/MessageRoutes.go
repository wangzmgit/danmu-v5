package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetMessageRoutes(route *gin.RouterGroup) {
	message := route.Group("/message")
	{
		message.GET("/ws", controller.GetMsgConnect)
		message.GET("/announce/list", controller.GetAnnounce) //获取公告

		//需要用户登录
		userAuth := message.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/list", controller.GetMessageList)       //获取消息列表
			userAuth.GET("/details", controller.GetMessageDetails) //获取消息详情
			userAuth.POST("/send", controller.SendMessage)         //发送消息
			userAuth.POST("/read", controller.ReadMessageService)  //已读消息
		}

		//管理接口
		adminAuth := message.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.POST("/announce/add", controller.AddAnnounce)       //添加公告
			adminAuth.POST("/announce/delete", controller.DeleteAnnounce) //删除公告
		}

	}
}
