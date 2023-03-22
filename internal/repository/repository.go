package repository

import (
	"fmt"
	"os"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetSecretByClientID(string) (string, error) {
	return "abc123", nil
}

func (r *Repository) GetKeypairByClientID(string) (privkey, pubkey string, _ error) {
	fmt.Println(os.Getwd())
	priKey, _ := os.ReadFile("/etc/krakend/test/data/id_ecdsa")
	pubKey, _ := os.ReadFile("/etc/krakend/test/data/id_ecdsa.pub")

	fmt.Println(priKey, pubKey)

	return string(priKey), string(pubKey), nil
}
