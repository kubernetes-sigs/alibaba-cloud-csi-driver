# Alibaba Cloud Kubernetes CSI Plugin

[![Build Status](https://travis-ci.org/AliyunContainerService/csi-plugin.svg?branch=master)](https://travis-ci.org/AliyunContainerService/csi-plugin)
[![CircleCI](https://circleci.com/gh/AliyunContainerService/csi-plugin.svg?style=svg)](https://circleci.com/gh/AliyunContainerService/csi-plugin)
[![Go Report Card](https://goreportcard.com/badge/github.com/AliyunContainerService/csi-plugin)](https://goreportcard.com/report/github.com/AliyunContainerService/csi-plugin)

## Overview

Alibaba cloud CSI plugins implement an interface between CSI enabled Container
Orchestrator and Alibaba Cloud Storage. It allows dynamically provision Disk
volumes and attach it to workloads.
Current implementation of CSI plugins was tested in Kubernetes environment (requires Kubernetes 1.10+).

Current Support: ***Alibaba cloud Disk, OSS, NAS***;

## Disk CSI Plugin

An Disk CSI plugin is available to help simplify storage management.
Once user creates PVC with the reference to a Disk storage class, disk and
corresponding PV object gets dynamically created and becomes ready to be used by
workloads.

### Configuration Requirements

* Secret object with the authentication key for Disk
* StorageClass with diskplugin (default csi-diskplugin name) as a provisioner name
  and information about disk(zoneId, regionId, type)
* Service Accounts with required RBAC permissions   

### Feature Status
Alpha

### Compiling and Package
csi-diskplugin can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-disk.sh
```

### Demo


[![](demo.png)](http://cloud.video.taobao.com/play/u/1962692024/p/1/e/6/t/1/50224108448.mp4)


### Usage

#### Prerequisite

##### Get Kubernetes cluster

You can create a Kubernetes Cluster on [Alibaba cloud Container Service](https://help.aliyun.com/product/25972.html?spm=a2c4g.750001.2.3.A7g9FZ)

##### Config Kubelet

Set --enable-controller-attach-detach=true for kubelet:

```
# vi /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

Kubernetes default to true if you use Kubeadm to create Cluster, Alibaba cloud set it to false for the Flexvolume feature;
If you want to use Flexvolume, set it to false again;


#### Step 1: Create CSI Attacher
```
# kubectl create -f ./deploy/disk/diskattacher.yaml
```

#### Step 2: Create CSI Provisioner
```
# kubectl create -f ./deploy/disk/diskprovisioner.yaml
```

#### Step 3: Create CSI Plugin
If the cluster not in STS mode, you need to config AK info to plugin; Set ACCESS_KEY_ID, ACCESS_KEY_SECRET to environment;

```
# kubectl create -f ./deploy/disk/diskplugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/csi-diskplugin.log);

> "stdout": default option, logs will be printed to stdout, can be printed by docker logs or kubectl logs.

#### Step 4: Create StorageClass
```
# kubectl create -f ./deploy/disk/storageclass.yaml
```
**Important:** storageclass.yaml, must be customized to match your environment: zoneId, zoneId;

#### Step 5: Check Status of CSI plugin
```
# kubectl get pods | grep csi
```

The following output should be displayed:

```
NAME                                 READY     STATUS    RESTARTS   AGE
csi-attacher-0                       1/1       Running   11         2d
csi-diskplugin-568pb                 2/2       Running   0          19h
csi-diskplugin-f4tsn                 2/2       Running   0          19h
csi-diskplugin-rq8tj                 2/2       Running   0          19h
csi-diskplugin-tc6rj                 2/2       Running   0          19h
csi-provisioner-0                    1/1       Running   0          2d
```

#### Step 7: Create PVC & Deployments
```
# kubectl create -f ./deploy/disk/deploy.yaml
```

#### Step 8: Check status of PVC/PV
```
# kubectl get pvc
NAME       STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS           AGE
disk-pvc   Bound     pvc-64b3d1bd-96c0-11e8-89b1-00163e0c412f   25Gi       RWO            csi-disk               36m
```

```
# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS        CLAIM                 STORAGECLASS               REASON    AGE
pvc-64b3d1bd-96c0-11e8-89b1-00163e0c412f   25Gi       RWO            Delete           Terminating   default/disk-pvc      csi-disk                             35m
```

```
# kubectl describe pv pvc-64b3d1bd-96c0-11e8-89b1-00163e0c412f
Name:            pvc-64b3d1bd-96c0-11e8-89b1-00163e0c412f
Labels:          <none>
Annotations:     pv.kubernetes.io/provisioned-by=csi-diskplugin
Finalizers:      [kubernetes.io/pv-protection external-attacher/csi-diskplugin]
StorageClass:    csi-disk
Status:          Terminating (lasts 2m)
Claim:           default/disk-pvc
Reclaim Policy:  Delete
Access Modes:    RWO
Capacity:        25Gi
Node Affinity:   <none>
Message:
Source:
    Type:          CSI (a Container Storage Interface (CSI) volume source)
    Driver:        csi-diskplugin
    VolumeHandle:  d-2ze47lce65lv5g7zsb4y
    ReadOnly:      false
Events:            <none>
```

#### Step 9: Check status of Deployment
```
# kubectl get pod
NAME                                 READY     STATUS    RESTARTS   AGE
nginx-deployment1-5879d9db88-49n8m   1/1       Running   0          37m
```

## NAS CSI Plugin

An NAS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

### Configuration Requirements

* Service Accounts with required RBAC permissions

### Feature Status
Alpha

### Compiling and Package
csi-nasplugin can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-nas.sh
```


### Usage

#### Prerequisite

Same as csi-diskplugin;


#### Step 1: Create CSI Attacher
```
# kubectl create -f ./deploy/nas/nasattacher.yaml
```

#### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/nas/nasplugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/csi-nasplugin.log);

> "stdout": default option, logs will be printed to stdout, can be printed by docker logs or kubectl logs.

#### Step 3: Create nginx deploy with csi
```
# kubectl create -f ./deploy/nas/deploy.yaml
```

#### Step 4: Check status of PVC/PV
```
# kubectl get pvc
NAME      STATUS    VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
nas-pvc   Bound     nas-csi-pv   5Gi        RWO                           3m
```

```
# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS        CLAIM              STORAGECLASS   REASON    AGE
nas-csi-pv                                 5Gi        RWO            Retain           Bound         default/nas-pvc                             3m
```

```
# kubectl describe pv nas-csi-pv
Name:            nas-csi-pv
Labels:          <none>
Annotations:     pv.kubernetes.io/bound-by-controller=yes
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:
Status:          Bound
Claim:           default/nas-pvc
Reclaim Policy:  Retain
Access Modes:    RWO
Capacity:        5Gi
Node Affinity:   <none>
Message:
Source:
    Type:          CSI (a Container Storage Interface (CSI) volume source)
    Driver:        csi-nasplugin
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```


## OSS CSI Plugin


An OSS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

### Configuration Requirements

* Service Accounts with required RBAC permissions

### Feature Status
Alpha

### Compiling and Package
csi-ossplugin can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-oss.sh
```


### Usage

#### Prerequisite

Same as csi-diskplugin;


#### Step 1: Create CSI Attacher
```
# kubectl create -f ./deploy/oss/ossattacher.yaml
```

#### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/oss/ossplugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/csi-ossplugin.log);

> "stdout": default option, logs will be printed to stdout, can be printed by docker logs or kubectl logs.

#### Step 3: Create nginx deploy with csi
```
# kubectl create -f ./deploy/oss/deploy.yaml
```

#### Step 4: Check status of PVC/PV
```
# kubectl get pvc
NAME      STATUS    VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
oss-pvc   Bound     oss-csi-pv   5Gi        RWO                           1m
```

```
# kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS        CLAIM              STORAGECLASS   REASON    AGE
oss-csi-pv                                 5Gi        RWO            Retain           Bound         default/oss-pvc                             1m
```

```
# kubectl describe pv oss-csi-pv
Name:            oss-csi-pv
Labels:          <none>
Annotations:     pv.kubernetes.io/bound-by-controller=yes
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:
Status:          Bound
Claim:           default/oss-pvc
Reclaim Policy:  Retain
Access Modes:    RWO
Capacity:        5Gi
Node Affinity:   <none>
Message:
Source:
    Type:          CSI (a Container Storage Interface (CSI) volume source)
    Driver:        csi-ossplugin
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```



## Troubleshooting

Please submit an issue at: [Issues](https://github.com/AliyunContainerService/csi-plugin/issues)
