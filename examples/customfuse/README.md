# CustomFuse CSI Driver

The CustomFuse driver runs **any FUSE client** as a managed pod. You supply the
image + entrypoint; the driver handles pod lifecycle, mount propagation, and
credential injection.

## How It Works

```
┌─────────────────────────────────────────────────────────────────┐
│ Fuse Pod (created by CSI controller in ack-csi-customfuse ns)   │
│                                                                 │
│  PID 1: csi-mount-proxy-server                                  │
│    ↓ receives mount request from CSI node                       │
│    ↓ injects env vars from PV volumeAttributes + Secret         │
│    ↓ calls /entrypoint.sh (or /etc/fuse-config/entrypoint.sh)   │
│                                                                 │
│  entrypoint.sh:                                                 │
│    - reads env vars ($source, $mountpoint, $bucket, $url, ...)  │
│    - runs FUSE mount in foreground at $mountpoint               │
└─────────────────────────────────────────────────────────────────┘
```

**Key concepts:**

- The container ENTRYPOINT is `csi-mount-proxy-server`, NOT your entrypoint.sh.
  mount-proxy is the process manager; it calls your entrypoint.sh when a mount
  request arrives.
- `$mountpoint` is injected by mount-proxy — you do NOT set it in PV. Your
  entrypoint just mounts the filesystem there.
- All `volumeAttributes` keys and all `Secret` keys become env vars in
  entrypoint.sh (same key name, no prefix).

### What goes into the fuse pod image

Your image must contain:

1. **csi-mount-proxy-server** + **csi-mount-proxy-client** — built from this repo
   (see Stage 1 in any demo Dockerfile)
2. **Your FUSE client binary** — the base image in Stage 2 (e.g., `juicedata/mount`,
   `jindofs/smartdata`, or any image with your FUSE client installed)
3. **`/entrypoint.sh`** — your mount script that reads env vars and runs the FUSE
   command in foreground

## Configuration Reference

### PV volumeAttributes → Entrypoint Env Vars

| Key | Env Var | Description |
|-----|---------|-------------|
| `fuseType` | — | Selects fuse pod image from ConfigMap key `fuse-<fuseType>` |
| `source` | `$source` | Mount source (META-URL, bucket:path, or any string). Falls back to `bucket:path` if empty |
| `bucket` | `$bucket` | Object storage bucket name |
| `path` | — | Sub-path; only used for `source` fallback (`bucket:path`) |
| `url` | `$url` | Storage endpoint |
| `otherOpts` | `$otherOpts` | Opaque options string passed through to entrypoint |
| `entrypointConfig` | — | ConfigMap name (in `ack-csi-customfuse` ns) to override `/entrypoint.sh` |
| `entrypointKey` | — | Key in ConfigMap (default: `entrypoint.sh`) |
| `dnsPolicy` | — | Pod DNS policy (`ClusterFirst`, `Default`, etc.) |

Additionally, mount-proxy **always** injects these env vars (not from PV):

| Env Var | Description |
|---------|-------------|
| `$mountpoint` | The target path where your FUSE client must mount. Managed by CSI — do NOT hardcode. |

### Secret → Entrypoint Env Vars

All Secret keys are passed through as env vars with the **same key name**.

| Key | Env Var | Note |
|-----|---------|------|
| `accessKeyId` | `$accessKeyId` | |
| `accessKeySecret` | `$accessKeySecret` | |
| `akId` | `$akId` + `$accessKeyId` | Legacy; auto-mapped if `accessKeyId` absent |
| `akSecret` | `$akSecret` + `$accessKeySecret` | Legacy; auto-mapped if `accessKeySecret` absent |
| *(any key)* | `$<key>` | Arbitrary keys are passed through |

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

## Quick Start (Demo 2 — Standard)

### Step 1: Build the fuse pod image

All Dockerfiles share the same structure:

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
docker build -f examples/customfuse/2-standard/Dockerfile \
    -t registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1 .
docker push registry.cn-hangzhou.cr.aliyuncs.com/my-repo/csi-fuse-juicefs:v0.1
```

### Step 2: Configure the fuse pod image in csi-plugin ConfigMap

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

### Step 3: Create Secret, PV, PVC

```bash
# Edit secret.yaml with your real credentials first!
kubectl apply -f examples/customfuse/2-standard/secret.yaml
kubectl apply -f examples/customfuse/2-standard/pv.yaml
kubectl apply -f examples/customfuse/2-standard/pvc.yaml
```

### Step 4: Create a consumer Pod

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

## Demos (JuiceFS CE + Alibaba Cloud OSS)

The demos under this directory use [JuiceFS Community Edition](https://github.com/juicedata/juicefs)
as a concrete example. See individual demo directories for full YAML and Dockerfiles.

| Demo | Description | Image Rebuild? |
|------|-------------|----------------|
| [1-baked-in](1-baked-in/) | All config hardcoded in image ENV | Yes |
| [2-standard](2-standard/) | Config from PV volumeAttributes + Secret (recommended) | No |
| [3-advanced-entrypoint](3-advanced-entrypoint/) | Custom resource controls: cache quota, credential cleanup | Yes (if logic changes) |
| [4-configmap](4-configmap/) | Entrypoint via ConfigMap — iterate without image rebuild | No |

References:
- [JuiceFS — Object Storage Setup (Alibaba Cloud OSS)](https://juicefs.com/docs/community/reference/how_to_set_up_object_storage/#alibaba-cloud-oss)
- [JuiceFS — Command Reference](https://juicefs.com/docs/community/command_reference)

## Other FUSE Clients — Quick Reference

To adapt the driver for a different FUSE client, you need to change 3 things:

1. **Dockerfile Stage 2 `FROM`** → your FUSE client's base image
2. **entrypoint.sh** → translate env vars into your client's CLI arguments
3. **PV volumeAttributes + Secret** → fill in values your entrypoint expects

Below are minimal examples. All follow the **standard** pattern (Demo 2).

### s3fs (S3-compatible / OSS)

```
# PV volumeAttributes:
fuseType: s3fs
source: my-bucket
url: oss-cn-hangzhou-internal.aliyuncs.com
otherOpts: -o parallel_count=5 -o multipart_size=10

# Secret:
accessKeyId: <ak>
accessKeySecret: <sk>

# Entrypoint mount command:
echo "$accessKeyId:$accessKeySecret" > /tmp/passwd && chmod 600 /tmp/passwd
s3fs "$source" "$mountpoint" -f -o url="http://$url" -o passwd_file=/tmp/passwd $otherOpts
```

Ref: [s3fs-fuse](https://github.com/s3fs-fuse/s3fs-fuse)

### JindoFS (Alibaba Cloud)

```
# PV volumeAttributes:
fuseType: jindofs
source: oss://my-bucket/path
url: oss-cn-hangzhou-internal.aliyuncs.com
otherOpts: -oattr_timeout=7 -oentry_timeout=7

# Secret:
accessKeyId: <ak>
accessKeySecret: <sk>

# Entrypoint mount command:
export JINDOSDK_ACCESS_KEY_ID="$accessKeyId"
export JINDOSDK_ACCESS_KEY_SECRET="$accessKeySecret"
exec jindo-fuse -f -ouri="$source" -oendpoint="$url" $otherOpts "$mountpoint"
```

Ref: [JindoFS FUSE](https://github.com/aliyun/alibabacloud-jindodata)

## Important Notes

1. **Fuse pod namespace**: Fuse pods run in `ack-csi-customfuse` namespace, which is
   automatically created by the Helm chart when `csi.customfuse.enabled: true`.

2. **`entrypointConfig` ConfigMap** must also be in `ack-csi-customfuse` namespace
   (Kubernetes does not support cross-namespace ConfigMap volume mounts).

3. **Do NOT bake credentials into images** in production — use Kubernetes Secrets.

4. **FUSE mount command must run in foreground** (e.g., `juicefs mount -f`,
   `s3fs -f`). FUSE is a userspace filesystem — the kernel forwards all VFS
   operations to the FUSE process. If the process exits, the mount becomes stale
   (`Transport endpoint is not connected`). Foreground mode keeps the process as a
   child of mount-proxy, which tracks its lifecycle and detects crashes.

   If your FUSE client does not support foreground mode, add a `wait` or blocking
   loop at the end of entrypoint.sh to prevent it from exiting:
   ```bash
   my-fuse-client $source $mountpoint &
   FUSE_PID=$!
   wait $FUSE_PID
   ```
   Note that daemonized FUSE clients have additional drawbacks:
   - Logs go to internal files instead of stdout — invisible via `kubectl logs`.
     If the client supports configuring a log output path, set it to `/dev/stdout`.
   - mount-proxy cannot distinguish between "mount succeeded and is running" vs
     "mount failed silently" since the entrypoint returns immediately in both cases.

   **Recommendation**: Always prefer `-f` (foreground) when available. Configure
   the log destination to stdout/stderr so that `kubectl logs` works for
   debugging.

5. **Validate `otherOpts` to prevent command injection**. `otherOpts` comes from
   PV `volumeAttributes` — a user-configurable field that is passed directly into
   the mount command in entrypoint.sh. Without validation, a malicious value like
   `--cache-size=100; rm -rf /` could be injected.

   The driver cannot perform a generic check because valid option formats vary
   across FUSE clients (JuiceFS: `--key=value`, s3fs: `-o key=value`, etc.).
   **You must implement input validation in your entrypoint.sh** that matches your
   target client's flag format.

   The demos use a simple whitelist check — every token must start with `--`:
   ```bash
   for arg in $otherOpts; do
       if [[ "$arg" != --* ]]; then
           echo "ERROR: rejected invalid option in otherOpts: $arg" >&2
           exit 1
       fi
   done
   ```

   For production use, consider stricter validation: maintain an explicit allowlist
   of permitted flags, validate value ranges (e.g., `--cache-size` must be numeric),
   and reject flags that could alter security-sensitive behavior.
