#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


VERSION="v1.16.9"


docker build -t=registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_SHA-arm64 --build-arg ARCH=arm64/ . -f ./build/arm/Dockerfile.arm
docker push registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:$VERSION-$GIT_SHA 

