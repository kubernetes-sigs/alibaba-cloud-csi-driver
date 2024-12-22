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
OS := $(shell go env GOHOSTOS)
ARCH := $(shell go env GOHOSTARCH)

bin: export PATH := $(CURDIR)/bin:$(PATH)
bin:
	mkdir -p bin

bin/helm: | bin
	curl -L https://get.helm.sh/helm-$(HELM_VERSION)-$(OS)-$(ARCH).tar.gz | tar -xz -C bin --strip-components=1 '*/helm'

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

.PHONY: clean
clean:
	rm -rf bin
