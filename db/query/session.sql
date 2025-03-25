-- name: SelectOpenedSessions :many
select
  gym_id,
  camera_id
from sessions
where user_id = $1;

-- name: InsertSession :exec
insert into sessions (user_id, camera_id)
values ($1, $2);
