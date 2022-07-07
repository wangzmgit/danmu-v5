package dto

type CollectionDto struct {
	Name string //收藏夹名称
}

type ModifyCollectionDto struct {
	ID    uint
	Cover string //封面图
	Name  string //收藏夹名称
	Desc  string //简介
	Open  bool   //是否公开
}

type AddCollectDto struct {
	VID        uint   //视频ID
	AddList    []uint //添加的收藏夹数组
	CancelList []uint //取消的收藏夹数组
}
