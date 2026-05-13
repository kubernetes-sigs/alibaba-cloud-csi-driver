#!/bin/bash
set -e

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && git rev-parse --show-toplevel)"
VALUES_FILE="$ROOT/.local/test-values.yaml"

REGISTRY=$(yq '.images.registry' "$VALUES_FILE")
SECRET_NAME=$(yq '.imagePullSecrets[0]' "$VALUES_FILE")

AUTH=$(jq -r --arg server "$REGISTRY" '.auths[$server].auth // empty' ~/.docker/config.json)
if [ -z "$AUTH" ]; then
    echo "No docker credentials found for $REGISTRY"
    echo "Please run: docker login $REGISTRY"
    exit 1
fi

jq -n --arg server "$REGISTRY" --arg auth "$AUTH" \
    '{"auths": {($server): {"auth": $auth}}}' | \
kubectl create secret generic "$SECRET_NAME" \
    --from-file=.dockerconfigjson=/dev/stdin \
    --type=kubernetes.io/dockerconfigjson \
    "$@"
