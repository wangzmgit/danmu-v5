package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/ws"
)

/*********************************************************
** 函数功能: 获取消息Websocket连接
** 日    期: 2022年2月25日09:31:10
**********************************************************/
func GetMsgConnect(ctx *gin.Context) {
	token := ctx.Query("token")
	auth, uid := service.AnalysisToken(token)
	if !auth {
		response.Fail(ctx, nil, response.InsufficientAuthority)
		return
	}
	util.LogInfo("User " + strconv.Itoa(int(uid)) + " connection successful")
	// 升级为websocket长链接
	ws.WsMsgHandler(ctx.Writer, ctx.Request, uid)
}

/*********************************************************
** 函数功能: 发送私信
** 日    期: 2022年6月25日11:08:19
**********************************************************/
func SendMessage(ctx *gin.Context) {
	var request dto.MessageDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	fid := request.Fid
	content := request.Content
	uid := ctx.GetUint("uid")

	//验证数据
	if fid == 0 {
		response.CheckFail(ctx, nil, response.SendFail)
		return
	}
	if fid == uid {
		response.CheckFail(ctx, nil, response.CantSendYourself)
		return
	}
	if len(content) == 0 {
		response.CheckFail(ctx, nil, response.UserMsgContentCheck)
		return
	}

	res := service.SendMessageService(request, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取消息列表
** 日    期: 2022年6月25日11:24:46
**********************************************************/
func GetMessageList(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	res := service.GetMessageListService(uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取消息详细信息
** 日    期: 2022年2月26日17:45:38
**********************************************************/
func GetMessageDetails(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	fid := util.StringToUint(ctx.Query("fid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if fid == 0 {
		response.Fail(ctx, nil, response.MessageNotExist)
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

	res := service.GetMessageDetailsService(uid, fid, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 已读消息
** 日    期: 2022年3月3日19:38:30
**********************************************************/
func ReadMessageService(ctx *gin.Context) {
	var request dto.ID
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	fid := request.ID
	uid := ctx.GetUint("uid")
	//验证数据
	if fid == 0 {
		response.CheckFail(ctx, nil, response.SendFail)
		return
	}
	if fid == uid {
		response.CheckFail(ctx, nil, response.CantSendYourself)
		return
	}

	res := service.ReadMessageService(uid, fid)
	response.HandleResponse(ctx, res)
}
