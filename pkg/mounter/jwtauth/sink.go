package jwtauth

// CredentialSink delivers a freshly fetched STS credential to the consumer of
// a mount. Deliver persists or pushes the credential; Notify tells the client
// that a rotation happened (e.g. by running a refresh hook or command).
// Cleanup is called once after the Refresher has stopped.
//
// Separating Deliver from Notify lets the refresh loop advance its expiry
// tracking as soon as the credential is persisted, even when notification
// fails. A failing hook therefore triggers backoff-limited retries instead of
// redundant STS exchanges.
type CredentialSink interface {
	Deliver(cred *STSToken) error
	Notify() error
	Cleanup()
}
