-- name: SelectOpenedSessions :many
select
  gym_id,
  camera_id
from sessions
where user_id = $1;
