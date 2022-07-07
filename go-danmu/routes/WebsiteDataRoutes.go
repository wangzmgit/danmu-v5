package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetWebsiteDataRoutes(route *gin.RouterGroup) {
	wbsiteData := route.Group("/website/data")
	{
		//管理员接口
		adminAuth := wbsiteData.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.GET("/total", controller.GetTotalWebsiteData)   //获取总数据
			adminAuth.GET("/recent", controller.GetRecentWebsiteData) //获取近五天数据
		}

	}
}
