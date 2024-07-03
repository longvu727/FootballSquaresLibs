-- name: CreateSquare :execlastid
INSERT INTO squares (square_guid, square_size) VALUES (?, ?);

-- name: GetSquare :one
SELECT square_id, square_guid, square_size, row_points, column_points
FROM squares
WHERE square_id = ?;
