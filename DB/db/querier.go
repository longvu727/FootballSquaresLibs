// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateFootballSquareGame(ctx context.Context, arg CreateFootballSquareGameParams) (int64, error)
	CreateGames(ctx context.Context, arg CreateGamesParams) (sql.Result, error)
	CreateSquare(ctx context.Context, arg CreateSquareParams) (sql.Result, error)
	GetFootballSquareGameByGameID(ctx context.Context, gameID sql.NullInt32) ([]GetFootballSquareGameByGameIDRow, error)
	GetGame(ctx context.Context, gameID int32) (GetGameRow, error)
	GetSquare(ctx context.Context, squareID int32) (GetSquareRow, error)
}

var _ Querier = (*Queries)(nil)
