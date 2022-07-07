package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetConfigRoutes(route *gin.RouterGroup) {
	wbsiteData := route.Group("/config")
	{
		//管理员接口
		rootAuth := wbsiteData.Group("")
		rootAuth.Use(middleware.AuthMiddleware(common.Root))
		{
			rootAuth.GET("/server/info", controller.GetServerInfo)       //获取服务信息
			rootAuth.GET("/database/get", controller.GetDatabaseConfig)  //获取数据库配置
			rootAuth.GET("/email/get", controller.GetEmailConfig)        //获取邮箱配置
			rootAuth.GET("/oss/get", controller.GetOssConfig)            //获取oss配置
			rootAuth.GET("/other/get", controller.GetOtherConfig)        //获取其他配置
			rootAuth.POST("/other/set", controller.SetOtherConfig)       //修改其他配置
			rootAuth.POST("/oss/set", controller.SetOssConfig)           //修改oss配置
			rootAuth.POST("/email/set", controller.SetEmailConfig)       //修改邮箱配置
			rootAuth.POST("/database/set", controller.SetDatabaseConfig) //修改数据库配置
		}

	}
}
