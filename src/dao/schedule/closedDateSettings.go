package schedule

import (
	"OpenSchedule/src/model/doctor"
	"fmt"
	"time"
)

type ClosedDate struct {
	Npi int64
	ClosedDate time.Time
	AmStartTime string
	AmEndTime string
	PmStartTime string
	PmEndTime string
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
			ClosedDate: date,
			AmStartTime: d.GetHourMinuteFromTimestamp(st.AmStartDateTime),
			AmEndTime: d.GetHourMinuteFromTimestamp(st.AmEndDateTime),
			PmStartTime: d.GetHourMinuteFromTimestamp(st.PmStartDateTime),
			PmEndTime: d.GetHourMinuteFromTimestamp(st.PmEndTimeDateTime),
		}
		return cd, nil
	}
}

func (d *Dao) GetHourMinuteFromTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	hour := t.Hour()
	minute := t.Minute()
	var hourStr, minuteStr string
	if hour < 10 {
		hourStr = fmt.Sprintf("0%d", hour)
	}else {
		hourStr = fmt.Sprintf("0%d", hour)
	}

	if minute < 10 {
		minuteStr = fmt.Sprintf("0%d", minute)
	}else {
		minuteStr = fmt.Sprintf("0%d", minute)
	}
	return hourStr + "-" + minuteStr
}