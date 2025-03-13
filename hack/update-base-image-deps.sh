#!/bin/bash

set -e

SYFT=(syft scan -o template -t ./hack/base-image-deps.tpl)

# Get visibility on what is included in our base image
scan() {
    if command -v buildctl > /dev/null; then
        buildctl build --frontend dockerfile.v0 \
            --local context=. \
            --local dockerfile=build/multi \
            --opt filename=Dockerfile.multi \
            --opt platform=linux/amd64 \
            --opt target=dep-list \
            --output type=oci | "${SYFT[@]}" oci-archive:/dev/stdin
    else
        docker buildx build . \
            --platform linux/amd64 \
            -f ./build/multi/Dockerfile.multi \
            --target dep-list \
            --output type=image,name=csi-dep-list
        "${SYFT[@]}" docker:csi-dep-list
    fi
}

scan > ./hack/base-image-deps.txt
