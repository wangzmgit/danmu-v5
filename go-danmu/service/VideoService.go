package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/dto"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/response"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/vo"
)

/*********************************************************
** 函数功能: 上传视频信息
** 日    期: 2022年5月23日14:17:44
**********************************************************/
func UploadVideoInfoService(video dto.UploadVideoDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	if !isSubpartition(DB, video.Partition) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.PartitionNotExist
		return res
	}

	newVideo := model.Video{
		Title:       video.Title,
		Cover:       video.Cover,
		Desc:        video.Desc,
		Copyright:   video.Copyright,
		Uid:         uid,
		PartitionID: video.Partition,
		Review:      common.CreateVideo,
	}

	if err := DB.Create(&newVideo).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UploadFail
		return res
	}
	res.Data = gin.H{"vid": newVideo.ID}
	return res
}

/*********************************************************
** 函数功能: 修改视频信息
** 日    期: 2021年12月9日17:50:35
**********************************************************/
func ModifyVideoInfoService(video dto.ModifyVideoInfoDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()

	if err := DB.Model(&model.Video{}).Where("id = ? and uid = ?", video.VID, uid).Updates(
		map[string]interface{}{
			"title":     video.Title,
			"cover":     video.Cover,
			"desc":      video.Desc,
			"copyright": video.Copyright,
			"review":    common.WaitingReview,
		},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.ModifyFail
		return res
	}

	return res
}

/*********************************************************
** 函数功能: 修改资源标题
** 日    期: 2022年7月7日11:59:55
**********************************************************/
func ModifyResourceTitleService(id, uid uint, title string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	//查询资源作者
	resource := getVideoResourceByID(DB, id)
	authorId := getVideoAuthorID(DB, resource.Vid)
	if authorId != uid {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.NotExistResources
		return res
	}

	if err := DB.Model(&model.Resource{}).Where("id = ? ", id).Updates(
		map[string]interface{}{"title": title},
	).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.ModifyFail
		return res
	}

	return res
}

/*********************************************************
** 函数功能: 删除视频
** 日    期: 2022年5月23日15:14:53
**********************************************************/
func DeleteVideoService(vid uint, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	if getVideoAuthorID(DB, vid) != uid {
		//该视频不属于这个用户
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.DeleteFail
		return res
	}
	DB.Where("id = ?", vid).Delete(&model.Video{})
	return res
}

/*********************************************************
** 函数功能: 删除资源
** 日    期: 2022年5月23日15:14:53
**********************************************************/
func DeleteResourceService(id uint, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	vid := getVidByResource(DB, id)
	if getVideoAuthorID(DB, vid) != uid {
		//该视频不属于这个用户
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.DeleteFail
		return res
	}
	DB.Where("id = ?", id).Delete(&model.Resource{})
	return res
}

/*********************************************************
** 函数功能: 获取自己上传的视频
** 日    期: 2021年11月17日12:59:09
**********************************************************/
func GetUploadVideoListService(page int, pageSize int, uid uint) response.ResponseStruct {
	var total int64
	var videos []model.Video

	DB := common.GetDB()
	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)
	Pagination.Where("uid = ?", uid).Find(&videos)
	DB.Model(&model.Video{}).Where("uid = ?", uid).Count(&total)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "videos": vo.ToUploadVideoVo(videos)},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 用户获取上传资源列表
** 日    期: 2022年6月6日17:28:54
**********************************************************/
func GetUploadResourceListService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	var resources []model.Resource
	if err := DB.Where("vid = ?", vid).Find(&resources).Error; err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UpdateStatusFail
		return res
	}
	res.Data = gin.H{"resources": vo.ToReviewResourceListVo(resources)}
	return res
}

/*********************************************************
** 函数功能: 通过视频ID获取视频
** 日    期: 2021年11月17日12:59:45
**********************************************************/
func GetVideoByIDService(vid int, ip string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var video model.Video
	DB := common.GetDB()
	DB.Model(&model.Video{}).Where("id = ? and review = ?", vid, common.AuditApproved).First(&video)
	if video.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.CheckFailCode
		res.Msg = response.VideoNotExist
		return res
	}
	//获取作者信息
	author := getUserInfo(DB, video.Uid)
	//获取视频资源
	resource := getVideoResource(DB, uint(vid))
	//获取视频统计数据
	like, collect := getCollectAndLikeCount(DB, vid)
	//增加播放量(一个ip在同一个视频下，每30分钟可重新增加1播放量)
	if Cache := common.GetCache(); Cache != nil {
		clicksLimit, _ := Cache.Get(util.VideoClicksLimitKey(vid, ip)).Result()
		if clicksLimit == "" {
			DB.Model(&video).UpdateColumn("clicks", gorm.Expr("clicks + 1"))
			Cache.Set(util.VideoClicksLimitKey(vid, ip), 1, time.Minute*30)
		}
	}

	res.Data = gin.H{
		"video": vo.ToVideoVo(video, author, resource),
		"stat": vo.ArchiveCountVo{
			Like:    like,
			Collect: collect,
		},
	}
	return res
}

/*********************************************************
** 函数功能: 获取推荐视频
** 日    期: 2021年11月17日13:00:01
**********************************************************/
func GetRecommendVideoService() response.ResponseStruct {
	var videos []vo.RecommendVideoVo
	DB := common.GetDB()

	const sql = "select videos.id,title,cover,name as author,clicks from users,videos where " +
		"users.id=videos.uid and review = ? and videos.deleted_at is null order by clicks desc limit 6"

	DB.Raw(sql, common.AuditApproved).Scan(&videos)

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"videos": videos},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取视频列表
** 日    期: 2021年11月17日13:00:26
**********************************************************/
func GetVideoListService(query dto.GetVideoListDto) response.ResponseStruct {
	var videos []vo.VideoListVo

	DB := common.GetDB()
	selectVideo := DB.Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).
		Model(&model.Video{}).Select("id,title,cover,created_at").Where("review = ?", common.AuditApproved)

	if query.Partition == 0 {
		//不传分区参数默认查询全部
		selectVideo.Scan(&videos)
	} else if isSubpartition(DB, uint(query.Partition)) {
		//判断是否为子分区
		selectVideo.Where("partition_id = ?", query.Partition).Scan(&videos)
	} else {
		//获取该分区下的子分区
		list := getSubpartitionList(DB, uint(query.Partition))
		selectVideo.Where("partition_id in (?)", list).Scan(&videos)
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"videos": videos},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 通过用户ID获取视频列表
** 日    期: 2021年11月17日13:00:41
**********************************************************/
func GetVideoListByUIDService(uid int, page int, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var videos []vo.VideoListVo

	DB := common.GetDB()
	if !isUserExist(DB, uint(uid)) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.UserNotExist
		return res
	}
	//记录总数
	var total int64
	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&model.Video{}).Where("review = ? and uid = ?", common.AuditApproved, uid).Count(&total)
	Pagination.Model(&model.Video{}).Select("id,title,cover,created_at").
		Where("review = ? and uid = ?", common.AuditApproved, uid).Scan(&videos)
	res.Data = gin.H{"count": total, "videos": videos}
	return res
}

/*********************************************************
** 函数功能: 搜索视频
** 日    期: 2022年3月24日19:32:17
**********************************************************/
func SearchVideoService(keyword string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var videos []vo.SearchVideoListVo

	DB := common.GetDB()
	DB.Model(model.Video{}).Where("review = ? and (title like ? or id = ?)",
		common.AuditApproved, keyword+"%", util.StringToUint(keyword)).Limit(15).Scan(&videos)

	for i := 0; i < len(videos); i++ {
		videos[i].Partition = getPartitionName(DB, videos[i].PartitionID)
	}

	res.Data = gin.H{"videos": videos}
	return res
}

/*********************************************************
** 函数功能: 管理员获取视频列表
** 日    期: 2021年11月12日15:30:26
**********************************************************/
func AdminGetVideoListService(page, pageSize int) response.ResponseStruct {
	var total int64 //记录总数
	var videos []vo.SearchVideoListVo

	DB := common.GetDB()
	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(model.Video{}).Where("review = ?", common.AuditApproved).Count(&total)
	Pagination.Model(model.Video{}).Where("review = ?", common.AuditApproved).Scan(&videos)

	for i := 0; i < len(videos); i++ {
		videos[i].Partition = getPartitionName(DB, videos[i].PartitionID)
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "videos": videos},
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 管理员删除视频
** 日    期: 2021年11月12日15:33:20
**********************************************************/
func AdminDeleteVideoService(id uint) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ?", id).Delete(&model.Video{})
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
}

/*********************************************************
** 函数功能: 获取视频资源
** 日    期: 2022年1月6日10:33:53
**********************************************************/
func getVideoResource(db *gorm.DB, vid uint) []model.Resource {
	var resource []model.Resource
	db.Where("vid = ? and review = ?", vid, common.AuditApproved).Find(&resource)
	return resource
}

/*********************************************************
** 函数功能: 通过ID获取视频资源
** 日    期: 2022年7月7日12:04:47
**********************************************************/
func getVideoResourceByID(db *gorm.DB, id uint) model.Resource {
	var resource model.Resource
	db.First(&resource, id)
	return resource
}

/*********************************************************
** 函数功能: 视频是否存在
** 日    期:2021/7/22
**********************************************************/
func isVideoExist(db *gorm.DB, vid interface{}) bool {
	var video model.Video
	db.First(&video, vid)
	if video.ID != 0 {
		return true
	}
	return false
}

/*********************************************************
** 函数功能: 获取视频作者uid
** 日    期: 2022年5月23日15:12:57
**********************************************************/
func getVideoAuthorID(db *gorm.DB, vid uint) uint {
	var video model.Video
	db.Where("id = ?", vid).First(&video)
	return video.Uid
}

/*********************************************************
** 函数功能: 通过资源id获取vid
** 日    期: 2022年7月1日10:12:14
**********************************************************/
func getVidByResource(db *gorm.DB, id uint) uint {
	var resource model.Resource
	db.First(&resource, id)
	return resource.Vid
}
