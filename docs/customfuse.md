# CustomFuse CSI Driver

The CustomFuse driver provides a generic framework for running **any FUSE client** as a managed Kubernetes pod. You bring your own FUSE client image and entrypoint; the driver handles pod lifecycle, mount propagation, and credential injection.

For a complete working example, see the [JuiceFS CE demo](../examples/customfuse/).

## Installation

CustomFuse can be installed independently or alongside other CSI drivers (OSS, NAS, etc.).
The controller and node plugin deploy in `kube-system` (same as other drivers);
only the managed **fuse pods** run in the dedicated `ack-csi-customfuse` namespace.

### Install CustomFuse only

```shell
helm upgrade --install csi-customfuse ./deploy/charts/alibaba-cloud-csi-driver \
  --namespace kube-system \
  --set csi.customfuse.enabled=true \
  --set csi.disk.enabled=false \
  --set csi.nas.enabled=false \
  --set csi.oss.enabled=false \
  --set csi.bmcpfs.enabled=false
```

### Install alongside other drivers

```shell
helm upgrade --install alibaba-cloud-csi-driver ./deploy/charts/alibaba-cloud-csi-driver \
  --namespace kube-system \
  --set csi.customfuse.enabled=true \
  --set csi.oss.enabled=true \
  --set csi.disk.enabled=true
```

### Verify

```shell
# Controller and node plugin (in kube-system, like other CSI drivers)
kubectl get deploy -n kube-system csi-customfuse-provisioner
kubectl get ds -n kube-system csi-customfuse-plugin

# Fuse pods (in ack-csi-customfuse, created automatically on mount)
kubectl get pods -n ack-csi-customfuse
```

## Overall Flow

Once the Helm chart is deployed, mounting a CustomFuse volume involves:

1. **Build a fuse pod image** — your FUSE client + mount-proxy + entrypoint.sh
2. **Register the image** in the `csi-plugin` ConfigMap (`kube-system` namespace)
3. **Create PV + Secret + PVC** — volumeAttributes carry config, Secret carries credentials
4. **Deploy a consumer pod** — the CSI driver creates a fuse pod automatically

All steps with complete Dockerfiles, YAML, and entrypoint examples are documented in [examples/customfuse/](../examples/customfuse/).

## Configuration

### Fuse pod image (csi-plugin ConfigMap)

Key format: `fuse-<fuseType>`, where `fuseType` matches the PV's `volumeAttributes.fuseType`.

```shell
kubectl -n kube-system create configmap csi-plugin \
  --from-literal="fuse-juicefs=image=registry.example.com/csi-fuse-juicefs:v1.0"
```

Changes take effect on the next mount via informer — no restart needed.

See [examples/customfuse/](../examples/customfuse/) for volumeAttributes reference, Secret passthrough, entrypoint examples, and a complete working demo with JuiceFS CE.
