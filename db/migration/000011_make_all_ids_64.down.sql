alter table sessions
alter column gym_id type integer;

alter table sessions
alter column user_id type bigserial;

alter table sessions
alter column camera_id type bigserial;

alter table access_grants
alter column gym_id type bigserial;

alter table access_grants
alter column user_id type bigserial;

