package schedule

import (
	"OpenSchedule/src/model/doctor"
)

func (d *Dao) AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error {
	db := d.engine.Create(closeDateSettings)
	return db.Error
}

func (d *Dao) DeleteClosedDate(id int) error {
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("id = ?", id).Delete(st)
	return db.Error
}

func (d *Dao) GetClosedDate(npi int64) *doctor.ClosedDateSettings {
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ?", npi).First(st)
	if db.Error != nil {
		return nil
	}
	return st
}