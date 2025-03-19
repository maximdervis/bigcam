alter table gyms alter column updated_at
set default current_timestamp;

alter table access_grants alter column updated_at
set default current_timestamp;

alter table sessions alter column updated_at
set default current_timestamp;

alter table users alter column updated_at
set default current_timestamp;
