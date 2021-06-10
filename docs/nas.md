
## NAS CSI Plugin

An NAS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Beta

## Compiling and Package
nasplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-nas.sh
```


## Usage

### Prerequisite

Same as diskplugin.csi.alibabacloud.com;

### Step 1: Create RBAC resource

```
# kubectl create -f ./deploy/rbac.yaml
```

### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/nas/nas-plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/nasplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.


If your kubelet RootDir is not /var/lib/kubelet, the deploy template(nas-plugin.yaml) should be changed:

> Replace all /var/lib/kubelet into /rootdir/kubelet for the nas-plugin.yaml file.

### Step 3: Create CSI Plugin

```
# kubectl create -f ./deploy/nas/nas-provisioner.yaml
```

### Step 4: Create nginx deploy with csi volume
```
# kubectl create -f ./examples/nas/deploy.yaml
```

### Step 3: Check status of PVC/PV
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
    Driver:        nasplugin.csi.alibabacloud.com
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```

## features

**Nas Dynamic:** [dynamic volume](./nas-dynamic.md)