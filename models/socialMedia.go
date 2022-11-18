package models

import "github.com/jinzhu/gorm"

type SocialMed struct {
	gorm.Model
	Name             string `gorm:"not null" json:"name" form:"name" valid:"required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required"`
	UserId           uint   `json:"UserId" gorm:"index"`
	User             []User `gorm:"foreignKey:UserId" json:"user"`
}
