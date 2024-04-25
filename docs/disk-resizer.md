# Alibaba Cloud Disk Resizer Plugin

## Overview

Aliyun EBS support to expand size of disk, this feature will show how to resize ebs in kubernetes platform.

*The CSI-Resizer is Alpha feature in k8s 1.14, Beta feature in k8s 1.16+, GA in k8s 1.24+*

## Requirements

Working Disk CSI driver,
Please refer to the [disk driver doc](./disk.md) for details.

## Usage

### 1. Create pvc with Disk
Create pvc, using disk volume.

```shell
kubectl get pvc
```
Expected output:
```
NAME           STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS           AGE
pvc-disk-new   Bound    pvc-8db30f1a-ad23-11e9-ae51-00163e105050   20Gi       RWO            alicloud-disk-expand   7s
```
```shell
kubectl get pv
```
Expected output:
```
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                  STORAGECLASS           REASON   AGE
pvc-8db30f1a-ad23-11e9-ae51-00163e105050   20Gi       RWO            Retain           Bound    default/pvc-disk-new   alicloud-disk-expand            15s
```


### 2. Create pod with disk
```shell
kubectl get pod
```
```
NAME                              READY   STATUS    RESTARTS   AGE
dynamic-create-6d5dc9bb7d-lvhgz   1/1     Running   0          23s
```

current, disk size is 20G；

```shell
kubectl exec dynamic-create-6d5dc9bb7d-lvhgz df | grep data
```
Expected output:
```
/dev/vdd        20511312    45080  20449848   1% /data
```


### 3. Expand pvc from 20G to 30G:

Expand pvc from 20G to 30G

```shell
kubectl patch pvc pvc-disk-new -p '{"spec":{"resources":{"requests":{"storage":"30Gi"}}}}'
```

phase1: resizer finish disk expand, pvc is in FileSystemResizePending;
```shell
kubectl get pv
```
```
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                  STORAGECLASS           REASON   AGE
pvc-8db30f1a-ad23-11e9-ae51-00163e105050   30Gi       RWO            Retain           Bound    default/pvc-disk-new   alicloud-disk-expand            5m1s
```
```
Name:          pvc-disk-new
Namespace:     default
StorageClass:  alicloud-disk-expand
Status:        Bound
Volume:        pvc-8db30f1a-ad23-11e9-ae51-00163e105050
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: diskplugin.csi.alibabacloud.com
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      20Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Conditions:
  Type                      Status  LastProbeTime                     LastTransitionTime                Reason  Message
  ----                      ------  -----------------                 ------------------                ------  -------
  FileSystemResizePending   True    Mon, 01 Jan 0001 00:00:00 +0000   Tue, 23 Jul 2019 16:31:25 +0800           Waiting for user to (re-)start a pod to finish file system resize of volume on node.
```

### 4. Expand FileSystem with restart pod:

phase2: restart Pod, expand filesystem;

```shell
kubectl delete pod dynamic-create-6d5dc9bb7d-lvhgz
```
Check pvc status
```shell
kubectl get pvc
```
Expected output:
```
NAME           STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS           AGE
pvc-disk-new   Bound    pvc-8db30f1a-ad23-11e9-ae51-00163e105050   30Gi       RWO            alicloud-disk-expand   6m13s
```
Run `df` to check the disk size
```shell
kubectl exec -ti dynamic-create-6d5dc9bb7d-5gzq2 df | grep data
```
Expected output:
```
/dev/vdd        30832548    45036  30771128   1% /data
```
## Troubleshooting
