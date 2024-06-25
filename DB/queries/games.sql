-- name: CreateGames :execresult
INSERT INTO games (
  game_guid, sport, team_a, team_b
) VALUES (
  ?, ?, ?, ?
);

-- name: GetGame :one
SELECT game_id, game_guid, sport, team_a, team_b
FROM games 
WHERE
  games.game_id = ?;