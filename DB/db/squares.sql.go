// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: squares.sql

package db

import (
	"context"
	"database/sql"
)

const createSquare = `-- name: CreateSquare :execlastid
INSERT INTO squares (square_guid, square_size) VALUES (?, ?)
`

type CreateSquareParams struct {
	SquareGuid string
	SquareSize sql.NullInt32
}

func (q *Queries) CreateSquare(ctx context.Context, arg CreateSquareParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createSquare, arg.SquareGuid, arg.SquareSize)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getSquare = `-- name: GetSquare :one
SELECT square_id, square_guid, square_size, row_points, column_points
FROM squares
WHERE square_id = ?
`

type GetSquareRow struct {
	SquareID     int32
	SquareGuid   string
	SquareSize   sql.NullInt32
	RowPoints    sql.NullString
	ColumnPoints sql.NullString
}

func (q *Queries) GetSquare(ctx context.Context, squareID int32) (GetSquareRow, error) {
	row := q.db.QueryRowContext(ctx, getSquare, squareID)
	var i GetSquareRow
	err := row.Scan(
		&i.SquareID,
		&i.SquareGuid,
		&i.SquareSize,
		&i.RowPoints,
		&i.ColumnPoints,
	)
	return i, err
}
