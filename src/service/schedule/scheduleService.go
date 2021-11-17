package schedule

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/dao/schedule"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"time"
)

type Service interface {
	SetScheduleSettings (settings *doctor.ScheduleSettings) error
	GetScheduleSettings(npi int64) *doctor.ScheduleSettings
	AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error
	DeleteClosedDate(npi int64, id int) error
	GetClosedDate(npi int64) []doctor.ClosedDateSettings
}

func NewService() Service {
	return &service{dao: schedule.NewDao(database.GetMySqlEngine(), database.GetElasticSearchEngine())}
}

type service struct {
	dao *schedule.Dao
}

func (s *service) SetScheduleSettings(setting *doctor.ScheduleSettings) error {
	err := s.dao.SetScheduleSettings(setting)
	if err != nil {
		return err
	}
	//begin to sync the certain doctor next available date.
	err = s.SyncCertainDoctorScheduleNextAvailableDateToES(setting)
	return err
}

func (s *service) GetScheduleSettings(npi int64) *doctor.ScheduleSettings {
	return s.dao.GetScheduleSettings(npi)
}

func (s *service) SyncCertainDoctorScheduleNextAvailableDateToES(settings *doctor.ScheduleSettings) error  {
	currentTime := time.Now().UTC()
	nextAvailableDateInClinic := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, settings)
	nextAvailableDateVirtual := s.dao.CalcNextAvailableDate(currentTime, constant.Virtual, settings)

	isInClinicBookEnable := nextAvailableDateInClinic != constant.InvalidDateTime
	isVirtualBookEnable := nextAvailableDateVirtual != constant.InvalidDateTime
	isOnlineScheduleEnable := isInClinicBookEnable || isVirtualBookEnable

	err := s.dao.SyncCertainDoctorNextAvailableDateToES(settings.Npi,
		isOnlineScheduleEnable, isInClinicBookEnable, isVirtualBookEnable,
		nextAvailableDateInClinic, nextAvailableDateVirtual)
	return err
}

func (s *service) AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error  {
	return s.dao.AddClosedDate(closeDateSettings)
}

func (s *service) DeleteClosedDate(npi int64, id int) error {
	return s.dao.DeleteClosedDateByID(npi, id)
}

func (s *service) GetClosedDate(npi int64) []doctor.ClosedDateSettings {
	return s.dao.GetClosedDate(npi)
}
