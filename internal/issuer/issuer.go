package issuer

// TokenIssuer is the interface responsible for
// issue and sign a new JWT token and return it
type EmitToken interface {
	EmitToken(clientId string) (SucessResponse, error)
}
