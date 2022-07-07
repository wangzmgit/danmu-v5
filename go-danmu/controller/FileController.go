package controller

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/service"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 上传头像
** 日    期: 2021/7/13
**********************************************************/
func UploadAvatar(ctx *gin.Context) {
	avatar, err := ctx.FormFile("avatar")
	if err != nil {
		response.Fail(ctx, nil, response.FileUploadFail)
		return
	}
	suffix := path.Ext(avatar.Filename)
	if suffix != ".jpg" && suffix != ".jpeg" && suffix != ".png" {
		response.CheckFail(ctx, nil, response.FileCheckFail)
		return
	}
	avatar.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + util.RandomCode(6) + suffix
	errSave := ctx.SaveUploadedFile(avatar, "./file/avatar/"+avatar.Filename)
	if errSave != nil {
		response.Fail(ctx, nil, response.FileSaveFail)
		return
	}
	fileInfo, err := os.Stat("./file/avatar/" + avatar.Filename)
	//文件大小限制
	if fileInfo == nil || fileInfo.Size() > 1024*1024*viper.GetInt64("file.max_img_size") || err != nil {
		response.CheckFail(ctx, nil, response.FileCheckFail)
		return
	}

	uid := ctx.GetUint("uid")
	// 拼接上传图片的路径信息
	localFileName := "./file/avatar/" + avatar.Filename
	objectName := "avatar/" + avatar.Filename

	//记录日志
	util.LogUploadInfo(strconv.Itoa(int(uid)), "avatar", objectName, ctx.ClientIP())

	res := service.UploadAvatarService(localFileName, objectName, uid)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 上传视频封面图
** 日    期: 2021/7/14
**********************************************************/
func UploadCover(ctx *gin.Context) {
	cover, err := ctx.FormFile("cover")
	if err != nil {
		response.Fail(ctx, nil, response.FileUploadFail)
		return
	}
	suffix := path.Ext(cover.Filename)
	if suffix != ".jpg" && suffix != ".jpeg" && suffix != ".png" {
		response.CheckFail(ctx, nil, response.FileCheckFail)
		return
	}
	//储存文件
	cover.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + util.RandomCode(6) + suffix
	errSave := ctx.SaveUploadedFile(cover, "./file/cover/"+cover.Filename)
	if errSave != nil {
		response.Fail(ctx, nil, response.FileSaveFail)
		return
	}
	fileInfo, err := os.Stat("./file/cover/" + cover.Filename)
	//文件大小限制
	if fileInfo == nil || fileInfo.Size() > 1024*1024*viper.GetInt64("file.max_img_size") || err != nil {
		response.CheckFail(ctx, nil, response.FileSizeCheckFail)
		return
	}

	uid := ctx.GetUint("uid")
	localFileName := "./file/cover/" + cover.Filename
	objectName := "cover/" + cover.Filename
	//记录日志
	util.LogUploadInfo(strconv.Itoa(int(uid)), "cover", objectName, ctx.ClientIP())

	res := service.UploadCoverService(localFileName, objectName)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 上传轮播图
** 日    期: 2021/8/4
**********************************************************/
func UploadCarousel(ctx *gin.Context) {
	carousel, err := ctx.FormFile("carousel")
	if err != nil {
		response.Fail(ctx, nil, response.FileUploadFail)
		return
	}
	suffix := path.Ext(carousel.Filename)
	if suffix != ".jpg" && suffix != ".jpeg" && suffix != ".png" {
		response.CheckFail(ctx, nil, response.FileCheckFail)
		return
	}
	carousel.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + util.RandomCode(6) + suffix
	errSave := ctx.SaveUploadedFile(carousel, "./file/carousel/"+carousel.Filename)
	if errSave != nil {
		response.Fail(ctx, nil, response.FileSaveFail)
		return
	}
	fileInfo, err := os.Stat("./file/carousel/" + carousel.Filename)
	//大小限制到5M
	if fileInfo == nil || fileInfo.Size() > 1024*1024*5 || err != nil {
		response.CheckFail(ctx, nil, response.FileSizeCheckFail)
		return
	}
	// 拼接上传图片的路径信息
	localFileName := "./file/carousel/" + carousel.Filename
	objectName := "carousel/" + carousel.Filename

	res := service.UploadCarouselService(localFileName, objectName)
	response.HandleResponse(ctx, res)
}

/*********************************************************
** 函数功能: 上传视频
** 日    期: 2021/7/14
**********************************************************/
func UploadVideo(ctx *gin.Context) {
	uid := ctx.GetUint("uid")
	vid := util.StringToInt(ctx.PostForm("vid"))
	if vid <= 0 {
		response.Fail(ctx, nil, response.ParameterError)
		return
	}
	video, err := ctx.FormFile("video")
	if err != nil {
		response.Fail(ctx, nil, response.FileUploadFail)
		return
	}
	suffix := path.Ext(video.Filename)
	if suffix != ".mp4" {
		response.CheckFail(ctx, nil, response.FileCheckFail)
		return
	}
	//保存文件名并重命名文件
	title := video.Filename
	video.Filename = "upload.mp4"
	//生成目录名称并创建目录
	dirName := strconv.FormatInt(time.Now().UnixNano(), 10) + util.RandomCode(6)
	if err := os.Mkdir("./file/video/"+dirName, os.ModePerm); err != nil {
		response.ServerError(ctx, nil, response.SystemError)
		return
	}
	// 拼接上传文件的路径信息
	objectName := "video/" + dirName + "/" + video.Filename
	localFileName := "./file/" + objectName
	//保存文件
	if err := ctx.SaveUploadedFile(video, localFileName); err != nil {
		response.Fail(ctx, nil, response.FileSaveFail)
		return
	}
	fileInfo, err := os.Stat(localFileName)
	//文件大小限制
	if fileInfo == nil || fileInfo.Size() > 1024*1024*viper.GetInt64("file.max_video_size") || err != nil {
		response.CheckFail(ctx, nil, response.FileSizeCheckFail)
		return
	}
	//生成视频url和最大分辨率信息
	maxRes, duration, urls, success := service.PreTreatmentVideo(localFileName, dirName)
	if !success {
		response.Fail(ctx, nil, response.TranscodingError)
		return
	}
	//记录日志
	util.LogUploadInfo(strconv.Itoa(int(uid)), "video", objectName, ctx.ClientIP())
	//启动转码服务或上传服务
	res, resourceID := service.UploadVideoService(urls, title, duration, uint(vid), uid)
	go service.Transcoding(dirName, resourceID, maxRes)
	response.HandleResponse(ctx, res)
}
