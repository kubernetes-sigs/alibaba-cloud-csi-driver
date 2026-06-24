#!/bin/bash
set -e

# Demo 2: OSS-compatible — config from PV volumeAttributes + Kubernetes Secret.
#
# This entrypoint follows the same pattern as OSS static volumes:
# all parameters come from volumeAttributes (source, url, bucket, otherOpts)
# and credentials from Secret.
#
# Environment variables (set by mount-proxy):
#   source     - mount source (volumeAttributes.source)
#   url        - storage endpoint (volumeAttributes.url)
#   bucket     - bucket name (volumeAttributes.bucket)
#   otherOpts  - mount options string (volumeAttributes.otherOpts)
#   path       - sub-path (volumeAttributes.path)
#   readOnly   - "true" if PV is read-only
#   mountpoint - mount target path (managed by CSI)
#
# Secret env vars:
#   accessKeyId     - access key
#   accessKeySecret - secret key

# ── Input validation ──
validate_opts() {
    for opt in "$@"; do
        if [[ "$opt" =~ [;\|\&\`\$\(\)$'\n'] ]]; then
            echo "ERROR: invalid character in option: $opt" >&2
            exit 1
        fi
    done
}
if [ -n "$otherOpts" ]; then
    IFS=',' read -ra _OPTS <<< "$otherOpts"
    validate_opts "${_OPTS[@]}"
fi

echo "Formatting JuiceFS volume: $source"
juicefs format \
    --storage=oss \
    --bucket="http://$bucket.$url" \
    --access-key="$accessKeyId" \
    --secret-key="$accessKeySecret" \
    "$source" \
    myjfs

# ── Credential cleanup ──
unset accessKeyId accessKeySecret akId akSecret

# Build -o options string for mount.juicefs.
# foreground: keep FUSE process in foreground so mount-proxy can track its lifecycle.
# no-update: skip JuiceFS version check on startup — version is pinned by the container image.
# These two are always set here; no need to duplicate them in $mountOptions.
MOUNT_OPTS="foreground,no-update"
[ -n "$path" ] && MOUNT_OPTS="${MOUNT_OPTS},subdir=${path}"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS},ro"
[ -n "$otherOpts" ] && MOUNT_OPTS="${MOUNT_OPTS},${otherOpts}"

echo "Mounting at $mountpoint with -o $MOUNT_OPTS"
exec mount.juicefs "$source" "$mountpoint" -o "$MOUNT_OPTS"
