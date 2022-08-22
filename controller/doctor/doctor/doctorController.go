package doctor

import (
	"OpenSchedule/constant"
	"OpenSchedule/model/doctor"
	"OpenSchedule/model/viewModel"
	"OpenSchedule/response"
	"OpenSchedule/router"
	"OpenSchedule/service/doctorService"
	"OpenSchedule/service/scheduleService"
	"OpenSchedule/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"time"
)

type Controller struct {
	Ctx             iris.Context
	DoctorService   doctorService.Service
	ScheduleService scheduleService.Service
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, router.SearchDoctor, "SearchDoctor")
	b.Handle(http.MethodPost, router.GetDoctor, "GetDoctor")
	b.Handle(http.MethodPost, router.SaveDoctor, "SaveDoctor")
	b.Handle(http.MethodPost, router.GetTimeSlots, "GetTimeSlots")
	b.Handle(http.MethodPost, router.GetDoctorDetailInfo, "GetDoctorDetailInfo")
}

func (c *Controller) GetDoctor() {
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

func (c *Controller) SearchDoctor() {
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
		Keyword         string
		AppointmentType constant.AppointmentType
		StartDate       string
		EndDate         string
		DateRange       int `validate:"required,gt=1"`
		Gender          constant.Gender
		Specialty       string
		City            string
		Lat             float64
		Lon             float64
		Distance        int
		Page            int
		PageSize        int
		SortByType      constant.SortByType
	}

	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	total, doctorInfoList := c.DoctorService.SearchDoctor(
		p.Keyword,
		p.AppointmentType,
		p.City,
		p.Specialty,
		p.Lat,
		p.Lon,
		p.Gender,
		p.Page,
		p.PageSize,
		p.SortByType,
		p.Distance)

	type DoctorDetailInfo struct {
		*viewModel.DoctorInfo
		TimeSlots []viewModel.TimeSlotPerDay `json:"timeSlotsPerDay"`
	}
	data := make([]DoctorDetailInfo, 0)
	startDate, _ := time.Parse(time.RFC3339, p.StartDate)
	endDate, _ := time.Parse(time.RFC3339, p.EndDate)

	startDateUTC := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDateUTC := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)

	npiList := make([]int64, 0)
	for _, docInfo := range doctorInfoList {
		npiList = append(npiList, docInfo.Npi)
	}
	settingMap := c.ScheduleService.GetScheduleSettingByNpiList(npiList)
	closedDateMap := c.ScheduleService.GetClosedDateByNpiList(npiList, startDateUTC, endDateUTC)
	bookedTimeSlots := c.ScheduleService.GetBookedAppointmentsTimeSlotsByNpiList(npiList, startDateUTC, endDateUTC)
	for _, docInfo := range doctorInfoList {
		setting, ok := settingMap[docInfo.Npi]
		if ok {
			bookedTimeSlotsForNpi, ok := bookedTimeSlots[setting.Npi]
			if !ok {
				bookedTimeSlotsForNpi = make(map[string][]doctor.TimeSlot)
			}
			closeDateForNpi, ok := closedDateMap[setting.Npi]
			if !ok {
				closeDateForNpi = []doctor.ClosedDateSettings{}
			}
			timeSlots := c.ScheduleService.GetDoctorTimeSlotsByDate(setting, startDateUTC, endDateUTC, bookedTimeSlotsForNpi, closeDateForNpi)
			data = append(data, DoctorDetailInfo{
				DoctorInfo: docInfo,
				TimeSlots:  timeSlots,
			})
		}
	}

	response.Success(c.Ctx, response.Successful, struct {
		Total int64              `json:"total"`
		Data  []DoctorDetailInfo `json:"data"`
	}{
		Total: total,
		Data:  data,
	})
}

func (c *Controller) getScheduleSettingByNpi(npi []int64) map[int64]doctor.ScheduleSettings {
	list := c.ScheduleService.GetSettingsByNpiList(npi)
	settingMap := make(map[int64]doctor.ScheduleSettings)
	for _, setting := range list {
		settingMap[setting.Npi] = setting
	}
	return settingMap
}

func (c *Controller) GetTimeSlots() {
	type Param struct {
		Npi       int64  `json:"npi"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	startDate, err := time.Parse(time.RFC3339, p.StartDate)
	endDate, err := time.Parse(time.RFC3339, p.EndDate)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "param error: start date", nil)
		return
	}
	startDateUTC := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDateUTC := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)

	setting := c.ScheduleService.GetScheduleSettings(p.Npi)
	if setting.Npi == 0 {
		response.Fail(c.Ctx, response.Error, response.NotFound, nil)
		return
	}

	npiList := []int64{setting.Npi}
	allBookedTimeSlotsMap := c.ScheduleService.GetBookedAppointmentsTimeSlotsByNpiList(npiList, startDateUTC, endDateUTC)
	closedDateMap := c.ScheduleService.GetClosedDateByNpiList(npiList, startDateUTC, endDateUTC)
	bookedTimeSlotsForNpi, ok := allBookedTimeSlotsMap[setting.Npi]
	if !ok {
		bookedTimeSlotsForNpi = make(map[string][]doctor.TimeSlot)
	}
	closeDateForNpi, ok := closedDateMap[setting.Npi]
	if !ok {
		closeDateForNpi = []doctor.ClosedDateSettings{}
	}
	timeSlots := c.ScheduleService.GetDoctorTimeSlotsByDate(setting, startDateUTC, endDateUTC, bookedTimeSlotsForNpi, closeDateForNpi)
	response.Success(c.Ctx, response.Successful, timeSlots)
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

func (c *Controller) GetDoctorDetailInfo() {
	type Param struct {
		Npi int64
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	doc := c.DoctorService.GetDoctorDetail(p.Npi)
	doc.Awards = c.DoctorService.GetAwards(p.Npi)
	doc.Certifications = c.DoctorService.GetCertification(p.Npi)
	doc.Educations = c.DoctorService.GetEducation(p.Npi)
	doc.Insurances = c.DoctorService.GetInsurance(p.Npi)

	response.Success(c.Ctx, response.Successful, doc)
}
