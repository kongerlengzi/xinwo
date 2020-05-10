package handler

import (
	"2_ihome/model"
	"2_ihome/pkg"
	"2_ihome/servies"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateHouse (ctx *gin.Context) {
	var houseService servies.HouseService

	err := ctx.ShouldBind(&houseService)
	if err != nil {
		servies.Response(ctx,100,"参数缺少")
		return
	}

	errStr := pkg.ValidateParams(houseService)
	if errStr != "" {
		fmt.Printf("%v\n",err)
		servies.Response(ctx,600,errStr)
		return
	}

	u :=ctx.MustGet("userid").(uint)
	houseService.Userid = u
	file, err := ctx.FormFile("imagea")
	if err != nil {
		ctx.String(500, "上传图片出错")
	}
	ctx.SaveUploadedFile(file, file.Filename)
	houseService.Image = "/home/ixxa/GoWork/2_ihome/"+file.Filename
	create := houseService.Create()
	servies.Response(ctx,100,create)
	return
}

func GetUserRentHouse(ctx *gin.Context)  {
	userid :=ctx.MustGet("userid").(uint)
	page,_ := strconv.Atoi(ctx.PostForm("page"))
	num,_ := strconv.Atoi(ctx.PostForm("num"))
	houses := model.GetHouseByUser(page, num, userid)
	servies.Response(ctx,100,houses)
	return
}

func GetIndexHouse(ctx *gin.Context)  {
    houses := model.GetHouseIndex()
	servies.Response(ctx,100,houses)
	return
}

func GetHouseById(ctx *gin.Context)  {
	id ,_:= strconv.Atoi(ctx.Param("id"))
	house := model.GetHouseById(id)
	servies.Response(ctx,100,house)
	return
}

func GetHouseBySearch(ctx *gin.Context)  {
	roomCount,_ := strconv.Atoi(ctx.PostForm("roomcount"))
	rentPrice,_ := strconv.Atoi(ctx.PostForm("rentprice"))
	hosues,count := model.GetHousesBySearch(roomCount, rentPrice)
	servies.Response(ctx,100,gin.H{
		"houses":hosues,
		"count":count,
	})
	return
}