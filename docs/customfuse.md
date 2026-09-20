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

## Where the driver runs

Those four steps describe the driver as the Helm chart installs it: a controller
Deployment and a node DaemonSet, and every field in this document is written against
that shape.

Some runtimes inject the node plugin into the workload's own sandbox rather than
running it on the node. customfuse serves that case from a separate entry point
(`pkg/customfuse/csi_agent.go`), which implements the Node service only. The volume
contract is identical — same `volumeAttributes`, same `pv.spec.mountOptions`, same
Secret passthrough, same environment for the entrypoint and same names it cannot take
over — so an adapter written against one runs on the other. What differs is who
provides the client process, and therefore which settings have anything to act on:

| | node DaemonSet | injected sidecar |
|---|---|---|
| Controller service | yes | no, Node service only |
| `CreateVolume`, dynamic provisioning | yes | no |
| Fuse pod | one per volume per node, created at mount time | none: mount-proxy is the sidecar itself |
| Fuse pod image, from the `csi-plugin` ConfigMap | resolved by the controller | not consulted — the injecting runtime chooses the sidecar image, which is where the client has to be |
| `entrypointConfig`, `entrypointKey`, `dnsPolicy`, `serviceAccountName` | applied to the fuse pod | no effect: they describe a pod that is never created |
| Mount layout | one shared mount per node, bound into each consumer | mounted directly on the requested target path and propagated to the workload container |
| Read-only | `$readOnly`, and a read-only bind | `$readOnly` alone |

The last row is the one to plan around. On a node the driver makes the bind into the
consumer read-only, so a client that ignores `$readOnly` still cannot be written
through. Under injection there is no bind to make read-only — the mount is the
client's own, and the workload sees it through propagation — so honouring `$readOnly`
is the adapter's job.

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
supported route. `authType` is parsed alongside it: the empty default and
`agent-identity` are accepted, and any other value fails the mount with
`unsupported authType`, so the `authType: rrsa` an OSS volume would carry has no
equivalent here yet. The same ServiceAccount field is what RRSA-style
authentication will be configured through when it arrives.

[examples/customfuse/README.md](../examples/customfuse/README.md#private-registry)
has the full recipe.

### Entrypoint process contract

mount-proxy starts the entrypoint as a child process and tracks exactly one PID:
the entrypoint's. It waits on that PID to tell a live mount from a dead one, and
when a mount times out or is torn down it signals that PID and nothing else.
There is no process group, so a child of the entrypoint is invisible to both
halves.

The entrypoint therefore has to `exec` the client as its last act, which makes
the client *be* the PID mount-proxy is watching. Every demo here ends that way.
The two ways to get it wrong fail differently:

* **Start the client and let the script exit.** mount-proxy sees its one PID go
  away and reports `entrypoint exited unexpectedly`, whether or not the mount
  came up — and if the client is still alive, nothing is tracking it any more.
* **Keep a shell in front of the client** — background it and `wait`, for
  instance. The mount works and stays tracked while the script lives, but on
  teardown only the shell is signalled. The client survives it holding a mount
  that nothing will clean up.

The client also has to run in the foreground, since a mount-proxy that outlives
the client would report a healthy mount over a dead one. How a client asks for
that is its own business: a `-f` flag, an environment variable, or nothing
because foreground is already the default.

A client that can only daemonize leaves the script no `exec` to make, so the
script has to forward the signals itself — see
[Important Notes](../examples/customfuse/README.md#important-notes) for the
shape that does.

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

`hostNetwork` has one consequence that is easy to miss, because it changes name
resolution rather than reachability. A pod on the host network that leaves
`dnsPolicy` unset gets `ClusterFirst`, which there means the *node's* resolver,
not the cluster's — so a service name in a PV, `redis.default.svc` say, does not
resolve even though the same name works from a workload pod. Set
`dnsPolicy: ClusterFirstWithHostNet` on the volume when the client has to reach
an in-cluster endpoint.

**Secret entries reach the entrypoint as environment variables.** Every key in
the volume's `nodePublishSecretRef` becomes `$key` in the entrypoint's
environment, because the driver cannot know how an arbitrary client wants to be
given a credential. Consequences worth planning for:

* Anyone who can exec into the fuse pod can read it, and it may appear in a core
  dump of the client. Workload containers cannot: the fuse pod is a separate pod,
  and under sandbox injection the sidecar shares no process namespace with them.
* An entrypoint that echoes its environment for debugging will write credentials
  to the pod log.
* A key arrives under its own name and no other. The driver derives no alternative
  spelling from one, because which names a client accepts is the client's business:
  an adapter that takes more than one resolves them itself. The OSS-compatible demo
  does that for the older `akId`/`akSecret`, so a Secret written for OSS still works
  there — through its entrypoint, not through the driver.

If your client does not need the credential once it is running, `unset` the
variables before `exec`ing it — only before an `exec`, since a script that stays
alive keeps the environment it was started with.

`authType: agent-identity` avoids this entirely on the sandbox path, where the
driver runs in an injected sidecar rather than a fuse pod: the credential is
exchanged per mount, delivered as files rather than environment variables, and
rotated before it expires. See
[customfuse-agent-identity.md](./customfuse-agent-identity.md).

**`otherOpts` reaches the entrypoint without interpretation**, as `$otherOpts`,
since only the entrypoint knows what its client's options look like. The driver
never splits it, so it also cannot merge anything into it.

`pv.spec.mountOptions` is the other channel, and it is open: an entry the driver
does not know becomes the environment variable of its own name, spelled exactly as
the PV wrote it. An entry that *does* name a driver field — `source`, `bucket`,
`url`, `path`, `otherOpts` — is applied to that field instead and not passed
through, and only where `volumeAttributes` left it empty. `volumeAttributes`
cannot be edited once the PV exists while `mountOptions` can, so this is how an
existing volume gains a missing parameter; it is not how one redefines a value a
driver consumer — the `bucket_name` metrics label — has already read. Ignored
entries are logged rather than dropped silently.

Two fields are outside that rule. `capacity` may be raised but never lowered,
which is the only way a quota can change at all on an existing volume: the driver
implements no expansion RPC, and Kubernetes itself refuses to shrink a claim.
`readOnly` is never settable from either channel, because it comes from the PV's
`accessModes` and from the publish request.

The control fields are outside it too — the ones that decide how the fuse pod is
built rather than what it mounts. They are read from `volumeAttributes` alone;
named in `mountOptions` they are logged and dropped, rather than becoming
environment variables nothing would read.

So are the names the driver fills in itself: `mountpoint`, the path it tells the
entrypoint to mount on and then waits for a mount point to appear at; and
`readOnly`, refused in both directions, so a volume can neither clear a read-only
request nor impose one on a writable volume. Nor can a volume redefine a variable the
mount-proxy already has in its own environment, since the client runs inside that
process.

One exception runs the other way: under `authType: agent-identity`, two control
fields do reach the entrypoint, because the script has to know which credential
flow ran and where the credential landed. It receives `$authType`, and
`$credentialDir` resolved to the directory actually used rather than the value the
PV named. The options configuring the exchange itself are stripped before the
entrypoint runs. See
[customfuse-agent-identity.md](./customfuse-agent-identity.md).

Those refusals are not all enforced in the same place, and the plugin and
mount-proxy are separate images, built and rolled out independently — so which log
to read depends on which rule fired:

| rule | plugin, parsing the PV | mount-proxy, building the environment |
|---|---|---|
| `mountpoint`, `source`, `readOnly` refused | yes, case-insensitively | yes, case-insensitively |
| a name mount-proxy's own environment already has | no — it cannot see that environment | yes |
| a Secret key colliding with either of the above | no — Secrets pass through unfiltered | yes |
| control fields named in `mountOptions` | yes | no |
| `capacity` lowered | yes | no |

The reserved names are the one rule both halves enforce, because a Secret reaches
mount-proxy without the plugin ever parsing it.

`source` is emitted by the driver but is *not* refused from `mountOptions`: it is one
of the driver fields above, so an entry may fill it where `volumeAttributes` left it
empty — the only way to add one to an existing PV. It then travels to mount-proxy as
the resolved value rather than as an option entry, and mount-proxy emits `$source`
from it. What the second refusal catches is a *second* `source` — a Secret key of that
name, which the plugin passes through unfiltered, or an option entry, which
mount-proxy refuses rather than assume the plugin already consumed it.

**Read-only is enforced by the driver, not left to the entrypoint.** `$readOnly`
tells the client what to do, and the bind into the consumer pod is made read-only
as well, so a client that ignores the variable still cannot be written through.
The volume is mounted once per node and bound into every pod using it, so a pod
asking for read-write after a read-only one inherits that read-only mount; give the
two access modes separate volumes.

Both halves of that belong to the node deployment. Under an injected sidecar there is
no bind to make read-only, so `$readOnly` is the whole mechanism — see
[Where the driver runs](#where-the-driver-runs).

The full precedence table is in
[examples/customfuse/README.md](../examples/customfuse/README.md#entrypoint-env-vars).

The driver refuses a newline, a carriage return, a semicolon, a backtick or `$(`
in `otherOpts` and in every `mountOptions` entry, because an entrypoint that
splices them onto a command line would run them in a privileged container. That
refusal is a floor, not a validation of your options: everything else is left
alone, since the driver cannot tell an exotic option value from an attack and
guessing wrong breaks mounts that were working. So an entrypoint that expands
`$otherOpts` unquoted is still exposed to word splitting and to the
metacharacters that remain. Keep the expansion quoted, or check the value against
what your client accepts — the examples under
[examples/customfuse/](../examples/customfuse/) show one way.

`source`, `path` and Secret values read from `volumeAttributes` are deliberately
not checked. They can legitimately contain any of those characters — a metadata
engine URL commonly embeds a password, and a password may contain a `;` — so
quoting them stays the entrypoint's job.

The same value arriving through `mountOptions` *is* checked, because that channel
is refused by entry rather than by field and cannot tell a `source=` from an
option. A `source` that has to carry those characters therefore belongs in
`volumeAttributes`, which is where the driver reads it first anyway.
