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
	"encoding/pem"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/substrate/internal/ateapipb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type actorControlFixture struct {
	ateapipb.UnimplementedControlServer
	actor    *ateapipb.Actor
	template *ateapipb.ActorTemplate
	tokens   chan string
}

func (s *actorControlFixture) GetActor(ctx context.Context, req *ateapipb.GetActorRequest) (*ateapipb.Actor, error) {
	if req.GetActor().GetAtespace() != s.actor.Metadata.Atespace || req.GetActor().GetName() != s.actor.Metadata.Name {
		return nil, status.Error(codes.NotFound, "actor reference not found")
	}
	if s.tokens != nil {
		md, _ := metadata.FromIncomingContext(ctx)
		values := md.Get("authorization")
		if len(values) != 1 {
			return nil, status.Error(codes.Unauthenticated, "missing bearer")
		}
		s.tokens <- values[0]
	}
	return s.actor, nil
}

func (s *actorControlFixture) GetActorTemplate(_ context.Context, req *ateapipb.GetActorTemplateRequest) (*ateapipb.ActorTemplate, error) {
	if s.template == nil || req.GetActorTemplate().GetAtespace() != s.template.Metadata.Atespace || req.GetActorTemplate().GetName() != s.template.Metadata.Name {
		return nil, status.Error(codes.NotFound, "template reference not found")
	}
	return s.template, nil
}

func TestActorLookupTLSAndTokenRotation(t *testing.T) {
	seed := httptest.NewTLSServer(http.NotFoundHandler())
	certificate := seed.TLS.Certificates[0]
	ca := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: seed.Certificate().Raw})
	seed.Close()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	tokens := make(chan string, 2)
	server := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{certificate}, MinVersion: tls.VersionTLS12})))
	ateapipb.RegisterControlServer(server, &actorControlFixture{
		actor:  &ateapipb.Actor{Metadata: &ateapipb.ResourceMetadata{Uid: testUID, Atespace: "storage-test", Name: "actor", Annotations: map[string]string{PublishRequestsAnnotation: "[]"}}},
		tokens: tokens,
	})
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	t.Cleanup(func() { server.Stop(); require.NoError(t, <-done) })
	dir := t.TempDir()
	caPath, tokenPath := filepath.Join(dir, "ca.pem"), filepath.Join(dir, "token")
	require.NoError(t, os.WriteFile(caPath, ca, 0600))
	require.NoError(t, os.WriteFile(tokenPath, []byte("first"), 0600))
	client, err := NewActorClient(ActorClientOptions{Endpoint: listener.Addr().String(), CAFile: caPath, TokenFile: tokenPath})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Close()) })
	for _, token := range []string{"first", "rotated"} {
		require.NoError(t, os.WriteFile(tokenPath, []byte(token+"\n"), 0600))
		got, err := client.Lookup(t.Context(), ActorReference{UID: testUID, Name: "actor", Atespace: "storage-test"})
		require.NoError(t, err)
		require.Equal(t, testUID, got.UID)
		require.Equal(t, "[]", got.Annotation)
		require.False(t, got.Golden)
		require.Equal(t, "Bearer "+token, <-tokens)
	}
}

func TestActorLookupVerifiesGoldenTemplate(t *testing.T) {
	for _, tc := range []struct {
		name        string
		actorName   string
		templateUID string
		wantError   bool
	}{
		{"verified", "template-uid", "template-uid", false},
		{"mismatched template", "another-template", "template-uid", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			listener, err := net.Listen("tcp", "127.0.0.1:0")
			require.NoError(t, err)
			server := grpc.NewServer()
			ateapipb.RegisterControlServer(server, &actorControlFixture{
				actor:    &ateapipb.Actor{Metadata: &ateapipb.ResourceMetadata{Uid: testUID, Atespace: "ate-golden", Name: tc.actorName}, ActorTemplate: &ateapipb.ObjectRef{Atespace: "storage-test", Name: "template"}},
				template: &ateapipb.ActorTemplate{Metadata: &ateapipb.ResourceMetadata{Uid: tc.templateUID, Atespace: "storage-test", Name: "template"}},
			})
			done := make(chan error, 1)
			go func() { done <- server.Serve(listener) }()
			t.Cleanup(func() { server.Stop(); require.NoError(t, <-done) })
			conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			t.Cleanup(func() { require.NoError(t, conn.Close()) })
			client := &ActorClient{conn: conn, control: ateapipb.NewControlClient(conn)}
			info, err := client.Lookup(t.Context(), ActorReference{UID: testUID, Name: tc.actorName, Atespace: "ate-golden"})
			if tc.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, info.Golden)
			require.Equal(t, tc.templateUID, info.TemplateUID)
		})
	}
}

func TestActorLookupRejectsRecreatedActorUID(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	server := grpc.NewServer()
	ateapipb.RegisterControlServer(server, &actorControlFixture{actor: &ateapipb.Actor{Metadata: &ateapipb.ResourceMetadata{Uid: "replacement-uid", Name: "actor", Atespace: "storage-test"}}})
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	t.Cleanup(func() { server.Stop(); require.NoError(t, <-done) })
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, conn.Close()) })
	client := &ActorClient{conn: conn, control: ateapipb.NewControlClient(conn)}
	_, err = client.Lookup(t.Context(), ActorReference{UID: testUID, Name: "actor", Atespace: "storage-test"})
	require.Equal(t, codes.PermissionDenied, status.Code(err))
}

func TestActorClientRejectsMissingTrust(t *testing.T) {
	_, err := NewActorClient(ActorClientOptions{Endpoint: "localhost:443"})
	require.Error(t, err)
}

func TestActorLookupLive(t *testing.T) {
	endpoint := os.Getenv("SUBSTRATE_LIVE_ENDPOINT")
	if endpoint == "" {
		t.Skip("set SUBSTRATE_LIVE_* to run the read-only cluster probe")
	}
	client, err := NewActorClient(ActorClientOptions{
		Endpoint: endpoint, CAFile: os.Getenv("SUBSTRATE_LIVE_CA_FILE"),
		TokenFile: os.Getenv("SUBSTRATE_LIVE_TOKEN_FILE"), ServerName: os.Getenv("SUBSTRATE_LIVE_SERVER_NAME"),
	})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Close()) })
	ref := ActorReference{UID: os.Getenv("SUBSTRATE_LIVE_ACTOR_UID"), Name: os.Getenv("SUBSTRATE_LIVE_ACTOR_NAME"), Atespace: os.Getenv("SUBSTRATE_LIVE_ATESPACE")}
	actor, err := client.Lookup(t.Context(), ref)
	require.NoError(t, err)
	require.Equal(t, ref.UID, actor.UID)
	require.Equal(t, ref.Name, actor.Name)
	require.Equal(t, ref.Atespace, actor.Atespace)
	if ref.Atespace == "ate-golden" {
		require.True(t, actor.Golden)
	} else {
		require.NotEmpty(t, actor.Annotation)
	}
}
