package vo

//视频点赞收藏数据
type ArchiveCountVo struct {
	Collect int64 `json:"collect"`
	Like    int64 `json:"like"`
}

//用户点赞收藏关注信息
type ArchiveVo struct {
	IsCollect bool `json:"is_collect"`
	IsLike    bool `json:"is_like"`
}
