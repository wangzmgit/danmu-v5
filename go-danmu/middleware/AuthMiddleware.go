package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
)

/*********************************************************
** 函数功能: 用户认证中间件
** 日    期: 2022年5月22日21:38:09
**********************************************************/
func AuthMiddleware(role int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       response.UnauthorizedCode,
			Data:       nil,
			Msg:        response.Unauthorized,
		}

		//获取Authorization，header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := parseTokenString(tokenString, common.AccessTypeToken)
		if !parseSuccess {
			if isExpired {
				//如果是token过期
				res.Msg = response.TokenExpried
				res.Code = response.TokenExpriedCode
			}
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}

		//验证用户是否存在
		userId := claims.UserId
		var user model.User
		if claims.Role != common.Root {
			DB := common.GetDB()
			if err := DB.First(&user, userId).Error; err != nil || user.ID == 0 {
				response.HandleResponse(ctx, res)
				//抛弃请求
				ctx.Abort()
				return
			}
		}

		//如果用户存在，比对目标权限
		if claims.Role < role {
			response.HandleResponse(ctx, res)
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

/*********************************************************
** 函数功能: 验证refreshToken
** 日    期: 2022年8月25日15:08:49
**********************************************************/
func RefreshTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       response.UnauthorizedCode,
			Data:       nil,
			Msg:        response.Unauthorized,
		}

		//获取Authorization，header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := parseTokenString(tokenString, common.RefreshTypeToken)
		if !parseSuccess {
			if isExpired {
				//如果是token过期
				res.Msg = response.TokenExpried
				res.Code = response.TokenExpriedCode
			}
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}

		//验证用户是否存在
		userId := claims.UserId
		var user model.User
		if claims.Role != common.Root {
			DB := common.GetDB()
			if err := DB.First(&user, userId).Error; err != nil || user.ID == 0 {
				response.HandleResponse(ctx, res)
				//抛弃请求
				ctx.Abort()
				return
			}
		}

		//将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

/*********************************************************
** 函数功能: 解析token字符串
** 日    期: 2022年8月25日15:12:28
**********************************************************/
func parseTokenString(tokenString, tokenType string) (bool, *common.Claims, bool) {
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		return false, nil, false
	}

	tokenString = tokenString[7:]
	token, claims, err, isExpired := common.ParseUserToken(tokenString, tokenType)
	if err != nil || !token.Valid {
		return false, nil, isExpired
	}

	if tokenType != claims.Subject {
		return false, nil, isExpired
	}

	return true, claims, isExpired
}
