package db

import (
	"context"
	"testing"
	"time"

	"github.com/abs/bestinter/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.RandomString(6))

	arg := CreateUserParams{
		Username:        util.RandomOwner(),
		Email:           util.RandomEmail(),
		Fullname:        util.RandomOwner(),
		Password:        hashedPassword,
		Role:            util.RandomRole(),
		Enable:          int32(util.RandomeEnable()),
		CreatedDatetime: time.Now(),
		UpdatedDatetime: time.Now(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.Enable, user.Enable)
	require.NotZero(t, user.CreatedDatetime)

	return user
}
