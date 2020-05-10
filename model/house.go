package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type House struct {
	gorm.Model
	UserID      uint		`gorm:"not null" json:"user_id" xml:"user_id" form:"user_id"`
	Title 		string 		`gorm:"not null" json:"title" xml:"title"`
	Detail 		string		`gorm:"not null" json:"detail" xml:"detail"`
	Address 	string		`gorm:"not null" json:"address" xml:"address"`
	Room    	int		    `gorm:"not null" json:"room" xml:"room"`
	Salon   	int		    `gorm:"not null" json:"salon" xml:"salon"`
	Toilet   	int		    `gorm:"not null" json:"toilet" xml:"toilet"`
	Acreage 	int		    `gorm:"not null" json:"acreage" xml:"acreage"`
	Floor 		int		    `gorm:"not null" json:"floor" xml:"floor"`
	FloorTotal  int		    `gorm:"not null" json:"floor_total" xml:"floor_total"`
	LetPeople   int		    `gorm:"not null" json:"let_people" xml:"let_people"`
	RentType    int		    `gorm:"not null" json:"rent_type" xml:"rent_type"`
	Sublet 		bool		`gorm:"not null" json:"sublet" xml:"sublet"`
	Expires 	time.Time   `gorm:"not null" json:"expires" xml:"expires"`
	CheckIn 	time.Time	`gorm:"not null" json:"check_in" xml:"check_in"`
	Decoration 	string		`gorm:"not null" json:"decoration" xml:"decoration"`
	Direction 	string		`gorm:"not null" json:"direction" xml:"direction"`
	Type 		string		`gorm:"not null" json:"type" xml:"type"`
	ShowTime 	string		`gorm:"not null" json:"show_time" xml:"show_time"`
	PayMethod 	int			`gorm:"not null" json:"pay_method" xml:"pay_method"`
	RentPrice 	int			`gorm:"not null" json:"'rent_price'" xml:"rent_price"`
	RenterSex 	bool		`gorm:"not null" json:"renter_sex" xml:"renter_sex"`
	AllowPet 	bool		`gorm:"not null" json:"allow_pet" xml:"allow_pet"`
	Facilities 	[]Facility 	`gorm:"many2many:house_facilities;" json:"facilities" xml:"facilities"`
	Image 		string	`json:"image" xml:"image"`
}

func (house *House) CreateHouse() House {
	Db.Create(house)
	return *house
}

func GetHouseByUser(page,num int,userid uint) []House {
	var houses []House
	if page <1 {
		page = 1
	}
	switch{
		case num >20:
			num = 20
		case num <0:
			num = 10
	}
	Db.Where("user_id=?",userid).Limit(num).Offset((page-1)*num).Find(&houses)
	return houses
}

func GetHouseIndex() []House {
	var houses []House
	Db.Order("updated_at").Limit(10).Find(&houses)
	return houses
}

func GetHouseById(id int) House {
	var house House
	Db.Where("id=?",id).First(&house)
	return house
}

func GetHousesBySearch(roomCount int,rentPrice int) ([]House,int) {
	var houses []House
	var rentPricelow,rentPricehigh,count int
	if roomCount <=0 {
		roomCount =0
	}
	switch  {
	case rentPrice<=0:
		rentPricelow =0
		rentPricehigh =100000
	case rentPrice ==1:
		rentPricelow =0
		rentPricehigh =1500
	case rentPrice ==2:
		rentPricelow =1500
		rentPricehigh =3000
	}
	Db.Where("room=? AND rent_price BETWEEN ? AND ?",roomCount,rentPricelow ,rentPricehigh).Find(&houses).Count(&count)
	return houses,count
}