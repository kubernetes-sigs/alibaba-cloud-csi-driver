
## NAS CSI Plugin

An NAS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Alpha

## Compiling and Package
nasplugin.csi.aliyun.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-nas.sh
```


## Usage

### Prerequisite

Same as diskplugin.csi.aliyun.com;


### Step 1: Create CSI Attacher
```
# kubectl create -f ./deploy/nas/nas-attacher.yaml
```

### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/nas/nas-plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/nasplugin.csi.aliyun.com.log);

> "stdout": default option, logs will be printed to stdout, can be printed by docker logs or kubectl logs.

### Step 3: Create nginx deploy with csi
```
# kubectl create -f ./deploy/nas/deploy.yaml
```

### Step 4: Check status of PVC/PV
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
    Driver:        nasplugin.csi.aliyun.com
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```
