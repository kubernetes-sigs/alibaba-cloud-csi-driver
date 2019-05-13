#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

branch="v1.0.0"
version="v1.13.2"
commitId=$GIT_SHA
buildTime=`date "+%Y-%m-%d-%H:%M:%S"`

CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o diskplugin.csi.alibabacloud.com
mv diskplugin.csi.alibabacloud.com ./build/

if [ "$1" == "" ]; then
  cd ./build
  cp ./disk/Dockerfile ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA
fi
