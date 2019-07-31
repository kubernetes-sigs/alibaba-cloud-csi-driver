#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/oss/csiplugin-connector build/oss/plugin.csi.alibabacloud.com

export GOARCH="amd64"
export GOOS="linux"

branch="v1.0.0"
version="v1.14.3"
commitId=$GIT_SHA
buildTime=`date "+%Y-%m-%d-%H:%M:%S"`

CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/build/oss/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA
fi

rm -rf csiplugin-connector plugin.csi.alibabacloud.com