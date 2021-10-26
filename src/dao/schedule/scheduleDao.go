package schedule

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/utils"
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Dao struct {
	engine *gorm.DB
	elasticSearchEngine *elastic.Client
}

func NewDao(engine *gorm.DB, elasticSearchEngine *elastic.Client) *Dao {
	return &Dao{engine: engine, elasticSearchEngine: elasticSearchEngine}
}

func (d *Dao) AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error {
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ?", closeDateSettings.Npi).First(st)
	if db.Error != nil {
		db := d.engine.Create(closeDateSettings)
		return db.Error
	}else {
		return errors.New("can't update existing closed date")
	}
}

func (d *Dao) DeleteClosedDate(npi int64) error {
	st := &doctor.ClosedDateSettings{}
	db := d.engine.Where("npi = ?", npi).Delete(st)
	return db.Error
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

func (d *Dao) SyncCertainDoctorNextAvailableDateToES(npi int64, nextAvailableDateInClinic string, nextAvailableDateVirtual string) error {
	esId := d.GetDoctorInfoFromES(npi)
	req := elastic.NewBulkUpdateRequest().Index(database.DoctorIndexName).Id(esId).Doc(struct {
		NextAvailableDateInClinic string
		NextAvailableDateVirtual string
	}{
		NextAvailableDateInClinic: nextAvailableDateInClinic,
		NextAvailableDateVirtual: nextAvailableDateVirtual,
	})
	_, err := d.elasticSearchEngine.Bulk().Add(req).Do(context.TODO())
	return err
}

func (d *Dao) GetDoctorInfoFromES(npi int64) string {
	q := elastic.NewTermQuery("Npi", npi)
	result, err := d.elasticSearchEngine.Search().Index(database.DoctorIndexName).
		Size(1).Query(q).Pretty(true).Do(context.Background())
	esId := ""
	if err != nil {
		fmt.Println("search failed")
		return esId
	}
	for _, hit := range result.Hits.Hits {
		esId = hit.Id
	}
	return esId
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
	durationOfSlot int, numberOfSlot int, settings doctor.ClosedDateSettings) (bool, string) {


	//calc the next available date by the closed date.
	newAmStartTime, newAmEndTime := d.CalcAvailableTimeByClosedDate(weekDayAMStartTime, weekDayAMEndTime, settings.AmStartTime, settings.AmEndTime)
	newPmStartTime, newPmEndTime := d.CalcAvailableTimeByClosedDate(weekDayPMStartTime, weekDayPMEndTime, settings.PmStartTime, settings.PmEndTime)

	amStartTime, err := d.ParseScheduleTimeToUTC(nextTime, newAmStartTime, true)
	amEndTime, err := d.ParseScheduleTimeToUTC(nextTime, newAmEndTime, true)

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

func (d *Dao) CalcAvailableTimeByClosedDate(startTime string, endTime string,
	closedStartTime string, closedEndTime string) (string, string) {
	if !utils.CheckDateTime(startTime) || !utils.CheckDateTime(endTime) ||
		!utils.CheckDateTime(closedStartTime) || !utils.CheckDateTime(closedEndTime) {
		return "", ""
	}
	startTimeMinutes, _ := d.ConvertHourMinToMinutes(startTime)
	endTimeMinutes, _ := d.ConvertHourMinToMinutes(endTime)
	closedStartTimeMinutes, _ := d.ConvertHourMinToMinutes(closedStartTime)
	closedEndTimeMinutes, _ := d.ConvertHourMinToMinutes(closedEndTime)

	if startTimeMinutes >= closedStartTimeMinutes && startTimeMinutes <= closedEndTimeMinutes &&
		endTimeMinutes > closedEndTimeMinutes {
		startTimeMinutes = closedEndTimeMinutes
	}else if endTimeMinutes > closedStartTimeMinutes && endTimeMinutes <= closedEndTimeMinutes &&
		startTimeMinutes < closedStartTimeMinutes {
		endTimeMinutes = closedStartTimeMinutes
	}else if startTimeMinutes < closedStartTimeMinutes && endTimeMinutes > closedEndTimeMinutes {
		endTimeMinutes = closedStartTimeMinutes
	}else {
		return "", ""
	}
	return d.ReverseMinutesToHourMin(startTimeMinutes), d.ReverseMinutesToHourMin(endTimeMinutes)
}

func (d *Dao) ConvertHourMinToMinutes(dateTime string) (int, error) {
	hour, min, err := d.ConvertDateTimeToHourMin(dateTime)
	if err != nil {
		return 0, err
	}
	return hour*60 + min, nil
}

func (d *Dao) ReverseMinutesToHourMin (minutes int) string {
	hour := minutes/60
	min := minutes%60
	hourStr := ""
	if hour < 10 {
		hourStr = fmt.Sprintf("0%d", hour)
	}else {
		hourStr = fmt.Sprintf("%d", hour)
	}
	minStr := ""
	if min < 10 {
		minStr = fmt.Sprintf("0%d", min)
	}else {
		minStr = fmt.Sprintf("%d", min)
	}
	return fmt.Sprintf("%s:%s", hourStr, minStr)
}

func (d *Dao) ConvertDateTimeToHourMin(dateTime string) (int, int, error) {
	if !utils.CheckDateTime(dateTime) {
		errStr := fmt.Sprintf("param error: %s", dateTime)
		return 0, 0, errors.New(errStr)
	}
	dateTimes := strings.Split(dateTime, ":")
	hour, _ := strconv.Atoi(dateTimes[0])
	min, _ := strconv.Atoi(dateTimes[1])
	return hour, min, nil
}