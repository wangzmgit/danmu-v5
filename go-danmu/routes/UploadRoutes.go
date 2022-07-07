package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

//点赞收藏
func GetUploadRoutes(route *gin.RouterGroup) {
	upload := route.Group("/upload")
	{
		//需要用户登录
		userAuth := upload.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.POST("/avatar", controller.UploadAvatar)
			userAuth.POST("/cover", controller.UploadCover)
			userAuth.POST("/video", controller.UploadVideo)
		}

		//管理员接口
		adminAuth := upload.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.POST("/carousel", controller.UploadCarousel)
		}
	}
}
