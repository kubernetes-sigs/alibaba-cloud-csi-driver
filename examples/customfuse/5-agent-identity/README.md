# Demo 5: Agent identity on an ACS sandbox, scripts baked into the image

Run a self-built FUSE client in an Alibaba Cloud ACS sandbox with no credential
anywhere in the cluster. mount-proxy exchanges the sandbox's internal credential
for a scoped STS credential, delivers it via stdin, and rotates it
before it expires.

Here both scripts live in your image, which keeps this the smallest working
setup. **To work out what your own client needs from rotation, read
[Which case is your client?](#which-case-is-your-client)** below — it covers the
three behaviours a client can have. This demo ships the middle one, because it is
the case that needs the most work. For the mechanism and the responsibility
boundary see
[docs/customfuse-agent-identity.md](../../../docs/customfuse-agent-identity.md).

This demo covers the credential and nothing else. If your volumes also carry a
capacity, [../3-standard/entrypoint.sh](../3-standard/entrypoint.sh) shows
converting `$capacity` into a `juicefs quota set` — add that step here too.

## How this runs in a sandbox

The CSI driver is not a DaemonSet here — it is injected into the sandbox pod:

```
the sandbox is created with the customfuse runtime
    │
    └─ the platform injects into every sandbox pod:
         csi-sidecar                    CSI plugin, serves NodePublish
         csi-customfuse-agent-sidecar   ← YOUR IMAGE, runs mount-proxy

the sandbox attaches the PV as a volume
    │
    └─ the PV resolves to driver=customfuseplugin
         └─ NodePublish on the in-pod socket
              └─ csi-sidecar → mount-proxy socket
                   └─ your /entrypoint.sh runs, credential already on disk
                        └─ FUSE mount, visible to the agent container
```

Your image supplies the `csi-customfuse-agent-sidecar`, so it must carry
`csi-mount-proxy-server` as its ENTRYPOINT, your FUSE client, and the scripts —
see [Dockerfile](Dockerfile). That image, and how the sandbox comes to run it, are
the sandbox product's side; this demo covers only what the driver does once your
`/entrypoint.sh` is reached.

## Which case is your client?

The driver's side of rotation is fixed: a new credential appears in the same
directory before the old one expires. Everything after that is your client's
behaviour, and there are only three. This demo ships the middle one — JuiceFS —
because it is the case that needs the most work.

| | JindoFS | JuiceFS (this demo) | s3fs |
|---|---|---|---|
| **Notices a rotation?** | **Yes**, if your hook writes files it polls | **No**, must be told | **No, and cannot be told** |
| **Ship `/refresh-hook.sh`?** | **yes** — writes files from stdin, JindoFS polls them | **yes** — runs `juicefs config` | no — a hook cannot help |
| **`/entrypoint.sh` reads stdin?** | yes — writes files for the SDK, then mounts | yes — passes values as arguments | yes — writes an AWS ini profile |

Find the row that matches your client and start from this demo's
[entrypoint.sh](entrypoint.sh) / [refresh-hook.sh](refresh-hook.sh).

Both the entrypoint and the hook receive the credential on **stdin** as four
lines: AccessKeyId, AccessKeySecret, SecurityToken, Expiration. Read them with
`read -r`.

### If your client polls credential files

Your entrypoint and hook both write the stdin lines to files. The SDK polls them.

If you write credential files, consider mounting a `tmpfs` at the path you
write to (e.g. `emptyDir` with `medium: Memory`), so the files do not land in
the container's writable layer.

**entrypoint.sh** (JindoFS):
```bash
read -r ak; read -r sk; read -r token; read -r expiration
mkdir -p /run/creds
echo "$ak" > /run/creds/AccessKeyId
echo "$sk" > /run/creds/AccessKeySecret
echo "$token" > /run/creds/SecurityToken

cat > /jindosdk.cfg <<EOF
[jindosdk]
fs.oss.endpoint = ${url}
aliyun.oss.provider.url = secrets:///run/creds/
EOF

exec 0<&-
exec jindo-fuse -f -o "uri=$source" "$mountpoint"
```

**refresh-hook.sh** (JindoFS):
```bash
read -r ak; read -r sk; read -r token; read -r expiration
echo "$ak" > /run/creds/AccessKeyId
echo "$sk" > /run/creds/AccessKeySecret
echo "$token" > /run/creds/SecurityToken
echo "rotated credential files for JindoFS (expires $expiration)"
```

### If your client must be told

The entrypoint reads stdin and mounts once. The hook reads stdin on each rotation
and pushes the credential to the client. The hook is an arbitrary executable, so
"tell it" is whatever your client supports — its CLI, a signal, an admin endpoint,
rewriting its own config. It receives the volume's attributes (`$mountpoint`,
`$source`, etc.) as environment variables.

Note what [entrypoint.sh](entrypoint.sh) has to do twice: `juicefs format` creates
the volume but **skips itself** once the volume exists, credential included, so
`juicefs config` is what actually installs the current credential on every mount.
That kind of client-specific quirk is exactly what the adapter absorbs.

### If your client cannot rotate at all

Read stdin, write whatever your client needs, and accept that a rotation will not
reach the running client; plan for the mount to be re-established, or give it a
credential lifetime longer than the mount. s3fs is the example: it reads the
standard AWS credentials file, so the entrypoint reshapes the four stdin lines
into one profile —

```bash
mkdir -p /root/.aws
umask 077
cat > /root/.aws/credentials <<EOF
[default]
aws_access_key_id = $ak
aws_secret_access_key = $sk
aws_session_token = $token
EOF
```

— and then nothing can update it. It is worth knowing this case exists, because no
hook fixes it.

**A mount that works and then loses access some minutes later is the signature of a
rotation your client did not pick up**: it belongs in the "must be told" row, and
you shipped no hook.

## What is assumed ready

The sandbox itself: set up per the
[Agent Sandbox](https://help.aliyun.com/en/cs/user-guide/agent-sandbox) guide, and
running the image you build below as its customfuse sidecar. None of that is
configured from this repo, and none of it is per-volume.

What this demo adds on top of it is two things: the PV, and the two scripts. If a
sandbox comes up without the sidecars, or the volume never attaches, that is the
sandbox side rather than anything here.

## Setup

```bash
# 1. build your image — the demo directory is the whole build context, and
#    mount-proxy comes from an image you pass in rather than being built here.
#    Where that image comes from: ../README.md, "Step 2: Build the fuse pod image".
docker build --build-arg MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag> \
    -t <your-image>:<tag> examples/customfuse/5-agent-identity
docker push <your-image>:<tag>

# 2. the volume — no Secret, no ConfigMap
kubectl apply -f examples/customfuse/5-agent-identity/pv.yaml

# 3. attach the PV to your sandbox — how to do that is the sandbox product's own
#    step, not something this repo configures. Two things about it matter here:
#    the sandbox attaches the PV rather than a PVC of yours binding it, and that
#    attach is what triggers NodePublish; and the credential provider is named on
#    the sandbox side, which is why pv.yaml declares none.
```

## Verifying

```bash
NS=default                       # your sandbox's namespace
POD=<your sandbox pod>           # the pod running your agent workload
AGENT=<your agent container>     # the container the mount should be visible in

# The two sidecar names below are what the sandbox injects; read them off your own
# pod rather than trusting them here:
#   kubectl -n "$NS" get "$POD" -o jsonpath='{.spec.containers[*].name}'

# the credential is delivered via stdin, so there are no files to inspect.
# check the sidecar logs to confirm the credential was exchanged and delivered.

# the exchange, the rotations, and the hook. Failures always print; the success
# lines are klog V(2)/V(4) and the sidecar's verbosity is set by the platform,
# so an empty result is not evidence of a problem. Expiration moving forward is.
kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar | grep -i 'credential\|refresh\|entrypoint'

# NodePublish handling
kubectl -n "$NS" logs "$POD" -c csi-sidecar --tail=50

# the mount as the agent sees it
kubectl -n "$NS" exec "$POD" -c "$AGENT" -- df -h /mnt/data
```

`Expiration` should move forward roughly `AGENT_IDENTITY_TOKEN_REFRESH_MARGIN`
(default `20m`) before the previous value is reached.

## Testing the read pattern

`entrypoint.sh` and `refresh-hook.sh` read the credential from stdin as four
lines. All four fields arrive atomically in a single write, so there is no risk
of mixing fields from two different rotations.

`TestRotateTokenFiles_PublishedGenerationIsNeverRewritten` in
[ossfs_secret_test.go](../../../pkg/mounter/interceptors/ossfs_secret_test.go) pins a
published generation, rotates twenty more past it, and requires that nothing
reachable through the pin changed. `make test` runs it.
