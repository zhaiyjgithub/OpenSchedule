package schedule

import (
	"OpenSchedule/model/doctor"
	"OpenSchedule/response"
	"OpenSchedule/router"
	"OpenSchedule/service/scheduleService"
	"OpenSchedule/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Controller struct {
	Ctx             iris.Context
	ScheduleService scheduleService.Service
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, router.SetScheduleSettings, "SetScheduleSettings")
	b.Handle(iris.MethodPost, router.GetScheduleSettings, "GetScheduleSettings")
	b.Handle(iris.MethodPost, router.AddClosedDateSettings, "AddClosedDateSettings")
	b.Handle(iris.MethodPost, router.DeleteClosedDateSettings, "DeleteClosedDateSettings")
	b.Handle(iris.MethodPost, router.GetClosedDateSettings, "GetClosedDateSettings")
}

func (c *Controller) SetScheduleSettings() {
	var p doctor.ScheduleSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	err := c.ScheduleService.SetScheduleSettings(p)
	if err != nil {
		err = c.ScheduleService.SyncCertainDoctorScheduleNextAvailableDateToES(p)
	}
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetScheduleSettings() {
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

func (c *Controller) AddClosedDateSettings() {
	var p doctor.ClosedDateSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.AddClosedDate(p); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) DeleteClosedDateSettings() {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
		Sid int   `json:"sid"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.DeleteClosedDate(p.Npi, p.Sid); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetClosedDateSettings() {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	list := c.ScheduleService.GetClosedDate(p.Npi)
	response.Success(c.Ctx, response.Successful, list)
}
