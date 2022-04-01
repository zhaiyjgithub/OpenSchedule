package doctorService

import (
	"OpenSchedule/src/constant"
	doctor2 "OpenSchedule/src/dao/doctor"
	"OpenSchedule/src/database"
	"OpenSchedule/src/model/doctor"
	"OpenSchedule/src/model/viewModel"
)

type Service interface {
	SearchDoctor(keyword string,
		appointmentType constant.AppointmentType,
		city string,
		specialty string,
		lat float64,
		lon float64,
		gender constant.Gender,
		page int,
		pageSize int,
		sortType constant.SortByType, distance int) (int64, []*viewModel.DoctorInfo)
	GetDoctorByPage(page int, pageSize int) []*doctor.Doctor
	IsExist(npi int64) bool
	GetDoctor(npi int64) doctor.Doctor
	SaveDoctor(doc *doctor.Doctor) error
	GetClinic(npi int64) []doctor.Clinicals
	GetAwards(npi int64) []doctor.Awards
	GetCertification(npi int64) []doctor.Certifications
	GetEducation(npi int64) []doctor.Educations
	GetInsurance(npi int64) []doctor.Insurances
	GetDoctorDetail(npi int64) viewModel.DoctorDetailInfo

}

type doctorService struct {
	dao *doctor2.Dao
}

func NewService() Service {
	return &doctorService{dao: doctor2.NewDoctorDao(database.GetElasticSearchEngine(), database.GetMySqlEngine())}
}

func (s *doctorService) SearchDoctor(keyword string,
	appointmentType constant.AppointmentType,
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int, sortType constant.SortByType, distance int) (int64, []*viewModel.DoctorInfo) {
	return s.dao.SearchDoctor(
		keyword,
		appointmentType,
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

func (s *doctorService) IsExist(npi int64) bool {
	return s.dao.IsExist(npi)
}

func (s *doctorService) GetDoctor(npi int64) doctor.Doctor  {
	return s.dao.GetDoctor(npi)
}

func (s *doctorService) SaveDoctor(doc *doctor.Doctor) error  {
	return s.dao.SaveDoctor(doc)
}

func (s *doctorService) GetClinic(npi int64) []doctor.Clinicals {
	return s.dao.GetClinic(npi)
}

func (s *doctorService) GetAwards(npi int64) []doctor.Awards {
	return s.dao.GetAwards(npi)
}

func (s *doctorService) GetCertification(npi int64) []doctor.Certifications {
	return s.dao.GetCertification(npi)
}

func (s *doctorService) GetEducation(npi int64) []doctor.Educations {
	return s.dao.GetEducation(npi)
}

func (s *doctorService) GetInsurance(npi int64) []doctor.Insurances {
	return s.dao.GetInsurance(npi)
}

func (s *doctorService) GetDoctorDetail(npi int64) viewModel.DoctorDetailInfo  {
	return s.dao.GetDoctorDetail(npi)
}