package doctor

import (
	"OpenSchedule/src/constant"
	doctor2 "OpenSchedule/src/dao/doctor"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/model/viewModel"
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
		pageSize int,
		sortType constant.SortType, distance int) []*viewModel.DoctorInfo
	GetDoctorByPage(page int, pageSize int) []*doctor.Doctor
}

type doctorService struct {
	dao *doctor2.Dao
}

func NewDoctorService() DoctorService {
	return &doctorService{dao: doctor2.NewDoctorDao(database.GetElasticSearchEngine(), database.GetMySqlEngine())}
}

func (s *doctorService) SearchDoctor(keyword string,
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
	pageSize int, sortType constant.SortType, distance int) []*viewModel.DoctorInfo {
	return s.dao.SearchDoctor(keyword,
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
		pageSize, sortType, distance)
}

func (s *doctorService) GetDoctorByPage(page int, pageSize int) []*doctor.Doctor {
	return s.dao.GetDoctorByPage(page, pageSize)
}