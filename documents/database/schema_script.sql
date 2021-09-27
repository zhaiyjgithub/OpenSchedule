
create table schedule_settings
(
    id int(11) unsigned auto_increment,
    npi int(12) not null,
    duration_per_slot int null,
    number_per_slot int null,
    monday_am_is_enable bool null,
    monday_am_start_time varchar(5) null,
    monday_am_end_time varchar(5) null,
    monday_am_appointment_type int null,
    monday_pm_is_enable bool null,
    monday_pm_start_time varchar(5) null,
    monday_pm_end_time varchar(5) null,
    monday_pm_appointment_type varchar(5) null,
    tuesday_am_is_enable bool null,
    tuesday_am_start_time varchar(5) null,
    tuesday_am_end_time varchar(5) null,
    tuesday_am_appointment_type int null,
    tuesday_pm_is_enable bool null,
    tuesday_pm_start_time varchar(5) null,
    tuesday_pm_end_time varchar(5) null,
    tuesday_pm_appointment_type varchar(5) null,
    wednesday_am_is_enable bool null,
    wednesday_am_start_time varchar(5) null,
    wednesday_am_end_time varchar(5) null,
    wednesday_am_appointment_type int null,
    wednesday_pm_is_enable bool null,
    wednesday_pm_start_time varchar(5) null,
    wednesday_pm_end_time varchar(5) null,
    wednesday_pm_appointment_type varchar(5) null,
    thursday_am_is_enable bool null,
    thursday_am_start_time varchar(5) null,
    thursday_am_end_time varchar(5) null,
    thursday_am_appointment_type int null,
    thursday_pm_is_enable bool null,
    thursday_pm_start_time varchar(5) null,
    thursday_pm_end_time varchar(5) null,
    thursday_pm_appointment_type varchar(5) null,
    friday_am_is_enable bool null,
    friday_am_start_time varchar(5) null,
    friday_am_end_time varchar(5) null,
    friday_am_appointment_type int null,
    friday_pm_is_enable bool null,
    friday_pm_start_time varchar(5) null,
    friday_pm_end_time varchar(5) null,
    friday_pm_appointment_type varchar(5) null,
    saturday_am_is_enable bool null,
    saturday_am_start_time varchar(5) null,
    saturday_am_end_time varchar(5) null,
    saturday_am_appointment_type int null,
    saturday_pm_is_enable bool null,
    saturday_pm_start_time varchar(5) null,
    saturday_pm_end_time varchar(5) null,
    saturday_pm_appointment_type varchar(5) null,
    sunday_am_is_enable bool null,
    sunday_am_start_time varchar(5) null,
    sunday_am_end_time varchar(5) null,
    sunday_am_appointment_type int null,
    sunday_pm_is_enable bool null,
    sunday_pm_start_time varchar(5) null,
    sunday_pm_end_time varchar(5) null,
    sunday_pm_appointment_type varchar(5) null,
    updated_at datetime null,
    created_at datetime null,
    constraint schedule_settings_pk
    primary key (id)
);

