package schedule

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/dao/schedule"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"errors"
	"time"
)

type Service interface {
	SetScheduleSettings (settings *doctor.ScheduleSettings) error
	GetScheduleSettings(npi int64) *doctor.ScheduleSettings
	AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error
	DeleteClosedDate(id int) error
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

func (s *service) SyncCertainDoctorScheduleNextAvailableDateToES(setting *doctor.ScheduleSettings) error  {
	//get the next available date
	//begin to update the es for certain doctor
	currentTime := time.Now().UTC()
	nextAvailableDateInClinic := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, setting)
	if len(nextAvailableDateInClinic) == 0 {
		return errors.New("calc nextAvailable Date in clinic failed")
	}
	nextAvailableDateVirtual := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, setting)
	if len(nextAvailableDateInClinic) == 0 {
		return errors.New("calc nextAvailable date virtual failed")
	}
	err := s.dao.SyncCertainDoctorNextAvailableDateToES(setting.Npi, nextAvailableDateInClinic, nextAvailableDateVirtual)
	return err
}

func (s *service) AddClosedDate(closeDateSettings *doctor.ClosedDateSettings) error  {
	return s.dao.AddClosedDate(closeDateSettings)
}

func (s *service) DeleteClosedDate(id int) error {
	return s.dao.DeleteClosedDate(id)
}

func (s *service) GetClosedDate(npi int64) []doctor.ClosedDateSettings {
	return s.dao.GetClosedDate(npi)
}
