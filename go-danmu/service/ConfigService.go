package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 获取服务信息
** 日    期: 2022年7月5日19:09:48
**********************************************************/
func GetServerInfoService() response.ResponseStruct {
	var info vo.ServerInfoVo

	//版本、redis、ffmpeg
	info.Version = common.Version
	info.Redis = getRedisStatus()
	info.FFmpeg = getFFmpegStatus()
	info.MySQl = getMysqlStatus()
	info.Web = getWebInfo()

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"info": info},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取redis状态
** 日    期: 2021年11月26日
**********************************************************/
func getRedisStatus() string {
	Cache := common.GetCache()
	if Cache != nil {
		_, err := Cache.Ping().Result()
		if err != nil {
			return ""
		}

		info, _ := Cache.Info().Result()
		pattern := `redis_version:(.*)`
		reg := regexp.MustCompile(pattern)
		return reg.FindStringSubmatch(info)[1]
	}

	return ""
}

/*********************************************************
** 函数功能: 获取ffmpeg状态
** 日    期: 2021年11月26日
**********************************************************/
func getFFmpegStatus() string {
	cmd := exec.Command("ffmpeg", "-version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ""
	}

	version := out.String()
	pattern := `version ((n?[0-9]|\.)*)`
	reg := regexp.MustCompile(pattern)
	return reg.FindStringSubmatch(version)[1]
}

/*********************************************************
** 函数功能: 获取mysql状态
** 日    期: 2022年7月5日19:50:04
**********************************************************/
func getMysqlStatus() string {
	var version string
	DB := common.GetDB()
	DB.Raw("select version()").Pluck("version()", &version)
	return version
}

/*********************************************************
** 函数功能: 获取前端信息
** 日    期: 2022年7月5日19:54:32
**********************************************************/
func getWebInfo() vo.FEInfoVo {
	var feInfo vo.FEInfoVo
	infoFile, err := os.Open(viper.GetString("server.fe_path") + "/info.json")
	if err != nil {
		println(err.Error())
		return feInfo
	}
	defer infoFile.Close()

	byteValue, _ := ioutil.ReadAll(infoFile)

	json.Unmarshal(byteValue, &feInfo)

	return feInfo
}
