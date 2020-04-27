#!/usr/bin/env bash
set -e

# PROJECT_DIR=${GOPATH}/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
PROJECT_DIR=/Users/etby/Dev/TMP/alibaba-cloud-csi-driver

cd $PROJECT_DIR
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

rm -rf build/oss/csiplugin-connector build/oss/plugin.csi.alibabacloud.com

export GOARCH="amd64"
export GOOS="linux"

branch="v1.0.0"
version="v1.14.5"
commitId=$GIT_SHA
buildTime=`date "+%Y-%m-%d-%H:%M:%S"`

CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o plugin.csi.alibabacloud.com

cd ${PROJECT_DIR}/build/oss/
CGO_ENABLED=0 go build csiplugin-connector.go

if [ "$1" == "" ]; then
  mv ${PROJECT_DIR}/plugin.csi.alibabacloud.com ./

  # centos
  docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-ossplugin:$version-$GIT_SHA ./
  docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-ossplugin:$version-$GIT_SHA

  # ubuntu
  docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-ossplugin:$version-$GIT_SHA-ubuntu-1804 ./ -f dockerfiles/ubuntu-1804
  docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-ossplugin:$version-$GIT_SHA-ubuntu-1804
fi

rm -rf csiplugin-connector plugin.csi.alibabacloud.com
