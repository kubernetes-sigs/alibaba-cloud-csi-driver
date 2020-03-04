package agent

const (
	// CsiPluginRunTimeFlagFile tag
	CsiPluginRunTimeFlagFile = "alibabacloudcsiplugin.json"
	Runc_RunTime_Tag         = "runc"
	Runv_RunTime_Tag         = "runv"
	Unkown_RunTime_Tag       = "unkown"
)

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
