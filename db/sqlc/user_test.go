package db

import (
	"context"
	"testing"
	"time"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		HashedPassword:    util.RandomName(8),
		FullName:          util.RandomName(8),
		Email:             util.RandomName(8),
		PasswordChangedAt: time.Now().UTC(),
		CreatedAt:         time.Now().UTC(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.WithinDuration(t, arg.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, arg.CreatedAt, user.CreatedAt, time.Second)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser0(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser0(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUser3(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser3(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)

	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserParams{
		Username:          util.RandomName(8),
		HashedPassword:    util.RandomName(8),
		FullName:          util.RandomName(8),
		Email:             util.RandomName(8),
		PasswordChangedAt: time.Now().UTC(),
		CreatedAt:         time.Now().UTC(),
	}
	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.Email, user2.Email)
	require.WithinDuration(t, arg.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, arg.CreatedAt, user2.CreatedAt, time.Second)

}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)
	user2, err := testStore.GetUser0(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)

}
