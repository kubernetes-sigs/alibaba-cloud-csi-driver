package server

type FuseMountResult struct {
	PID      int
	ExitChan chan error
}
