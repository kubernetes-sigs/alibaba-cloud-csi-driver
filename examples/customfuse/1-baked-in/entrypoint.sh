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

echo "Mounting JuiceFS at $mountpoint"
# foreground: required by mount-proxy; no-update: skip version check (pinned by image).
MOUNT_OPTS="foreground,no-update,cache-size=$CACHE_SIZE"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS},ro"
exec mount.juicefs "$META_URL" "$mountpoint" -o "$MOUNT_OPTS"
