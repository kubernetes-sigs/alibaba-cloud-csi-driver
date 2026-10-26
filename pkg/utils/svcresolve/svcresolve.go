// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Package svcresolve resolves Kubernetes Service addresses to dialable IPs
// for pods that run with hostNetwork and cannot use cluster DNS.
package svcresolve

import (
	"context"
	"fmt"
	"strings"

	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ResolveServiceIP resolves a Service to a single IP address. It always calls
// Get Service first; if the Service has a ClusterIP it is returned directly.
// For headless Services (ClusterIP: None) the lookupEndpoints parameter
// controls behavior: when false an error is returned, when true
// EndpointSlices are listed and the first ready endpoint address is returned.
func ResolveServiceIP(ctx context.Context, client kubernetes.Interface, namespace, name string, lookupEndpoints bool) (string, error) {
	svc, err := client.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("get Service %s/%s: %w", namespace, name, err)
	}
	if svc.Spec.ClusterIP != "" && svc.Spec.ClusterIP != "None" {
		return svc.Spec.ClusterIP, nil
	}
	if !lookupEndpoints {
		return "", fmt.Errorf("service %s/%s has no ClusterIP", namespace, name)
	}
	slices, err := client.DiscoveryV1().EndpointSlices(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "kubernetes.io/service-name=" + name,
	})
	if err != nil {
		return "", fmt.Errorf("list EndpointSlices for headless Service %s/%s: %w", namespace, name, err)
	}
	ip, err := PickReadyEndpoint(slices.Items)
	if err != nil {
		return "", fmt.Errorf("headless Service %s/%s: %w", namespace, name, err)
	}
	return ip, nil
}

// ResolveServiceAddresses resolves a Service to a list of IP addresses. For
// ClusterIP Services the ClusterIP is returned as a single-element slice.
// For headless Services with lookupEndpoints=true all ready endpoint
// addresses are returned, for client-side load balancing.
func ResolveServiceAddresses(ctx context.Context, client kubernetes.Interface, namespace, name string, lookupEndpoints bool) ([]string, error) {
	svc, err := client.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get Service %s/%s: %w", namespace, name, err)
	}
	if svc.Spec.ClusterIP != "" && svc.Spec.ClusterIP != "None" {
		return []string{svc.Spec.ClusterIP}, nil
	}
	if !lookupEndpoints {
		return nil, fmt.Errorf("service %s/%s has no ClusterIP", namespace, name)
	}
	slices, err := client.DiscoveryV1().EndpointSlices(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "kubernetes.io/service-name=" + name,
	})
	if err != nil {
		return nil, fmt.Errorf("list EndpointSlices for headless Service %s/%s: %w", namespace, name, err)
	}
	addrs := ListReadyEndpoints(slices.Items)
	if len(addrs) == 0 {
		return nil, fmt.Errorf("headless Service %s/%s: no ready endpoints", namespace, name)
	}
	return addrs, nil
}

// PickReadyEndpoint returns the first ready endpoint address from a list of
// EndpointSlices. Endpoints with Ready == false are skipped; endpoints with
// Ready == nil (unknown) are included, matching the Kubernetes convention
// where unknown readiness is treated as potentially ready.
func PickReadyEndpoint(slices []discoveryv1.EndpointSlice) (string, error) {
	for _, slice := range slices {
		for _, ep := range slice.Endpoints {
			if ep.Conditions.Ready != nil && !*ep.Conditions.Ready {
				continue
			}
			if len(ep.Addresses) > 0 {
				return ep.Addresses[0], nil
			}
		}
	}
	return "", fmt.Errorf("no ready endpoints")
}

// ListReadyEndpoints returns all ready endpoint addresses from a list of
// EndpointSlices. Used by gRPC resolvers that need multiple addresses for
// client-side load balancing.
func ListReadyEndpoints(slices []discoveryv1.EndpointSlice) []string {
	var addrs []string
	for _, slice := range slices {
		for _, ep := range slice.Endpoints {
			if ep.Conditions.Ready != nil && !*ep.Conditions.Ready {
				continue
			}
			addrs = append(addrs, ep.Addresses...)
		}
	}
	return addrs
}

// SplitServiceHost parses <name>.<namespace>.svc[.cluster.local] and returns
// the name and namespace. Returns ok=false if the host is not a .svc address.
func SplitServiceHost(host string) (name, namespace string, ok bool) {
	parts := strings.Split(host, ".")
	if len(parts) < 3 || parts[2] != "svc" {
		return "", "", false
	}
	return parts[0], parts[1], true
}
