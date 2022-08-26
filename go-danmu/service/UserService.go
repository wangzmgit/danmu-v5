package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/vo"

	"golang.org/x/crypto/bcrypt"
)

/*********************************************************
** 函数功能: 注册
** 日    期: 2022年5月22日21:22:14
**********************************************************/
func RegisterService(user dto.RegisterDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()

	//邮箱是否存在
	if isEmailExist(DB, user.Email) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.EmailRegistered
		return res
	}

	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		//记录日志
		util.LogError("hashed password " + err.Error())
		return res
	}

	newUser := model.User{
		Name:     viper.GetString("register.prefix") + strconv.FormatInt(time.Now().Unix(), 10) + util.RandomCode(3),
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	return res
}

/*********************************************************
** 函数功能: 登录
** 日    期: 2022年5月22日21:32:29
**********************************************************/
func LoginService(login dto.LoginDto, userIP string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	//判断邮箱是否存在
	var user model.User
	DB := common.GetDB()
	DB.Where("email = ?", login.Email).First(&user)
	if user.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UserNotExist
		return res
	}
	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.NameOrPasswordError
		return res
	}
	//发放token
	refreshToken, accessToken, err := common.ReleaseToken(user)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		util.LogError("token generate error " + err.Error())
		return res
	}
	util.LogInfo("token issued successfully uid " + strconv.Itoa(int(user.ID)) + " | " + userIP)
	//返回数据
	res.Data = gin.H{"refresh_token": refreshToken, "access_token": accessToken, "user": vo.ToUserVo(user)}
	return res
}

/*********************************************************
** 函数功能: 超级管理员登录
** 日    期: 2022年7月2日15:59:51
**********************************************************/
func RootLoginService(login dto.LoginDto, userIP string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	email := viper.GetString("admin.email")
	password := viper.GetString("admin.password")
	if login.Email == email && util.GetStringMD5(login.Password) == password {
		rootInfo := model.User{
			Name:  "超级管理员",
			Email: email,
			Role:  common.Root,
		}
		//发放token
		refreshToken, accessToken, err := common.ReleaseToken(rootInfo)
		if err != nil {
			res.HttpStatus = http.StatusInternalServerError
			res.Code = response.ServerErrorCode
			res.Msg = response.SystemError
			util.LogError("token generate error " + err.Error())
			return res
		}
		util.LogInfo("token issued successfully root  | " + userIP)

		//返回数据
		res.Data = gin.H{"refresh_token": refreshToken, "access_token": accessToken, "user": vo.ToUserVo(rootInfo)}
		return res
	}

	res.HttpStatus = http.StatusBadRequest
	res.Code = response.FailCode
	res.Msg = response.UserNotExist
	return res
}

/*********************************************************
** 函数功能: 修改用户信息
** 日    期: 2022年5月22日21:43:51
**********************************************************/
func UserModifyService(modify dto.ModifyUserDto, uid uint, tBirthday time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	//查询用户名是否存在
	var user model.User
	DB.Where("name = ?", modify.Name).First(&user)
	if user.ID != 0 && user.ID != uid {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.NameExist
		return res
	}

	err := DB.Model(model.User{}).Where("id = ?", uid).Updates(
		map[string]interface{}{
			"name":     modify.Name,
			"gender":   modify.Gender,
			"birthday": tBirthday,
			"sign":     modify.Sign,
		},
	).Error
	if err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.ModifyFail
		return res
	}
	return res
}

/*********************************************************
** 函数功能: 通过用户ID获取用户信息
** 日    期: 2021/11/10
**********************************************************/
func GetUserInfoByIDService(uid interface{}) response.ResponseStruct {
	var user model.User
	DB := common.GetDB()
	DB.Select("id,name,sign,avatar,gender").Where("id = ?", uid).First(&user)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"user": vo.ToUserVo(user)},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过用户名获取用户ID
** 日    期: 2022年6月24日15:21:28
**********************************************************/
func GetUIDByNameService(name string) response.ResponseStruct {
	var id uint
	DB := common.GetDB()
	DB.Model(&model.User{}).Where("name = ?", name).Pluck("id", &id)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"uid": id},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 搜索用户
** 日    期: 2022年5月22日22:04:16
**********************************************************/
func AdminSearchUserService(keyword string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var users []vo.AdminUserListVo

	DB := common.GetDB()
	DB.Model(model.User{}).Where("name like ? or id = ? or email = ?", keyword+"%", keyword, keyword).Scan(&users).Limit(10)

	res.Data = gin.H{"users": users}
	return res
}

/*********************************************************
** 函数功能: 管理员获取用户列表
** 日    期: 2022年5月23日11:28:35
**********************************************************/
func GetUserListService(page int, pageSize int) response.ResponseStruct {
	//记录总数
	var total int64
	var users []vo.AdminUserListVo
	DB := common.GetDB()
	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&model.User{}).Count(&total)
	Pagination.Model(&model.User{}).Select("id,name,created_at,email,avatar,sign,gender,role").Scan(&users)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "users": users},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员修改用户信息
** 日    期: 22022年5月23日11:33:19
**********************************************************/
func AdminModifyUserService(newInfo dto.AdminModifyUserDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var user model.User
	DB := common.GetDB()

	//如果用户名已经存在且不属于当前用户,返回
	DB.Where("name = ?", newInfo.Name).First(&user)
	if user.ID != 0 && user.ID != newInfo.ID {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.NameExist
		return res
	}
	//如果邮箱已经存在且不属于当前用户,返回
	DB.Where("email = ?", newInfo.Email).First(&user)
	if user.ID != 0 && user.ID != newInfo.ID {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.EmailRegistered
		return res
	}

	DB.Model(&model.User{}).Where("id = ?", newInfo.ID).Updates(
		map[string]interface{}{
			"email": newInfo.Email,
			"name":  newInfo.Name,
			"sign":  newInfo.Sign,
		},
	)
	return res
}

/*********************************************************
** 函数功能: 管理员删除用户
** 日    期: 2021年11月12日15:26:42
**********************************************************/
func AdminDeleteUserService(id uint) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ?", id).Delete(&model.User{})

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 修改用户角色
** 日    期: 2022年7月7日10:56:59
**********************************************************/
func ModifyUserRoleService(uid uint, role int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	DB.Model(&model.User{}).Where("id = ?", uid).Updates(map[string]interface{}{"role": role})

	return res
}

/*********************************************************
** 函数功能: 邮箱是否存在
** 日    期: 2021/7/10
**********************************************************/
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

/*********************************************************
** 函数功能: 用户是否存在
** 日    期: 2021/7/10
**********************************************************/
func isUserExist(db *gorm.DB, id uint) bool {
	var user model.User
	db.First(&user, id)
	if user.ID != 0 {
		return true
	}
	return false
}

/*********************************************************
** 函数功能: 获取用户信息
** 日    期: 2022年6月10日15:47:00
**********************************************************/
func getUserInfo(db *gorm.DB, uid uint) model.User {
	var user model.User
	db.First(&user, uid)
	return user
}
