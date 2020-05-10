package router

import (
	"2_ihome/handler"
	"2_ihome/middlerware"
	"github.com/gin-gonic/gin"
)

func InitHouseRouter(r *gin.Engine) {
	houseRouter := r.Group("/house")
	houseRouter.POST("",middlerware.JWTAuthMiddleware(),handler.CreateHouse)
	houseRouter.POST("/user",middlerware.JWTAuthMiddleware(),handler.GetUserRentHouse)
	houseRouter.POST("/search",handler.GetHouseBySearch)
	houseRouter.GET("/index",handler.GetIndexHouse)
	houseRouter.GET("/info/:id",handler.GetHouseById)
}
