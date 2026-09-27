# Substrate CSI bridge

The bridge exposes `substrate.csi.alibabacloud.com` and translates Substrate's
logical Actor volumes into pre-generated NAS CSI publish requests. It does not
create AgenticFS data on behalf of an E2B client. The annotation producer and
the Substrate caller must implement the contracts below.

The first release forwards only `nasplugin.csi.alibabacloud.com` requests using
AgenticFS (`mountProtocol=alinas`) and Agent Identity. It does not implement OSS
forwarding or worker-side OSS deployment. This restriction belongs to the bridge,
not to the other drivers in the shared CSI process. The NAS AccessPoint
status-casing fix in this branch is a separate NAS change, not bridge routing.

## Helm deployment

The chart defaults to `enableSubstrate: false`, preserving the existing
deployment. With the option enabled it extends the existing workloads:

- The existing `csi-provisioner` and `csi-plugin` processes also serve the bridge
  driver, on a separate socket. No additional Deployment, DaemonSet or bridge
  process is created.
- The existing `alicloud-csi-provisioner` and `alicloud-csi-node` ServiceAccounts
  are retained. Only the Node workload receives an audience-bound token for the
  Substrate API; the bridge Controller does not query Actors or templates.
- One Envoy sidecar and one `csi-provisioner` Service expose the native NAS
  Controller on port 443 and the bridge Controller on port 444. Both require
  mTLS and accept Controller/Identity RPCs only.
- One `csi-provisioner-substrate` ConfigMap defines both TLS listeners.
- One additional bridge registrar is added to each existing Node Pod.
- Actor hostPath mounts with bidirectional propagation and persistent Golden
  source directories under the bridge's CSI socket directory; no binding journal.
- Both workloads retain Pod identity projections for TLS. Only the Node workload
  projects the Actor API token and Service DNS server trust; the Controller keeps
  its certificate and Pod identity client trust for inbound mTLS.
- A bridge `CSIDriver`, per-driver `CSIDriverConfig` objects, and optional
  `ate-storage` class. Native NAS configuration is emitted only when its
  Controller is enabled; two driver names cannot share one `CSIDriverConfig`.

```sh
helm upgrade --install csi deploy/charts/alibaba-cloud-csi-driver \
  --namespace kube-system \
  --set enableSubstrate=true \
  --set deploy.featureGates=AlinasMountProxy=true \
  --set images.plugin.tag=<image-containing-this-change> \
  --set images.controller.tag=<image-containing-this-change> \
  --set deploy.regionID=<region>
```

Published images predating this driver do not gain bridge support by enabling
the chart option. Use images built with the implementation. The chart does not
publish images or install the separate CNFS NAS mount broker.

Key values:

| Value | Default | Meaning |
|---|---|---|
| `enableSubstrate` | `false` | Enable all bridge resources |
| `substrate.apiEndpoint` | `api.ate-system.svc:443` | Substrate API TLS endpoint, required only for Node |
| `substrate.apiAudience` | `api.ate-system.svc` | Audience of the Node's projected API token |
| `substrate.actorRoot` | `/var/lib/ateom-gvisor/actors` | Must match atelet's host target layout; identity carrier, see Identity keys |
| `substrate.createDriverConfig` | `true` | Create the native NAS and bridge configs for enabled Controllers |
| `substrate.storageClass.create` | `true` | Create the bridge reference class |
| `substrate.storageClass.name` | `ate-storage` | Must match the annotation producer's template slots |
| `images.envoy.repo` | `acs/envoy` | TLS frontend repository, resolved through the chart's standard registry selection |
| `images.envoy.tag` | `v1.39-latest` | Published ACK mirror tag; override with your validated release tag if needed |

`controller.enabled`, `controller.replicas`, `plugin.enabled`, `nodePools`,
`deploy.kubeletRootDir`, and `imagePullSecrets` are also honored. Node pools must
be non-overlapping and use the same kubelet root for the global driver config.
The original process health ports are unchanged; there is no second bridge
process competing for them. NAS uses the existing socket selection: an explicit
`--mount-proxy-sock` takes precedence; otherwise `AlinasMountProxy=true` selects
`/run/cnfs/alinas-mounter.sock`. The feature gate is off by default, so enable it
for this deployment as shown above. The chart does not inject a socket override
into the shared CSI process. OSS retains its original flag/per-volume selection.

The `StateDir` is `/csi/substrate.csi.alibabacloud.com/substrate-state`
inside the shared Node container. It is used only for Golden source data, not
NAS metadata or binding JSON. Keep its existing node-local hostPath lifetime;
an ephemeral replacement can lose a live Golden source. Never run the old standalone Node
driver and the shared Node driver on the same node/socket during migration.

The shared Service selects `app: csi-provisioner`. `substrate-nas` points the NAS
driver to `dns:///csi-provisioner.<namespace>.svc:443`; `substrate-csi-bridge`
points the bridge to port 444. Each config retains its own Node socket. Disable
`substrate.createDriverConfig` when these resources are managed externally,
and do not leave competing configs for the same driver name.

### External prerequisites

These are not supplied by this CSI chart:

1. Substrate control plane and the `CSIDriverConfig` CRD.
2. `PodCertificate` and `ClusterTrustBundle` projection support, with the
   `servicedns.podcert.ate.dev/identity` and
   `podidentity.podcert.ate.dev/identity` signers and live trust bundles.
3. Substrate authorization permitting `alicloud-csi-node` to query the intended
   Actors and templates. The bridge Controller needs no such permission.
   Kubernetes RBAC alone does not
   configure Substrate's application authorization policy.
4. An operational NAS mount broker, AgentIdentity credential service, NAS/AP
   connectivity and the required per-Actor permissions.
5. A Substrate API with `ResourceMetadata.annotations`. Public upstream does
   not currently expose that field; see the API client provenance below.

Do not turn off TLS verification to work around missing CA/signers. No CA
private key, cloud key or business token is placed in this chart.

The shared identity projection puts the certificate at
`/run/podidentity.podcert.ate.dev/credential-bundle.pem`. The Node additionally
projects Service DNS server trust at `trust-bundle.pem` in that directory,
matching the native Substrate credential-path contract. Envoy uses a separate `client-trust-bundle.pem` from
the Pod identity signer for inbound client verification. The directory name
does not determine which issuer a trust bundle must contain.

The CNFS mount broker is a separate workload: these CSI Pod projections do not
automatically configure its credentials. The broker must independently have the
client identity and correct trust for its credential endpoint.

### Actor API trust rotation

The Node's Actor API client loads its CA bundle once at construction. Kubelet
updates the projected file, but the client does not reload its in-memory roots.
Roll the CSI Node Pods after adding or removing trusted CAs. For a planned root
rotation, publish a bundle containing both roots, roll the Node Pods, then switch
the API certificate; after removing the old root, roll the Node Pods again.
Renewing a server leaf certificate under the same trusted CA does not by itself
require this restart.

Existing TLS connections may remain usable while new connections fail against
an unrecognized CA. The existing liveness/readiness probes check the shared CSI
process, not Actor API connectivity, and do not automatically recover stale
trust. Do not turn an ateapi outage into shared CSI liveness failure: that would
restart unrelated drivers too. Dynamic CA reload is not implemented here.

The projected bearer token is read on each RPC. TLS retains a minimum of 1.2 and
can negotiate 1.3; changing the protocol minimum is separate from trust rotation.

## Two requests, one authority

Substrate's request describes a virtual bridge volume. The entry in
`ate.dev/csi-volume-publish-requests` describes the actual NAS volume:

```json
[
  {
    "volumeName": "data",
    "driver": "nasplugin.csi.alibabacloud.com",
    "request": {"volumeId": "backend-id", "targetPath": "/data"}
  }
]
```

This is a structural example; a real request also needs filesystem capability,
AgenticFS/AgentIdentity attributes and its credential provider. The bridge puts no
size limit on the annotation of its own accord. Never put credentials in
annotations/VolumeContext.

The actual request remains authoritative for backend ID, server/path, fsType,
mount options, provider and PublishContext. The bridge changes only routing and
identity metadata: `targetPath` becomes the host directory the caller gave,
`substrate-mode` is added so that NAS skips kubelet semantics, and the Actor
identity keys are set from the verified Actor, while standard PodInfo comes from
the current worker assignment. Storage ownership and credential ResourceID remain
Actor-scoped; the worker UID/name/namespace describe placement and can change on
resume. Stored PodInfo is not a source of current placement labels; missing
outer Pod name/namespace labels remove their stored counterparts.
Nothing else is merged from the outer request: no arbitrary VolumeContext entry,
no fsType, no mount flag.

Read-only configuration belongs to the annotation request. Its `Readonly`,
access mode, mount flags and `VolumeContext.options` are forwarded unchanged.
The NAS driver applies their semantics and option precedence; the bridge does
not normalize conflicting `rw` options or reinterpret options overridden by
mount flags.

The outer request contributes no read-only configuration. The bridge does not
read its `Readonly`, access mode, mount flags or `VolumeContext.options`, and it
does not reject them either: whatever the Substrate StorageClass, PVC or caller
asks for, the forwarded request keeps the read-only state of the Actor publish
request. This is the passthrough rule of the design document, whose purpose is
that the Substrate and ACS Sandbox forms reach the NAS driver with the same
read-only configuration; a bridge that "corrects" the request would make the two
forms differ exactly where neither form can see the difference. The cost is a
deviation from the CSI wording that the SP MUST honour `readonly` on publish: an
outer read-only request against a writable Actor request is published writable
rather than honoured or refused. Configure business read-only access in the
annotation producer, and treat a bridge StorageClass `mountOptions: [ro]`, a
`ReadOnlyMany` PVC or a caller-set `readonly` as unconfigured, not as
enforcement. Golden placeholders are writable bind mounts.

For an existing mount at the target, the bridge asks one question: is the mount
visible there one it may mount over? NFS, NFS4 and alinas are (an alinas AccessPoint
mount is performed as fstype `alinas` but shows up as `nfs` in the kernel mount
table, so both spellings belong to the answer), anything else is refused, because
stacking a NAS mount over foreign storage would hide it and would break the
unpublish rule "a mount the bridge owns is either a NAS mount or its own Golden
placeholder". Whether a live NAS mount matches the new request is not compared:
NAS itself answers an already-mounted target with success, so the mount keeps the
read-only state it was created with until it is unpublished. That is the caller's
obligation, not a check the bridge adds, and adding it would turn a retryable
publish into one that only a manual unpublish can clear.

## Identity keys

| Key | Meaning / producer |
|---|---|
| `csi.alibabacloud.com/actor.uid` | Actor UID; at publish checked against the Actor directory in the mount target and against the Actor API response |
| `csi.alibabacloud.com/actor.name` | Actor name used for `GetActor` |
| `csi.alibabacloud.com/actor.namespace` | Actor atespace, not the Kubernetes namespace |
| `csi.storage.k8s.io/pod.uid` | Current worker Pod UID, required on publish; supplied by atelet's worker assignment |
| `csi.storage.k8s.io/pod.name` | Current worker Pod name, when supplied by the control plane |
| `csi.storage.k8s.io/pod.namespace` | Current worker Pod Kubernetes namespace, not the Actor atespace |
| `csi.alibabacloud.com/substrate-mode` | Selects Substrate path/credential handling |

The Actor keys are sent as CreateVolume parameters and returned in VolumeContext.
Actor identity is per-volume-stable; worker PodInfo is supplied on each Run/Restore
and must not be confused with Actor identity. On the storage side, `sandboxId`/credential `ResourceID`,
volume limits and EFC ownership all resolve to the Actor UID
(`utils.MountOwnerUID`), which is what keeps the Substrate and the ACS Sandbox forms
on the same data-access identity.

Required caller integration:

- ateapi supplies Actor UID/name/atespace to bridge CreateVolume and persists the
  returned opaque handle and VolumeContext without reconstructing them.
- Before Run/Restore, the current worker assignment supplies PodInfo. Atelet
  requires `TargetAteomUid` and uses that worker UID, not the Actor UID.
- Native NAS requests select Substrate mode only for
  `nasplugin.csi.alibabacloud.com` with `volumeAs=Agentic`; bridge references
  receive the mode marker from bridge CreateVolume.
- The annotation producer writes the backend request without credentials. Actor
  identity is validated and injected by the bridge at publish; a stored Actor
  identity that disagrees with the lookup is rejected. No additional role key is
  required. `GetActor` uses atespace/name, and Golden verification also uses
  `GetActorTemplate`; no UID-only lookup extension is required.

`attachRequired=false` controls Kubernetes attachment, not Substrate's direct
calls. Its CSI adapter calls ControllerPublishVolume and skips only
`Unimplemented` errors; the bridge declares PUBLISH_UNPUBLISH and returns a no-op
success. NodeStage/NodeUnstage are also successful no-ops. The resulting outer
staging path is not substituted into the stored inner request.

The mount target, not the VolumeId, carries volume identity. Substrate has not
stabilized its logical volume naming, so `substrate.actorRoot` is the only layout
the bridge reads: a publish or unpublish target must be a clean absolute path
inside the actor root with at least two elements below it, whose first element is
the Actor UID and whose last element selects the Actor publish-request entry.
The VolumeId is opaque to the bridge; it is echoed back from `CreateVolume` and
forwarded to NAS as a lock and log key. This keeps unpublish workable without a
VolumeContext or an Actor lookup, which `NodeUnpublishVolumeRequest` provides
neither of.

CreateVolume validates only its request: capacity range,
absence of content source/secrets, and non-empty Actor parameters. It does not
interpret the caller's declared volume capabilities: the shape that is actually
mounted is the inner request validated at publish, and only
`ValidateVolumeCapabilities` answers the filesystem-only question. Nor does it
compare the declared identity with the Actor API; the
parameters are echoed into VolumeContext and verified at publish, where
`GetActor` is the authority and a value that disagrees with the Actor directory
in the target is rejected. It returns the logical ID,
requested capacity, Substrate mode and Actor keys. It performs no Actor/template
lookup, annotation parsing, backend provisioning or driver-side state write.
Missing or invalid annotations and unavailable Actor APIs are handled at publish,
not at create. Reaching this RPC requires a client certificate from the Substrate
pod-identity CA (`require_client_certificate` on the bridge listener), and Envoy
applies no per-method authorization beyond that, so every holder of such a
certificate may create logical volumes for any Actor. Placement authority is the
trust boundary here: the bridge treats "the Actor API confirms this identity owns
that template volume" as sufficient, and does not decide whether the requester
may access the Actor.

The bridge is stateless. NodePublish reads the current Actor annotation, validates
it and forwards the inner request to the existing NAS Node service. Golden uses a
deterministically named source directory and a normal bind mount. The in-memory
per-target lock (`utils.VolumeLocks`, the same helper the NAS, disk, OSS, BMCPFS
and customfuse servers use) serializes operations on one target within the CSI
process.

There is no configuration digest, write-ahead binding or persistent intent.
Failed calls cannot leave such metadata behind to block a corrected request.
The only retained SHA256 is a deterministic Golden directory name derived from
the validated target alone; it is not a configuration fingerprint or conflict
check. A different caller VolumeId at the same target does not select another
source directory. This does not require Substrate to change its stored handle.

Unpublish needs no Actor lookup. It validates the target against the actor root
and reads the live
mount table: an absent mount succeeds; NFS/NFS4/alinas is passed to NAS unpublish
with logicalID as its lock/log key; a verified Golden source is unmounted and
cleaned up; other mounts are rejected. No broker protocol or server change is
required. NAS does not depend on StateDir or any old JSON files.

Golden source directories hold temporary volume data written during template
warmup, not the Golden snapshot itself. Normal unpublish removes that source
after checking the visible mount and remaining references. Ownership refusals
return FailedPrecondition; filesystem inspection and cleanup failures return
Internal with the failed operation.

There is no startup or periodic placeholder GC. A failed bind, interrupted
cleanup, or missing unpublish can leave source data behind. If another mount is
revealed after unmounting the Golden bind, cleanup is refused and the source is
retained. A retry with no live target mount succeeds without sweeping old source
directories. Operators must verify that the actor/runtime no longer uses the
source and that no mount references remain before removing residual data; age or
absence from one mount namespace alone is not proof. Never automatically unmount
foreign storage or recursively delete an uncertain target. Monitor the node's
disk and inode use when running repeated Golden builds.

This unreleased source-key change does not read the old VolumeId-based layout.
Before testing an upgrade with old Golden mounts, unpublish them with their
owning version and safely clean any remaining test sources. Do not delete the
whole StateDir or retain old mounts merely to make a restore pass.

The trusted controller/atelet must perform Unpublish before changing the
configuration of an already-mounted target. Backend idempotence does not prove
that an existing NAS mount matches a new annotation; the stateless bridge does
not attempt that comparison. The
caller must preserve the `substrate-mode` and Actor identity keys returned by
CreateVolume and add the current worker PodInfo. The integrated atelet writes
`GetTargetAteomUid()` into `pod.uid`; older PoC builds that wrote the Actor UID
must not be used as evidence for this worker-identity contract.

When upgrading from the cross-stage-digest implementation, upgrade all Node
instances before the Controller. New Nodes ignore an old digest left in persisted
VolumeContext; old Nodes still require one and cannot consume new Controller
responses. A single Helm upgrade rolls both workloads concurrently, so a volume
created during that window may fail its first publish until the Node restarts and
the publish is retried. Old binding JSON is ignored by this version. Do not remove
the containing StateDir wholesale: its Golden source data is still needed.

The broker-owned lifecycle alternative is archived on
`backup/substrate-lifecycle-v1`, with its design and implementation evidence at
`examples/substrate/bridge/LIFECYCLE-DESIGN.md` and
`examples/substrate/bridge/LIFECYCLE-IMPLEMENTATION.md` on that branch. It is not included
in this PR. Reconsider it only with production evidence or a changed lifecycle
contract, rather than adding permanent broker interfaces to a temporary bridge.

Substrate skips pod-oriented filesystem metrics files. It does not redirect them
into actor directories. Generic CSI RPC metrics and the reviewed short
`substrate` metric label remain enabled.

## API client provenance

The driver uses generated `ControlClient.GetActor` and `GetActorTemplate`, checks
the returned UID, and verifies Golden Actor/template association. There are no
handwritten RPC paths.

`ActorLookup` returns the resolved Actor's metadata. A Golden result requires a
verified `ate-golden` Actor whose name matches its template UID. Node repeats the
invariant checks for alternative lookup implementations; these are not separate
failure paths through a successful built-in client lookup. Actor and optional
template queries share a ten-second lookup budget, bounded by an earlier caller
deadline when present.

`pkg/substrate/internal/ateapipb` is a minimal generated projection of the
deployed Substrate API: only the two lookups above and the message closure they
need, with verbatim message, field and RPC names and numbers and omitted field
numbers marked `reserved`. `SOURCE.md` records the source revision, the
extraction rule and the refresh procedure. The projection exists because public
upstream currently omits annotations; it is not advertised as compatibility
with an unextended public server. Once an annotation-capable lightweight API
module is published, replace the projection with that dependency; refresh it
only by re-extraction, never by hand-editing.

## Validation and remaining system work

The following earlier pre-release checks were deliberately removed. These are
observable relaxations, not behavior-preserving refactors:

| Removed bridge check | Current behavior |
|---|---|
| 256 KiB annotation ceiling | No additional bridge ceiling; producer/API limits remain independent |
| Stored TargetPath must be a clean absolute path | Replaced with the validated outer host target |
| Stored StagingTargetPath must be empty | Retained in the inner request; NAS does not consume it at publish |
| Live mounted readonly check | No access-mode comparison against an existing mount; configuration changes require unpublish first |

Local tests cover inner read-only passthrough, outer read-only signals being
ignored, the live-mount filesystem check including a mount stacked over a NAS
mount, the Actor identity of the target directory
against the caller's and the Actor API's claims,
current annotation forwarding, live-based cleanup, corrected requests after
failed publish/restart, target serialization and Golden source reuse,
and generated-client TLS/token rotation. Helm tests cover disabled output,
enabled resources, node-only API trust/token wiring, Controller-only operation
without Actor API configuration, socket paths and component switches.
After changing Substrate-related code or chart templates, run the local self-check
with Go and Helm installed:

```sh
bash hack/check-substrate-helm.sh
```

It runs the render tests and Helm lint with Substrate disabled and enabled. It
does not require a cluster or deploy resources, and has no dedicated workflow.

On an isolated privileged Linux container, set `BRIDGE_REAL_MOUNT_TEST=1` and
select `-test.run 'Real'` to run the real Golden bind/unbind, changed-VolumeId,
writable access and stacked-mount ownership tests. Confirm that they ran rather
than being skipped; `RealMount` alone misses the stacked-mount test. The optional
`TestActorLookupLive` performs only API reads when `SUBSTRATE_LIVE_*` is set.

These do not prove the full system sequence:

```text
Golden publish → checkpoint → placeholder unpublish
→ business Actor publishes real storage → restore Golden snapshot
```

That sequence needs the Substrate/runtime integration environment. Previously
observed restore failures depended on deleted source paths, and same-Actor FULL
restore could invalidate guest bind references after backend replacement. This
CSI change does not claim to fix those VMM/runtime issues. Preserve a separate
system-level result for the exact deployed versions; do not weaken cleanup or
leave source placeholders mounted merely to make restore pass.

## File responsibilities

The bridge has no binding persistence implementation. Its internal `ateapipb` package is a minimal generated
projection of the deployed API plus one provenance file; most of its added lines
are generated API definitions, not additional runtime components. The small
metadata checks were folded into `utils.go`, and shared identity helpers now
live in the existing `pkg/mounter/utils/agentidentity` package rather than a new
`pkg/volumecontext`.

| Files | Responsibility |
|---|---|
| `actor_client.go` | TLS/token connection and generated API calls |
| `internal/ateapipb/*` | Minimal generated API projection and provenance |
| `controller.go` | Logical CSI Controller lifecycle |
| `node.go` | Node publish/unpublish and target validation |
| `utils.go` | Publish-time Actor validation and annotation selection, live mount inspection and Golden source verification; no readonly union or live readonly comparison |
| `*_test.go` | One test file per unit above, plus `integration_test.go` and `integration_linux_test.go` for cross-stage and real-mount coverage |
| Helm `plugin.yaml` / `controller.yaml` | Shared processes, registrar, TLS sidecar and mounts |
| Helm `_substrate.tpl` / `substrate-support.yaml` | One identity/flag definition, shared Service/TLS configuration and per-driver configs |

`utils.SubstrateModeKey`, `utils.IsSubstrateVolumeContext` and
`utils.MountOwnerUID` in `pkg/utils` are the single definition of the Substrate
mode key and its parser; `jwtauth`, `pkg/common`, `pkg/nas` and the bridge all
share them. Only the exact string `true` enables the mode.

Stateless forwarding, live mount inspection, Actor lookup on publish, no-op
logical Create/Attach/Stage and Actor-independent Unpublish are the selected
design.
