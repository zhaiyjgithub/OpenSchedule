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

func (d *Dao) NewBulkUpdateRequest(esId string,
	isOnlineScheduleEnable bool, isInClinicBookEnable bool, isVirtualBookEnable bool,
	nextAvailableDateInClinic string, nextAvailableDateVirtual string) *elastic.BulkUpdateRequest {
	return elastic.NewBulkUpdateRequest().Index(database.DoctorIndexName).Id(esId).Doc(struct {
		IsOnlineScheduleEnable bool
		IsInClinicBookEnable bool
		IsVirtualBookEnable bool
		NextAvailableDateInClinic string
		NextAvailableDateVirtual string
	}{
		IsOnlineScheduleEnable: isOnlineScheduleEnable,
		IsInClinicBookEnable: isInClinicBookEnable,
		IsVirtualBookEnable: isVirtualBookEnable,
		NextAvailableDateInClinic: nextAvailableDateInClinic,
		NextAvailableDateVirtual: nextAvailableDateVirtual,
	})
}

func (d *Dao) SyncCertainDoctorNextAvailableDateToES(npi int64,
	isOnlineScheduleEnable bool, isInClinicBookEnable bool, isVirtualBookEnable bool,
	nextAvailableDateInClinic string, nextAvailableDateVirtual string) error {
	esId := d.GetDoctorInfoFromES(npi)
	if len(esId) == 0 {
		errs := fmt.Sprintf("esid is empty: %d", npi)
		return errors.New(errs)
	}
	req := d.NewBulkUpdateRequest(esId,
		isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
		nextAvailableDateInClinic, nextAvailableDateVirtual)
	_, err := d.elasticSearchEngine.Bulk().Add(req).Do(context.TODO())
	return err
}

func (d *Dao) GetESBulkUpdateRequest(npi int64,
	isOnlineScheduleEnable bool, isInClinicBookEnable bool, isVirtualBookEnable bool,
	nextAvailableDateInClinic string, nextAvailableDateVirtual string) (error, *elastic.BulkUpdateRequest) {
	esId := d.GetDoctorInfoFromES(npi)
	if len(esId) == 0 {
		errs := fmt.Sprintf("esid is empty: %d", npi)
		return errors.New(errs), nil
	}
	req := d.NewBulkUpdateRequest(esId,
		isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
	nextAvailableDateInClinic, nextAvailableDateVirtual)
	return nil, req
}

func (d *Dao) BulkUpdateToES(reqs []*elastic.BulkUpdateRequest) error {
	bulkService := d.elasticSearchEngine.Bulk()
	for _, req := range reqs {
		bulkService.Add(req)
	}
	_, err := bulkService.Do(context.TODO())
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

func (d *Dao) DeleteESDoctorById(esId string) error {
	if len(esId) == 0 {
		errs := fmt.Sprintf("esid is empty")
		return errors.New(errs)
	}
	req := elastic.NewBulkDeleteRequest().Index(database.DoctorIndexName).Id(esId)
	_, err := d.elasticSearchEngine.Bulk().Add(req).Do(context.TODO())
	return err
}

func (d *Dao) SyncDoctorToES(doctor * doctor.Doctor) error {
	return nil
}

func (d *Dao) GetDuplicateDoctorInfoFromES(npi int64) []string {
	q := elastic.NewTermQuery("Npi", npi)
	result, err := d.elasticSearchEngine.Search().Index(database.DoctorIndexName).
		Size(2).Query(q).Pretty(true).Do(context.Background())
	var esIds []string
	if err != nil {
		fmt.Println("search failed")
		return esIds
	}
	for _, hit := range result.Hits.Hits {
		esIds = append(esIds, hit.Id)
	}
	return esIds
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
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.SundayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.SundayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.SundayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.SundayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.SundayAmAppointmentType,  settings.SundayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.SundayPmAppointmentType, settings.SundayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Monday && (settings.MondayAmIsEnable || settings.MondayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.MondayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.MondayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.MondayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.MondayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.MondayAmAppointmentType,  settings.MondayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.MondayPmAppointmentType, settings.MondayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Tuesday && (settings.TuesdayAmIsEnable || settings.TuesdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.TuesdayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.TuesdayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.TuesdayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.TuesdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.TuesdayAmAppointmentType,  settings.TuesdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.TuesdayPmAppointmentType, settings.TuesdayPmIsEnable, pmStartDateTime, pmEndDateTime,duration, number, closedDateSettings)
		}else if weekDay == time.Wednesday && (settings.WednesdayAmIsEnable || settings.WednesdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.WednesdayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.WednesdayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.WednesdayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.WednesdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.WednesdayAmAppointmentType,  settings.WednesdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.WednesdayPmAppointmentType, settings.WednesdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Thursday && (settings.ThursdayAmIsEnable || settings.ThursdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.ThursdayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.ThursdayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.ThursdayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.ThursdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.ThursdayAmAppointmentType,  settings.ThursdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.ThursdayPmAppointmentType, settings.ThursdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Friday && (settings.FridayAmIsEnable || settings.FridayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.FridayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.FridayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.FridayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.FridayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.FridayAmAppointmentType,  settings.FridayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.FridayPmAppointmentType, settings.FridayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}else if weekDay == time.Saturday && (settings.SaturdayAmIsEnable || settings.SaturdayPmIsEnable) {
			amStartDateTime, err := d.ConvertScheduleTimeOffsetToDateTime(nextTime, settings.SaturdayAmStartTimeOffset)
			if err != nil {
				return ""
			}
			amEndDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.SaturdayAmEndTimeOffset))
			pmStartDateTime := amStartDateTime.Add(time.Minute * time.Duration(settings.SaturdayPmStartTimeOffset))
			if err != nil {
				return ""
			}
			pmEndDateTime := pmStartDateTime.Add(time.Minute * time.Duration(settings.SaturdayPmEndTimeOffset))
			nextAvailableDate = d.CalcNextAvailableDateForEachWeekDay(currentTime, appointmentType, settings.SaturdayAmAppointmentType,  settings.SaturdayAmIsEnable, amStartDateTime, amEndDateTime,
				settings.SaturdayPmAppointmentType, settings.SaturdayPmIsEnable, pmStartDateTime, pmEndDateTime, duration, number, closedDateSettings)
		}
		if nextAvailableDate != constant.InvalidDateTime {
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
		amStartTime, amEndTime = d.CalcAvailableTimeRangeByClosedDate(currentTime, weekDayAMStartDateTime, weekDayAMEndTime, closedDateSettings.AmStartDateTime, closedDateSettings.AmEndDateTime)
		pmStartTime, pmEndTime = d.CalcAvailableTimeRangeByClosedDate(currentTime, weekDayPMStartDateTime, weekDayPMEndTime, closedDateSettings.PmStartDateTime, closedDateSettings.PmEndDateTime)
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
	return constant.InvalidDateTime
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

func (d *Dao) ConvertScheduleTimeOffsetToDateTime(t time.Time, offset int) (time.Time, error) {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	utcTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Add(time.Minute * time.Duration(offset))
	return utcTime, nil
}

func (d *Dao) CalcAvailableTimeRangeByClosedDate(currentDateTime time.Time, startDateTime time.Time, endDateTime time.Time,
	closedStartDateTime time.Time, closedEndDateTime time.Time) (*time.Time, *time.Time) {
	if closedStartDateTime.Equal(constant.DefaultTimeStamp) || closedEndDateTime.Equal(constant.DefaultTimeStamp) {
		return nil, nil
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
		if currentDateTime.Before(closedStartDateTime) {
			return &startDateTime, &closedStartDateTime
		} else {
			return &closedStartDateTime, &endDateTime
		}
	}else if (endDateTime.Before(closedStartDateTime) || endDateTime.Equal(closedStartDateTime)) ||
		(startDateTime.Before(closedEndDateTime) || startDateTime.Equal(closedEndDateTime)) {
		return &startDateTime, &endDateTime
	}else if startDateTime.After(closedEndDateTime) || startDateTime.Equal(closedEndDateTime)  {
		return &startDateTime, &endDateTime
	} else {
		return nil, nil
	}
}

func (d *Dao) AddAppointment(appointment doctor.Appointments) error {
	db := d.engine.Create(&appointment)
	return db.Error
}

func (d *Dao) GetAppointmentByRange(
	npi int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
	) []*doctor.Appointments {
	appts := make([]*doctor.Appointments, 0)
	_ = d.engine.Where("npi =? AND appointment_status = ? AND appointment_date >= ? AND appointment_date <= ?",
			npi, appointmentStatus, startDate, endDate).Find(&appts)
	return appts
}

func (d *Dao) GetSettingsByNpiList(npiList []int64) []*doctor.ScheduleSettings {
	list := make([]*doctor.ScheduleSettings, 0)
	_ = d.engine.Where("npi IN ?", npiList).Find(&list)
	return list
}