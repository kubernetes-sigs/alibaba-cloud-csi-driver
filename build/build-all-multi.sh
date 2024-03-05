#!/usr/bin/env bash
set -ex

BUILD_ARGS=( \
    --frontend dockerfile.v0 \
    --local context=. \
    --local dockerfile=build/multi \
    --opt filename=Dockerfile.multi \
    --opt platform=linux/amd64,linux/arm64 \
    --opt build-arg:CSI_VERSION=$(git describe --tags --always --dirty)
)
BUILD_ARGS+=("$@")

buildctl build "${BUILD_ARGS[@]}" \
    --output type=image,name=alibaba-cloud-csi-driver:latest
buildctl build "${BUILD_ARGS[@]}" \
    --opt target=init \
    --output type=image,name=alibaba-cloud-csi-driver:latest-init
