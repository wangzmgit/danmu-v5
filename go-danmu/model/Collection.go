package model

import "gorm.io/gorm"

//收藏夹
type Collection struct {
	gorm.Model
	Uid   uint   `gorm:"not null"`           //所属用户
	Name  string `gorm:"type:varchar(20);"`  //收藏夹名称
	Desc  string `gorm:"type:varchar(150);"` //简介
	Cover string `gorm:"size:255;"`          //封面
	Open  bool   `gorm:"default:false"`      //是否公开
}
