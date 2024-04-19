
## NAS CSI Plugin

An NAS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements
* Service Accounts with required RBAC permissions

## Compiling and Package
nasplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-nas.sh
```

## Usage

### Prerequisite

Same as diskplugin.csi.alibabacloud.com;

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/nasplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 1: Create a statically provisioned PV.
```
# kubectl apply -f ./examples/nas/pv.yaml
```

### Step 2: Create a statically provisioned PVC.
```
# kubectl create -f ./deploy/nas/pvc.yaml
```
### Step 3: Deploy an application named nas-static and associate the PVC with the application.

```
# kubectl create -f ./examples/nas/deploy.yaml
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
    Driver:        nasplugin.csi.alibabacloud.com
    VolumeHandle:  data-id
    ReadOnly:      false
Events:            <none>
```

### Step 5: Run the following command to query the pods that run the application:
```
# kubectl get pod
```
```
NAME                          READY   STATUS    RESTARTS   AGE
deployment-nas-*****-n****    1/1     Running   0          32s
deployment-nas-*****-4****    1/1     Running   0          32s
```
## features

**Nas Dynamic:** [dynamic volume](./nas-dynamic.md)