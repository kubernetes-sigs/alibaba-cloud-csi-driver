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
# The driver puts no credential value in the environment or in this script's
# arguments. What happens after you read them is up to this script; the
# patterns below are what we suggest.
#
# Read all four lines. Anything left unread stays in the pipe and is inherited
# by whatever this script exec's into.
read -r ak
read -r sk
read -r token
read -r expiration

# juicefs takes the credential as command-line flags, so it appears in this
# container's /proc/<pid>/cmdline while the command runs. Keep xtrace off
# around these calls so the values do not also reach the log when someone
# debugs this script with `bash -x`. The `{ ...; } 2>/dev/null` form keeps the
# `set +x` itself from being traced.
configure_juicefs() {
    { set +x; } 2>/dev/null

    juicefs format \
        --storage=oss \
        --bucket="http://$bucket.$url" \
        --access-key="$ak" \
        --secret-key="$sk" \
        --session-token="$token" \
        "$source" myjfs || return $?

    juicefs config --yes "$source" \
        --access-key="$ak" \
        --secret-key="$sk" \
        --session-token="$token"
}

if configure_juicefs; then
    echo "juicefs configured (credential expires $expiration)"
else
    echo "juicefs configuration failed" >&2
    exit 1
fi

unset ak sk token expiration

MOUNT_OPTS=""
[ -n "$path" ] && MOUNT_OPTS="subdir=${path}"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}ro"

echo "Mounting at $mountpoint with options: ${MOUNT_OPTS:-none}"
export JFS_FOREGROUND=1
# stdin is drained above; close it so the mount process does not inherit it.
exec 0<&-
exec /bin/mount.juicefs "$source" "$mountpoint" ${MOUNT_OPTS:+-o "$MOUNT_OPTS"}
