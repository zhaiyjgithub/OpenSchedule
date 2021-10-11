package schedule

import (
	"OpenSchedule/src/dao/schedule"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"github.com/olivere/elastic/v7"
)

type Service interface {
	SetScheduleSettings (settings *doctor.ScheduleSettings) error
	GetScheduleSettings(npi int64) *doctor.ScheduleSettings
}

func NewService() Service {
	return &service{dao: schedule.NewDao(database.GetMySqlEngine())}
}

type service struct {
	dao *schedule.Dao
	esEngine *elastic.Client
}

func (s *service) SetScheduleSettings(setting *doctor.ScheduleSettings) error {
	return s.dao.SetScheduleSettings(setting)
	//begin to sync the certain doctor next available date.
}

func (s *service) GetScheduleSettings(npi int64) *doctor.ScheduleSettings {
	return s.dao.GetScheduleSettings(npi)
}

func (s *service) SyncCertainDoctorScheduleNextAvailableDateToES(setting *doctor.ScheduleSettings)  {
	//get the next available date
	//begin to update the es for certain doctor
}