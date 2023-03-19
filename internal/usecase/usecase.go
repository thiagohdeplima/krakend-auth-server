package usecase

import (
	"context"

	"github.com/thiagohdeplima/krakend-auth-server/internal/issuer"
)

type IssueToken interface {
	Run(_ context.Context, clientId, clientSecret string) (issuer.SucessResponse, error)
}
