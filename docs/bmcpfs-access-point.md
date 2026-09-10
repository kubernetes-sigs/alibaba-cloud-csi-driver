# Mounting BMCPFS via an Access Point (AP)

This document describes how to mount a file system through a CPFS Access Point with the bmcpfs driver (`bmcpfsplugin.csi.alibabacloud.com`). It covers anonymous mounts and RAM-authenticated mounts (static AccessKey, or STS credentials with rotation).

## Overview

| Capability | Description |
| --- | --- |
| AP mount | The access point is specified by `accessPointId`. Both `tcp` and `vsc` network types are supported; the driver picks one based on the node type. |
| Anonymous mount | Mount through an AP only, without configuring `nodePublishSecretRef`. |
| RAM auth: static AK | The Secret holds two keys (AK/SK), written to `g_unas_AKFile`. Hot reload is not supported. |
| RAM auth: STS rotation | The Secret holds an STS credential triple and can be rotated externally. It is written to `g_unas_STSFile`, and the EFC client re-reads it and refreshes the signature every 10 minutes. |
| Automatic auth mode detection | No `authType` parameter is needed. The driver infers AK vs. STS mode from the set of keys in the Secret referenced by `nodePublishSecretRef`. |
| Multiple APs in one Pod | A single Pod can mount several APs of the same file system at once (one PV per AP). |

STS rotation builds on the upstream Kubernetes `CSIDriver requiresRepublish` mechanism: kubelet periodically calls NodePublishVolume again with the latest Secret contents, the driver atomically rewrites the STS credential file when it detects a change, and the EFC client picks up the new credentials in its next 10-minute cycle. No Pod restart and no mount interruption are involved.

## Prerequisites

1. **CSI component version**: install a csi-plugin / csi-provisioner version that includes bmcpfs AP support, and enable it in the helm values:

   ```yaml
   csi:
     bmcpfs:
       enabled: true
   ```

   The bmcpfs CSIDriver object renders `requiresRepublish: true` by default (required for STS rotation); no extra switch is needed.

2. **EFC client version**: the EFC client on the node must support the `accesspoint`, `g_unas_AKFile`, and `g_unas_STSFile` mount options.
3. **fileserver cluster configuration**: the CPFS fileserver cluster must have the flag `efc_pov_UmmSigningRegion=<current region>` set. This is the server-side prerequisite for RAM auth signature verification; contact the CPFS team to have it configured.
4. **AP created**: create the Access Point beforehand via the NAS console or OpenAPI and obtain its `ap-` prefixed ID.

## Volume specification

### volumeHandle uniqueness (important)

When a single Pod references multiple APs of the same file system, kubelet deduplicates volumes by `volumeHandle`. Therefore **the volumeHandle of each AP PV must be unique**, in the following format:

```
volumeHandle: "<bmcpfsId>+<unique suffix>"
# Using the AP ID as the suffix is recommended: cpfs-0123456789+ap-aaaaaaaa
# The existing <bmcpfsId>+<filesetId> format for fileset volumes stays compatible and needs no change
```

The first segment must be the file system ID, which the driver uses to perform the attach. The suffix exists purely to guarantee uniqueness and may be anything (the AP ID is recommended, as it makes operations easier to correlate). The AP that actually gets mounted is determined by `volumeAttributes.accessPointId`, not by the suffix.

### volumeAttributes

| Key | Required | Description |
| --- | --- | --- |
| `vpcMountTarget` | One of this or vsc, or both | VPC mount target domain (used for the tcp link) |
| `vscMountTarget` | One of this or vpc, or both | VSC mount target domain (used for the vsc link) |
| `accessPointId` | Required for AP mounts | Access Point ID. A non-empty value enables AP mounting. |

### Auth mode detection and validation

The driver does not use an `authType` parameter. It infers the mode from the set of keys in the Secret referenced by `nodePublishSecretRef`:

| Secret contents | Detected mode |
| --- | --- |
| No `nodePublishSecretRef` configured | Anonymous mount |
| Exactly `accessKeyId` and `accessKeySecret` | AK mode |
| `accessKeyId`, `accessKeySecret`, `securityToken` (optionally `expiration`) | STS mode |
| Any other shape (missing keys, empty values, unknown keys) | Mount rejected (`InvalidArgument`) |

Notes:

- Key-set matching is a strict allowlist. A typo (such as `security_token`) is rejected outright rather than silently degrading to AK mode or an anonymous mount.
- When a Secret is configured, `accessPointId` must be non-empty (RAM auth applies to AP mounts only).
- The auth mode is fixed at first mount. If the Secret shape changes afterwards (for example AK keys replaced by an STS triple), the driver logs a warning and ignores it; recreate the Pod to remount with the new mode.
- Writing `g_unas_AKFile` / `g_unas_STSFile` by hand in the PV `mountOptions` is not supported. Such options are stripped by the driver with a warning, since credential file paths are managed centrally by the driver.

## Usage

### Example 1: anonymous AP mount

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: bmcpfs-ap-pvc
  namespace: default
spec:
  accessModes: ["ReadWriteMany"]
  storageClassName: ""
  resources:
    requests:
      storage: 500Gi
  volumeName: bmcpfs-ap-pv
```

The mount command generated by the driver is equivalent to (tcp link):

```bash
mount -t alinas -o efc,protocol=efc,net=tcp,fstype=cpfs \
  -o accesspoint=ap-aaaaaaaa \
  cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com:/ <targetPath>
```

### Example 2: RAM auth with a static AccessKey

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: bmcpfs-ak-secret
  namespace: kube-system
type: Opaque
stringData:
  accessKeyId: "LTAI5t********"
  accessKeySecret: "********"
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-ak-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
    nodePublishSecretRef:
      name: bmcpfs-ak-secret
      namespace: kube-system
```

The driver detects that the Secret holds only the two AK keys and mounts in AK mode automatically.

**Note**: AK configuration does not support hot reload. After changing the AK in the Secret, recreate the Pods using the volume to trigger a remount.

### Example 3: RAM auth with rotating STS credentials

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: bmcpfs-sts-secret
  namespace: kube-system
type: Opaque
stringData:               # rotated periodically by an external system
  accessKeyId: "STS.xxx"
  accessKeySecret: "********"
  securityToken: "********"
  expiration: "2026-08-10T12:00:00Z"   # optional, used for warnings only
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-sts-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
    nodePublishSecretRef:
      name: bmcpfs-sts-secret
      namespace: kube-system
```

The driver detects `securityToken` in the Secret and mounts in STS mode automatically.

**STS rotation notes**:

- The external system only needs to update the credential triple in the Secret; no action on the node or the Pod is required.
- An md5 field is **not** needed in the Secret. The driver computes `md5(accessKeyId + accessKeySecret + securityToken)` itself when writing the STS credential file.
- Propagation path for new credentials: Secret update → picked up by the next periodic kubelet republish → driver atomically rewrites the STS file → EFC loads it in the next 10-minute cycle. **End-to-end propagation is bounded by roughly the republish interval plus 10 minutes**, so make sure the new STS credential stays valid well beyond that window when rotating (30 minutes or more is recommended).
- Pods on the same node sharing one PV also share a single credential file, as credentials are configured per PV.

### Example 4: multiple APs of the same file system in one Pod

Create one PV/PVC pair per AP (only the volumeHandle suffix needs to differ) and reference both from the Pod:

```yaml
# PV/PVC definitions are the same as in Examples 1-3, with volumeHandles:
#   cpfs-0123456789+ap-aaaaaaaa
#   cpfs-0123456789+ap-bbbbbbbb
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bmcpfs-multi-ap-app
spec:
  replicas: 2
  selector:
    matchLabels: {app: bmcpfs-multi-ap-app}
  template:
    metadata:
      labels: {app: bmcpfs-multi-ap-app}
    spec:
      containers:
        - name: app
          image: registry.cn-hangzhou.aliyuncs.com/acs/busybox:latest
          command: ["sleep", "infinity"]
          volumeMounts:
            - {name: data-a, mountPath: /data-a}
            - {name: data-b, mountPath: /data-b}
      volumes:
        - name: data-a
          persistentVolumeClaim: {claimName: bmcpfs-ap-a-pvc}
        - name: data-b
          persistentVolumeClaim: {claimName: bmcpfs-ap-b-pvc}
```

The two APs are mounted and authenticated independently (they may reference different Secrets) and do not affect each other.

## Behavior and constraints

### Unmount and attach retention

Multiple AP/fileset PVs of the same file system on one node share a single CPFS↔VSC attach. To keep unmounting safe in multi-AP scenarios, this version **does not detach** volumes whose volumeHandle carries a `+` suffix (AP volumes and fileset volumes). Instead it returns an error directly (`InvalidArgument`: `Volume with suffix is not detachable, please use the skip detach feature`):

- Unmount requests keep failing and are retried, the VolumeAttachment is retained, mounts of other APs on the same node are not broken by mistake, and running Pods are unaffected.
- To let AP/fileset volumes complete unmounting normally (on Pod deletion or PVC deletion), set `SKIP_BMCPFS_DETACH=true`: unmount succeeds immediately without performing a detach.
- The attach between the file system and the node VSC is retained and cleaned up automatically when the VSC or the node is reclaimed.
- To reclaim it earlier (for example when hitting the VSC attach quota), confirm the node has no active mount of that file system left, then detach manually via the NAS OpenAPI `DetachVscFromFilesystems`.
- A later version plans to restore the standard reference-counted detach behavior through an enhanced external-attacher approach. That constraint will then be lifted, with no change required to PV configuration.

### Other constraints

| Constraint | Description |
| --- | --- |
| accessModes | `ReadWriteMany` is recommended for AP volumes |
| AK hot reload | Not supported; changing the AK requires recreating the Pod |
| Credential files | Managed by the driver under `/run/cnfs/efc-credentials/<volumeId>/` on the node (directory 0700, files 0600, on tmpfs) and cleaned up automatically after the volume is unmounted. Do not modify them by hand. |
| Secret contents | The driver does not verify that the credentials match the AP permissions. Insufficient permissions show up as a failed mount or IO errors, so check the AP's RAM policy. |

## Troubleshooting

| Symptom | Where to look |
| --- | --- |
| Mount fails with an `InvalidArgument` event | Check that the Secret key set matches the AK / STS allowlist (spelling, extra keys, empty values) and that `accessPointId` is configured |
| Only one of several APs is mounted, or the wrong AP is mounted | Check that each PV's volumeHandle is unique (kubelet deduplicates by handle, so duplicates result in only one mount) |
| Mount fails with an EFC auth error | Confirm the fileserver has `efc_pov_UmmSigningRegion` configured; verify the AK/STS credentials are valid and the AP RAM policy grants access |
| A new STS credential does not take effect after rotation | Check in order: the Secret is updated → `/run/cnfs/efc-credentials/<volumeId>/sts.json` on the node is updated (driver side) → wait for the EFC 10-minute read cycle (client side). If the file is not updated, check that the CSIDriver object has `requiresRepublish: true` |
| IO failures caused by an expired STS credential | Check whether the external rotation system stopped updating the Secret; confirm the rotation period satisfies "remaining validity ≥ republish interval + 10 minutes" |
| Driver logs | Node side: the csi-plugin DaemonSet Pod (bmcpfs-related logs). Mount execution side: the alinas mount-proxy logs. |

Quick way to verify the STS file was updated (run on the node):

```bash
ls -l --time-style=full-iso /run/cnfs/efc-credentials/<volumeId>/sts.json   # check the modification time
```
