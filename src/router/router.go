package router
const (
	AnyHealthService = "/AnyHealth"
)

const (
	Doctor = AnyHealthService + "/Doctor"
	SearchDoctor = "/SearchDoctor"
)

const (
	scheduleSettings = AnyHealthService + "/scheduleSettings"
	SetScheduleSettings = "/SetScheduleSettings"
)
