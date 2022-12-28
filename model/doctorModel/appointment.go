/**
 * @author zhaiyuanji
 * @date 2022年02月22日 1:43 下午
 */
package doctorModel

import "time"

type Appointment struct {
	ID                     int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	DoctorID               int       `gorm:"column:doctor_id;type:int(11);not null" json:"doctorId"`
	Npi                    int64     `gorm:"column:npi;type:int(12);not null" json:"npi"`
	AppointmentType        int       `gorm:"column:appointment_type;type:int(11)" json:"appointmentType"`
	AppointmentDate        time.Time `gorm:"column:appointment_date;type:datetime" json:"appointmentDate"`
	AppointmentStatus      int       `gorm:"column:appointment_status;type:int(11)" json:"appointmentStatus"`
	Memo                   string    `gorm:"column:memo;type:text" json:"memo"`
	Offset                 int       `gorm:"column:offset;type:int(11)" json:"offset"`
	PatientID              int       `gorm:"column:patient_id;type:int(12)" json:"patientId"`
	LegalGuardianPatientID int       `gorm:"column:legal_guardian_patient_id;type:int(12)" json:"legalGuardianPatientId"`
	FirstName              string    `gorm:"column:first_name;type:varchar(50)" json:"firstName"`
	LastName               string    `gorm:"column:last_name;type:varchar(50)" json:"lastName"`
	Dob                    string    `gorm:"column:dob;type:char(10)" json:"dob"`
	Gender                 string    `gorm:"column:gender;type:char(1)" json:"gender"`
	Email                  string    `gorm:"column:email;type:varchar(100)" json:"email"`
	Phone                  string    `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Insurance              int       `gorm:"column:insurance;type:int(11)" json:"insurance"`
	VisitReason            string    `gorm:"column:visit_reason;type:varchar(100)" json:"visitReason"`
	IsNewPatient           bool      `gorm:"column:is_new_patient;type:tinyint(1)" json:"isNewPatient"`
	CreatedDate            time.Time `gorm:"column:created_date;type:datetime" json:"createdDate"`
	CreatedAt              time.Time `gorm:"column:created_at;type:datetime" json:"createdAt"`
	UpdatedAt              time.Time `gorm:"column:updated_at;type:datetime" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (d *Appointment) TableName() string {
	return "appointments"
}
