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

// TokenKindAapp is a token used to auth to CredStore.
const TokenKindAapp = "app"

// AppToken is JWT session token for CredStore.
type AppToken struct {
	Client string `json:"cli"`
	Kind   string `json:"kind"`
}

// Verify verifies token structure.
func (tok AppToken) Verify() error {
	if tok.Kind != TokenKindAapp {
		return fmt.Errorf("expected app token, got %s token", tok.Kind)
	}
	return nil
}

// BuildAppToken creates and serializes a session token.
func BuildAppToken(client string) ([]byte, error) {
	tok := AppToken{
		Client: client,
		Kind:   TokenKindAapp,
	}

	return json.Marshal(tok)
}
