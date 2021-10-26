package schedule

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"fmt"
	"testing"
	"time"
)

var dao = NewDao(database.GetMySqlEngine(), database.GetElasticSearchEngine())

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
	currentTime := time.Date(2021, 10, 26, 9, 50, 0, 0, time.UTC)
	nextTime := currentTime.Add(time.Hour*24*time.Duration(0))
	closedDateSettings := &doctor.ClosedDateSettings{
		AmStartTime: "09:00",
		AmEndTime: "11:00",
	}
	isOk, nextAvailableDate := dao.CalcNextAvailableDateForEachWeekDay(currentTime, constant.Virtual, nextTime, constant.InClinic, true, "09:00",
			"12:00", constant.Virtual, true,"01:00", "06:00", 15, 1, closedDateSettings)
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

func TestDao_GetDoctorInfoFromES(t *testing.T) {
	id := dao.GetDoctorInfoFromES(1619970365)
	if len(id) == 0 {
		t.Errorf("get fail")
	}
	fmt.Println(id)
}

func TestDao_SyncCertainDoctorNextAvailableDateToES(t *testing.T) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	err := dao.SyncCertainDoctorNextAvailableDateToES(1902809254, currentTime, currentTime)
	if err != nil {
		t.Errorf("sync failed")
	}
}

func TestDao_AddClosedDate(t *testing.T) {
	startDate := time.Date(2021, 10, 23, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, 10, 23, 23, 59, 0, 0, time.UTC)
	st := &doctor.ClosedDateSettings{
		Npi: 2222,
		StartDate: startDate,
		EndDate: endDate,
		AmStartTime: "09:00",
		AmEndTime: "12:00",
	}
	err := dao.AddClosedDate(st)
	if err != nil {
		t.Errorf("add closed date setting failed")
	}
}

func TestDao_DeleteClosedDate(t *testing.T) {
	npi := 2222
	err := dao.DeleteClosedDate(int64(npi))
	if err != nil {
		t.Errorf("delete closed date setting failed")
	}
}

func TestDao_ReverseMinutesToHourMin(t *testing.T) {
	time := dao.ReverseMinutesToHourMin(15)
	if time != "00:15" {
		t.Errorf("expected failed %s", time)
	}
}

func TestDao_CalcAvailableTimeByClosedDate(t *testing.T) {
	startTime := "09:00"
	endTime := "11:00"

	closedStartTime := "10:00"
	closedEndTime := "11:00"
	newStartTime, newEndTime := dao.CalcAvailableTimeByClosedDate(startTime, endTime, closedStartTime, closedEndTime)
	if newStartTime != "09:00" {
		t.Errorf("calc failed")
	}
	fmt.Println(newStartTime, newEndTime)
}