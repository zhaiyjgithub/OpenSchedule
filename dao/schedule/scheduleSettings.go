package schedule

import "OpenSchedule/model/doctor"

func (d *Dao) SetScheduleSettings(setting doctor.ScheduleSettings) error {
	var st doctor.ScheduleSettings
	db := d.engine.Where("npi = ?", setting.Npi).First(st)
	if db.Error != nil { // not found
		db = d.engine.Create(setting)
		return db.Error
	} else {
		setting.ID = st.ID
		db = d.engine.Model(&doctor.ScheduleSettings{}).Where("id = ?", st.ID).
			Select("*").Omit("created_at").Updates(setting)
		return db.Error
	}
}

func (d *Dao) GetScheduleSettings(npi int64) doctor.ScheduleSettings {
	var setting doctor.ScheduleSettings
	_ = d.engine.Where("npi = ?", npi).First(&setting)
	return setting
}

func (d *Dao) IsExist(npi int64) bool {
	var count int64
	d.engine.Model(&doctor.ScheduleSettings{}).Where("npi = ?", npi).Count(&count)
	return count > 0
}
