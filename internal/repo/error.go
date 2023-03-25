package repo

import (
	"fmt"
)

type KeyNotFoundError struct{}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprint("required key not found")
}
