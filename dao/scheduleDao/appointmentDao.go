/**
 * @author zhaiyuanji
 * @date 2022年08月15日 3:26 下午
 */
package scheduleDao

import (
	"OpenSchedule/constant"
	"OpenSchedule/model/doctorModel"
	"OpenSchedule/model/viewModel"
	"time"
)

func (d *Dao) AddAppointment(appointment doctorModel.Appointment) error {
	db := d.engine.Create(&appointment)
	return db.Error
}

func (d *Dao) GetAppointmentByDate(npi int64, startDate time.Time, endDate time.Time) ([]doctorModel.Appointment, error) {
	var l []doctorModel.Appointment
	db := d.engine.Where("npi = ? AND appointment_date >= ? AND appointment_date <= ?", npi, startDate, endDate).Find(&l)
	return l, db.Error
}

func (d *Dao) GetAppointmentByRange(
	npi int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctorModel.Appointment {
	appts := make([]*doctorModel.Appointment, 0)
	_ = d.engine.Where("npi = ? AND appointment_status = ? AND appointment_date >= ? AND appointment_date <= ?",
		npi, appointmentStatus, startDate, endDate).Find(&appts)
	return appts
}

func (d *Dao) GetAppointmentsByRange(
	npi []int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctorModel.Appointment {
	appts := make([]*doctorModel.Appointment, 0)
	_ = d.engine.Where("npi IN ? AND appointment_status = ? AND appointment_date >= ? AND appointment_date <= ?",
		npi, appointmentStatus, startDate, endDate).Find(&appts)
	return appts
}

func (d *Dao) GetAppointmentInfo(patientID int, page int, pageSize int) ([]viewModel.AppointmentInfo, error) {
	var appts []viewModel.AppointmentInfo
	db := d.engine.Model(&doctorModel.Appointment{}).Select("appointments.*, doctors.full_name as doctor_full_name, doctors.phone as doctor_phone, doctors.specialty as doctor_specialty, doctors.sub_specialty as doctor_sub_specialty, doctors.address as doctor_address").Joins("left join doctors on appointments.npi = doctors.npi").
		Order("appointments.created_at").Where("appointments.patient_id = ?", patientID).Limit(pageSize).Offset((page - 1) * pageSize).Find(&appts)
	return appts, db.Error
}
