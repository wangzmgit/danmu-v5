package vo

//单个分区
type PartitionVo struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

//所有分区
type AllPartitionVo struct {
	ID           uint          `json:"id"`
	Content      string        `json:"content"`
	Subpartition []PartitionVo `json:"subpartition" gorm:"-"`
}
