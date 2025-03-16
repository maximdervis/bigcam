create table if not exists gyms (
  id integer primary key,
  name varchar not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null
);
