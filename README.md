# GO-Poll

## Poll API
| Method | URL Pattern         | Handler             | Action                                 |
|--------|---------------------|---------------------|----------------------------------------|
| GET    | /api/v1/healthcheck | healthcheckHandler  | Show application information           |
| GET    | /api/v1/polls       | ListPollsHandler    | Show all created polls                 |
| POST   | /api/v1/polls       | CreatePollHander    | Create a new poll                      |
| GET    | /api/v1/polls/:id   | GetPollHander       | Show the details of a poll and options |
| PUT    | /api/v1/movies/:id  | updateMovieHandler  | Update the details of a specific movie |
| DELETE | /api/v1/movies/:id  | deleteMovieHandler  | Delete a specific movie                |
