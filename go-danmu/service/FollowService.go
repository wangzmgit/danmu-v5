package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 关注
** 日    期:2021/11/10
**********************************************************/
func FollowingService(fid uint, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//验证关注的人是否存在
	DB := common.GetDB()
	if !isUserExist(DB, fid) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.UserNotExist
		return res
	}
	//没有记录自动创建记录
	var followInfo model.Follow
	DB.Where("uid = ? and fid = ?", uid, fid).Attrs(model.Follow{Uid: uid, Fid: fid}).FirstOrCreate(&followInfo)
	return res
}

/*********************************************************
** 函数功能: 取消关注
** 日    期:2021/11/11
**********************************************************/
func UnFollowService(fid uint, uid uint) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("uid = ? and fid = ?", uid, fid).Delete(&model.Follow{})
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取关注状态
** 日    期:2021/11/11
**********************************************************/
func GetFollowStatusService(uid uint, fid uint) response.ResponseStruct {
	DB := common.GetDB()
	follow := isFollow(DB, uid, fid)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"follow": follow},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过UID获取关注列表
** 日    期:2021/11/11
**********************************************************/
func GetFollowingByIDService(uid interface{}, page int, pageSize int) response.ResponseStruct {
	var users []vo.FollowVo
	DB := common.GetDB()
	sql := "select id,name,sign,avatar from users where deleted_at is null and id in " +
		"(select fid from follows where uid = ? and deleted_at is null) limit ? offset ?"
	DB.Raw(sql, uid, pageSize, (page-1)*pageSize).Scan(&users)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"users": users},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过UID获取粉丝列表
** 日    期:2021/11/11
**********************************************************/
func GetFollowersByIDService(uid interface{}, page int, pageSize int) response.ResponseStruct {
	var users []vo.FollowVo
	DB := common.GetDB()
	sql := "select id,name,sign,avatar from users where deleted_at is null and id in " +
		"(select uid from follows where fid = ? and deleted_at is null) limit ? offset ?"
	DB.Raw(sql, uid, pageSize, (page-1)*pageSize).Scan(&users)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"users": users},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过UID获取粉丝数
** 日    期:2021/7/25
**********************************************************/
func GetFollowCountService(uid int) response.ResponseStruct {
	var following int64
	var followers int64
	DB := common.GetDB()
	DB.Model(&model.Follow{}).Where("uid = ?", uid).Count(&following)
	DB.Model(&model.Follow{}).Where("fid = ?", uid).Count(&followers)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"following": following, "followers": followers},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 是否关注
** 日    期:2021/7/24
**********************************************************/
func isFollow(db *gorm.DB, uid uint, fid uint) bool {
	var follow model.Follow
	db.Where("uid = ? and fid = ?", uid, fid).First(&follow)
	if follow.ID != 0 {
		return true
	}
	return false
}
