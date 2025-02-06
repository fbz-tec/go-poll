-- name: CreateUser :one
insert into users (
   username,"role",hashed_password,full_name,email
) values (
    $1,$2,$3,$4,$5
)
returning *;

-- name: GetUserByUserName :one
select * from users
where username ilike $1;

-- name: GetUserByUserEmail :one
select * from users
where email ilike $1;