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
** 函数功能: 获取公告
** 日    期: 2021年11月11日22:06:33
**********************************************************/
func GetAnnounceService(page, pageSize int) response.ResponseStruct {
	var total int64
	var announceList []vo.AnnounceVo
	DB := common.GetDB()

	//拉取用户公告
	DB.Model(&model.Announce{}).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&model.Announce{}).Select("id,created_at,title,content,url").Scan(&announceList)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "announces": announceList},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员添加公告
** 日    期: 2021年11月12日12:17:08
**********************************************************/
func AddAnnounceService(announce dto.AddAnnounceDto) response.ResponseStruct {
	newAnnounce := model.Announce{
		Title:   announce.Title,
		Content: announce.Content,
		Url:     announce.Url,
	}
	DB := common.GetDB()
	DB.Create(&newAnnounce)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员删除公告
** 日    期: 2021年11月12日12:20:35
**********************************************************/
func DeleteAnnounceService(id uint) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ?", id).Delete(&model.Announce{})

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}
