package schedule

import (
	"OpenSchedule/src/constant"
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

func (d *Dao) SyncCertainDoctorScheduleNextAvailableDateToES()  {

}

func (d *Dao) CalcNextAvailableDate(currentTime time.Time, appointmentType constant.AppointmentType, settings *doctor.ScheduleSettings) (bool, string)  {
	duration := settings.DurationPerSlot
	number := settings.NumberPerSlot
	isOk := false
	nextAvailableDate := ""
	for i := 0; i < 14; i ++ {// future 2 weeks
		nextTime := currentTime.Add(time.Hour*24*time.Duration(i))
		weekDay := nextTime.Weekday()
		if weekDay == time.Sunday && (settings.SundayAmIsEnable || settings.SundayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.SundayAmAppointmentType,  settings.SundayAmIsEnable, settings.SundayAmStartTime, settings.SundayAmEndTime,
				settings.SundayPmAppointmentType, settings.SundayPmIsEnable, settings.SundayPmStartTime, settings.SundayPmEndTime,duration, number)
		}else if weekDay == time.Monday && (settings.MondayAmIsEnable || settings.MondayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.MondayAmAppointmentType, settings.MondayAmIsEnable, settings.MondayAmStartTime, settings.MondayAmEndTime,
				settings.MondayPmAppointmentType, settings.MondayPmIsEnable, settings.MondayPmStartTime, settings.MondayPmEndTime,duration, number)
		}else if weekDay == time.Tuesday && (settings.TuesdayAmIsEnable || settings.TuesdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.TuesdayAmAppointmentType, settings.TuesdayAmIsEnable, settings.TuesdayAmStartTime, settings.TuesdayAmEndTime,
				settings.TuesdayPmAppointmentType, settings.TuesdayPmIsEnable, settings.TuesdayPmStartTime, settings.TuesdayPmEndTime,duration, number)
		}else if weekDay == time.Wednesday && (settings.WednesdayAmIsEnable || settings.WednesdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.WednesdayAmAppointmentType, settings.WednesdayAmIsEnable, settings.WednesdayAmStartTime, settings.WednesdayAmEndTime,
				settings.WednesdayPmAppointmentType, settings.WednesdayPmIsEnable, settings.WednesdayPmStartTime, settings.WednesdayPmEndTime,duration, number)
		}else if weekDay == time.Thursday && (settings.ThursdayAmIsEnable || settings.ThursdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.ThursdayAmAppointmentType, settings.ThursdayAmIsEnable, settings.ThursdayAmStartTime, settings.ThursdayAmEndTime,
				settings.ThursdayPmAppointmentType,settings.ThursdayPmIsEnable, settings.ThursdayPmStartTime, settings.ThursdayPmEndTime,duration, number)
		}else if weekDay == time.Friday && (settings.FridayAmIsEnable || settings.FridayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.FridayAmAppointmentType, settings.FridayAmIsEnable, settings.FridayAmStartTime, settings.FridayAmEndTime,
				settings.FridayPmAppointmentType, settings.FridayPmIsEnable, settings.FridayPmStartTime, settings.FridayPmEndTime,duration, number)
		}else if weekDay == time.Saturday && (settings.SaturdayAmIsEnable || settings.SaturdayPmIsEnable) {
			isOk, nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, nextTime, settings.SaturdayAmAppointmentType, settings.SaturdayAmIsEnable, settings.SaturdayAmStartTime, settings.SaturdayAmEndTime,
				settings.SaturdayPmAppointmentType ,settings.SaturdayPmIsEnable, settings.SaturdayPmStartTime, settings.SaturdayPmEndTime,duration, number)
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

func (d *Dao) CalcNextAvailableDateForEachWeekDay(currentTime time.Time, appointmentType constant.AppointmentType, nextTime time.Time,
	amAppointmentType constant.AppointmentType, isAmEnable bool, weekDayAMStartTime string, weekDayAMEndTime string,
	pmAppointmentType constant.AppointmentType, isPmEnable bool, weekDayPMStartTime string, weekDayPMEndTime string,
	durationOfSlot int, numberOfSlot int) (bool, string) {
	amStartTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayAMStartTime, true)
	amEndTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayAMEndTime, true)

	pmStartTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayPMStartTime, false)
	pmEndTime, err := d.ParseScheduleTimeToUTC(nextTime, weekDayPMEndTime, false)
	if err != nil {
		return false, ""
	}
	if appointmentType == amAppointmentType && isAmEnable && currentTime.Before(amStartTime)  {
		return true, amStartTime.Format(time.RFC3339)
	}else if appointmentType == amAppointmentType && isAmEnable && currentTime.After(amStartTime) && currentTime.Before(amEndTime) {
		nextAvailableDateTime := d.CalcNextAvailableDateForTimeRange(currentTime, amStartTime, durationOfSlot)
		return true, nextAvailableDateTime
	}else if appointmentType == pmAppointmentType && isPmEnable && currentTime.Before(pmStartTime)  {
		return true, pmStartTime.Format(time.RFC3339)
	} else if appointmentType == pmAppointmentType && isPmEnable && currentTime.After(amEndTime) && currentTime.Before(pmEndTime) {
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
