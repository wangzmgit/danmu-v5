package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 上传头像
** 日    期: 2021年11月11日17:31:42
**********************************************************/
func UploadAvatarService(localFileName string, objectName string, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	if viper.GetBool("file.oss") {
		success := UploadFileToOSS(localFileName, objectName)
		if !success {
			res.HttpStatus = http.StatusBadRequest
			res.Code = response.FailCode
			return res
		}
	}

	url := getUrl() + objectName
	DB := common.GetDB()
	DB.Model(model.User{}).Where("id = ?", uid).Update("avatar", url)
	res.Data = gin.H{"url": url}
	return res
}

/*********************************************************
** 函数功能: 上传封面
** 日    期: 2021年11月11日17:38:07
**********************************************************/
func UploadCoverService(localFileName string, objectName string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	if viper.GetBool("file.oss") {
		success := UploadFileToOSS(localFileName, objectName)
		if !success {
			res.HttpStatus = http.StatusBadRequest
			res.Code = response.FailCode
			return res
		}
	}

	res.Data = gin.H{"url": getUrl() + objectName}
	return res
}

/*********************************************************
** 函数功能: 上传视频
** 日    期: 2021年11月11日17:38:07
**********************************************************/
func UploadVideoService(urls dto.ResDto, title string, duration float64, vid, uid uint) (response.ResponseStruct, uint) {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	if !isVideoExist(DB, vid) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.VideoNotExist
		return res, 0
	}

	var newResource = model.Resource{
		Vid:      vid,
		Title:    title,
		Res360:   urls.Res360,
		Res480:   urls.Res480,
		Res720:   urls.Res720,
		Res1080:  urls.Res1080,
		Original: urls.Original,
		Duration: duration,
		Review:   common.VideoProcessing,
	}

	//创建新的资源
	if err := DB.Model(&model.Resource{}).Create(&newResource).Error; err != nil {
		util.LogError("upload video error " + err.Error())
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.FileUploadFail
		return res, 0
	}

	return res, newResource.ID
}

/*********************************************************
** 函数功能: 上传轮播图
** 日    期: 2021年11月12日12:24:55
**********************************************************/
func UploadCarouselService(localFileName string, objectName string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	if viper.GetBool("file.oss") {
		success := UploadFileToOSS(localFileName, objectName)
		if !success {
			res.HttpStatus = http.StatusBadRequest
			res.Code = response.FailCode
			res.Msg = response.FileUploadFail
			return res
		}
	}

	res.Data = gin.H{"url": getUrl() + objectName}
	return res
}

/*********************************************************
** 函数功能: 完成视频上传
** 日    期: 2021/9/16
**********************************************************/
func completeUpload(resourceID uint) {
	DB := common.GetDB()
	if err := DB.Model(&model.Resource{}).Where("id = ?", resourceID).Updates(
		map[string]interface{}{"review": common.WaitingReview},
	).Error; err != nil {
		util.LogError("update review fail uuid " + strconv.Itoa(int(resourceID)) + " error " + err.Error())
	}
}

/*********************************************************
** 函数功能: 获取上传视频文件的url
** 日    期: 2022年2月16日17:11:08
**********************************************************/
func getUploadVideoUrls(dirName string, maxRes int) dto.ResDto {
	var urls dto.ResDto
	baseDir := getUrl() + "video/" + dirName
	if viper.GetInt("transcoding.max_res") == 0 {
		// hls-不处理分辨率
		urls.Original = baseDir + "/original.m3u8"
	} else {
		// hls-处理分辨率
		switch maxRes {
		case 1080:
			urls.Res1080 = baseDir + "/1080p.m3u8"
			fallthrough
		case 720:
			urls.Res720 = baseDir + "/720p.m3u8"
			fallthrough
		case 480:
			urls.Res480 = baseDir + "/480p.m3u8"
			fallthrough
		case 360:
			urls.Res360 = baseDir + "/360p.m3u8"
		}
	}
	return urls
}

/*********************************************************
** 函数功能: 获取文件的URL
** 日    期: 2022年1月5日16:49:02
**********************************************************/
func getUrl() string {
	protocol := "http://"
	if viper.GetBool("file.https") {
		return "https://"
	}
	if viper.GetBool("file.oss") {
		if len(viper.GetString("file.domain")) == 0 {
			return protocol + viper.GetString("aliyunoss.bucket") + "." + viper.GetString("aliyunoss.endpoint") + "/"
		} else {
			return protocol + viper.GetString("file.domain") + "/"
		}
	} else {
		if len(viper.GetString("file.domain")) == 0 {
			return "/api/"
		} else {
			return protocol + viper.GetString("file.domain") + "/api/"
		}
	}
}
