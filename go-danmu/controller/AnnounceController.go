package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 获取公告
** 日    期:2021/7/29
**********************************************************/
func GetAnnounce(ctx *gin.Context) {
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	if pageSize >= 30 {
		response.CheckFail(ctx, nil, response.TooManyRequests)
		return
	}
	res := service.GetAnnounceService(page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 添加公告
** 日    期: 2021/8/4
**********************************************************/
func AddAnnounce(ctx *gin.Context) {
	//获取参数
	var announce dto.AddAnnounceDto
	err := ctx.Bind(&announce)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	title := announce.Title
	content := announce.Content

	if len(title) == 0 {
		response.CheckFail(ctx, nil, response.TitleCheck)
		return
	}
	if len(content) == 0 {
		response.CheckFail(ctx, nil, response.AnnounceContentCheck)
		return
	}

	res := service.AddAnnounceService(announce)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除公告
** 日    期: 2021/8/4
**********************************************************/
func DeleteAnnounce(ctx *gin.Context) {
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID

	res := service.DeleteAnnounceService(id)
	response.HandleResponse(ctx, res)
}
