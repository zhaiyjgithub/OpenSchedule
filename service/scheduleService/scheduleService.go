package scheduleService

import (
	"OpenSchedule/constant"
	"OpenSchedule/dao/scheduleDao"
	"OpenSchedule/database"
	"OpenSchedule/model/doctorModel"
	"OpenSchedule/model/viewModel"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type Service interface {
	SetScheduleSettings(settings doctorModel.ScheduleSettings) error
	GetScheduleSettings(npi int64) doctorModel.ScheduleSettings
	AddClosedDate(closeDateSettings doctorModel.ClosedDateSettings) error
	DeleteClosedDate(npi int64, id int) error
	GetClosedDate(npi int64) []doctorModel.ClosedDateSettings
	SyncCertainDoctorScheduleNextAvailableDateToES(settings doctorModel.ScheduleSettings) error
	SyncMultiDoctorsScheduleNextAvailableDateToES(doctors []*doctorModel.Doctor) error
	IsExist(npi int64) bool
	SyncDoctorToES(doctor *doctorModel.Doctor) error
	AddAppointment(appointment doctorModel.Appointment) error
	GetAppointmentByRange(
		npi int64,
		appointmentStatus constant.AppointmentStatus,
		startDate time.Time,
		endDate time.Time,
	) []*doctorModel.Appointment
	GetAppointmentsByDate(
		npi int64,
		startDate time.Time,
		endDate time.Time,
	) ([]doctorModel.Appointment, error)
	GetAppointmentsByRange(
		npi []int64,
		appointmentStatus constant.AppointmentStatus,
		startDate time.Time,
		endDate time.Time,
	) []*doctorModel.Appointment
	GetSettingsByNpiList(npiList []int64) []doctorModel.ScheduleSettings
	GetClosedDateByRange(npi []int64, from time.Time, to time.Time) []doctorModel.ClosedDateSettings
	GetBookedAppointmentsTimeSlotsByNpiList(npiList []int64, startDate time.Time, endTime time.Time) map[int64]map[string][]doctorModel.TimeSlot
	GetClosedDateByNpiList(npiList []int64, startDate time.Time, endDate time.Time) map[int64][]doctorModel.ClosedDateSettings
	GetScheduleSettingByNpiList(npiList []int64) map[int64]doctorModel.ScheduleSettings
	GetDoctorTimeSlotsByDate(setting doctorModel.ScheduleSettings, startDate time.Time, endDate time.Time, bookedTimeSlots map[string][]doctorModel.TimeSlot,
		closeDateSetting []doctorModel.ClosedDateSettings) []viewModel.TimeSlotPerDay
	GetOneDayTimeSlotByNpi(npi int64, targetDate time.Time) ([]doctorModel.TimeSlot, error)
	CheckTimeSlotIsAvailable(npi int64, targetDateTime time.Time) (bool, error)
	GetAppointmentInfo(patientID int, page int, pageSize int) ([]viewModel.AppointmentInfo, error)
}

func NewService() Service {
	return &service{dao: scheduleDao.NewDao(database.GetMySqlEngine(), database.GetElasticSearchEngine())}
}

type service struct {
	dao *scheduleDao.Dao
}

func (s *service) SetScheduleSettings(setting doctorModel.ScheduleSettings) error {
	return s.dao.SetScheduleSettings(setting)
}

func (s *service) GetScheduleSettings(npi int64) doctorModel.ScheduleSettings {
	return s.dao.GetScheduleSettings(npi)
}

func (s *service) SyncCertainDoctorScheduleNextAvailableDateToES(settings doctorModel.ScheduleSettings) error {
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

func (s *service) SyncMultiDoctorsScheduleNextAvailableDateToES(doctors []*doctorModel.Doctor) error {
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

func (s *service) AddClosedDate(closeDateSettings doctorModel.ClosedDateSettings) error {
	return s.dao.AddClosedDate(closeDateSettings)
}

func (s *service) DeleteClosedDate(npi int64, id int) error {
	return s.dao.DeleteClosedDateByID(npi, id)
}

func (s *service) GetClosedDate(npi int64) []doctorModel.ClosedDateSettings {
	return s.dao.GetClosedDate(npi)
}

func (s *service) IsExist(npi int64) bool {
	return s.dao.IsExist(npi)
}

func (s *service) SyncDoctorToES(doctor *doctorModel.Doctor) error {
	return s.dao.SyncDoctorToES(doctor)
}

func (s *service) AddAppointment(appointment doctorModel.Appointment) error {
	ok, err := s.CheckTimeSlotIsAvailable(appointment.Npi, appointment.AppointmentDate)
	if err != nil {
		return err
	}
	if ok {
		return s.dao.AddAppointment(appointment)
	} else {
		return errors.New("this appointment is unavailable")
	}
}

func (s *service) GetAppointmentByRange(
	npi int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctorModel.Appointment {
	return s.dao.GetAppointmentByRange(npi, appointmentStatus, startDate, endDate)
}

func (s *service) GetAppointmentsByRange(
	npi []int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctorModel.Appointment {
	return s.dao.GetAppointmentsByRange(npi, appointmentStatus, startDate, endDate)
}

func (s *service) GetAppointmentsByDate(
	npi int64,
	startDate time.Time,
	endDate time.Time,
) ([]doctorModel.Appointment, error) {
	return s.dao.GetAppointmentByDate(npi, startDate, endDate)
}

func (s *service) GetSettingsByNpiList(npiList []int64) []doctorModel.ScheduleSettings {
	return s.dao.GetSettingsByNpiList(npiList)
}

func (s *service) GetClosedDateByRange(npi []int64, from time.Time, to time.Time) []doctorModel.ClosedDateSettings {
	return s.dao.GetClosedDateByRange(npi, from, to)
}

// todo: Refactor appoint type

func (s *service) GetBookedAppointmentsTimeSlotsByNpiList(npiList []int64, startDate time.Time, endTime time.Time) map[int64]map[string][]doctorModel.TimeSlot {
	appointments := s.GetAppointmentsByRange(npiList, constant.Requested, startDate, endTime)
	bTimeSlots := make(map[int64]map[string][]doctorModel.TimeSlot)
	for _, appt := range appointments {
		bookedTimeSlotsPerNpi, ok := bTimeSlots[appt.Npi]
		offset := appt.AppointmentDate.Hour()*60 + appt.AppointmentDate.Minute()
		dateKey := fmt.Sprintf("%d-%d-%d", appt.AppointmentDate.Year(), appt.AppointmentDate.Month(), appt.AppointmentDate.Day())
		if !ok {
			bookedTimeSlotsPerNpi = make(map[string][]doctorModel.TimeSlot)
			bookedTimeSlotsPerNpi[dateKey] = []doctorModel.TimeSlot{{Offset: offset, AvailableSlotsNumber: 1}}
		} else {
			bookedTimeSlotsPerNpi[dateKey] = append(bookedTimeSlotsPerNpi[dateKey], doctorModel.TimeSlot{Offset: offset, AvailableSlotsNumber: 1})
		}
		bTimeSlots[appt.Npi] = bookedTimeSlotsPerNpi
	}
	return bTimeSlots
}

func (s *service) GetClosedDateByNpiList(npiList []int64, startDate time.Time, endDate time.Time) map[int64][]doctorModel.ClosedDateSettings {
	closeDateSettings := s.GetClosedDateByRange(npiList, startDate, endDate)
	settingMap := make(map[int64][]doctorModel.ClosedDateSettings)
	for _, setting := range closeDateSettings {
		settingMap[setting.Npi] = append(settingMap[setting.Npi], setting)
	}
	return settingMap
}

func (s *service) GetScheduleSettingByNpiList(npiList []int64) map[int64]doctorModel.ScheduleSettings {
	list := s.GetSettingsByNpiList(npiList)
	settingMap := make(map[int64]doctorModel.ScheduleSettings)
	for _, setting := range list {
		settingMap[setting.Npi] = setting
	}
	return settingMap
}

func (s *service) GetDoctorTimeSlotsByDate(setting doctorModel.ScheduleSettings, startDate time.Time, endDate time.Time, bookedTimeSlots map[string][]doctorModel.TimeSlot,
	closeDateSetting []doctorModel.ClosedDateSettings) []viewModel.TimeSlotPerDay {
	timeSlots := make([]viewModel.TimeSlotPerDay, 0)
	dateRange := int(endDate.Sub(startDate).Hours() / 24)
	for i := 0; i < dateRange; i++ {
		targetDate := startDate.AddDate(0, 0, i)
		dateKey := fmt.Sprintf("%d-%d-%d", targetDate.Year(), targetDate.Month(), targetDate.Day())
		bookedTimeSlots, ok := bookedTimeSlots[dateKey]
		timeSlotsPerDay := make([]doctorModel.TimeSlot, 0)
		if ok {
			timeSlotsPerDay = s.GetDoctorTimeSlotsPerDay(setting, targetDate, bookedTimeSlots)
		} else {
			timeSlotsPerDay = s.GetDoctorTimeSlotsPerDay(setting, targetDate, make([]doctorModel.TimeSlot, 0))
		}
		targetClosetDate := s.getClosedDateByDate(closeDateSetting, targetDate)
		filterTimeSlotsPerDay := s.filterTimeSlotsByClosedDate(targetDate, timeSlotsPerDay, targetClosetDate)
		timeSlots = append(timeSlots, viewModel.TimeSlotPerDay{Date: targetDate, TimeSlots: filterTimeSlotsPerDay})
	}
	return timeSlots
}

func (s *service) getClosedDateByDate(closedDateList []doctorModel.ClosedDateSettings, targetDate time.Time) doctorModel.ClosedDateSettings {
	var targetClosedDate doctorModel.ClosedDateSettings
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

func (s *service) filterTimeSlotsByClosedDate(targetDate time.Time, timeSlots []doctorModel.TimeSlot, closedDate doctorModel.ClosedDateSettings) []doctorModel.TimeSlot {
	if closedDate.AmStartDateTime.IsZero() && closedDate.PmStartDateTime.IsZero() {
		return timeSlots
	}
	var filterList []doctorModel.TimeSlot
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

func (s *service) GetDoctorTimeSlotsPerDay(setting doctorModel.ScheduleSettings, targetDate time.Time, bookedTimeSlots []doctorModel.TimeSlot) []doctorModel.TimeSlot {
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

	timeSlots := make([]doctorModel.TimeSlot, 0)
	for i := amStartTimeOffset; i <= amEndTimeOffset+amStartTimeOffset; i += setting.DurationPerSlot {
		if i < currentOffSet {
			continue
		}
		timeSlot := doctorModel.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookedNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
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
		timeSlot := doctorModel.TimeSlot{Offset: i, AvailableSlotsNumber: setting.NumberPerSlot}
		numberOfBooked := getBookedNumberOfTimeSlot(timeSlot.Offset, setting.DurationPerSlot, bookedTimeSlots)
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

func getBookedNumberOfTimeSlot(currentOffset int, duration int, bookedTimeSlots []doctorModel.TimeSlot) int {
	bookedNumber := 0
	for _, ts := range bookedTimeSlots {
		if ts.Offset <= currentOffset && ts.Offset > currentOffset-duration {
			bookedNumber = bookedNumber + 1
		}
	}
	return bookedNumber
}

func (s *service) GetOneDayTimeSlotByNpi(npi int64, targetDate time.Time) ([]doctorModel.TimeSlot, error) {
	l := []int64{npi}
	var timeSlots []doctorModel.TimeSlot
	setting := s.GetScheduleSettings(npi)
	if setting.Npi == 0 {
		return timeSlots, errors.New("npi not found")
	}
	endDate := targetDate
	allBookedTimeSlotsMap := s.GetBookedAppointmentsTimeSlotsByNpiList(l, targetDate, endDate)
	closedDateMap := s.GetClosedDateByNpiList(l, targetDate, endDate)
	dateKey := fmt.Sprintf("%d-%d-%d", targetDate.Year(), targetDate.Month(), targetDate.Day())
	bookedTimeSlotsForNpi, ok := allBookedTimeSlotsMap[npi]
	var bookedTimeSlotInTargetDate []doctorModel.TimeSlot
	if !ok {
		bookedTimeSlotInTargetDate = make([]doctorModel.TimeSlot, 0)
	} else {
		bookedTimeSlotInTargetDate = bookedTimeSlotsForNpi[dateKey]
	}
	closeDateList, ok := closedDateMap[npi]
	if !ok {
		closeDateList = []doctorModel.ClosedDateSettings{}
	}
	allTimeSlots := s.GetDoctorTimeSlotsPerDay(setting, targetDate, bookedTimeSlotInTargetDate)
	targetClosetDate := s.getClosedDateByDate(closeDateList, targetDate)
	timeSlots = s.filterTimeSlotsByClosedDate(targetDate, allTimeSlots, targetClosetDate)
	return timeSlots, nil
}

func (s *service) CheckTimeSlotIsAvailable(npi int64, targetDateTime time.Time) (bool, error) {
	allTimeSlots, err := s.GetOneDayTimeSlotByNpi(npi, targetDateTime)
	if err != nil {
		return false, err
	}
	offset := targetDateTime.Hour()*60 + targetDateTime.Minute()
	for idx, _slot := range allTimeSlots {
		if _slot.AvailableSlotsNumber > 0 && _slot.Offset >= offset && idx < (len(allTimeSlots)-1) {
			return true, nil
		}
	}
	return false, nil
}

func (s *service) GetAppointmentInfo(patientID int, page int, pageSize int) ([]viewModel.AppointmentInfo, error) {
	return s.dao.GetAppointmentInfo(patientID, page, pageSize)
}
