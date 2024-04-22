## OSS CSI Plugin

An OSS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Compiling and Package
ossplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```shell
cd build && sh build-oss.sh
```

## Usage

### Prerequisite

Same as diskplugin.csi.alibabacloud.com;

### Features

> Oss Plugin support to use Secret for Access Authorization;
>
> Oss Plugin support to mount remote subpath under oss bucket;
>
> Oss Plugin support to upgrade online;

### Step 1: Create pv/pvc/deploy with csi

```shell
kubectl create -f ./examples/oss/pv.yaml
kubectl create -f ./examples/oss/pvc.yaml
kubectl create -f ./examples/oss/deploy.yaml
```

### Step 3: Check status of PVC/PV

Check pvc status
```shell
kubectl get pvc | grep oss
```
Expected output:
```
oss-pvc    Bound    oss-csi-pv      5Gi        RWO        5m18s
```
Check pv status
```shell
kubectl get pv | grep oss
```
Expected output:
```
oss-csi-pv      5Gi        RWO        Retain       Bound    default/oss-pvc                                 5m21s
```
```shell
kubectl get pv oss-csi-pv -o yaml
```
Expected output:
```
apiVersion: v1
kind: PersistentVolume
metadata:
  labels:
    alicloud-pvname: oss-csi-pv
  name: oss-csi-pv
spec:
  accessModes:
  - ReadWriteOnce
  capacity:
    storage: 5Gi
  csi:
    driver: ossplugin.csi.alibabacloud.com
    volumeAttributes:
      akId: *
      akSecret: *
      bucket: oss-docker
      otherOpts: -o max_stat_cache_size=0 -o allow_other
      path: /docker
      url: oss-cn-shenzhen-internal.aliyuncs.com
    volumeHandle: oss-csi-pv
  persistentVolumeReclaimPolicy: Retain
```
