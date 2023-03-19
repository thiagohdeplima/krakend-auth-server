package authenticator

import (
	"context"
)

// ValidateCredentials is the interface implemented by any class used
// to validate if clientID and clientSecret match with existing ones
type ValidateCredentials interface {

	// ValidateCredentials is where the "de facto" validation occurs
	ValidateCredentials(_ context.Context, clientId, clientSecret string) error
}
