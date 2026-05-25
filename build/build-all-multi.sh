#!/usr/bin/env bash
set -ex

PLATFORM=${PLATFORM:-"linux/amd64,linux/arm64"}
IMAGE_TAG=${IMAGE_TAG:-"latest"}

if [ -n "$IMAGE_REPO" ]; then
    PUSH=true
else
    IMAGE_REPO="local"
    PUSH=false
fi

BUILD_ARGS=(
    --frontend dockerfile.v0
    --local context=.
    --opt "platform=${PLATFORM}"
    --opt "build-arg:CSI_VERSION=$(git describe --tags --match='v*' --always --dirty)"
    --opt "build-arg:SOURCE_DATE_EPOCH=$(git log -1 --pretty=%ct)"
)
BUILD_ARGS+=("$@")

should_build() {
    if [ -z "$IMAGES" ]; then
        return 0
    fi
    local name=${1%%:*}
    for img in $IMAGES; do
        if [ "$img" = "$name" ]; then
            return 0
        fi
    done
    return 1
}

build_image() {
    local image=$1
    shift
    if ! should_build "$image"; then
        return 0
    fi
    buildctl build "${BUILD_ARGS[@]}" "$@" \
        --output "type=image,push=$PUSH,name=${IMAGE_REPO}/${image}"
}

MAIN_ARGS=(--local dockerfile=build/multi --opt filename=Dockerfile.multi)
build_image "csi-plugin:${IMAGE_TAG}"             "${MAIN_ARGS[@]}"
build_image "csi-plugin:${IMAGE_TAG}-controller"  "${MAIN_ARGS[@]}" --opt target=csi-controller
build_image "csi-plugin:${IMAGE_TAG}-init"        "${MAIN_ARGS[@]}" --opt target=init

AGENT_ARGS=(--local dockerfile=build/csi-agent --opt filename=Dockerfile)
build_image "csi-agent:${IMAGE_TAG}" "${AGENT_ARGS[@]}"

PROXY_ARGS=(--local dockerfile=build/mount-proxy --opt filename=Dockerfile)
build_image "csi-ossfs:${IMAGE_TAG}"      "${PROXY_ARGS[@]}" --opt target=ossfs
build_image "csi-ossfs:${IMAGE_TAG}-1.88" "${PROXY_ARGS[@]}" --opt target=ossfs-1.88
build_image "csi-ossfs2:${IMAGE_TAG}"     "${PROXY_ARGS[@]}" --opt target=ossfs2
build_image "csi-alinas:${IMAGE_TAG}"     "${PROXY_ARGS[@]}" --opt target=alinas
build_image "mount-proxy:${IMAGE_TAG}"    "${PROXY_ARGS[@]}" --opt target=aio
