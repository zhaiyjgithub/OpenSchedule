/**
 * @author zhaiyuanji
 * @date 2022年08月15日 3:26 下午
 */
package schedule

import (
	"OpenSchedule/constant"
	"OpenSchedule/model/doctor"
	"time"
)

func (d *Dao) AddAppointment(appointment doctor.Appointment) error {
	db := d.engine.Create(&appointment)
	return db.Error
}

func (d *Dao) GetAppointmentByDate(npi int64, startDate time.Time, endDate time.Time) ([]doctor.Appointment, error) {
	var l []doctor.Appointment
	db := d.engine.Where("npi = ? AND appointment_date >= ? AND appointment_date <= ?", npi, startDate, endDate).Find(&l)
	return l, db.Error
}

func (d *Dao) GetAppointmentByRange(
	npi int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctor.Appointment {
	appts := make([]*doctor.Appointment, 0)
	_ = d.engine.Where("npi = ? AND appointment_status = ? AND appointment_date >= ? AND appointment_date <= ?",
		npi, appointmentStatus, startDate, endDate).Find(&appts)
	return appts
}

func (d *Dao) GetAppointmentsByRange(
	npi []int64,
	appointmentStatus constant.AppointmentStatus,
	startDate time.Time,
	endDate time.Time,
) []*doctor.Appointment {
	appts := make([]*doctor.Appointment, 0)
	_ = d.engine.Where("npi IN ? AND appointment_status = ? AND appointment_date >= ? AND appointment_date <= ?",
		npi, appointmentStatus, startDate, endDate).Find(&appts)
	return appts
}
