// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package k8sresolver

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/svcresolve"
	"google.golang.org/grpc/resolver"
	"k8s.io/client-go/kubernetes"
)

const Scheme = "k8s"

type Builder struct{ client kubernetes.Interface }

func NewBuilder(client kubernetes.Interface) *Builder { return &Builder{client: client} }
func (b *Builder) Scheme() string                     { return Scheme }

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	if b == nil || b.client == nil {
		return nil, fmt.Errorf("k8sresolver: kubernetes client is required")
	}
	ns, svc, port, err := ParseTarget(target)
	if err != nil {
		return nil, fmt.Errorf("k8sresolver: %w", err)
	}
	r := &k8sResolver{
		client:  b.client,
		ns:      ns,
		svc:     svc,
		port:    port,
		cc:      cc,
		trigger: make(chan struct{}, 1),
	}
	go r.run()
	r.trigger <- struct{}{}
	return r, nil
}

func ParseTarget(target resolver.Target) (ns, svc, port string, err error) {
	ep := target.Endpoint()
	if target.URL.Host != "" && target.URL.Host != "localhost" {
		ep = target.URL.Host + "/" + strings.TrimPrefix(target.URL.Path, "/")
	}
	if parts := strings.SplitN(ep, "/", 2); len(parts) == 2 {
		ns, ep = parts[0], parts[1]
	} else {
		ns = "default"
	}
	svc, port, err = net.SplitHostPort(ep)
	if err != nil {
		svc, port = ep, "443"
	}
	if parts := strings.Split(svc, "."); len(parts) > 1 {
		svc, ns = parts[0], parts[1]
	}
	if ns == "" || svc == "" {
		return "", "", "", fmt.Errorf("invalid target %q, expected k8s:///<namespace>/<service>[:port]", target.String())
	}
	return ns, svc, port, nil
}

type k8sResolver struct {
	client  kubernetes.Interface
	ns      string
	svc     string
	port    string
	cc      resolver.ClientConn
	ctx     context.Context
	cancel  context.CancelFunc
	trigger chan struct{}
}

func (r *k8sResolver) ResolveNow(resolver.ResolveNowOptions) {
	select {
	case r.trigger <- struct{}{}:
	default:
	}
}

func (r *k8sResolver) Close() {
	if r.cancel != nil {
		r.cancel()
	}
}

func (r *k8sResolver) run() {
	r.ctx, r.cancel = context.WithCancel(context.Background())
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-r.trigger:
			r.updateState()
		}
	}
}

func (r *k8sResolver) updateState() {
	if r.ctx.Err() != nil {
		return
	}
	ips, err := svcresolve.ResolveServiceAddresses(r.ctx, r.client, r.ns, r.svc, true)
	if err != nil {
		r.cc.ReportError(err)
		return
	}
	targetPort, _ := strconv.Atoi(r.port)
	if targetPort == 0 {
		targetPort = 443
	}
	var addrs []resolver.Address
	for _, ip := range ips {
		addrs = append(addrs, resolver.Address{Addr: net.JoinHostPort(ip, strconv.Itoa(targetPort))})
	}
	if len(addrs) == 0 {
		r.cc.ReportError(fmt.Errorf("k8sresolver: no addresses for %s/%s", r.ns, r.svc))
		return
	}
	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil && r.ctx.Err() == nil {
		r.cc.ReportError(err)
	}
}
