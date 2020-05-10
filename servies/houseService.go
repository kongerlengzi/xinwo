package servies

import (
	"2_ihome/model"
	"2_ihome/pkg"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type HouseService struct {
	Userid      uint
	Title 		string 		`json:"title" xml:"title" form:"title" valid:"required"`
	Detail 		string		`json:"detail" xml:"detail" form:"detail" valid:"required"`
	Address 	string		`json:"address" xml:"address" form:"address" valid:"required"`
	Room    	int		    `json:"room" xml:"room" form:"room" valid:"required"`
	Salon   	int		    `json:"salon" xml:"salon" form:"salon" valid:"required"`
	Toilet   	int		    `json:"toilet" xml:"toilet" form:"toilet" valid:"required"`
	Acreage 	int		    `json:"acreage" xml:"acreage" form:"acreage" valid:"required"`
	Floor 		int		    `json:"floor" xml:"floor" form:"floor" valid:"required"`
	FloorTotal  int		    `json:"floor_total" xml:"floor_total" form:"floor_total" valid:"required"`
	LetPeople   int		    `json:"let_people" xml:"let_people" form:"let_people" valid:"required"`
	RentType    int		    `json:"rent_type" xml:"rent_type" form:"rent_type" valid:"required"`
	Sublet 		bool		`json:"sublet" xml:"sublet" form:"sublet"`
	ExpiresS 	string      `json:"expires" xml:"expires" form:"expires" valid:"required|date"`
	CheckInS	string      `json:"check_in" xml:"check_in" form:"check_in" valid:"required|date"`
	Decoration 	string		`json:"decoration" xml:"decoration" form:"decoration" valid:"required"`
	Direction 	string		`json:"direction" xml:"direction" form:"direction" valid:"required|in:朝东,朝西,朝北,朝南"`
	Type 		string	 	`json:"type" xml:"type" form:"type" valid:"required"`
	ShowTime 	string	    `json:"show_time" xml:"show_time" form:"show_time" valid:"required"`
	PayMethod 	int			`json:"pay_method" xml:"pay_method" form:"pay_method" valid:"required"`
	RentPrice 	int			`json:"rent_price" xml:"rent_price" form:"rent_price" valid:"required"`
	RenterSex 	bool		`json:"renter_sex" xml:"renter_sex" form:"renter_sex" valid:"required"`
	AllowPet 	bool		`json:"allow_pet" xml:"allow_pet" form:"allow_pet" valid:"required"`
	Facilities 	string      `json:"facilities" xml:"facilities" form:"facilities" valid:"required"`
	Image 		string      `json:"image" xml:"image" form:"image" valid:"required"`
}

func (houseService *HouseService) Create() model.House {
	var house model.House
	var facilities []model.Facility
	copier.Copy(&house,houseService)
	Expires := pkg.StringTotime(houseService.ExpiresS)
	CheckIn := pkg.StringTotime(houseService.CheckInS)
	house.CheckIn = CheckIn
	house.Expires = Expires
	facilitiesStr := houseService.Facilities
	sep:=","
	arr:= strings.Split(facilitiesStr,sep)
	for _,str := range arr{
		id,_ := strconv.Atoi(str)
		fmt.Println(id)
		facility := model.GetFacilityByID(id)
		if facility.ID <=0{
			continue
		}
		facilities = append(facilities, facility)
	}
	house.Facilities = facilities
	house.UserID = houseService.Userid
	createHouse := house.CreateHouse()
	return createHouse
}