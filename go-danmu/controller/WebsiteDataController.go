package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
)

/*********************************************************
** 函数功能: 获取网站数据
** 日    期: 2021年11月25日18:00:40
**********************************************************/
func GetTotalWebsiteData(ctx *gin.Context) {
	res := service.GetTotalWebsiteDataService()
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取网站近期(5天)数据
** 日    期: 2021年11月25日
**********************************************************/
func GetRecentWebsiteData(ctx *gin.Context) {
	res := service.GetRecentWebsiteDataService()
	response.HandleResponse(ctx, res)
}
