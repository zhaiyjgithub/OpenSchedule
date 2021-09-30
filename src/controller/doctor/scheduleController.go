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
