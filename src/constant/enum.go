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
	Trans Gender = ""
	Female Gender = "F"
	Male Gender = "M"
)

type SortByType int

const  (
	ByDefault  SortByType = 0
	ByDistance SortByType = 1
)

var DefaultTimeStamp = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
var InvalidDateTime = DefaultTimeStamp.Format(time.RFC3339)

const (
	YYYY_MM_DD_HH_mm_SS = "2006-01-02 15:00:00"
	YYYMMDD = "2006-01-02"
)

type AppointmentStatus int
const (
	Requested AppointmentStatus = 0
	Confirmed AppointmentStatus = 1
	Canceled AppointmentStatus = 2
)