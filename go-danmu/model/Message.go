package model

import "gorm.io/gorm"

type UserMessage struct {
	gorm.Model
	Uid     uint   `gorm:"not null;"` //用户ID
	Fid     uint   `gorm:"not null;"` //关联ID
	FromId  uint   `gorm:"not null;"` // 发送者
	ToId    uint   `gorm:"not null;"` // 接受者
	Content string `gorm:"size:255;"`
	Status  bool   `gorm:"default:false"` //已读状态
}
