// +build

package main

import (
	"OpenSchedule/src/controller/doctor"
	"OpenSchedule/src/database"
	"OpenSchedule/src/router"
	"OpenSchedule/src/service"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)


func main()  {
	fmt.Println("Version: 1.0")
	fmt.Println("Hell, AnyHealth.")
	database.SetupElasticSearchEngine()

	app := iris.New()
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M

}

func configureAnyHealthService(app *iris.Application)  {
	mvc.Configure(app.Party(router.Doctor), configureDoctorMVC)
}

func configureDoctorMVC(app *mvc.Application)  {
	doctorService := service.NewDoctorService()
	app.Register(doctorService)
	app.Handle(new(doctor.FindDoctorController))
}