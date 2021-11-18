package schedule

import (
	"OpenSchedule/src/constant"
	dao2 "OpenSchedule/src/dao/doctor"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"fmt"
	"testing"
	"time"
)

var dao = NewDao(database.GetMySqlEngine(), database.GetElasticSearchEngine())
var doctorDao = dao2.NewDoctorDao(database.GetElasticSearchEngine(), database.GetMySqlEngine())

var testNpi = int64(1902809254)

func TestDao_SetScheduleSettings(t *testing.T) {
	//settings := new(doctor.ScheduleSettings)
	//settings.Npi = testNpi
	//settings.NumberPerSlot = 1
	//settings.DurationPerSlot = 15
	//settings.ThursdayAmIsEnable = true
	//settings.ThursdayAmAppointmentType = constant.InClinic
	//settings.ThursdayAmStartTime = "08:00"
	//settings.ThursdayAmEndTimeOffset = 240
	//
	//settings.ThursdayPmIsEnable = true
	//settings.ThursdayPmAppointmentType = constant.InClinic
	//settings.ThursdayPmStartTime = "13:00"
	//settings.ThursdayPmEndTimeOffset = 240
	//err := dao.SetScheduleSettings(settings)
	//if err != nil {
	//	t.Errorf("test failed")
	//}
}

func TestDao_ParseScheduleTimeToUTC(t *testing.T) {
	currentTime := time.Date(2021, 10, 8, 9, 42, 0, 0, time.UTC)
	stime, err := dao.ParseScheduleTimeToUTC(currentTime,"09:45", false)
	if err != nil {
		t.Errorf("test failed")
	}
	fmt.Println(stime.Format(time.RFC3339))
}

func TestDao_MatchDateTimeByDuration(t *testing.T) {
	startTime := time.Date(2021, 10, 8, 9, 0, 0, 0, time.UTC)
	currentTime := time.Date(2021, 10, 8, 9, 42, 0, 0, time.UTC)
	nextAvailableDate := dao.MatchDateTimeByDuration(currentTime, startTime, 15)
	if nextAvailableDate == constant.InvalidDateTime {
		t.Errorf("test failed")
	}
	fmt.Println(nextAvailableDate)
}

func TestDao_CalcNextAvailableDateForEachWeekDay(t *testing.T) {
	currentTime := time.Date(2021, 10, 28, 11, 8, 0, 0, time.UTC)
	nextTime := currentTime.Add(time.Hour*24*time.Duration(0))
	closedDate, _ := dao.GetClosedDateByDateTime(testNpi, nextTime)

	amStartDateTime := time.Date(2021, 10, 28, 8, 0, 0, 0, time.UTC)
	amEndDateTime := time.Date(2021, 10, 28, 12, 0, 0, 0, time.UTC)

	pmStartDateTime := constant.DefaultTimeStamp//time.Date(2021, 10, 29, 9, 0, 0, 0, time.UTC)
	pmEndDateTime := constant.DefaultTimeStamp//time.Date(2021, 10, 29, 12, 0, 0, 0, time.UTC)

	nextAvailableDate := dao.CalcNextAvailableDateForEachWeekDay(currentTime, constant.InClinic, constant.InClinic, true, amStartDateTime,
			amEndDateTime, constant.InClinic, true, pmStartDateTime, pmEndDateTime, 15, 1, closedDate)
	if nextAvailableDate == constant.InvalidDateTime {
		t.Errorf("CalcNextAvailableDateForEachWeekDay failed")
	}
	fmt.Println(nextAvailableDate)
}

func TestDao_CalcNextAvailableDate(t *testing.T) {
	currentTime := time.Now().UTC()//time.Date(2021, 10, 28, 11, 36, 0, 0, time.UTC)
	st := dao.GetScheduleSettings(testNpi)
	nextAvailableDate := dao.CalcNextAvailableDate(currentTime, constant.Virtual, st)
	if nextAvailableDate == constant.InvalidDateTime {
		t.Errorf("calc failed")
	}
	fmt.Println(nextAvailableDate)
}

func TestDao_GetDoctorInfoFromES(t *testing.T) {
	id := dao.GetDoctorInfoFromES(testNpi)
	if len(id) == 0 {
		t.Errorf("get fail")
	}
	fmt.Println(id)
}

func TestDao_SyncCertainDoctorNextAvailableDateToES(t *testing.T) {
	currentTime := time.Now().UTC()
	settings := dao.GetScheduleSettings(testNpi)

	nextAvailableDateInClinic := dao.CalcNextAvailableDate(currentTime, constant.InClinic, settings)
	nextAvailableDateVirtual := dao.CalcNextAvailableDate(currentTime, constant.Virtual, settings)

	isInClinicBookEnable := nextAvailableDateInClinic != constant.InvalidDateTime
	isVirtualBookEnable := nextAvailableDateVirtual != constant.InvalidDateTime
	isOnlineScheduleEnable := isInClinicBookEnable || isVirtualBookEnable

	err := dao.SyncCertainDoctorNextAvailableDateToES(testNpi,
		isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
		nextAvailableDateInClinic, nextAvailableDateVirtual)
	if err != nil {
		t.Errorf("sync failed")
	}
}

func TestDao_AddClosedDate(t *testing.T) {
	closeStartDate := time.Date(2021, 10, 28, 0, 0, 0, 0, time.UTC)
	closeEndDate := time.Date(2021, 10, 29, 0, 0, 0, 0, time.UTC)
	amStartDateTime := time.Date(2021, 10, 28, 10, 0, 0, 0, time.UTC)
	amEndDateTime := time.Date(2021, 10, 28, 11, 0, 0, 0, time.UTC)
	st := &doctor.ClosedDateSettings{
		Npi: testNpi,
		StartDate: closeStartDate,
		EndDate: closeEndDate,
		AmStartDateTime: amStartDateTime,
		AmEndDateTime: amEndDateTime,
		PmStartDateTime: constant.DefaultTimeStamp,
		PmEndDateTime: constant.DefaultTimeStamp,
	}
	err := dao.AddClosedDate(st)
	if err != nil {
		t.Errorf("add closed date setting failed")
	}
}

func TestDao_DeleteClosedDate(t *testing.T) {
	id := 1
	err := dao.DeleteClosedDateByID(testNpi, id)
	if err != nil {
		t.Errorf("delete closed date setting failed")
	}
}

func TestDao_CalcAvailableTimeByClosedDate(t *testing.T) {
	startDateTime := time.Date(2021, 10, 29, 9, 0, 0, 0, time.UTC)
	endDateTime := time.Date(2021, 10, 29, 12, 0, 0, 0, time.UTC)

	closedStartDateTime := time.Date(2021, 10, 28, 10, 0, 0, 0, time.UTC)
	closedEndDateTime := time.Date(2021, 10, 28, 11, 0, 0, 0, time.UTC)

	currentTime := time.Date(2021, 10, 28, 11, 8, 0, 0, time.UTC)

	newStartTime, newEndTime := dao.CalcAvailableTimeRangeByClosedDate(currentTime, startDateTime, endDateTime, closedStartDateTime, closedEndDateTime)
	expected := time.Date(2021, 10, 29, 9, 0, 0, 0, time.UTC)
	if newStartTime == nil || !newStartTime.Equal(expected) {
		t.Errorf("TestDao_CalcAvailableTimeByClosedDate failed")
	}
	fmt.Println(newStartTime, newEndTime)

}

func TestDao_GetClosedDateByDateTime(t *testing.T) {
	cd := time.Date(2021, 10, 28, 9, 0, 0, 0, time.UTC)
	d, _ := dao.GetClosedDateByDateTime(testNpi, cd)
	fmt.Printf("%v", d)
}

func TestDao_GetClosedDate(t *testing.T) {
	l := dao.GetClosedDate(454344)
	fmt.Println(l)

	var a []doctor.ClosedDateSettings
	println(a == nil)

	c := make([]doctor.ClosedDateSettings, 0)
	println(c == nil)
}

func TestDao_GetDuplicateDoctorInfoFromES(t *testing.T) {
	//clear the database

	page := 1
	pageSize := 500
	count := 0
	for ;; {
		doctors := doctorDao.GetDoctorByPage(page, pageSize)
		for _, doc := range doctors {
			ids := dao.GetDuplicateDoctorInfoFromES(doc.Npi)
			if len(ids) == 2 {
				count = count + 1
				fmt.Println(doc.Npi, ids)
				//err := dao.DeleteESDoctorById(ids[1])
				//if err != nil {
				//	fmt.Println(err.Error())
				//}
			}
		}
		time.Sleep(time.Millisecond * 500)
	}

}