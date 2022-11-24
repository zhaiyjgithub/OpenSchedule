package user

import (
	"OpenSchedule/model/userModel"
	"OpenSchedule/response"
	"OpenSchedule/router"
	"OpenSchedule/service/userService"
	"OpenSchedule/utils"
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
	b.Handle(http.MethodPost, router.Login, "Login")
	b.Handle(http.MethodPost, router.CreateSubUser, "CreateSubUser")
	b.Handle(http.MethodPost, router.GetSubUsers, "GetSubUsers")
	b.Handle(http.MethodPost, router.UpdateSubUserPhone, "UpdateSubUserPhone")
	b.Handle(http.MethodPost, router.UpdateUserProfile, "UpdateUserProfile")
}

func (c *Controller) CreateUser() {
	type Param struct {
		FirstName string ` json:"firstName"`
		LastName  string ` json:"lastName"`
		Gender    string `validate:"required,oneof=F M" json:"gender"`
		Birthday  string `validate:"required,len=10" json:"birthday"`
		Email     string `validate:"email,required" json:"email"`
		Password  string `json:"password"`
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	u := userModel.User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		Birthday:  p.Birthday,
		Email:     p.Email,
		Password:  p.Password,
	}
	err, newUser := c.UserService.CreateUser(u)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "create failed", nil)
	} else {
		response.Success(c.Ctx, response.Successful, newUser)
	}
}

func (c *Controller) GetUserByEmail() {
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

func (c *Controller) Login() {
	type Param struct {
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required,len=32" json:"password"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	u := c.UserService.GetUserByEmail(p.Email)
	if u.Password == p.Password {
		response.Success(c.Ctx, response.Successful, u)
	} else {
		response.Fail(c.Ctx, response.Error, "Login failed", nil)
	}
}

func (c *Controller) CreateSubUser() {
	type Param struct {
		FirstName string `json:"firstName,required"`
		LastName  string `json:"lastName,required"`
		Email     string `json:"email,omitempty"`
		Phone     string `json:"phone,required"`
		Birthday  string `json:"birthday,required" validate:"datetime=2006-01-02"`
		Gender    string `json:"gender,required"`
		UserID    int    `json:"userID,required"`
		IsLegal   bool   `json:"is_legal,required"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	u := userModel.SubUsers{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     p.Phone,
		Birthday:  p.Birthday,
		UserID:    p.UserID,
		Gender:    p.Gender,
		IsLegal:   p.IsLegal,
	}
	err = c.UserService.CreateSubUser(u)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) UpdateUserProfile()  {
	var p userModel.UserProfile
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	err := c.UserService.UpdateUserProfile(p)
	response.Success(c.Ctx, response.Successful, err)
}

func (c *Controller) GetSubUsers() {
	type Param struct {
		UserID int `json:"userID,required"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	u := c.UserService.GetSubUsers(p.UserID)
	response.Success(c.Ctx, response.Successful, u)
}

func (c *Controller) UpdateSubUserPhone() {
	type Param struct {
		UserID int    `json:"user_id,required"`
		Phone  string `json:"phone,required"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	err = c.UserService.UpdateSubUserPhone(p.UserID, p.Phone)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}
