-- name: CreatePoll :one
insert into polls (
    question,"owner"
) values (
    $1,$2
)
returning *;

-- name: ListPolls :many
select * from polls
limit $1
offset $2;

-- name: GetPoll :one
select * from polls 
where poll_id = $1 limit 1;

