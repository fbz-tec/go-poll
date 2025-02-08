# GO-Poll

## Polls API
| Method | URL Pattern         | Handler             | Action                                 |
|--------|---------------------|---------------------|----------------------------------------|
| GET    | /api/v1/healthcheck | healthcheckHandler  | Show application information           |
| GET    | /api/v1/polls       | ListPollsHandler    | Show all created polls                 |
| POST   | /api/v1/polls       | CreatePollHander    | Create a new poll                      |
| GET    | /api/v1/polls/:pollId  | GetPollHander       | Show the details of a poll and options |
| GET    | /api/v1/polls/:pollId/votes       | ListVotesHandler    | Show all poll votes      |
| GET    | /api/v1/polls/:pollId/results   | GetVoteResultHandler       | Show the current votes result for poll |

## Votes API
| Method    | URL Pattern                | Handler                | Action                          |
|-----------|----------------------------|------------------------|---------------------------------|
| POST      | /api/v1/votes              | CreateVoteHandler      | Create a new vote               |



## Users API
| Method | URL Pattern         | Handler             | Action                                 |
|--------|---------------------|---------------------|----------------------------------------|
| POST    | /api/v1/users      | CreateUserHandler   | Create a new user                      |
| GET    | /api/v1/users/:user_id/votes       | GetVoteByUserHandler    | Show all user votes polls                 |