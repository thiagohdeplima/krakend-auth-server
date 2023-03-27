package repo

type KeyNotFoundError struct{}

func (e *KeyNotFoundError) Error() string {
	return "required key not found"
}
