package schedule

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/utils"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Dao struct {
	engine *gorm.DB
}

func NewDao(engine *gorm.DB) *Dao {
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

func (d *Dao) GetScheduleSettings(npi int64) *doctor.ScheduleSettings {
	st := &doctor.ScheduleSettings{}
	db := d.engine.Where("npi = ?", npi).First(st)
	if db.Error != nil {
		return nil
	}
	return st
}

func (d *Dao) SyncScheduleNextAvailableDateToES()  {

}

func (d *Dao) CalcNextAvailableDate(t time.Time, settings *doctor.ScheduleSettings) (bool, string)  {
	weekDay := t.Weekday()
	duration := settings.DurationPerSlot
	number := settings.NumberPerSlot
	isOk := false
	nextAvailableDate := ""
	if weekDay == time.Sunday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.SundayAmStartTime, settings.SundayAmEndTime,
			settings.SundayPmStartTime, settings.SundayPmEndTime,duration, number)
	}else if weekDay == time.Friday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.FridayAmStartTime, settings.FridayAmEndTime,
			settings.FridayPmStartTime, settings.FridayPmEndTime,duration, number)
	}
	return isOk, nextAvailableDate
}

func (d *Dao) CalcNextAvailableDateForEachWeekDay(t time.Time,
	weekDayAMStartTime string, weekDayAMEndTime string, weekDayPMStartTime string, weekDayPMEndTime string,
	durationOfSlot int, numberOfSlot int) (bool, string) {
	AMStartTime, err := d.ParseScheduleTimeToUTC(weekDayAMStartTime)
	AMEndTime, err := d.ParseScheduleTimeToUTC(weekDayAMEndTime)

	PMStartTime, err := d.ParseScheduleTimeToUTC(weekDayPMStartTime)
	PMEndTime, err := d.ParseScheduleTimeToUTC(weekDayPMEndTime)
	if err != nil {
		return false, ""
	}
	if t.Before(AMStartTime)  {
		return true, time.Now().UTC().Format(time.RFC3339)
	}else if t.After(AMStartTime) && t.Before(AMEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(t, AMStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	}else if t.After(PMStartTime) && t.Before(PMEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(t, PMStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	}
	return false, ""
}

func (d *Dao) CalcNextAvailableDateForTimeRange(t time.Time, startTime time.Time, durationOfSlot int) string {
	availableTimeRangeMinutes := t.Sub(startTime).Minutes()
	slotTimeNumber := int(availableTimeRangeMinutes) / durationOfSlot
	slotTime := slotTimeNumber*durationOfSlot
	nextAvailableDateTime := startTime.Add(time.Duration(slotTime)*time.Minute).Format(time.RFC3339)
	return nextAvailableDateTime
}

func (d *Dao)ParseScheduleTimeToUTC(scheduleTime string) (time.Time, error) {
	if !utils.CheckDateTime(scheduleTime) {
		return time.Now().UTC(), errors.New("param error")
	}

	t := time.Now().UTC()
	year := t.Year()
	month := t.Month()
	day := t.Day()

	dateTime := strings.Split(scheduleTime, ":")
	hour, _ := strconv.Atoi(dateTime[0])
	min, _ := strconv.Atoi(dateTime[1])
	utcTime := time.Date(year, month, day, hour, min, 0, 0, time.UTC)
	return utcTime, nil
}