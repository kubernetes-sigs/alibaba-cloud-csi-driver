#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

branch="v1.0.0"
version="v1.14.5"
commitId=$GIT_SHA
buildTime=`date "+%Y-%m-%d-%H:%M:%S"`

CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o plugin.csi.alibabacloud.com
mv plugin.csi.alibabacloud.com ./build/nas/

if [ "$1" == "" ]; then
  cd ./build/nas/
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-nasplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-nasplugin:$version-$GIT_SHA
fi

rm -rf plugin.csi.alibabacloud.com