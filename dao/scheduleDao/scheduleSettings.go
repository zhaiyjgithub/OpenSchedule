package scheduleDao

import "OpenSchedule/model/doctorModel"

func (d *Dao) SetScheduleSettings(setting doctorModel.ScheduleSettings) error {
	var st doctorModel.ScheduleSettings
	db := d.engine.Where("npi = ?", setting.Npi).First(st)
	if db.Error != nil { // not found
		db = d.engine.Create(setting)
		return db.Error
	} else {
		setting.ID = st.ID
		db = d.engine.Model(&doctorModel.ScheduleSettings{}).Where("id = ?", st.ID).
			Select("*").Omit("created_at").Updates(setting)
		return db.Error
	}
}

func (d *Dao) GetScheduleSettings(npi int64) doctorModel.ScheduleSettings {
	var setting doctorModel.ScheduleSettings
	_ = d.engine.Where("npi = ?", npi).First(&setting)
	return setting
}

func (d *Dao) IsExist(npi int64) bool {
	var count int64
	d.engine.Model(&doctorModel.ScheduleSettings{}).Where("npi = ?", npi).Count(&count)
	return count > 0
}
