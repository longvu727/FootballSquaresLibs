// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: games.sql

package db

import (
	"context"
	"database/sql"
)

const createGame = `-- name: CreateGame :execlastid
INSERT INTO games (
  game_guid, sport, team_a, team_b
) VALUES (
  ?, ?, ?, ?
)
`

type CreateGameParams struct {
	GameGuid string
	Sport    sql.NullString
	TeamA    sql.NullString
	TeamB    sql.NullString
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createGame,
		arg.GameGuid,
		arg.Sport,
		arg.TeamA,
		arg.TeamB,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getGame = `-- name: GetGame :one
SELECT game_id, game_guid, sport, team_a, team_b
FROM games 
WHERE
  games.game_id = ?
`

type GetGameRow struct {
	GameID   int32
	GameGuid string
	Sport    sql.NullString
	TeamA    sql.NullString
	TeamB    sql.NullString
}

func (q *Queries) GetGame(ctx context.Context, gameID int32) (GetGameRow, error) {
	row := q.db.QueryRowContext(ctx, getGame, gameID)
	var i GetGameRow
	err := row.Scan(
		&i.GameID,
		&i.GameGuid,
		&i.Sport,
		&i.TeamA,
		&i.TeamB,
	)
	return i, err
}

const getGameByGUID = `-- name: GetGameByGUID :one
SELECT game_id, game_guid, sport, team_a, team_b
FROM games 
WHERE
  games.game_guid = ?
`

type GetGameByGUIDRow struct {
	GameID   int32
	GameGuid string
	Sport    sql.NullString
	TeamA    sql.NullString
	TeamB    sql.NullString
}

func (q *Queries) GetGameByGUID(ctx context.Context, gameGuid string) (GetGameByGUIDRow, error) {
	row := q.db.QueryRowContext(ctx, getGameByGUID, gameGuid)
	var i GetGameByGUIDRow
	err := row.Scan(
		&i.GameID,
		&i.GameGuid,
		&i.Sport,
		&i.TeamA,
		&i.TeamB,
	)
	return i, err
}
