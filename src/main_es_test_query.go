package main

import (
	"OpenSchedule/src/constant"
	"OpenSchedule/src/database"
	"OpenSchedule/src/service"
	"github.com/olivere/elastic/v7"
)

var elasticSearchEngine *elastic.Client

func main()  {
	database.SetupElasticSearchEngine()

	s := service.NewDoctorService()
	keyword := "Richard"
	isInClinicEnable := true
	isVirtualEnable:= true
	appointmentType:= constant.InClinic
	nextAvailableDate:= "2021-07-05T14:36:41Z"
	city := ""
	specialty := ""
	lat := 40.747898
	lon := -73.324025
	gender := constant.Female
	page := 1
	pageSize := 20

	s.SearchDoctor(keyword,
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


