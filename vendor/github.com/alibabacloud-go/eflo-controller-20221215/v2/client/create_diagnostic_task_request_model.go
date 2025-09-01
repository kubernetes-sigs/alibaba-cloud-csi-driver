// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiJobLogInfo(v *CreateDiagnosticTaskRequestAiJobLogInfo) *CreateDiagnosticTaskRequest
	GetAiJobLogInfo() *CreateDiagnosticTaskRequestAiJobLogInfo
	SetClusterId(v string) *CreateDiagnosticTaskRequest
	GetClusterId() *string
	SetDiagnosticType(v string) *CreateDiagnosticTaskRequest
	GetDiagnosticType() *string
	SetNodeIds(v []*string) *CreateDiagnosticTaskRequest
	GetNodeIds() []*string
}

type CreateDiagnosticTaskRequest struct {
	// The log information.
	AiJobLogInfo *CreateDiagnosticTaskRequestAiJobLogInfo `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The diagnostics type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The IDs of the nodes.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s CreateDiagnosticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequest) GetAiJobLogInfo() *CreateDiagnosticTaskRequestAiJobLogInfo {
	return s.AiJobLogInfo
}

func (s *CreateDiagnosticTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDiagnosticTaskRequest) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticTaskRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *CreateDiagnosticTaskRequest) SetAiJobLogInfo(v *CreateDiagnosticTaskRequestAiJobLogInfo) *CreateDiagnosticTaskRequest {
	s.AiJobLogInfo = v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetClusterId(v string) *CreateDiagnosticTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetNodeIds(v []*string) *CreateDiagnosticTaskRequest {
	s.NodeIds = v
	return s
}

func (s *CreateDiagnosticTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDiagnosticTaskRequestAiJobLogInfo struct {
	// The task logs.
	AiJobLogs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs `json:"AiJobLogs,omitempty" xml:"AiJobLogs,omitempty" type:"Repeated"`
	// The end time. The value is in the timestamp format. Unit: seconds.
	//
	// >  This timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time. The value is in the timestamp format. Unit: seconds.
	//
	// >  This timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 2024-10-11T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetAiJobLogs() []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	return s.AiJobLogs
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetAiJobLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.AiJobLogs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetEndTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetStartTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) Validate() error {
	return dara.Validate(s)
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs struct {
	// The instance ID.
	//
	// example:
	//
	// null
	AiInstance *string `json:"AiInstance,omitempty" xml:"AiInstance,omitempty"`
	// The logs.
	Logs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-p2p2al5u1hn
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetAiInstance() *string {
	return s.AiInstance
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetLogs() []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	return s.Logs
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetAiInstance(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.AiInstance = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.Logs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetNodeId(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.NodeId = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) Validate() error {
	return dara.Validate(s)
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs struct {
	// The sending date in the yyyymmdd format.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	Datetime *string `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	// The log content.
	//
	// example:
	//
	// success
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GetDatetime() *string {
	return s.Datetime
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GetLogContent() *string {
	return s.LogContent
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetDatetime(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.Datetime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetLogContent(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.LogContent = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) Validate() error {
	return dara.Validate(s)
}
