package dao

import (
	"OpenSchedule/src/constant"
	"github.com/olivere/elastic/v7"
	"time"
)

type DoctorDao struct {
	elasticSearchEngine *elastic.Client
}

/*
|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|keyword |yes  |string |   |
|isInClinicEnable |no  |bool |  default true  |
|isVirtualEnable     |no  |bool |    |
|appointmentType | yes | int | all = 0 inClinic=1 virtual = 2
|nextAvailableDate |no |date | UTC format
|city |no |string |
|specialty |no | string |
|lat | yes | float |
|lon |yes | float |
|gender |no | bool|
|page |yes | int| default index from 1
|pageSize |yes | int | default size from 50*/
func (d *DoctorDao) SearchDoctor(keyword string,
	isInClinicEnable bool,
	isVirtualEnable bool,
	appointmentType constant.AppointmentType,
	nextAvailableDate time.Time,
	city string,
	specialty string,
	lat float64,
	lon float64,
	gender constant.Gender,
	page int,
	pageSize int)  {

	bq := elastic.NewBoolQuery()

	if len(keyword) > 0 {

	}
}
