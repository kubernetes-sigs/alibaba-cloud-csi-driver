#!/bin/bash
set -e

# Demo 2: OSS-compatible — config from PV volumeAttributes + Kubernetes Secret.
#
# This entrypoint follows the same pattern as OSS static volumes:
# all parameters come from volumeAttributes (source, url, bucket, otherOpts)
# and credentials from Secret.
#
# Environment variables (set by mount-proxy):
#   source     - mount source (volumeAttributes.source; composed below when absent)
#   url        - storage endpoint (volumeAttributes.url)
#   bucket     - bucket name (volumeAttributes.bucket)
#   otherOpts  - mount options string (volumeAttributes.otherOpts)
#   path       - sub-path (volumeAttributes.path)
#   capacity   - volume quota (volumeAttributes.capacity, or the PVC's size)
#   readOnly   - "true" if PV is read-only
#   mountpoint - mount target path (managed by CSI)
#
#   extraOpts  - appended after $otherOpts. Comes from an "extraOpts=..." entry
#                in pv.spec.mountOptions, which unlike volumeAttributes stays
#                editable after the PV exists, so it is how an existing volume
#                gains a mount option without being recreated. Same format as
#                $otherOpts: comma separated, no "-o" prefix, because this
#                script passes the whole thing as one -o value below.
#
# Secret env vars, one per Secret key under that same name:
#   accessKeyId     - access key
#   accessKeySecret - secret key
#   akId, akSecret  - older spellings, resolved to the two above right below and
#                     cleaned up together with them

# ── OSS-compatible conventions ──
# The driver hands volume data over exactly as it arrived: it composes nothing and
# knows no credential spellings. Both of the following belong to the shape of an
# OSS static volume rather than to mounting a FUSE filesystem, so this adapter
# applies them here, where they are known, and a PV written for OSS works unchanged.
: "${accessKeyId:=$akId}"
: "${accessKeySecret:=$akSecret}"
if [ -z "$source" ] && [ -n "$bucket" ]; then
    source="${bucket}${path:+:$path}"
fi

# ── Input validation ──
# The character set lives in a variable: bash cannot parse `$'\n'` inside a [[ =~ ]] class.
FORBIDDEN=';|&`$()'
FORBIDDEN+=$'\n'
validate_opts() {
    for opt in "$@"; do
        case "$opt" in
            *["$FORBIDDEN"]*)
                echo "ERROR: invalid character in option: $opt" >&2
                exit 1 ;;
        esac
    done
}
for _OPTSTR in "$otherOpts" "$extraOpts"; do
    if [ -n "$_OPTSTR" ]; then
        IFS=',' read -ra _OPTS <<< "$_OPTSTR"
        validate_opts "${_OPTS[@]}"
    fi
done

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
# Only real JuiceFS CE mount options belong here; keeping the client in the
# foreground is done with JFS_FOREGROUND below, not with an -o value.
MOUNT_OPTS=""
[ -n "$path" ] && MOUNT_OPTS="subdir=${path}"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}ro"
[ -n "$otherOpts" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}${otherOpts}"
# Last, matching the order the two channels were written in: $otherOpts is what
# the PV was created with, $extraOpts what was added to it later. No "-o" prefix:
# everything above becomes a single -o argument on the exec line below.
[ -n "$extraOpts" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}${extraOpts}"

# $capacity may have units (e.g. "100Gi" from PV or auto-capacity feature gate).
# juicefs quota --capacity takes a GiB integer, convert common Quantity suffixes.
# NOTE: This only handles the most common units. Extend as needed for your use case.
if [ -n "$capacity" ]; then
    case "$capacity" in
        *TiB|*Ti) capacity=$(( ${capacity%%[A-Za-z]*} * 1024 )) ;;
        *GiB|*Gi) capacity=${capacity%%[A-Za-z]*} ;;
        *MiB|*Mi) capacity=$(( ${capacity%%[A-Za-z]*} / 1024 )) ;;
        *[A-Za-z]*)
            echo "ERROR: unsupported capacity unit: $capacity (expected GiB/TiB/MiB or plain integer)" >&2
            exit 1 ;;
    esac
    echo "Setting quota: capacity=${capacity}GiB path=${path:-/}"
    juicefs quota set "$source" --path "${path:-/}" --capacity "$capacity"
fi

echo "Mounting at $mountpoint with options: ${MOUNT_OPTS:-none}"
# Absolute path: JuiceFS only acts as the mount helper when argv[0] ends in
# "/mount.juicefs". JFS_FOREGROUND: the helper daemonizes otherwise, and
# mount-proxy tracks the entrypoint process.
export JFS_FOREGROUND=1
exec /bin/mount.juicefs "$source" "$mountpoint" ${MOUNT_OPTS:+-o "$MOUNT_OPTS"}
