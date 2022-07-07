package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

//点赞收藏
func GetArchiveRoutes(route *gin.RouterGroup) {
	archive := route.Group("/archive")
	{

		//需要用户登录
		userAuth := archive.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/video", controller.GetArchiveByVID)               //获取视频点赞收藏信息
			userAuth.GET("/collect/status", controller.GetCollectStatus)     //获取收藏状态
			userAuth.GET("/collection/list", controller.GetCollectionList)   //获取收藏夹列表
			userAuth.POST("/collection/add", controller.AddCollection)       //添加收藏夹
			userAuth.POST("/collection/modify", controller.ModifyCollection) //修改收藏夹
			userAuth.POST("/collection/delete", controller.DeleteCollection) //删除收藏夹
			userAuth.POST("/collect/add", controller.Collect)                //收藏
			userAuth.POST("/like/add", controller.Like)                      //点赞
			userAuth.POST("/like/cancel", controller.Dislike)                //取消赞
		}

		uidAuth := archive.Group("")
		uidAuth.Use(middleware.UidMiddleware())
		{
			uidAuth.GET("/collection/info", controller.GetCollectionByID) //获取收藏夹信息
			uidAuth.GET("/collect/get", controller.GetCollectVideo)       //获取收藏的视频
		}
	}
}
