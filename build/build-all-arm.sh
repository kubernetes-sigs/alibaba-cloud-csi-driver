#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`

# rm -rf build/ack/csiplugin-connector.go build/ack/csiplugin-connector-svc build/ack/csiplugin-connector

# cp build/oss/csiplugin-connector.go build/ack/csiplugin-connector.go
# cp build/oss/csiplugin-connector.service build/ack/csiplugin-connector.service
# cp build/oss/ossfs_1.80.6_centos7.0_x86_64.rpm build/ack/ossfs_1.80.6_centos7.0_x86_64.rpm
# cp build/oss/nsenter build/ack/nsenter

# export GOARCH="amd64"
# export GOOS="linux"

VERSION="v1.16.9"
# GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
# GIT_BRANCH=`git symbolic-ref --short -q HEAD`
# BUILD_TIME=`date +%FT%T%z`


docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_SHA ./ -f ./build/arm/Dockerfile.arm
docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_SHA 

