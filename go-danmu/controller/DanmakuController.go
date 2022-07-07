package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 获取弹幕
** 日    期: 2022年6月29日
**********************************************************/
func GetDanmaku(ctx *gin.Context) {
	vid := util.StringToInt(ctx.Query("vid"))
	part := util.StringToInt(ctx.Query("part"))
	if vid == 0 {
		response.Fail(ctx, nil, response.ParameterError)
		return
	}

	res := service.GetDanmakuService(vid, part)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 发送弹幕
** 日    期: 2022年6月29日
**********************************************************/
func SendDanmaku(ctx *gin.Context) {
	var danmaku dto.DanmakuDto
	err := ctx.Bind(&danmaku)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	uid := ctx.GetUint("uid")
	if danmaku.Vid == 0 {
		response.CheckFail(ctx, nil, response.SendFail)
		return
	}
	if danmaku.Text == "" {
		response.CheckFail(ctx, nil, response.DanmakuCheck)
		return
	}

	if danmaku.Part == 0 {
		danmaku.Part = 1
	}

	res := service.SendDanmakuService(danmaku, uid)
	response.HandleResponse(ctx, res)
}
