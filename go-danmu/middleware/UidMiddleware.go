package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
)

/*********************************************************
** 函数功能: 如果登录获取用户UID，未登录则不获取
** 日    期: 2022年1月6日14:59:53
**********************************************************/
func UidMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization，header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.Set("uid", uint(0))
			return
		}

		tokenString = tokenString[7:]
		token, claims, err, _ := common.ParseUserToken(tokenString, common.AccessTypeToken)
		if err != nil || !token.Valid {
			ctx.Set("uid", uint(0))
			return
		}

		ctx.Set("uid", claims.UserId)
		return
	}
}
