// +build

package main

import (
	"OpenSchedule/src/controller/doctor"
	"OpenSchedule/src/controller/job"
	"OpenSchedule/src/database"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service"
	"OpenSchedule/src/service/schedule"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main()  {
	fmt.Println("Hello, AnyHealth.")
	database.SetupElasticSearchEngine()

	doOnceSyncSettings()
	addJobWorker()

	app := iris.New()
	configureAnyHealthService(app)
	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M


}

func doOnceSyncSettings()  {
	doctorService := service.NewDoctorService()
	scheduleService := schedule.NewService()
	j := job.NewJob()
	j.RegisterService(doctorService, scheduleService)
	j.SyncDefaultScheduleSettingsToAllDoctor()
}

func addJobWorker()  {
	doctorService := service.NewDoctorService()
	scheduleService := schedule.NewService()
	j := job.NewJob()
	j.RegisterService(doctorService, scheduleService)
	j.SyncDoctorsNextAvailableDate()
}

func configureAnyHealthService(app *iris.Application)  {
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
	mvc.Configure(app.Party(router.ScheduleSettings), configureScheduleSettingsMVC)
}

func configureDoctorMVC(app *mvc.Application)  {
	doctorService := service.NewDoctorService()
	app.Register(doctorService)
	app.Handle(new(doctor.FindDoctorController))
}

func configureScheduleSettingsMVC(app *mvc.Application)  {
	scheduleService := schedule.NewService()
	app.Register(scheduleService)
	app.Handle(new(doctor.ScheduleController))
}