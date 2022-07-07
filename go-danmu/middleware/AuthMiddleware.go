package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
)

/*********************************************************
** 函数功能: 用户认证中间件
** 日    期: 2022年5月22日21:38:09
**********************************************************/
func AuthMiddleware(role int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization，header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃请求
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseUserToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//验证用户是否存在
		userId := claims.UserId
		var user model.User
		if claims.Role != common.Root {
			DB := common.GetDB()
			if err := DB.First(&user, userId).Error; err != nil || user.ID == 0 {
				ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
				//抛弃请求
				ctx.Abort()
				return
			}
		}

		//如果用户存在，比对目标权限
		if claims.Role < role {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃请求
			ctx.Abort()
			return
		}

		//将用户信息写入上下文
		ctx.Set("user", user)
		//将用户id和权限写入上下文
		ctx.Set("uid", userId)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
