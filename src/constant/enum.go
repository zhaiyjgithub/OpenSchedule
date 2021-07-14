package constant

type AppointmentType int
const (
	All AppointmentType = 0
	InClinic AppointmentType = 1
	Virtual AppointmentType = 2
)

type Gender string
const (
	Unspecified Gender = ""
	Female Gender = "F"
	Male Gender = "M"
)