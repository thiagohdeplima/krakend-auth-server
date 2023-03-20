package repository

type Repository interface {
	GetSecretByClientID(string) (string, error)
}
