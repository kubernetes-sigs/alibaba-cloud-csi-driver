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

// Package client provides helper APIs for clients.
//
// Based on https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/auth, licensed under Apache-2.0
package client

import (
	"crypto"
	"encoding/json"
	"strings"

	"github.com/google/credstore/jwt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	jose "gopkg.in/square/go-jose.v2"
)

type tokenKeyType struct{}

// TokenKey is the context key for token interceptor payload
var TokenKey = tokenKeyType{}

// CredStoreTokenInterceptor returns a new unary server interceptor that performs per-request auth.
func CredStoreTokenInterceptor(publicKey crypto.PublicKey) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var newCtx context.Context
		var err error

		md, ok := metadata.FromIncomingContext(ctx)
		if ok == false {
			return nil, grpc.Errorf(codes.Unauthenticated, "cannot read metadata for request")
		}

		tok := md["authorization"]
		if len(tok) != 1 {
			return nil, grpc.Errorf(codes.Unauthenticated, "bad authorization string")
		}
		if tok[0] == "" {
			return nil, grpc.Errorf(codes.Unauthenticated, "authorization header is missing")
		}

		splits := strings.SplitN(tok[0], " ", 2)
		if len(splits) < 2 {
			return nil, grpc.Errorf(codes.Unauthenticated, "bad authorization string")
		}

		if strings.ToLower(splits[0]) != strings.ToLower("bearer") {
			return nil, grpc.Errorf(codes.Unauthenticated, "request unauthenticated with 'bearer'")
		}

		jwtTokString := splits[1]
		jwtTok, err := jose.ParseSigned(jwtTokString)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "failed to parse token: %v", err)
		}

		jwtPayload, err := jwtTok.Verify(publicKey)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "failed to verify token: %v", err)
		}

		newCtx = context.WithValue(ctx, TokenKey, jwtPayload)

		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

// CredStoreMethodAuthInterceptor returns a new unary server interceptor that verifies rpc token.
func CredStoreMethodAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		jwtTokString := ctx.Value(TokenKey).([]byte)
		var jwtTok jwt.RPCToken
		err := json.Unmarshal(jwtTokString, &jwtTok)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "cannot deserialize JWT token: %v", err)
		}

		if err := jwtTok.Verify(); err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "cannot verify JWT token: %v", err)
		}

		splits := strings.SplitN(info.FullMethod, "/", 3)
		if len(splits) != 3 {
			return nil, grpc.Errorf(codes.Internal, "failed to split method %s", info.FullMethod)
		}

		if jwtTok.Service != splits[1] {
			return nil, grpc.Errorf(codes.PermissionDenied, "not authrized to access service %s", jwtTok.Service)
		}

		if jwtTok.Method != "*" {
			if jwtTok.Method != splits[2] {
				return nil, grpc.Errorf(codes.PermissionDenied, "not authrized to access method %s of service %s", jwtTok.Method, jwtTok.Service)
			}
		}

		return handler(ctx, req)
	}
}
