#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi
GOPATH=/go

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/amd/csiplugin-connector.go build/amd/csiplugin-connector-svc build/amd/csiplugin-connector

cp build/lib/csiplugin-connector.go build/amd/csiplugin-connector.go
cp build/lib/csiplugin-connector.service build/amd/csiplugin-connector.service
cp build/lib/amd64-nsenter build/amd/nsenter
cp build/lib/freezefs.sh build/amd/freezefs.sh
cp build/lib/jindofs-fuse-3.7.3-20211207.tar.gz build/amd/jindofs-fuse-3.7.3-20211207.tar.gz
cp build/lib/aliyun-alinas-utils-1.1-1.al.noarch.rpm build/amd/aliyun-alinas-utils-1.1-1.al.noarch.rpm
cp build/lib/amd64-entrypoint.sh build/amd/amd64-entrypoint.sh

export GOARCH="amd64"
export GOOS="linux"

VERSION="v1.16.9"
GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_TIME=`date +%FT%T%z`

CGO_ENABLED=0 go build -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${GIT_BRANCH} -X main.REVISION=${GIT_HASH} -X main.BUILDTIME=${BUILD_TIME}" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/build/amd/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-plugin:$VERSION-$GIT_HASH ./
  rm -rf csiplugin-connector.go csiplugin-connector.service csiplugin-connector plugin.csi.alibabacloud.com ossfs_1.80.6_centos7.0_x86_64.rpm nsenter jindofs-fuse jindofs-fuse-3.7.3-20211207.tar.gz aliyun-alinas-utils-1.1-1.al.noarch.rpm amd64-entrypoint.sh freezefs.sh
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-plugin:$VERSION-$GIT_HASH
fi

