package doctorModel

import "time"

type DoctorUser struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Npi       int       `gorm:"index:index_npi;column:npi;type:int(20)" json:"npi"`
	FullName  string    `gorm:"column:full_name;type:varchar(50)" json:"fullName"`
	FirstName string    `gorm:"column:first_name;type:varchar(50)" json:"firstName"`
	LastName  string    `gorm:"column:last_name;type:varchar(50)" json:"lastName"`
	Email     string    `gorm:"column:email;type:varchar(100)" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(32)" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"-"`
}

func (u DoctorUser) TableName() string {
	return "doctor_users"
}
