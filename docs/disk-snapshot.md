# Alibaba Cloud Disk Snapshot Plugin

## Overview

An Disk snapshot plugin is available to help simplify storage management.
Once user creates VolumeSnapshot with the reference to a snapshot class, snapshot and
corresponding VolumeSnapshotContent object gets dynamically created and becomes ready to be used by
workloads.

The csi-disksnapshot is Beta feature in k8s 1.17.

Disk Snapshot Plugin will create 3 crds as below.

```
volumesnapshotclasses.snapshot.storage.k8s.io:  define details to create volumesnapshotcontents, like: storageclass;
volumesnapshotcontents.snapshot.storage.k8s.io: define one disk snapshot with the backend, like: pv;
volumesnapshots.snapshot.storage.k8s.io:        the claim of one snapshot, like: pvc;
```

## Requirements

* CSI Snapshot controller (registry.cn-hangzhou.aliyuncs.com/acs/snapshot-controller:v2.0.1)
* CSI Snapshot external runner (registry.cn-hangzhou.aliyuncs.com/acs/csi-snapshotter:v2.1.1).
* Disk Snapshot plugin depends on csi-plugin (registry.cn-hangzhou.aliyuncs.com/acs/csi-diskplugin).
* Secret object with the authentication key for Disk
* Service Accounts with required RBAC permissions

### Feature Status
Beta

## Compiling and Package
disk-snapshot can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-disk.sh
```

## Deploy

### 1. Requirements
Kubernetes cluster, api-server, kubelet configuration, please refer to [disk-plugin](./disk.md)


### 2. Create Snapshot CRD

```
# kubectl create -f ./deploy/disk/snapshot/crd.yaml
```

### 3. Deploy Snapshot Components

```
# kubectl create -f ./deploy/disk/snapshot/csi-snapshotter.yaml
```


## Usage

### 1. Create SnapshotClass
```
# kubectl create -f ./examples/disk/snapshot/snapshotclass.yaml

# kubectl get volumesnapshotclasses.snapshot.storage.k8s.io
NAME                AGE
default-snapclass   15s
```

### 2. Create pvc with Disk
Create pvc, using disk volume.

```
# kubectl get pvc
NAME       STATUS   VOLUME                   CAPACITY   ACCESS MODES   STORAGECLASS        AGE
pvc-disk   Bound    d-wz9hhnhxu66zs6r6yyyq   20Gi       RWO            alicloud-disk-ssd   25s
```

### 3. Create Snapshot with pvc
```
# kubectl create -f ./examples/disk/snapshot/snapshot.yaml
volumesnapshot.snapshot.storage.k8s.io/new-snapshot-demo created

# kubectl get volumesnapshots.snapshot.storage.k8s.io
NAME                AGE
new-snapshot-demo   2m

# kubectl get volumesnapshotcontents.snapshot.storage.k8s.io
NAME                                               AGE
snapcontent-715d2ae7-0096-4eb8-bfc7-34ffaa624043   2m
```

> snapshot.yaml define the data source pvc name, and snapshotClass name.

> You can check disk snapshot in alibaba cloud ecs console.

## Troubleshooting
