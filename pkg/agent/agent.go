package agent

// Agent for node server
type Agent struct {
}

// NewAgent new Agent object
func NewAgent() *Agent {
	return &Agent{}
}

// RunAgent run agent
func (ag *Agent) RunAgent() {
	// current only query server support
	queryServ := NewQueryServer()
	queryServ.RunQueryServer()
}
