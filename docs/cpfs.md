
## CPFS CSI Plugin

An CPFS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Alpha

## Compiling and Package
cpfsplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-cpfs.sh
```


## Usage

### Prerequisite

Same as cpfsplugin.csi.alibabacloud.com;


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
# kubectl describe pv cpfs-csi-pv
Name:            cpfs-csi-pv
Labels:          <none>
Annotations:     <none>
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:
Status:          Available
Claim:
Reclaim Policy:  Retain
Access Modes:    RWO
VolumeMode:      Filesystem
Capacity:        5Gi
Node Affinity:   <none>
Message:
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            cpfsplugin.csi.alibabacloud.com
    VolumeHandle:      cpfs-csi-pv
    ReadOnly:          false
    VolumeAttributes:      fileSystem=kfs
                           server=192.168.0.214@tcp
                           subPath=/k8s
Events:                <none>
```
