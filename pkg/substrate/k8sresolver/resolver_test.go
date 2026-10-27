// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package k8sresolver

import (
	"context"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func boolPtr(b bool) *bool { return &b }

type fakeConn struct {
	mu      sync.Mutex
	addrs   []resolver.Address
	lastErr error
	updated chan struct{}
}

func newFakeConn() *fakeConn { return &fakeConn{updated: make(chan struct{}, 1)} }

func (c *fakeConn) UpdateState(state resolver.State) error {
	c.mu.Lock()
	c.addrs = state.Addresses
	c.mu.Unlock()
	select {
	case c.updated <- struct{}{}:
	default:
	}
	return nil
}
func (c *fakeConn) ReportError(err error)                                { c.mu.Lock(); c.lastErr = err; c.mu.Unlock() }
func (c *fakeConn) NewAddress(addrs []resolver.Address)                  {}
func (c *fakeConn) ParseServiceConfig(string) *serviceconfig.ParseResult { return nil }
func (c *fakeConn) lastError() error                                     { c.mu.Lock(); defer c.mu.Unlock(); return c.lastErr }

func (c *fakeConn) waitAddresses(t *testing.T) []resolver.Address {
	t.Helper()
	select {
	case <-c.updated:
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for resolver update")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.addrs
}

func TestParseTarget(t *testing.T) {
	cases := []struct {
		target string
		ns     string
		svc    string
		port   string
	}{
		{"k8s:///ate-system/api:443", "ate-system", "api", "443"},
		{"k8s:///api.ate-system:443", "ate-system", "api", "443"},
		{"k8s:///ate-system/api", "ate-system", "api", "443"},
	}
	for _, tc := range cases {
		t.Run(tc.target, func(t *testing.T) {
			ns, svc, port, err := ParseTarget(resolver.Target{URL: mustParseURL(tc.target)})
			require.NoError(t, err)
			assert.Equal(t, tc.ns, ns)
			assert.Equal(t, tc.svc, svc)
			assert.Equal(t, tc.port, port)
		})
	}
}

func TestResolveClusterIPService(t *testing.T) {
	client := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "api", Namespace: "ate-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.1"},
	})
	b := NewBuilder(client)
	conn := newFakeConn()
	_, err := b.Build(resolver.Target{URL: mustParseURL("k8s:///ate-system/api:443")}, conn, resolver.BuildOptions{})
	require.NoError(t, err)
	addrs := conn.waitAddresses(t)
	require.Len(t, addrs, 1)
	assert.Contains(t, addrs[0].Addr, "10.0.0.1")
}

func TestResolveHeadlessService(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "api", Namespace: "ate-system"},
			Spec:       corev1.ServiceSpec{ClusterIP: "None"},
		},
		&discoveryv1.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name: "api-es", Namespace: "ate-system",
				Labels: map[string]string{"kubernetes.io/service-name": "api"},
			},
			Endpoints: []discoveryv1.Endpoint{
				{Addresses: []string{"10.0.0.2", "10.0.0.3"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(true)}},
				{Addresses: []string{"10.0.0.4"}, Conditions: discoveryv1.EndpointConditions{Ready: boolPtr(false)}},
			},
		},
	)
	b := NewBuilder(client)
	conn := newFakeConn()
	_, err := b.Build(resolver.Target{URL: mustParseURL("k8s:///ate-system/api:443")}, conn, resolver.BuildOptions{})
	require.NoError(t, err)
	addrs := conn.waitAddresses(t)
	require.Len(t, addrs, 1, "ResolveServiceIP returns a single ready endpoint")
	assert.Contains(t, addrs[0].Addr, "10.0.0.2")
}

func TestResolveMissingService(t *testing.T) {
	client := fake.NewSimpleClientset()
	b := NewBuilder(client)
	conn := newFakeConn()
	_, err := b.Build(resolver.Target{URL: mustParseURL("k8s:///ate-system/absent:443")}, conn, resolver.BuildOptions{})
	require.NoError(t, err)
	select {
	case <-conn.updated:
		t.Fatal("should not update state for missing service")
	case <-time.After(500 * time.Millisecond):
	}
	require.Error(t, conn.lastError())
}

func TestResolveNowReResolvesAfterEndpointChange(t *testing.T) {
	client := fake.NewSimpleClientset(
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "api", Namespace: "ate-system"},
			Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.1"},
		},
	)
	b := NewBuilder(client)
	conn := newFakeConn()
	r, err := b.Build(resolver.Target{URL: mustParseURL("k8s:///ate-system/api:443")}, conn, resolver.BuildOptions{})
	require.NoError(t, err)
	first := conn.waitAddresses(t)
	require.Len(t, first, 1)
	assert.Contains(t, first[0].Addr, "10.0.0.1")

	_, err = client.CoreV1().Services("ate-system").Update(context.Background(),
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "api", Namespace: "ate-system"},
			Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.99"},
		}, metav1.UpdateOptions{})
	require.NoError(t, err)

	conn.updated = make(chan struct{}, 1)
	r.ResolveNow(resolver.ResolveNowOptions{})
	second := conn.waitAddresses(t)
	require.Len(t, second, 1)
	assert.Contains(t, second[0].Addr, "10.0.0.99")
}

func TestCloseStopsResolution(t *testing.T) {
	client := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "api", Namespace: "ate-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.1"},
	})
	b := NewBuilder(client)
	conn := newFakeConn()
	r, err := b.Build(resolver.Target{URL: mustParseURL("k8s:///ate-system/api:443")}, conn, resolver.BuildOptions{})
	require.NoError(t, err)
	_ = conn.waitAddresses(t)

	r.Close()

	conn.updated = make(chan struct{}, 1)
	r.ResolveNow(resolver.ResolveNowOptions{})
	select {
	case <-conn.updated:
		t.Fatal("should not update after Close")
	case <-time.After(300 * time.Millisecond):
	}
}

func mustParseURL(raw string) url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return *u
}
