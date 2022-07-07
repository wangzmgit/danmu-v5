package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 发送注册验证码
** 日    期: 2022年3月24日12:18:40
**********************************************************/
func SendRegisterCodeService(email string) response.ResponseStruct {
	DB := common.GetDB()
	//邮箱是否存在
	if isEmailExist(DB, email) {
		return response.ResponseStruct{
			HttpStatus: http.StatusUnprocessableEntity,
			Code:       response.CheckFailCode,
			Data:       nil,
			Msg:        response.EmailRegistered,
		}

	}
	Cache := common.GetCache()
	code := util.RandomCode(6)
	res, success := getCode(Cache, code, util.RegCode(email))
	if success {
		if viper.GetBool("register.verify_email") {
			send := util.SendEmail(email, code, util.RegisterCode)
			if !send {
				Cache.Del(util.RegCode(email))
				res.HttpStatus = http.StatusBadRequest
				res.Code = response.FailCode
				res.Msg = response.SendFail
				return res
			}
		} else {
			res.Data = gin.H{"code": code}
			return res
		}
	}
	return res
}

/*********************************************************
** 函数功能: 验证验证码
** 日    期:2021/7/24
**********************************************************/
func VerificationCode(emailKey string, code string) bool {
	if len(code) == 0 {
		return false
	}
	Cache := common.GetCache()
	if Cache == nil {
		util.LogError("Verification code redis error")
		return false
	}
	dbCode, _ := Cache.Get(emailKey).Result()
	if dbCode == "" || dbCode != code {
		return false
	}
	Cache.Del(emailKey)
	return true
}

/*********************************************************
** 函数功能: 获取验证码
** 日    期: 2022年6月3日17:44:02
**********************************************************/
func getCode(cache *redis.Client, randomCode, key string) (response.ResponseStruct, bool) {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//存储code到redis
	if cache == nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res, false
	}
	code, _ := cache.Get(key).Result()
	if code != "" {
		//如果时间小于一分钟则不能重新发送
		duration, _ := cache.TTL(key).Result()
		if duration >= 240000000000 {
			res.HttpStatus = http.StatusBadRequest
			res.Code = response.FailCode
			res.Msg = response.OperationTooFrequently
			return res, false
		}
	}

	err := cache.Set(key, randomCode, time.Second*300).Err()
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SendFail
		return res, false
	}
	return res, true
}
