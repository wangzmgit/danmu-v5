package util

import (
	"os"
)

/********************************************************************************
* @功    能：判断文件夹或文件是否存在
* @输入参数：path 文件名称；文件夹或文件路径
* @返 回 值：文件是否存在，错误
* @日    期：2022年2月28日18:07:30
* @版    本：1.0
********************************************************************************/
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/********************************************************************************
* @功    能：检查并创建需要的目录
* @输入参数：无
* @返 回 值：无
* @日    期：2022年5月22日18:13:05
* @版    本：1.0
********************************************************************************/
func CheckFolder() {
	if exist, _ := pathExists("./file/"); !exist {
		os.Mkdir("./file", os.ModePerm)
	}

	var folderNames = [...]string{"avatar", "carousel", "cover", "logs", "video"}
	for _, item := range folderNames {
		if exist, _ := pathExists("./file/" + item); !exist {
			os.Mkdir("./file/"+item, os.ModePerm)
		}
	}
}
