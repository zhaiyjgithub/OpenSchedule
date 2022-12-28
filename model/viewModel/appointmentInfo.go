package viewModel

import "OpenSchedule/model/doctorModel"

type AppointmentInfo struct {
	doctorModel.Appointment
	DoctorFullName     string `gorm:"column:doctor_full_name;type:varchar(255)" json:"doctorFullName"`
	DoctorAddress      string `gorm:"column:doctor_address;type:varchar(255)" json:"doctorAddress"`
	DoctorSpecialty    string `gorm:"column:doctor_specialty;type:varchar(255)" json:"doctorSpecialty"`
	DoctorSubSpecialty string `gorm:"column:doctor_sub_specialty;type:varchar(255)" json:"doctorSubSpecialty"`
	DoctorPhone        string `gorm:"column:doctor_phone;type:varchar(255)" json:"doctorPhone"`
}
