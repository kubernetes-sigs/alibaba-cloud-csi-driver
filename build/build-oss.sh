#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

CGO_ENABLED=0 go build -o ossplugin.csi.alibabacloud.com

if [ "$1" == "" ]; then
  mv ossplugin.csi.alibabacloud.com ./build/oss/
  cd ./build/oss
  version="v1.13"
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA
fi
