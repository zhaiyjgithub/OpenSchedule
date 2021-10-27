package schedule

import "OpenSchedule/src/model/doctor"

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

func (d *Dao) GetScheduleSettings(npi int64) *doctor.ScheduleSettings {
	st := &doctor.ScheduleSettings{}
	db := d.engine.Where("npi = ?", npi).First(st)
	if db.Error != nil {
		return nil
	}
	return st
}