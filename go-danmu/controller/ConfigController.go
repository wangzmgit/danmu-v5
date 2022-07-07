package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 获取服务信息
** 日    期: 2022年7月5日19:23:04
**********************************************************/
func GetServerInfo(ctx *gin.Context) {
	res := service.GetServerInfoService()
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 获取数据库配置信息
** 日    期: 2022年2月24日11:51:31
**********************************************************/
func GetDatabaseConfig(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"db_name":    viper.GetString("datasource.database"),
		"db_host":    viper.GetString("datasource.host"),
		"db_port":    viper.GetInt("datasource.port"),
		"db_user":    viper.GetString("datasource.username"),
		"cache_host": viper.GetString("redis.host"),
		"cache_port": viper.GetInt("redis.port"),
	}, response.OK)
}

/*********************************************************
** 函数功能: 修改数据库配置信息
** 日    期: 2022年7月6日12:30:08
**********************************************************/
func SetDatabaseConfig(ctx *gin.Context) {
	var db dto.DatabasegDto
	err := ctx.Bind(&db)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	viper.Set("datasource.database", db.DbName)
	viper.Set("datasource.host", db.DbHost)
	viper.Set("datasource.port", db.DbPort)
	viper.Set("datasource.username", db.DbUser)
	viper.Set("redis.host", db.CacheHost)
	viper.Set("redis.port", db.CachePort)

	if len(db.DbPwd) != 0 {
		viper.Set("datasource.password", db.DbPwd)
	}

	if len(db.CachePwd) != 0 {
		viper.Set("redis.password", db.CachePwd)
	}

	viper.WriteConfig()

	//重新加载MySQL和Redis
	common.InitDB()
	common.InitCache()

	response.Success(ctx, nil, response.OK)
}

/*********************************************************
** 函数功能: 获取邮箱配置信息
** 日    期: 2022年2月24日16:11:27
**********************************************************/
func GetEmailConfig(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"open":    viper.GetBool("register.verify_email"),
		"name":    viper.GetString("email.name"),
		"host":    viper.GetString("email.host"),
		"port":    viper.GetInt("email.port"),
		"address": viper.GetString("email.address"),
	}, response.OK)
}

/*********************************************************
** 函数功能: 配置邮箱
** 日    期: 2022年7月6日13:22:39
**********************************************************/
func SetEmailConfig(ctx *gin.Context) {
	var email dto.EmailConfigDto
	err := ctx.Bind(&email)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	viper.Set("register.verify_email", email.Open)
	viper.Set("email.name", email.Name)
	viper.Set("email.host", email.Host)
	viper.Set("email.port", email.Port)
	viper.Set("email.address", email.Address)
	if len(email.Password) != 0 {
		viper.Set("email.password", email.Password)
	}

	viper.WriteConfig()
	response.Success(ctx, nil, response.OK)
}

/*********************************************************
** 函数功能: 获取oss配置信息
** 日    期: 2022年2月24日11:51:31
**********************************************************/
func GetOssConfig(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"storage":        viper.GetBool("file.storage"),
		"https":          viper.GetBool("file.https"),
		"domain":         viper.GetString("file.domain"),
		"max_img_size":   viper.GetInt("file.max_img_size"),
		"max_video_size": viper.GetInt("file.max_video_size"),

		"bucket":      viper.GetString("aliyunoss.bucket"),
		"endpoint":    viper.GetString("aliyunoss.endpoint"),
		"accesskeyId": viper.GetString("aliyunoss.accesskey_id"),
	}, response.OK)
}

/*********************************************************
** 函数功能: 配置oss
** 日    期: 2022年7月6日13:22:35
**********************************************************/
func SetOssConfig(ctx *gin.Context) {
	var storage dto.StorageConfigDto
	err := ctx.Bind(&storage)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	viper.Set("file.https", storage.Https)
	viper.Set("file.storage", storage.Storage)
	viper.Set("file.domain", storage.Domain)
	viper.Set("file.max_img_size", storage.MaxImgSize)
	viper.Set("file.max_video_size", storage.MaxVideoSize)

	viper.Set("aliyunoss.bucket", storage.Bucket)
	viper.Set("aliyunoss.endpoint", storage.Endpoint)
	viper.Set("aliyunoss.accesskey_id", storage.AccesskeyId)
	if len(storage.AccesskeySecret) != 0 {
		viper.Set("aliyunoss.accesskey_secret", storage.AccesskeySecret)
	}

	viper.WriteConfig()
	response.Success(ctx, nil, response.OK)
}

/*********************************************************
** 函数功能: 获取其他配置信息
** 日    期: 2022年7月6日15:29:32
**********************************************************/
func GetOtherConfig(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"prefix":  viper.GetString("register.prefix"),
		"fe_path": viper.GetString("server.fe_path"),
		"max_res": viper.GetInt("transcoding.max_res"),
		"email":   viper.GetString("admin.email"),
	}, response.OK)
}

/*********************************************************
** 函数功能: 设置其他配置
** 日    期: 2022年7月6日15:29:15
**********************************************************/
func SetOtherConfig(ctx *gin.Context) {
	var other dto.OtherConfigDto
	err := ctx.Bind(&other)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	viper.Set("register.prefix", other.Prefix)
	viper.Set("server.fe_path", other.FePath)
	viper.Set("transcoding.max_res", other.MaxRes)

	viper.Set("admin.email", other.Email)
	if len(other.Password) != 0 {
		viper.Set("admin.password", util.GetStringMD5(other.Password))
	}

	viper.WriteConfig()
	response.Success(ctx, nil, response.OK)
}
