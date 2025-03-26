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

-- name: UpdateUserEmail :exec
UPDATE users
SET
  email = $1,
  updated_at = now()
WHERE id = $2;

-- name: UpdateUserName :exec
UPDATE users
SET
  name = $1,
  updated_at = now()
WHERE id = $2;

-- name: UpdateUserDob :exec
UPDATE users
SET
  dob = $1,
  updated_at = now()
WHERE id = $2;

-- name: UpdateUserAvatarId :exec
UPDATE users
SET
  avatar_id = $1,
  updated_at = now()
WHERE id = $2;
