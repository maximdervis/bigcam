-- name: SelectGymInfo :one
select
  name
from gyms u
where id = $1
limit 1;
