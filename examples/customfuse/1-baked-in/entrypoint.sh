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
    myjfs 2>/dev/null || true

echo "Mounting JuiceFS at $mountpoint"
exec juicefs mount \
    -f \
    --cache-size="$CACHE_SIZE" \
    "$META_URL" \
    "$mountpoint"
