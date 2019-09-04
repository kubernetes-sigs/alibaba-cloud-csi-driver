
## OSS CSI Plugin


An OSS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.


## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Alpha

## Compiling and Package
ossplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-oss.sh
```

## Usage

### Prerequisite
Same as diskplugin.csi.alibabacloud.com;


### Step 1: Create CSI Plugin
```
# kubectl create -f ./deploy/oss/oss-plugin.yaml
```


> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/ossplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 2: Create nginx deploy with csi
```
# kubectl create -f ./examples/oss/deploy.yaml
```

### Step 3: Check status of PVC/PV
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
    Driver:        ossplugin.csi.alibabacloud.com
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```
