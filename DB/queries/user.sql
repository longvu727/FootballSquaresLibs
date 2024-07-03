-- name: CreateUser :execresult
INSERT INTO users (user_guid, ip, device_name, user_name, alias)
    VALUES (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT user_id, user_guid, ip, device_name, user_name, alias
FROM users 
WHERE
  users.user_id = ?