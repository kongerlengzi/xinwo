package model

import (
	"2_ihome/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
	err error
)

func init()  {
	Db,err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.IhomeConfig.User,
		config.IhomeConfig.Password,
		config.IhomeConfig.Host,
		config.IhomeConfig.Dbname))
	if err != nil {
		fmt.Printf("连接数据库出错，err： %v,",err)
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	Db.LogMode(true)

	Db.AutoMigrate(&User{},&House{},&Facility{})
}