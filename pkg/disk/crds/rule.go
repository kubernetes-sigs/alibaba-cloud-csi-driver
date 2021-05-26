package crds

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GVR is rules GroupVersionResource
var GVR = schema.GroupVersionResource{
	Group:    "storage.alibabacloud.com",
	Version:  "v1alpha1",
	Resource: "rules",
}

// Rule ...
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RuleSpec   `json:"spec,omitempty"`
	Status RuleStatus `json:"status,omitempty"`
}

// RuleSpec ...
type RuleSpec struct {
	ContainerName string       `json:"containerName"`
	Rules         []RuleAction `json:"rules"`
}

// RuleAction defines rule actions
type RuleAction struct {
	Type  string `json:"actionType"`
	Value string `json:"action"`
}

// RuleStatus ...
type RuleStatus struct {
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty" protobuf:"bytes,4,opt,name=lastScheduleTime"`
}

// RuleList ...
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Rule `json:"items"`
}
