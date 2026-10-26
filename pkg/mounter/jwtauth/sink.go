package jwtauth

// CredentialSink delivers a freshly fetched STS credential to the consumer of
// a mount. Apply is called for the initial credential and again on every
// rotation; Cleanup is called once after the Refresher has stopped.
type CredentialSink interface {
	Apply(cred *STSToken) error
}
