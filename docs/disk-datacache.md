# Accelerating a Cloud Disk with a Local Disk (Data Cache)

The Disk CSI driver can place a fast **local disk** in front of a slower
**cloud disk** to accelerate I/O. This is implemented with the Linux
device-mapper `dm-cache` target: the cloud disk is the persistent *origin*, and
a file on the node's local storage acts as the cache. Pod I/O is served from the
local cache whenever possible, while data ultimately persists on the cloud disk.

This feature is a capability of the standard Cloud Disk CSI driver
(`diskplugin.csi.alibabacloud.com`); it is **not** a separate plugin.

## When to Use It

Use the data cache when your workload:

- is read-heavy or has a hot working set that fits in the local disk, and
- runs on nodes that have fast local storage (e.g. local NVMe SSD), and
- can tolerate the cache being local to the node (a cache-backed volume must be
  re-cached after it moves to another node).

For latency-insensitive or write-once workloads the cache adds little benefit.

## How It Works

1. **Provisioning** — The `StorageClass` carries `dataCacheSize` (and optionally
   `dataCacheMode`). These values are validated at `CreateVolume` and preserved
   in the PV's `VolumeContext` so they reach the node.
2. **Node stage/publish** — When the volume is staged on a node, the driver:
   - `fallocate`s two backing files on the node's local storage under
     `/var/alibaba-cloud-csi/data-cache/`: `<volumeID>.data` (the cache, sized
     by `dataCacheSize`) and `<volumeID>.meta` (dm-cache metadata, sized
     automatically);
   - attaches each file to a loop device;
   - creates a device-mapper `cache` device named `csi-datacache-<volumeID>`,
     exposed at `/dev/mapper/csi-datacache-<volumeID>`, with the cloud disk as
     the origin;
   - formats/mounts **the dm-cache device** instead of the raw cloud disk, so
     all pod I/O flows through the local cache.
3. **Persistence & crash safety** — Backing files are never truncated or deleted
   while they may hold un-flushed writeback data. After a reboot the dm-cache
   superblock is re-opened (not reformatted), so dirty blocks are written back to
   the cloud disk on the next teardown.
4. **Expansion** — On `NodeExpandVolume` the dm-cache target is grown to cover
   the enlarged cloud disk before the filesystem is resized.
5. **Teardown** — When the volume is unstaged, the driver switches the cache to
   the `cleaner` policy to flush all dirty blocks back to the cloud disk, then
   removes the dm-cache device and deletes the backing files.

## Prerequisites

- The **node** must be Linux with device-mapper (`/dev/mapper/control`), loop
  devices, and the `dm-cache` kernel target available.
- The node plugin must host-mount `/var/alibaba-cloud-csi` (this is configured
  by the Helm chart by default). The cache `.data`/`.meta` files live here.
- The node must have enough free local disk space for `dataCacheSize` per
  cache-backed volume scheduled to it.

If device-mapper is unavailable on a node, or the cache file cannot be created,
the driver **falls back** to using the raw cloud disk and emits a
`DataCacheFallback` warning event on the pod. The volume still works, just
without acceleration.

## StorageClass Parameters

| Parameter       | Required | Values                        | Default        | Description |
|-----------------|----------|-------------------------------|----------------|-------------|
| `dataCacheSize` | Yes (to enable) | A Kubernetes quantity, e.g. `10Gi` | — | Size of the local cache. A non-zero value **enables** the data cache. |
| `dataCacheMode` | No       | `writethrough`, `writeback`   | `writethrough` | Cache write mode. See below. |

Validation rules:

- `dataCacheSize` must parse as a quantity; an unrecognized `dataCacheMode` is
  rejected.
- Setting `dataCacheMode` without a non-zero `dataCacheSize` is an error.
- If `dataCacheSize` is set but `dataCacheMode` is omitted, the mode defaults to
  `writethrough`.

### Write modes

- **`writethrough`** (default, recommended) — writes go to both the cache and
  the cloud disk before being acknowledged. Accelerates reads; safe against node
  failure because the cloud disk is always up to date.
- **`writeback`** — writes are acknowledged once they land in the local cache
  and are flushed to the cloud disk later. Accelerates writes as well, but dirty
  data lives only on the node's local disk until flushed. **If the node's local
  disk is lost before a flush, that data is lost.** Use only when the local disk
  is durable enough for your needs and the performance gain is required.

## Usage

### 1. Create a StorageClass

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-disk-datacache
provisioner: diskplugin.csi.alibabacloud.com
parameters:
  type: cloud_essd
  dataCacheSize: "10Gi"
  dataCacheMode: "writethrough"   # or "writeback"
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
```

```shell
kubectl apply -f storageclass.yaml
```

### 2. Create a PVC

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: datacache-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: alicloud-disk-datacache
  resources:
    requests:
      storage: 100Gi
```

```shell
kubectl apply -f pvc.yaml
```

### 3. Mount it in a workload

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: datacache-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: datacache-demo
  template:
    metadata:
      labels:
        app: datacache-demo
    spec:
      containers:
        - name: app
          image: nginx:1.25
          volumeMounts:
            - name: data
              mountPath: /data
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: datacache-pvc
```

```shell
kubectl apply -f deploy.yaml
```

## Verifying the Cache Is Active

On the node where the pod runs:

```shell
# The dm-cache device should exist for the volume
ls -l /dev/mapper/csi-datacache-<volumeID>

# Inspect the cache table and runtime status
dmsetup status csi-datacache-<volumeID>

# The backing files
ls -lh /var/alibaba-cloud-csi/data-cache/
```

`dmsetup status` reports the cache block usage, hit/miss counters, the number of
dirty blocks, the io mode, and the active policy (`mq` during normal operation).

If the cache is **not** active, check for a `DataCacheFallback` warning event:

```shell
kubectl describe pod <pod-name>
```

and confirm that the node meets the [Prerequisites](#prerequisites).

## Expansion

Volume expansion is supported. Ensure the `StorageClass` has
`allowVolumeExpansion: true`, then edit the PVC's requested size. The driver
grows the cloud disk, extends the dm-cache target to match, and resizes the
filesystem.

## Notes and Limitations

- The cache is **node-local**. When a pod (and its `ReadWriteOnce` volume) moves
  to a new node, a fresh cache is built there; the previous node's cache files
  are cleaned up on unstage.
- With `writeback`, un-flushed writes reside only on the node's local disk until
  drained. Prefer `writethrough` unless you specifically need write
  acceleration and accept the durability trade-off.
- Enough free local disk space must be available on the node for each
  cache-backed volume.
- Requires a Linux node with the `dm-cache` kernel target; otherwise the driver
  transparently falls back to the raw cloud disk.

## See Also

- [Cloud Disk](./disk.md)
- [Disk — Volume Expansion](./disk-resizer.md)
- [Installation](./install.md)
