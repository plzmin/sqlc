// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package querytest

import (
	"context"
)

const textArray = `-- name: TextArray :many
SELECT tags FROM bar
`

func (q *Queries) TextArray(ctx context.Context) ([][][]string, error) {
	rows, err := q.db.Query(ctx, textArray)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][][]string
	for rows.Next() {
		var tags [][]string
		if err := rows.Scan(&tags); err != nil {
			return nil, err
		}
		items = append(items, tags)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}