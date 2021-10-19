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
GOPROXY=direct
REPONAME="$(shell pwd | rev | awk -F \/ '{ print $$2 }' | rev)"
VERSION="v1.20.4"
GIT_SHA=$(shell git rev-parse --short HEAD || echo "HEAD")

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

.PHONY: build-disk-image
build-disk-image:
	docker build -t registry.cn-hangzhou.aliyuncs.com/acs/csi-diskplugin:$(VERSION)-$(GIT_SHA) .
