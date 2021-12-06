package router
const (
	AnyHealthService = "/AnyHealth"
)

const (
	Doctor = AnyHealthService + "/Doctor"
	SearchDoctor = "/SearchDoctor"
	GetDoctor = "/GetDoctor"
	SaveDoctor = "/SaveDoctor"
)

const (
	ScheduleSettings = AnyHealthService + "/Schedule"
	SetScheduleSettings = "/SetScheduleSettings"
	GetScheduleSettings = "/GetScheduleSettings"
	AddClosedDateSettings = "/AddClosedDateSettings"
	DeleteClosedDateSettings = "/DeleteClosedDateSettings"
	GetClosedDateSettings = "/GetClosedDateSettings"
)
