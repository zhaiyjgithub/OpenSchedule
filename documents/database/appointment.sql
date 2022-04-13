
create table appointments (
    id  int auto_increment,
    doctor_id int not null ,
    npi int(12) not null,
    appointment_type int,
    appointment_date datetime,
    appointment_status int,
    memo text,
    time_slot int,
    patient_id int(12),
    legal_guardian_patient_id int(12),
    first_name varchar(50),
    last_name varchar(50),
    dob char(10),
    gender char(1),
    email varchar(100),
    phone varchar(20),
    insurance int,
    visit_reason varchar(100),
    is_new_patient tinyint(1),
    created_date datetime,
    created_at datetime,
    updated_at datetime,
    constraint appointments_pk
    primary key (id)
)