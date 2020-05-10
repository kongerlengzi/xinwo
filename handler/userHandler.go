package handler

import (
	"2_ihome/model"
	"2_ihome/pkg"
	"2_ihome/servies"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context)  {
	userRegister := servies.UserRegister{}
	user := model.User{}
	err := ctx.ShouldBind(&userRegister)
	if err != nil {
		servies.Response(ctx,100,"参数缺少")
		return
	}
	errStr := pkg.ValidateParams(userRegister)
	if errStr != "" {
		servies.Response(ctx,200,errStr)
		return
	}

	if userRegister.Password != userRegister.Password2 {
		servies.Response(ctx,600,err)
		return
	}
    user.Email = userRegister.Email
	user.GetUserByName()
	if user.ID >0 {
		servies.Response(ctx,300,"用户已经存在!")
		return
	}
	user,err = userRegister.CreateUser()
	if err != nil {
		servies.Response(ctx,400,"注册出错")
		return
	}
	servies.Response(ctx,500,user)
	return
}

func Login(ctx *gin.Context)  {
	userLogin := servies.UserLogin{}

	err := ctx.ShouldBind(&userLogin)
	if err != nil {
		servies.Response(ctx,200,"")
		return
	}
	errStr := pkg.ValidateParams(userLogin)
	if errStr != "" {
		servies.Response(ctx,200,errStr)
		return
	}

	token,err := userLogin.Login()
	if err != nil {
		servies.Response(ctx,200,err)
		return
	}

	servies.Response(ctx,200,token)
	return
}

func UpdateUserProfile(ctx *gin.Context)  {

}

func GetUser(ctx *gin.Context)  {
	user := model.User{}
	param := ctx.Param("id")
	user.Name = param
	user.GetUserByName()
	servies.Response(ctx,200,user)
	return
}