package vo

import (
	"time"
)

//获取审核状态的视频信息
type ReviewVideoVo struct {
	ID        uint   `json:"vid"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Desc      string `json:"desc"`
	State     int    `json:"state"`
	Partition string `json:"partition"`
	Copyright bool   `json:"copyright"`
}

//审核视频列表
type ReviewListVo struct {
	ID        uint      `json:"vid"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Desc      string    `json:"desc"`
	Uid       uint      `json:"uid"`
	Copyright bool      `json:"copyright"`
	Partition string    `json:"partition"`
}
