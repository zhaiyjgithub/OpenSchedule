package schedule

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"fmt"
	"testing"
	"time"
)

var dao = NewDao(database.GetMySqlEngine())

func TestDao_SetScheduleSettings(t *testing.T) {
	settings := new(doctor.ScheduleSettings)
	settings.Npi = 5
	settings.NumberPerSlot = 5
	err := dao.SetScheduleSettings(settings)
	if err != nil {
		t.Errorf("test failed")
	}
}

func TestDao_ParseScheduleTimeToUTC(t *testing.T) {
	currentTime := time.Date(2021, 10, 8, 9, 42, 0, 0, time.UTC)
	stime, err := dao.ParseScheduleTimeToUTC(currentTime,"09:45", false)
	if err != nil {
		t.Errorf("test failed")
	}
	fmt.Println(stime.Format(time.RFC3339))
}

func TestDao_CalcNextAvailableDateForTimeRange(t *testing.T) {
	startTime := time.Date(2021, 10, 8, 9, 0, 0, 0, time.UTC)
	currentTime := time.Date(2021, 10, 8, 9, 42, 0, 0, time.UTC)
	nextAvailableDate := dao.CalcNextAvailableDateForTimeRange(currentTime, startTime, 15)
	if len(nextAvailableDate) == 0 {
		t.Errorf("test failed")
	}
	fmt.Println(nextAvailableDate)
}

func TestDao_CalcNextAvailableDateForEachWeekDay(t *testing.T) {
	currentTime := time.Date(2021, 10, 11, 9, 50, 0, 0, time.UTC)
	nextTime := currentTime.Add(time.Hour*24*time.Duration(0))
	isOk, nextAvailableDate := dao.CalcNextAvailableDateForEachWeekDay(currentTime, constant.Virtual, nextTime, constant.InClinic, true, "09:00",
			"12:00", constant.Virtual, true,"01:00", "06:00", 15, 1)
	if isOk != true {
		t.Errorf("calc failed")
	}
	fmt.Println(nextAvailableDate)
}

func TestDao_CalcNextAvailableDate(t *testing.T) {
	currentTime := time.Date(2021, 10, 11, 9, 36, 0, 0, time.UTC)
	st := dao.GetScheduleSettings(3)
	st.MondayAmIsEnable = true
	st.MondayPmIsEnable = true
	st.MondayAmAppointmentType = constant.InClinic
	st.MondayPmAppointmentType = constant.Virtual
	isOk, nextAvailableDate := dao.CalcNextAvailableDate(currentTime, constant.Virtual, st)
	if isOk != true {
		t.Errorf("calc failed")
	}
	fmt.Println(nextAvailableDate)
}