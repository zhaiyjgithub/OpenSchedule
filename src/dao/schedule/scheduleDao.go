package schedule

import (
	"OpenSchedule/src/model/doctor"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func NewScheduleDao(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) SetScheduleSettings(setting *doctor.ScheduleSettings) error {
	st := &doctor.ScheduleSettings{}
	db := d.engine.Where("npi = ?", setting.Npi).First(st)
	if db.Error != nil { // not found
		db = d.engine.Create(setting)
		return db.Error
	}else {
		setting.ID = st.ID
		db = d.engine.Model(&doctor.ScheduleSettings{}).Where("id = ?", st.ID).
			Select("*").Omit("created_at").Updates(setting)
		return db.Error
	}
}
