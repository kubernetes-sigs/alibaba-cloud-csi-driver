# Demo 4: Entrypoint via ConfigMap

## Overview

Complex format and mount configurations are **hardcoded into ConfigMaps**. Each
PV only provides instance-specific values (`source`, `bucket`, `url`) and
points to a ConfigMap via `entrypointConfig`. This separates **what to run**
(ConfigMap) from **where to run it** (PV).

Adding a new JuiceFS instance = copy a PV, change `source`/`bucket`/`url`.
No image rebuild, no entrypoint editing.

```
ConfigMap "juicefs-production"     ConfigMap "juicefs-development"
┌──────────────────────────┐      ┌──────────────────────────┐
│ format: block-size=4096  │      │ format: block-size=4096  │
│         trash-days=0     │      │         trash-days=2     │
│ mount:  cache-size=4096  │      │ mount:  cache-size=256   │
│         writeback        │      │                          │
│         prefetch=1       │      │                          │
└──────────┬───────────────┘      └──────────┬───────────────┘
           │                                  │
    ┌──────┴──────┐                    ┌──────┴──────┐
    │ PV prod-a   │                    │ PV dev-a    │
    │ source=...  │                    │ source=...  │
    │ bucket=...  │                    │ bucket=...  │
    └─────────────┘                    └─────────────┘
```

## How It Works

1. Create ConfigMaps in `ack-csi-customfuse` namespace with key `entrypoint.sh`
2. Set `entrypointConfig` in PV volumeAttributes to the ConfigMap name
3. FusePodManager mounts the ConfigMap at `/etc/fuse-config/` (read-only)
4. mount-proxy checks `/etc/fuse-config/entrypoint.sh` on each mount request;
   if it exists, runs that instead of the baked-in `/entrypoint.sh`

All parameter passing methods from previous demos are fully compatible
(volumeAttributes, mountOptions, Secret — see
[Configuration Reference](../README.md#configuration-reference)). This demo
puts `source`, `bucket`, `url` in PV volumeAttributes so that the ConfigMap
is not tied to a specific instance and can be reused across different
filesystems.

**The ConfigMap must be in `ack-csi-customfuse` namespace.** Kubernetes does not
support cross-namespace ConfigMap volume mounts.

## Setup

```bash
# 1. Create ConfigMaps (one per environment)
kubectl apply -f configmap-production.yaml
kubectl apply -f configmap-development.yaml

# 2. Create the Secret (credentials, shared across instances)
kubectl apply -f secret.yaml

# 3. Create PVs and PVCs
kubectl apply -f pv.yaml -f pvc.yaml
```

## Iterating without image rebuild

1. `kubectl edit configmap juicefs-production -n ack-csi-customfuse`
2. Ensure no consumer pods on the node are using this PV (the fuse pod is
   reused as long as any consumer on the same node references the PV).
   Delete all consumer pods (or the PVC) so that ControllerUnpublish cleans
   up the fuse pod.
3. Re-create the consumer pod — ControllerPublish creates a new fuse pod
   that picks up the updated ConfigMap.
4. Find and verify the new fuse pod:
   ```bash
   kubectl -n ack-csi-customfuse get pods \
       -l csi.alibabacloud.com/volume-id=juicefs-prod-pv
   kubectl -n ack-csi-customfuse logs \
       -l csi.alibabacloud.com/volume-id=juicefs-prod-pv \
       -c customfuse
   ```

No `docker build`, no `docker push`, no PV edit.
