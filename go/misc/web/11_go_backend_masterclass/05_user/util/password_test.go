package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashed1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed1)

	err = CheckPassword(password, hashed1)
	require.NoError(t, err)

	wrong := RandomString(6)
	err = CheckPassword(wrong, hashed1)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword)

	hashed2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed2)
	require.NotEqual(t, hashed1, hashed2)
}
