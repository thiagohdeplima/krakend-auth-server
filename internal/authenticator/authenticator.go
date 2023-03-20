package authenticator

import (
	"context"

	"github.com/thiagohdeplima/krakend-auth-server/internal/repository"
)

// ValidateCredentials is the interface implemented by any class used
// to validate if clientID and clientSecret match with existing ones
type ValidateCredentials interface {

	// ValidateCredentials is where the "de facto" validation occurs
	ValidateCredentials(_ context.Context, clientId, clientSecret string) error
}

// Repository is the interface that must be implemented
// by clientId and ClientSecret datasources
type Repository interface {
	GetSecretByClientID(string) (string, error)
}

type Authenticator struct {
	repo Repository
}

func NewAuthenticator(r Repository) *Authenticator {
	return &Authenticator{r}
}

func (a Authenticator) ValidateCredentials(_ context.Context, clientId, clientSecret string) error {
	gottenSecret, err := a.repo.GetSecretByClientID(clientId)

	if err != nil {
		switch err.(type) {
		default:
			return err

		case *repository.KeyNotFoundError:
			return &InvalidCredentialsError{}
		}
	}

	if gottenSecret == clientSecret {
		return nil
	}

	return &InvalidCredentialsError{}
}
