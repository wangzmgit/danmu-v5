package vo

import (
	"time"

	"kuukaa.fun/danmu-v5/model"
)

type RoomVo struct {
	Number int `json:"number"`
}

//上传的视频
type UploadVideoVo struct {
	ID        uint      `json:"vid"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Cover     string    `json:"cover"`
	Review    int       `json:"review"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VideoVo struct {
	ID        uint         `json:"vid"`
	Title     string       `json:"title"`
	Cover     string       `json:"cover"`
	Desc      string       `json:"desc"`
	CreatedAt time.Time    `json:"created_at"`
	Copyright bool         `json:"copyright"`
	Author    UserVo       `json:"author"`
	Resource  []ResourceVo `json:"resources"`
	Clicks    int          `json:"clicks"`
}

//推荐视频
type RecommendVideoVo struct {
	ID     uint   `json:"vid"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
	Clicks string `json:"clicks"`
}

type VideoListVo struct {
	ID        uint      `json:"vid"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	CreatedAt time.Time `json:"created_at"`
}

//搜索视频列表
type SearchVideoListVo struct {
	ID          uint      `json:"vid"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Desc        string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
	Copyright   bool      `json:"copyright"`
	Uid         uint      `json:"uid"`
	Partition   string    `json:"partition"` //分区
	PartitionID uint      `json:"-"`
}

//用户上传视频列表
func ToUploadVideoVo(videos []model.Video) []UploadVideoVo {
	length := len(videos)
	newVideos := make([]UploadVideoVo, length)
	for i := 0; i < length; i++ {
		newVideos[i].ID = videos[i].ID
		newVideos[i].Title = videos[i].Title
		newVideos[i].Desc = videos[i].Desc
		newVideos[i].Cover = videos[i].Cover
		newVideos[i].Review = videos[i].Review
		newVideos[i].CreatedAt = videos[i].CreatedAt
		newVideos[i].UpdatedAt = videos[i].UpdatedAt
		newVideos[i].Clicks = videos[i].Clicks
	}
	return newVideos
}

//视频详情
func ToVideoVo(video model.Video, author model.User, resource []model.Resource) VideoVo {

	return VideoVo{
		ID:        video.ID,
		Title:     video.Title,
		Cover:     video.Cover,
		Desc:      video.Desc,
		CreatedAt: video.CreatedAt,
		Copyright: video.Copyright,
		Author: UserVo{
			ID:     author.ID,
			Name:   author.Name,
			Sign:   author.Sign,
			Avatar: author.Avatar,
		},
		Resource: ToReviewResourceListVo(resource),
		Clicks:   video.Clicks,
	}
}
