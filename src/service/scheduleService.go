package service

import "OpenSchedule/src/model/doctor"

type ScheduleService interface {
	SetScheduleSettings (settings *doctor.ScheduleSettings) error
}
