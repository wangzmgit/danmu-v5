package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
)

/*********************************************************
** 函数功能: 点赞
** 日    期: 2021/7/22
**********************************************************/
func Like(ctx *gin.Context) {
	//获取参数
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	vid := int(request.ID)
	uid := ctx.GetUint("uid")

	res := service.LikeService(vid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 取消赞
** 日    期: 2021/7/22
**********************************************************/
func Dislike(ctx *gin.Context) {
	//获取参数
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	vid := int(request.ID)
	uid := ctx.GetUint("uid")

	res := service.DislikeService(vid, uid)
	response.HandleResponse(ctx, res)
}
