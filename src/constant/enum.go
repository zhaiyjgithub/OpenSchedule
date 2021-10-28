package constant

import "time"

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

type SortType int

const  (
	ByDefault SortType = 0
	ByDistance SortType = 1
)

var DefaultTimeStamp = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

const (
	YYYY_MM_DD_HH_mm_SS = "2006-01-02 15:00:00"
)