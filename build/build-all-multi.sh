#!/usr/bin/env bash
set -ex
REPO_NAME=$2

if [ "$REPO_NAME" == "" ]; then
    REPO_NAME="kubernetes-sigs"
fi

cd ${GOPATH}/src/github.com/$REPO_NAME/alibaba-cloud-csi-driver/
GIT_SHA=`git rev-parse --short HEAD || echo "HEAD"`


VERSION="v1.16.9"
# GIT_HASH=`git rev-parse --short HEAD || echo "HEAD"`
# GIT_BRANCH=`git symbolic-ref --short -q HEAD`
# BUILD_TIME=`date +%FT%T%z`

docker buildx build --platform linux/amd64,linux/arm64 . -f ./build/multi/Dockerfile.multi
