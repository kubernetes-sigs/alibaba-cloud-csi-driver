#!/bin/bash
set -e

# Demo 3: pv.spec.mountOptions style.
#
# Parameters come from pv.spec.mountOptions. The two below are this script's own
# names, which the driver does not know, so each arrives as an env var of that name:
#   $formatOptions  — comma-separated flags, parsed into --flag args
#   $mountOptions   — comma-separated options, passed directly to mount -o
#
# The driver's own fields reach the script under the same names whichever channel
# the PV used, so $source, $capacity and $readOnly are read below as in Demo 2.
#
# Credentials come from Kubernetes Secret:
#   $accessKeyId / $accessKeySecret

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
if [ -n "$formatOptions" ]; then
    IFS=',' read -ra _OPTS <<< "$formatOptions"
    validate_opts "${_OPTS[@]}"
fi
if [ -n "$mountOptions" ]; then
    IFS=',' read -ra _OPTS <<< "$mountOptions"
    validate_opts "${_OPTS[@]}"
fi

# ── 1. Parse $formatOptions into --flag arguments ──
# Each comma-separated part becomes a --flag:
#   "key=value" → "--key=value"
#   "flag"      → "--flag"
FORMAT_ARGS=(
    --access-key="$accessKeyId"
    --secret-key="$accessKeySecret"
)

if [ -n "$formatOptions" ]; then
    IFS=',' read -ra PARTS <<< "$formatOptions"
    for part in "${PARTS[@]}"; do
        FORMAT_ARGS+=("--${part}")
    done
fi

echo "Formatting volume: $source"
juicefs format "${FORMAT_ARGS[@]}" "$source" myjfs

# ── 2. Credential cleanup ──
unset accessKeyId accessKeySecret akId akSecret

# ── 3. Build mount options ──
# Only real JuiceFS CE mount options belong here; keeping the client in the
# foreground is done with JFS_FOREGROUND below, not with an -o value.
MOUNT_OPTS=""
[ "$readOnly" = "true" ] && MOUNT_OPTS="ro"
[ -n "$mountOptions" ] && MOUNT_OPTS="${MOUNT_OPTS:+${MOUNT_OPTS},}${mountOptions}"

# $capacity may have units (e.g. "100Gi" from auto-capacity feature gate).
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
    echo "Setting quota: capacity=${capacity}GiB"
    juicefs quota set "$source" --path / --capacity "$capacity"
fi

echo "Mounting at $mountpoint with options: ${MOUNT_OPTS:-none}"
# Absolute path: JuiceFS only acts as the mount helper when argv[0] ends in
# "/mount.juicefs". JFS_FOREGROUND: the helper daemonizes otherwise, and
# mount-proxy tracks the entrypoint process.
export JFS_FOREGROUND=1
exec /bin/mount.juicefs "$source" "$mountpoint" ${MOUNT_OPTS:+-o "$MOUNT_OPTS"}
