package config

import (
	"github.com/go-ini/ini"
	"log"
)

var IhomeConfig IhomeConf

type IhomeConf struct {
	MysqlConf		  `ini:"mysql"`
	IhomeAPP          `ini:"ihome"`

}

type MysqlConf struct {
	User 	 string   `ini:"User"`
	Password string   `ini:"Password"`
	Dbname 	 string   `ini:"Dbname"`
	Host	 string   `ini:"Host"`
}

type IhomeAPP struct {
	JwtSecret []byte  `ini:"JwtSecret"`
}

func init()  {
	cfg,err := ini.Load("E:/go_work/Project/19_ihome/config/conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	cfg.MapTo(&IhomeConfig)
}