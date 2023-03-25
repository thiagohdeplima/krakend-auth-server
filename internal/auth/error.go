package auth

import "fmt"

// InvalidCredentialsError indicates that the clientID
// and the clientSecret doesn't match with any existing
type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("the clientID/clientSecret doesn't match")
}
