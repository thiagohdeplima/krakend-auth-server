package issuer

import (
	"crypto/ecdsa"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

const TOKEN_TTL = 30

// TokenIssuer is the interface responsible for
// issue and sign a new JWT token and return it
type EmitToken interface {
	EmitToken(clientId string) (SucessResponse, error)
}

// Repository is the interface that must be implemented
// by any keypair datasource
type IssuerRepository interface {
	GetKeypairByClientID(string) (privkey, pubkey string, _ error)
}

type TokenEmissor struct {
	repo IssuerRepository
}

func NewTokenEmissor(repo IssuerRepository) *TokenEmissor {
	return &TokenEmissor{repo}
}

func (te *TokenEmissor) EmitToken(clientId string) (SucessResponse, error) {
	privkey, pubkey, err := te.repo.GetKeypairByClientID(clientId)

	if err != nil {
		return SucessResponse{}, err
	}

	if pubkey == "" || privkey == "" {
		return SucessResponse{}, &ClientWithoutKeyPairError{}
	}

	// from here, fill with key signature validation
	return te.signToken(privkey, pubkey, clientId)
}

func (te *TokenEmissor) signToken(privKey, pubKey, clientId string) (SucessResponse, error) {
	var exp = time.Now().Add(TOKEN_TTL * time.Minute).Unix()

	// TODO: add iss claim in order to help it to get JWKS URL
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"sub": clientId, "exp": exp,
	})

	parsedKey, err := ssh.ParseRawPrivateKey([]byte(privKey))

	if err != nil {
		return SucessResponse{}, err
	}

	if token, err := token.SignedString(parsedKey.(*ecdsa.PrivateKey)); err != nil {
		return SucessResponse{}, err
	} else {
		return SucessResponse{TokenType: "Bearer", AccessToken: token, ExpiresIn: TOKEN_TTL * 60}, nil
	}
}
