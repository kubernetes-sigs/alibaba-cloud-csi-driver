
## NAS CSI Plugin

An NAS CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

## Configuration Requirements
* Service Accounts with required RBAC permissions

## Compiling and Package
nasplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```shell
cd build && sh build-nas.sh
```

## Usage

### Prerequisite

Same as diskplugin.csi.alibabacloud.com;

### Step 1: Create a statically provisioned PV.
```shell
kubectl apply -f ./examples/nas/pv.yaml
```

### Step 2: Create a statically provisioned PVC.
```shell
kubectl create -f ./deploy/nas/pvc.yaml
```
### Step 3: Deploy an application named nas-static and associate the PVC with the application.

```shell
kubectl create -f ./examples/nas/deploy.yaml
```

### Step 4: Check status of PVC/PV
Check status of PVC
```shell
kubectl get pvc
```
Expected output:
```
NAME      STATUS    VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
nas-pvc   Bound     nas-csi-pv   5Gi        RWO                           3m
```
Check status of PV
```shell
kubectl get pv
```
Expected output:
```
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS        CLAIM              STORAGECLASS   REASON    AGE
nas-csi-pv                                 5Gi        RWO            Retain           Bound         default/nas-pvc                             3m
```
```shell
kubectl describe pv nas-csi-pv
```
Expected output:
```
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
```shell
kubectl get pod
```
```
NAME                          READY   STATUS    RESTARTS   AGE
deployment-nas-*****-n****    1/1     Running   0          32s
deployment-nas-*****-4****    1/1     Running   0          32s
```
## features

**Nas Dynamic:** [dynamic volume](./nas-dynamic.md)