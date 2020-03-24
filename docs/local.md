
## Local CSI Plugin


An Local CSI plugin is available to help simplify local storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

Local volume is host based, and not a high availability solution for your application. Use local volume should accept effect for host down.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Beta

## Compiling and Package
localplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-local.sh
```

## Usage

### Prerequisite
Same as diskplugin.csi.alibabacloud.com;


### Step 1: Create CSI Provisioner
```
# kubectl create -f ./deploy/local/provisioner.yaml
```

### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/local/plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/localplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 3: Create storageclass
```
# kubectl create -f ./deploy/local/storageclass.yaml
```

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: csi-lvm-striping
provisioner: localplugin.csi.alibabacloud.com
parameters:
    volumeType: "LVM"
    vgName: volumegroup1
    fsType: ext4
    pvType: clouddisk
    lvmType: linear
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
```
Usage:

> volumeType: define the local volume type, current support LVM;
>
> vgName: define the volume group name of the storageclass;
>
> fsType: default as ext4, define the lvm file system type, support ext4, ext3, xfs;
>
> pvType: optional, default as clouddisk. Define the used physics disk type, support clouddisk, localdisk;
>
> ----> clouddisk: default, not support auto create vg, and vg should pre-created by hands.
>
> ----> localdisk: support auto create vg using localdisks, and localdisks should not be inited.
>
> lvmType: optional, default as linear; Define the expect lvm type, supporting linear, striping;
>
> ----> linear: default, create linear lvm volume.
>
> ----> striping: only support pvType(localdisk) now, and the striped number is the localdisk number
>
> nodeAffinity: optional, default as true. Decide add nodeAffinity in PV or not.
> ----> true: default, create PV with nodeAffinity configuration;
> ----> false: create PV without nodeAffinity configuration, pod can be scheduled to any node
>
> volumeBindingMode: support Immediate/WaitForFirstConsumer
> ----> Immediate: means volume will be provisioned while pvc is created, in this config nodeAffinity will be available;
> ----> WaitForFirstConsumer: means volume will not be created until the related pod created; In the config, nodeAffinity will be unavailable;

### Step 4: Create nginx deploy with local
```
# kubectl create -f ./examples/local/pvc.yaml
# kubectl create -f ./examples/local/deploy.yaml
```

### Step 5: Check status of PVC/PV
```
# kubectl get pvc | grep local
local-pvc     Bound    local-0762fd86-5156-11e9-983f-00163e0b8d64   2Gi        RWO            csi-local        13h

# kubectl get pv | grep local
local-0762fd86-5156-11e9-983f-00163e0b8d64   2Gi        RWO            Delete           Bound    default/local-pvc     csi-local                 13h
```

### Step 6: check status of pod
check directory in pod

```
# kubectl get pod | grep deployment-local
deployment-local-64657b57b6-wdnwj   1/1     Running   0          12h

# kubectl exec -ti deployment-local-64657b57b6-wdnwj sh
# df -h | grep data
/dev/mapper/volumegroup1-local--0762fd86--5156--11e9--983f--00163e0b8d64  2.9G  9.0M  2.8G   1% /data

```
Check Directory in host:

```
# kubectl describe pod deployment-local-64657b57b6-wdnwj | grep Node:
Node:               izbp1cs36nqp29umu67zslz/192.168.1.11

# ifconfig | grep 192.168.1.11
        inet 192.168.1.11  netmask 255.255.255.0  broadcast 192.168.1.255

# lvdisplay  | grep pvc
  LV Path                /dev/volumegroup1/local-0762fd86-5156-11e9-983f-00163e0b8d64
  LV Name                local-0762fd86-5156-11e9-983f-00163e0b8d64
  
# mount | grep volumegroup
/dev/mapper/volumegroup1-local--0762fd86--5156--11e9--983f--00163e0b8d64 on /var/lib/kubelet/pods/f48f3715-515d-11e9-983f-00163e0b8d64/volumes/kubernetes.io~csi/local-0762fd86-5156-11e9-983f-00163e0b8d64/mount type ext4 (rw,relatime,data=ordered)

```
