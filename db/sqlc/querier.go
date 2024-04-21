package db

import (
	"context"

)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)