package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 获取视频状态
** 日    期: 2021/11/10
**********************************************************/
func GetVideoStatusService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var video model.Video
	DB := common.GetDB()
	DB.Model(&model.Video{}).Where("id = ? and uid = ?", vid, uid).First(&video)
	if video.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.VideoNotExist
		return res
	}
	//获取分区名称
	partition := getPartitionName(DB, video.PartitionID)

	res.Data = gin.H{"video": vo.ReviewVideoVo{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		State:     video.Review,
		Partition: partition,
		Copyright: video.Copyright,
	}}
	return res
}

/*********************************************************
** 函数功能: 获取审核列表
** 日    期: 2021年11月12日14:52:09
**********************************************************/
func GetReviewListService(page int, pageSize int) response.ResponseStruct {
	var total int64
	var videos []vo.ReviewListVo

	DB := common.GetDB()
	//统计数量
	DB.Model(&model.Video{}).Where("review = ?", common.WaitingReview).Count(&total)
	sql := "select videos.id,videos.created_at,title,cover,`desc`,uid,copyright,`partitions`.content as `partition` " +
		"from videos,`partitions` where videos.deleted_at is null and `partitions`.id = videos.partition_id " +
		"and review = ? limit ? offset ?"

	DB.Raw(sql, common.WaitingReview, pageSize, (page-1)*pageSize).Scan(&videos)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "videos": videos},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过视频ID获取待审核视频资源
** 日    期: 2022年1月6日15:16:50
**********************************************************/
func GetReviewVideoByIDService(vid int) response.ResponseStruct {
	var videos []model.Resource
	DB := common.GetDB()
	DB.Model(&model.Resource{}).Where("vid = ?", vid).Find(&videos)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"video": vo.ToReviewResourceListVo(videos)},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 提交审核
** 日    期: 2022年7月1日12:12:02
**********************************************************/
func SubmitReviewService(vid, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	//查询视频是否属于用户
	if getVideoAuthorID(DB, vid) != uid {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.VideoNotExist
		return res
	}

	//查询是否存在视频资源
	var resource model.Resource
	DB.Where("vid = ? ", vid).First(&resource)
	if resource.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.NotExistResources
		return res
	}

	//更新视频审核状态
	if err := DB.Model(&model.Video{}).Where("id = ?", vid).Updates(
		map[string]interface{}{"review": common.WaitingReview},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}
	return res
}

/*********************************************************
** 函数功能: 审核视频
** 日    期: 2021/8/4
**********************************************************/
func ReviewVideoService(review dto.ReviewDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	status := common.AuditApproved
	if !review.Status {
		status = common.WrongVideoInfo
	}

	DB := common.GetDB()
	//查询是否存在视频资源
	var resource model.Resource
	DB.Where("vid = ? and review = ?", review.ID, common.AuditApproved).First(&resource)
	if resource.ID == 0 && review.Status {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.NotExistResources
		return res
	}
	//更新视频审核状态
	if err := DB.Model(&model.Video{}).Where("id = ?", review.ID).Updates(
		map[string]interface{}{"review": status},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}
	return res
}

/*********************************************************
** 函数功能: 审核资源
** 日    期: 2022年6月10日09:22:39
**********************************************************/
func ReviewResourceService(review dto.ReviewDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	status := common.AuditApproved
	if !review.Status {
		status = common.WrongVideoContent
	}

	DB := common.GetDB()
	if err := DB.Model(&model.Resource{}).Where("id = ?", review.ID).Updates(
		map[string]interface{}{"review": status},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}
	return res
}

/*********************************************************
** 函数功能: 视频处理失败未通过审核
** 日    期: 2022年1月5日17:12:27
**********************************************************/
func videoReviewFail(resourceID uint, state int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	if err := DB.Model(&model.Resource{}).Where("id = ?", resourceID).Updates(
		map[string]interface{}{
			"review": state,
		},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}
	return res
}
