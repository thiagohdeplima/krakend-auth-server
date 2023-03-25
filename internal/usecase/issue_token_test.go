package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagohdeplima/krakend-auth-server/internal/auth"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
	"github.com/thiagohdeplima/krakend-auth-server/mocks"
)

func Test_TokenIssuer_Run(t *testing.T) {
	var cid = "clientID"
	var cst = "clientSecret"
	var ctx = context.Background()

	t.Run("when ValidateCredentials returns error return error", func(t *testing.T) {
		validator := mocks.NewValidateCredentials(t)
		issr := mocks.NewEmitToken(t)
		target := NewTokenIssuer(validator, issr)

		validator.
			On("ValidateCredentials", ctx, cid, cst).
			Return(&auth.InvalidCredentialsError{})

		_, actualErr := target.Run(ctx, cid, cst)

		assert.ErrorIs(t, actualErr, &auth.InvalidCredentialsError{})
	})

	t.Run("when TokenIssuer returns error return error", func(t *testing.T) {
		validator := mocks.NewValidateCredentials(t)
		issr := mocks.NewEmitToken(t)
		target := NewTokenIssuer(validator, issr)

		validator.
			On("ValidateCredentials", ctx, cid, cst).
			Return(nil)

		issr.
			On("EmitToken", cid).
			Return(issuer.SucessResponse{}, &issuer.ClientWithoutKeyPairError{})

		_, actualErr := target.Run(ctx, cid, cst)

		assert.ErrorIs(t, actualErr, &issuer.ClientWithoutKeyPairError{})
	})
}
