package doctor

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/response"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/doctorService"
	"OpenSchedule/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type FindDoctorController struct {
	Ctx           iris.Context
	DoctorService doctorService.Service
}

func (c *FindDoctorController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(http.MethodPost, router.SearchDoctor, "SearchDoctor")
	b.Handle(http.MethodPost, router.GetDoctor, "GetDoctor")
}

func (c *FindDoctorController) GetDoctor()  {
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

func (c *FindDoctorController) SearchDoctor()  {
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
		IsInClinicEnable bool
		IsVirtualEnable bool
		AppointmentType constant.AppointmentType
		NextAvailableDate string
		Gender constant.Gender
		Specialty string
		City string
		Lat float64
		Lon float64
		Distance int
		Page int
		PageSize int
		SortType constant.SortType
	}

	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}

	docs := c.DoctorService.SearchDoctor(p.Keyword,
		p.IsInClinicEnable,
		p.IsVirtualEnable,
		p.AppointmentType,
		p.NextAvailableDate,
		p.City,
		p.Specialty,
		p.Lat,
		p.Lon,
		p.Gender,
		p.Page,
		p.PageSize,
		p.SortType,
		p.Distance)
	response.Success(c.Ctx, response.Successful, docs)
}
