-- name: InsertGym :exec
insert into gyms (name, auth_key)
values ($1, $2);

