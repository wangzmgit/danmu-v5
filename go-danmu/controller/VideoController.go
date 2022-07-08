package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/ws"
)

/*********************************************************
** 函数功能: 上传视频信息
** 日    期: 2022年5月23日14:17:30
**********************************************************/
func UploadVideoInfo(ctx *gin.Context) {
	var video dto.UploadVideoDto
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	title := video.Title
	cover := video.Cover
	partition := video.Partition
	uid := ctx.GetUint("uid")

	//验证数据
	if len(title) == 0 {
		response.CheckFail(ctx, nil, response.TitleCheck)
		return
	}
	if len(cover) == 0 {
		response.CheckFail(ctx, nil, response.CoverCheck)
		return
	}
	if partition == 0 {
		response.CheckFail(ctx, nil, response.PartitionCheck)
		return
	}

	res := service.UploadVideoInfoService(video, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 修改视频信息
** 日    期: 2021/7/17
**********************************************************/
func ModifyVideoInfo(ctx *gin.Context) {
	//获取参数
	var videoInfo dto.ModifyVideoInfoDto
	err := ctx.Bind(&videoInfo)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	title := videoInfo.Title
	cover := videoInfo.Cover

	if len(title) == 0 {
		response.CheckFail(ctx, nil, response.TitleCheck)
		return
	}
	if len(cover) == 0 {
		response.CheckFail(ctx, nil, response.CoverCheck)
		return
	}

	//从上下文中获取用户id
	uid := ctx.GetUint("uid")
	res := service.ModifyVideoInfoService(videoInfo, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 修改资源标题
** 日    期: 2022年7月7日11:59:55
**********************************************************/
func ModifyResourceTitle(ctx *gin.Context) {
	//获取参数
	var request dto.ModifyResourceTitleDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	if request.ID == 0 {
		response.CheckFail(ctx, nil, response.NotExistResources)
		return
	}
	if len(request.Title) == 0 {
		response.CheckFail(ctx, nil, response.TitleCheck)
		return
	}

	uid := ctx.GetUint("uid")
	res := service.ModifyResourceTitleService(request.ID, uid, request.Title)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 用户删除视频
** 日    期: 2021/7/17
**********************************************************/
func DeleteVideo(ctx *gin.Context) {
	//获取参数
	var video dto.ID
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := video.ID
	uid := ctx.GetUint("uid")
	//数据验证
	if id == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.DeleteVideoService(id, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除资源
** 日    期: 2022年7月1日10:02:44
**********************************************************/
func DeleteResource(ctx *gin.Context) {
	//获取参数
	var video dto.ID
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := video.ID
	uid := ctx.GetUint("uid")
	//数据验证
	if id == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.DeleteResourceService(id, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取自己的视频
** 日    期: 2021/7/17
**********************************************************/
func GetUploadVideoList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	//获取分页信息
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}
	res := service.GetUploadVideoListService(page, pageSize, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 用户获取资源列表
** 日    期: 2022年6月6日17:28:17
**********************************************************/
func GetUploadResourceList(ctx *gin.Context) {
	vid := util.StringToInt(ctx.Query("vid"))

	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	uid := ctx.GetUint("uid")

	res := service.GetUploadResourceListService(vid, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过ID获取视频
** 日    期: 2021/10/31
**********************************************************/
func GetVideoByID(ctx *gin.Context) {
	vid := util.StringToInt(ctx.Query("vid"))
	if vid == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.GetVideoByIDService(vid, ctx.ClientIP())
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取推荐视频
** 日    期: 2021/10/26
**********************************************************/
func GetRecommendVideo(ctx *gin.Context) {
	//因为视频比较少，就直接按播放量排名
	res := service.GetRecommendVideoService()
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取视频列表
** 日    期: 2021年12月11日17:03:06
**********************************************************/
func GetVideoList(ctx *gin.Context) {
	var request dto.GetVideoListDto
	request.Page = util.StringToInt(ctx.Query("page"))
	request.PageSize = util.StringToInt(ctx.Query("page_size"))
	request.Partition = util.StringToInt(ctx.DefaultQuery("partition", "0")) //分区

	if request.Page <= 0 || request.PageSize <= 0 {
		response.Fail(ctx, nil, response.PageOrSizeError)
		return
	}
	if request.PageSize >= 30 {
		response.Fail(ctx, nil, response.TooManyRequests)
		return
	}

	res := service.GetVideoListService(request)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过用户ID获取视频列表
** 日    期: 2021/10/26
**********************************************************/
func GetVideoListByUID(ctx *gin.Context) {
	uid := util.StringToInt(ctx.Query("uid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))

	if page <= 0 || pageSize <= 0 {
		response.Fail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.GetVideoListByUIDService(uid, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 搜索视频
** 日    期: 2022年3月24日19:30:51
**********************************************************/
func SearchVideo(ctx *gin.Context) {
	keyword := ctx.Query("keywords")
	if len(keyword) == 0 {
		response.Fail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.SearchVideoService(keyword)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取视频列表
** 日    期: 2021/8/4
**********************************************************/
func AdminGetVideoList(ctx *gin.Context) {
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.AdminGetVideoListService(page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除视频
** 日    期:2021/8/3
**********************************************************/
func AdminDeleteVideo(ctx *gin.Context) {
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID

	if id == 0 {
		response.CheckFail(ctx, nil, response.VideoNotExist)
		return
	}

	res := service.AdminDeleteVideoService(id)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 视频Websocket连接(统计在线人数)
** 日    期: 2022年5月16日22:24:19
**********************************************************/
func GetRoomConnect(ctx *gin.Context) {
	vid := ctx.Query("vid")
	if util.StringToInt(vid) == 0 {
		return
	}

	// 升级为websocket长链接
	ws.WsHandler(ctx.Writer, ctx.Request, vid)
}
