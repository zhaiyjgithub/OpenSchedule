package doctor

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/model/viewModel"
	"OpenSchedule/src/response"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/doctorService"
	"OpenSchedule/src/service/scheduleService"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type Controller struct {
	Ctx           iris.Context
	DoctorService doctorService.Service
	ScheduleService scheduleService.Service
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(http.MethodPost, router.SearchDoctor, "SearchDoctor")
	b.Handle(http.MethodPost, router.GetDoctor, "GetDoctor")
	b.Handle(http.MethodPost, router.SaveDoctor, "SaveDoctor")
}

func (c *Controller) GetDoctor()  {
	type Param struct {
		Npi int64 `json:"npi"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	doc := c.DoctorService.GetDoctor(p.Npi)
	response.Success(c.Ctx, response.Successful, doc)
}

func (c *Controller) SearchDoctor()  {
	//keyword := ""
	//isInClinicEnable := true
	//isVirtualEnable:= true
	//appointmentType:= constant.InClinic
	//nextAvailableDate:= "2021-07-05T14:36:41Z"
	//city := ""
	//specialty := "Obstetrics & Gynecology"
	//lat := 40.747898
	//lon := -73.324025
	//gender := constant.Male
	//page := 1
	//pageSize := 200
	//sortType := constant.ByDistance
	//distance:= 200

	type Param struct {
		Keyword string
		AppointmentType constant.AppointmentType
		StartDate string
		EndDate string
		Gender constant.Gender
		Specialty string
		City string
		Lat float64
		Lon float64
		Distance int
		Page int
		PageSize   int
		SortByType constant.SortByType
	}

	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	total, docs := c.DoctorService.SearchDoctor(
		p.Keyword,
		p.AppointmentType,
		p.StartDate,
		p.EndDate,
		p.City,
		p.Specialty,
		p.Lat,
		p.Lon,
		p.Gender,
		p.Page,
		p.PageSize,
		p.SortByType,
		p.Distance)
	response.Success(c.Ctx, response.Successful, struct {
		Total int64 `json:"total"`
		Data []*viewModel.DoctorInfo `json:"data"`
	}{
		Total: total,
		Data: docs,
	})
}

func (c *Controller) SaveDoctor() {
	type Param struct {
		Doctor doctor.Doctor `json:"doctor"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.DoctorService.SaveDoctor(&p.Doctor); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
		return
	}
	if err := c.ScheduleService.SyncDoctorToES(&p.Doctor); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
		return
	}
	response.Success(c.Ctx, response.Successful, nil)
}
