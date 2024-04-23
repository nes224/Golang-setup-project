-- name: CreateUser :one
INSERT INTO "user" (
  username,
  email,
  fullname,
  password,
  role,
  enable,
  created_datetime,
  updated_datetime
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;



