// packate authenticator have all the things that we need
// to perform clientId and clientSecret validation
package auth

import (
	"context"

	repo "github.com/thiagohdeplima/krakend-auth-server/internal/repo"
)

// ValidateCredentials is the interface implemented by any class used
// to validate if clientID and clientSecret match with existing ones
type ValidateCredentials interface {

	// ValidateCredentials is where the "de facto" validation occurs
	ValidateCredentials(_ context.Context, clientId, clientSecret string) error
}

type Authenticator struct {
	repo repo.ClientSecretRepository
}

// NewAuthenticator generates a instance of Authenticator, that implements

func NewAuthenticator(r repo.ClientSecretRepository) *Authenticator {
	return &Authenticator{r}
}

func (a Authenticator) ValidateCredentials(_ context.Context, clientId, clientSecret string) error {
	gottenSecret, err := a.repo.GetSecretByClientID(clientId)

	if err != nil {
		switch err.(type) {
		default:
			return err

		case *repo.KeyNotFoundError:
			return &InvalidCredentialsError{}
		}
	}

	if gottenSecret == clientSecret {
		return nil
	}

	return &InvalidCredentialsError{}
}
