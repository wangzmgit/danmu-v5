package dto

//上传视频信息
type UploadVideoDto struct {
	Title     string
	Cover     string
	Desc      string
	Copyright bool
	Partition uint //分区ID
}

//修改视频信息
type ModifyVideoInfoDto struct {
	VID       uint
	Title     string
	Cover     string
	Desc      string
	Copyright bool
}

//修改资源标题
type ModifyResourceTitleDto struct {
	ID    uint
	Title string
}

//获取视频列表
type GetVideoListDto struct {
	Page      int
	PageSize  int
	Partition int
}
