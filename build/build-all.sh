#!/usr/bin/env bash
set -e

cd ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/all/csiplugin-connector.go build/all/csiplugin-connector-svc build/all/csiplugin-connector

cp build/oss/csiplugin-connector.go build/all/csiplugin-connector.go
cp build/oss/csiplugin-connector.service build/all/csiplugin-connector.service
cp build/oss/ossfs_1.80.6_centos7.0_x86_64.rpm build/all/ossfs_1.80.6_centos7.0_x86_64.rpm
cp build/oss/nsenter build/all/nsenter

export GOARCH="amd64"
export GOOS="linux"

branch="v1.0.0"
version="v1.14.5"
commitId=$GIT_SHA
buildTime=`date "+%Y-%m-%d-%H:%M:%S"`

CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/build/all/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$version-$GIT_SHA
fi


rm -rf csiplugin-connector.go csiplugin-connector.service csiplugin-connector plugin.csi.alibabacloud.com ossfs_1.80.6_centos7.0_x86_64.rpm nsenter
