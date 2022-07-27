package schedule

import (
	"OpenSchedule/constant"
	"OpenSchedule/model/doctor"
	"fmt"
	"time"
)

type ClosedDate struct {
	Npi         int64
	ClosedDate  time.Time
	AmStartTime string
	AmEndTime   string
	PmStartTime string
	PmEndTime   string
}

func (d *Dao) AddClosedDate(closeDateSettings doctor.ClosedDateSettings) error {
	db := d.engine.Create(&closeDateSettings)
	return db.Error
}

func (d *Dao) DeleteClosedDateByID(npi int64, id int) error {
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("id = ? && npi = ?", id, npi).Delete(st)
	return db.Error
}

func (d *Dao) GetClosedDate(npi int64) []doctor.ClosedDateSettings {
	var list []doctor.ClosedDateSettings
	_ = d.engine.Where("npi = ?", npi).Find(&list)
	return list
}

func (d *Dao) GetClosedDateByDateTime(npi int64, t time.Time) (*doctor.ClosedDateSettings, error) {
	ft := t.Format(constant.YYYY_MM_DD_HH_mm_SS)
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ? and start_date <= ? and end_date >= ?", npi, ft, ft).First(st)
	if db.Error != nil {
		return nil, db.Error
	}
	return st, nil
}

func (d *Dao) GetClosedDateByRange(npi []int64, from time.Time, to time.Time) []doctor.ClosedDateSettings {
	var closedDateSettings []doctor.ClosedDateSettings
	ft := from.Format(constant.YYYMMDD)
	tt := to.Format(constant.YYYMMDD)
	_ = d.engine.Raw("SELECT * FROM closed_date_settings WHERE npi IN (?) AND ( DATE_FORMAT(start_date, '%Y-%M-%D') BETWEEN DATE_FORMAT(?, '%Y-%M-%D') AND DATE_FORMAT(?, '%Y-%M-%D') OR DATE_FORMAT(end_date, '%Y-%M-%D') BETWEEN DATE_FORMAT(?, '%Y-%M-%D') AND DATE_FORMAT(?, '%Y-%M-%D'));", npi, ft, tt, ft, tt).Scan(&closedDateSettings)
	return closedDateSettings
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
	} else {
		hourStr = fmt.Sprintf("%d", hour)
	}

	if minute < 10 {
		minuteStr = fmt.Sprintf("0%d", minute)
	} else {
		minuteStr = fmt.Sprintf("%d", minute)
	}
	return hourStr + "-" + minuteStr
}
