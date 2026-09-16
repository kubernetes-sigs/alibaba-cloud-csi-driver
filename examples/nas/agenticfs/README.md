# AgenticFS dynamic provisioning

## 1. Prerequisites

The normal resource model is **one PVC → one AgenticSpace → one RAM-enabled
AccessPoint**, on an existing AgenticFS filesystem (`StorageType=Agentic`). The
CSI driver does not create or delete the filesystem in this mode. Retries recover
the same volume with the same placement parameters; this is not a mechanism for
selecting among cross-VPC access points or migrating an existing volume's network.

Use the canonical StorageClass value `volumeAs: Agentic` (case-sensitive).
The filesystem, AgenticSpace zone, VPC and vSwitch must be available to your
account, and the access-point endpoint must be reachable from the mounting nodes.
Nodes need the `alinas` mount helper with TLS and RAM authentication configured.

### Controller RAM permissions

Update the RAM role or credentials used by **csi-provisioner**, not just the node
plugin. The [NAS policy example](../../../docs/ram-policies/nas.json) includes
these AgenticFS operations:

| Operation | RAM actions |
| --- | --- |
| Provision a space | `nas:CreateAgenticSpace` |
| Discover/reuse and wait for access points | `nas:ListAccessPoints`, `nas:DescribeAccessPoint` |
| Create an access point | `nas:CreateAccessPoint` |
| Delete a volume | `nas:DeleteAccessPoint`, `nas:DeleteAgenticSpace` |
| Read and expand quota; verify absence during deletion | `nas:GetAgenticSpace`, `nas:SetAgenticSpaceQuota` |

The policy is a broad driver example using `Resource: "*"`. For a least-privilege
installation, grant only the operations needed by enabled volume modes and scope
resources where supported by NAS RAM authorization. These control-plane actions
do not replace the mount identity's data-access permissions. CNFS filesystem
creation also needs the permissions required by the separate CNFS controller.
If provisioning reports `Forbidden`, check the action, resource scope and identity
in the NAS error before retrying.

## 2. Choose a filesystem source

- **Without CNFS:** edit `storageclass-direct.yaml` with an existing AgenticFS
  `fileSystemId`, `zoneId`, `vpcId` and `vSwitchId`. Apply that StorageClass and
  change `pvc.yaml` to use `alicloud-nas-agenticfs-direct`.
- **With CNFS:** edit and apply `cnfs.yaml`, wait for the CNFS controller to populate
  `status.fsAttributes.filesystemId` and `storageType: Agentic`, then edit and
  apply `storageclass.yaml`. Apply `pvc.yaml` and `pod.yaml` afterwards.

When both are specified, `fileSystemId` takes precedence over
`containerNetworkFileSystem`. The filesystem's placement and the AgenticSpace's
`zoneId` are separate settings.

## 3. Quotas and mounting

The PVC request is rounded up to a whole GiB. The result must be at least 10 GiB,
no greater than the NAS maximum (1024000 GiB), and within both CSI `limit_bytes`
and the optional `agenticSpaceSizeLimit` StorageClass cap. A request that fits
before rounding but exceeds the cap afterwards is rejected; for example,
`10Gi + 1 byte` cannot be provisioned under a `10.5Gi` cap. Prefer whole-GiB caps.
When the request omits capacity, the cap is used as the requested size and is
subject to the same checks.

`agenticSpaceFileCountLimit` defaults to 1000000, with the NAS-supported range
10000–100000000. Expansion updates the existing space's size quota and does not
shrink it.

The controller supplies missing `tls,vers=3,ram` options. Explicit mount options
must retain `tls` and `ram`: the node validates them rather than silently repairing
a security configuration. The PV mounts `/` through the access-point domain,
which is the root of this AgenticSpace, not the entire filesystem.

## 4. Deletion and failed provisioning

With `reclaimPolicy: Delete`, normal PVC/PV reclamation removes the access points
bound to the space, waits for them to disappear, then deletes the AgenticSpace.
It never deletes the parent filesystem. Deletion is destructive; use `Retain`
when data must be preserved after the PVC is removed.

**CreateVolume never deletes resources on failure**, including terminal errors.
Creation proceeds forward on retries:

1. `CreateAgenticSpace` reuses `ClientToken=<volume name>` to replay creation.
2. `CreateAccessPoint` has no ClientToken in the current public API. The driver
   lists access points bound to the space and reuses one before attempting a new
   creation; it does not blindly create another access point on every retry.
3. The driver waits for the access point to become Active before returning a volume.

| Access point found for the Space | CreateVolume behavior |
| --- | --- |
| None | Create one access point. |
| Active | Reuse it; do not create another. |
| Pending | Reuse it and wait for Active. |
| Inactive or Deleting | Return a retryable error; do not create a replacement. |

The one-access-point model is not a claim that NAS enforces uniqueness on repeated
CreateAccessPoint calls. If recovery encounters multiple entries, it logs the
unexpected state and prefers an existing Active entry rather than creating more.
Ownership validation and pagination remain in place. DeleteVolume enumerates and
removes all access points belonging to the Space so leftovers cannot block deletion.

An Inactive access point may require operator intervention; see §4.1. Deleting a
Pending PVC alone does not guarantee that NAS resources will be released.

A timeout or missing response does not prove that creation failed in NAS. Likewise,
a later terminal error does not prove that an earlier attempt created nothing.
Failed requests log the known identifiers; a missing space ID is recorded as an
unknown resource state. Error classification changes the diagnostics and CSI
response, not whether resources are deleted.

This deliberately leaves spaces and access points for recovery. If the PVC is
abandoned before a PV is delivered, DeleteVolume may never be called. Retained
resources can consume quota and incur costs: arrange external reconciliation or
use the manual procedure below. This driver does not include an automatic reaper,
and ClientToken idempotency does not garbage-collect abandoned resources.

### 4.1 Manual cleanup

Use this procedure when provisioning leaves resources without a PV, or when a
malformed PV is missing `agenticSpaceId` / `accesspointId` and automatic deletion
cannot safely proceed.

1. **Confirm ownership and stop concurrent use.** Save the PVC/PV definitions,
   events and CSI logs. Check that no workload, PV or pending provisioning request
   still needs the resources. Suspend provisioning/retries for the affected
   volume before manual cleanup. Do not delete resources solely because a log
   mentions them.
2. **Locate the exact space.** Collect `region`, `fileSystemId`, `agenticSpaceId`,
   `accesspointId`, `fileSystemPath` and `volumeHandle` from the PV or CSI logs.
   If the space ID is missing, locate it by the exact filesystem/path in the NAS
   console or with `DescribeAgenticSpaces`, following all pages and verifying
   ownership. This operator lookup may require additional
   `nas:DescribeAgenticSpaces` permission. A volume name alone is not sufficient
   across regions or filesystems.
3. **Enumerate and verify every bound access point.** Use NAS `ListAccessPoints`
   filtered by the confirmed `AgenticSpaceId`, follow pagination, and verify that
   each returned access point belongs to that space. Include a known PV access
   point even if a recent change has not yet appeared in the list. Never delete
   an access point based only on a matching name or tag.
4. **Back up any required data, then remove the access points.** Delete only the
   verified access points and poll `DescribeAccessPoint` until NAS confirms they
   are absent. Acceptance of a delete request is not proof of completion.
5. **Delete the AgenticSpace and verify absence.** Once all access points are
   detached, call `DeleteAgenticSpace` for the verified filesystem/space IDs and
   confirm absence with `GetAgenticSpace`. Do not delete the parent filesystem.
   If any response is ambiguous, re-query rather than assume cleanup succeeded.
6. **Reconcile Kubernetes only after cloud cleanup.** If possible, repair missing
   PV attributes from the verified identifiers and let CSI complete reclamation.
   Do not remove PV finalizers merely to hide a deletion error. Resume provisioning
   only after resolving the old resources; recreating a failed PVC gives it a new
   volume name and idempotency token.

### Log interpretation

- `agenticfs-resource-created`: a space exists but has not yet been delivered;
  successful provisioning also emits this message.
- `agenticfs-resource-retained-for-retry`: intentionally retained for a retry;
  **not** a cleanup instruction.
- `agenticfs-orphan-resource`: a terminal provisioning failure requires
  reconciliation. The historical prefix does not prove a resource exists or is
  orphaned. Check live NAS and Kubernetes state before deleting anything.
- `agenticfs-orphan-resolved`: legacy message from versions that attempted partial
  rollback; no longer emitted by CreateVolume. In historical logs it referred
  only to the access point, not the AgenticSpace.

These prefixes describe attempt outcomes, not globally unique resource events.
An external reaper must correlate identifiers and current Kubernetes/NAS state;
grepping a prefix alone is not a safe deletion policy.
