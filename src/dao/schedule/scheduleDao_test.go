package schedule

import (
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"testing"
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
