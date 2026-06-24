# CustomFuse CSI Driver

The CustomFuse driver runs **any FUSE client** as a managed pod. You supply the
image + entrypoint; the driver handles pod lifecycle, mount propagation, and
credential injection.

> **Note**: CustomFuse currently supports static PVs only (manually created).
> Dynamic provisioning via StorageClass + PVC is not supported.

## Quick Start

### Step 1: Write your entrypoint.sh

Take the mount command you normally run on a node, put it in a script, and
replace hardcoded values with env vars. The driver calls this script with all
PV parameters and credentials injected as env vars. The script must mount at
`$mountpoint` and **stay in the foreground**:

```bash
#!/bin/bash
set -e

# All PV volumeAttributes, Secret keys, and mountOptions entries
# are available as env vars. $mountpoint is always injected.

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

```dockerfile
# Stage 1: build mount-proxy (same for ALL FUSE clients, copy as-is)
FROM golang AS builder
RUN go build -o /out/csi-mount-proxy-server ./cmd/mount-proxy-server && \
    go build -o /out/csi-mount-proxy-client ./cmd/mount-proxy-client

# Stage 2: YOUR fuse client image + mount-proxy + entrypoint
FROM <your-fuse-client-base-image>          # ← change this
COPY --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --from=builder /out/csi-mount-proxy-client /usr/local/bin/
COPY entrypoint.sh /entrypoint.sh           # ← your mount script
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
```

Build from the **repo root** (mount-proxy source code needed as context):

```bash
docker build -f examples/customfuse/2-oss-compatible/Dockerfile \
    -t registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1 .
docker push registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1
```

### Step 3: Configure the fuse pod image in csi-plugin ConfigMap

If the ConfigMap does not exist yet, create it:

```bash
kubectl -n kube-system create configmap csi-plugin \
  --from-literal="fuse-juicefs=image=registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1"
```

Or if it already exists, patch it:

```yaml
# kubectl -n kube-system edit configmap csi-plugin
data:
  fuse-juicefs: |
    image=registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1
```

### Step 4: Create Secret, PV, PVC

```bash
# Edit secret.yaml with your real credentials first!
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
      claimName: juicefs-standard-pvc
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
| [2-oss-compatible](2-oss-compatible/) | PV volumeAttributes | Recreate PV | You want the same PV format as `ossplugin` driver — easy migration from existing OSS volumes |
| [3-standard](3-standard/) | PV mountOptions | Edit PV | Not tied to OSS field names — entrypoint defines its own env var schema, so it adapts to any FUSE client without reshaping the PV around `otherOpts` |
| [4-configmap](4-configmap/) | ConfigMap | Edit ConfigMap | Solidify complex format/mount configs (cache policy, writeback, etc.) into a reusable template — adding a new instance is just a PV with `source`/`bucket`/`url` pointing to the ConfigMap |

> **Note**: volumeAttributes are immutable after PV creation. To change them you
> must delete and recreate the PV. <!-- TODO: support CSI ModifyVolume to allow
> in-place volumeAttributes updates -->

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
request arrives, with all `volumeAttributes` keys, `Secret` keys, and
`mountOptions` entries injected as env vars (same key name, no prefix).

## Configuration Reference

### Entrypoint Env Vars

Your entrypoint receives env vars from **four sources** (merged, last-wins on
duplicate keys):

| Source | How it maps |
|--------|-------------|
| `volumeAttributes` | Each key → `$key` (except control fields below) |
| `pv.spec.mountOptions` | `key=value` → `$key="value"`;  bare flag `key` → `$key=""` (detect with `${key+set}`) |
| Secret (`nodePublishSecretRef`) | Each key → `$key` |
| mount-proxy (always injected) | `$mountpoint`, `$readOnly` |

Common volumeAttributes keys (all optional):

| Key | Description |
|-----|-------------|
| `source` | Mount source (META-URL, bucket:path, etc.). Falls back to `bucket:path` if empty |
| `bucket` | Object storage bucket name |
| `path` | Sub-path; also used for `source` fallback |
| `url` | Storage endpoint |
| `otherOpts` | Comma-separated options string → `$otherOpts` |
| `capacity` | Volume quota passed as `$capacity` to the entrypoint. Plain integer or Kubernetes Quantity (e.g. `100`, `100Gi`). Values with units are validated; passed through as-is. Entrypoint strips suffix if needed: `capacity=${capacity%Gi}`. With `CustomFuseAutoCapacity` feature gate enabled, auto-filled from `pv.spec.capacity.storage` (with units). **Note:** auto-fill uses volumeHandle as PV name for the API lookup — if they differ, `$capacity` will not be set (a warning is logged). Set this field explicitly for guaranteed behavior. |

Control fields (consumed by the driver, NOT passed as env vars):

| Key | Description |
|-----|-------------|
| `fuseType` | FUSE client type for metrics and image resolution. Can also be set via `pv.spec.csi.fsType` (a PV spec field, not a volumeAttribute); if both are set, they must match. |
| `entrypointConfig` | ConfigMap name (in `ack-csi-customfuse` ns) to override `/entrypoint.sh` |
| `entrypointKey` | Key in ConfigMap (default: `entrypoint.sh`) |
| `dnsPolicy` | Pod DNS policy (`ClusterFirst`, `Default`, etc.) |

Secret key mapping:

| Key | Behavior |
|-----|----------|
| `accessKeyId` / `accessKeySecret` | Passed as-is |
| `akId` / `akSecret` | Legacy; auto-mapped to `$accessKeyId` / `$accessKeySecret` if absent |
| *(any key)* | Passed through as `$<key>` |

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

S3FS_OPTS="-o url=http://$url -o passwd_file=/tmp/passwd"
[ -n "$otherOpts" ] && S3FS_OPTS="$S3FS_OPTS -o $otherOpts"

exec s3fs "$source" "$mountpoint" -f $S3FS_OPTS
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

**Step 1 — Dockerfile**:
```dockerfile
FROM registry.cn-shanghai.aliyuncs.com/jindofs/smartdata:6.4.0
```

**Step 2 — entrypoint.sh**:
```bash
#!/bin/bash
set -e

export JINDOSDK_ACCESS_KEY_ID="$accessKeyId"
export JINDOSDK_ACCESS_KEY_SECRET="$accessKeySecret"

JINDO_OPTS="-ouri=$source -oendpoint=$url"
[ -n "$otherOpts" ] && JINDO_OPTS="$JINDO_OPTS -o $otherOpts"

exec jindo-fuse -f $JINDO_OPTS "$mountpoint"
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

4. **FUSE mount command must run in foreground** (e.g., `-f` flag). mount-proxy
   tracks the entrypoint process — if it exits, the mount becomes stale
   (`Transport endpoint is not connected`). If your client doesn't support
   foreground mode, background it and `wait`:
   ```bash
   my-fuse-client "$source" "$mountpoint" &
   wait $!
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
   validate_opts() {
       for opt in "$@"; do
           if [[ "$opt" =~ [;\|\&\`\$\(\)$'\n'] ]]; then
               echo "ERROR: invalid character in option: $opt" >&2
               exit 1
           fi
       done
   }

   # For comma-separated $otherOpts:
   IFS=',' read -ra OPTS <<< "$otherOpts"
   validate_opts "${OPTS[@]}"
   ```
