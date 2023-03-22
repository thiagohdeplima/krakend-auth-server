package repository

import "os"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetSecretByClientID(string) (string, error) {
	return "abc123", nil
}

func (r *Repository) GetKeypairByClientID(string) (privkey, pubkey string, _ error) {
	priKey, _ := os.ReadFile("../../test/data/id_ecdsa")
	pubKey, _ := os.ReadFile("../../test/data/id_ecdsa.pub")

	return string(priKey), string(pubKey), nil
}
