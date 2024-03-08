
## CPFS CSI Plugin

**Deprecated**

An CPFS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Beta

## Compiling and Package
cpfsplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-cpfs.sh
```


## Usage

### Prerequisite

Same as cpfsplugin.csi.alibabacloud.com;

Current, csi-cpfsplugin only support centos7 os and linux kernel versions as belows:

> 3.10.0-1062.9.1
>
> 3.10.0-957.21.3
>
> 3.10.0-957.5.1


### Step 1: Create CSI Plugin
```
# kubectl create -f ./deploy/cpfs/cpfs-plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/cpfsplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 2: Create nginx deploy with csi
```
# kubectl create -f ./examples/cpfs/deploy.yaml
```

### Step 3: Check status of PVC/PV
```
# kubectl get pvc
NAME      STATUS    VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
cpfs-pvc   Bound     cpfs-csi-pv   5Gi        RWO                           3m
```

```
# kubectl get pv
NAME          CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
cpfs-csi-pv   5Gi        RWO            Retain           Available                                   2s
```

```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: cpfs-csi-pv
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 5Gi
  csi:
    driver: cpfsplugin.csi.alibabacloud.com
    volumeAttributes:
      fileSystem: 0237cc41
      server: cpfs-0237cc41-**.cn-shenzhen.cpfs.nas.aliyuncs.com@tcp:cpfs-0237cc41-**.cn-shenzhen.cpfs.nas.aliyuncs.com@tcp
      subPath: /shenzhen
    volumeHandle: cpfs-csi-pv
  persistentVolumeReclaimPolicy: Retain
  volumeMode: Filesystem
```
