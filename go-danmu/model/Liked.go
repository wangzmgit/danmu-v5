package model

import "gorm.io/gorm"

type Liked struct {
	gorm.Model
	Uid    uint  `gorm:"not null"`
	Vid    uint  `gorm:"not null"`
	IsLike bool  `gorm:"default:false"` //是否点赞
	Video  Video `gorm:"ForeignKey:id;AssociationForeignKey:vid"`
}
