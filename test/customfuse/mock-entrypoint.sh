#!/bin/bash
# Mock FUSE client, for exercising the customfuse CSI code paths without a real one.
#
# A real client needs a bucket, credentials and a network path to storage, none of
# which say anything about whether the driver did its job. This stands in for it:
# it reports what the driver delivered, then satisfies the one thing mount-proxy
# actually waits for — the target becoming a mount point — by mounting a tmpfs.
#
# Output is prefixed so a test can assert on it instead of a human reading logs:
#
#   CSITEST_ENV <name>=<value>            one line per variable a test asserts on, or the
#                                         literal <unset> if nothing set it. Most are the
#                                         driver's outputs; HOME and HOSTNAME are the
#                                         opposite — a volume is refused the right to
#                                         redefine them, so they report what the
#                                         mount-proxy's own environment gave them.
#   CSITEST_ENVKEY <name>                 every variable present, names only
#   CSITEST_CREDFILE <name>               files under $credentialDir, names only
#   CSITEST_MOUNTED <path>                the tmpfs is up; the mount is now observable
#   CSITEST_CREDROTATED n=<i> at=<time>   the credential on disk was replaced
#   CSITEST_OSSPROBE <label>=…            whether the delivered credential actually works
#
# Values are printed only for the known-safe variables above. Everything else is
# reported by name alone, because Secret entries arrive as environment variables
# too and a test image must never print a credential.

set -euo pipefail

for var in source bucket url path readOnly otherOpts extraOpts capacity fuseType authType \
    credentialDir entrypointConfig mountpoint HOME HOSTNAME; do
    printf 'CSITEST_ENV %s=%s\n' "$var" "${!var-<unset>}"
done

env | cut -d= -f1 | sort | sed 's/^/CSITEST_ENVKEY /'

# For agent-identity: prove the credential landed before this ran, without
# revealing it. The driver writes one file per field.
if [ -n "${credentialDir:-}" ] && [ -d "$credentialDir" ]; then
    ls -1 "$credentialDir" | sed 's/^/CSITEST_CREDFILE /'
fi

# oss_probe reports whether the delivered credential actually works, which is the
# one thing the files alone cannot show: a credential can arrive complete and
# still be scoped to nothing usable.
#
# The credential goes into a 0600 config rather than onto ossutil's argv, because
# argv is world-readable through /proc inside this container and a test image must
# not be the thing that leaks an STS token.
oss_probe() { # $1=label
    local cfg out rc prefix
    for f in AccessKeyId AccessKeySecret SecurityToken; do
        if [ ! -s "$credentialDir/$f" ]; then
            printf 'CSITEST_OSSPROBE %s=skip:no-%s\n' "$1" "$f"
            return
        fi
    done
    if [ -z "${bucket:-}" ] || [ -z "${url:-}" ]; then
        printf 'CSITEST_OSSPROBE %s=skip:no-bucket-or-url\n' "$1"
        return
    fi

    cfg="$(mktemp)"
    chmod 600 "$cfg"
    {
        echo '[Credentials]'
        echo 'language=EN'
        printf 'endpoint=%s\n'        "${url#https://}"
        printf 'accessKeyID=%s\n'     "$(cat "$credentialDir/AccessKeyId")"
        printf 'accessKeySecret=%s\n' "$(cat "$credentialDir/AccessKeySecret")"
        printf 'stsToken=%s\n'        "$(cat "$credentialDir/SecurityToken")"
    } > "$cfg"

    # The provider's policy conditions ListObjects on an oss:Prefix of the form
    # "<sub-path>/*", so the probe has to ask about exactly that: the trailing
    # slash is what makes the requested prefix match the condition. Listing the
    # bare path, or the bucket root, is denied by that condition and would read as
    # an invalid credential when it is only the wrong question.
    prefix="${path:-/}"
    prefix="${prefix#/}"
    [ -n "$prefix" ] && prefix="$prefix/"
    # `out=$(...)` on its own would take errexit down with it when the probe fails,
    # and this process dying tears the mount down with it — mount-proxy tracks the
    # pid. The probe is diagnostic and must never be able to do that.
    rc=0
    out="$(timeout 20 ossutil -c "$cfg" ls "oss://$bucket/$prefix" --limited-num=1 2>&1)" || rc=$?
    rm -f "$cfg"

    if [ "$rc" = 0 ]; then
        printf 'CSITEST_OSSPROBE %s=ok prefix=%s\n' "$1" "$prefix"
    else
        # Only the error code is echoed: ossutil prints the request it made, and
        # that output is not worth trusting with a credential in scope.
        printf 'CSITEST_OSSPROBE %s=fail:prefix=%s:%s\n' "$1" "$prefix" \
            "$(printf '%s' "$out" | grep -oE 'ErrorCode=[A-Za-z]+|StatusCode=[0-9]+' | head -2 | tr '\n' ',' || echo "rc=$rc")"
    fi
}

: "${mountpoint:?mount-proxy did not provide mountpoint}"

# Honouring $readOnly is what makes the driver's own enforcement observable: with the
# client itself mounting read-only, a consumer that binds afterwards inherits that,
# which is the ordering dependence the driver documents rather than hides.
roflag=""
[ "${readOnly:-}" = "true" ] && roflag="ro,"

# mount-proxy polls IsLikelyNotMountPoint until this succeeds, then records the
# process and treats the mount as established. tmpfs needs no backing store, so
# this reaches that state without any storage being involved.
mount -t tmpfs -o "${roflag}size=16m" tmpfs "$mountpoint"
printf 'CSITEST_MOUNTED %s\n' "$mountpoint"

# Leave a marker the test can read through the mount, which is what proves
# propagation reached the workload container. A read-only mount gets none: the test
# asserts on the write being denied instead.
if [ -z "$roflag" ]; then
    echo "customfuse-mock" > "$mountpoint/.csitest-marker"
fi

# Probed only now that the mount is up. mount-proxy waits for the target to become
# a mount point, so anything that can stall — an unreachable endpoint above all —
# has to run after that, or a probe failure becomes a mount failure.
if [ -n "${credentialDir:-}" ] && [ -d "$credentialDir" ] && command -v ossutil >/dev/null 2>&1; then
    oss_probe initial
fi

# Rotation is where an agent-identity setup fails quietly: the refresh loop can
# replace the files on schedule and still hand over a credential that no longer
# authorises anything. Watching for a replacement and probing again is the only way
# to tell a working rotation from a merely punctual one.
#
# The fingerprint covers every field, not just Expiration: a provider that caches
# hands back a byte-identical credential until its own refresh window opens, and
# a rotation that changes nothing is not a rotation worth reporting.
if [ -n "${credentialDir:-}" ] && [ -s "$credentialDir/AccessKeyId" ] && command -v ossutil >/dev/null 2>&1; then
    (
        cred_fingerprint() {
            cat "$credentialDir/AccessKeyId" "$credentialDir/Expiration" 2>/dev/null | tr -d '\n'
        }
        seen="$(cred_fingerprint)"
        n=0
        while sleep 15; do
            now="$(cred_fingerprint)"
            [ -n "$now" ] && [ "$now" != "$seen" ] || continue
            seen="$now"
            n=$((n + 1))
            printf 'CSITEST_CREDROTATED n=%d at=%s\n' "$n" "$(date -u +%H:%M:%SZ)"
            oss_probe "rotated$n"
        done
    ) &
fi

# The mount lives as long as this process: mount-proxy tracks its pid and tears
# the mount down when it exits.
exec sleep infinity
