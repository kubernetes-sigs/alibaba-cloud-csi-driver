#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

export GOARCH="amd64"
export GOOS="linux"

CGO_ENABLED=0 go build -o lvmplugin.csi.alibabacloud.com

if [ "$1" == "" ]; then
  version="v1.13"
  cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/build/lvm/
  mv ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/lvmplugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-lvmplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-lvmplugin:$version-$GIT_SHA
fi
