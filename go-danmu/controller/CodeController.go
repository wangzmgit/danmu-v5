package controller

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 发送注册验证码
** 日    期: 2021/7/23
**********************************************************/
func SendRegisterCode(ctx *gin.Context) {
	var request dto.Email
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	email := request.Email

	//数据验证
	if !util.VerifyEmailFormat(email) {
		response.CheckFail(ctx, nil, response.EmailFormatCheck)
		return
	}

	res := service.SendRegisterCodeService(email)
	response.HandleResponse(ctx, res)
}
