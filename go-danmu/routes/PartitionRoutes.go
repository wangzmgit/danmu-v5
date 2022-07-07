package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetPartitionRoutes(route *gin.RouterGroup) {
	partition := route.Group("/partition")
	{
		partition.GET("/list", controller.GetPartitionList) //获取分区列表
		partition.GET("/all", controller.GetAllPartition)   //获取所有分区

		//管理接口
		rootAuth := partition.Group("")
		rootAuth.Use(middleware.AuthMiddleware(common.Root))
		{
			rootAuth.POST("/add", controller.AddPartition)       //添加分区
			rootAuth.POST("/delete", controller.DeletePartition) //删除分区
		}

	}
}
