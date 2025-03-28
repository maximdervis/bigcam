alter table sessions
alter column gym_id type bigint;

alter table sessions
alter column user_id type bigint;

alter table sessions
alter column camera_id type bigint;

alter table access_grants
alter column gym_id type bigint;

alter table access_grants
alter column user_id type bigint;

