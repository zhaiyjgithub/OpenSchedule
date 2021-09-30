package router
const (
	AnyHealthService = "/AnyHealth"
)

const (
	Doctor = AnyHealthService + "/Doctor"
	SearchDoctor = "/SearchDoctor"
)

const (
	ScheduleSettings = AnyHealthService + "/scheduleSettings"
	SetScheduleSettings = "/SetScheduleSettings"
)
