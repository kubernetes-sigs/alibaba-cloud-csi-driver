# Alibaba Cloud Lvm Resizer Plugin

## Overview

*The CSI-Resizer is Alpha feature in k8s 1.14.* and beta in 1.16.

## Requirements

* CSI resizer external runner (registry.cn-hangzhou.aliyuncs.com/acs/csi-resizer:v0.1.0).
* Lvm resizer plugin depends on csi-lvm-plugin (registry.cn-hangzhou.aliyuncs.com/acs/csi-lvm-plugin).
* Service Accounts with required RBAC permissions.
* Feature Gate Enable: ExpandCSIVolumes=true(kube-controller, kubelet), this is Alpha feature in kubernetes 1.14;


### Feature Status
Alpha

## Compiling and Package
Csi-lvm-resizer can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-lvm.sh
```

## Deploy

### 1. Requirements
Kubernetes cluster, api-server, kubelet configuration, please refer to [lvm-plugin](./lvm.md)

The resizer runner is working with CSI Plugin, so you should deploy the base plugin first. Please refer to [lvm-plugin](./lvm.md)

### 2. Deploy Resizer Runner

You can use below command to deploy csi resizer.

```
# kubectl create -f ./deploy/lvm/resizer/csi-resizer.yaml
```

### 3. Create storageclass
```
# kubectl create -f ./examples/lvm/resizer/storageclass.yaml
```

## Usage

### 1. dependence

Create vg in all nodes, which expect to create lvm volume. The vg name should be configured in storageclass.

StorageClass configure allowVolumeExpansion as true;

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: csi-lvm
provisioner: lvmplugin.csi.alibabacloud.com
parameters:
    vgName: vgtest
reclaimPolicy: Delete
allowVolumeExpansion: true
```

### 2. Create pvc with Lvm
Create pvc, provision lvm volume.

*# examples/lvm/pvc.yaml*

```
# kubectl get pvc
NAME      STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
lvm-pvc   Bound    lvm-dcfd087a-ea93-11e9-a442-00163e07fb69   2Gi        RWO            csi-lvm        28s

# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM             STORAGECLASS   REASON   AGE
lvm-dcfd087a-ea93-11e9-a442-00163e07fb69   2Gi        RWO            Delete           Bound    default/lvm-pvc   csi-lvm                 21s
```


### 3. Create pod with lvm

*# examples/lvm/deploy.yaml*

```
# kubectl get pod
NAME                              READY   STATUS    RESTARTS   AGE
deployment-lvm-546996b647-k9cww   1/1     Running   0          51s

// lvm volume size is 2G;
# kubectl exec deployment-lvm-546996b647-k9cww df | grep data
/dev/mapper/vgtest-lvm--dcfd087a--ea93--11e9--a442--00163e07fb69   1998672    6144   1871288   1% /data
```

Check lvm volume in the node pod running:

```
# lvdisplay /dev/vgtest/lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  --- Logical volume ---
  LV Path                /dev/vgtest/lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  LV Name                lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  VG Name                vgtest
  LV UUID                eDXhM9-vKOp-jCOF-9GCR-l8nb-GvzW-Ej5YfE
  LV Write Access        read/write
  LV Creation host, time iZ8vb1wy4teeoa0ql4bjq3Z, 2019-10-09 20:54:17 +0800
  LV Status              available
  # open                 1
  LV Size                2.00 GiB
  Current LE             512
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           252:1
```

### 4. Expand pvc from 2G to 3G:

```
// expand pvc from 2G to 3G
# kubectl patch pvc lvm-pvc -p '{"spec":{"resources":{"requests":{"storage":"3Gi"}}}}'
persistentvolumeclaim/lvm-pvc patched

# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM             STORAGECLASS   REASON   AGE
lvm-dcfd087a-ea93-11e9-a442-00163e07fb69   3Gi        RWO            Delete           Bound    default/lvm-pvc   csi-lvm                 2m41s
```
### 5. Expand FileSystem with restart pod:

```
// phase2: restart Pod, expand filesystem;
# kubectl delete pod deployment-lvm-546996b647-k9cww
pod "deployment-lvm-546996b647-k9cww" deleted

# kubectl get pod
NAME                              READY   STATUS    RESTARTS   AGE
deployment-lvm-546996b647-t8z6v   1/1     Running   0          10m

// pvc expand to 3G;
# kubectl get pvc
NAME      STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
lvm-pvc   Bound    lvm-dcfd087a-ea93-11e9-a442-00163e07fb69   3Gi        RWO            csi-lvm        7m58s

// filesystem expand to 3G;
# kubectl exec deployment-lvm-546996b647-t8z6v df | grep data
/dev/mapper/vgtest-lvm--dcfd087a--ea93--11e9--a442--00163e07fb69   3030800    9216   2847916   1% /data
```

Check lvm volume in node pod running.

```
# lvdisplay /dev/vgtest/lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  --- Logical volume ---
  LV Path                /dev/vgtest/lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  LV Name                lvm-dcfd087a-ea93-11e9-a442-00163e07fb69
  VG Name                vgtest
  LV UUID                uaCe6k-3LmH-3JzB-0bjG-kZSX-yD0S-iDKyej
  LV Write Access        read/write
  LV Creation host, time iZ8vb1wy4teeoa0ql4bjq2Z, 2019-10-09 21:01:38 +0800
  LV Status              available
  # open                 1
  LV Size                3.00 GiB
  Current LE             768
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           252:1
```

## Troubleshooting
