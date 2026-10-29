# Demo 6 — Non-OSS backend with agent-identity

This demo shows that the customfuse driver is **not limited to OSS-compatible
storage**. Any FUSE client that accepts credentials and a mount point can be
used — the driver passes arbitrary volume attributes as environment variables
and delivers agent-identity credentials via stdin.

The example uses [agent-memory-fuse](https://help.aliyun.com/en/tablestore/memory-storage-service-sub-product-test/file-memory-1)
(Tablestore Agent Memory) as a concrete non-OSS backend. The same pattern
applies to any other non-OSS FUSE client.

This demo uses agent-identity for credential delivery. For fixed credentials
(AccessKey in a Secret), see [Demo 2](../2-oss-compatible/) — the same
Secret-passthrough approach works for non-OSS FUSE clients unchanged.

## How it differs from Demo 5 (JuiceFS on OSS)

| | Demo 5 (JuiceFS / OSS) | Demo 6 (agent-memory-fuse / Tablestore) |
|---|---|---|
| **bucket** | Required (OSS bucket name) | Not used |
| **url** | Required (OSS endpoint) | Not used |
| **source** | JuiceFS metadata URL | Not used |
| **endpoint** | — | Tablestore HTTPS endpoint |
| **instance** | — | Tablestore instance name |
| **store** | — | Memory store name |
| **scope** | — | `appId#tenantId#agentId#runId` |
| **Credential delivery** | stdin (4 lines) | stdin (4 lines) — identical |
| **Credential rotation** | `juicefs config` via hook | `agent-memory-fuse rotate-credential` via hook |

The driver does not validate or interpret any of these parameters. They are
passed through as environment variables (`$endpoint`, `$instance`, etc.) to the
entrypoint and hook scripts, which construct the appropriate command lines.

## Volume attributes

```yaml
volumeAttributes:
  fuseType: agent-memory-fuse        # selects the container image
  authType: agent-identity           # use sandbox token exchange
  sandboxCredProviderName: my-provider

  # Tablestore-specific — no bucket, no url, no source
  endpoint: https://my-instance.cn-hangzhou.ots.aliyuncs.com
  instance: my-instance
  store: agent_files
  scope: "app-001#user-001#assistant#session-001"
  session: fuse-session-001
```

## Entrypoint

[entrypoint.sh](entrypoint.sh) reads the credential from stdin and starts
`agent-memory-fuse mount` with the `--credential sts:<ak>:<sk>:<token>` flag:

```bash
read -r ak
read -r sk
read -r token
read -r expiration

# The credential appears in /proc/<pid>/cmdline because agent-memory-fuse
# takes it as a flag. Keep xtrace off so the value does not reach the log.
{ set +x; } 2>/dev/null
exec 0<&-
exec /usr/local/bin/agent-memory-fuse mount \
    --endpoint "$endpoint" \
    --instance "$instance" \
    --credential "sts:${ak}:${sk}:${token}" \
    --store "$store" \
    --scope "$scope" \
    --session "$session" \
    "$mountpoint"
```

Read all four stdin lines before handing off, and close stdin with `exec 0<&-`
so nothing is inherited by the mount process. The driver itself does not put a
credential value in the filesystem, argv or environment; what happens after you
read stdin is up to the entrypoint script.

## Credential rotation

[refresh-hook.sh](refresh-hook.sh) uses `agent-memory-fuse rotate-credential`,
which reads a JSON credential from stdin — matching the tool's native rotation
API. The credential never reaches the command line in this path:

```bash
read -r ak
read -r sk
read -r token
read -r expiration

printf '{"accessKeyId":"%s","accessKeySecret":"%s","stsToken":"%s"}' \
    "$ak" "$sk" "$token" | \
    /usr/local/bin/agent-memory-fuse rotate-credential --mountpoint "$mountpoint"
```

## Building the image

See [Dockerfile](Dockerfile). Build `csi-mount-proxy` once from the repo's
`build/mount-proxy/Dockerfile` (customfuse-base target), then pass it in:

```bash
# Build mount-proxy (once, shared across all demos)
docker build --target customfuse-base -f build/mount-proxy/Dockerfile \
    -t <registry>/csi-mount-proxy:<tag> .

# Build this demo's image
docker build --build-arg MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag> \
    -t <image>:<tag> examples/customfuse/6-agent-memory-fuse
```

The mount-proxy-server is the container's ENTRYPOINT — it listens on a Unix
socket for mount requests from the CSI plugin, then starts `/entrypoint.sh`
with volume attributes as environment variables and the credential on stdin.

## What this proves

The customfuse driver's volume attribute handling is generic:

- `bucket`, `url`, `source` are optional — empty when not provided
- Arbitrary keys (`endpoint`, `instance`, `store`, `scope`, `session`) pass
  through as environment variables without any driver-side registration
- Agent-identity credential delivery via stdin works identically for OSS and
  non-OSS backends
- The refresh hook mechanism (stdin credential → client-specific rotation
  command) is backend-agnostic
