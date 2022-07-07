package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 获取网站数据
** 日    期: 2021年11月25日
**********************************************************/
func GetTotalWebsiteDataService() response.ResponseStruct {
	var data vo.TotalData
	DB := common.GetDB()

	DB.Model(&model.User{}).Count(&data.User)
	DB.Model(&model.Video{}).Count(&data.Video)
	DB.Model(&model.Video{}).Where("review = ?", common.WaitingReview).Count(&data.Review)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"data": data},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取网站近期(5天)数据
** 日    期: 2021年11月25日
**********************************************************/
func GetRecentWebsiteDataService() response.ResponseStruct {
	DB := common.GetDB()
	data := make([]vo.OneDayData, 5)

	for i := 0; i < 5; i++ {
		data[i] = getOneDayData(DB, i)
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"data": data},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取一天的视频和用户
** 日    期: 2021年11月25日
**********************************************************/
func getOneDayData(db *gorm.DB, offset int) vo.OneDayData {
	var data vo.OneDayData

	t := time.Now()
	startTime := t.AddDate(0, 0, -(offset + 1))
	endTime := t.AddDate(0, 0, -offset)
	data.Date = startTime.Format("2006/1/02")

	db.Model(&model.User{}).Where("created_at > ? and created_at < ?", startTime, endTime).Count(&data.User)
	db.Model(&model.Video{}).Where("created_at > ? and created_at < ?", startTime, endTime).Count(&data.Video)

	return data
}
