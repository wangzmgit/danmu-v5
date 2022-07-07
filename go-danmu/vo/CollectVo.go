package vo

import (
	"time"

	"kuukaa.fun/danmu-v5/model"
)

//收藏夹
type CollectionVo struct {
	ID        uint      `json:"id"`
	Cover     string    `json:"cover"`
	Name      string    `json:"name"` //收藏夹名称
	Desc      string    `json:"desc"` //简介
	Open      bool      `json:"open"` //是否公开
	CreatedAt time.Time `json:"created_at"`
}

//收藏夹信息
type CollectionInfoVo struct {
	ID        uint      `json:"id"`
	Cover     string    `json:"cover"`
	Name      string    `json:"name"` //收藏夹名称
	Desc      string    `json:"desc"` //简介
	CreatedAt time.Time `json:"created_at"`
	Open      bool      `json:"open"` //是否公开
	Author    UserVo    `json:"author"`
}

//收藏的视频
type CollectVideoVo struct {
	ID     uint   `json:"vid"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Desc   string `json:"desc"`
	Clicks int    `json:"clicks"`
}

func ToCollectionInfoVo(collection model.Collection, user model.User) CollectionInfoVo {
	return CollectionInfoVo{
		ID:        collection.ID,
		Cover:     collection.Cover,
		Name:      collection.Name,
		Desc:      collection.Desc,
		CreatedAt: collection.CreatedAt,
		Open:      collection.Open,
		Author: UserVo{
			ID:     user.ID,
			Name:   user.Name,
			Avatar: user.Avatar,
		},
	}
}
