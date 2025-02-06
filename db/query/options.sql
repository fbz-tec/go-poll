-- name: CreateOption :one
insert into options (
    option_value,poll_id
) values (
    $1,$2
) returning *;

-- name: GetOptions :many
select * from options
where poll_id = $1 order by option_id;