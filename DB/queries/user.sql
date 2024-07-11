-- name: CreateUser :execlastid
INSERT INTO users (user_guid, ip, device_name, user_name, alias)
    VALUES (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT user_id, user_guid, ip, device_name, user_name, alias
FROM users 
WHERE
  users.user_id = ?;  

-- name: GetUserByGUID :one
SELECT user_guid, ip, device_name, user_name, alias
FROM users 
WHERE
  users.user_guid = ?;