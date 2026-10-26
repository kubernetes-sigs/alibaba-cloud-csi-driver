#!/bin/bash
# Mock credential refresh hook, run by mount-proxy after every rotation.
#
# It receives the new credential on stdin as four lines and reports whether it
# differs from the last one and whether it still authorises anything. Its own
# stdout is captured by mount-proxy rather than logged, so the findings go to
# $ROTATION_LOG, which mock-entrypoint.sh relays into the log a test reads.

set -euo pipefail

ROTATION_LOG=/tmp/csitest-rotation.log
CRED_STATE=/tmp/csitest-cred
ROTATION_COUNT=/tmp/csitest-rotation-count

read -r ak || true
read -r sk || true
read -r token || true
read -r expiration || true

umask 077

# A provider that caches hands back byte-identical material until its own
# refresh window opens, and a rotation that changes nothing is not a rotation
# worth reporting. The fingerprint covers every field, not just Expiration.
fingerprint="$(printf '%s|%s' "$ak" "$expiration")"
previous=""
[ -s "$CRED_STATE" ] && previous="$(printf '%s|%s' \
    "$(sed -n 1p "$CRED_STATE")" "$(sed -n 4p "$CRED_STATE")")"

if [ "$fingerprint" = "$previous" ]; then
    exit 0
fi

printf '%s\n%s\n%s\n%s\n' "$ak" "$sk" "$token" "$expiration" > "$CRED_STATE"

n=1
[ -s "$ROTATION_COUNT" ] && n=$(($(cat "$ROTATION_COUNT") + 1))
printf '%s' "$n" > "$ROTATION_COUNT"

printf 'CSITEST_CREDROTATED n=%d at=%s\n' "$n" "$(date -u +%H:%M:%SZ)" >> "$ROTATION_LOG"

# Punctuality is not the point: a loop can deliver on schedule and hand over a
# credential that no longer authorises anything, which is what this probe is
# here to catch. The credential goes into a 0600 config rather than onto
# ossutil's argv, which is world-readable through /proc inside this container.
if ! command -v ossutil >/dev/null 2>&1 \
        || [ -z "${bucket:-}" ] || [ -z "${url:-}" ] \
        || [ -z "$ak" ] || [ -z "$sk" ] || [ -z "$token" ]; then
    printf 'CSITEST_OSSPROBE rotated%d=skip:incomplete-credential\n' "$n" >> "$ROTATION_LOG"
    exit 0
fi

cfg="$(mktemp)"
chmod 600 "$cfg"
{
    echo '[Credentials]'
    echo 'language=EN'
    printf 'endpoint=%s\n'        "${url#https://}"
    printf 'accessKeyID=%s\n'     "$ak"
    printf 'accessKeySecret=%s\n' "$sk"
    printf 'stsToken=%s\n'        "$token"
} > "$cfg"

prefix="${path:-/}"
prefix="${prefix#/}"
[ -n "$prefix" ] && prefix="$prefix/"

rc=0
out="$(timeout 20 ossutil -c "$cfg" ls "oss://$bucket/$prefix" --limited-num=1 2>&1)" || rc=$?
rm -f "$cfg"

if [ "$rc" = 0 ]; then
    printf 'CSITEST_OSSPROBE rotated%d=ok prefix=%s\n' "$n" "$prefix" >> "$ROTATION_LOG"
else
    # Only the error code is echoed: ossutil prints the request it made, and
    # that output is not worth trusting with a credential in scope.
    printf 'CSITEST_OSSPROBE rotated%d=fail:prefix=%s:%s\n' "$n" "$prefix" \
        "$(printf '%s' "$out" | grep -oE 'ErrorCode=[A-Za-z]+|StatusCode=[0-9]+' | head -2 | tr '\n' ',' || echo "rc=$rc")" \
        >> "$ROTATION_LOG"
fi
