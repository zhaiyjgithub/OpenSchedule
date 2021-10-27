package doctor

import "time"

type ClosedDateSettings struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-" validate:"omitempty"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi" validate:"gt=0"`
	ClosedDate int64 `gorm:"column:closed_date;type:int(12);not null" json:"closedDate" validate:"gt=0"`
	AmStartDateTime int64 `gorm:"column:am_start_date_tme;type:int(12);not null" json:"amStartDateTime" validate:"gt=0"`
	AmEndDateTime int64 `gorm:"column:am_end_date_tme;type:int(12);not null" json:"amEndDateTime" validate:"gt=0"`
	PmStartDateTime int64 `gorm:"column:pm_start_date_tme;type:int(12);not null" json:"pmStartDateTime" validate:"gt=0"`
	PmEndTimeDateTime int64 `gorm:"column:pm_end_date_tme;type:int(12);not null" json:"pmEndDateTime" validate:"gt=0"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:dateTime" json:"updatedAt" validate:"omitempty"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:dateTime" json:"createdAt" validate:"omitempty"`
}

func (d *ClosedDateSettings) TableName() string  {
	return "schedule_closed_date"
}