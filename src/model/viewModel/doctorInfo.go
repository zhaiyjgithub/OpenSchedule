package viewModel

import (
	"OpenSchedule/src/model"
	"github.com/olivere/elastic/v7"
)

type Location struct {
	Lat float64
	Lng float64
}
type DoctorInfo struct {
	model.Doctor
	Location elastic.GeoPoint
	Distance float64
}

