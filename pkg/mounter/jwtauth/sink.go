package jwtauth

// CredentialSink delivers a freshly fetched STS credential to the consumer of
// a mount. Apply is called for the initial credential and again on every
// rotation; Cleanup is called once after the Refresher has stopped.
//
// Implementations choose the delivery mechanism appropriate for their mount
// client: ExecSink pushes credentials to a live alinas mount via the vendor
// refresh command, while the interceptors package provides a file-based sink
// for FUSE clients that read credential files from disk.
type CredentialSink interface {
	Apply(cred *STSToken) error
	Cleanup()
}
