package schedule

import (
	"OpenSchedule/src/dao/schedule"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
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
}

func (s *service) SetScheduleSettings(setting *doctor.ScheduleSettings) error {
	return s.dao.SetScheduleSettings(setting)
}

func (s *service) GetScheduleSettings(npi int64) *doctor.ScheduleSettings {
	return s.dao.GetScheduleSettings(npi)
}