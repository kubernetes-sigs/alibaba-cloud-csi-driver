package server

// OssfsMountResult contains the result of an ossfs mount operation.
type OssfsMountResult struct {
	PID      int
	ExitChan chan error
}
