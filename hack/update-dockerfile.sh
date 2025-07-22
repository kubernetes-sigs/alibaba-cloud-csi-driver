#!/bin/bash

set -e

skopeo() {
    command skopeo --override-os linux --override-arch amd64 "$@"
}

DOCKERFILE=build/multi/Dockerfile.multi

DISTROLESS=registry-cn-hangzhou.ack.aliyuncs.com/dev/ack-base/distroless/base-debian12
DEBIAN=registry-cn-hangzhou.ack.aliyuncs.com/dev/debian

if [ "$UPSTREAM" ]; then
    DISTROLESS=gcr.io/distroless/base-debian12
    DEBIAN=docker.io/debian
fi

DISTROLESS_DIGEST=$(skopeo inspect docker://$DISTROLESS --format '{{.Digest}}')
echo "The latest distroless digest is $DISTROLESS_DIGEST"

DEBIAN_TAG=$(skopeo list-tags docker://$DEBIAN | jq -r '.Tags|map(select(test("^bookworm-.+-slim$"))) | sort | last')
echo "The latest debian tag is $DEBIAN_TAG"

sed -i "
    s|@sha[0-9a-f:]* as distroless-base|@$DISTROLESS_DIGEST as distroless-base|;
    s|debian:[0-9a-z-]* as |debian:$DEBIAN_TAG as |;
    " $DOCKERFILE

./hack/update-base-image-deps.sh
