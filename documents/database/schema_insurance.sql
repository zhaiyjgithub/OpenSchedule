drop table insurances;
create table insurances (
  id  int auto_increment,
  npi int(12) not null,
  name text,
  created_at datetime,
  updated_at datetime,
  constraint appointments_pk
      primary key (id)
)