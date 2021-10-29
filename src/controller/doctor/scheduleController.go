package doctor

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/response"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/schedule"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ScheduleController struct {
	Ctx iris.Context
	ScheduleService schedule.Service
}

func (c *ScheduleController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, router.SetScheduleSettings, "SetScheduleSettings")
	b.Handle(iris.MethodPost, router.GetScheduleSettings, "GetScheduleSettings")
	b.Handle(iris.MethodPost, router.AddClosedDateSettings, "AddClosedDateSettings")
	b.Handle(iris.MethodPost, router.DeleteClosedDateSettings, "DeleteClosedDateSettings")
}

func (c *ScheduleController) SetScheduleSettings()  {
	var p doctor.ScheduleSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	err := c.ScheduleService.SetScheduleSettings(&p)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(),nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *ScheduleController) GetScheduleSettings()  {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	settings := c.ScheduleService.GetScheduleSettings(p.Npi)
	response.Success(c.Ctx, response.Successful, settings)
}

func (c *ScheduleController) AddClosedDateSettings()  {
	var p doctor.ClosedDateSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.AddClosedDate(&p); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *ScheduleController) DeleteClosedDateSettings()  {
	type Param struct {
		ID int `json:"id"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.DeleteClosedDate(p.ID); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}
