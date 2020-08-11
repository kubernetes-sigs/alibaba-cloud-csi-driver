# Alibaba Cloud Disk Volume Restore

## Overview

An Disk snapshot plugin is available to help simplify storage management, and backup the volume with snapshot.
You can restore your volume with volumesnapshot, which use dataSource in pvc template.


## Requirements

Same Requirements with VolumeSnapshot feature.

* CSI Snapshot external runner (registry.cn-hangzhou.aliyuncs.com/plugins/csi-snapshotter:v1.0.0).
* Disk Snapshot plugin depends on csi-plugin (registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin).
* Secret object with the authentication key for Disk
* Service Accounts with required RBAC permissions

### Feature Status
Beta


## Usage

### 1. Create pod with disk volume
Create pvc, using disk volume.

```
# kubectl get pod
NAME                                  READY   STATUS    RESTARTS   AGE
web-0                                 1/1     Running   0          79m

# kubectl get pvc
NAME              STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS        AGE
disk-ssd-web-0    Bound    pvc-59dc9df7-aaa1-4395-8d63-0dc34b3ac5e0   20Gi       RWO            alicloud-disk-ssd   79m
```


### 2. Create Snapshot with pvc
```
# kubectl apply -f examples/disk/snapshot-restore/snapshot.yaml

# kubectl get volumesnapshots.snapshot.storage.k8s.io
NAME                AGE
new-snapshot-demo   2m53s

# kubectl get volumesnapshotcontents.snapshot.storage.k8s.io
NAME                                               AGE
snapcontent-8de7d7fc-f53e-47fe-b468-cafba124c3bf   3m22s

```

### 3. Restore new statefulset with VolumeSnapshot

```
# kubectl get pod
NAME                                  READY   STATUS    RESTARTS   AGE
web-restore-0                         1/1     Running   0          44s

# kubectl get pvc
NAME                     STATUS   VOLUME                                      CAPACITY   ACCESS MODES   STORAGECLASS        AGE
disk-ssd-web-restore-0   Bound    disk-e2b037f4-1cb3-42d7-9753-684566b4c1e9   20Gi       RWO            alicloud-disk-ssd   100s


# kubectl get pv
NAME                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                            STORAGECLASS        REASON   AGE
disk-e2b037f4-1cb3-42d7-9753-684566b4c1e9   20Gi       RWO            Delete           Bound    default/disk-ssd-web-restore-0   alicloud-disk-ssd            114s
```

## Troubleshooting
