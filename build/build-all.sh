#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/all/csiplugin-connector.go build/all/csiplugin-connector-svc build/all/csiplugin-connector

cp build/oss/csiplugin-connector.go build/all/csiplugin-connector.go
cp build/oss/csiplugin-connector.service build/all/csiplugin-connector.service
cp build/oss/ossfs_1.80.6_centos7.0_x86_64.rpm build/all/ossfs_1.86.1_centos7.0_x86_64.rpm
cp build/oss/nsenter build/all/nsenter
cp build/ack/Dockerfile build/all/Dockerfile
cp build/ack/entrypoint.sh build/all/entrypoint.sh

export GOARCH="amd64"
export GOOS="linux"

VERSION="v1.14.5"
GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_USER=`git log -1 --pretty=format:'%an'`
BUILD_TIME=`date +%FT%T%z`

CGO_ENABLED=0 go build -ldflags "-X main.Version=${VERSION} -X main.Branch=${GIT_BRANCH} -X main.Revision=${GIT_HASH} -X main.BuildTime=${BUILD_TIME} -X main.BuildUser=${BUILD_USER}" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/build/all/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-shenzhen.aliyuncs.com/acs1/csi-metric-demo:$VERSION-$GIT_HASH ./
  docker push registry.cn-shenzhen.aliyuncs.com/acs1/csi-metric-demo:$VERSION-$GIT_HASH
fi


rm -rf csiplugin-connector.go csiplugin-connector.service csiplugin-connector plugin.csi.alibabacloud.com ossfs_1.86.1_centos7.0_x86_64.rpm nsenter
