package servies

import (
	"2_ihome/model"
	"2_ihome/pkg"
	"errors"

	"github.com/jinzhu/copier"
)

type UserRegister struct {
	Name 	 string		`json:"name" xml:"name" form:"name" valid:"required#请输入用户名"`
	Mobile 	 string		`json:"mobile" xml:"mobile" form:"mobile" valid:"required|phone#请输入手机|手机格式非法"`
	Email 	 string		`json:"email" xml:"email" form:"email" valid:"required|email#请输入邮箱|邮箱格式非法"`
	Password string		`json:"password" xml:"password" form:"password" valid:"required|length:6,12#请输入密码|密码长度非法"`
	Password2 string	`json:"password2" xml:"password2" form:"password2" valid:"required|same:Password#||两次密码不一致，请重新输入"`
}

type UserLogin struct {
	LoginName 	 string		`json:"name" xml:"name" form:"name" valid:"required#请输入登录名"`
	Password     string		`json:"password" xml:"password" form:"password" valid:"required|length:6,12#请输入密码|密码长度非法"`
}

func (userRegister *UserRegister) CreateUser() (model.User,error) {
	user := model.User{}
	err := copier.Copy(&user, userRegister)
	if err != nil {
		return user, err
	}
	passwordHash := pkg.Md5Encrept(user.Password)
	user.Password = passwordHash

	err = user.Create()
	if err != nil {
		return user, err
	}
	return user,nil
}

func (userLogin * UserLogin) Login () (string, error) {
	user := model.User{}
	user.Email = userLogin.LoginName
	user.Mobile = userLogin.LoginName
	user.Name = userLogin.LoginName
	user.GetUserByName()

	if user.ID <=0 {
		return "",errors.New("用户不存在")

	}
	if user.Password == pkg.Md5Encrept(userLogin.Password) {
		token, _ := pkg.GenToken( user.Name,user.Mobile)
		return token,nil
	}

	return "",errors.New("密码错误")
}