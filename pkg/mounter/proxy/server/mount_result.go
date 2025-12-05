package server

type OssfsMountResult struct {
	PID      int
	ExitChan chan error
}
