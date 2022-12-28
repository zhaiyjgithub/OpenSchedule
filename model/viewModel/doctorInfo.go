package viewModel

import (
	"OpenSchedule/model/doctorModel"
	"github.com/olivere/elastic/v7"
	"time"
)

type Location struct {
	Lat float64
	Lng float64
}
type DoctorInfo struct {
	doctorModel.Doctor
	Location                  elastic.GeoPoint `json:"location"`
	Distance                  float64          `json:"distance"`
	NextAvailableDateInClinic string           `json:"nextAvailableDateInClinic"`
	NextAvailableDateVirtual  string           `json:"nextAvailableDateVirtual"`
}

type DoctorDetailInfo struct {
	doctorModel.Doctor
	Lat            float64                      `json:"lat"`
	Lng            float64                      `json:"lng"`
	Language       string                       `json:"language"`
	Awards         []doctorModel.Awards         `json:"awards"`
	Certifications []doctorModel.Certifications `json:"certifications"`
	Educations     []doctorModel.Educations     `json:"educations"`
	Insurances     string                       `json:"insurances"`
}

type TimeSlotPerDay struct {
	Date      time.Time              `json:"date"`
	TimeSlots []doctorModel.TimeSlot `json:"timeSlots"`
}

type BackgroundInfo struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
