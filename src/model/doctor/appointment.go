/**
 * @author zhaiyuanji
 * @date 2022年02月22日 1:43 下午
 */
package doctor

import (
	"OpenSchedule/src/constant"
	"time"
)

// Appointments [...]
type	Appointments	struct {
	ID	int	`gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi"`
	AppointmentType constant.AppointmentType	`gorm:"column:appointment_type;type:int(11)" json:"appointmentType"`
	AppointmentDate	time.Time	`gorm:"column:appointment_date;type:datetime" json:"appointmentDate"`
	AppointmentStatus	constant.AppointmentStatus	`gorm:"column:appointment_status;type:int(11)" json:"appointmentStatus"`
	Memo	string	`gorm:"column:memo;type:text" json:"memo"`
	TimeSlot	int	`gorm:"column:time_slot;type:int(11)" json:"timeSlot"`
	PatientID	int	`gorm:"column:patient_id;type:int(12)" json:"patientId"`
	CreatedDate	time.Time	`gorm:"column:created_date;type:datetime" json:"createdDate"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}
