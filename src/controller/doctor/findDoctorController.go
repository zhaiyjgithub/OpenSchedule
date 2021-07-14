package doctor

import (
	"OpenSchedule/src/router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type FindDoctorController struct {
	Ctx iris.Context
}

func (c *FindDoctorController) AnyLogout(b mvc.BeforeActivation)  {
	b.Handle(http.MethodPost, router.SearchDoctor, "SearchDoctor")
}

func (c *FindDoctorController) SearchDoctor()  {

}
