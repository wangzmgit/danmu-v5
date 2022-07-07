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
** 函数功能: 通过视频ID获取点赞收藏数据
** 日    期: 2022年5月29日15:33:56
**********************************************************/
func GetArchiveByVIDService(vid int, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	//验证点赞是否存在
	DB := common.GetDB()
	isLike, _ := isLike(DB, uid, vid)
	isCollect := isCollect(DB, uid, vid)
	res.Data = gin.H{"archive": vo.ArchiveVo{
		IsLike:    isLike,
		IsCollect: isCollect,
	}}

	return res
}

/*********************************************************
** 函数功能: 获取点赞和收藏数据
** 日    期: 2021/11/6
**********************************************************/
func getCollectAndLikeCount(db *gorm.DB, vid int) (int64, int64) {
	var like int64
	var collect int64
	collectSQL := "select count(id) from (select id from collects where vid = ? and deleted_at is null group by uid) a"
	db.Model(&model.Liked{}).Where("vid = ? and `is_like` = 1", vid).Count(&like)
	db.Raw(collectSQL, vid).Pluck("count(id)", &collect)
	return like, collect
}
