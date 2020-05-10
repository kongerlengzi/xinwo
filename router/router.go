package router

import (
	"2_ihome/middlerware"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	r := gin.Default()
	r.Use(middlerware.LoggerToFile())
	InitUserRouter(r)
	InitHouseRouter(r)

	_ = r.Run(":8085")
}
