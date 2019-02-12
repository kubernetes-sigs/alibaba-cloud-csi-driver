# Alibaba Cloud Disk Snapshot Plugin

## Overview

An Disk snapshot plugin is available to help simplify storage management.
Once user creates VolumeSnapshot with the reference to a snapshot class, snapshot and
corresponding VolumeSnapshotContent object gets dynamically created and becomes ready to be used by
workloads.

The csi-disksnapshot is Alpha feature in k8s 1.13.

Disk Snapshot Plugin will create 3 crds as below.

```
volumesnapshotclasses.snapshot.storage.k8s.io:  define details to create volumesnapshotcontents, like: storageclass;
volumesnapshotcontents.snapshot.storage.k8s.io: define one disk snapshot with the backend, like: pv;
volumesnapshots.snapshot.storage.k8s.io:        the claim of one snapshot, like: pvc;
```

## Requirements

* CSI Snapshot external runner (registry.cn-hangzhou.aliyuncs.com/plugins/csi-snapshotter:v1.0.0).
* Disk Snapshot plugin depends on csi-plugin (registry.cn-hangzhou.aliyuncs.com/plugins/csi-diskplugin:v1.0.0).
* Secret object with the authentication key for Disk
* Service Accounts with required RBAC permissions

### Feature Status
Alpha

## Compiling and Package
disk-snapshot can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-disk.sh
```

## Deploy

### 1. Requirements
Kubernetes cluster, api-server, kubelet configuration, please refer to [disk-plugin](./README-disk.md)


### 2. Deploy Snapshot Runner

```
# kubectl create -f ./deploy/snapshot/csi-external-snapshot.yaml
```

### 3. Create SnapshotClass
```
# kubectl create -f ./deploy/snapshot/snapshotclass.yaml

# kubectl get volumesnapshotclasses.snapshot.storage.k8s.io
NAME                AGE
default-snapclass   15s
```

3 CRDs will be created：

```
# kubectl get crd | grep snap
volumesnapshotclasses.snapshot.storage.k8s.io    2019-02-01T06:04:49Z
volumesnapshotcontents.snapshot.storage.k8s.io   2019-02-01T06:04:49Z
volumesnapshots.snapshot.storage.k8s.io          2019-02-01T06:04:49Z
```

## Usage

### 1. Create pvc with Disk
Create pvc, using disk volume.

```
# kubectl get pvc
NAME       STATUS   VOLUME                   CAPACITY   ACCESS MODES   STORAGECLASS        AGE
pvc-disk   Bound    d-wz9hhnhxu66zs6r6yyyq   20Gi       RWO            alicloud-disk-ssd   25s

```


### 2. Create Snapshot with pvc
```
# kubectl create -f ./deploy/snapshot/snapshot.yaml
volumesnapshot.snapshot.storage.k8s.io/snapshot-test created

# kubectl get volumesnapshots.snapshot.storage.k8s.io
NAME            AGE
snapshot-test   2m

# kubectl get volumesnapshotcontents.snapshot.storage.k8s.io
NAME                                               AGE
snapcontent-77a75b63-1987-11e9-a520-00163e024341   2m
```

> snapshot.yaml define the data source pvc name, and snapshotClass name.

> You can check disk snapshot in aliyun ecs console.

## Troubleshooting
