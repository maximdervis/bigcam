-- name: SelectUserInfo :one
select
  name,
  email,
  dob,
  avatar_id
from users u
where id = $1
limit 1;

-- name: SelectUserInfoByEmail :one
select
  id,
  name,
  email,
  password
from users
where email = $1;

-- name: ContainsUserWithEmail :one
select
  count(email) <> 0
from users u
where email = $1;

-- name: InsertUserInfo :exec
insert into users (name, email, password)
values ($1, $2, $3);

-- name: UpdateUserInfo :exec
WITH input_data AS (
  SELECT
    sqlc.arg(update_data)::jsonb as data
)
UPDATE users
SET
  email = COALESCE(input_data.data->>'email', email),
  name = COALESCE(input_data.data->>'name', name),
  dob = COALESCE((input_data.data->>'dob')::date, dob),
  avatar_id = COALESCE((input_data.data->>'avatar_id')::int, avatar_id),
  updated_at = now()
FROM input_data
WHERE id = $1;
