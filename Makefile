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

PKG=github.com/kubernetes-sigs/alibaba-cloud-csi-driver

GO111MODULE=on
REPONAME="$(shell pwd | rev | awk -F \/ '{ print $$2 }' | rev)"

.EXPORT_ALL_VARIABLES:

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

.PHONY: build
build:
	./build/build-all-multi.sh "" $(REPONAME)

./PHONY: build-amd
build-amd:
	./build/build-all.sh "" $(REPONAME)

.PHONY: build-arm
build-arm:
	./build/build-all-arm.sh "" $(REPONAME)

.PHONY: build-nas
build-nas:
	./build/build-nas.sh "" $(REPONAME)

.PHONY: build-disk
build-disk:
	./build/build-disk.sh "" $(REPONAME)

pkg/disk/ecsmock.go: pkg/disk/ecsinterface.go
	mockgen -source pkg/disk/ecsinterface.go -destination $@ -package disk
