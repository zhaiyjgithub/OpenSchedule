package router

const (
	AnyHealthService = "/AnyHealth"
)

const (
	Doctor              = AnyHealthService + "/Doctor"
	SearchDoctor        = "/SearchDoctor"
	GetDoctor           = "/GetDoctor"
	SaveDoctor          = "/SaveDoctor"
	GetTimeSlots        = "/GetTimeSlots"
	GetDoctorDetailInfo = "/GetDoctorDetailInfo"
)

const (
	ScheduleSettings         = AnyHealthService + "/Schedule"
	SetScheduleSettings      = "/SetScheduleSettings"
	GetScheduleSettings      = "/GetScheduleSettings"
	AddClosedDateSettings    = "/AddClosedDateSettings"
	DeleteClosedDateSettings = "/DeleteClosedDateSettings"
	GetClosedDateSettings    = "/GetClosedDateSettings"
	AddAppointment           = "/AddAppointment"
	GetAppointmentByPage = "/GetAppointmentByPage"
)

const (
	User               = AnyHealthService + "/User"
	CreateUser         = "/CreateUser"
	GetUserByEmail     = "/GetUserByEmail"
	Login              = "/Login"
	CreateSubUser      = "/CreateSubUser"
	GetSubUsers        = "/GetSubUsers"
	UpdateSubUserPhone = "/UpdateSubUserPhone"
	UpdateUserProfile = "/UpdateUserProfile"
	GetUserByID = "/GetUserByID"
	GetUserInsurance = "/GetUserInsurance"
	UpdateUserInsurance = "/UpdateUserInsurance"
)
