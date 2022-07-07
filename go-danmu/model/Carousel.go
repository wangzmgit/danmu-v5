package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	Img string `gorm:"size:255;"`
	Url string `gorm:"type:varchar(100)"`
}
