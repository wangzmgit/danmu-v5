package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 用户注册
** 日    期: 2022年5月22日21:22:22
**********************************************************/
func Register(ctx *gin.Context) {
	//获取参数
	var request dto.RegisterDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	email := request.Email
	password := request.Password
	code := request.Code
	//数据验证
	if !util.VerifyEmailFormat(email) {
		response.CheckFail(ctx, nil, response.EmailFormatCheck)
		return
	}
	if len(password) < 6 {
		response.CheckFail(ctx, nil, response.PasswordCheck)
		return
	}
	if !service.VerificationCode(util.RegCode(email), code) {
		response.CheckFail(ctx, nil, response.VerificationCodeError)
		return
	}

	res := service.RegisterService(request)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 用户登录
** 日    期: 2022年5月22日21:32:07
**********************************************************/
func Login(ctx *gin.Context) {
	//获取参数
	var request dto.LoginDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	email := request.Email
	password := request.Password

	//数据验证
	if !util.VerifyEmailFormat(email) {
		response.CheckFail(ctx, nil, response.EmailFormatCheck)
		return
	}
	if len(password) < 6 {
		response.CheckFail(ctx, nil, response.PasswordCheck)
		return
	}

	res := service.LoginService(request, ctx.ClientIP())
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 超级管理员登录
** 日    期: 2022年7月2日15:58:49
**********************************************************/
func RootLogin(ctx *gin.Context) {
	//获取参数
	var request dto.LoginDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	email := request.Email
	password := request.Password

	//数据验证
	if !util.VerifyEmailFormat(email) {
		response.CheckFail(ctx, nil, response.EmailFormatCheck)
		return
	}
	if len(password) < 6 {
		response.CheckFail(ctx, nil, response.PasswordCheck)
		return
	}

	res := service.RootLoginService(request, ctx.ClientIP())
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 用户获取个人信息
** 日    期: 2022年5月22日21:32:00
**********************************************************/
func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": vo.ToUserVo(user.(model.User))}, response.OK)
}

/*********************************************************
** 函数功能: 用户修改个人信息
** 日    期: 2022年5月22日21:44:42
**********************************************************/
func ModifyInfo(ctx *gin.Context) {
	var request dto.ModifyUserDto
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	name := request.Name
	birthday := request.Birthday

	if len(name) == 0 {
		response.CheckFail(ctx, nil, response.NickCheck)
		return
	}
	//判断日期
	tBirthday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		response.CheckFail(ctx, nil, response.BirthdayFormatCheck)
		return
	}

	//从上下文中获取用户id
	uid := ctx.GetUint("uid")

	res := service.UserModifyService(request, uid, tBirthday)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过用户ID获取用户信息
** 日    期: 2021/7/10
** 说    明: 用于获取其他用户信息
**********************************************************/
func GetUserInfoByID(ctx *gin.Context) {
	uid := util.StringToInt(ctx.Query("uid"))
	res := service.GetUserInfoByIDService(uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 通过用户名获取用户ID
** 日    期: 2022年6月24日15:17:47
**********************************************************/
func GetUIDByName(ctx *gin.Context) {
	name := ctx.Query("name")
	res := service.GetUIDByNameService(name)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 搜索用户
** 日    期: 2022年5月22日22:04:41
**********************************************************/
func AdminSearchUser(ctx *gin.Context) {
	keyword := ctx.Query("keywords")

	if len(keyword) == 0 {
		response.Fail(ctx, nil, response.UserNotExist)
		return
	}

	res := service.AdminSearchUserService(keyword)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取用户列表
** 日    期: 2022年5月23日11:28:04
**********************************************************/
func GetUserList(ctx *gin.Context) {
	//获取分页信息
	page := util.StringToInt(ctx.Query("page"))
	pageSize := util.StringToInt(ctx.Query("page_size"))
	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, response.PageOrSizeError)
		return
	}
	if pageSize > 30 {
		response.CheckFail(ctx, nil, response.TooManyRequests)
		return
	}

	res := service.GetUserListService(page, pageSize)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 管理员修改用户信息
** 日    期: 2022年5月23日11:32:47
**********************************************************/
func AdminModifyUserInfo(ctx *gin.Context) {
	//获取参数
	var requestUser dto.AdminModifyUserDto
	err := ctx.Bind(&requestUser)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	email := requestUser.Email
	name := requestUser.Name

	if len(name) == 0 {
		response.CheckFail(ctx, nil, response.NickCheck)
		return
	}
	if !util.VerifyEmailFormat(email) {
		response.CheckFail(ctx, nil, response.EmailFormatCheck)
		return
	}

	res := service.AdminModifyUserService(requestUser)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 修改用户角色
** 日    期: 2022年7月7日10:56:59
**********************************************************/
func ModifyUserRole(ctx *gin.Context) {
	var request dto.ModifyRoleDto
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	roles := map[int]bool{
		common.User:    true,
		common.Auditor: true,
		common.Admin:   true,
	}
	if _, ok := roles[request.Role]; !ok {
		response.Fail(ctx, nil, response.RoleNotExist)
		return
	}

	res := service.ModifyUserRoleService(request.UID, request.Role)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 删除用户
** 日    期: 2021/8/3
**********************************************************/
func AdminDeleteUser(ctx *gin.Context) {
	var request dto.ID
	if err := ctx.Bind(&request); err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	res := service.AdminDeleteUserService(request.ID)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取accessToken
** 日    期: 2022年8月25日15:42:03
**********************************************************/
func GetAccessToken(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if exist {
		token, _ := common.ReleaseAccessToken(user.(model.User))
		response.Success(ctx, gin.H{"token": token}, response.OK)
		return
	}

	response.Fail(ctx, nil, response.RequestError)
}
