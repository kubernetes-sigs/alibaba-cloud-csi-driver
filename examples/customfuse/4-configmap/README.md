# Demo 4: Entrypoint via ConfigMap

## Use Case

Instead of baking the entrypoint script into a custom Docker image, store it in
a ConfigMap. This lets you update mount logic (e.g. add flags, change auth flow)
without rebuilding the container image — just edit the ConfigMap.

## How It Works

1. Create a ConfigMap in `ack-csi-customfuse` namespace with key `entrypoint.sh`
2. Set `entrypointConfig` in PV volumeAttributes to the ConfigMap name
3. FusePodManager mounts the ConfigMap at `/etc/fuse-config/` (read-only)
4. mount-proxy checks `/etc/fuse-config/entrypoint.sh` on each mount request;
   if it exists, runs that instead of the baked-in `/entrypoint.sh`

No custom Docker image needed — use the base image directly. The fuse pod image
only needs the FUSE client binary and mount-proxy; the startup logic lives in
the ConfigMap.

## Important

**The ConfigMap must be in `ack-csi-customfuse` namespace.** Kubernetes does not
support cross-namespace ConfigMap volume mounts.

## Setup

```bash
# 1. Create the ConfigMap with your entrypoint script
kubectl apply -f configmap.yaml

# 2. Create the Secret
kubectl apply -f secret.yaml

# 3. Create PV, PVC, and consumer
kubectl apply -f pv.yaml -f pvc.yaml -f deploy.yaml
```

## ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: juicefs-entrypoint
  namespace: ack-csi-customfuse
data:
  entrypoint.sh: |
    #!/bin/bash
    set -e
    juicefs format --storage=oss \
        --bucket="http://$bucket.$url" \
        --access-key="$accessKeyId" --secret-key="$accessKeySecret" \
        "$source" myjfs 2>/dev/null || true
    exec juicefs mount -f $otherOpts "$source" "$mountpoint"
```
