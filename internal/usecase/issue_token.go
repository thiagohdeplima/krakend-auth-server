package usecase

import (
	"context"

	"github.com/thiagohdeplima/krakend-auth-server/internal/authenticator"
	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
)

type TokenIssuer struct {
	v authenticator.ValidateCredentials
	e issuer.EmitToken
}

func NewTokenIssuer(v authenticator.ValidateCredentials, e issuer.EmitToken) *TokenIssuer {
	return &TokenIssuer{v, e}
}

func (t *TokenIssuer) Run(ctx context.Context, clientId, clientSecret string) (a issuer.SucessResponse, _ error) {
	if err := t.v.ValidateCredentials(ctx, clientId, clientSecret); err != nil {
		return a, err
	}

	return t.e.EmitToken(clientId)
}
