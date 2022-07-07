package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 通过视频ID获取点赞收藏数据
** 日    期: 2022年5月24日13:11:49
**********************************************************/
func GetArchiveByVID(ctx *gin.Context) {
	vid := util.StringToInt(ctx.Query("vid"))
	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}
	uid := ctx.GetUint("uid")

	res := service.GetArchiveByVIDService(vid, uid)
	response.HandleResponse(ctx, res)
}
