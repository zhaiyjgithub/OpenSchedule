package main

import (
	doctor2 "OpenSchedule/controller/doctor/doctor"
	"OpenSchedule/controller/doctor/schedule"
	"OpenSchedule/controller/job"
	"OpenSchedule/controller/user"
	"OpenSchedule/database"
	"OpenSchedule/router"
	"OpenSchedule/service/doctorService"
	"OpenSchedule/service/scheduleService"
	"OpenSchedule/service/userService"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	fmt.Println("Hello, Zen.")
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
	app.Handle(new(schedule.Controller))
}

func configureUerMVC(app *mvc.Application) {
	usersService := userService.NewService()
	app.Register(usersService)
	app.Handle(new(user.Controller))
}
