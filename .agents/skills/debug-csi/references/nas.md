# NAS / cnfs-nas-daemon Debugging

NAS EFC/CPFS mounts are served by cnfs-nas-daemon (a DaemonSet in `cnfs-system` namespace).
The CSI plugin sends mount requests over a Unix socket to the daemon, which runs `mount -t alinas`.

## Quick setup

```bash
# Set up mount proxy with alinas image (default)
scripts/setup-mount-proxy-debug.sh [image-tag]

# Set up with the all-in-one image
scripts/setup-mount-proxy-debug.sh [image-tag] aio
```

This enables the `AlinasMountProxy` feature gate on csi-plugin, patches cnfs-nas-daemon with the
test image, and restarts the pods.

## Test PV

PV volumeAttributes must trigger the `alinas` fstype. Use `useClient: efcclient`:
```yaml
csi:
  driver: nasplugin.csi.alibabacloud.com
  volumeHandle: my-nas-pv
  volumeAttributes:
    server: "<fs-id>-<suffix>.cn-<region>.nas.aliyuncs.com"
    path: "/"
    useClient: "efcclient"
```

## Architecture

```
csi-plugin (NodePublishVolume)
  → Unix socket /run/cnfs/alinas-mounter.sock
    → cnfs-nas-daemon pod (mount-proxy-server --driver=alinas)
      → mount -t alinas (EFC client, with auto_fallback_nfs to NFS)
```

Code path: `useClient: efcclient` → `ClientType=efcclient` → `mountFstype=alinas` → ProxyMounter.

## Finding NAS filesystems

Try `kubectl get pv test-nas-proxy` first to find existing test resources that you can reuse.
Also prefer creating PV with this name to reuse it next time.
If not found, try:
```bash
aliyun nas DescribeFileSystems --RegionId cn-beijing | jq '.FileSystems.FileSystem[] | {FileSystemId, FileSystemType, Status}'
aliyun nas DescribeMountTargets --RegionId cn-beijing --FileSystemId <fs-id> | jq '.MountTargets.MountTarget[]'
```

## Logs

```bash
kubectl -n cnfs-system logs -l app=cnfs-nas-daemon --tail=50
kubectl -n kube-system logs -l app=csi-plugin -c csi-plugin --tail=50 | grep -i "proxy\|alinas\|mount"
```

## Key source files

- `pkg/nas/nodeserver.go` — NodePublishVolume, volumeAttribute parsing, ClientType/MountProtocol determination
- `pkg/nas/utils.go:doMount()` — fstype routing, subpath auto-creation
- `pkg/nas/mounter.go` — NasMounter construction (ProxyMounter vs ConnectorMounter vs AdaptorMounter)
- `pkg/mounter/proxy_mounter.go` — ProxyMounter (sends requests over Unix socket)
- `pkg/mounter/proxy/server/alinas/driver.go` — server-side alinas mount handler
- `pkg/features/features.go` — AlinasMountProxy feature gate definition
