#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/ack/csiplugin-connector.go build/ack/csiplugin-connector-svc build/ack/csiplugin-connector

cp build/oss/csiplugin-connector.go build/ack/csiplugin-connector.go
cp build/oss/csiplugin-connector.service build/ack/csiplugin-connector.service
cp build/oss/ossfs_1.80.6_centos7.0_x86_64.rpm build/ack/ossfs_1.80.6_centos7.0_x86_64.rpm
cp build/oss/nsenter build/ack/nsenter

export GOARCH="amd64"
export GOOS="linux"

VERSION="v1.16.9"
GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_TIME=`date +%FT%T%z`

CGO_ENABLED=0 go build -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${GIT_BRANCH} -X main.REVISION=${GIT_HASH} -X main.BUILDTIME=${BUILD_TIME}" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/build/ack/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_HASH ./
  rm -rf csiplugin-connector.go csiplugin-connector.service csiplugin-connector plugin.csi.alibabacloud.com ossfs_1.80.6_centos7.0_x86_64.rpm nsenter
  docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_HASH
fi

