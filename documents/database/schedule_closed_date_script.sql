#drop table schedule_closed_date;
create table schedule_closed_date
(
    id  int auto_increment,
    npi int(12) not null,
    start_date datetime,
    end_date datetime,
    am_start_time varchar (5),
    am_end_time varchar (5),
    pm_start_time varchar (5),
    pm_end_time varchar (5),
    updated_at datetime,
    created_at datetime,
    constraint schedule_closed_date_pk
    primary key (id)
);

