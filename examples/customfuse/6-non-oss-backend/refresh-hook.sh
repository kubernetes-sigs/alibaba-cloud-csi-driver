#!/bin/bash
set -e

# Demo 6: credential rotation hook for agent-memory-fuse.
#
# agent-memory-fuse supports live credential rotation via its
# rotate-credential subcommand, which reads a JSON object from stdin — so the
# credential never reaches its command line here.
#
# Read all four lines before handing off; see entrypoint.sh for why.

read -r ak
read -r sk
read -r token
read -r expiration

rotate_credential() {
    { set +x; } 2>/dev/null

    printf '{"accessKeyId":"%s","accessKeySecret":"%s","stsToken":"%s"}' \
        "$ak" "$sk" "$token" | \
        /usr/local/bin/agent-memory-fuse rotate-credential --mountpoint "$mountpoint"
}

if rotate_credential; then
    echo "rotated agent-memory-fuse credential (expires $expiration)"
else
    echo "failed to rotate agent-memory-fuse credential" >&2
    exit 1
fi
