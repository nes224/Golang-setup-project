// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_pool.sql

package db

import (
	"context"
)

const createUserPool = `-- name: CreateUserPool :one
INSERT INTO user_pool (
  username,
  login_retry_count,
  duplicated_login_count
) VALUES (
  $1, $2, $3
) RETURNING username, login_retry_count, duplicated_login_count
`

type CreateUserPoolParams struct {
	Username             string `json:"username"`
	LoginRetryCount      int32  `json:"login_retry_count"`
	DuplicatedLoginCount int32  `json:"duplicated_login_count"`
}

func (q *Queries) CreateUserPool(ctx context.Context, arg CreateUserPoolParams) (UserPool, error) {
	row := q.db.QueryRow(ctx, createUserPool, arg.Username, arg.LoginRetryCount, arg.DuplicatedLoginCount)
	var i UserPool
	err := row.Scan(&i.Username, &i.LoginRetryCount, &i.DuplicatedLoginCount)
	return i, err
}

const getUserPool = `-- name: GetUserPool :one
SELECT username, login_retry_count, duplicated_login_count FROM user_pool
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserPool(ctx context.Context, username string) (UserPool, error) {
	row := q.db.QueryRow(ctx, getUserPool, username)
	var i UserPool
	err := row.Scan(&i.Username, &i.LoginRetryCount, &i.DuplicatedLoginCount)
	return i, err
}
