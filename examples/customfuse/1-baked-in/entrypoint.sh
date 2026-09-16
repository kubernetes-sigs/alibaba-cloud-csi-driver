#!/bin/bash
set -e

# Demo 1: Baked-in — all config comes from Dockerfile ENV, not from PV/Secret.
#
# Only $mountpoint is set by mount-proxy (the mount target path).
# Everything else is baked into the image via Dockerfile ENV directives.

echo "Formatting/authenticating JuiceFS volume..."
juicefs format \
    --storage=oss \
    --bucket="http://$BUCKET.$URL" \
    --access-key="$ACCESS_KEY" \
    --secret-key="$SECRET_KEY" \
    "$META_URL" \
    myjfs

MOUNT_OPTS="cache-size=$CACHE_SIZE"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS},ro"

if [ -n "$CAPACITY" ]; then
    echo "Setting quota: capacity=${CAPACITY}GiB"
    juicefs quota set "$META_URL" --path / --capacity "$CAPACITY"
fi

echo "Mounting JuiceFS at $mountpoint"
# Absolute path: JuiceFS only acts as the mount helper when argv[0] ends in
# "/mount.juicefs". JFS_FOREGROUND: the helper daemonizes otherwise, and
# mount-proxy tracks the entrypoint process.
export JFS_FOREGROUND=1
exec /bin/mount.juicefs "$META_URL" "$mountpoint" -o "$MOUNT_OPTS"
