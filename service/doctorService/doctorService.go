package doctorService

import (
	"OpenSchedule/constant"
	doctor2 "OpenSchedule/dao/doctorDao"
	"OpenSchedule/database"
	"OpenSchedule/model/doctorModel"
	"OpenSchedule/model/viewModel"
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
	GetDoctorByPage(page int, pageSize int) []*doctorModel.Doctor
	IsExist(npi int64) bool
	GetDoctor(npi int64) doctorModel.Doctor
	SaveDoctor(doc *doctorModel.Doctor) error
	GetClinic(npi int64) []doctorModel.Clinicals
	GetAwards(npi int64) []doctorModel.Awards
	GetCertification(npi int64) []doctorModel.Certifications
	GetEducation(npi int64) []doctorModel.Educations
	GetInsurance(npi int64) string
	GetDoctorDetail(npi int64) viewModel.DoctorDetailInfo
	GetDoctorUser(email string, password string) (doctorModel.DoctorUser, error)
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

func (s *doctorService) GetDoctorByPage(page int, pageSize int) []*doctorModel.Doctor {
	return s.dao.GetDoctorByPage(page, pageSize)
}

func (s *doctorService) IsExist(npi int64) bool {
	return s.dao.IsExist(npi)
}

func (s *doctorService) GetDoctor(npi int64) doctorModel.Doctor {
	return s.dao.GetDoctor(npi)
}

func (s *doctorService) SaveDoctor(doc *doctorModel.Doctor) error {
	return s.dao.SaveDoctor(doc)
}

func (s *doctorService) GetClinic(npi int64) []doctorModel.Clinicals {
	return s.dao.GetClinic(npi)
}

func (s *doctorService) GetAwards(npi int64) []doctorModel.Awards {
	return s.dao.GetAwards(npi)
}

func (s *doctorService) GetCertification(npi int64) []doctorModel.Certifications {
	return s.dao.GetCertification(npi)
}

func (s *doctorService) GetEducation(npi int64) []doctorModel.Educations {
	return s.dao.GetEducation(npi)
}

func (s *doctorService) GetInsurance(npi int64) string {
	return s.dao.GetInsurance(npi)
}

func (s *doctorService) GetDoctorDetail(npi int64) viewModel.DoctorDetailInfo {
	return s.dao.GetDoctorDetail(npi)
}

func (s *doctorService) GetDoctorUser(email string, password string) (doctorModel.DoctorUser, error) {
	return s.dao.GetUser(email, password)
}
