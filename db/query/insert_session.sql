-- name: InsertSession :exec
insert into sessions (user_id, camera_id)
values ($1, $2);

