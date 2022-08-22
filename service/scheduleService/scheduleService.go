package scheduleService

import (
	"OpenSchedule/constant"
	"OpenSchedule/dao/schedule"
	"OpenSchedule/database"
	"OpenSchedule/model/doctor"
	"OpenSchedule/model/viewModel"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type Service interface {
	SetScheduleSettings(settings doctor.ScheduleSettings) error
	GetScheduleSettings(npi int64) doctor.ScheduleSettings
	AddClosedDate(closeDateSettings doctor.ClosedDateSettings) error
	DeleteClosedDate(npi int64, id int) error
	GetClosedDate(npi int64) []doctor.ClosedDateSettings
	SyncCertainDoctorScheduleNextAvailableDateToES(settings doctor.ScheduleSettings) error
	SyncMultiDoctorsScheduleNextAvailableDateToES(doctors []*doctor.Doctor) error
	IsExist(npi int64) bool
	SyncDoctorToES(doctor *doctor.Doctor) error
	AddAppointment(appointment doctor.Appointment) error
	GetAppointmentByRange(
		npi int64,
		appointmentStatus constant.AppointmentStatus,
		startDate time.Time,
		endDate time.Time,
	) []*doctor.Appointment
	GetAppointmentsByRange(
		npi []int64,
		appointmentStatus constant.AppointmentStatus,
		startDate time.Time,
		endDate time.Time,
	) []*doctor.Appointment
	GetSettingsByNpiList(npiList []int64) []doctor.ScheduleSettings
	GetClosedDateByRange(npi []int64, from time.Time, to time.Time) []doctor.ClosedDateSettings
	GetBookedAppointmentsTimeSlotsByNpiList(npiList []int64, startDate time.Time, endTime time.Time) map[int64]map[string][]doctor.TimeSlot
	GetClosedDateByNpiList(npiList []int64, startDate time.Time, endDate time.Time) map[int64][]doctor.ClosedDateSettings
	GetScheduleSettingByNpiList(npiList []int64) map[int64]doctor.ScheduleSettings
	GetDoctorTimeSlotsByDate(setting doctor.ScheduleSettings, startDate time.Time, endDate time.Time, bookedTimeSlots map[string][]doctor.TimeSlot,
		closeDateSetting []doctor.ClosedDateSettings) []viewModel.TimeSlotPerDay
}

func NewService() Service {
	return &service{dao: schedule.NewDao(database.GetMySqlEngine(), database.GetElasticSearchEngine())}
}

type service struct {
	dao *schedule.Dao
}

func (s *service) SetScheduleSettings(setting doctor.ScheduleSettings) error {
	return s.dao.SetScheduleSettings(setting)
}

func (s *service) GetScheduleSettings(npi int64) doctor.ScheduleSettings {
	return s.dao.GetScheduleSettings(npi)
}

func (s *service) SyncCertainDoctorScheduleNextAvailableDateToES(settings doctor.ScheduleSettings) error {
	if settings.Npi == 0 {
		return errors.New("param is nil")
	}
	currentTime := time.Now().UTC()
	nextAvailableDateInClinic := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, settings)
	nextAvailableDateVirtual := s.dao.CalcNextAvailableDate(currentTime, constant.Virtual, settings)

	isInClinicBookEnable := nextAvailableDateInClinic != constant.InvalidDateTime
	isVirtualBookEnable := nextAvailableDateVirtual != constant.InvalidDateTime
	isOnlineScheduleEnable := isInClinicBookEnable || isVirtualBookEnable

	return s.dao.SyncCertainDoctorNextAvailableDateToES(settings.Npi,
		isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
		nextAvailableDateInClinic, nextAvailableDateVirtual)
}

func (s *service) SyncMultiDoctorsScheduleNextAvailableDateToES(doctors []*doctor.Doctor) error {
	var reqs []*elastic.BulkUpdateRequest
	for _, doc := range doctors {
		settings := s.GetScheduleSettings(doc.Npi)
		if settings.Npi == 0 {
			fmt.Println("settings not found: ", doc.Npi)
			continue
		}
		currentTime := time.Now().UTC()
		nextAvailableDateInClinic := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, settings)
		nextAvailableDateVirtual := s.dao.CalcNextAvailableDate(currentTime, constant.Virtual, settings)

		isInClinicBookEnable := nextAvailableDateInClinic != constant.InvalidDateTime
		isVirtualBookEnable := nextAvailableDateVirtual != constant.InvalidDateTime
		isOnlineScheduleEnable := isInClinicBookEnable || isVirtualBookEnable

		err, req := s.dao.GetESBulkUpdateRequest(settings.Npi,
			isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
			nextAvailableDateInClinic, nextAvailableDateVirtual)
		if err != nil {
			fmt.Println("sync multi doctor error: ", err.Error(), doc.Npi)
		}
		if req != nil {
			reqs = append(reqs, req)
		}
	}

	return s.dao.BulkUpdateToES(reqs)
}

func (s *service) AddClosedDate(closeDateSettings doctor.ClosedDateSettings) error {
	return s.dao.AddClosedDate(closeDateSettings)
}

func (s *service) DeleteClosedDate(npi int64, id int) error {
	return s.dao.DeleteClosedDateByID(npi, id)
}

func (s *service) GetClosedDate(npi int64) []doctor.ClosedDateSettings {
	return s.dao.GetClosedDate(npi)
}

func (s *service) IsExist(npi int64) bool {
	return s.dao.IsExist(npi)
}

func (s *service) SyncDoctorToES(doctor *doctor.Doctor) error {
	return s.dao.SyncDoctorToES(doctor)
}

func (s *service) AddAppointment(appointment doctor.Appointment) error {
	return s.dao.AddAppointment(appointment)
}

func (s *service) GetAppointmentByRange(
	npi int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctor.Appointment {
	return s.dao.GetAppointmentByRange(npi, appointmentStatus, startDate, endDate)
}

func (s *service) GetAppointmentsByRange(
	npi []int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctor.Appointment {
	return s.dao.GetAppointmentsByRange(npi, appointmentStatus, startDate, endDate)
}

func (s *service) GetSettingsByNpiList(npiList []int64) []doctor.ScheduleSettings {
	return s.dao.GetSettingsByNpiList(npiList)
}

func (s *service) GetClosedDateByRange(npi []int64, from time.Time, to time.Time) []doctor.ClosedDateSettings {
	return s.dao.GetClosedDateByRange(npi, from, to)
}

// todo: Refactor appoint type

func (s *service) GetBookedAppointmentsTimeSlotsByNpiList(npiList []int64, startDate time.Time, endTime time.Time) map[int64]map[string][]doctor.TimeSlot {
	appointments := s.GetAppointmentsByRange(npiList, constant.Requested, startDate, endTime)
	bTimeSlots := make(map[int64]map[string][]doctor.TimeSlot)
	for _, appt := range appointments {
		bookedTimeSlotsPerNpi, ok := bTimeSlots[appt.Npi]
		offset := appt.AppointmentDate.Hour()*60 + appt.AppointmentDate.Minute()
		dateKey := fmt.Sprintf("%d-%d-%d", appt.AppointmentDate.Year(), appt.AppointmentDate.Month(), appt.AppointmentDate.Day())
		if !ok {
			bookedTimeSlotsPerNpi = make(map[string][]doctor.TimeSlot)
			bookedTimeSlotsPerNpi[dateKey] = []doctor.TimeSlot{{Offset: offset, AvailableSlotsNumber: 1}}
		} else {
			bookedTimeSlotsPerNpi[dateKey] = append(bookedTimeSlotsPerNpi[dateKey], doctor.TimeSlot{Offset: offset, AvailableSlotsNumber: 1})
		}
		bTimeSlots[appt.Npi] = bookedTimeSlotsPerNpi
	}
	return bTimeSlots
}

func (s *service) GetClosedDateByNpiList(npiList []int64, startDate time.Time, endDate time.Time) map[int64][]doctor.ClosedDateSettings {
	closeDateSettings := s.GetClosedDateByRange(npiList, startDate, endDate)
	settingMap := make(map[int64][]doctor.ClosedDateSettings)
	for _, setting := range closeDateSettings {
		settingMap[setting.Npi] = append(settingMap[setting.Npi], setting)
	}
	return settingMap
}

func (s *service) GetScheduleSettingByNpiList(npiList []int64) map[int64]doctor.ScheduleSettings {
	list := s.GetSettingsByNpiList(npiList)
	settingMap := make(map[int64]doctor.ScheduleSettings)
	for _, setting := range list {
		settingMap[setting.Npi] = setting
	}
	return settingMap
}

func (s *service) GetDoctorTimeSlotsByDate(setting doctor.ScheduleSettings, startDate time.Time, endDate time.Time, bookedTimeSlots map[string][]doctor.TimeSlot,
	closeDateSetting []doctor.ClosedDateSettings) []viewModel.TimeSlotPerDay {
	timeSlots := make([]viewModel.TimeSlotPerDay, 0)
	dateRange := int(endDate.Sub(startDate).Hours() / 24)
	for i := 0; i < dateRange; i++ {
		targetDate := startDate.AddDate(0, 0, i)
		dateKey := fmt.Sprintf("%d-%d-%d", targetDate.Year(), targetDate.Month(), targetDate.Day())
		bookedTimeSlots, ok := bookedTimeSlots[dateKey]
		timeSlotsPerDay := make([]doctor.TimeSlot, 0)
		if ok {
			timeSlotsPerDay = s.GetDoctorTimeSlotsPerDay(setting, targetDate, bookedTimeSlots)
		} else {
			timeSlotsPerDay = s.GetDoctorTimeSlotsPerDay(setting, targetDate, make([]doctor.TimeSlot, 0))
		}
		targetClosetDate := s.getClosedDateByDate(closeDateSetting, targetDate)
		filterTimeSlotsPerDay := s.filterTimeSlotsByClosedDate(targetDate, timeSlotsPerDay, targetClosetDate)
		timeSlots = append(timeSlots, viewModel.TimeSlotPerDay{Date: targetDate, TimeSlots: filterTimeSlotsPerDay})
	}
	return timeSlots
}

func (s *service) getClosedDateByDate(closedDateList []doctor.ClosedDateSettings, targetDate time.Time) doctor.ClosedDateSettings {
	var targetClosedDate doctor.ClosedDateSettings
	for i := 0; i < len(closedDateList); i++ {
		closedDate := closedDateList[i]
		if closedDate.StartDate.Year() == targetDate.Year() &&
			closedDate.StartDate.Month() == targetDate.Month() &&
			closedDate.StartDate.Day() == targetDate.Day() {
			targetClosedDate = closedDate
			break
		}
	}
	return targetClosedDate
}

func (s *service) filterTimeSlotsByClosedDate(targetDate time.Time, timeSlots []doctor.TimeSlot, closedDate doctor.ClosedDateSettings) []doctor.TimeSlot {
	if closedDate.AmStartDateTime.IsZero() && closedDate.PmStartDateTime.IsZero() {
		return timeSlots
	}
	var filterList []doctor.TimeSlot
	targetDateZero := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), 0, 0, 0, 0, time.UTC)
	for _, timeSlot := range timeSlots {
		timeSlotDateTime := targetDateZero.Add(time.Minute * time.Duration(timeSlot.Offset))
		if (timeSlotDateTime.Equal(closedDate.AmStartDateTime) || timeSlotDateTime.After(closedDate.AmStartDateTime) &&
			(timeSlotDateTime.Equal(closedDate.AmEndDateTime) || timeSlotDateTime.Before(closedDate.AmEndDateTime))) ||
			((timeSlotDateTime.Equal(closedDate.PmStartDateTime) || timeSlotDateTime.After(closedDate.PmStartDateTime)) &&
				(timeSlotDateTime.Equal(closedDate.PmEndDateTime) || timeSlotDateTime.Before(closedDate.PmEndDateTime))) {
			continue
		} else {
			filterList = append(filterList, timeSlot)
		}
	}
	return filterList
}

func (s *service) GetDoctorTimeSlotsPerDay(setting doctor.ScheduleSettings, targetDate time.Time, bookedTimeSlots []doctor.TimeSlot) []doctor.TimeSlot {
	weekDay := targetDate.Weekday()
	amStartTimeOffset := 0
	amEndTimeOffset := 0
	pmStartTimeOffset := 0
	pmEndTimeOffset := 0
	if weekDay == time.Sunday {
		amStartTimeOffset = setting.SundayAmStartTimeOffset
		amEndTimeOffset = setting.SundayAmEndTimeOffset
		pmStartTimeOffset = setting.SundayPmStartTimeOffset
		pmEndTimeOffset = setting.SundayPmEndTimeOffset
	} else if weekDay == time.Monday {
		amStartTimeOffset = setting.MondayAmStartTimeOffset
		amEndTimeOffset = setting.MondayAmEndTimeOffset
		pmStartTimeOffset = setting.MondayPmStartTimeOffset
		pmEndTimeOffset = setting.MondayPmEndTimeOffset
	} else if weekDay == time.Tuesday {
		amStartTimeOffset = setting.TuesdayAmStartTimeOffset
		amEndTimeOffset = setting.TuesdayAmEndTimeOffset
		pmStartTimeOffset = setting.TuesdayPmStartTimeOffset
		pmEndTimeOffset = setting.TuesdayPmEndTimeOffset
	} else if weekDay == time.Wednesday {
		amStartTimeOffset = setting.WednesdayAmStartTimeOffset
		amEndTimeOffset = setting.WednesdayAmEndTimeOffset
		pmStartTimeOffset = setting.WednesdayPmStartTimeOffset
		pmEndTimeOffset = setting.WednesdayPmEndTimeOffset
	} else if weekDay == time.Thursday {
		amStartTimeOffset = setting.ThursdayAmStartTimeOffset
		amEndTimeOffset = setting.ThursdayAmEndTimeOffset
		pmStartTimeOffset = setting.ThursdayPmStartTimeOffset
		pmEndTimeOffset = setting.ThursdayPmEndTimeOffset
	} else if weekDay == time.Friday {
		amStartTimeOffset = setting.FridayAmStartTimeOffset
		amEndTimeOffset = setting.FridayAmEndTimeOffset
		pmStartTimeOffset = setting.FridayPmStartTimeOffset
		pmEndTimeOffset = setting.FridayPmEndTimeOffset
	} else if weekDay == time.Saturday {
		amStartTimeOffset = setting.SaturdayAmStartTimeOffset
		amEndTimeOffset = setting.SaturdayAmEndTimeOffset
		pmStartTimeOffset = setting.SaturdayPmStartTimeOffset
		pmEndTimeOffset = setting.SaturdayPmEndTimeOffset
	}

	currentOffSet := 0
	if time.Now().UTC().Day() == targetDate.Day() {
		currentOffSet = targetDate.Hour()*60 + targetDate.Minute()
	}

	timeSlots := make([]doctor.TimeSlot, 0)
	for i := amStartTimeOffset; i <= amEndTimeOffset+amStartTimeOffset; i += setting.DurationPerSlot {
		if i < currentOffSet {
			continue
		}
		timeSlot := doctor.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
		availableNumber := setting.NumberPerSlot
		if numberOfBooked >= timeSlot.AvailableSlotsNumber {
			availableNumber = 0
		} else {
			availableNumber = timeSlot.AvailableSlotsNumber - numberOfBooked
		}
		timeSlot.AvailableSlotsNumber = availableNumber
		timeSlots = append(timeSlots, timeSlot)
	}
	for i := pmStartTimeOffset + amStartTimeOffset; i <= pmEndTimeOffset+pmStartTimeOffset; i += setting.DurationPerSlot {
		if i < currentOffSet {
			continue
		}
		timeSlot := doctor.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
		availableNumber := setting.NumberPerSlot
		if numberOfBooked >= timeSlot.AvailableSlotsNumber {
			availableNumber = 0
		} else {
			availableNumber = timeSlot.AvailableSlotsNumber - numberOfBooked
		}
		timeSlot.AvailableSlotsNumber = availableNumber
		timeSlots = append(timeSlots, timeSlot)
	}
	return timeSlots
}

func getBookNumberOfTimeSlot(currentOffset int, duration int, bookedTimeSlots []doctor.TimeSlot) int {
	bookedNumber := 0
	for _, ts := range bookedTimeSlots {
		if ts.Offset <= currentOffset && ts.Offset > currentOffset-duration {
			bookedNumber = bookedNumber + 1
		}
	}
	return bookedNumber
}
