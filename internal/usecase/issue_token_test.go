package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagohdeplima/krakend-auth-server/internal/authenticator"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
	"github.com/thiagohdeplima/krakend-auth-server/mocks"
)

func Test_TokenIssuer_Run(t *testing.T) {
	var cid = "clientID"
	var cst = "clientSecret"
	var ctx = context.Background()

	t.Run("when ValidateCredentials returns error return error", func(t *testing.T) {
		auth := mocks.NewValidateCredentials(t)
		issr := mocks.NewTokenIssuer(t)
		target := NewTokenIssuer(auth, issr)

		auth.
			On("ValidateCredentials", ctx, cid, cst).
			Return(&authenticator.InvalidCredentialsError{})

		_, actualErr := target.Run(ctx, cid, cst)

		assert.ErrorIs(t, actualErr, &authenticator.InvalidCredentialsError{})
	})

	t.Run("when TokenIssuer returns error return error", func(t *testing.T) {
		auth := mocks.NewValidateCredentials(t)
		issr := mocks.NewEmitToken(t)
		target := NewTokenIssuer(auth, issr)

		auth.
			On("ValidateCredentials", ctx, cid, cst).
			Return(nil)

		issr.
			On("EmitToken", cid).
			Return(issuer.SucessResponse{}, &issuer.ClientWithoutKeyPairError{})

		_, actualErr := target.Run(ctx, cid, cst)

		assert.ErrorIs(t, actualErr, &issuer.ClientWithoutKeyPairError{})
	})
}
