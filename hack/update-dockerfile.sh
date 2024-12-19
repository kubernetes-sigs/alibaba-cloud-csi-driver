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

./hack/update-base-image-deps.sh
