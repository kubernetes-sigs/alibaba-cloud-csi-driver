package jwtauth

// CredentialSink delivers a freshly fetched STS credential to the consumer of
// a mount. Apply is called for the initial credential and again on every
// rotation; Cleanup is called once after the Refresher has stopped.
//
// This is the seam where the storage-specific part of the credential flow
// lives: how a credential is handed to a running mount differs per storage
// client, so this package owns only the exchange and the refresh schedule and
// each driver supplies its own sink next to its mount wiring (see the
// interceptors package: NAS/CPFS pushes the rotated credential to the live
// alinas mount by executing the vendor refresh command, while FUSE clients
// read credential files from a directory rotated atomically on disk).
type CredentialSink interface {
	Apply(cred *STSToken) error
	Cleanup()
}
