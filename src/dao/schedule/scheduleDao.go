package schedule

import (
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/utils"
	"errors"
	"fmt"
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

func (d *Dao) SyncScheduleNextAvailableDateToES()  {

}

func (d *Dao) CalcNextAvailableDate(t time.Time, settings *doctor.ScheduleSettings) (bool, string)  {
	weekDay := t.Weekday()
	duration := settings.DurationPerSlot
	number := settings.NumberPerSlot
	isOk := false
	nextAvailableDate := ""
	weekDayLine := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday,
		time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday,
		}
	weekSchedules := d.GetWeekSchedule(settings)
	i := int(weekDay)
	length := int(weekDay) + 7
	for ; i < length; i ++ {
		wl := weekDayLine[i]
		ws := weekSchedules[wl]
		if ws.AmIsEnable || ws.PmIsEnable {
			startTime := ""
			endTime := ""
			if ws.AmIsEnable && ws.PmIsEnable {
				startTime = ws.AmStartTime
				endTime = ws.PmEndTime
			}else if !ws.AmIsEnable && ws.PmIsEnable {
				startTime = ws.PmStartTime
				endTime = ws.PmEndTime
			}else {
				startTime = ws.AmStartTime
				endTime = ws.AmEndTime
			}
			isContain := d.ContainTimeInRange(t, startTime, endTime, ws.AmIsEnable, ws.PmIsEnable)
			if isContain {
				break
			}
		}
	}

	if i == length {
		fmt.Println("currently unavailable")
	}


	if weekDay == time.Sunday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.SundayAmStartTime, settings.SundayAmEndTime,
			settings.SundayPmStartTime, settings.SundayPmEndTime,duration, number)
	}else if weekDay == time.Monday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.MondayAmStartTime, settings.MondayAmEndTime,
			settings.MondayPmStartTime, settings.MondayPmEndTime,duration, number)
	}else if weekDay == time.Tuesday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.TuesdayAmStartTime, settings.TuesdayAmEndTime,
			settings.TuesdayPmStartTime, settings.TuesdayPmEndTime,duration, number)
	}else if weekDay == time.Wednesday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.WednesdayAmStartTime, settings.WednesdayAmEndTime,
			settings.WednesdayPmStartTime, settings.WednesdayPmEndTime,duration, number)
	}else if weekDay == time.Thursday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.ThursdayAmStartTime, settings.ThursdayAmEndTime,
			settings.ThursdayPmStartTime, settings.ThursdayPmEndTime,duration, number)
	}else if weekDay == time.Friday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.FridayAmStartTime, settings.FridayAmEndTime,
			settings.FridayPmStartTime, settings.FridayPmEndTime,duration, number)
	}else if weekDay == time.Saturday {
		isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(t, settings.SaturdayAmStartTime, settings.SaturdayAmEndTime,
			settings.SaturdayPmStartTime, settings.SaturdayPmEndTime,duration, number)
	}
	return isOk, nextAvailableDate
}

func (d *Dao) ContainTimeInRange(t time.Time, startTime string, endTime string, isAmEnable bool, isPmEnable bool) bool  {
	startTimeUTC, _ := d.ParseScheduleTimeToUTC(t, startTime, isAmEnable)
	endTimeUTC, _ := d.ParseScheduleTimeToUTC(t, endTime, !isPmEnable)
	return t.After(startTimeUTC) && t.Before(endTimeUTC)
}

func (d *Dao) GetWeekSchedule(settings *doctor.ScheduleSettings) map[time.Weekday]WeekSchedule {
	weekSchedules := make(map[time.Weekday]WeekSchedule)
	weekSchedules[time.Sunday] = WeekSchedule{
		AmIsEnable: settings.SundayAmIsEnable,
		AmStartTime: settings.SundayAmStartTime,
		AmEndTime: settings.SundayAmEndTime,

		PmIsEnable: settings.SundayPmIsEnable,
		PmStartTime: settings.SundayPmStartTime,
		PmEndTime: settings.SundayAmEndTime,
	}
	weekSchedules[time.Monday] = WeekSchedule{
		AmIsEnable: settings.MondayAmIsEnable,
		AmStartTime: settings.MondayAmStartTime,
		AmEndTime: settings.MondayAmEndTime,

		PmIsEnable: settings.MondayPmIsEnable,
		PmStartTime: settings.MondayPmStartTime,
		PmEndTime: settings.MondayAmEndTime,
	}
	weekSchedules[time.Tuesday] = WeekSchedule{
		AmIsEnable: settings.TuesdayAmIsEnable,
		AmStartTime: settings.TuesdayAmStartTime,
		AmEndTime: settings.TuesdayAmEndTime,

		PmIsEnable: settings.TuesdayPmIsEnable,
		PmStartTime: settings.TuesdayPmStartTime,
		PmEndTime: settings.TuesdayAmEndTime,
	}
	weekSchedules[time.Wednesday] = WeekSchedule{
		AmIsEnable: settings.WednesdayAmIsEnable,
		AmStartTime: settings.WednesdayAmStartTime,
		AmEndTime: settings.WednesdayAmEndTime,

		PmIsEnable: settings.WednesdayPmIsEnable,
		PmStartTime: settings.WednesdayPmStartTime,
		PmEndTime: settings.WednesdayAmEndTime,
	}
	weekSchedules[time.Thursday] = WeekSchedule{
		AmIsEnable: settings.ThursdayAmIsEnable,
		AmStartTime: settings.ThursdayAmStartTime,
		AmEndTime: settings.ThursdayAmEndTime,

		PmIsEnable: settings.ThursdayPmIsEnable,
		PmStartTime: settings.ThursdayPmStartTime,
		PmEndTime: settings.ThursdayAmEndTime,
	}
	weekSchedules[time.Friday] = WeekSchedule{
		AmIsEnable: settings.FridayAmIsEnable,
		AmStartTime: settings.FridayAmStartTime,
		AmEndTime: settings.FridayAmEndTime,

		PmIsEnable: settings.FridayPmIsEnable,
		PmStartTime: settings.FridayPmStartTime,
		PmEndTime: settings.FridayAmEndTime,
	}
	weekSchedules[time.Saturday] = WeekSchedule{
		AmIsEnable: settings.SaturdayAmIsEnable,
		AmStartTime: settings.SaturdayAmStartTime,
		AmEndTime: settings.SaturdayAmEndTime,

		PmIsEnable: settings.SaturdayPmIsEnable,
		PmStartTime: settings.SaturdayPmStartTime,
		PmEndTime: settings.SaturdayAmEndTime,
	}
	return weekSchedules
}

func (d *Dao) CalcNextAvailableDateForEachWeekDay(t time.Time,
	weekDayAMStartTime string, weekDayAMEndTime string, weekDayPMStartTime string, weekDayPMEndTime string,
	durationOfSlot int, numberOfSlot int) (bool, string) {
	AMStartTime, err := d.ParseScheduleTimeToUTC(weekDayAMStartTime, true)
	AMEndTime, err := d.ParseScheduleTimeToUTC(weekDayAMEndTime, true)

	PMStartTime, err := d.ParseScheduleTimeToUTC(weekDayPMStartTime, false)
	PMEndTime, err := d.ParseScheduleTimeToUTC(weekDayPMEndTime, false)
	if err != nil {
		return false, ""
	}
	if t.Before(AMStartTime)  {
		return true, time.Now().UTC().Format(time.RFC3339)
	}else if t.After(AMStartTime) && t.Before(AMEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(t, AMStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	} else if t.After(AMEndTime) && t.Before(PMEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(t, PMStartTime, durationOfSlot)
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
