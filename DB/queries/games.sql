-- name: CreateGame :execlastid
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

-- name: GetGameByGUID :one
SELECT game_id, game_guid, sport, team_a, team_b
FROM games 
WHERE
  games.game_guid = ?;

