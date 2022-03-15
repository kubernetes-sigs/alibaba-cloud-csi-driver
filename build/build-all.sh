#!/usr/bin/env bash
set -ex

cd /Users/wangjiao/go/src/github.com/JiaoDean/alibaba-cloud-csi-driver
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/ack/csiplugin-connector.go build/ack/csiplugin-connector-svc build/ack/csiplugin-connector

cp build/oss/csiplugin-connector.go build/ack/csiplugin-connector.go
cp build/oss/csiplugin-connector.service build/ack/csiplugin-connector.service
cp build/oss/nsenter build/ack/nsenter

export GOARCH="amd64"
export GOOS="linux"

VERSION="v1.16.9"
GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_TIME=`date +%FT%T%z`

CGO_ENABLED=0 go build -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${GIT_BRANCH} -X main.REVISION=${GIT_HASH} -X main.BUILDTIME=${BUILD_TIME}" -o plugin.csi.alibabacloud.com

cd /Users/wangjiao/go/src/github.com/JiaoDean/alibaba-cloud-csi-driver/build/ack/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv /Users/wangjiao/go/src/github.com/JiaoDean/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/acs1/csi-plugin-nas-rich-client:$VERSION-$GIT_HASH ./
  rm -rf csiplugin-connector.go csiplugin-connector.service csiplugin-connector plugin.csi.alibabacloud.com ossfs_1.80.6_centos7.0_x86_64.rpm nsenter jindofs-fuse
  docker push registry.cn-hangzhou.aliyuncs.com/acs1/csi-plugin-nas-rich-client:$VERSION-$GIT_HASH
fi

