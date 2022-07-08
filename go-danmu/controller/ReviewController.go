package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 获取视频状态
** 日    期: 2021/7/16
**********************************************************/
func GetVideoStatus(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	vid := util.StringToInt(ctx.Query("vid"))

	res := service.GetVideoStatusService(vid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取待审核视频列表
** 日    期: 2021年11月12日14:55:29
**********************************************************/
func GetReviewVideoList(ctx *gin.Context) {
	//获取参数
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))

	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.GetReviewListService(page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 提交审核
** 日    期: 2022年7月1日12:09:29
**********************************************************/
func SubmitReview(ctx *gin.Context) {
	var request dto.ID
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	vid := request.ID
	uid := ctx.GetUint("uid")

	res := service.SubmitReviewService(vid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 审核视频
** 日    期: 2021年11月12日15:01:04
**********************************************************/
func ReviewVideo(ctx *gin.Context) {
	var request dto.ReviewDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	vid := request.ID
	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.ReviewVideoService(request)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 审核视频资源
** 日    期: 2022年6月10日09:19:27
**********************************************************/
func ReviewResource(ctx *gin.Context) {
	var request dto.ReviewDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	id := request.ID
	if id == 0 {
		response.CheckFail(ctx, nil, response.NotExistResources)
		return
	}

	res := service.ReviewResourceService(request)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过视频ID获取待审核视频资源
** 日    期: 2022年1月10日11:02:51
**********************************************************/
func GetReviewVideoByID(ctx *gin.Context) {
	vid := util.StringToInt(ctx.Query("vid"))
	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.GetReviewVideoByIDService(vid)
	response.HandleResponse(ctx, res)
}
