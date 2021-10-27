#drop table schedule_closed_date;
use openSchedule;
drop table schedule_closed_date;
create table schedule_closed_date
(
    id  int auto_increment,
    npi int(12) not null,
    closed_date int(12) not null ,
    am_start_date_time int(12),
    am_end_date_time int(12),
    pm_start_date_time int(12),
    pm_end_date_time int(12),
    updated_at datetime,
    created_at datetime,
    constraint schedule_closed_date_pk
    primary key (id)
);

