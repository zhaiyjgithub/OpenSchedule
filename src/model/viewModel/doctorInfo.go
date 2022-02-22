package viewModel

import (
	"OpenSchedule/src/model/doctor"
	"github.com/olivere/elastic/v7"
	"time"
)

type Location struct {
	Lat float64
	Lng float64
}
type DoctorInfo struct {
	doctor.Doctor
	Location elastic.GeoPoint `json:"location"`
	Distance float64 `json:"distance"`
	NextAvailableDateInClinic string `json:"nextAvailableDateInClinic"`
	NextAvailableDateVirtual string `json:"nextAvailableDateVirtual"`
}

type TimeSlotPeerDay struct {
	Date time.Time `json:"date"`
	TimeSlots []doctor.TimeSlot `json:"timeSlots"`
}