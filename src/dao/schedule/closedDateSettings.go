package schedule

import (
	"OpenSchedule/src/model/doctor"
	"time"
)

type ClosedDate struct {
	Npi int64
	AmStartDateTime time.Time
	AmEndDateTime time.Time
	PmStartDateTime time.Time
	PmEndDateTime time.Time
}

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

func (d *Dao) GetClosedDateByDateTime(npi int64, currentTime time.Time) (*ClosedDate, error) {
	date := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	cn := date.UnixNano()
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ? and closed_date = ?", npi, cn).First(st)
	if db.Error != nil {
		return nil, db.Error
	}else {
		cd := &ClosedDate{
			Npi: st.Npi,
			AmStartDateTime: time.Unix(st.AmStartDateTime, 0),
			AmEndDateTime: time.Unix(st.AmEndDateTime, 0),
			PmStartDateTime: time.Unix(st.PmStartDateTime, 0),
			PmEndDateTime: time.Unix(st.PmEndTimeDateTime, 0),
		}
		return cd, nil
	}
}