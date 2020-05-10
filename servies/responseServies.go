package servies

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context,code int,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":data,
	})
	return
}
