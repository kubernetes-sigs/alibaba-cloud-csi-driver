#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


export GOARCH="amd64"
export GOOS="linux"

VERSION="v1.18.8"
GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
GIT_BRANCH=`git symbolic-ref --short -q HEAD`
BUILD_TIME=`date +%FT%T%z`

CGO_ENABLED=0 go build -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${GIT_BRANCH} -X main.REVISION=${GIT_HASH} -X main.BUILDTIME=${BUILD_TIME}" -o plugin.csi.alibabacloud.com

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/build/dbfs/

if [ "$1" == "" ]; then
  mv ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com ./
  docker build -t=registry.cn-hangzhou.aliyuncs.com/plugins/csi-plugin:$VERSION-$GIT_HASH-dbfs ./
  rm -rf plugin.csi.alibabacloud.com
  docker push registry.cn-hangzhou.aliyuncs.com/plugins/csi-plugin:$VERSION-$GIT_HASH-dbfs
fi

