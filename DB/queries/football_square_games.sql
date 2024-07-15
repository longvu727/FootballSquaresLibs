-- name: CreateFootballSquareGame :execlastid
INSERT INTO football_square_games (
  game_id, square_id, row_index, column_index
) VALUES (
  ?, ?, ?, ?
);

-- name: GetFootballSquareGameByGameID :many
SELECT football.football_square_game_id, football.game_id, football.square_id,
        football.user_id, users.user_name, users.alias, football.winner,
        football.winner_quarter_number, football.row_index, football.column_index
FROM football_square_games football 
  left join users on(football.user_id = users.user_id)
WHERE
  football.game_id = ?;


-- name: GetFootballSquareGame :one
SELECT football_square_game_id, game_id, square_id, user_id, winner, winner_quarter_number, row_index, column_index
FROM football_square_games football 
WHERE
  football.football_square_game_id = ?;

-- name: ReserveFootballSquareByGameIDRowIndexColumnIndex :exec
UPDATE football_square_games football
SET football.user_id = ?
WHERE football.game_id = ?
  AND football.row_index = ?
  AND football.column_index = ?
  AND football.user_id is NULL;