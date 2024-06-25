-- name: CreateFootballSquareGame :execlastid
INSERT INTO football_square_games (
  game_id, square_id, row_index, column_index
) VALUES (
  ?, ?, ?, ?
);

-- name: GetFootballSquareGameByGameID :many
SELECT game_id, square_id, user_id, winner, winner_quarter_number, row_index, column_index
FROM football_square_games football 
WHERE
  football.game_id = ?