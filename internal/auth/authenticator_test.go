package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	repo "github.com/thiagohdeplima/krakend-auth-server/internal/repo"
	"github.com/thiagohdeplima/krakend-auth-server/mocks"
)

func Test_Authenticator_ValidateCredentials(t *testing.T) {
	var clientID = "clientID"
	var clientSecret = "clientSecret"

	t.Run("when clientID doesn't exists returns error", func(t *testing.T) {
		var fake = mocks.NewClientSecretRepository(t)
		var auth = NewAuthenticator(fake)

		fake.
			On("GetSecretByClientID", clientID).
			Return("", &repo.KeyNotFoundError{})

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.ErrorIs(t, actualErr, &InvalidCredentialsError{})
	})

	t.Run("when clientSecret doen't match return error", func(t *testing.T) {
		var fake = mocks.NewClientSecretRepository(t)
		var auth = NewAuthenticator(fake)
		var wrong = "A different client secret"

		fake.
			On("GetSecretByClientID", clientID).
			Return("another", nil)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, wrong)

		assert.ErrorIs(t, actualErr, &InvalidCredentialsError{})
	})

	t.Run("when fake returns error return the error", func(t *testing.T) {
		var fake = mocks.NewClientSecretRepository(t)
		var auth = NewAuthenticator(fake)
		var err = errors.New("a random error")

		fake.
			On("GetSecretByClientID", clientID).
			Return("", err)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.ErrorIs(t, actualErr, err)
	})

	t.Run("when credentials match return no error", func(t *testing.T) {
		var fake = mocks.NewClientSecretRepository(t)
		var auth = NewAuthenticator(fake)

		fake.
			On("GetSecretByClientID", clientID).
			Return(clientSecret, nil)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.NoError(t, actualErr)
	})
}
