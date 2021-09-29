package doctor

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ScheduleController struct {
	Ctx iris.Context
	
}

func (c *ScheduleController) BeforeActivation(b mvc.BeforeActivation)  {
	
}

func (c *ScheduleController) SetScheduleSettings()  {
	var p doctor.ScheduleSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
}
