package router

import (
	"2_ihome/handler"
	"2_ihome/middlerware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("",handler.Register)
	userRouter.POST("/login",handler.Login)
	userRouter.GET("/:id",middlerware.JWTAuthMiddleware(),handler.GetUser)
	userRouter.PUT("")
}
