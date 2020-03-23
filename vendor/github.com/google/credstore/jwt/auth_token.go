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

// TokenKindAuth is a token used to query CredStore.
const TokenKindAuth = "auth"

// AuthToken is JWT session token for CredStore.
type AuthToken struct {
	Client string `json:"cli"`
	Kind   string `json:"kind"`
}

// Verify verifies token structure.
func (tok AuthToken) Verify() error {
	if tok.Kind != TokenKindAuth {
		return fmt.Errorf("expected auth token, got %s token", tok.Kind)
	}
	return nil
}

// BuildAuthToken creates and serializes a session token.
func BuildAuthToken(client string) ([]byte, error) {
	tok := AuthToken{
		Client: client,
		Kind:   TokenKindAuth,
	}

	return json.Marshal(tok)
}
