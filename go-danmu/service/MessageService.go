package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
	"kuukaa.fun/danmu-v5/ws"
)

/*********************************************************
** 函数功能: 发送私信
** 日    期: 2021年11月12日11:18:49
**********************************************************/
func SendMessageService(msg dto.MessageDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	newMsg := [2]model.UserMessage{
		{
			Uid:     uid,
			Fid:     msg.Fid,
			FromId:  uid,
			ToId:    msg.Fid,
			Content: msg.Content,
			Status:  true,
		},
		{
			Uid:     msg.Fid,
			Fid:     uid,
			FromId:  uid,
			ToId:    msg.Fid,
			Content: msg.Content,
		},
	}

	//写入数据库
	if err := DB.Create(&newMsg).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.SendMsgFail
		return res
	}

	//推送消息给接收者
	data, _ := json.Marshal(&vo.MessageVo{
		Fid:     uid,
		Content: msg.Content,
	})
	ws.SendMsgToUser(msg.Fid, data)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取消息列表
** 日    期: 2022年6月25日11:24:57
**********************************************************/
func GetMessageListService(uid uint) response.ResponseStruct {
	DB := common.GetDB()
	var messageList []vo.MessagesListVo
	sql := "select user_messages.id,user_messages.created_at,`status`,users.id as uid,users.name,users.avatar "
	sql += "from user_messages,users where user_messages.id in (select Max(id) from user_messages where deleted_at is null "
	sql += "group by fid) and user_messages.fid = users.id and uid = ? order by id desc"
	DB.Raw(sql, uid).Scan(&messageList)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"messages": messageList},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取消息详细信息
** 日    期: 2022年2月26日17:47:01
**********************************************************/
func GetMessageDetailsService(uid, fid uint, page, pageSize int) response.ResponseStruct {
	var messages []vo.MessageDetailsVo

	DB := common.GetDB()
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize).Order("id desc")
	DB.Model(&model.UserMessage{}).Select("fid,from_id,content,created_at").
		Where("uid = ? AND fid = ?", uid, fid).Scan(&messages)

	// 此时查询到的消息为为倒叙，需要进行反转
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"messages": messages},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 已读消息
** 日    期: 2022年3月3日19:40:20
**********************************************************/
func ReadMessageService(uid, fid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	if err := DB.Model(&model.UserMessage{}).
		Where("uid = ? and fid = ?", uid, fid).Update("status", true).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}

	return res
}

/*********************************************************
** 函数功能: 手动解析token
** 日    期: 2022年2月25日09:47:02
**********************************************************/
func AnalysisToken(tokenString string) (bool, uint) {
	if tokenString == "" {
		return false, uint(0)
	}

	token, claims, err := common.ParseUserToken(tokenString)
	if err != nil || !token.Valid {
		return false, uint(0)
	}

	return true, claims.UserId
}
