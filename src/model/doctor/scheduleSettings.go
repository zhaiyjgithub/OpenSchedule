package doctor

import (
	"OpenSchedule/src/constant"
	"time"
)

// ScheduleSettings [...]	
type	ScheduleSettings	struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-"`
	Npi	int64	`gorm:"column:npi;type:int(12);not null" json:"npi"`
	DurationPerSlot	int	`gorm:"column:duration_per_slot;type:int(11);default:15" json:"durationPerSlot"`
	NumberPerSlot	int	`gorm:"column:number_per_slot;type:int(11);default:1" json:"numberPerSlot"`
	MondayAmIsEnable	bool	`gorm:"column:monday_am_is_enable;type:tinyint(1);default:0" json:"mondayAmIsEnable"`
	MondayAmStartTimeOffset	int	`gorm:"column:monday_am_start_time_offset;type:int(11)" json:"mondayAmStartTimeOffset"`
	MondayAmEndTimeOffset	int	`gorm:"column:monday_am_end_time_offset;type:int(11)" json:"mondayAmEndTimeOffset"`
	MondayAmAppointmentType	constant.AppointmentType	`gorm:"column:monday_am_appointment_type;type:int(11)" json:"mondayAmAppointmentType"`
	MondayPmIsEnable	bool	`gorm:"column:monday_pm_is_enable;type:tinyint(1);default:0" json:"mondayPmIsEnable"`
	MondayPmStartTimeOffset	int	`gorm:"column:monday_pm_start_time_offset;type:int(11)" json:"mondayPmStartTimeOffset"`
	MondayPmEndTimeOffset	int	`gorm:"column:monday_pm_end_time_offset;type:int(11)" json:"mondayPmEndTimeOffset"`
	MondayPmAppointmentType	constant.AppointmentType	`gorm:"column:monday_pm_appointment_type;type:int(11)" json:"mondayPmAppointmentType"`
	TuesdayAmIsEnable	bool	`gorm:"column:tuesday_am_is_enable;type:tinyint(1);default:0" json:"tuesdayAmIsEnable"`
	TuesdayAmStartTimeOffset	int	`gorm:"column:tuesday_am_start_time_offset;type:int(11)" json:"tuesdayAmStartTimeOffset"`
	TuesdayAmEndTimeOffset	int	`gorm:"column:tuesday_am_end_time_offset;type:int(11)" json:"tuesdayAmEndTimeOffset"`
	TuesdayAmAppointmentType	constant.AppointmentType	`gorm:"column:tuesday_am_appointment_type;type:int(11)" json:"tuesdayAmAppointmentType"`
	TuesdayPmIsEnable	bool	`gorm:"column:tuesday_pm_is_enable;type:tinyint(1);default:0" json:"tuesdayPmIsEnable"`
	TuesdayPmStartTimeOffset	int	`gorm:"column:tuesday_pm_start_time_offset;type:int(11)" json:"tuesdayPmStartTimeOffset"`
	TuesdayPmEndTimeOffset	int	`gorm:"column:tuesday_pm_end_time_offset;type:int(11)" json:"tuesdayPmEndTimeOffset"`
	TuesdayPmAppointmentType	constant.AppointmentType	`gorm:"column:tuesday_pm_appointment_type;type:int(11)" json:"tuesdayPmAppointmentType"`
	WednesdayAmIsEnable	bool	`gorm:"column:wednesday_am_is_enable;type:tinyint(1);default:0" json:"wednesdayAmIsEnable"`
	WednesdayAmStartTimeOffset	int	`gorm:"column:wednesday_am_start_time_offset;type:int(11)" json:"wednesdayAmStartTimeOffset"`
	WednesdayAmEndTimeOffset	int	`gorm:"column:wednesday_am_end_time_offset;type:int(11)" json:"wednesdayAmEndTimeOffset"`
	WednesdayAmAppointmentType	constant.AppointmentType	`gorm:"column:wednesday_am_appointment_type;type:int(11)" json:"wednesdayAmAppointmentType"`
	WednesdayPmIsEnable	bool	`gorm:"column:wednesday_pm_is_enable;type:tinyint(1);default:0" json:"wednesdayPmIsEnable"`
	WednesdayPmStartTimeOffset	int	`gorm:"column:wednesday_pm_start_time_offset;type:int(11)" json:"wednesdayPmStartTimeOffset"`
	WednesdayPmEndTimeOffset	int	`gorm:"column:wednesday_pm_end_time_offset;type:int(11)" json:"wednesdayPmEndTimeOffset"`
	WednesdayPmAppointmentType	constant.AppointmentType	`gorm:"column:wednesday_pm_appointment_type;type:int(11)" json:"wednesdayPmAppointmentType"`
	ThursdayAmIsEnable	bool	`gorm:"column:thursday_am_is_enable;type:tinyint(1);default:0" json:"thursdayAmIsEnable"`
	ThursdayAmStartTimeOffset	int	`gorm:"column:thursday_am_start_time_offset;type:int(11)" json:"thursdayAmStartTimeOffset"`
	ThursdayAmEndTimeOffset	int	`gorm:"column:thursday_am_end_time_offset;type:int(11)" json:"thursdayAmEndTimeOffset"`
	ThursdayAmAppointmentType	constant.AppointmentType	`gorm:"column:thursday_am_appointment_type;type:int(11)" json:"thursdayAmAppointmentType"`
	ThursdayPmIsEnable	bool	`gorm:"column:thursday_pm_is_enable;type:tinyint(1);default:0" json:"thursdayPmIsEnable"`
	ThursdayPmStartTimeOffset	int	`gorm:"column:thursday_pm_start_time_offset;type:int(11)" json:"thursdayPmStartTimeOffset"`
	ThursdayPmEndTimeOffset	int	`gorm:"column:thursday_pm_end_time_offset;type:int(11)" json:"thursdayPmEndTimeOffset"`
	ThursdayPmAppointmentType	constant.AppointmentType	`gorm:"column:thursday_pm_appointment_type;type:int(11)" json:"thursdayPmAppointmentType"`
	FridayAmIsEnable	bool	`gorm:"column:friday_am_is_enable;type:tinyint(1);default:0" json:"fridayAmIsEnable"`
	FridayAmStartTimeOffset	int	`gorm:"column:friday_am_start_time_offset;type:int(11)" json:"fridayAmStartTimeOffset"`
	FridayAmEndTimeOffset	int	`gorm:"column:friday_am_end_time_offset;type:int(11)" json:"fridayAmEndTimeOffset"`
	FridayAmAppointmentType	constant.AppointmentType	`gorm:"column:friday_am_appointment_type;type:int(11)" json:"fridayAmAppointmentType"`
	FridayPmIsEnable	bool	`gorm:"column:friday_pm_is_enable;type:tinyint(1);default:0" json:"fridayPmIsEnable"`
	FridayPmStartTimeOffset	int	`gorm:"column:friday_pm_start_time_offset;type:int(11)" json:"fridayPmStartTimeOffset"`
	FridayPmEndTimeOffset	int	`gorm:"column:friday_pm_end_time_offset;type:int(11)" json:"fridayPmEndTimeOffset"`
	FridayPmAppointmentType	constant.AppointmentType	`gorm:"column:friday_pm_appointment_type;type:int(11)" json:"fridayPmAppointmentType"`
	SaturdayAmIsEnable	bool	`gorm:"column:saturday_am_is_enable;type:tinyint(1);default:0" json:"saturdayAmIsEnable"`
	SaturdayAmStartTimeOffset	int	`gorm:"column:saturday_am_start_time_offset;type:int(11)" json:"saturdayAmStartTimeOffset"`
	SaturdayAmEndTimeOffset	int	`gorm:"column:saturday_am_end_time_offset;type:int(11)" json:"saturdayAmEndTimeOffset"`
	SaturdayAmAppointmentType	constant.AppointmentType	`gorm:"column:saturday_am_appointment_type;type:int(11)" json:"saturdayAmAppointmentType"`
	SaturdayPmIsEnable	bool	`gorm:"column:saturday_pm_is_enable;type:tinyint(1);default:0" json:"saturdayPmIsEnable"`
	SaturdayPmStartTimeOffset	int	`gorm:"column:saturday_pm_start_time_offset;type:int(11)" json:"saturdayPmStartTimeOffset"`
	SaturdayPmEndTimeOffset	int	`gorm:"column:saturday_pm_end_time_offset;type:int(11)" json:"saturdayPmEndTimeOffset"`
	SaturdayPmAppointmentType	constant.AppointmentType	`gorm:"column:saturday_pm_appointment_type;type:int(11)" json:"saturdayPmAppointmentType"`
	SundayAmIsEnable	bool	`gorm:"column:sunday_am_is_enable;type:tinyint(1);default:0" json:"sundayAmIsEnable"`
	SundayAmStartTimeOffset	int	`gorm:"column:sunday_am_start_time_offset;type:int(11)" json:"sundayAmStartTimeOffset"`
	SundayAmEndTimeOffset	int	`gorm:"column:sunday_am_end_time_offset;type:int(11)" json:"sundayAmEndTimeOffset"`
	SundayAmAppointmentType	constant.AppointmentType	`gorm:"column:sunday_am_appointment_type;type:int(11)" json:"sundayAmAppointmentType"`
	SundayPmIsEnable	bool	`gorm:"column:sunday_pm_is_enable;type:tinyint(1);default:0" json:"sundayPmIsEnable"`
	SundayPmStartTimeOffset	int	`gorm:"column:sunday_pm_start_time_offset;type:int(11)" json:"sundayPmStartTimeOffset"`
	SundayPmEndTimeOffset	int	`gorm:"column:sunday_pm_end_time_offset;type:int(11)" json:"sundayPmEndTimeOffset"`
	SundayPmAppointmentType	constant.AppointmentType	`gorm:"column:sunday_pm_appointment_type;type:int(11)" json:"sundayPmAppointmentType"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:datetime" json:"-"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:datetime" json:"-"`
}

// TableName sets the insert table name for this struct type
func (d *ScheduleSettings) TableName() string {
	return "schedule_settings"
}
