package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

//评论回复
func GetCommentRoutes(route *gin.RouterGroup) {
	comment := route.Group("/comment")
	{
		comment.GET("/list", controller.GetCommentAndReplyList) //获取评论列表
		comment.GET("/reply/list", controller.GetReplyDetails)  //获取回复列表
		//需要用户登录
		userAuth := comment.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/list/manage", controller.GetCommentList) //获取当前视频下的评论
			userAuth.POST("/add", controller.AddComment)            //发布评论
			userAuth.POST("/delete", controller.DeleteComment)      //删除评论
		}
	}
}
