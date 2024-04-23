-- name: CreateUserPool :one
INSERT INTO user_pool (
  username,
  login_retry_count,
  duplicated_login_count
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserPool :one
SELECT * FROM user_pool
WHERE username = $1 LIMIT 1;
