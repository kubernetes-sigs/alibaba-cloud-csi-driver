#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

CGO_ENABLED=0 go build -o diskplugin.csi.aliyun.com
mv diskplugin.csi.aliyun.com ./build/

if [ "$1" == "" ]; then
  cd ./build
  version="v1.13"
  cp ./disk/Dockerfile ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA
fi
