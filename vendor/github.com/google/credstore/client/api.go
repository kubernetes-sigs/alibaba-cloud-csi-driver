/*

Copyright 2017 Google Inc.

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

package client

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-microservice-helpers/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/google/credstore/api"
)

// GetSigningKey requests the current signing key from CredStore.
func GetSigningKey(ctx context.Context, conn *grpc.ClientConn, tok string) (crypto.PublicKey, error) {
	ctx = WithBearerToken(ctx, tok)

	cli := pb.NewCredStoreAuthClient(conn)
	repl, err := cli.SigningKey(ctx, &pb.SigningKeyRequest{})
	if err != nil {
		return nil, err
	}

	pubkey, err := x509.ParsePKIXPublicKey(repl.GetSigningKey())
	if k, ok := pubkey.(crypto.PublicKey); ok {
		return k, nil
	}
	return nil, fmt.Errorf("cannot parse the public key")
}

// GetAuthToken returns a session JWT token.
func GetAuthToken(ctx context.Context, conn *grpc.ClientConn, tok string) (string, error) {
	ctx = WithBearerToken(ctx, tok)

	cli := pb.NewCredStoreAuthClient(conn)
	repl, err := cli.Auth(ctx, &pb.AuthRequest{})
	if err != nil {
		return "", err
	}

	return repl.GetAuthJwt(), nil
}

const appTokenEnv = "CREDSTORE_APP_TOKEN"

// GetAppToken returns the app token for currently running app.
func GetAppToken() (string, error) {
	tok := os.Getenv(appTokenEnv)
	if tok == "" {
		return "", fmt.Errorf("app token not present or malformed")
	}
	return tok, nil
}

// GetTokenForRemote returns a JWT token for given remote.
func GetTokenForRemote(ctx context.Context, conn *grpc.ClientConn, sessTok string, target string) (string, error) {
	ctx = WithBearerToken(ctx, sessTok)

	cli := pb.NewCredStoreClient(conn)
	repl, err := cli.GetToken(ctx, &pb.GetTokenRequest{Target: target})
	if err != nil {
		return "", err
	}

	return repl.GetSessionJwt(), nil
}

type CredstoreClient struct {
	conn       *grpc.ClientConn
	appTok     string
	sessTok    string
	signingKey crypto.PublicKey
}

// NewCredstoreClient creates a new CredstoreClient.
func NewCredstoreClient(ctx context.Context, credStoreAddress string, credStoreCA string) (*CredstoreClient, error) {
	appTok, err := GetAppToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get app token: %v", err)
	}

	conn, err := clienthelpers.NewGRPCConn(credStoreAddress, credStoreCA, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to create connection to credstore: %v", err)
	}

	credStoreSessionTok, err := GetAuthToken(ctx, conn, appTok)
	if err != nil {
		return nil, fmt.Errorf("failed to get session token: %v", err)
	}

	credStoreKey, err := GetSigningKey(ctx, conn, appTok)
	if err != nil {
		return nil, fmt.Errorf("failed to get signing key: %v", err)
	}

	return &CredstoreClient{
		conn:       conn,
		appTok:     appTok,
		sessTok:    credStoreSessionTok,
		signingKey: credStoreKey,
	}, nil
}

// SigningKey returns credstore public key for JWT verification.
func (c CredstoreClient) SigningKey() crypto.PublicKey {
	return c.signingKey
}

// GetTokenForRemote returns a token for given remote host:port.
func (c CredstoreClient) GetTokenForRemote(ctx context.Context, remoteHostPort string) (string, error) {
	dnsname := strings.Split(remoteHostPort, ":")[0]
	return GetTokenForRemote(ctx, c.conn, c.sessTok, dnsname)
}
