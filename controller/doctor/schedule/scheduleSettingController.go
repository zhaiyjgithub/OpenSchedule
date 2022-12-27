package schedule

import (
	"OpenSchedule/model/doctor"
	"OpenSchedule/response"
	"OpenSchedule/router"
	"OpenSchedule/service/scheduleService"
	"OpenSchedule/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type Controller struct {
	Ctx             iris.Context
	ScheduleService scheduleService.Service
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, router.SetScheduleSettings, "SetScheduleSettings")
	b.Handle(iris.MethodPost, router.GetScheduleSettings, "GetScheduleSettings")
	b.Handle(iris.MethodPost, router.AddClosedDateSettings, "AddClosedDateSettings")
	b.Handle(iris.MethodPost, router.DeleteClosedDateSettings, "DeleteClosedDateSettings")
	b.Handle(iris.MethodPost, router.GetClosedDateSettings, "GetClosedDateSettings")
	b.Handle(iris.MethodPost, router.AddAppointment, "AddAppointment")
	b.Handle(iris.MethodPost, router.GetAppointmentByPage, "GetAppointmentByPage")
}

func (c *Controller) SetScheduleSettings() {
	var p doctor.ScheduleSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	err := c.ScheduleService.SetScheduleSettings(p)
	if err != nil {
		err = c.ScheduleService.SyncCertainDoctorScheduleNextAvailableDateToES(p)
	}
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetScheduleSettings() {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	settings := c.ScheduleService.GetScheduleSettings(p.Npi)
	response.Success(c.Ctx, response.Successful, settings)
}

func (c *Controller) AddClosedDateSettings() {
	var p doctor.ClosedDateSettings
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.AddClosedDate(p); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) DeleteClosedDateSettings() {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
		Sid int   `json:"sid"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	if err := c.ScheduleService.DeleteClosedDate(p.Npi, p.Sid); err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetClosedDateSettings() {
	type Param struct {
		Npi int64 `json:"npi" validate:"gt=0"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	list := c.ScheduleService.GetClosedDate(p.Npi)
	response.Success(c.Ctx, response.Successful, list)
}

func (c *Controller) AddAppointment() {
	type Param struct {
		DoctorID               int       `json:"doctorId" validate:"required"`
		Npi                    int64     `json:"npi" validate:"required"`
		AppointmentType        int       `json:"appointmentType"`
		AppointmentDate        string `json:"appointmentDate" validate:"required"`
		AppointmentStatus      int       `json:"appointmentStatus" validate:"required"`
		Memo                   string    `json:"memo"`
		Offset                 int       `json:"offset"`
		PatientID              int       `json:"patientId" validate:"required"`
		LegalGuardianPatientID int       `json:"legalGuardianPatientId"`
		FirstName              string    `json:"firstName" validate:"required"`
		LastName               string    `json:"lastName" validate:"required"`
		Dob                    string    `json:"dob" validate:"required"`
		Gender                 string    `json:"gender"`
		Email                  string    `json:"email" validate:"required"`
		Phone                  string    `json:"phone" validate:"required"`
		Insurance              int       `json:"insurance"`
		VisitReason            string    `json:"visitReason"`
		IsNewPatient           bool      `json:"isNewPatient"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	createdAt := time.Now().UTC()
	aptDate, err := time.Parse(time.RFC3339, p.AppointmentDate)
	if err != nil {
		fmt.Println(err.Error())
	}
	appt := doctor.Appointment{
		DoctorID:               p.DoctorID,
		Npi:                    p.Npi,
		AppointmentType:        p.AppointmentType,
		AppointmentDate:        aptDate,
		AppointmentStatus:      p.AppointmentStatus,
		Memo:                   p.Memo,
		Offset:                 p.Offset,
		PatientID:              p.PatientID,
		LegalGuardianPatientID: p.LegalGuardianPatientID,
		FirstName:              p.FirstName,
		LastName:               p.LastName,
		Dob:                    p.Dob,
		Gender:                 p.Gender,
		Email:                  p.Email,
		Phone:                  p.Phone,
		Insurance:              p.Insurance,
		VisitReason:            p.VisitReason,
		IsNewPatient:           p.IsNewPatient,
		CreatedDate: createdAt,
		UpdatedAt: createdAt,
	}
	err = c.ScheduleService.AddAppointment(appt)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetAppointmentByPage() {
	type Param struct {
		PatientID int `validate:"gt=1"`
		Page int `validate:"gte=1"`
		PageSize int `validate:"gte=5"`
	}
	var p Param
	if err := utils.ValidateParam(c.Ctx, &p); err != nil {
		return
	}
	appts, err := c.ScheduleService.GetAppointment(p.PatientID, p.Page, p.PageSize)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	} else {
		response.Success(c.Ctx, response.Successful, appts)
	}
}
