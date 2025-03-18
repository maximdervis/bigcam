-- name: GetAvailableGyms :many
select
  u.id as user_id,
  g.gym_id as gym_id,
  g.access_type as access_type
from users u
inner join access_grants g
on u.id = g.user_id;
