package scheduleService

import (
	"OpenSchedule/constant"
	"OpenSchedule/dao/schedule"
	"OpenSchedule/database"
	"OpenSchedule/model/doctor"
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
