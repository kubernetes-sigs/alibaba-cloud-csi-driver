#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

CGO_ENABLED=0 go build -o csi-diskplugin
mv csi-diskplugin ./build/

if [ "$1" == "" ]; then
  cd ./build
  version="v1.10"
  cp ./disk/Dockerfile ./
  cp ./disk/nsenter ./
  #cp ./disk/diskplugin ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:$version-$GIT_SHA
fi
