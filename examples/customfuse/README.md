# CustomFuse CSI Driver

The CustomFuse driver runs **any FUSE client** as a managed pod. You supply the
image + entrypoint; the driver handles pod lifecycle, mount propagation, and
credential injection.

> **Note**: Both static PVs and dynamic provisioning (StorageClass + PVC) work.
> `CreateVolume` provisions no storage — the volume is whatever your entrypoint
> mounts, so there is nothing to create — it turns the StorageClass `parameters`
> into the PV's `volumeAttributes` and the PVC's requested size into `capacity`.
> Two consequences: `DeleteVolume` is a matching no-op, so a StorageClass needs
> `reclaimPolicy: Retain` and the reason is in
> [Dynamic provisioning](#dynamic-provisioning); and there is no expansion, so a
> quota can be raised on an existing PV through `spec.mountOptions` but nothing
> resizes storage underneath it. The demos here are all static PVs.

## Quick Start

### Step 1: Write your entrypoint.sh

Take the mount command you normally run on a node, put it in a script, and
replace hardcoded values with env vars. The driver calls this script with the PV
parameters it knows and the Secret's credentials injected as env vars — see
[Entrypoint Env Vars](#entrypoint-env-vars) for exactly which. The script must
mount at `$mountpoint`, **stay in the foreground**, and **`exec` the client as its
last act** so that the client is the process mount-proxy tracks — see
[Important Notes](#important-notes) item 4 for what happens otherwise:

```bash
#!/bin/bash
set -e

# Secret keys and pv.spec.mountOptions entries arrive as env vars of their own
# name. volumeAttributes forwards only the keys the driver knows — see the
# reference below. $mountpoint is always injected.

# Example: a generic FUSE mount
# Credential passing varies by client — command-line flags, environment
# variables, or a passwd file all work; see the s3fs and JindoFS examples.
exec my-fuse-client "$source" "$mountpoint" \
    --access-key="$accessKeyId" \
    --secret-key="$accessKeySecret" \
    -f
```

See the [demos](#choosing-a-configuration-style) for real entrypoint examples
(JuiceFS, s3fs, JindoFS).

### Step 2: Build the fuse pod image

mount-proxy is the same for every FUSE client, so build it once from this
repository and reuse the image:

```bash
docker build --target customfuse-base -f build/mount-proxy/Dockerfile \
    -t registry.example.com/my-repo/csi-mount-proxy:v0.1 .
```

> **Note**: that build runs in this repository's own build environment. The stages
> `customfuse-base` derives from pull base and toolchain images that are not
> generally reachable, so it will not succeed on an arbitrary machine. If you
> already have a `csi-mount-proxy` image — one your driver installation ships, or
> one a colleague built — use it in place of the build; nothing below depends on
> where it came from, only on its holding the two binaries copied out of it.

Your own image is then just your FUSE client, those two binaries, and your mount
script:

```dockerfile
ARG MOUNT_PROXY_IMAGE                        # ← the image built above
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

FROM <your-fuse-client-base-image>           # ← change this
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/
COPY entrypoint.sh /entrypoint.sh            # ← your mount script
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
```

The build context is the directory holding this Dockerfile and `entrypoint.sh`,
so a copy of it anywhere builds the same image:

```bash
docker build \
    --build-arg MOUNT_PROXY_IMAGE=registry.example.com/my-repo/csi-mount-proxy:v0.1 \
    -t registry.example.com/my-repo/csi-fuse-juicefs:v0.1 .
docker push registry.example.com/my-repo/csi-fuse-juicefs:v0.1
```

### Step 3: Configure the fuse pod image in csi-plugin ConfigMap

If the ConfigMap does not exist yet, create it:

```bash
kubectl -n kube-system create configmap csi-plugin \
  --from-literal="fuse-juicefs=image=registry.example.com/my-repo/csi-fuse-juicefs:v0.1"
```

Or if it already exists, patch it:

```yaml
# kubectl -n kube-system edit configmap csi-plugin
data:
  fuse-juicefs: |
    image=registry.example.com/my-repo/csi-fuse-juicefs:v0.1
```

### Step 4: Create Secret, PV, PVC

```bash
# Edit secret.yaml with your real credentials first!
# Edit pv.yaml's `source` too: `redis-host.default.svc` is a placeholder and this
# repo ships no Redis, so applying the demos as-is fails to connect. Substitute
# your own address — Redis is not required.
kubectl apply -f examples/customfuse/2-oss-compatible/secret.yaml
kubectl apply -f examples/customfuse/2-oss-compatible/pv.yaml
kubectl apply -f examples/customfuse/2-oss-compatible/pvc.yaml
```

### Step 5: Create a consumer Pod and verify

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: juicefs-app
spec:
  containers:
  - name: app
    image: busybox
    command: ["sleep", "infinity"]
    volumeMounts:
    - name: data
      mountPath: /data
  volumes:
  - name: data
    persistentVolumeClaim:
      claimName: juicefs-oss-pvc
```

```bash
kubectl wait --for=condition=Ready pod/juicefs-app --timeout=60s
kubectl exec juicefs-app -- ls /data
```

> The Quick Start above uses the [Demo 2 (OSS-compatible)](#choosing-a-configuration-style)
> pattern. See **Choosing a Configuration Style** below to pick the right
> pattern for your use case.

## Debugging

Fuse pods run in the `ack-csi-customfuse` namespace. Find and inspect them
using the `csi.alibabacloud.com/volume-id` label (value = PV's `volumeHandle`):

```bash
# Find the fuse pod for a specific PV
kubectl -n ack-csi-customfuse get pods \
    -l csi.alibabacloud.com/volume-id=<volumeHandle>

# Check fuse pod logs
kubectl -n ack-csi-customfuse logs \
    -l csi.alibabacloud.com/volume-id=<volumeHandle> \
    -c customfuse
```

Common issues:
- **Pod stuck in Pending**: image pull failure — check `csi-plugin` ConfigMap
  in `kube-system` and verify image registry access.
- **Pod Running but mount fails**: check logs above for entrypoint errors
  (wrong credentials, unreachable metadata server, etc.).
- **`Transport endpoint is not connected`** on consumer pod: the fuse process
  crashed. Check logs for root cause. Delete the consumer pod — the new pod's
  NodePublish triggers mount-proxy to re-run the entrypoint within the same
  fuse pod.

## Operations

### Image upgrade

1. Build and push a new image version
2. Update `csi-plugin` ConfigMap in `kube-system`:
   ```bash
   kubectl -n kube-system edit configmap csi-plugin
   # Change: fuse-juicefs: "image=registry.../csi-fuse-juicefs:v0.2"
   ```
3. Existing fuse pods keep using the old image. New mounts use the new image.
4. To upgrade an existing volume: delete all consumer pods on the node →
   ControllerUnpublish cleans up the fuse pod → re-create consumers →
   ControllerPublish creates a new fuse pod with the updated image.

### Mount recovery

The fuse pod stays running as long as at least one consumer pod on the same
node references the PV. If the fuse process crashes inside the pod,
mount-proxy detects it. Delete the consumer pod to trigger a new NodePublish,
which re-runs the entrypoint within the same fuse pod.

If the fuse pod itself is deleted (node drain, OOM kill, etc.), the CSI
controller automatically re-creates it on the next ControllerPublish (triggered
when a new consumer pod is scheduled to that node).

## Choosing a Configuration Style

The demos use JuiceFS CE as an example, but the patterns apply to any FUSE
client. See [Other FUSE Clients](#other-fuse-clients) for s3fs, JindoFS, etc.

Pick based on **where you want the configuration to live**:

| Demo | Config lives in... | Change config by... | Use when... |
|------|--------------------|---------------------|-------------|
| [1-baked-in](1-baked-in/) | Image ENV | Rebuild image | Config is fixed and never changes |
| [2-oss-compatible](2-oss-compatible/) | PV volumeAttributes | Recreate PV, or add to `spec.mountOptions` | You want the same PV format as `ossplugin` driver — easy migration from existing OSS volumes |
| [3-standard](3-standard/) | PV mountOptions | Edit PV | Not tied to OSS field names — entrypoint defines its own env var schema, so it adapts to any FUSE client without reshaping the PV around `otherOpts` |
| [4-configmap](4-configmap/) | ConfigMap | Edit ConfigMap | Solidify complex format/mount configs (cache policy, block size, trash retention, etc.) into a reusable template — adding a new instance is just a PV with `source`/`bucket`/`url` pointing to the ConfigMap |
| [5-agent-identity](5-agent-identity/) | Sandbox identity, scripts in the image | Rebuild image | Self-built FUSE client on an **ACS sandbox** with no credential in the cluster: mount-proxy exchanges the sandbox token for a scoped STS credential, delivers it as files, and rotates it before expiry. Also **how to adapt a client to rotation** — JindoFS / JuiceFS / s3fs, one per behaviour a client can have |

Demo 5 is for the **ACS sandbox** scenario rather than the DaemonSet one used by
demos 1–4. It assumes the sandbox itself is already set up — see
[Agent Sandbox](https://help.aliyun.com/en/cs/user-guide/agent-sandbox) — and
covers only what the driver does with the credential:
[docs/customfuse-agent-identity.md](../../docs/customfuse-agent-identity.md).

> **Note**: `volumeAttributes` are immutable after PV creation — to change one you
> must delete and recreate the PV. `spec.mountOptions` is not: it can add what
> `volumeAttributes` left out and raise `capacity`, without recreating anything.
> See [Entrypoint Env Vars](#entrypoint-env-vars) for exactly what it can and
> cannot do.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│ Fuse Pod (created by CSI controller in ack-csi-customfuse ns)   │
│                                                                 │
│  PID 1: csi-mount-proxy-server                                  │
│    ↓ receives mount request from CSI node                       │
│    ↓ sets env vars from volumeAttributes, mountOptions, Secret  │
│    ↓ calls /entrypoint.sh (or /etc/fuse-config/entrypoint.sh)   │
│                                                                 │
│  entrypoint.sh:                                                 │
│    - reads env vars ($source, $mountpoint, $bucket, $url, ...)  │
│    - runs FUSE mount in foreground at $mountpoint               │
└─────────────────────────────────────────────────────────────────┘
```

The container ENTRYPOINT is `csi-mount-proxy-server` (NOT your entrypoint.sh).
mount-proxy is the process manager — it calls your entrypoint.sh when a mount
request arrives, with the `Secret` keys and the parameters the driver resolved
injected as env vars (same key name, no prefix). Which `volumeAttributes` keys
those are, and what happens to the rest, is in the reference below.

## Configuration Reference

### Entrypoint Env Vars

Your entrypoint receives env vars from **four sources**:

| Source | How it maps |
|--------|-------------|
| `volumeAttributes` | Only the keys listed below. Each becomes a driver field and then `$key`. A key the driver does not know is **dropped**: no env var, no error |
| `pv.spec.mountOptions` | Each entry → its own `$key`, spelled exactly as written. A bare flag `key` with no `=` becomes `$key=""` (detect with `${key+set}`) — but only for a name the driver does not know; a bare entry naming one is refused, since a field cannot be applied from a value that is not there, and the mount fails with `InvalidArgument` |
| Secret (`nodePublishSecretRef`) | Every key → `$key` |
| mount-proxy | `$mountpoint` always; `$source` when the volume has one; `$readOnly` only when the volume is read-only |

`volumeAttributes` is a whitelist and `mountOptions` is open. The asymmetry is
deliberate: `volumeAttributes` feeds the driver's own logic — the `bucket_name`
metrics label — so a key it does not understand cannot be forwarded safely.
`mountOptions` exists precisely to reach an entrypoint the driver knows nothing
about, so it forwards everything it does not consume.

**Names the driver owns.** `$mountpoint`, `$source` and `$readOnly` describe what the
driver does rather than what the volume said, so a volume parameter or a Secret key
cannot redefine them. `$mountpoint` is the path the driver tells your entrypoint to
mount on and then waits for a mount point to appear at; `$source` it resolves from the
fields below and emits itself; `$readOnly` comes from the PV's `accessModes` and the
publish request, and is refused in both directions — a read-only volume cannot be made
writable, nor a writable one read-only. A name the mount-proxy already has in its own
environment is not overridable either — the client runs inside that process, and its
environment carries the settings that process runs on. Such an entry is dropped, and
mount-proxy logs the dropped names together with the full list of names your
entrypoint will see, so a variable that does not turn up is one log line away from
being explained.

**Precedence.** Kubernetes rejects edits to `spec.csi.volumeAttributes` once a PV
exists, while `spec.mountOptions` stays editable and takes effect on the next
mount. So `mountOptions` can *add* what `volumeAttributes` left out — that is how
an existing volume gains a parameter without being recreated — but it cannot
*redefine* what `volumeAttributes` already set:

| `mountOptions` entry | `volumeAttributes` | Result |
|---|---|---|
| `bucket=b` | no `bucket` | `$bucket=b`, and the `bucket_name` metrics label sees it too |
| `bucket=b` | `bucket: a` | `$bucket=a` — the entry is dropped and the plugin logs a warning |
| `capacity=100Gi` | `capacity: 50Gi` | `$capacity=100Gi` — the one field that may be redefined, and only upward |
| `capacity=10Gi` | `capacity: 50Gi` | `$capacity=50Gi` — dropped with a warning, as Kubernetes itself refuses to shrink a claim |
| `readOnly=false` | accessModes read-only | `$readOnly=true` — dropped with a warning |
| `mountpoint=/x` | — | dropped with a warning — see [Names the driver owns](#entrypoint-env-vars) |
| `bucket` with no `=` | — | the mount fails with `InvalidArgument`: a bare entry can only be passed through for a name the driver does not know, and `bucket` is one it does. Same for the other fields above and for `capacity` |
| `verbose` with no `=` | — | `$verbose=""` — the bare-flag form works, for a name the driver does not know |
| `extraOpts=a,b` | — | `$extraOpts=a,b` — a name the driver does not know, passed through untouched |

`readOnly` is not settable from `mountOptions` in either direction. It comes from
the PV's `accessModes` and from the publish request, neither of which is a volume
parameter.

**Read-only is enforced by the driver, not left to your entrypoint.** `$readOnly`
tells the client what to do, and the driver additionally makes the bind into your
pod read-only, so a client that ignores the variable still cannot be written
through. One consequence to plan for: the volume is mounted once per node and
bound into every pod using it, so a pod asking for read-write after a read-only
one inherits that read-only mount. Give the two access modes separate volumes.
That bind is a property of the node deployment: where the node plugin is injected
as a sidecar there is no bind to make read-only, so honouring `$readOnly` is your
entrypoint's job — see
[Where the driver runs](../../docs/customfuse.md#where-the-driver-runs).

Nothing is dropped silently — every ignored entry is logged by the node plugin —
but the corollary is that `mountOptions` is not a way to *change* an existing
`volumeAttributes` value. In particular it cannot add one flag to an existing
`otherOpts`: `$otherOpts` is a single opaque string whose internal format only
your entrypoint knows, so the driver cannot merge into it. To add an option to a
volume that already has some, pick a name of your own (`extraOpts` above) and
have your entrypoint concatenate the two — that is what
[2-oss-compatible/entrypoint.sh](2-oss-compatible/entrypoint.sh) does.

Common volumeAttributes keys (all optional):

| Key | Description |
|-----|-------------|
| `source` | Mount source → `$source`, opaque to the driver and passed through exactly as written: a metadata engine URL, a `bucket:path` pair, anything your entrypoint understands. The driver composes nothing, so a client that wants a derived form builds it from `$bucket` and `$path` — [2-oss-compatible/entrypoint.sh](2-oss-compatible/entrypoint.sh) does |
| `bucket` | The storage location the client mounts, in whatever form it names one → `$bucket`. Also the `bucket_name` metrics label |
| `path` | Sub-path within the volume → `$path` |
| `url` | The service endpoint the client talks to → `$url` |
| `otherOpts` | Client options → `$otherOpts`, one whole value the driver never splits or merges into, so its internal format is whatever your entrypoint expects: comma separated, `-o` prefixed, anything. An OSS volume's `otherOpts` therefore carries over unchanged. Cannot be extended from `spec.mountOptions` — add a name of your own and concatenate it in your entrypoint, as in the [precedence](#entrypoint-env-vars) section |
| `capacity` | Volume quota passed as `$capacity` to the entrypoint. Plain integer or Kubernetes Quantity (e.g. `100`, `100Gi`), validated and passed through as-is; the entrypoint strips the suffix if its client needs a bare number: `capacity=${capacity%Gi}`. A dynamically provisioned volume gets it from the PVC's requested size. To raise a quota after the PV exists, put `capacity=<value>` in `spec.mountOptions` — the one field it may redefine, and only upward, because the driver implements no expansion and Kubernetes likewise refuses to shrink a claim |

Control fields (consumed by the driver rather than forwarded as client mount
options — except `authType` and `credentialDir`, which do reach the entrypoint as
env vars):

| Key | Description |
|-----|-------------|
| `fuseType` | FUSE client type for metrics and image resolution. Can also be set via `pv.spec.csi.fsType` (a PV spec field, not a volumeAttribute). Where both name a client they have to agree; `customfuse` on either side is exempt, being the generic marker rather than a client name. That exemption is what makes dynamic provisioning work: the controller stamps `fsType: customfuse` on every PV it creates, so a StorageClass parameter naming your client is not a conflict. |
| `entrypointConfig` | ConfigMap name (in `ack-csi-customfuse` ns) projected into the **fuse pod** to override `/entrypoint.sh` — per volume. On the sandbox path a PV selects nothing: the sidecar still runs `/etc/fuse-config/entrypoint.sh` ahead of the image's, but that directory is one copy shared by every sandbox |
| `entrypointKey` | Key in that ConfigMap (default: `entrypoint.sh`) |
| `dnsPolicy` | Fuse pod DNS policy: `ClusterFirst`, `ClusterFirstWithHostNet` or `Default`, matched case-insensitively. Anything else is logged and left unset. The fuse pod is on the host network, so unset means `ClusterFirst`, which there resolves through the *node* rather than the cluster: an in-cluster endpoint — a service name, say — will not resolve until you set `ClusterFirstWithHostNet`. Endpoints reached by IP or by public name are unaffected. |
| `serviceAccountName` | ServiceAccount the fuse pod runs as. Must exist in `ack-csi-customfuse`, not in the consumer's namespace. Defaults to that namespace's `default`. The fuse pod calls no API server, so this is for reaching a private registry — see [Private registry](#private-registry) |
| `authType` | `agent-identity` to exchange a sandbox token for a scoped, rotated STS credential delivered as files. Empty (default) passes Secret entries through as env vars. Also reaches the entrypoint as `$authType`, so one entrypoint can serve both flows. See [Demo 5](5-agent-identity/). |
| `sandboxCredProviderName` | Supplied by the sandbox side rather than declared on the PV; `credentialProviderName` is accepted as an alias, as in OSS. |
| `credentialDir` | With `authType: agent-identity`, pins where the credential files are written. Default: a per-mount directory. Either way the entrypoint gets `$credentialDir` resolved to the directory actually used, so it never has to reconstruct the default. |
| `credentialRefreshHookKey` | With `authType: agent-identity`, key in `entrypointConfig` holding a script run after each rotation, for clients that cannot reload the credential files themselves. Same fuse-pod projection as `entrypointConfig`, so on the sandbox path a PV cannot pick the hook: the sidecar runs `/etc/fuse-config/refresh-hook.sh` if anything is mounted there, otherwise `/refresh-hook.sh` from the image. |

These decide how the driver builds the fuse pod, not what the client mounts, so
`spec.mountOptions` cannot set them — an entry naming one is logged and ignored.
They come from `volumeAttributes` only, which means changing one on an existing
volume means recreating the PV.

**Secret keys.** Every key is passed through as `$<key>` — no prefix, no
transformation, and no second spelling derived from one. Which names a client accepts
is the client's business, so an adapter that takes more than one resolves them itself:
[2-oss-compatible/entrypoint.sh](2-oss-compatible/entrypoint.sh) maps the older
`akId`/`akSecret` onto `$accessKeyId`/`$accessKeySecret` in two lines, which is why a
Secret written for OSS works there unchanged.

### Dynamic provisioning

The demos here are all static PVs. A StorageClass for this driver carries the
volume's parameters, which arrive as `volumeAttributes` on the PV it creates:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: customfuse
provisioner: customfuseplugin.csi.alibabacloud.com
reclaimPolicy: Retain
parameters:
  fuseType: juicefs
  bucket: my-jfs-data
  url: oss-cn-hangzhou-internal.aliyuncs.com
```

`reclaimPolicy: Retain` is not a style choice. `CreateVolume` makes no storage and
`DeleteVolume` reclaims none, so under `Delete` the PV is removed and whatever the
entrypoint created — a formatted filesystem, its metadata engine — is left behind
with nothing pointing at it. The driver does reject a `Delete` it can see: it reads
the policy from a `csi.alibabacloud.com/reclaimPolicy` parameter, the same
convention the OSS and NAS drivers use. A provisioner that passes the
StorageClass's own field through supplies that parameter; the stock
external-provisioner does not, and then the check never runs. Set the policy on the
StorageClass rather than counting on the check.

Two things the controller adds to every PV it creates, both of which the static
demos spell out by hand: `capacity`, from the PVC's requested size, and
`fsType: customfuse`, from its `--default-fstype` flag. The second is why a
`fuseType` parameter naming your client is not a conflict — see the `fuseType` row
in [Control fields](#entrypoint-env-vars).

### Image Resolution (csi-plugin ConfigMap)

The controller watches `csi-plugin` ConfigMap in `kube-system` via informer.
Key format: `fuse-<fuseType>`, content: `image=<full-image-path>`.

Changes take effect on the next mount — no restart needed. Already-running fuse
pods retain their current image until the volume is remounted (a new
ControllerPublish is triggered).

```bash
# Create (if not exists):
kubectl -n kube-system create configmap csi-plugin \
  --from-literal="fuse-juicefs=image=registry.example.com/csi-fuse-juicefs:v1.0"
```

### Private registry

The driver puts no `imagePullSecrets` on the fuse pod, and the chart's
`imagePullSecrets` apply only to the workloads the chart itself renders — not to
pods created at mount time. So an image in a private registry needs a
ServiceAccount that carries the credentials, named from the volume:

```bash
# 1. A docker-registry Secret in the fuse pod namespace. Read the password from a
#    file or a prompt rather than a flag, so it does not land in shell history.
kubectl -n ack-csi-customfuse create secret docker-registry my-registry-cred \
  --docker-server=registry.example.com \
  --docker-username=<user> --docker-password=<password>

# 2. A ServiceAccount that uses it.
kubectl -n ack-csi-customfuse create serviceaccount my-fuse-sa
kubectl -n ack-csi-customfuse patch serviceaccount my-fuse-sa \
  -p '{"imagePullSecrets": [{"name": "my-registry-cred"}]}'
```

```yaml
# 3. Name it from the PV.
spec:
  csi:
    volumeAttributes:
      serviceAccountName: my-fuse-sa
```

Both objects belong in `ack-csi-customfuse`, because that is where the fuse pod is
created; Kubernetes does not look in the consumer's namespace. Patching that
namespace's `default` ServiceAccount works too, but applies to every customfuse
volume in the cluster.

## Other FUSE Clients

### How to adapt for your FUSE client

Start from **Demo 2** if your FUSE client follows OSS-like conventions
(`source`, `bucket`, `url`, `otherOpts`), or from **Demo 3** if you want to
define your own parameter names via `mountOptions`:

```bash
# OSS-like clients:
cp -r examples/customfuse/2-oss-compatible/ examples/customfuse/my-fuse/
# Or for custom parameter names:
# cp -r examples/customfuse/3-standard/ examples/customfuse/my-fuse/
```

1. **Dockerfile** — change `FROM` in Stage 2 to your FUSE client's base image
2. **entrypoint.sh** — replace the mount command with yours, using env vars
3. **pv.yaml** — change `fuseType` and volumeAttributes to match your client;
   create a Secret with your credentials and reference it via `nodePublishSecretRef`
4. **csi-plugin ConfigMap** — add `fuse-<your-fuseType>=image=...`

### s3fs example

Following the 4 steps above with `cp -r 2-oss-compatible/ my-s3fs/`:

**Step 1 — Dockerfile** (change Stage 2 base image):
```dockerfile
FROM ubuntu:22.04
RUN apt-get update && apt-get install -y s3fs && rm -rf /var/lib/apt/lists/*
```

**Step 2 — entrypoint.sh** (replace mount command):
```bash
#!/bin/bash
set -e

echo "$accessKeyId:$accessKeySecret" > /tmp/passwd
chmod 600 /tmp/passwd

# One comma-separated -o value, expanded quoted. s3fs parses its options through
# libfuse, which splits that value on commas, so the flags still arrive as
# separate options. Building a string of several "-o x -o y" flags instead would
# have to be expanded unquoted to work at all, and unquoted means word splitting
# plus every metacharacter the driver does not refuse.
S3FS_OPTS="url=http://$url,passwd_file=/tmp/passwd"
[ -n "$otherOpts" ] && S3FS_OPTS="$S3FS_OPTS,$otherOpts"

exec s3fs "$source" "$mountpoint" -f -o "$S3FS_OPTS"
```

**Step 3 — pv.yaml** (change `fuseType` and volumeAttributes):
```yaml
csi:
  driver: customfuseplugin.csi.alibabacloud.com
  volumeHandle: s3fs-pv
  volumeAttributes:
    fuseType: "s3fs"
    source: "my-bucket"
    url: "oss-cn-hangzhou-internal.aliyuncs.com"
    otherOpts: "parallel_count=5,multipart_size=10"
  nodePublishSecretRef:        # ← credentials from Secret
    name: s3fs-creds
    namespace: default
```

**Step 4 — csi-plugin ConfigMap** (register the image):
```bash
kubectl -n kube-system patch configmap csi-plugin --type merge \
  -p '{"data":{"fuse-s3fs":"image=registry.example.com/csi-fuse-s3fs:v1.0"}}'
```

Ref: [s3fs-fuse](https://github.com/s3fs-fuse/s3fs-fuse)

### JindoFS example

Same 4 steps with `cp -r 2-oss-compatible/ my-jindofs/`:

**Step 1 — Dockerfile** (install the JindoFuse client on your own base image):
```dockerfile
# JindoFuse ships through Alibaba Cloud's official JindoData distribution, not a
# package registry, so pin the version and base image your deployment supports
# rather than a tag this README cannot vouch for. See the JindoData download page
# (https://github.com/aliyun/alibabacloud-jindodata) for the current release.
FROM <your-base-image>
COPY jindofsx-<version>/bin/jindo-fuse /usr/local/bin/
```

**Step 2 — entrypoint.sh**:
```bash
#!/bin/bash
set -e

export JINDOSDK_ACCESS_KEY_ID="$accessKeyId"
export JINDOSDK_ACCESS_KEY_SECRET="$accessKeySecret"

# As above: one comma-separated -o value, expanded quoted, rather than a string of
# separate flags that only works unquoted.
JINDO_OPTS="uri=$source,endpoint=$url"
[ -n "$otherOpts" ] && JINDO_OPTS="$JINDO_OPTS,$otherOpts"

exec jindo-fuse -f -o "$JINDO_OPTS" "$mountpoint"
```

**Step 3 — pv.yaml**:
```yaml
csi:
  driver: customfuseplugin.csi.alibabacloud.com
  volumeHandle: jindofs-pv
  volumeAttributes:
    fuseType: "jindofs"
    source: "oss://my-bucket/path"
    url: "oss-cn-hangzhou-internal.aliyuncs.com"
    otherOpts: "attr_timeout=7,entry_timeout=7"
  nodePublishSecretRef:
    name: jindofs-creds
    namespace: default
```

**Step 4 — csi-plugin ConfigMap**:
```bash
kubectl -n kube-system patch configmap csi-plugin --type merge \
  -p '{"data":{"fuse-jindofs":"image=registry.example.com/csi-fuse-jindofs:v1.0"}}'
```

Ref: [JindoFS FUSE](https://github.com/aliyun/alibabacloud-jindodata)

## Important Notes

1. **Fuse pod namespace**: Fuse pods run in `ack-csi-customfuse` namespace, which is
   automatically created by the Helm chart when `csi.customfuse.enabled: true`.

2. **`entrypointConfig` ConfigMap** must also be in `ack-csi-customfuse` namespace
   (Kubernetes does not support cross-namespace ConfigMap volume mounts).

3. **Do NOT bake credentials into images** in production — use Kubernetes Secrets.

4. **The client must be the entrypoint's last `exec`, and must run in the
   foreground.** mount-proxy tracks one PID — the entrypoint's — and signals only
   that PID, so `exec` is what makes the client the process being tracked. Start
   the client and let the script exit, and the mount is reported as
   `entrypoint exited unexpectedly` whether or not it came up. Keep a shell in
   front of it, and teardown signals the shell alone, leaving the client holding a
   mount nothing tracks. Every demo here ends in an `exec`;
   [the full contract](../../docs/customfuse.md#entrypoint-process-contract)
   states both failure modes.

   Foreground is the other half, and how a client asks for that varies: a `-f`
   flag, an environment variable, or nothing because foreground is already the
   default. The JuiceFS CE examples here need `JFS_FOREGROUND=1`, because its
   `mount.juicefs` helper daemonizes without it — and that helper must be invoked
   by absolute path, since JuiceFS only recognises it as a mount helper when
   `argv[0]` ends in `/mount.juicefs`.

   A client that can only daemonize leaves no `exec` to make, so the script has to
   forward the signals itself and outlive the client it started:
   ```bash
   my-fuse-client "$source" "$mountpoint" &
   client=$!
   # mount-proxy signals this script, not the client. Without the trap the shell
   # dies alone and the client keeps the mount. The inner wait holds the shell
   # until the client is actually gone, which is the window mount-proxy gives.
   trap 'kill -TERM "$client" 2>/dev/null; wait "$client"' TERM INT
   wait "$client"
   ```

5. **Validate user-supplied options to prevent command injection**. Parameters
   from `volumeAttributes.otherOpts` or `pv.spec.mountOptions` are
   user-configurable. Your entrypoint should validate them before use.

   **Best practices**:
   - Reject values containing shell metacharacters (`;`, `|`, `&`, `` ` ``,
     `$`, `(`, `)`, newlines).
   - If options are `--flag` style, validate each argument starts with `--`.
   - Always quote variables when passing to commands (`"$opts"`, not `$opts`).

   Example validation snippet:
   ```bash
   # The character set lives in a variable: bash cannot parse `$'\n'`
   # inside a [[ =~ ]] class.
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

   # For comma-separated $otherOpts:
   IFS=',' read -ra OPTS <<< "$otherOpts"
   validate_opts "${OPTS[@]}"
   ```

6. **The fuse pod mounts no cache volume.** Its only volumes are the mount
   target's parent directory, the metrics directory, and — when you set
   `entrypointConfig` — that ConfigMap. Anything else your client writes, its
   cache directory included, lands in the container's writable layer: it is
   destroyed when the pod is replaced, which [Operations](#operations) describes
   as a routine action, and it counts against the container's ephemeral storage,
   bounded by whatever cache size limit your client honours.

   For a read cache that only costs re-reads. What it must not carry is a cache
   mode that acknowledges a write before the data reaches the backend — JuiceFS's
   `writeback` is one such mode, and its own CSI driver advises against it under
   pod lifecycles for exactly this reason. Those writes would be lost silently,
   with no error anywhere. The demos here therefore configure a read cache only.
