package authenticator

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagohdeplima/krakend-auth-server/internal/repository"
	"github.com/thiagohdeplima/krakend-auth-server/mocks"
)

func Test_Authenticator_ValidateCredentials(t *testing.T) {
	var clientID = "clientID"
	var clientSecret = "clientSecret"

	t.Run("when clientID doesn't exists returns error", func(t *testing.T) {
		var repo = mocks.NewRepository(t)
		var auth = NewAuthenticator(repo)

		repo.
			On("GetSecretByClientID", clientID).
			Return("", &repository.KeyNotFoundError{})

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.ErrorIs(t, actualErr, &InvalidCredentialsError{})
	})

	t.Run("when clientSecret doen't match return error", func(t *testing.T) {
		var repo = mocks.NewRepository(t)
		var auth = NewAuthenticator(repo)
		var wrong = "A different client secret"

		repo.
			On("GetSecretByClientID", clientID).
			Return("another", nil)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, wrong)

		assert.ErrorIs(t, actualErr, &InvalidCredentialsError{})
	})

	t.Run("when repository returns error return the error", func(t *testing.T) {
		var repo = mocks.NewRepository(t)
		var auth = NewAuthenticator(repo)
		var err = errors.New("a random error")

		repo.
			On("GetSecretByClientID", clientID).
			Return("", err)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.ErrorIs(t, actualErr, err)
	})

	t.Run("when credentials match return no error", func(t *testing.T) {
		var repo = mocks.NewRepository(t)
		var auth = NewAuthenticator(repo)

		repo.
			On("GetSecretByClientID", clientID).
			Return(clientSecret, nil)

		actualErr := auth.ValidateCredentials(context.TODO(), clientID, clientSecret)

		assert.NoError(t, actualErr)
	})
}
