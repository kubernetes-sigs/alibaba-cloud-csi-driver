package agent

import restclient "k8s.io/client-go/rest"

// Agent for node server
type Agent struct {
}

// NewAgent new Agent object
func NewAgent() *Agent {
	return &Agent{}
}

// RunAgent run agent
func (ag *Agent) RunAgent(kubeconfig *restclient.Config) {
	// current only query server support
	queryServ := NewQueryServer(kubeconfig)
	queryServ.RunQueryServer()
}
