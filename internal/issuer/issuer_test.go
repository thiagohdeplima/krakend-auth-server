package issuer_test

import (
	"errors"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
	"github.com/thiagohdeplima/krakend-auth-server/mocks"
)

func Test_TokenEmissor_EmitToken(t *testing.T) {
	var clientID = "clientID"

	t.Run("when repository returns error return error", func(t *testing.T) {
		var repo = mocks.NewIssuerRepository(t)
		var issr = issuer.NewTokenEmissor(repo)
		var err = errors.New("It is an error")

		repo.
			On("GetKeypairByClientID", clientID).
			Return("", "", err)

		_, actualErr := issr.EmitToken(clientID)

		assert.ErrorIs(t, actualErr, err)
	})

	t.Run("when private key is empty return error", func(t *testing.T) {
		var repo = mocks.NewIssuerRepository(t)
		var issr = issuer.NewTokenEmissor(repo)

		repo.
			On("GetKeypairByClientID", clientID).
			Return("", "not-empty", nil)

		_, actualErr := issr.EmitToken(clientID)

		assert.ErrorIs(t, actualErr, &issuer.ClientWithoutKeyPairError{})
	})

	t.Run("when private key is empty return error", func(t *testing.T) {
		var repo = mocks.NewIssuerRepository(t)
		var issr = issuer.NewTokenEmissor(repo)

		repo.
			On("GetKeypairByClientID", clientID).
			Return("not-empty", "", nil)

		_, actualErr := issr.EmitToken(clientID)

		assert.ErrorIs(t, actualErr, &issuer.ClientWithoutKeyPairError{})
	})

	t.Run("when repo returns valid keys return sucessful response", func(t *testing.T) {
		var repo = mocks.NewIssuerRepository(t)
		var issr = issuer.NewTokenEmissor(repo)

		priKey, _ := os.ReadFile("../../test/data/id_ecdsa")
		pubKey, _ := os.ReadFile("../../test/data/id_ecdsa.pub")

		repo.
			On("GetKeypairByClientID", clientID).
			Return(string(priKey), string(pubKey), nil)

		sucessful, actualErr := issr.EmitToken(clientID)

		assert.NoError(t, actualErr)
		assert.Regexp(t, regexp.MustCompile(`^((\w|\-)+\.){2}(\w|\-)+$`), sucessful.AccessToken)
		assert.Equal(t, sucessful.ExpiresIn, issuer.TOKEN_TTL*60)
		assert.Equal(t, sucessful.TokenType, "Bearer")
	})
}
