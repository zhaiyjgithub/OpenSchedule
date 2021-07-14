package main

import (
	"OpenSchedule/src/controller/doctor"
	"OpenSchedule/src/database"
	"OpenSchedule/src/router"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)


func main()  {
	fmt.Println("hello go")
	database.SetupElasticSearchEngine()

	//fmt.Println(time.Now().UTC().Format(time.RFC3339))
	app := iris.New()
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
}

func configureDoctorMVC(app *mvc.Application)  {
	app.Handle(new(doctor.FindDoctorController))
}