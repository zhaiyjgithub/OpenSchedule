package router
const (
	AnyHealthService = "/AnyHealth"
)

const (
	Doctor = AnyHealthService + "/Doctor"
	SearchDoctor = "/SearchDoctor"
)

const (
	ScheduleSettings = AnyHealthService + "/Schedule"
	SetScheduleSettings = "/SetScheduleSettings"
	GetScheduleSettings = "/GetScheduleSettings"
)
