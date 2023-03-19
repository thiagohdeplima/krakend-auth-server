package issuer

import (
	"fmt"
)

type ClientWithoutKeyPairError struct{}

func (e *ClientWithoutKeyPairError) Error() string {
	return fmt.Sprint("the client doesn't have a keypair")
}
