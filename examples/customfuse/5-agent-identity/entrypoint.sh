#!/bin/bash
set -e

# Demo 5: Agent identity — the credential arrives as files, not as env vars.
#
# mount-proxy has already written the credential before this runs, and keeps
# rewriting it before it expires. $credentialDir points at the directory; no
# AccessKey is ever passed as an environment variable or an argument.
#
# JuiceFS is used here as the example of a client that does not read credential
# files at all: it takes the credential as command line arguments and keeps it in
# its metadata engine. That is also why this demo ships refresh-hook.sh — see
# there for what a rotation needs.

: "${credentialDir:?mount-proxy did not provide credentialDir}"

# credentialDir is a symlink a rotation replaces. Resolve it once and read every
# field from the result, or the fields could come from two different rotations —
# a credential that never existed, with no error raised. Retry because the
# directory a rotation replaced is deleted soon after.
read_credential() {
    local dir
    dir=$(readlink -f "$credentialDir") || return 1
    ak=$(cat "$dir/AccessKeyId") || return 1
    sk=$(cat "$dir/AccessKeySecret") || return 1
    token=$(cat "$dir/SecurityToken") || return 1
}

for attempt in 1 2 3; do
    read_credential && break
    [ "$attempt" = 3 ] && { echo "cannot read credential from $credentialDir" >&2; exit 1; }
    sleep 1
done

# Creates the volume on first use. On every later mount the volume already exists
# and format SKIPS itself, credential included — so this cannot be relied on to
# install the current credential.
juicefs format \
    --storage=oss \
    --bucket="http://$bucket.$url" \
    --access-key="$ak" \
    --secret-key="$sk" \
    --session-token="$token" \
    "$source" myjfs

# Hence this: config is what actually writes the credential to the metadata
# engine, and it is required on every mount — by the time a pod is rescheduled,
# the token that formatted the volume is long gone.
# --yes because there is no tty here to answer a prompt on.
juicefs config --yes "$source" \
    --access-key="$ak" \
    --secret-key="$sk" \
    --session-token="$token"

unset ak sk token

# No credential here: the volume owns it, not the mount.
# Only real JuiceFS CE mount options belong here; keeping the client in the
# foreground is done with JFS_FOREGROUND below, not with an -o value.
MOUNT_OPTS=""
[ -n "$path" ] && MOUNT_OPTS="subdir=${path}"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}ro"

echo "Mounting at $mountpoint with options: ${MOUNT_OPTS:-none}, credentials from $credentialDir"
# Absolute path: JuiceFS only acts as the mount helper when argv[0] ends in
# "/mount.juicefs". JFS_FOREGROUND: the helper daemonizes otherwise, and
# mount-proxy tracks the entrypoint process.
export JFS_FOREGROUND=1
exec /bin/mount.juicefs "$source" "$mountpoint" ${MOUNT_OPTS:+-o "$MOUNT_OPTS"}
