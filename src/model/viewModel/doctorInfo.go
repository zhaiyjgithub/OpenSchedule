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

type DoctorDetailInfo struct {
	doctor.Doctor
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Language string `json:"language"`
	Awards []doctor.Awards `json:"awards"`
	Certifications []doctor.Certifications `json:"certifications"`
	Educations []doctor.Educations `json:"educations"`
	Insurances []doctor.Insurances `json:"insurances"`
}

type TimeSlotPerDay struct {
	Date time.Time `json:"date"`
	TimeSlots []doctor.TimeSlot `json:"timeSlots"`
}

type BackgroundInfo struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}