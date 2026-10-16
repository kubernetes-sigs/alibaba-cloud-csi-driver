#!/bin/bash

# Copyright 2026 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

# These tests populate fixed credential paths; run them only in a disposable container.
runtime=${CONTAINER_RUNTIME:-docker}
image=${CSI_INTEGRATION_IMAGE:-docker.io/library/busybox:1.37.0}
build_dir=$(mktemp -d)
trap 'rm -rf "$build_dir"' EXIT

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -tags=integration -c \
  -o "$build_dir/jwtauth.test" ./pkg/mounter/jwtauth

"$runtime" run --rm --platform linux/amd64 --network none \
  --cap-drop=ALL --security-opt=no-new-privileges \
  -e CSI_JWTAUTH_CONTAINER_TEST=1 \
  -v "$build_dir/jwtauth.test:/qa.test:ro" \
  --entrypoint /qa.test "$image" \
  -test.run '^TestSubstrateExchangeAndRefreshIntegration$' -test.v -test.timeout=30s
