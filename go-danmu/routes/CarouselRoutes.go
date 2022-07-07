package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

//轮播图
func GetCarouselRoutes(route *gin.RouterGroup) {
	carousel := route.Group("/carousel")
	{
		carousel.GET("/get", controller.GetCarousel)

		//需要管理员登录
		adminAuth := carousel.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.POST("/upload/info", controller.UploadCarouselInfo) //上传轮播图信息
			adminAuth.POST("/delete", controller.DeleteCarousel)          //删除轮播图
		}
	}
}
