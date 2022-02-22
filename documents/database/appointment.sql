

create table appointments (
    id  int auto_increment,
    npi int(12) not null,
    appointment_type int,
    appointment_date datetime,
    appointment_status int,
    memo text,
    time_slot int,
    patient_id int(12),
    created_date datetime,
    created_at datetime,
    updated_at datetime,
    constraint appointments_pk
    primary key (id)
)