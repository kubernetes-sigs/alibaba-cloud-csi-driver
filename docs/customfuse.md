# CustomFuse CSI Driver

The CustomFuse driver provides a generic framework for running **any FUSE client** as a managed Kubernetes pod. You bring your own FUSE client image and entrypoint; the driver handles pod lifecycle, mount propagation, and credential injection.

For a complete working example, see the [JuiceFS CE demo](../examples/customfuse/).

## Installation

CustomFuse can be installed independently or alongside other CSI drivers (OSS, NAS, etc.).
The controller and node plugin deploy in `kube-system` (same as other drivers);
only the managed **fuse pods** run in the dedicated `ack-csi-customfuse` namespace.

### Requirements

Kubernetes **1.32 or newer**.

The controller learns the fuse pod image from the `csi-plugin` ConfigMap through an
informer that lists it by name. RBAC's `resourceNames` scopes a `list` or a `watch` only
from 1.32, when the `AuthorizeWithSelectors` feature gate turned on by default; before
that it applies to `get` alone. So on an older cluster the chart's `customfuse-controller`
Role authorizes nothing the informer can use, its list stays Forbidden, and the controller
exits after ~30s with an error naming the gate rather than retrying forever.

The node plugin is unaffected — it reads the same ConfigMap with a `get`, which
`resourceNames` has always scoped.

If you cannot upgrade, drop `resourceNames` from that Role's `configmaps` rule so the
informer can list every ConfigMap in `kube-system`. That gives up the scoping the rule was
written with, so prefer upgrading.

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

### Fuse pod ServiceAccount

A volume can name the ServiceAccount its fuse pod runs as:

```yaml
spec:
  csi:
    volumeAttributes:
      serviceAccountName: my-fuse-sa
```

Unset, the pod uses `ack-csi-customfuse`'s `default`. The account has to exist in
that namespace, because that is where the fuse pod is created — Kubernetes does not
look in the consumer's namespace.

The fuse pod never calls the API server, so this is not about permissions. It is
how the pod reaches a **private registry**: the driver puts no `imagePullSecrets`
on the pod, and the chart's `imagePullSecrets` cover only the workloads the chart
renders, not pods created at mount time. Kubernetes honours `imagePullSecrets`
from the ServiceAccount or the pod spec, which leaves the ServiceAccount as the
supported route. The same field will carry RRSA-style authentication once
`authType` supports it.

[examples/customfuse/README.md](../examples/customfuse/README.md#private-registry)
has the full recipe.

## Security model

customfuse runs a FUSE client you supply, and hands it credentials you supply.
Both are more exposed than in a driver that knows its client, so the trade-offs
are stated here rather than left to be discovered.

**The fuse pod is privileged.** Establishing a FUSE mount that other pods can see
requires it, and `hostNetwork: true` lets the client reach storage endpoints the
way the node does. Treat the choice of its image and its entrypoint as a
privileged operation, and restrict who can write the `csi-plugin` ConfigMap (which
selects the image) and the `entrypointConfig` ConfigMaps (which supply the
script).

**Secret entries reach the entrypoint as environment variables.** Every key in
the volume's `nodePublishSecretRef` becomes `$key` in the entrypoint's
environment, because the driver cannot know how an arbitrary client wants to be
given a credential. Consequences worth planning for:

* Anyone who can exec into the fuse pod can read it, and it may appear in a core
  dump of the client. Workload containers cannot: the fuse pod is a separate pod,
  and under sandbox injection the sidecar shares no process namespace with them.
* An entrypoint that echoes its environment for debugging will write credentials
  to the pod log.
* For compatibility with the OSS convention, `akId` and `akSecret` are also
  exposed as `accessKeyId` and `accessKeySecret` when the latter are absent, so a
  Secret written for OSS works unchanged. Both spellings are then present.

If your client does not need the credential once it is running, `unset` the
variables before `exec`ing it — only before an `exec`, since a script that stays
alive keeps the environment it was started with.

`authType: agent-identity` avoids this entirely where it is available: the
credential is exchanged per mount, delivered as files rather than environment
variables, and rotated before it expires. See
[customfuse-agent-identity.md](./customfuse-agent-identity.md).

**`otherOpts` reaches the entrypoint without interpretation**, as `$otherOpts`,
since only the entrypoint knows what its client's options look like. Every
`pv.spec.mountOptions` entry travels the same way, as the environment variable its
own name.

The driver refuses a newline, a semicolon, a backtick or `$(` in either, because
an entrypoint that splices them onto a command line would run them in a privileged
container. That refusal is a floor, not a validation of your options: everything
else is left alone, since the driver cannot tell an exotic option value from an
attack and guessing wrong breaks mounts that were working. So an entrypoint that
expands `$otherOpts` unquoted is still exposed to word splitting and to the
metacharacters that remain. Keep the expansion quoted, or check the value against
what your client accepts — the examples under
[examples/customfuse/](../examples/customfuse/) show one way.

`source`, `path` and Secret values are deliberately not checked. They can
legitimately contain any of those characters, so quoting them stays the
entrypoint's job.
