# Copyright 2019 The Kubernetes Authors.
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

HELM_VERSION := v3.16.4
SYFT_VERSION := 1.18.1
OS := $(shell go env GOHOSTOS)
ARCH := $(shell go env GOHOSTARCH)

export PATH := $(CURDIR)/bin:$(PATH)

bin:
	mkdir -p bin

bin/helm: | bin
	curl -L https://get.helm.sh/helm-$(HELM_VERSION)-$(OS)-$(ARCH).tar.gz | tar -xz -C bin --strip-components=1 $(OS)-$(ARCH)/helm

bin/syft: | bin
	curl -L https://github.com/anchore/syft/releases/download/v$(SYFT_VERSION)/syft_$(SYFT_VERSION)_$(OS)_$(ARCH).tar.gz | tar -xz -C bin syft

.PHONY: fmt
fmt:
	./hack/check-format.sh

.PHONY: lint
lint:
	./hack/check-golint.sh

.PHONY: vet
vet:
	./hack/check-govet.sh

.PHONY: test
test:
	./hack/check-unittest.sh

.PHONY: update-deps check-deps
update-deps:
	go mod tidy
	go mod vendor
check-deps: update-deps
	git diff --exit-code go.mod go.sum vendor

.PHONY: update-dockerfile update-base-image-deps check-base-image-deps
update-dockerfile: bin/syft
	./hack/update-dockerfile.sh
update-base-image-deps: bin/syft
	./hack/update-base-image-deps.sh
check-base-image-deps: update-base-image-deps
	git diff --exit-code hack/base-image-deps.txt

.PHONY: check-helm-kind
check-helm-kind: bin/helm
check-helm-kind: export KUBECONFIG := /tmp/kubeconfig
check-helm-kind:
# Use the oldest k8s version we support
	kind create cluster --image kindest/node:v1.26.15
	./hack/check-helm.sh
	kind delete cluster

.PHONY: build
build:
	./build/build-all-multi.sh

pkg/cloud/ecsmock.go: pkg/cloud/ecsinterface.go
	mockgen -source pkg/cloud/ecsinterface.go -destination $@ -package cloud

PROTOC=protoc
pkg/disk/proto/disk.pb.go pkg/disk/proto/disk_ttrpc.pb.go: pkg/disk/disk.proto
	go install -mod=mod \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		github.com/containerd/ttrpc/cmd/protoc-gen-go-ttrpc
	$(PROTOC) -I pkg/disk disk.proto --go_out=pkg/disk --go-ttrpc_out=pkg/disk

CSI_VERSION := $(shell git describe --tags --always)
csi-agent-bin-linux-x86_64: output/csi-agent-bin-linux-amd64
	cp $< output/$@
csi-agent-bin-linux-aarch64: output/csi-agent-bin-linux-arm64
	cp $< output/$@
output/csi-agent-bin-%:
	CGO_ENABLED=0 \
	GOOS="$(firstword $(subst -, ,$*))" \
	GOARCH="$(lastword $(subst -, ,$*))" \
	go build -trimpath \
		-ldflags "-s -w -X github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version.VERSION=$(CSI_VERSION)" \
		-o output/csi-agent-bin-$* ./cmd/csi-agent

.PHONY: clean
clean:
	rm -rf bin output
