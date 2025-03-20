alter table gyms alter column updated_at
drop default;

alter table access_rights alter column updated_at
drop default;

alter table sessions alter column updated_at
drop default;

alter table users alter column updated_at
drop default;
