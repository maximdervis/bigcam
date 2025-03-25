alter table users
add column password varchar not null;

alter table users
drop column first_name;

alter table users
drop column last_name;

alter table users
add column name varchar not null;
