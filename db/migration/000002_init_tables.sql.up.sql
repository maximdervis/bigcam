drop table if exists gyms;

create table if not exists gyms (
  id integer primary key,
  name varchar not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null
);

create table if not exists users (
  id integer primary key,
  login varchar not null,
  first_name varchar not null,
  last_name varchar not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null
);

create type access_type as enum (
  'readonly',
  'admin'
);

create table if not exists access_grants (
  id integer primary key,
  user_id integer not null,
  gym_id integer not null,
  access_type access_type not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null
);

create table if not exists sessions (
  id integer primary key,
  user_id integer not null,
  camera_id integer not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null
);
