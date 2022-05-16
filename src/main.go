package main

import (
	doctor2 "OpenSchedule/src/controller/doctor/doctor"
	"OpenSchedule/src/controller/doctor/schedule"
	"OpenSchedule/src/controller/job"
	"OpenSchedule/src/controller/user"
	"OpenSchedule/src/database"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service/doctorService"
	"OpenSchedule/src/service/scheduleService"
	"OpenSchedule/src/service/userService"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	fmt.Println("Hello, AnyHealth.")
	database.SetupElasticSearchEngine()
	addJobWorker()
	app := iris.New()
	configureAnyHealthService(app)
	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M
}

func addJobWorker() {
	doctorService := doctorService.NewService()
	scheduleService := scheduleService.NewService()
	j := job.NewJob()
	j.RegisterService(doctorService, scheduleService)
	j.StartToSyncDoctorNextAvailableDateAsync()
}

func configureAnyHealthService(app *iris.Application) {
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
	mvc.Configure(app.Party(router.ScheduleSettings), configureScheduleSettingsMVC)
	mvc.Configure(app.Party(router.User), configureUerMVC)
}

func configureDoctorMVC(app *mvc.Application) {
	doctorService := doctorService.NewService()
	scheduleService := scheduleService.NewService()
	app.Register(doctorService, scheduleService)
	app.Handle(new(doctor2.Controller))
}

func configureScheduleSettingsMVC(app *mvc.Application) {
	scheduleService := scheduleService.NewService()
	app.Register(scheduleService)
	app.Handle(new(schedule.ScheduleController))
}

func configureUerMVC(app *mvc.Application) {
	usersService := userService.NewService()
	app.Register(usersService)
	app.Handle(new(user.Controller))
}
