package service

import (
	"bufio"
	"bytes"
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/util"
)

var resInfo = map[int]string{
	360:  "600:360",
	480:  "720:480",
	720:  "1280:720",
	1080: "1920:1080",
}

var resName = map[int]string{
	0:    "original",
	360:  "360p",
	480:  "480p",
	720:  "720p",
	1080: "1080p",
}

/********************************************************************************
* @功    能：转码
* @输入参数：目录名，资源id，最大分辨率
* @返 回 值：无
* @日    期：2022年6月5日13:52:10
* @版    本：1.0
********************************************************************************/
func Transcoding(dirName string, resourceID uint, maxRes int) {
	baseUrl := getUrl() + "video/" + dirName + "/"
	targetDir := "./file/video/" + dirName + "/"

	if maxRes == 0 {
		//调整码率
		toBitRate(targetDir, 0)
		toHlsDifferentRes(baseUrl, targetDir, dirName, 0, resourceID)
		//删除临时文件
		os.Remove(targetDir + "temp_original.ts")
		os.Remove(targetDir + "temp_original.m3u8")
		completeUpload(resourceID)
	} else {
		var wg sync.WaitGroup
		switch maxRes {
		case 1080:
			wg.Add(1)
			go func() {
				file1080, _ := toTargetRes(targetDir, 1080)
				toBitRate(file1080, 1080)
				toHlsDifferentRes(baseUrl, targetDir, dirName, 1080, resourceID)
				//删除临时文件
				os.Remove(targetDir + "temp_1080p.ts")
				os.Remove(targetDir + "temp_1080p.mp4")
				os.Remove(targetDir + "temp_1080p.m3u8")
				wg.Done()
			}()
			fallthrough
		case 720:
			wg.Add(1)
			go func() {
				file720, _ := toTargetRes(targetDir, 720)
				toBitRate(file720, 720)
				toHlsDifferentRes(baseUrl, targetDir, dirName, 720, resourceID)
				//删除临时文件
				os.Remove(targetDir + "temp_720p.ts")
				os.Remove(targetDir + "temp_720p.mp4")
				os.Remove(targetDir + "temp_720p.m3u8")
				wg.Done()
			}()
			fallthrough
		case 480:
			wg.Add(1)
			go func() {
				file480, _ := toTargetRes(targetDir, 480)
				toBitRate(file480, 480)
				toHlsDifferentRes(baseUrl, targetDir, dirName, 480, resourceID)
				//删除临时文件
				os.Remove(targetDir + "temp_480p.ts")
				os.Remove(targetDir + "temp_480p.mp4")
				os.Remove(targetDir + "temp_480p.m3u8")
				wg.Done()
			}()
			fallthrough
		case 360:
			wg.Add(1)
			go func() {
				file360, _ := toTargetRes(targetDir, 360)
				toBitRate(file360, 360)
				toHlsDifferentRes(baseUrl, targetDir, dirName, 360, resourceID)
				//删除临时文件
				os.Remove(targetDir + "temp_360p.ts")
				os.Remove(targetDir + "temp_360p.mp4")
				os.Remove(targetDir + "temp_360p.m3u8")
				wg.Done()
			}()
		}
		wg.Wait()
		completeUpload(resourceID)
	}

}

/********************************************************************************
* @功    能：预处理视频
* @输入参数：视频文件路径,视频目录名
* @返 回 值：最大分辨率，视频时长，视频url，是否成功
* @日    期：2022年6月15日18:37:56
* @版    本：1.0
********************************************************************************/
func PreTreatmentVideo(input, dirName string) (int, float64, dto.ResDto, bool) {
	var err error
	var videoData dto.VideoInfoData
	videoData, err = getVideoInfo(input)
	if err != nil {
		return 0, 0, dto.ResDto{}, false
	}

	if videoData.Stream[0].CodecName != "h264" {
		return 0, 0, dto.ResDto{}, false
	}
	//计算最大分辨率
	width := videoData.Stream[0].Width
	height := videoData.Stream[0].Height
	maxRes := util.Min(getWidthRes(width), getHeigthRes(height))
	maxRes = util.Min(maxRes, viper.GetInt("transcoding.max_res"))
	//获取视频时长
	duration, _ := strconv.ParseFloat(videoData.Stream[0].Duration, 64)
	//生成文件url
	urls := getUploadVideoUrls(dirName, maxRes)

	return maxRes, duration, urls, true
}

/*********************************************************
** 函数功能: 获取视频信息
** 日    期: 2022年1月4日17:23:40
**********************************************************/
func getVideoInfo(input string) (dto.VideoInfoData, error) {
	var err error
	var videoData dto.VideoInfoData
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", input)
	// 执行命令，返回命令是否执行成功
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err = cmd.Run()
	if err != nil {
		return videoData, errors.New(fmt.Sprint(err) + ":" + stderr.String())
	}
	// fmt.Println("Result: " + stdout.String())

	err = json.Unmarshal(stdout.Bytes(), &videoData)
	if err != nil {
		return videoData, err
	}

	return videoData, nil
}

/*********************************************************
** 函数功能: 获取宽度支持的最大分辨率
** 日    期: 2022年1月5日9:28:36
** 参    数: 视频宽度
**********************************************************/
func getWidthRes(width int) int {
	//1920*1080
	if width >= 1920 {
		return 1080
	}
	// 1280*720
	if width >= 1280 {
		return 720
	}
	//720*480
	if width >= 720 {
		return 480
	}
	return 360
}

/*********************************************************
** 函数功能: 获取高度支持的最大分辨率
** 日    期: 2022年1月5日9:30:05
** 参    数: 视频高度
**********************************************************/
func getHeigthRes(height int) int {
	//1920*1080
	if height >= 1080 {
		return 1080
	}
	// 1280*720
	if height >= 720 {
		return 720
	}
	//720*480
	if height >= 480 {
		return 480
	}
	return 360
}

/*********************************************************
** 函数功能: 降低码率转到ts
** 日    期: 2022年1月5日12:54:38
** 参    数: 输入文件，输出目录，分辨率
**********************************************************/
func toBitRate(targetDir string, res int) error {
	inputFile := targetDir + "upload.mp4"
	outputFile := targetDir + "temp_" + resName[res] + ".ts"
	cmd := exec.Command("ffmpeg", "-y", "-i", inputFile,
		"-vcodec", "copy", "-acodec", "copy",
		"-vbsf", "h264_mp4toannexb", "-crf",
		"20", outputFile,
	)
	// 执行命令，返回命令是否执行成功
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		util.LogError("transcoding error" + err.Error() + stderr.String())
		return err
	}
	return nil
}

/********************************************************************************
* @功    能：转到hls(指定分辨率)
* @输入参数：文件基础链接，文件目录，文件名，分辨率，资源uuid
* @返 回 值：无
* @日    期：2022年6月5日14:08:31
* @版    本：1.0
********************************************************************************/
func toHlsDifferentRes(baseUrl, baseDir, fileName string, res int, resourceID uint) {
	outputM3U8 := baseDir + "temp_" + resName[res] + ".m3u8"
	inputFile := baseDir + "temp_" + resName[res] + ".ts"
	outputTs := baseDir + fileName + "_" + resName[res] + "%03d.ts"
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c", "copy",
		"-map", "0", "-f", "segment", "-segment_list",
		outputM3U8, "-segment_time", "30", outputTs,
	)
	// fmt.Println(cmd.Args) //查看当前执行命令
	// 执行命令，返回命令是否执行成功
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		util.LogError("transcoding error" + err.Error() + stderr.String())
		return
	}

	//重写m3u8文件
	rewriteDifferentRes(baseUrl, baseDir, fileName, res, resourceID)
}

/********************************************************************************
* @功    能：重写m3u8文件(指定分辨率)
* @输入参数：文件基础链接，文件目录，文件名，分辨率，资源id
* @返 回 值：无
* @日    期：2022年6月5日14:29:36
* @版    本：1.0
********************************************************************************/
func rewriteDifferentRes(baseUrl, baseDir, fileName string, res int, resourceID uint) {
	newM3U8 := baseDir + resName[res] + ".m3u8"
	oldM3U8 := baseDir + "temp_" + resName[res] + ".m3u8"
	file, err := os.OpenFile(oldM3U8, os.O_RDONLY, 0666)
	if err != nil {
		util.LogError("open file filed." + err.Error())
		return
	}
	//创建目标文件
	newFile, _ := os.OpenFile(newM3U8, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	//文件列表
	fileList := list.New()
	//读取文件内容到io中
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//读到末尾
			if err == io.EOF {
				break
			} else {
				util.LogError("transcoding err " + err.Error())
				return
			}
		}
		//根据关键词覆盖当前行
		if strings.Contains(line, ".ts") {
			fileList.PushBack(strings.Replace(line, "\n", "", -1))
			newFile.WriteString(baseUrl + line)
		} else {
			newFile.WriteString(line)
		}
	}

	//关闭文件
	file.Close()
	newFile.Close()

	//将生成的m3u8文件加入文件列表
	fileList.PushBack(resName[res] + ".m3u8")

	if viper.GetBool("file.oss") {
		//将文件上传到oss
		success := UploadHlsToOSS(fileName, fileList)
		if !success {
			//上传失败，调用未通过审核
			videoReviewFail(resourceID, common.ProcessingFail)
			return
		}
	}
}

/*********************************************************
** 函数功能: 转到目标分辨率
** 日    期: 2022年2月13日17:46:59
** 参    数: 文件目录,分辨率
**********************************************************/
func toTargetRes(baseDir string, res int) (string, error) {
	inputFile := baseDir + "upload.mp4"
	outputFile := baseDir + "temp_" + resName[res] + ".mp4"
	cmd := exec.Command("ffmpeg", "-i", inputFile,
		"-vf", "scale="+resInfo[res], outputFile,
		"-hide_banner",
	)
	// 执行命令，返回命令是否执行成功
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	return outputFile, nil
}
