package issuer

type EmptyKeyError struct{}

type ClientWithoutKeyPairError struct{}

func (e *EmptyKeyError) Error() string {
	return "one or both keys are empty"
}

func (e *ClientWithoutKeyPairError) Error() string {
	return "the client doesn't have a keypair"
}
