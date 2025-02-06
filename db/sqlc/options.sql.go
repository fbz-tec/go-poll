// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: options.sql

package db

import (
	"context"
)

const createOption = `-- name: CreateOption :one
insert into options (
    option_value,poll_id
) values (
    $1,$2
) returning option_id, option_value, poll_id
`

type CreateOptionParams struct {
	OptionValue string `json:"option_value"`
	PollID      int64  `json:"poll_id"`
}

func (q *Queries) CreateOption(ctx context.Context, arg CreateOptionParams) (Option, error) {
	row := q.db.QueryRow(ctx, createOption, arg.OptionValue, arg.PollID)
	var i Option
	err := row.Scan(&i.OptionID, &i.OptionValue, &i.PollID)
	return i, err
}

const getOptions = `-- name: GetOptions :many
select option_id, option_value, poll_id from options
where poll_id = $1 order by option_id
`

func (q *Queries) GetOptions(ctx context.Context, pollID int64) ([]Option, error) {
	rows, err := q.db.Query(ctx, getOptions, pollID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Option{}
	for rows.Next() {
		var i Option
		if err := rows.Scan(&i.OptionID, &i.OptionValue, &i.PollID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
