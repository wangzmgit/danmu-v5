package service

import (
	"net/http"

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
** 函数功能: 添加收藏夹
** 日    期: 2022年6月18日11:35:57
**********************************************************/
func AddCollectionService(collection dto.CollectionDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var newCollection = model.Collection{
		Uid:  uid,
		Name: collection.Name,
	}
	DB := common.GetDB()
	if err := DB.Create(&newCollection).Error; err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.CreateCollectionFail
		return res
	}

	res.Data = gin.H{"id": newCollection.ID}
	return res
}

/*********************************************************
** 函数功能: 修改收藏夹
** 日    期: 2022年6月19日21:02:12
**********************************************************/
func ModifyCollectionService(modify dto.ModifyCollectionDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	//查询收藏夹名是否存在
	var collection model.Collection
	DB.First(&collection, modify.ID)
	if collection.ID != 0 && collection.Uid != uid {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.CollectionNotExist
		return res
	}

	err := DB.Model(model.Collection{}).Where("id = ?", modify.ID).Updates(
		map[string]interface{}{
			"name":  modify.Name,
			"cover": modify.Cover,
			"desc":  modify.Desc,
			"open":  modify.Open,
		},
	).Error
	if err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.ModifyFail
		return res
	}
	return res
}

/*********************************************************
** 函数功能: 删除收藏夹
** 日    期: 2022年6月19日21:21:15
**********************************************************/
func DeleteCollectionService(vid uint, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	DB.Where("id = ? and uid = ?", vid, uid).Delete(&model.Collection{})
	return res
}

/*********************************************************
** 函数功能: 获取收藏夹列表
** 日    期: 2022年6月18日11:20:48
**********************************************************/
func GetCollectionListService(uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	var collections []vo.CollectionVo
	DB.Model(&model.Collection{}).Where("uid = ?", uid).Select("id,cover,name,`desc`,open,created_at").Scan(&collections)

	res.Data = gin.H{"collections": collections}

	return res
}

/*********************************************************
** 函数功能: 添加收藏
** 日    期: 2022年6月18日14:42:50
**********************************************************/
func CollectService(collect dto.AddCollectDto, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//验证视频是否存在
	DB := common.GetDB()
	if !isVideoExist(DB, collect.VID) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.VideoNotExist
		return res
	}
	//处理收藏部分
	addVideoToCollection(DB, collect, uid)
	//处理取消收藏部分
	DB.Where("uid = ? and vid = ? and collection_id in ?", uid, collect.VID, collect.CancelList).Delete(&model.Collect{})

	return res
}

/*********************************************************
** 函数功能: 获取收藏信息
** 日    期: 2022年6月18日16:02:412022年6月18日16:01:12
**********************************************************/
func GetCollectInfoService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()

	res.Data = gin.H{"collect": getVideoCollection(DB, uid, vid)}
	return res
}

/*********************************************************
** 函数功能: 获取收藏夹信息
** 日    期: 2022年6月20日09:32:46
**********************************************************/
func GetCollectionByIDService(uid uint, id int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	//验证视频是否存在
	DB := common.GetDB()
	var collection model.Collection
	DB.First(&collection, id)               //查询收藏夹信息
	user := getUserInfo(DB, collection.Uid) //获取所有者信息
	if collection.Open || collection.Uid == uid {
		res.Data = gin.H{"collection": vo.ToCollectionInfoVo(collection, user)}
	}

	return res
}

/*********************************************************
** 函数功能: 获取收藏列表
** 日    期: 2022年6月20日12:50:02
**********************************************************/
func GetCollectVideoService(uid uint, id, page, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var total int64
	var collects []vo.CollectVideoVo
	var collection model.Collection

	DB := common.GetDB()
	DB.First(&collection, id)
	if !collection.Open && collection.Uid != uid {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.CollectionNotExist
		return res
	}

	DB.Model(&model.Collect{}).Where("collection_id = ?", id).Count(&total)
	sql := "select id,cover,title,`desc`,clicks from videos where deleted_at is NULL and review = ? and " +
		"id in (select vid from collects where deleted_at is NULL and collection_id = ?) limit ? offset ?"
	DB.Raw(sql, common.AuditApproved, id, pageSize, (page-1)*pageSize).Scan(&collects)

	res.Data = gin.H{"count": total, "videos": collects}
	return res
}

/*********************************************************
** 函数功能: 是否已经收藏
** 日    期: 2021/7/22
**********************************************************/
func isCollect(db *gorm.DB, uid uint, vid interface{}) bool {
	// bool 返回值 是否收藏
	var favorites model.Collect
	db.Where("uid = ? AND vid = ?", uid, vid).First(&favorites)
	if favorites.ID == 0 {
		return false
	}

	return true
}

/*********************************************************
** 函数功能: 获取视频所在的收藏夹id
** 日    期: 2022年6月18日15:02:16
**********************************************************/
func getVideoCollection(db *gorm.DB, uid uint, vid interface{}) []uint {
	var id []uint
	db.Model(&model.Collect{}).Where("uid = ? AND vid = ?", uid, vid).Pluck("collection_id", &id)
	return id
}

/*********************************************************
** 函数功能: 添加到收藏夹
** 日    期: 2022年6月18日15:02:16
**********************************************************/
func addVideoToCollection(db *gorm.DB, collect dto.AddCollectDto, uid uint) {
	// 取用户要添加的收藏夹id数组和已经存在收藏夹id数组的差集
	idList := util.DifferenceSet(collect.AddList, getVideoCollection(db, uid, collect.VID))
	dataLen := len(idList)
	if dataLen > 0 {
		//处理要写入的数据
		newCollects := make([]model.Collect, len(idList))
		for i := 0; i < len(idList); i++ {
			newCollects[i].Uid = uid
			newCollects[i].Vid = collect.VID
			newCollects[i].CollectionID = idList[i]
		}

		db.Create(&newCollects)
	}
}
