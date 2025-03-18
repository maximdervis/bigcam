-- name: SelectUserInfo :one
select
  first_name,
  last_name
from users u
where id = $1
limit 1;
