-- name: SelectUserInfo :one
select
  name,
  email,
  password
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
