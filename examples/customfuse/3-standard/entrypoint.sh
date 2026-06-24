#!/bin/bash
set -e

# Demo 3: pv.spec.mountOptions style.
#
# Parameters come from pv.spec.mountOptions. Each entry becomes an env var:
#   $formatOptions  — comma-separated flags, parsed into --flag args
#   $mountOptions   — comma-separated options, passed directly to mount -o
#
# Credentials come from Kubernetes Secret:
#   $accessKeyId / $accessKeySecret

# ── Input validation ──
validate_opts() {
    for opt in "$@"; do
        if [[ "$opt" =~ [;\|\&\`\$\(\)$'\n'] ]]; then
            echo "ERROR: invalid character in option: $opt" >&2
            exit 1
        fi
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
# foreground: keep FUSE process in foreground so mount-proxy can track its lifecycle.
# no-update: skip JuiceFS version check on startup — version is pinned by the container image.
# These two are always set here; no need to duplicate them in $mountOptions.
MOUNT_OPTS="foreground,no-update"
[ "$readOnly" = "true" ] && MOUNT_OPTS="${MOUNT_OPTS},ro"
[ -n "$mountOptions" ] && MOUNT_OPTS="${MOUNT_OPTS},${mountOptions}"

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

echo "Mounting at $mountpoint with -o $MOUNT_OPTS"
exec mount.juicefs "$source" "$mountpoint" -o "$MOUNT_OPTS"
