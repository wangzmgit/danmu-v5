package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 添加收藏夹
** 日    期: 2022年6月18日11:35:36
**********************************************************/
func AddCollection(ctx *gin.Context) {
	var request dto.CollectionDto
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	if request.Name == "" {
		response.CheckFail(ctx, nil, response.CollectionNameNotExist)
		return
	}
	uid := ctx.GetUint("uid")

	res := service.AddCollectionService(request, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 修改收藏夹
** 日    期: 2022年6月19日21:01:40
**********************************************************/
func ModifyCollection(ctx *gin.Context) {
	var request dto.ModifyCollectionDto
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	if request.Name == "" {
		response.CheckFail(ctx, nil, response.CollectionNameNotExist)
		return
	}
	uid := ctx.GetUint("uid")

	res := service.ModifyCollectionService(request, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除收藏夹
** 日    期: 2022年6月19日21:20:12
**********************************************************/
func DeleteCollection(ctx *gin.Context) {
	var request dto.ID
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID
	uid := ctx.GetUint("uid")

	if id == 0 {
		response.CheckFail(ctx, nil, response.CollectionNotExist)
		return
	}

	res := service.DeleteCollectionService(id, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取收藏夹列表
** 日    期: 2022年6月18日11:20:48
**********************************************************/
func GetCollectionList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	res := service.GetCollectionListService(uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 添加收藏
** 日    期: 2021/7/22
**********************************************************/
func Collect(ctx *gin.Context) {
	//获取参数
	var request dto.AddCollectDto
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	uid := ctx.GetUint("uid")

	res := service.CollectService(request, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取收藏信息
** 日    期: 2022年6月18日16:01:12
**********************************************************/
func GetCollectStatus(ctx *gin.Context) {
	//获取参数
	vid := util.StringToInt(ctx.Query("vid"))
	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}
	uid := ctx.GetUint("uid")

	res := service.GetCollectInfoService(vid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取收藏夹信息
** 日    期: 2022年6月20日09:32:46
**********************************************************/
func GetCollectionByID(ctx *gin.Context) {
	id := util.StringToInt(ctx.Query("id"))
	if id == 0 {
		response.CheckFail(ctx, nil, response.CollectionNotExist)
		return
	}
	uid := ctx.GetUint("uid")

	res := service.GetCollectionByIDService(uid, id)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取收藏列表
** 日    期: 2021/7/22
**********************************************************/
func GetCollectVideo(ctx *gin.Context) {
	id := util.StringToInt(ctx.Query("id"))
	if id == 0 {
		response.CheckFail(ctx, nil, response.CollectionNotExist)
		return
	}
	uid := ctx.GetUint("uid")
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))

	if page <= 0 || pageSize <= 0 {
		response.Fail(ctx, nil, response.PageOrSizeError)
		return
	}
	if pageSize > 30 {
		response.CheckFail(ctx, nil, response.TooManyRequests)
		return
	}

	res := service.GetCollectVideoService(uid, id, page, pageSize)
	response.HandleResponse(ctx, res)
}
