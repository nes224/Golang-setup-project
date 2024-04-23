-- name: CreateUserConnection :one
INSERT INTO user_connection (
  username,
  last_connection_date_time,
  last_action_date_time,
  logged_on_date_time,
  expired_date_time,
  session_key,
  client_device,
  created_datetime,
  updated_datetime
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetUserConnection :one
SELECT * FROM user_connection
WHERE username = $1 LIMIT 1;
