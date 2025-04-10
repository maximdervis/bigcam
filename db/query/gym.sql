-- name: InsertGym :exec
insert into gyms (name, auth_key)
values ($1, $2);

-- name: SelectGymInfo :one
select
  name
from gyms u
where id = $1
limit 1;

-- name: SelectGymIdByAuthKey :one
select
  id
from gyms u
where auth_key = $1
limit 1;
