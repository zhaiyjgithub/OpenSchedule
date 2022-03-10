/**
 * @author zhaiyuanji
 * @date 2022年02月22日 2:29 下午
 */
package doctor

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/service/doctorService"
	"OpenSchedule/src/service/scheduleService"
	"fmt"
	"testing"
	"time"
)

var testNpi = int64(1902809254)
var c = &Controller{
	Ctx: nil,
	ScheduleService: scheduleService.NewService(),
	DoctorService: doctorService.NewService(),
	}

func TestController_GetDoctorTimeSlotsPeerDay(t *testing.T) {
	booked := make([]doctor.TimeSlot, 0)
	booked = append(booked, doctor.TimeSlot{
		AvailableSlotsNumber: 3,
		Offset:               55,
	})
	setting := c.ScheduleService.GetScheduleSettings(testNpi)
	allTimeSlots := c.GetDoctorTimeSlotsPeerDay(setting, time.Date(2022, 2, 23, 0, 45, 0, 0, time.UTC), booked)
	fmt.Println(allTimeSlots)
}

func TestController_ConvertBookedAppointmentsToTimeSlots(t *testing.T) {
	startDate := time.Date(2022, 2, 20, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC)
	bookedTimeSlots := c.ConvertBookedAppointmentsToTimeSlots(testNpi, startDate, endDate)
	fmt.Println(bookedTimeSlots)
}


