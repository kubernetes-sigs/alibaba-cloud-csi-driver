package cleanup

// Steps performs deferred cleanup on condition that an error
// was returned in the caller. This simplifies code where earlier
// steps need to be undone if a later step fails.  It is not currently
// resilient to panics as this library is not expected to panic.
type Steps []func() error

// Add appends the given cleanup function to those that will be called
// if an error occurs.
func (fns *Steps) Add(fn func() error) {
	*fns = append(*fns, fn)
}

// Unwind calls the cleanup funcions in LIFO order. It panics
// if any of them return an error as failure during recovery is
// itself unrecoverable.
func (fns *Steps) Unwind() {
	// There was some error. Preform cleanup and return. If any of
	// the cleanup functions return and error, we do the only
	// sensible thing and panic.
	for _, fn := range *fns {
		defer func(clean func() error) { checkError(clean) }(fn)
	}
}

// checkError calls `fn` and panics if it returns an error.
func checkError(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}
