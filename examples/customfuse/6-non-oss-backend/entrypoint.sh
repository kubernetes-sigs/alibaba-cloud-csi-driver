#!/bin/bash
set -e

# Demo 6: agent-memory-fuse with agent-identity credentials.
#
# The credential arrives on stdin as four lines:
#   AccessKeyId
#   AccessKeySecret
#   SecurityToken
#   Expiration
#
# The driver itself does not put a credential value in the filesystem, argv or
# environment; what happens after you read stdin is up to this script. Read all
# four lines before handing off, otherwise the unread lines stay in the pipe
# and are inherited by the process it exec's into.

read -r ak
read -r sk
read -r token
read -r expiration

: "${endpoint:?endpoint not set}"
: "${instance:?instance not set}"
: "${store:?store not set}"
: "${scope:?scope not set}"

session="${session:-$(hostname)-$$}"

echo "Starting agent-memory-fuse: endpoint=$endpoint instance=$instance store=$store"

# agent-memory-fuse takes the credential in --credential, so it appears in
# this container's /proc/<pid>/cmdline. Keep xtrace off around the exec so the
# value does not also reach the log when someone debugs this script with `bash -x`.
{ set +x; } 2>/dev/null
exec /usr/local/bin/agent-memory-fuse mount --endpoint "$endpoint" --instance "$instance" --credential "sts:${ak}:${sk}:${token}" --store "$store" --scope "$scope" --session "$session" "$mountpoint"
