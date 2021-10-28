package schedule

import (
	"OpenSchedule/src/constant"
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
	startDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	fStartDate := startDate.Format(constant.YYYY_MM_DD_HH_mm_SS)
	fEndDate := startDate.Add(time.Hour*24).Format(constant.YYYY_MM_DD_HH_mm_SS)
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ? and (closed_date >= UNIX_TIMESTAMP(?) or closed_date < UNIX_TIMESTAMP(?))", npi, fStartDate, fEndDate).First(st)
	if db.Error != nil {
		return nil, db.Error
	}else {
		cd := &ClosedDate{
			Npi: st.Npi,
			ClosedDate: st.ClosedDate,
			AmStartTime: d.GetHourMinuteFromTimestamp(st.AmStartDateTime, true),
			AmEndTime: d.GetHourMinuteFromTimestamp(st.AmEndDateTime, true),
			PmStartTime: d.GetHourMinuteFromTimestamp(st.PmStartDateTime, false),
			PmEndTime: d.GetHourMinuteFromTimestamp(st.PmEndDateTime, false),
		}
		return cd, nil
	}
}

func (d *Dao) GetHourMinuteFromTimestamp(t time.Time, isAM bool) string {
	if t.Equal(constant.DefaultTimeStamp) {
		return ""
	}
	hour := t.Hour()
	if !isAM {
		hour = hour - 12
	}
	minute := t.Minute()
	var hourStr, minuteStr string
	if hour < 10 {
		hourStr = fmt.Sprintf("0%d", hour)
	}else {
		hourStr = fmt.Sprintf("%d", hour)
	}

	if minute < 10 {
		minuteStr = fmt.Sprintf("0%d", minute)
	}else {
		minuteStr = fmt.Sprintf("%d", minute)
	}
	return hourStr + "-" + minuteStr
}