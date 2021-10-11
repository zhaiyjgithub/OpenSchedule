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
	isOk, nextAvailableDateInClinic := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, setting)
	if !isOk {
		return errors.New("calc nextAvailable Date in clinic failed")
	}
	isOk, nextAvailableDateVirtual := s.dao.CalcNextAvailableDate(currentTime, constant.InClinic, setting)
	if !isOk {
		return errors.New("calc nextAvailable date virtual failed")
	}
	err := s.dao.SyncCertainDoctorNextAvailableDateToES(setting.Npi, nextAvailableDateInClinic, nextAvailableDateVirtual)
	return err
}