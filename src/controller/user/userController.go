package user

import (
	"OpenSchedule/src/model/user"
	"OpenSchedule/src/response"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/userService"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type Controller struct {
	Ctx         iris.Context
	UserService userService.Service
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, router.CreateUser, "CreateUser")
	b.Handle(http.MethodPost, router.GetUserByEmail, "GetUserByEmail")
}

func (c *Controller) CreateUser() {
	type Param struct {
		FirstName	string	`validate:"required" json:"firstName"`
		LastName	string	`validate:"required" json:"lastName"`
		Gender	string	`validate:"required,oneof=F M" json:"gender"`
		Birthday string `validate:"required,len=10" json:"birthday"`
		Email	string	`validate:"email,required" json:"email"`
		Password	string	`json:"password"`
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	u := user.Users{
		FirstName: p.FirstName,
		LastName: p.LastName,
		Gender: p.Gender,
		Birthday: p.Birthday,
		Email: p.Email,
		Password: p.Password,
	}
	err, newUser := c.UserService.CreateUser(u)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "create failed", nil)
	} else {
		response.Success(c.Ctx, response.Successful, newUser)
	}
}

func (c *Controller) GetUserByEmail()  {
	type Param struct {
		Email string `validate:"required, email" json:"email"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	u := c.UserService.GetUserByEmail(p.Email)
	if u.Email != p.Email {
		response.Fail(c.Ctx, response.Error, response.NotFound, nil)
	} else {
		response.Success(c.Ctx, response.Successful, u)
	}
}
