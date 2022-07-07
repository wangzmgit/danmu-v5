package service

import (
	"container/list"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/util"
)

/*********************************************************
** 函数功能: 上传文件
** 日    期: 2022年6月2日19:27:37
**********************************************************/
func UploadFileToOSS(localFileName, objectName string) bool {
	return uploadToOSS(localFileName, objectName)
}

/*********************************************************
** 函数功能: 上传hls视频文件夹
** 日    期: 2022年6月2日19:31:41
**********************************************************/
func UploadHlsToOSS(dir string, files *list.List) bool {

	objectDir := "video/" + dir + "/"
	localDir := "./file/video/" + dir + "/"

	for file := files.Front(); file != nil; file = file.Next() {
		oss := objectDir + file.Value.(string)
		local := localDir + file.Value.(string)

		state := uploadToOSS(local, oss)
		if !state {
			return false
		}
	}
	return true
}

/*********************************************************
** 函数功能: 将文件上传oss
** 日    期: 2022年6月2日19:23:20
**********************************************************/
func uploadToOSS(localFileName, objectName string) bool {
	//储存到阿里云OSS
	client, err := oss.New(viper.GetString("aliyunoss.endpoint"), viper.GetString("aliyunoss.accesskey_id"), viper.GetString("aliyunoss.accesskey_secret"))
	if err != nil {
		util.LogError("OSS请求错误 " + err.Error())
		return false
	}
	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("aliyunoss.bucket"))
	if err != nil {
		util.LogError("OSS请求错误 " + err.Error())
		return false
	}

	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		util.LogError("OSS上传失败 " + err.Error())
		return false
	}

	return true
}
