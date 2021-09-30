package doctor

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/response"
	"OpenSchedule/src/router"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ScheduleController struct {
	Ctx iris.Context
	
}

func (c *ScheduleController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, router.SetScheduleSettings, "SetScheduleSettings")
}

func (c *ScheduleController) SetScheduleSettings()  {
	var p doctor.ScheduleSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}

	response.Success(c.Ctx, response.Successful, nil)
}
