package doctor

import "time"

type ClosedDateSettings struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-" validate:"omitempty"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi" validate:"gt=0"`
	StartDate	time.Time	`gorm:"column:start_date;type:dateTime" json:"startDate" validate:"omitempty"`
	EndDate	time.Time	`gorm:"column:end_date;type:dateTime" json:"endDate" validate:"omitempty"`
	AmStartTime	string	`gorm:"column:am_start_time;type:varchar(5)" json:"amStartTime" validate:"hh:mm"`
	AmEndTime	string	`gorm:"column:am_end_time;type:varchar(5)" json:"amEndTime" validate:"hh:mm"`
	PmStartTime	string	`gorm:"column:pm_start_time;type:varchar(5)" json:"pmStartTime" validate:"hh:mm"`
	PmEndTime	string	`gorm:"column:pm_end_time;type:varchar(5)" json:"pmEndTime" validate:"hh:mm"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:dateTime" json:"updatedAt" validate:"omitempty"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:dateTime" json:"createdAt" validate:"omitempty"`
}

func (d *ClosedDateSettings) TableName() string  {
	return "schedule_closed_date"
}