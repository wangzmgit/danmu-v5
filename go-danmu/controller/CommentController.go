package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 发布评论回复
** 日    期: 2022年6月20日16:15:03
**********************************************************/
func AddComment(ctx *gin.Context) {
	var comment dto.AddCommentDto
	err := ctx.Bind(&comment)
	if err != nil {
		response.Fail(ctx, nil, response.ParameterError)
		return
	}
	content := comment.Content
	uid := ctx.GetUint("uid")

	if len(content) == 0 {
		response.CheckFail(ctx, nil, response.CommentContentCheck)
		return
	}

	res := service.AddCommentService(comment, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取评论回复列表
** 日    期: 2022年6月20日14:17:24
**********************************************************/
func GetCommentAndReplyList(ctx *gin.Context) {
	//获取分页信息
	vid := util.StringToInt(ctx.Query("vid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	replyCount := util.StringToInt(ctx.DefaultQuery("reply", "0"))
	if vid <= 0 {
		response.CheckFail(ctx, nil, response.ParameterError)
		return
	}
	if page <= 0 || pageSize <= 0 || pageSize > 20 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.GetCommentAndReplyListService(vid, replyCount, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取评论列表(不区分评论回复)
** 日    期: 2022年6月30日14:54:04
**********************************************************/
func GetCommentList(ctx *gin.Context) {
	//获取分页信息
	vid := util.StringToInt(ctx.Query("vid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if vid <= 0 {
		response.CheckFail(ctx, nil, response.ParameterError)
		return
	}

	if page <= 0 || pageSize <= 0 || pageSize > 20 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.GetCommentListService(vid, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取回复详情
** 日    期: 2022年4月23日15:37:16
**********************************************************/
func GetReplyDetails(ctx *gin.Context) {
	cid := util.StringToInt(ctx.Query("cid"))
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	offset := util.StringToInt(ctx.DefaultQuery("offset", "0"))

	if cid <= 0 {
		response.CheckFail(ctx, nil, response.ParameterError)
		return
	}
	if page <= 0 || pageSize <= 0 || pageSize > 20 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}

	res := service.GetReplyDetailsService(cid, offset, page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除评论
** 日    期: 2022年6月20日17:46:51
**********************************************************/
func DeleteComment(ctx *gin.Context) {
	//获取参数
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID
	uid := ctx.GetUint("uid")

	res := service.DeleteCommentService(id, uid)
	response.HandleResponse(ctx, res)
}
