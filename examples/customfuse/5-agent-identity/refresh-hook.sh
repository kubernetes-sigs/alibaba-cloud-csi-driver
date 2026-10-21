#!/bin/bash
set -e

# Demo 5: run by mount-proxy after every credential rotation.
#
# The new credential arrives on stdin as four lines:
#   AccessKeyId
#   AccessKeySecret
#   SecurityToken
#   Expiration

: "${source:?source not set}"

read -r ak
read -r sk
read -r token
read -r expiration

juicefs config --yes "$source" \
    --access-key="$ak" \
    --secret-key="$sk" \
    --session-token="$token"

echo "pushed rotated credential to $source (expires $expiration)"
