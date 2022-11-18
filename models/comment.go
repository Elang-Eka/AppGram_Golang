package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	User_id  uint    `json:"user_id" gorm:"index"`
	User     []User  `gorm:"foreignKey:user_id" json:"user"`
	Photo_id uint    `json:"photo_id" gorm:"index"`
	Photo    []Photo `gorm:"foreignKey:photo_id" json:"photo"`
	Message  string  `gorm:"not null" json:"message" form:"message" valid:"required"`
}
