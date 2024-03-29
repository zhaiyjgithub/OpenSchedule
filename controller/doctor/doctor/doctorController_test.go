/**
 * @author zhaiyuanji
 * @date 2022年02月22日 2:29 下午
 */
package doctor

import (
	"OpenSchedule/model/doctorModel"
	"OpenSchedule/service/doctorService"
	"OpenSchedule/service/scheduleService"
	"fmt"
	"testing"
	"time"
)

var testNpi = int64(1902809254)
var c = &Controller{
	Ctx:             nil,
	ScheduleService: scheduleService.NewService(),
	DoctorService:   doctorService.NewService(),
}

func TestController_GetDoctorTimeSlotsPeerDay(t *testing.T) {
	booked := make([]doctorModel.TimeSlot, 0)
	booked = append(booked, doctorModel.TimeSlot{
		AvailableSlotsNumber: 3,
		Offset:               55,
	})
	setting := c.ScheduleService.GetScheduleSettings(testNpi)
	allTimeSlots := c.GetDoctorTimeSlotsPerDay(setting, time.Date(2022, 2, 23, 0, 45, 0, 0, time.UTC), booked)
	fmt.Println(allTimeSlots)
}

func TestController_ConvertBookedAppointmentsToTimeSlots(t *testing.T) {
	startDate := time.Date(2022, 2, 20, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC)
	npi := []int64{testNpi}
	bookedTimeSlots := c.ConvertBookedAppointmentsToTimeSlots(npi, startDate, endDate)
	fmt.Println(bookedTimeSlots)
}

func TestConvertClosedDateToTimeSlotsByDate(t *testing.T) {
	startDate := time.Date(2022, 4, 20, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 4, 20, 0, 0, 0, 0, time.UTC)
	targetDate := time.Date(2022, 4, 20, 0, 0, 0, 0, time.UTC)
	npi := []int64{testNpi}

	closedDateSettingsMap := c.MappingClosedDateByDateRange(npi, startDate, endDate)

	settingList, _ := closedDateSettingsMap[testNpi]
	scheduleSetting := c.ScheduleService.GetScheduleSettings(testNpi)
	timeSlots := c.GetDoctorTimeSlotsPerDay(scheduleSetting, targetDate, []doctorModel.TimeSlot{})
	closeDate := c.getClosedDateByDate(settingList, targetDate)
	filterList := c.FilterTimeSlotsByClosedDate(targetDate, timeSlots, closeDate)
	fmt.Println(filterList)
}
