package viewModel

import (
	"OpenSchedule/src/model/doctor"
	"github.com/olivere/elastic/v7"
)

type Location struct {
	Lat float64
	Lng float64
}
type DoctorInfo struct {
	doctor.Doctor
	Location elastic.GeoPoint
	Distance float64
}

