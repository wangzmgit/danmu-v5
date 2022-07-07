package model

import "gorm.io/gorm"

type Collect struct {
	gorm.Model
	Uid          uint `gorm:"not null"`
	Vid          uint `gorm:"not null"`
	CollectionID uint `gorm:"default:0"` //所属收藏夹ID
}
