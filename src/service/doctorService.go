package service

import (
	"OpenSchedule/src/dao"
	"OpenSchedule/src/database"
)

type DoctorService interface {

}

type newDoctorService struct {
	dao *dao.DoctorDao
}

func NewDoctorService() DoctorService  {
	return &newDoctorService{dao: dao.NewDoctorDao(database.GetElasticSearchEngine())}
}
