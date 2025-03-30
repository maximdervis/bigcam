-- name: SelectOpenedSessions :many
select
  id,
  gym_id,
  camera_id
from sessions
where 1=1
  and user_id = $1
  and opened = true;

-- name: CloseSession :exec
update sessions
set opened = false
where id = $1;

-- name: SelectOccupiedCams :many
select
  gym_id,
  camera_id,
  name
from sessions s
left join users u
on s.user_id = u.id
where 1=1
  and gym_id = $1
;

-- name: InsertSession :one
insert into sessions (user_id, gym_id, camera_id)
values ($1, $2, $3)
returning id;
