-- name: CreateVote :one
insert into votes (
    option_id,voter
) values (
    $1,$2
) returning *;

-- name: GetVotesByPoll :many
select vote_id, v.option_id, voter, option_value, poll_id 
from votes v 
inner join options o on v.option_id = o.option_id 
where o.poll_id = $1
limit $2
offset $3;

-- name: GetVotesByUser :many
select * from votes
where voter ilike $1
limit $2
offset $3;

-- name: GetTotalVotes :many
select o.option_value, count(v.vote_id) as vote_count
from options o
left join votes v on o.option_id = v.option_id
where o.poll_id = $1
group by o.option_id;

