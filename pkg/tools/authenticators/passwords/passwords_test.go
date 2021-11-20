package passwords_test

import (
	"github.com/SimpleOpenBadge/authentication-api/pkg/tools/authenticators/passwords"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Run("Tests correct password case", func(t *testing.T) {
		password := "asdffdsa"

		hashedPassword, err := passwords.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		err = passwords.CheckPassword(password, hashedPassword)
		require.NoError(t, err)
	})

	t.Run("Tests wrong password case", func(t *testing.T) {
		password := "asdffdsa"
		wrongPassword := "asdffdsafdas"

		hashedPassword, err := passwords.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		err = passwords.CheckPassword(wrongPassword, hashedPassword)
		require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})

	t.Run("Same password must be hashed differently each time", func(t *testing.T) {
		password := "asdffdsa"

		hashedPassword1, err := passwords.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword1)

		hashedPassword2, err := passwords.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword1)

		err = passwords.CheckPassword(password, hashedPassword1)
		require.NoError(t, err)

		err = passwords.CheckPassword(password, hashedPassword2)
		require.NoError(t, err)

		require.NotEqual(t, hashedPassword1, hashedPassword2)
	})

	t.Run("Tests hashed hash", func(t *testing.T) {
		password := "asdffdsa"

		hashedPassword, err := passwords.HashPassword(password)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		err = passwords.CheckPassword(password, hashedPassword)
		require.NoError(t, err)

		hashedHashedPassword, err := passwords.HashPassword(hashedPassword)
		require.NoError(t, err)
		require.NotEmpty(t, hashedPassword)

		err = passwords.CheckPassword(password, hashedHashedPassword)
		require.Error(t, err) // -> Check should fail
		err = passwords.CheckPassword(hashedPassword, hashedHashedPassword)
		require.NoError(t, err)
	})
}
