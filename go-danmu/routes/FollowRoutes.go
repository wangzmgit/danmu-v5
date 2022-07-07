package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

//关注
func GetFollowRoutes(route *gin.RouterGroup) {
	follow := route.Group("/follow")
	{
		follow.GET("/following", controller.GetFollowingByID) //关注列表
		follow.GET("/followers", controller.GetFollowersByID) //粉丝列表
		follow.GET("/count", controller.GetFollowCount)       //获取关注数和粉丝数

		//需要用户登录
		userAuth := follow.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/status", controller.GetFollowStatus) //获取关注状态
			userAuth.POST("/add", controller.Following)         //关注
			userAuth.POST("/cancel", controller.UnFollow)       //取消关注
		}
	}
}
