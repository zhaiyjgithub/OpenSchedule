package doctor

import "time"

type ClosedDateSettings struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-" validate:"omitempty"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi" validate:"gt=0"`
	ClosedDate time.Time `gorm:"column:closed_date;type:timestamp;not null" json:"closedDate" validate:"gt=0"`
	AmStartDateTime time.Time `gorm:"column:am_start_date_time;type:timestamp;not null" json:"amStartDateTime" validate:"gt=0"`
	AmEndDateTime time.Time `gorm:"column:am_end_date_time;type:timestamp;not null" json:"amEndDateTime" validate:"gt=0"`
	PmStartDateTime time.Time `gorm:"column:pm_start_date_time;type:timestamp;not null" json:"pmStartDateTime" validate:"gt=0"`
	PmEndDateTime time.Time `gorm:"column:pm_end_date_time;type:timestamp;not null" json:"pmEndDateTime" validate:"gt=0"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:dateTime" json:"updatedAt" validate:"omitempty"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:dateTime" json:"createdAt" validate:"omitempty"`
}

func (d *ClosedDateSettings) TableName() string  {
	return "schedule_closed_date"
}

