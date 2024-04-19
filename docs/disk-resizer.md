# Alibaba Cloud Disk Resizer Plugin

## Overview

Aliyun EBS support to expand size of disk, this feature will show how to resize ebs in kubernetes platform.

*The CSI-Resizer is Alpha feature in k8s 1.14, Beta feature in k8s 1.16+*

## Requirements

* CSI resizer external runner (registry.cn-hangzhou.aliyuncs.com/acs/csi-resizer).
* Disk resizer plugin depends on csi-plugin (registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin).
* Service Accounts with required RBAC permissions.
* Feature Gate Enable: ExpandCSIVolumes=true(kube-controller, kubelet), this is Beta feature in kubernetes 1.16+;

### Feature Status
Beta

## Compiling and Package
Csi-resizer can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-all.sh
```

## Deploy

### 1. Requirements
Kubernetes cluster, api-server, kubelet configuration, please refer to [disk-plugin](./README-disk.md)

The resizer runner is working with CSI Plugin, so you should deploy the base plugin first. Please refer to [disk-plugin](./README-disk.md)

The API Authority:

The process of expand disk need to call aliyun ecs api, and your AK should have the ability to do it.

If you use STS in the csi plugin, the RAM should have the Authority of ResizeDisk.

### 2. Deploy Resizer and StorageClass

You can use below command to deploy csi resizer.

```
helm install csi-driver ./deploy/chart --values deploy/chart/values-ecs.yaml --namespace kube-system
```

## Usage

### 1. Create pvc with Disk
Create pvc, using disk volume.

```
# kubectl get pvc
NAME           STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS           AGE
pvc-disk-new   Bound    pvc-8db30f1a-ad23-11e9-ae51-00163e105050   20Gi       RWO            alicloud-disk-expand   7s

# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                  STORAGECLASS           REASON   AGE
pvc-8db30f1a-ad23-11e9-ae51-00163e105050   20Gi       RWO            Retain           Bound    default/pvc-disk-new   alicloud-disk-expand            15s
```


### 2. Create pod with disk
```
# kubectl get pod
NAME                              READY   STATUS    RESTARTS   AGE
dynamic-create-6d5dc9bb7d-lvhgz   1/1     Running   0          23s

// current, disk size is 20G；
# kubectl exec dynamic-create-6d5dc9bb7d-lvhgz df | grep data
/dev/vdd        20511312    45080  20449848   1% /data

```
### 3. Expand pvc from 20G to 30G:

```
// expand pvc from 20G to 30G
# kubectl patch pvc pvc-disk-new -p '{"spec":{"resources":{"requests":{"storage":"30Gi"}}}}'

// phase1: resizer finish disk expand, pvc is in FileSystemResizePending;
# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                  STORAGECLASS           REASON   AGE
pvc-8db30f1a-ad23-11e9-ae51-00163e105050   30Gi       RWO            Retain           Bound    default/pvc-disk-new   alicloud-disk-expand            5m1s

# kubectl describe pvc pvc-disk-new
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

```
// phase2: restart Pod, expand filesystem;
# kubectl delete pod dynamic-create-6d5dc9bb7d-lvhgz
pod "dynamic-create-6d5dc9bb7d-lvhgz" deleted

# kubectl get pvc
NAME           STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS           AGE
pvc-disk-new   Bound    pvc-8db30f1a-ad23-11e9-ae51-00163e105050   30Gi       RWO            alicloud-disk-expand   6m13s
# kubectl exec -ti dynamic-create-6d5dc9bb7d-5gzq2 df | grep data
/dev/vdd        30832548    45036  30771128   1% /data

```

## Troubleshooting
