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

type WeekSchedule struct {
	AmIsEnable bool
	AmStartTime string
	AmEndTime string

	PmIsEnable bool
	PmStartTime string
	PmEndTime string

	DurationPerSlot	int
	NumberPerSlot	int
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

/*
  Update the next available date when update the settings for certain doctor.

*/
func (d *Dao) SyncCertainDoctorScheduleNextAvailableDateToES()  {

}

func (d *Dao) CalcNextAvailableDate(t time.Time, settings *doctor.ScheduleSettings) (bool, string)  {
	duration := settings.DurationPerSlot
	number := settings.NumberPerSlot
	isOk := false
	nextAvailableDate := ""
	for i := 0; i < 14; i ++ {// future 2 weeks
		nextTime := t.Add(time.Hour*24*time.Duration(i))
		weekDay := nextTime.Weekday()
		if weekDay == time.Sunday && (settings.SundayAmIsEnable || settings.SundayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.SundayAmIsEnable, settings.SundayAmStartTime, settings.SundayAmEndTime,
				settings.SundayPmIsEnable, settings.SundayPmStartTime, settings.SundayPmEndTime,duration, number)
		}else if weekDay == time.Monday && (settings.MondayAmIsEnable || settings.MondayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.MondayAmIsEnable, settings.MondayAmStartTime, settings.MondayAmEndTime,
				settings.MondayPmIsEnable, settings.MondayPmStartTime, settings.MondayPmEndTime,duration, number)
		}else if weekDay == time.Tuesday && (settings.TuesdayAmIsEnable || settings.TuesdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.TuesdayAmIsEnable, settings.TuesdayAmStartTime, settings.TuesdayAmEndTime,
				settings.TuesdayPmIsEnable, settings.TuesdayPmStartTime, settings.TuesdayPmEndTime,duration, number)
		}else if weekDay == time.Wednesday && (settings.WednesdayAmIsEnable || settings.WednesdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.WednesdayAmIsEnable, settings.WednesdayAmStartTime, settings.WednesdayAmEndTime,
				settings.WednesdayPmIsEnable, settings.WednesdayPmStartTime, settings.WednesdayPmEndTime,duration, number)
		}else if weekDay == time.Thursday && (settings.ThursdayAmIsEnable || settings.ThursdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.ThursdayAmIsEnable, settings.ThursdayAmStartTime, settings.ThursdayAmEndTime,
				settings.ThursdayPmIsEnable, settings.ThursdayPmStartTime, settings.ThursdayPmEndTime,duration, number)
		}else if weekDay == time.Friday && (settings.FridayAmIsEnable || settings.FridayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.FridayAmIsEnable, settings.FridayAmStartTime, settings.FridayAmEndTime,
				settings.FridayPmIsEnable, settings.FridayPmStartTime, settings.FridayPmEndTime,duration, number)
		}else if weekDay == time.Saturday && (settings.SaturdayAmIsEnable || settings.SaturdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, nextTime, settings.SaturdayAmIsEnable, settings.SaturdayAmStartTime, settings.SaturdayAmEndTime,
				settings.SaturdayPmIsEnable, settings.SaturdayPmStartTime, settings.SaturdayPmEndTime,duration, number)
		}
		if isOk {
			break
		}
	}
	
	return isOk, nextAvailableDate
}

func (d *Dao) ContainTimeInRange(t time.Time, startTime string, endTime string, isAmEnable bool, isPmEnable bool) bool  {
	startTimeUTC, _ := d.ParseScheduleTimeToUTC(t, startTime, isAmEnable)
	endTimeUTC, _ := d.ParseScheduleTimeToUTC(t, endTime, !isPmEnable)
	return t.After(startTimeUTC) && t.Before(endTimeUTC)
}

func (d *Dao) CalcNextAvailableDateForEachWeekDay(currentTime time.Time, nextTime time.Time,
	isAmEnable bool, weekDayAMStartTime string, weekDayAMEndTime string,
	isPmEnable bool, weekDayPMStartTime string, weekDayPMEndTime string,
	durationOfSlot int, numberOfSlot int) (bool, string) {
	amStartTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayAMStartTime, true)
	amEndTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayAMEndTime, true)

	pmStartTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayPMStartTime, false)
	pmEndTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayPMEndTime, false)
	if err != nil {
		return false, ""
	}
	if isAmEnable && currentTime.Before(amStartTime)  {
		return true, amStartTime.Format(time.RFC3339)
	}else if isAmEnable && currentTime.After(amStartTime) && currentTime.Before(amEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(currentTime, amStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	}else if isPmEnable && currentTime.Before(pmStartTime)  {
		return true, pmStartTime.Format(time.RFC3339)
	} else if isPmEnable && currentTime.After(amEndTime) && currentTime.Before(pmEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(currentTime, pmStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	}
	return false, ""
}

func (d *Dao) CalcNextAvailableDateForTimeRange(now time.Time, startTime time.Time, durationOfSlot int) string {
	if now.Before(startTime) {
		return startTime.Format(time.RFC3339)
	}
	availableTimeRangeMinutes := now.Sub(startTime).Minutes()
	slotTimeNumber := int(availableTimeRangeMinutes) / durationOfSlot
	if int(availableTimeRangeMinutes) % durationOfSlot > 0 {
		slotTimeNumber += 1
	}
	slotTime := slotTimeNumber*durationOfSlot
	nextAvailableDateTime := startTime.Add(time.Duration(slotTime)*time.Minute).Format(time.RFC3339)
	return nextAvailableDateTime
}

func (d *Dao)ParseScheduleTimeToUTC(t time.Time, scheduleTime string, isAM bool) (time.Time, error) {
	if !utils.CheckDateTime(scheduleTime) {
		return time.Now().UTC(), errors.New("param error")
	}
	year := t.Year()
	month := t.Month()
	day := t.Day()

	dateTime := strings.Split(scheduleTime, ":")
	hour, _ := strconv.Atoi(dateTime[0])
	min, _ := strconv.Atoi(dateTime[1])
	if !isAM {
		hour += 12
	}
	utcTime := time.Date(year, month, day, hour, min, 0, 0, time.UTC)
	return utcTime, nil
}
