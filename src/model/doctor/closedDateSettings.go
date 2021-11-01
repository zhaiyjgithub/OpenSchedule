package doctor

import "time"

// ClosedDateSettings [...]
type  	ClosedDateSettings	struct {
	ID	int	`gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi"`
	StartDate	time.Time	`gorm:"column:start_date;type:timestamp;not null;default:2000-01-01 00:00:00" json:"startDate"`
	EndDate	time.Time	`gorm:"column:end_date;type:timestamp;not null;default:2000-01-01 00:00:00" json:"endDate"`
	AmStartDateTime	time.Time	`gorm:"column:am_start_date_time;type:timestamp;not null;default:2000-01-01 00:00:00" json:"amStartDateTime"`
	AmEndDateTime	time.Time	`gorm:"column:am_end_date_time;type:timestamp;not null;default:2000-01-01 00:00:00" json:"amEndDateTime"`
	PmStartDateTime	time.Time	`gorm:"column:pm_start_date_time;type:timestamp;not null;default:2000-01-01 00:00:00" json:"pmStartDateTime"`
	PmEndDateTime	time.Time	`gorm:"column:pm_end_date_time;type:timestamp;not null;default:2000-01-01 00:00:00" json:"pmEndDateTime"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:datetime" json:"updatedAt"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:datetime" json:"createdAt"`
}

func (d *ClosedDateSettings) TableName() string  {
	return "closed_date_settings"
}
