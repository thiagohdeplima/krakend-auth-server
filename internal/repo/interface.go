package repo

// Repository is the interface that must be implemented
// by any keypair datasource
type KeypairRepository interface {
	GetKeypairByClientID(string) (privkey, pubkey string, _ error)
}

// Repository is the interface that must be implemented
// by clientId and ClientSecret datasources
type ClientSecretRepository interface {
	GetSecretByClientID(string) (string, error)
}
