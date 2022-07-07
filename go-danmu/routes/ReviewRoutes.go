package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetReviewRoutes(route *gin.RouterGroup) {
	review := route.Group("/review")
	{
		//需要用户登录
		userAuth := review.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/status", controller.GetVideoStatus)                      //获取视频状态
			userAuth.GET("/resource/upload/list", controller.GetUploadResourceList) // 获取用户上传资源
			userAuth.POST("/submit", controller.SubmitReview)                       //提交审核
		}

		//管理员接口
		auditorAuth := review.Group("")
		auditorAuth.Use(middleware.AuthMiddleware(common.Auditor))
		{
			auditorAuth.GET("/resource", controller.GetReviewVideoByID)  // 获取待审核视频资源
			auditorAuth.GET("/list", controller.GetReviewVideoList)      //获取审核列表
			auditorAuth.POST("/video/add", controller.ReviewVideo)       //审核视频
			auditorAuth.POST("/resource/add", controller.ReviewResource) //审核资源
		}
	}
}
