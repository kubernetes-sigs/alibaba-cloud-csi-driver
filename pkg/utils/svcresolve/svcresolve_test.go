// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package svcresolve

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func boolPtr(b bool) *bool { return &b }

func TestResolveServiceIPClusterIP(t *testing.T) {
	client := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "svc-1", Namespace: "ns-1"},
		Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.1"},
	})
	ip, err := ResolveServiceIP(context.Background(), client, "ns-1", "svc-1", false)
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.1", ip)
}

func TestResolveServiceIPHeadlessWithLookup(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "cred-prov", Namespace: "ack-agent-identity"},
			Spec:       corev1.ServiceSpec{ClusterIP: "None"},
		},
		&discoveryv1.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name: "cred-prov-abc", Namespace: "ack-agent-identity",
				Labels: map[string]string{"kubernetes.io/service-name": "cred-prov"},
			},
			Endpoints: []discoveryv1.Endpoint{
				{Addresses: []string{"10.0.0.99"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(true)}},
			},
		},
	)
	ip, err := ResolveServiceIP(context.Background(), client, "ack-agent-identity", "cred-prov", true)
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.99", ip)
}

func TestResolveServiceIPHeadlessWithoutLookup(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "svc-1", Namespace: "ns-1"},
			Spec:       corev1.ServiceSpec{ClusterIP: "None"},
		},
	)
	_, err := ResolveServiceIP(context.Background(), client, "ns-1", "svc-1", false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "has no ClusterIP")
}

func TestResolveServiceIPHeadlessSkipsNotReady(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "svc-1", Namespace: "ns-1"},
			Spec:       corev1.ServiceSpec{ClusterIP: "None"},
		},
		&discoveryv1.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name: "es-1", Namespace: "ns-1",
				Labels: map[string]string{"kubernetes.io/service-name": "svc-1"},
			},
			Endpoints: []discoveryv1.Endpoint{
				{Addresses: []string{"10.0.0.1"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(false)}},
				{Addresses: []string{"10.0.0.2"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(true)}},
			},
		},
	)
	ip, err := ResolveServiceIP(context.Background(), client, "ns-1", "svc-1", true)
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.2", ip)
}

func TestResolveServiceIPHeadlessNoReadyEndpoints(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "svc-1", Namespace: "ns-1"},
			Spec:       corev1.ServiceSpec{ClusterIP: "None"},
		},
		&discoveryv1.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name: "es-1", Namespace: "ns-1",
				Labels: map[string]string{"kubernetes.io/service-name": "svc-1"},
			},
			Endpoints: []discoveryv1.Endpoint{
				{Addresses: []string{"10.0.0.1"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(false)}},
			},
		},
	)
	_, err := ResolveServiceIP(context.Background(), client, "ns-1", "svc-1", true)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no ready endpoints")
}

func TestPickReadyEndpointNilReadyIncluded(t *testing.T) {
	slices := []discoveryv1.EndpointSlice{{
		Endpoints: []discoveryv1.Endpoint{
			{Addresses: []string{"10.0.0.1"}, Conditions: discoveryv1.EndpointConditions{Ready: nil}},
		},
	}}
	ip, err := PickReadyEndpoint(slices)
	require.NoError(t, err)
	assert.Equal(t, "10.0.0.1", ip)
}

func TestSplitServiceHost(t *testing.T) {
	cases := []struct {
		host      string
		name      string
		namespace string
		ok        bool
	}{
		{"credential-provider.ack-agent-identity.svc", "credential-provider", "ack-agent-identity", true},
		{"credential-provider.ack-agent-identity.svc.cluster.local", "credential-provider", "ack-agent-identity", true},
		{"10.0.0.1", "", "", false},
		{"example.com", "", "", false},
	}
	for _, tc := range cases {
		name, ns, ok := SplitServiceHost(tc.host)
		assert.Equal(t, tc.ok, ok, tc.host)
		if ok {
			assert.Equal(t, tc.name, name, tc.host)
			assert.Equal(t, tc.namespace, ns, tc.host)
		}
	}
}
