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
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"time"
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
		StartDate interface{}
		EndDate interface{}
		Gender constant.Gender
		Specialty string
		City string
		Lat float64
		Lon float64
		Distance int
		Page int
		PageSize  int
		SortByType constant.SortByType
	}

	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	total, doctorInfoList := c.DoctorService.SearchDoctor(
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

	type DoctorDetailInfo struct {
		*viewModel.DoctorInfo
		TimeSlots []viewModel.TimeSlotPeerDay `json:"timeSlots"`
	}
	data := make([]DoctorDetailInfo, 0)
	startDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", p.StartDate))
	if err != nil {
		response.Fail(c.Ctx, response.Error, "param error: start date", nil)
		return
	}
	dayLength := 5//endDate.Day() - startDate.Day() + 1
	npiList := make([]int64, 0)
	for _, docInfo := range doctorInfoList {
		npiList = append(npiList, docInfo.Npi)
	}
	settingsList := c.ScheduleService.GetSettingsByNpiList(npiList)
	settingMap := make(map[int64]*doctor.ScheduleSettings)
	for _, setting := range settingsList {
		settingMap[setting.Npi] = setting
	}
	for _, docInfo := range doctorInfoList {
		setting, ok := settingMap[docInfo.Npi]
		if ok {
			timeSlots := c.GetDoctorTimeSlotsInRange(setting, startDate, dayLength)
			data = append(data, DoctorDetailInfo{
				DoctorInfo: docInfo,
				TimeSlots: timeSlots,
			})
		}
		
	}

	response.Success(c.Ctx, response.Successful, struct {
		Total int64 `json:"total"`
		Data []DoctorDetailInfo`json:"data"`
	}{
		Total: total,
		Data: data,
	})
}

func (c *Controller)GetDoctorTimeSlotsInRange(setting *doctor.ScheduleSettings, startDate time.Time, len int) []viewModel.TimeSlotPeerDay {
	endDate := startDate.AddDate(0,0, len - 1)
	bookedTimeSlotsMap := c.ConvertBookedAppointmentsToTimeSlots(setting.Npi, startDate, endDate)
	timeSlots := make([]viewModel.TimeSlotPeerDay, 0)
	if setting == nil {
		return timeSlots
	}
	for i := 0 ; i < len; i ++ {
		targetDate := startDate.AddDate(0,0, i)
		dateKey := fmt.Sprintf( "%d-%d-%d", targetDate.Year(), targetDate.Month(), targetDate.Day())
		bookedTimeSlots, ok := bookedTimeSlotsMap[dateKey]
		timeSlotsPeerDay := make([]doctor.TimeSlot, 0)
		if ok {
			timeSlotsPeerDay = c.GetDoctorTimeSlotsPeerDay(setting, targetDate, bookedTimeSlots)
		} else {
			timeSlotsPeerDay = c.GetDoctorTimeSlotsPeerDay(setting, targetDate, make([]doctor.TimeSlot, 0))
		}
		timeSlots = append(timeSlots, viewModel.TimeSlotPeerDay{Date: targetDate, TimeSlots: timeSlotsPeerDay})
	}
	return timeSlots
}

func (c *Controller) ConvertBookedAppointmentsToTimeSlots(npi int64, startDate time.Time, endTime time.Time) map[string][]doctor.TimeSlot {
	appts := c.ScheduleService.GetAppointmentByRange(npi, constant.Requested, startDate, endTime)
	bookedTimeSlotsMap := make(map[string][]doctor.TimeSlot)
	for _, appt := range appts {
		offset := appt.AppointmentDate.Hour() * 60 + appt.AppointmentDate.Minute()
		dateKey := fmt.Sprintf( "%d-%d-%d", appt.AppointmentDate.Year(), appt.AppointmentDate.Month(), appt.AppointmentDate.Day())
		bookedTimeSlotsMap[dateKey] = append(bookedTimeSlotsMap[dateKey], doctor.TimeSlot{Offset: offset, AvailableSlotsNumber: 1})
	}
	return bookedTimeSlotsMap
}

func (c *Controller) GetDoctorTimeSlotsPeerDay(setting *doctor.ScheduleSettings, targetDate time.Time, bookedTimeSlots []doctor.TimeSlot) []doctor.TimeSlot  {
	bookApptTimeSlotsMap := make(map[int]int)
	for _, bts := range bookedTimeSlots {
		bookApptTimeSlotsMap[bts.Offset] = bts.AvailableSlotsNumber
	}

	weekDay := targetDate.Weekday()
	amStartTimeOffset := 0
	amEndTimeOffset := 0
	pmStartTimeOffset := 0
	pmEndTimeOffset := 0
	if weekDay == time.Sunday {
		amStartTimeOffset = setting.SundayAmStartTimeOffset
		amEndTimeOffset = setting.SundayAmEndTimeOffset
		pmStartTimeOffset = setting.SundayPmStartTimeOffset
		pmEndTimeOffset = setting.SundayPmEndTimeOffset
	} else if weekDay == time.Monday {
		amStartTimeOffset = setting.MondayAmStartTimeOffset
		amEndTimeOffset = setting.MondayAmEndTimeOffset
		pmStartTimeOffset = setting.MondayPmStartTimeOffset
		pmEndTimeOffset = setting.MondayPmEndTimeOffset
	}  else if weekDay == time.Tuesday {
		amStartTimeOffset = setting.TuesdayAmStartTimeOffset
		amEndTimeOffset = setting.TuesdayAmEndTimeOffset
		pmStartTimeOffset = setting.TuesdayPmStartTimeOffset
		pmEndTimeOffset = setting.TuesdayPmEndTimeOffset
	}  else if weekDay == time.Wednesday {
		amStartTimeOffset = setting.WednesdayAmStartTimeOffset
		amEndTimeOffset = setting.WednesdayAmEndTimeOffset
		pmStartTimeOffset = setting.WednesdayPmStartTimeOffset
		pmEndTimeOffset = setting.WednesdayPmEndTimeOffset
	}  else if weekDay == time.Thursday {
		amStartTimeOffset = setting.ThursdayAmStartTimeOffset
		amEndTimeOffset = setting.ThursdayAmEndTimeOffset
		pmStartTimeOffset = setting.ThursdayPmStartTimeOffset
		pmEndTimeOffset = setting.ThursdayPmEndTimeOffset
	}  else if weekDay == time.Friday {
		amStartTimeOffset = setting.FridayAmStartTimeOffset
		amEndTimeOffset = setting.FridayAmEndTimeOffset
		pmStartTimeOffset = setting.FridayPmStartTimeOffset
		pmEndTimeOffset = setting.FridayPmEndTimeOffset
	}  else if weekDay == time.Saturday {
		amStartTimeOffset = setting.SaturdayAmStartTimeOffset
		amEndTimeOffset = setting.SaturdayAmEndTimeOffset
		pmStartTimeOffset = setting.SaturdayPmStartTimeOffset
		pmEndTimeOffset = setting.SaturdayPmEndTimeOffset
	}

	currentOffSet := 0
	if time.Now().UTC().Day() == targetDate.Day() {
		currentOffSet = targetDate.Hour() * 60 + targetDate.Minute()
	}

	timeSlots := make([]doctor.TimeSlot, 0)
	for i := amStartTimeOffset; i <= amEndTimeOffset + amStartTimeOffset; i += setting.DurationPerSlot {
		if i < currentOffSet {
			continue
		}
		timeSlot := doctor.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
		availableNumber := setting.NumberPerSlot
		if numberOfBooked >= timeSlot.AvailableSlotsNumber {
			availableNumber = 0
		} else {
			availableNumber = timeSlot.AvailableSlotsNumber - numberOfBooked
		}
		timeSlot.AvailableSlotsNumber = availableNumber
		timeSlots = append(timeSlots, timeSlot)
	}
	for i := pmStartTimeOffset + amStartTimeOffset; i <= pmEndTimeOffset + pmStartTimeOffset; i += setting.DurationPerSlot {
		if i < currentOffSet {
			continue
		}
		timeSlot := doctor.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
		availableNumber := setting.NumberPerSlot
		if numberOfBooked >= timeSlot.AvailableSlotsNumber {
			availableNumber = 0
		} else {
			availableNumber = timeSlot.AvailableSlotsNumber - numberOfBooked
		}
		timeSlot.AvailableSlotsNumber = availableNumber
		timeSlots = append(timeSlots, timeSlot)
	}
	return timeSlots
}

func getBookNumberOfTimeSlot(currentOffset int, duration int, bookedTimeSlots []doctor.TimeSlot) int {
	bookedNumber := 0
	for _, ts := range bookedTimeSlots {
		if ts.Offset <= currentOffset && ts.Offset > currentOffset - duration {
			bookedNumber = bookedNumber + 1
		}
	}
	return bookedNumber
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
