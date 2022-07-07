package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 获取轮播图
** 日    期: 2021年11月11日22:17:35
**********************************************************/
func GetCarouselService() response.ResponseStruct {
	DB := common.GetDB()
	var carousels []vo.CarouselVo
	DB.Model(&model.Carousel{}).Select("id,img,url,created_at").Scan(&carousels)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"carousels": carousels},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员上传轮播图信息
** 日    期: 2021年11月12日13:50:28
**********************************************************/
func UploadCarouselInfoService(img string, url string) response.ResponseStruct {
	DB := common.GetDB()
	newCarousel := model.Carousel{
		Img: img,
		Url: url,
	}
	DB.Create(&newCarousel)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员删除轮播图
** 日    期: 2021年11月12日14:37:55
**********************************************************/
func DeleteCarouselService(id uint) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ?", id).Delete(&model.Carousel{})

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}
