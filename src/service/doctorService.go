package service

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/dao"
	"OpenSchedule/src/database"
)

type DoctorService interface {
	SearchDoctor(keyword string,
		isInClinicEnable bool,
		isVirtualEnable bool,
		appointmentType constant.AppointmentType,
		nextAvailableDate string,
		city string,
		specialty string,
		lat float64,
		lon float64,
		gender constant.Gender,
		page int,
		pageSize int)
}

type newDoctorService struct {
	dao *dao.DoctorDao
}

func NewDoctorService() DoctorService  {
	return &newDoctorService{dao: dao.NewDoctorDao(database.GetElasticSearchEngine())}
}

func (s *newDoctorService) SearchDoctor(keyword string,
	isInClinicEnable bool,
	isVirtualEnable bool,
	appointmentType constant.AppointmentType,
	nextAvailableDate string,
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int)  {
	s.dao.SearchDoctor(keyword,
		isInClinicEnable,
		isVirtualEnable,
		appointmentType,
		nextAvailableDate,
		city,
		specialty,
		lat,
		lon,
		gender,
		page,
		pageSize)
}