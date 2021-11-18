// +build

package main

import (
	"OpenSchedule/src/controller/doctor"
	"OpenSchedule/src/controller/job"
	"OpenSchedule/src/database"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/doctorService"
	"OpenSchedule/src/service/scheduleService"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main()  {
	fmt.Println("Hello, AnyHealth.")
	database.SetupElasticSearchEngine()
	addJobWorker()
	app := iris.New()
	configureAnyHealthService(app)
	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M
}

func addJobWorker()  {
	doctorService := doctorService.NewService()
	scheduleService := scheduleService.NewService()
	j := job.NewJob()
	j.RegisterService(doctorService, scheduleService)
	j.StartToSyncDoctorNextAvailableDateAsync()
}

func configureAnyHealthService(app *iris.Application)  {
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
	mvc.Configure(app.Party(router.ScheduleSettings), configureScheduleSettingsMVC)
}

func configureDoctorMVC(app *mvc.Application)  {
	doctorService := doctorService.NewService()
	app.Register(doctorService)
	app.Handle(new(doctor.FindDoctorController))
}

func configureScheduleSettingsMVC(app *mvc.Application)  {
	scheduleService := scheduleService.NewService()
	app.Register(scheduleService)
	app.Handle(new(doctor.ScheduleController))
}