/*
Copyright 2026 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package substrate

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/substrate/internal/ateapipb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const actorLookupTimeout = 10 * time.Second

type ActorClientOptions struct {
	Endpoint   string
	CAFile     string
	TokenFile  string
	ServerName string
}

type ActorClient struct {
	conn    *grpc.ClientConn
	control ateapipb.ControlClient
}

type ActorReference struct {
	UID      string
	Name     string
	Atespace string
}

type ActorInfo struct {
	UID         string
	Atespace    string
	Name        string
	Annotation  string
	Golden      bool
	TemplateUID string
}

func NewActorClient(opts ActorClientOptions) (*ActorClient, error) {
	if opts.Endpoint == "" || opts.CAFile == "" || opts.TokenFile == "" {
		return nil, fmt.Errorf("substrate endpoint, CA and projected token file are required")
	}
	// Trust is loaded once; projected CA bundle changes require a Node Pod rollout.
	// Unlike this pool, the projected token is read on every RPC below.
	ca, err := os.ReadFile(opts.CAFile)
	if err != nil {
		return nil, fmt.Errorf("read Substrate CA: %w", err)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("substrate CA contains no certificates")
	}
	conn, err := grpc.NewClient(opts.Endpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12, RootCAs: roots, ServerName: opts.ServerName})),
		grpc.WithPerRPCCredentials(fileToken(opts.TokenFile)),
	)
	if err != nil {
		return nil, fmt.Errorf("create Substrate connection: %w", err)
	}
	return &ActorClient{conn: conn, control: ateapipb.NewControlClient(conn)}, nil
}

func (c *ActorClient) Lookup(ctx context.Context, ref ActorReference) (ActorInfo, error) {
	if ref.UID == "" || ref.Name == "" || ref.Atespace == "" {
		return ActorInfo{}, status.Error(codes.InvalidArgument, "actor UID, name and namespace are required")
	}
	// Actor and optional Golden template queries share this budget; a shorter
	// caller deadline still takes precedence.
	ctx, cancel := context.WithTimeout(ctx, actorLookupTimeout)
	defer cancel()
	response, err := c.control.GetActor(ctx, &ateapipb.GetActorRequest{Actor: &ateapipb.ObjectRef{Atespace: ref.Atespace, Name: ref.Name}})
	if err != nil {
		return ActorInfo{}, err
	}
	if response.GetMetadata().GetUid() != ref.UID {
		return ActorInfo{}, status.Error(codes.PermissionDenied, "actor lookup returned a different UID")
	}
	meta := response.GetMetadata()
	if meta.GetAtespace() != ref.Atespace || meta.GetName() != ref.Name {
		return ActorInfo{}, status.Error(codes.PermissionDenied, "actor lookup returned a different reference")
	}
	info := ActorInfo{UID: meta.GetUid(), Atespace: meta.Atespace, Name: meta.Name, Annotation: meta.GetAnnotations()[PublishRequestsAnnotation]}
	if info.Atespace != "ate-golden" {
		return info, nil
	}
	templateRef := response.GetActorTemplate()
	if templateRef.GetAtespace() == "" || templateRef.GetName() == "" {
		return ActorInfo{}, status.Error(codes.FailedPrecondition, "golden actor lacks a template reference")
	}
	template, err := c.control.GetActorTemplate(ctx, &ateapipb.GetActorTemplateRequest{ActorTemplate: templateRef})
	if err != nil {
		return ActorInfo{}, err
	}
	templateMeta := template.GetMetadata()
	if templateMeta.GetAtespace() != templateRef.Atespace || templateMeta.GetName() != templateRef.Name || templateMeta.GetUid() == "" || templateMeta.GetUid() != info.Name {
		return ActorInfo{}, status.Error(codes.FailedPrecondition, "golden actor does not match its template UID")
	}
	info.Golden = true
	info.TemplateUID = templateMeta.Uid
	return info, nil
}

func (c *ActorClient) Close() error { return c.conn.Close() }

type fileToken string

func (f fileToken) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	raw, err := os.ReadFile(string(f))
	if err != nil {
		return nil, fmt.Errorf("read projected Substrate token: %w", err)
	}
	token := strings.TrimSpace(string(raw))
	if token == "" {
		return nil, fmt.Errorf("projected Substrate token is empty")
	}
	return map[string]string{"authorization": "Bearer " + token}, nil
}

func (fileToken) RequireTransportSecurity() bool { return true }
