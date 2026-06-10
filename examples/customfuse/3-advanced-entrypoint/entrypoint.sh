#!/bin/bash
set -e

# Demo 3: Advanced entrypoint — business-specific safety controls
#
# User story:
#   The platform team wants to give users control over JuiceFS cache parameters
#   via PV otherOpts, but needs two safety guarantees:
#
#   1. OOM PREVENTION: Users may set --cache-size too large (JuiceFS default is
#      102400 MiB = 100GB). On nodes with limited disk, this fills up /var and
#      causes kubelet eviction. The entrypoint caps cache-size to a safe maximum.
#
#   2. CREDENTIAL CLEANUP: In Demo 2, $accessKeyId and $accessKeySecret remain
#      as env vars for the entire lifetime of the mount process. Anyone with
#      access to the pod (kubectl exec, /proc/1/environ via hostPath) can read
#      them. However, JuiceFS only needs credentials during `format` — after
#      format, AK/SK are stored in the metadata engine and `mount` does not
#      need them. This entrypoint unsets credentials after format so the
#      long-running mount process has no secrets in its environment.
#
# This demonstrates that entrypoint.sh is the right place to implement
# business-specific safety logic that the CSI driver itself should not enforce.
#
# Environment variables (set by mount-proxy):
#   source          - JuiceFS META-URL (e.g. redis://host:6379/1)
#   bucket          - OSS bucket name
#   url             - OSS endpoint
#   mountpoint      - mount target path
#   otherOpts       - user options string, e.g. "--cache-size=2048 --writeback"
#   accessKeyId     - OSS access key (from Secret)
#   accessKeySecret - OSS secret key (from Secret)
#
# Reference:
#   --cache-size: https://juicefs.com/docs/community/command_reference#mount
#   --cache-dir:  https://juicefs.com/docs/community/command_reference#mount

# ── 1. Parse otherOpts into individual flags ──
# Users pass otherOpts as a flat string in PV volumeAttributes.
# Parsing it allows the entrypoint to inspect, override, or cap values.
declare -A OPTS
EXTRA_ARGS=()

parse_opts() {
    local args=($otherOpts)
    local i=0
    while [ $i -lt ${#args[@]} ]; do
        local arg="${args[$i]}"
        case "$arg" in
            --*=*)
                local key="${arg%%=*}"
                local val="${arg#*=}"
                OPTS["$key"]="$val"
                ;;
            --*)
                if [ $((i+1)) -lt ${#args[@]} ] && [[ ! "${args[$((i+1))]}" == --* ]]; then
                    OPTS["$arg"]="${args[$((i+1))]}"
                    i=$((i+1))
                else
                    OPTS["$arg"]="true"
                fi
                ;;
            *)
                EXTRA_ARGS+=("$arg")
                ;;
        esac
        i=$((i+1))
    done
}

parse_opts

# ── 2. OOM prevention: cap cache-size ──
# JuiceFS --cache-size default is 102400 MiB (100GB). On nodes with small disks
# (e.g. 40GB system disk), an uncapped cache fills /var and triggers kubelet
# eviction of all pods on the node.
#
# The platform team sets a safe maximum here (e.g. 20GB). Users can request
# any value via otherOpts, but it will be silently capped.
MAX_CACHE_SIZE_MB=20480
: "${OPTS[--cache-size]:=1024}"
: "${OPTS[--cache-dir]:=/var/jfsCache}"

if [ "${OPTS[--cache-size]}" -gt "$MAX_CACHE_SIZE_MB" ] 2>/dev/null; then
    echo "WARN: --cache-size=${OPTS[--cache-size]} exceeds node limit ${MAX_CACHE_SIZE_MB}MB, capping"
    OPTS[--cache-size]=$MAX_CACHE_SIZE_MB
fi

mkdir -p "${OPTS[--cache-dir]}"

# ── 3. Format (uses credentials, idempotent) ──
echo "Formatting JuiceFS volume: $source"
FORMAT_ARGS=(--storage=oss)
[ -n "$bucket" ] && [ -n "$url" ] && FORMAT_ARGS+=(--bucket="http://$bucket.$url")

juicefs format \
    "${FORMAT_ARGS[@]}" \
    --access-key="$accessKeyId" \
    --secret-key="$accessKeySecret" \
    "$source" \
    myjfs 2>/dev/null || true

# ── 4. Credential cleanup ──
# After format, AK/SK are persisted in the metadata engine. The long-running
# mount process does not need them. Unset from environment so they are not
# visible via /proc/1/environ or `kubectl exec <pod> -- env`.
#
# Comparison with Demo 2:
#   Demo 2: credentials remain in env for entire container lifetime.
#   Demo 3: credentials exist only during the brief format step above.
unset accessKeyId accessKeySecret akId akSecret

# ── 5. Rebuild mount arguments ──
MOUNT_ARGS=(-f)
for key in "${!OPTS[@]}"; do
    val="${OPTS[$key]}"
    if [ "$val" = "true" ]; then
        MOUNT_ARGS+=("$key")
    else
        MOUNT_ARGS+=("$key=$val")
    fi
done
MOUNT_ARGS+=("${EXTRA_ARGS[@]}")

echo "Mounting JuiceFS at $mountpoint (cache-size=${OPTS[--cache-size]}MB, cache-dir=${OPTS[--cache-dir]})"
exec juicefs mount \
    "${MOUNT_ARGS[@]}" \
    "$source" \
    "$mountpoint"
