# Alibaba Cloud Disk Volume Restore

## Overview
A Disk snapshot plugin is available to help simplify storage management, and backup the volume with snapshot.
Once user creates VolumeSnapshot with the reference to a snapshot class, snapshot and
corresponding VolumeSnapshotContent object gets dynamically created and becomes ready to be used by
workloads. You can restore your volume with volumesnapshot, which use dataSource in pvc template.

Disk Snapshot Plugin will create 3 crds as below.

```
volumesnapshotclasses.snapshot.storage.k8s.io:  define details to create volumesnapshotcontents, like: storageclass;
volumesnapshotcontents.snapshot.storage.k8s.io: define one disk snapshot with the backend, like: pv;
volumesnapshots.snapshot.storage.k8s.io:        the claim of one snapshot, like: pvc;
```

## Requirements
* CSI Snapshot controller (registry.cn-hangzhou.aliyuncs.com/acs/snapshot-controller)
* CSI Snapshot external runner (registry.cn-hangzhou.aliyuncs.com/acs/csi-snapshotter).
* Disk Snapshot plugin depends on csi-plugin (registry.cn-hangzhou.aliyuncs.com/acs/csi-diskplugin).
* Secret object with the authentication key for Disk
* Service Accounts with required RBAC permissions

## Usage
### 1. Create a VolumeSnapshotClass.
```
# kubectl apply -f examples/disk/snapshot-restore/snapshotclass.yaml

# kubectl get volumesnapshotclasses.snapshot.storage.k8s.io
NAME                AGE
default-snapclass   15s
```
### 2. Create an application and write data to the application.
Create pvc, using disk volume.

```
# kubectl apply -f examples/disk/snapshot-restore/sts.yaml
# kubectl exec -it mysql-0 -- touch /data/test
# kubectl exec -it mysql-0 -- ls /data
```
Expected output:
```
lost+found test
```

### 3. Create a VolumeSnapshot.
```
# kubectl apply -f examples/disk/snapshot-restore/snapshot.yaml

# kubectl get volumesnapshots.snapshot.storage.k8s.io
NAME                AGE
new-snapshot-demo   2m53s

# kubectl get volumesnapshotcontents.snapshot.storage.k8s.io
NAME               AGE
snapcontent-****   3m22s
```
> snapshot.yaml define the data source pvc name, and snapshotClass name.
> You can check disk snapshot in alibaba cloud ecs console.
### 4. Restore data.

```
# kubectl apply -f examples/disk/snapshot-restore/sts-restore.yaml           Delete           Bound    default/disk-ssd-web-restore-0   alicloud-disk-ssd            114s
```

### 5. Run the following command to view the application data in pod mysql-restore-0:
```
# kubectl exec -it mysql-restore-0 -- ls /data
```
Expected output:
```
lost+found test
```
## Troubleshooting
