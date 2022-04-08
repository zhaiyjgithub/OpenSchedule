drop table users;
create table users (
    id int auto_increment,
    first_name varchar(50),
    last_name varchar(50),
    gender char(1),
    email varchar(100),
    phone varchar(20),
    address text,
    password varchar(32),
    created_at datetime,
    updated_at datetime,
    constraint users_pk
    primary key (id)
)