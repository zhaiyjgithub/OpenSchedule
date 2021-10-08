package schedule

import (
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"fmt"
	"testing"
	"time"
)

func TestDao_SetScheduleSettings(t *testing.T) {
	settings := new(doctor.ScheduleSettings)
	settings.Npi = 5
	settings.NumberPerSlot = 5

	dao := NewDao(database.GetMySqlEngine())
	err := dao.SetScheduleSettings(settings)
	if err != nil {
		t.Errorf("test failed")
	}
}

func TestDao_ParseScheduleTimeToUTC(t *testing.T) {
	dao := NewDao(database.GetMySqlEngine())
	stime, err := dao.ParseScheduleTimeToUTC("09:45")
	if err != nil {
		t.Errorf("test failed")
	}
	fmt.Println(stime.Format(time.RFC3339))
}

