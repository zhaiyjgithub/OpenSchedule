package doctor

import (
	"OpenSchedule/src/constant"
	"time"
)

type ScheduleSettings struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-" validate:"omitempty"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi" validate:"gt=0"`
	DurationPerSlot	int	`gorm:"column:duration_per_slot;type:int(11)" json:"durationPerSlot" validate:"gt=0"`
	NumberPerSlot	int	`gorm:"column:number_per_slot;type:int(11)" json:"numberPerSlot" validate:"gt=0"`
	MondayAmIsEnable	bool	`gorm:"column:monday_am_is_enable;type:tinyint(1)" json:"mondayAmIsEnable"`
	MondayAmStartTime	string	`gorm:"column:monday_am_start_time;type:varchar(5)" json:"mondayAmStartTime" validate:"hh:mm"`
	MondayAmEndTime	string	`gorm:"column:monday_am_end_time;type:varchar(5)" json:"mondayAmEndTime" validate:"hh:mm"`
	MondayAmAppointmentType	constant.AppointmentType	`gorm:"column:monday_am_appointment_type;type:int(11)" json:"mondayAmAppointmentType"`
	MondayPmIsEnable	bool	`gorm:"column:monday_pm_is_enable;type:tinyint(1)" json:"mondayPmIsEnable"`
	MondayPmStartTime	string	`gorm:"column:monday_pm_start_time;type:varchar(5)" json:"mondayPmStartTime" validate:"hh:mm"`
	MondayPmEndTime	string	`gorm:"column:monday_pm_end_time;type:varchar(5)" json:"mondayPmEndTime" validate:"hh:mm"`
	MondayPmAppointmentType	constant.AppointmentType	`gorm:"column:monday_pm_appointment_type;type:varchar(5)" json:"mondayPmAppointmentType"`
	TuesdayAmIsEnable	bool	`gorm:"column:tuesday_am_is_enable;type:tinyint(1)" json:"tuesdayAmIsEnable"`
	TuesdayAmStartTime	string	`gorm:"column:tuesday_am_start_time;type:varchar(5)" json:"tuesdayAmStartTime" validate:"hh:mm"`
	TuesdayAmEndTime	string	`gorm:"column:tuesday_am_end_time;type:varchar(5)" json:"tuesdayAmEndTime" validate:"hh:mm"`
	TuesdayAmAppointmentType	constant.AppointmentType	`gorm:"column:tuesday_am_appointment_type;type:int(11)" json:"tuesdayAmAppointmentType"`
	TuesdayPmIsEnable	bool	`gorm:"column:tuesday_pm_is_enable;type:tinyint(1)" json:"tuesdayPmIsEnable"`
	TuesdayPmStartTime	string	`gorm:"column:tuesday_pm_start_time;type:varchar(5)" json:"tuesdayPmStartTime" validate:"hh:mm"`
	TuesdayPmEndTime	string	`gorm:"column:tuesday_pm_end_time;type:varchar(5)" json:"tuesdayPmEndTime" validate:"hh:mm"`
	TuesdayPmAppointmentType	constant.AppointmentType	`gorm:"column:tuesday_pm_appointment_type;type:varchar(5)" json:"tuesdayPmAppointmentType"`
	WednesdayAmIsEnable	bool	`gorm:"column:wednesday_am_is_enable;type:tinyint(1)" json:"wednesdayAmIsEnable"`
	WednesdayAmStartTime	string	`gorm:"column:wednesday_am_start_time;type:varchar(5)" json:"wednesdayAmStartTime" validate:"hh:mm"`
	WednesdayAmEndTime	string	`gorm:"column:wednesday_am_end_time;type:varchar(5)" json:"wednesdayAmEndTime" validate:"hh:mm"`
	WednesdayAmAppointmentType	constant.AppointmentType	`gorm:"column:wednesday_am_appointment_type;type:int(11)" json:"wednesdayAmAppointmentType"`
	WednesdayPmIsEnable	bool	`gorm:"column:wednesday_pm_is_enable;type:tinyint(1)" json:"wednesdayPmIsEnable"`
	WednesdayPmStartTime	string	`gorm:"column:wednesday_pm_start_time;type:varchar(5)" json:"wednesdayPmStartTime" validate:"hh:mm"`
	WednesdayPmEndTime	string	`gorm:"column:wednesday_pm_end_time;type:varchar(5)" json:"wednesdayPmEndTime" validate:"hh:mm"`
	WednesdayPmAppointmentType	constant.AppointmentType	`gorm:"column:wednesday_pm_appointment_type;type:varchar(5)" json:"wednesdayPmAppointmentType"`
	ThursdayAmIsEnable	bool	`gorm:"column:thursday_am_is_enable;type:tinyint(1)" json:"thursdayAmIsEnable"`
	ThursdayAmStartTime	string	`gorm:"column:thursday_am_start_time;type:varchar(5)" json:"thursdayAmStartTime" validate:"hh:mm"`
	ThursdayAmEndTime	string	`gorm:"column:thursday_am_end_time;type:varchar(5)" json:"thursdayAmEndTime" validate:"hh:mm"`
	ThursdayAmAppointmentType	constant.AppointmentType	`gorm:"column:thursday_am_appointment_type;type:int(11)" json:"thursdayAmAppointmentType"`
	ThursdayPmIsEnable	bool	`gorm:"column:thursday_pm_is_enable;type:tinyint(1)" json:"thursdayPmIsEnable"`
	ThursdayPmStartTime	string	`gorm:"column:thursday_pm_start_time;type:varchar(5)" json:"thursdayPmStartTime" validate:"hh:mm"`
	ThursdayPmEndTime	string	`gorm:"column:thursday_pm_end_time;type:varchar(5)" json:"thursdayPmEndTime" validate:"hh:mm"`
	ThursdayPmAppointmentType	constant.AppointmentType	`gorm:"column:thursday_pm_appointment_type;type:varchar(5)" json:"thursdayPmAppointmentType"`
	FridayAmIsEnable	bool	`gorm:"column:friday_am_is_enable;type:tinyint(1)" json:"fridayAmIsEnable"`
	FridayAmStartTime	string	`gorm:"column:friday_am_start_time;type:varchar(5)" json:"fridayAmStartTime" validate:"hh:mm"`
	FridayAmEndTime	string	`gorm:"column:friday_am_end_time;type:varchar(5)" json:"fridayAmEndTime" validate:"hh:mm"`
	FridayAmAppointmentType	constant.AppointmentType	`gorm:"column:friday_am_appointment_type;type:int(11)" json:"fridayAmAppointmentType"`
	FridayPmIsEnable	bool	`gorm:"column:friday_pm_is_enable;type:tinyint(1)" json:"fridayPmIsEnable"`
	FridayPmStartTime	string	`gorm:"column:friday_pm_start_time;type:varchar(5)" json:"fridayPmStartTime" validate:"hh:mm"`
	FridayPmEndTime	string	`gorm:"column:friday_pm_end_time;type:varchar(5)" json:"fridayPmEndTime" validate:"hh:mm"`
	FridayPmAppointmentType	constant.AppointmentType	`gorm:"column:friday_pm_appointment_type;type:varchar(5)" json:"fridayPmAppointmentType"`
	SaturdayAmIsEnable	bool	`gorm:"column:saturday_am_is_enable;type:tinyint(1)" json:"saturdayAmIsEnable"`
	SaturdayAmStartTime	string	`gorm:"column:saturday_am_start_time;type:varchar(5)" json:"saturdayAmStartTime" validate:"hh:mm"`
	SaturdayAmEndTime	string	`gorm:"column:saturday_am_end_time;type:varchar(5)" json:"saturdayAmEndTime" validate:"hh:mm"`
	SaturdayAmAppointmentType	constant.AppointmentType	`gorm:"column:saturday_am_appointment_type;type:int(11)" json:"saturdayAmAppointmentType"`
	SaturdayPmIsEnable	bool	`gorm:"column:saturday_pm_is_enable;type:tinyint(1)" json:"saturdayPmIsEnable"`
	SaturdayPmStartTime	string	`gorm:"column:saturday_pm_start_time;type:varchar(5)" json:"saturdayPmStartTime" validate:"hh:mm"`
	SaturdayPmEndTime	string	`gorm:"column:saturday_pm_end_time;type:varchar(5)" json:"saturdayPmEndTime" validate:"hh:mm"`
	SaturdayPmAppointmentType	constant.AppointmentType	`gorm:"column:saturday_pm_appointment_type;type:varchar(5)" json:"saturdayPmAppointmentType"`
	SundayAmIsEnable	bool	`gorm:"column:sunday_am_is_enable;type:tinyint(1)" json:"sundayAmIsEnable"`
	SundayAmStartTime	string	`gorm:"column:sunday_am_start_time;type:varchar(5)" json:"sundayAmStartTime" validate:"hh:mm"`
	SundayAmEndTime	string	`gorm:"column:sunday_am_end_time;type:varchar(5)" json:"sundayAmEndTime" validate:"hh:mm"`
	SundayAmAppointmentType	constant.AppointmentType	`gorm:"column:sunday_am_appointment_type;type:int(11)" json:"sundayAmAppointmentType"`
	SundayPmIsEnable	bool	`gorm:"column:sunday_pm_is_enable;type:tinyint(1)" json:"sundayPmIsEnable"`
	SundayPmStartTime	string	`gorm:"column:sunday_pm_start_time;type:varchar(5)" json:"sundayPmStartTime" validate:"hh:mm"`
	SundayPmEndTime	string	`gorm:"column:sunday_pm_end_time;type:varchar(5)" json:"sundayPmEndTime" validate:"hh:mm"`
	SundayPmAppointmentType	constant.AppointmentType	`gorm:"column:sunday_pm_appointment_type;type:varchar(5)" json:"sundayPmAppointmentType"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:dateTime" json:"updatedAt" validate:"omitempty"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:dateTime" json:"createdAt" validate:"omitempty"`
}

// TableName sets the insert table name for this struct type
func (d *ScheduleSettings) TableName() string {
	return "schedule_settings"
}
