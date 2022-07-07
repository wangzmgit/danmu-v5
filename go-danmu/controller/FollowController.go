package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 关注
** 日    期: 2021/7/24
**********************************************************/
func Following(ctx *gin.Context) {
	//获取参数
	var follow dto.ID
	err := ctx.Bind(&follow)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	//关注的人的id和自己的id
	fid := follow.ID
	uid := ctx.GetUint("uid")
	//判断关注的是否为自己
	if fid == uid {
		response.CheckFail(ctx, nil, response.CantFollowYourself)
		return
	}

	res := service.FollowingService(fid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 取消关注
** 日    期: 2021/7/24
**********************************************************/
func UnFollow(ctx *gin.Context) {
	var follow dto.ID
	err := ctx.Bind(&follow)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	//关注的人的id和自己的id
	fid := follow.ID
	uid := ctx.GetUint("uid")

	res := service.UnFollowService(fid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取关注状态
** 日    期: 2021/7/25
**********************************************************/
func GetFollowStatus(ctx *gin.Context) {
	fid := util.StringToInt(ctx.Query("fid"))
	if fid == 0 {
		response.CheckFail(ctx, nil, response.UserNotExist)
		return
	}
	uid := ctx.GetUint("uid")
	res := service.GetFollowStatusService(uid, uint(fid))
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过UID获取关注列表
** 日    期: 2021年11月29日14:49:48
**********************************************************/
func GetFollowingByID(ctx *gin.Context) {
	uid := util.StringToInt(ctx.Query("uid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))

	if uid == 0 {
		response.CheckFail(ctx, nil, response.UserNotExist)
		return
	}

	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	if pageSize >= 30 {
		response.CheckFail(ctx, nil, response.TooManyRequests)
		return
	}

	res := service.GetFollowingByIDService(uid, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过UID获取粉丝列表
** 日    期: 2021/7/25
**********************************************************/
func GetFollowersByID(ctx *gin.Context) {
	uid := util.StringToInt(ctx.Query("uid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))

	if uid == 0 {
		response.CheckFail(ctx, nil, response.UserNotExist)
		return
	}

	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	if pageSize >= 30 {
		response.CheckFail(ctx, nil, response.TooManyRequests)
		return
	}
	res := service.GetFollowersByIDService(uid, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过UID获取粉丝数
** 日    期: 2021/7/25
**********************************************************/
func GetFollowCount(ctx *gin.Context) {
	uid := util.StringToInt(ctx.Query("uid"))
	if uid == 0 {
		response.CheckFail(ctx, nil, response.UserNotExist)
		return
	}

	res := service.GetFollowCountService(uid)
	response.HandleResponse(ctx, res)
}
