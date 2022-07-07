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
** 函数功能: 获取弹幕
** 日    期: 2022年6月29日
**********************************************************/
func GetDanmakuService(vid, part int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var danmakuList []vo.DanmakuVo
	DB := common.GetDB()
	if !isVideoExist(DB, uint(vid)) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.VideoNotExist
		return res
	}
	DB.Model(&model.Danmaku{}).Select("time,type,color,text").
		Where("vid = ? and part = ?", vid, part).Order("time").Scan(&danmakuList)
	res.Data = gin.H{"danmaku": danmakuList}
	return res
}

/*********************************************************
** 函数功能: 发送弹幕
** 日    期: 2022年6月29日
**********************************************************/
func SendDanmakuService(danmaku dto.DanmakuDto, uid uint) response.ResponseStruct {
	DB := common.GetDB()
	newDanmaku := model.Danmaku{
		Vid:   danmaku.Vid,
		Part:  danmaku.Part,
		Time:  danmaku.Time,
		Type:  danmaku.Type,
		Color: danmaku.Color,
		Text:  danmaku.Text,
		Uid:   uid,
	}
	DB.Create(&newDanmaku)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}
