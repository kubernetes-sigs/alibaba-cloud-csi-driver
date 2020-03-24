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

package jwt

import (
	"encoding/json"
	"fmt"
)

// TokenKindRPC is a token used to query CredStore.
const TokenKindRPC = "rpc"

// RPCToken is JWT token to access other services.
type RPCToken struct {
	Client  string `json:"cli"`
	Kind    string `json:"kind"`
	Service string `json:"svc"`
	Method  string `json:"meth"`
}

// Verify verifies token structure.
func (tok RPCToken) Verify() error {
	if tok.Kind != TokenKindRPC {
		return fmt.Errorf("expected rpc token, got %s token", tok.Kind)
	}
	return nil
}

// BuildRPCToken creates and serializes a session token.
func BuildRPCToken(client, service, method string) ([]byte, error) {
	tok := RPCToken{
		Client:  client,
		Kind:    TokenKindRPC,
		Service: service,
		Method:  method,
	}

	return json.Marshal(tok)
}
