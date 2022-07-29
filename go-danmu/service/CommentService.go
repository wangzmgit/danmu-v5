package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 发布评论回复
** 日    期: 2022年6月20日16:12:22
**********************************************************/
func AddCommentService(comment dto.AddCommentDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	newComment := model.Comment{
		Vid:      comment.Vid,
		Content:  comment.Content,
		ParentID: comment.ParentID,
		Uid:      uid,
	}

	DB := common.GetDB()
	DB.Create(&newComment)

	res.Data = gin.H{"id": newComment.ID}
	return res
}

/*********************************************************
** 函数功能: 获取评论列表
** 日    期: 2022年6月20日14:16:16
**********************************************************/
func GetCommentAndReplyListService(vid, replyCount, page, pageSize int) response.ResponseStruct {
	var total int64        //总评论数
	var commentCount int64 //一级评论数
	var comments []vo.CommentVo

	DB := common.GetDB()
	DB.Model(model.Comment{}).Where("vid = ?", vid).Count(&total)
	DB.Model(model.Comment{}).Where("parent_id != 0 and vid = ?", vid).Count(&commentCount)
	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from users,comments " +
		"where comments.deleted_at is null and comments.uid = users.id and vid = ? and parent_id = 0 limit ? offset ?"
	sqlReply := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from comments,users " +
		"where comments.deleted_at is null and comments.uid = users.id and parent_id = ? limit ?"
	DB.Raw(sql, vid, pageSize, (page-1)*pageSize).Scan(&comments)

	if replyCount > 0 {
		//当需要的回复数大于0时，查询回复
		for i := 0; i < len(comments); i++ {
			DB.Raw(sqlReply, comments[i].ID, replyCount).Scan(&comments[i].Reply)
		}
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "comment_count": commentCount, "comments": comments},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取评论列表(不区分评论回复)
** 日    期: 2022年6月30日14:54:04
**********************************************************/
func GetCommentListService(vid, page, pageSize int) response.ResponseStruct {
	var total int64 //记录总数
	var comments []vo.CommentListVo

	DB := common.GetDB()
	DB.Model(model.Comment{}).Where("vid = ?", vid).Count(&total)
	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from users,comments " +
		"where comments.deleted_at is null and comments.uid = users.id and vid = ? limit ? offset ?"

	DB.Raw(sql, vid, pageSize, (page-1)*pageSize).Scan(&comments)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "comments": comments},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取回复详情
** 日    期: 2022年6月20日15:33:13
**********************************************************/
func GetReplyDetailsService(id, offset, page, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var replies []vo.ReplyVo

	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from comments,users " +
		"where comments.deleted_at is null and comments.uid = users.id and parent_id = ? limit ? offset ?"
	DB := common.GetDB()

	DB.Raw(sql, id, pageSize, (page-1)*pageSize+offset).Scan(&replies)
	res.Data = gin.H{
		"comments": replies,
	}
	return res
}

/*********************************************************
** 函数功能: 删除评论回复
** 日    期: 2022年6月20日17:35:50
**********************************************************/
func DeleteCommentService(id uint, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var comment model.Comment
	DB.Where("id = ?", id).First(&comment)
	if comment.ID != 0 {
		//uid为评论作者或视频作者
		if comment.Uid == uid || uid == getVideoAuthorID(DB, comment.Vid) {
			DB.Where("id = ? and uid = ?", id, uid).Delete(&comment)
		}
	}

	return res
}
