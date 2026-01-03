// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*DescribeDiagnosticMetricsResponseBodyMetrics) *DescribeDiagnosticMetricsResponseBody
	GetMetrics() []*DescribeDiagnosticMetricsResponseBodyMetrics
	SetNextToken(v string) *DescribeDiagnosticMetricsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDiagnosticMetricsResponseBody
	GetRequestId() *string
}

type DescribeDiagnosticMetricsResponseBody struct {
	// The diagnostic metrics.
	Metrics []*DescribeDiagnosticMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricsResponseBody) GetMetrics() []*DescribeDiagnosticMetricsResponseBodyMetrics {
	return s.Metrics
}

func (s *DescribeDiagnosticMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticMetricsResponseBody) SetMetrics(v []*DescribeDiagnosticMetricsResponseBodyMetrics) *DescribeDiagnosticMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBody) SetNextToken(v string) *DescribeDiagnosticMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBody) SetRequestId(v string) *DescribeDiagnosticMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBody) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosticMetricsResponseBodyMetrics struct {
	// The description of the diagnostic metric.
	//
	// example:
	//
	// CPU diagnostic
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the diagnostic metric needs to be assessed by running a Cloud Assistant command in a guest operating system.
	//
	// example:
	//
	// true
	GuestMetric *bool `json:"GuestMetric,omitempty" xml:"GuestMetric,omitempty"`
	// The category of the diagnostic metric.
	//
	// example:
	//
	// CPU
	MetricCategory *string `json:"MetricCategory,omitempty" xml:"MetricCategory,omitempty"`
	// The ID of the diagnostic metric.
	//
	// example:
	//
	// GuestOS.WinFirewall
	MetricId *string `json:"MetricId,omitempty" xml:"MetricId,omitempty"`
	// The name of the diagnostic metric.
	//
	// example:
	//
	// CPU diagnostic
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The resource type supported by the diagnostic metric.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The operating system type supported by the diagnostic metric. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// 	- All: Windows and Linux
	//
	// example:
	//
	// ALL
	SupportedOperatingSystem *string `json:"SupportedOperatingSystem,omitempty" xml:"SupportedOperatingSystem,omitempty"`
}

func (s DescribeDiagnosticMetricsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetGuestMetric() *bool {
	return s.GuestMetric
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetMetricCategory() *string {
	return s.MetricCategory
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetMetricId() *string {
	return s.MetricId
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) GetSupportedOperatingSystem() *string {
	return s.SupportedOperatingSystem
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetDescription(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetGuestMetric(v bool) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.GuestMetric = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetMetricCategory(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.MetricCategory = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetMetricId(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.MetricId = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetMetricName(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetResourceType(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) SetSupportedOperatingSystem(v string) *DescribeDiagnosticMetricsResponseBodyMetrics {
	s.SupportedOperatingSystem = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
