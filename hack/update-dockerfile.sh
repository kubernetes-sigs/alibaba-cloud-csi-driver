#!/bin/bash

set -e

DOCKERFILE=build/multi/Dockerfile.multi

DISTROLESS=gcr.io/distroless/base-debian12
DEBIAN=debian

if [ "$ACK" ]; then
    DISTROLESS=registry-cn-hangzhou.ack.aliyuncs.com/dev/ack-base/distroless/base-debian12
    DEBIAN=registry-cn-hangzhou.ack.aliyuncs.com/dev/debian
fi

DISTROLESS_DIGEST=$(crane digest $DISTROLESS)
echo "The latest distroless digest is $DISTROLESS_DIGEST"

DEBIAN_TAG=$(crane ls $DEBIAN | grep -E 'bookworm-.+-slim' | sort | tail -n1)
echo "The latest debian tag is $DEBIAN_TAG"

sed -i '' "
    s|@sha[0-9a-f:]* as distroless-base|@$DISTROLESS_DIGEST as distroless-base|;
    s|debian:[0-9a-z-]* as |debian:$DEBIAN_TAG as |;
    " $DOCKERFILE

# Get visibility on what is included in our base image
buildctl build --frontend dockerfile.v0 \
    --local context=. \
    --local dockerfile=build/multi \
    --opt filename=Dockerfile.multi \
    --opt platform=linux/amd64 \
    --opt target=dep-list \
    --output type=oci | \
syft scan --select-catalogers "+sbom-cataloger" -o syft-table oci-archive:/dev/stdin > ./hack/base-image-deps.txt
