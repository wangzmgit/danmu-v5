package service

import (
	"net/http"

	"gorm.io/gorm"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
)

/*********************************************************
** 函数功能: 点赞
** 日    期: 2021/11/11
**********************************************************/
func LikeService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//验证视频是否存在
	DB := common.GetDB()
	if !isVideoExist(DB, vid) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.VideoNotExist
		return res
	}

	//验证是否已经点赞
	_, status := isLike(DB, uid, vid)
	if status == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.IsLike
		return res
	}
	if status == -1 {
		newFavorites := model.Liked{
			Uid:    uid,
			Vid:    uint(vid),
			IsLike: true,
		}
		DB.Create(&newFavorites)
	} else {
		DB.Model(&model.Liked{}).Where("uid = ? AND vid = ?", uid, vid).Update("is_like", true)
	}
	return res
}

/*********************************************************
** 函数功能: 取消赞
** 日    期: 2021/11/11
**********************************************************/
func DislikeService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//验证点赞是否存在
	DB := common.GetDB()
	status, _ := isLike(DB, uid, vid)
	if status {
		DB.Model(&model.Liked{}).Where("uid = ? AND vid = ?", uid, vid).Update("is_like", false)
		return res
	} else {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.NotLike
		return res
	}
}

/*********************************************************
** 函数功能: 是否已经点赞
** 日    期: 2021/7/22
**********************************************************/
func isLike(db *gorm.DB, uid uint, vid interface{}) (bool, int) {
	// bool 返回值 是否点赞
	//不存在返回-1，存在但是没有点赞返回1，已经点赞返回0
	var liked model.Liked
	db.Where("uid = ? AND vid = ?", uid, vid).First(&liked)
	if liked.ID == 0 {
		return false, -1
	} else if !liked.IsLike {
		return false, 1
	}
	return true, 0
}
