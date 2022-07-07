package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
)

/*********************************************************
** 函数功能: 获取轮播图
** 日    期: 2022年7月1日11:49:16
**********************************************************/
func GetCarousel(ctx *gin.Context) {
	res := service.GetCarouselService()
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 上传轮播图信息
** 日    期: 2021/8/4
**********************************************************/
func UploadCarouselInfo(ctx *gin.Context) {
	//获取参数
	var carousel dto.CarouselDto
	err := ctx.Bind(&carousel)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	img := carousel.Img
	url := carousel.Url

	//验证数据
	if len(img) == 0 {
		response.CheckFail(ctx, nil, response.ImgNotExist)
		return
	}

	res := service.UploadCarouselInfoService(img, url)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除轮播图
** 日    期: 2021/8/4
**********************************************************/
func DeleteCarousel(ctx *gin.Context) {
	//获取参数
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID

	if id == 0 {
		response.CheckFail(ctx, nil, response.CarouselNotExist)
		return
	}

	res := service.DeleteCarouselService(id)
	response.HandleResponse(ctx, res)
}
