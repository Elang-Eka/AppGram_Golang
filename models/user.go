package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"not null;unique_index" json:"email" form:"email" valid:"email,required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
