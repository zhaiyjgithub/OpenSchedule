#drop table schedule_closed_date;
# use openSchedule;
drop table schedule_closed_date;




create table closed_date_settings
(
    id  int auto_increment,
    npi int(12) not null,
    start_date timestamp DEFAULT '2000-01-01 00:00:00',
    end_date timestamp DEFAULT '2000-01-01 00:00:00',
    am_start_date_time timestamp DEFAULT '2000-01-01 00:00:00',
    am_end_date_time timestamp DEFAULT '2000-01-01 00:00:00',
    pm_start_date_time timestamp DEFAULT '2000-01-01 00:00:00',
    pm_end_date_time timestamp DEFAULT '2000-01-01 00:00:00',
    updated_at datetime,
    created_at datetime,
    constraint schedule_closed_date_pk
    primary key (id)
);

desc closed_date_settings
