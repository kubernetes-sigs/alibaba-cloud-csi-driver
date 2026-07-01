// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyslogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListSyslogsResponseBodyLogs) *ListSyslogsResponseBody
	GetLogs() []*ListSyslogsResponseBodyLogs
	SetNextToken(v string) *ListSyslogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSyslogsResponseBody
	GetRequestId() *string
}

type ListSyslogsResponseBody struct {
	// The object that contains the paginated list of logs.
	Logs []*ListSyslogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The \\`NextToken\\` for the next page. To retrieve the next page of results, include this return value in the next request.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSyslogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponseBody) GetLogs() []*ListSyslogsResponseBodyLogs {
	return s.Logs
}

func (s *ListSyslogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSyslogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSyslogsResponseBody) SetLogs(v []*ListSyslogsResponseBodyLogs) *ListSyslogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListSyslogsResponseBody) SetNextToken(v string) *ListSyslogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSyslogsResponseBody) SetRequestId(v string) *ListSyslogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSyslogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSyslogsResponseBodyLogs struct {
	// The cluster ID.
	//
	// example:
	//
	// i119583961673208491760
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The security domain of the client.
	//
	// example:
	//
	// ALIYUN_PUBLIC
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// A standard field in Linux syslog.
	//
	// example:
	//
	// kern
	Facility *string `json:"Facility,omitempty" xml:"Facility,omitempty"`
	// The hostname of the client.
	//
	// example:
	//
	// damo-m53kr8kd-0008
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 114.55.254.44
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The content of the syslog.
	//
	// example:
	//
	// bond4: failed to get link speed/duplex for eth8
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-9lb36u4s601
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// A standard field in Linux syslog.
	//
	// example:
	//
	// warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The unique identifier of the machine.
	//
	// example:
	//
	// 21A401332
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The hostname of the syslog server.
	//
	// example:
	//
	// damo-m53kr8kd-0008
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// A standard field in Linux syslog.
	//
	// example:
	//
	// kernel
	Syslogtag *string `json:"Syslogtag,omitempty" xml:"Syslogtag,omitempty"`
	// Tag: Hostname.
	//
	// example:
	//
	// damo-m53kr8kd-0008
	TagHostname *string `json:"TagHostname,omitempty" xml:"TagHostname,omitempty"`
	// Tag: Package ID.
	//
	// example:
	//
	// D990314D3C25D7E8-1080
	TagPackId *string `json:"TagPackId,omitempty" xml:"TagPackId,omitempty"`
	// Tag: Path.
	//
	// example:
	//
	// /var/log/kern
	TagPath *string `json:"TagPath,omitempty" xml:"TagPath,omitempty"`
	// Tag: Receive time.
	//
	// example:
	//
	// 1687363348
	TagReceiveTime *string `json:"TagReceiveTime,omitempty" xml:"TagReceiveTime,omitempty"`
	// Tag: User-defined ID.
	//
	// example:
	//
	// application_b
	TagUserDefinedId *string `json:"TagUserDefinedId,omitempty" xml:"TagUserDefinedId,omitempty"`
	// The time in UNIX format.
	//
	// example:
	//
	// 1687363346
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic. Valid values are \\`logserver\\` and \\`log2sls\\`. Note: Both are syslog servers. \\`logserver\\` collects data over TCP. \\`log2sls\\` collects data over UDP.
	//
	// example:
	//
	// logserver
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListSyslogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponseBodyLogs) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSyslogsResponseBodyLogs) GetDomain() *string {
	return s.Domain
}

func (s *ListSyslogsResponseBodyLogs) GetFacility() *string {
	return s.Facility
}

func (s *ListSyslogsResponseBodyLogs) GetHostname() *string {
	return s.Hostname
}

func (s *ListSyslogsResponseBodyLogs) GetIp() *string {
	return s.Ip
}

func (s *ListSyslogsResponseBodyLogs) GetMsg() *string {
	return s.Msg
}

func (s *ListSyslogsResponseBodyLogs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSyslogsResponseBodyLogs) GetSeverity() *string {
	return s.Severity
}

func (s *ListSyslogsResponseBodyLogs) GetSn() *string {
	return s.Sn
}

func (s *ListSyslogsResponseBodyLogs) GetSource() *string {
	return s.Source
}

func (s *ListSyslogsResponseBodyLogs) GetSyslogtag() *string {
	return s.Syslogtag
}

func (s *ListSyslogsResponseBodyLogs) GetTagHostname() *string {
	return s.TagHostname
}

func (s *ListSyslogsResponseBodyLogs) GetTagPackId() *string {
	return s.TagPackId
}

func (s *ListSyslogsResponseBodyLogs) GetTagPath() *string {
	return s.TagPath
}

func (s *ListSyslogsResponseBodyLogs) GetTagReceiveTime() *string {
	return s.TagReceiveTime
}

func (s *ListSyslogsResponseBodyLogs) GetTagUserDefinedId() *string {
	return s.TagUserDefinedId
}

func (s *ListSyslogsResponseBodyLogs) GetTime() *string {
	return s.Time
}

func (s *ListSyslogsResponseBodyLogs) GetTopic() *string {
	return s.Topic
}

func (s *ListSyslogsResponseBodyLogs) SetClusterId(v string) *ListSyslogsResponseBodyLogs {
	s.ClusterId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetDomain(v string) *ListSyslogsResponseBodyLogs {
	s.Domain = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetFacility(v string) *ListSyslogsResponseBodyLogs {
	s.Facility = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetHostname(v string) *ListSyslogsResponseBodyLogs {
	s.Hostname = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetIp(v string) *ListSyslogsResponseBodyLogs {
	s.Ip = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetMsg(v string) *ListSyslogsResponseBodyLogs {
	s.Msg = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetNodeId(v string) *ListSyslogsResponseBodyLogs {
	s.NodeId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSeverity(v string) *ListSyslogsResponseBodyLogs {
	s.Severity = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSn(v string) *ListSyslogsResponseBodyLogs {
	s.Sn = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSource(v string) *ListSyslogsResponseBodyLogs {
	s.Source = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSyslogtag(v string) *ListSyslogsResponseBodyLogs {
	s.Syslogtag = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagHostname(v string) *ListSyslogsResponseBodyLogs {
	s.TagHostname = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagPackId(v string) *ListSyslogsResponseBodyLogs {
	s.TagPackId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagPath(v string) *ListSyslogsResponseBodyLogs {
	s.TagPath = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagReceiveTime(v string) *ListSyslogsResponseBodyLogs {
	s.TagReceiveTime = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagUserDefinedId(v string) *ListSyslogsResponseBodyLogs {
	s.TagUserDefinedId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTime(v string) *ListSyslogsResponseBodyLogs {
	s.Time = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTopic(v string) *ListSyslogsResponseBodyLogs {
	s.Topic = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
