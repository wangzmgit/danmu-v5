package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/controller"
	"kuukaa.fun/danmu-v5/middleware"
)

func GetUserRoutes(route *gin.RouterGroup) {
	user := route.Group("/user")
	{
		user.GET("/info/other", controller.GetUserInfoByID)      //获取其他用户信息
		user.GET("/info/uid", controller.GetUIDByName)           //通过用户名获取用户id
		user.POST("/register", controller.Register)              //用户注册
		user.POST("/register/code", controller.SendRegisterCode) //发送注册验证码
		user.POST("/login", controller.Login)                    //用户登录
		user.POST("/root/login", controller.RootLogin)           //超级管理员登录
		// 通过RefreshToken获取AccessToken
		user.GET("/token/refresh", middleware.RefreshTokenMiddleware(), controller.GetAccessToken)

		//需要用户登录
		userAuth := user.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/info/get", controller.UserInfo)       //用户获取个人信息
			userAuth.POST("/info/modify", controller.ModifyInfo) //用户修改个人信息
		}

		//管理员接口
		adminAuth := user.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.GET("/search/admin", controller.AdminSearchUser)           //管理员搜索
			adminAuth.GET("/list", controller.GetUserList)                       //获取用户列表
			adminAuth.POST("/info/modify/admin", controller.AdminModifyUserInfo) //修改用户信息
		}

		rootAuth := user.Group("")
		rootAuth.Use(middleware.AuthMiddleware(common.Root))
		{
			rootAuth.POST("/role/modify", controller.ModifyUserRole) //修改用户权限
			rootAuth.POST("/delete", controller.AdminDeleteUser)     //删除用户
		}
	}
}
