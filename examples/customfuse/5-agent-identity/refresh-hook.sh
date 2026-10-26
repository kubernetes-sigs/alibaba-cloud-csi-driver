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

# Same as in entrypoint.sh: juicefs takes the credential as flags, so keep
# xtrace off around the call and report only the outcome.
push_credential() {
    { set +x; } 2>/dev/null

    juicefs config --yes "$source" \
        --access-key="$ak" \
        --secret-key="$sk" \
        --session-token="$token"
}

if push_credential; then
    echo "pushed rotated credential to $source (expires $expiration)"
else
    echo "failed to push rotated credential to $source" >&2
    exit 1
fi
