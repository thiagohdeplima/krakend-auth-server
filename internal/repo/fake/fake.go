package fake

import (
	"fmt"
	"os"

	"github.com/thiagohdeplima/krakend-auth-server/internal/repo"
)

type FakeRepository struct{}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{}
}

func (r *FakeRepository) GetSecretByClientID(clientId string) (string, error) {
	if clientId == "" {
		return "", &repo.KeyNotFoundError{}
	}

	return "abc123", nil
}

func (r *FakeRepository) GetKeypairByClientID(string) (privkey, pubkey string, _ error) {

	priKey, _ := os.ReadFile("/etc/krakend/test/data/id_ecdsa")
	pubKey, _ := os.ReadFile("/etc/krakend/test/data/id_ecdsa.pub")

	fmt.Println(priKey, pubKey)

	return string(priKey), string(pubKey), nil
}
