#!/bin/bash
set -e

# Demo 2: Standard — config from PV volumeAttributes + Kubernetes Secret.
#
# Environment variables set by mount-proxy (names match PV volumeAttributes keys):
#   source     - mount source (from PV 'source'), here: JuiceFS META-URL
#   url        - storage endpoint (from PV 'url'), here: S3/OSS endpoint
#   mountpoint - mount target path (managed by CSI)
#   otherOpts  - extra mount options (from PV 'otherOpts')
#
# Secret env vars (passed directly, name = Secret key):
#   accessKeyId     - S3/OSS access key
#   accessKeySecret - S3/OSS secret key
#
# AK/SK compatibility: if the Secret uses OSS-style keys (akId/akSecret),
# mount-proxy also sets accessKeyId/accessKeySecret automatically.

echo "Formatting/authenticating JuiceFS volume: $source"
juicefs format \
    --storage=oss \
    --bucket="http://$bucket.$url" \
    --access-key="$accessKeyId" \
    --secret-key="$accessKeySecret" \
    "$source" \
    myjfs 2>/dev/null || true

echo "Mounting JuiceFS at $mountpoint"
exec juicefs mount \
    -f \
    $otherOpts \
    "$source" \
    "$mountpoint"
