#!/bin/bash
set -e

# Demo 5: Agent identity — the credential arrives on stdin, not as env vars.
#
# mount-proxy writes four lines to this script's stdin before it runs:
#   AccessKeyId
#   AccessKeySecret
#   SecurityToken
#   Expiration
#
# No credential value touches the filesystem, argv or environment.

read -r ak
read -r sk
read -r token
read -r expiration

juicefs format \
    --storage=oss \
    --bucket="http://$bucket.$url" \
    --access-key="$ak" \
    --secret-key="$sk" \
    --session-token="$token" \
    "$source" myjfs

juicefs config --yes "$source" \
    --access-key="$ak" \
    --secret-key="$sk" \
    --session-token="$token"

unset ak sk token expiration

MOUNT_OPTS=""
[ -n "$path" ] && MOUNT_OPTS="subdir=${path}"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}ro"

echo "Mounting at $mountpoint with options: ${MOUNT_OPTS:-none}"
export JFS_FOREGROUND=1
exec /bin/mount.juicefs "$source" "$mountpoint" ${MOUNT_OPTS:+-o "$MOUNT_OPTS"}
