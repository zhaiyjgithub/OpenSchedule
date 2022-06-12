drop table sub_users;
create table sub_users (
    id int auto_increment,
    first_name varchar(50),
    last_name varchar(50),
    email varchar(50),
    phone varchar(20),
    birthday char(10),
    gender char(1),
    user_id int,
    is_legal tinyint(1),
    constraint sub_users_pk
    primary key (id)
)
