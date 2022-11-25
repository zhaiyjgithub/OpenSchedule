package user

import (
	"OpenSchedule/model/userModel"
	"OpenSchedule/response"
	"OpenSchedule/utils"
)

func (c *Controller) GetUserInsurance()  {
	type Param struct {
		UserID int `json:"userID" validate:"required"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	ins, err := c.UserService.GetUserInsurance(p.UserID)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, ins)
	}
}

func (c *Controller) UpdateUserInsurance()  {
	type Insurance struct {
		ID int
		UserID int
		PlanID string `validate:"required"`
		MemberID string
		Photo string `validate:"omitempty,url"`
	}
	type Param struct {
		Insurances []Insurance `json:"insurances" validate:"required,dive"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	var ins []userModel.UserInsurance
	for _, in := range p.Insurances {
		ins = append(ins, userModel.UserInsurance{
			ID: in.ID,
			UserID: in.UserID,
			PlanID: in.PlanID,
			MemberID: in.MemberID,
			Photo: in.Photo,
		})
	}
	err := c.UserService.UpdateUserInsurance(ins)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}
