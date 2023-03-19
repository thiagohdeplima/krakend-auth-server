package issuer

// SucessResponse represents an
// OAuth2 sucessfully response
type SucessResponse struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int
	Scope       string
}
