package db

import (
	"context"
	"importer-api/util"

	"testing"

	"github.com/stretchr/testify/require"
)

// createRandomUser creates a random user for testing purposes.
func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomString(6),
		Password: util.RandomString(12),
		Email:    util.RamdomEmail(11),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)

}


func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}
