#!/bin/bash
set -e

# Demo 5: run by mount-proxy after every credential rotation.
#
# A client that re-reads the credential files needs no hook at all. This one
# exists because JuiceFS does not: its credential lives in the metadata engine,
# so rewriting the files reaches nothing that is watching them. Only
# `juicefs config` makes a new credential take effect, and running clients
# converge within about a minute of it, without interrupting IO.
#
# The hook receives $credentialDir plus this volume's attributes — $mountpoint,
# $source and whatever the PV set — so it can name the volume it is being asked
# about. No credential value reaches it, so nothing lands in /proc/<pid>/cmdline
# or in this script's environment.

: "${credentialDir:?credentialDir not set}"
: "${source:?source not set}"

# One resolved directory for all fields, same reason as in entrypoint.sh.
dir=$(readlink -f "$credentialDir")

juicefs config --yes "$source" \
    --access-key="$(cat "$dir/AccessKeyId")" \
    --secret-key="$(cat "$dir/AccessKeySecret")" \
    --session-token="$(cat "$dir/SecurityToken")"

echo "pushed rotated credential to $source (expires $(cat "$dir/Expiration"))"
