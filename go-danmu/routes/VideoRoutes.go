package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetVideoRoutes(route *gin.RouterGroup) {
	video := route.Group("/video")
	{
		video.GET("/search", controller.SearchVideo)              //搜索视频
		video.GET("/get", controller.GetVideoByID)                //获取视频详情
		video.GET("/recommend/get", controller.GetRecommendVideo) //获取推荐视频
		video.GET("/list/get", controller.GetVideoList)           //获取视频列表
		video.GET("/user/get", controller.GetVideoListByUID)      //获取用户视频列表
		video.GET("/online/ws", controller.GetRoomConnect)        //在线人数连接

		//需要用户登录
		userAuth := video.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/upload/get", controller.GetUploadVideoList)              //获取个人上传的视频
			userAuth.POST("/modify/info", controller.ModifyVideoInfo)               //更新视频信息
			userAuth.POST("/delete", controller.DeleteVideo)                        //删除视频
			userAuth.POST("/resource/delete", controller.DeleteResource)            //删除资源
			userAuth.POST("/upload/info", controller.UploadVideoInfo)               //上传视频信息
			userAuth.POST("/resource/title/modify", controller.ModifyResourceTitle) //修改资源标题
		}

		//管理员接口
		adminAuth := video.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.GET("/list/admin", controller.AdminGetVideoList)   //获取视频列表
			adminAuth.POST("/delete/admin", controller.AdminDeleteVideo) //删除视频
		}

	}
}
