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

func (d *Dao) CalcNextAvailableDate(currentTime time.Time, appointmentType constant.AppointmentType, settings *doctor.ScheduleSettings) (string)  {
	duration := settings.DurationPerSlot
	number := settings.NumberPerSlot
	nextAvailableDate := ""
	for i := 0; i < 14; i ++ {// future 2 weeks
		nextTime := currentTime.Add(time.Hour*24*time.Duration(i))
		weekDay := nextTime.Weekday()
		closedDateSettings, _ := d.GetClosedDateByDateTime(settings.Npi, nextTime)
		if weekDay == time.Sunday && (settings.SundayAmIsEnable || settings.SundayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.SundayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.SundayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.SundayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.SundayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.SundayAmAppointmentType,  settings.SundayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.SundayPmAppointmentType, settings.SundayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Monday && (settings.MondayAmIsEnable || settings.MondayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.MondayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.MondayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.MondayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.MondayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.MondayAmAppointmentType,  settings.MondayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.MondayPmAppointmentType, settings.MondayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Tuesday && (settings.TuesdayAmIsEnable || settings.TuesdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.TuesdayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.TuesdayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.TuesdayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.TuesdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.TuesdayAmAppointmentType,  settings.TuesdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.TuesdayPmAppointmentType, settings.TuesdayPmIsEnable, pmStartDateTime, pmEndDateTime,duration, number, closedDateSettings)
		}else if weekDay == time.Wednesday && (settings.WednesdayAmIsEnable || settings.WednesdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.WednesdayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.WednesdayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.WednesdayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.WednesdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.WednesdayAmAppointmentType,  settings.WednesdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.WednesdayPmAppointmentType, settings.WednesdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Thursday && (settings.ThursdayAmIsEnable || settings.ThursdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.ThursdayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.ThursdayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.ThursdayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.ThursdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.ThursdayAmAppointmentType,  settings.ThursdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.ThursdayPmAppointmentType, settings.ThursdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Friday && (settings.FridayAmIsEnable || settings.FridayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.FridayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.FridayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.FridayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.FridayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.FridayAmAppointmentType,  settings.FridayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.FridayPmAppointmentType, settings.FridayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Saturday && (settings.SaturdayAmIsEnable || settings.SaturdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.SaturdayAmStartTime)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Duration(settings.SaturdayAmEndTimeOffset))
			pmStartDateTime, err := d.ConvertScheduleTimeToTime(currentTime, settings.SaturdayPmStartTime)
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Duration(settings.SaturdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.SaturdayAmAppointmentType,  settings.SaturdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.SaturdayPmAppointmentType, settings.SaturdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}
		if len(nextAvailableDate) > 0 {
			break
		}
	}
	return nextAvailableDate
}

func (d *Dao) ContainTimeInRange(t time.Time, startTime string, endTime string, isAmEnable bool, isPmEnable bool) bool  {
	startTimeUTC, _ := d.ParseScheduleTimeToUTC(t, startTime, isAmEnable)
	endTimeUTC, _ := d.ParseScheduleTimeToUTC(t, endTime, !isPmEnable)
	return t.After(*startTimeUTC) && t.Before(*endTimeUTC)
}

func (d *Dao) CalcNextAvailableDateForEachWeekDay(currentTime time.Time, appointmentType constant.AppointmentType,
	amAppointmentType constant.AppointmentType, isAmEnable bool, weekDayAMStartDateTime time.Time, weekDayAMEndTime time.Time,
	pmAppointmentType constant.AppointmentType, isPmEnable bool, weekDayPMStartDateTime time.Time, weekDayPMEndTime time.Time,
	durationOfSlot int, numberOfSlot int,
	closedDateSettings *doctor.ClosedDateSettings) string {
	//calc the next available date by the closed date.
	var amStartTime, amEndTime, pmStartTime, pmEndTime *time.Time
	if closedDateSettings != nil {
		amStartTime, amEndTime = d.CalcAvailableTimeRangeByClosedDate(weekDayAMStartDateTime, weekDayAMEndTime, closedDateSettings.AmStartDateTime, closedDateSettings.AmEndDateTime)
		pmStartTime, pmEndTime = d.CalcAvailableTimeRangeByClosedDate(weekDayPMStartDateTime, weekDayPMEndTime, closedDateSettings.PmStartDateTime, closedDateSettings.PmEndDateTime)
	}else {
		amStartTime = &weekDayAMStartDateTime
		amEndTime= &weekDayAMEndTime
		pmStartTime= &weekDayPMStartDateTime
		pmEndTime = &weekDayPMEndTime
	}

	if appointmentType == amAppointmentType && isAmEnable && amStartTime != nil && currentTime.Before(*amStartTime)  {
		return amStartTime.Format(time.RFC3339)
	}else if appointmentType == amAppointmentType && isAmEnable && amStartTime != nil && currentTime.After(*amStartTime) && amEndTime != nil && currentTime.Before(*amEndTime) {
		nextAvailableDateTime := d.MatchDateTimeByDuration(currentTime, *amStartTime, durationOfSlot)
		return nextAvailableDateTime
	}else if appointmentType == pmAppointmentType && isPmEnable && pmStartTime != nil && currentTime.Before(*pmStartTime)  {
		return pmStartTime.Format(time.RFC3339)
	} else if appointmentType == pmAppointmentType && isPmEnable && amStartTime != nil && currentTime.After(*amEndTime) && pmStartTime != nil && currentTime.Before(*pmEndTime) {
		nextAvailableDateTime := d.MatchDateTimeByDuration(currentTime, *pmStartTime, durationOfSlot)
		return nextAvailableDateTime
	}
	return ""
}

func (d *Dao) MatchDateTimeByDuration(now time.Time, startTime time.Time, durationOfSlot int) string {
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

func (d *Dao)ParseScheduleTimeToUTC(t time.Time, scheduleTime string, isAM bool) (*time.Time, error) {
	if !utils.CheckDateTime(scheduleTime) {
		return nil, errors.New("param error")
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
	return &utcTime, nil
}

func (d *Dao) ConvertScheduleTimeToTime(t time.Time, scheduleTime string) (time.Time, error) {
	if !utils.CheckDateTime(scheduleTime) {
		return constant.DefaultTimeStamp, errors.New("param error")
	}
	year := t.Year()
	month := t.Month()
	day := t.Day()

	dateTime := strings.Split(scheduleTime, ":")
	hour, _ := strconv.Atoi(dateTime[0])
	min, _ := strconv.Atoi(dateTime[1])
	utcTime := time.Date(year, month, day, hour, min, 0, 0, time.UTC)
	return utcTime, nil
}

func (d *Dao) CalcAvailableTimeRangeByClosedDate(startDateTime time.Time, endDateTime time.Time,
	closedStartDateTime time.Time, closedEndDateTime time.Time) (*time.Time, *time.Time) {
	if closedStartDateTime.Equal(constant.DefaultTimeStamp) || closedEndDateTime.Equal(constant.DefaultTimeStamp) {
		return &startDateTime, &endDateTime
	}
	
	if (startDateTime.Equal(closedStartDateTime) || startDateTime.After(closedStartDateTime)) &&
		(startDateTime.Equal(closedEndDateTime) || startDateTime.Before(closedEndDateTime)) &&
		endDateTime.After(closedEndDateTime) {
		return &closedEndDateTime, &endDateTime
	}else if endDateTime.After(closedStartDateTime) &&
		(endDateTime.Equal(closedEndDateTime) || endDateTime.Before(closedEndDateTime)) &&
		startDateTime.Before(closedStartDateTime) {
		return &startDateTime, &closedStartDateTime
	}else if startDateTime.Before(closedStartDateTime) && endDateTime.After(closedEndDateTime) {
		return &startDateTime, &closedStartDateTime
	}else if (endDateTime.Before(closedStartDateTime) || endDateTime.Equal(closedStartDateTime)) ||
		(startDateTime.Before(closedEndDateTime) || startDateTime.Equal(closedEndDateTime)) {
		return &startDateTime, &endDateTime
	}else if startDateTime.After(closedEndDateTime) || startDateTime.Equal(closedEndDateTime)  {
		return &startDateTime, &endDateTime
	} else {
		return nil, nil
	}
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