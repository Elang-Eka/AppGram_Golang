package models

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Caption   string `gorm:"not null" json:"caption" form:"caption" valid:"required"`
	Photo_url string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required"`
	User_Id   uint   `json:"User_Id" gorm:"index"`
	User      []User `gorm:"foreignKey:User_Id" json:"user"`
}
