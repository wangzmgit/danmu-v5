package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 2000, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 4000, data, msg)
}

func CheckFail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, 4022, data, msg)
}

func ServerError(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusInternalServerError, 5000, data, msg)
}

/*********************************************************
** 函数功能: 处理返回信息
** 日    期: 2021/11/8
**********************************************************/
func HandleResponse(ctx *gin.Context, res ResponseStruct) {
	ctx.JSON(res.HttpStatus, gin.H{"code": res.Code, "data": res.Data, "msg": res.Msg})
}
