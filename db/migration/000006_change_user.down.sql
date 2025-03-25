alter table users
drop column password;

alter table users
add column first_name varchar not null;

alter table users
add column last_name varchar not null;

alter table users
drop column name;
