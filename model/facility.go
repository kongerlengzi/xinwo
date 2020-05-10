package model

import (
	"github.com/jinzhu/gorm"
)

type Facility struct {
	gorm.Model
	Name   string 		`gorm:"unique;not null" json:"name" xml:"name" form:"name"`
}

func GetFacilityByID(id int) Facility {
	var facility Facility
	Db.Where("id=?",id).First(&facility)
	return facility
}

func CreateFacility(name string) Facility {
	facility := Facility{
		Name: name,
	}
	Db.Create(&facility)
	return facility
}
