package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Uid uint `gorm:"not null;index"`
	Fid uint `gorm:"not null;index"`
}
