#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/AliyunContainerService/csi-plugin/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

CGO_ENABLED=0 go build -o csi-ossplugin
mv csi-ossplugin ./build/

if [ "$1" == "" ]; then
  cd ./build
  version="v1.13"
  cp ./oss/Dockerfile ./
  cp ./oss/ossfs_1.80.3_centos7.0_x86_64.rpm ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-ossplugin:$version-$GIT_SHA
fi
