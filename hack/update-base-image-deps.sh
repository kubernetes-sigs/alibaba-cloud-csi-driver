#!/bin/bash

set -e

# Get visibility on what is included in our base image
buildctl build --frontend dockerfile.v0 \
    --local context=. \
    --local dockerfile=build/multi \
    --opt filename=Dockerfile.multi \
    --opt platform=linux/amd64 \
    --opt target=dep-list \
    --output type=oci | \
syft scan --select-catalogers "+sbom-cataloger" -o syft-table oci-archive:/dev/stdin > ./hack/base-image-deps.txt
