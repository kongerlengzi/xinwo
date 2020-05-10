package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 	 string		`gorm:"unique;not null size:20" json:"name" xml:"name" form:"name"`
	Mobile 	 string		`gorm:"unique;not null size:15" json:"mobile" xml:"mobile" form:"mobile"`
	Email 	 string		`gorm:"unique;not null" json:"email" xml:"email" form:"email"`
	Password string		`gorm:"not null" json:"-" xml:"-" form:"-"`
}

func (user *User) GetUserByName() {
	Db.Where("name=?",user.Name).Or("mobile=?",user.Mobile).Or("email=?",user.Email).First(&user)
	return
}

func (user *User) Create() error {
	err := Db.Create(user).Error
	return err
}