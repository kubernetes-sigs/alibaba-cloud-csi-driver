
## Lvm CSI Plugin


An LVM CSI plugin is available to help simplify storage management.
You can create a pv with csi configuration, and the pvc, pod defines as usual.

Lvm is host based, and not a high availability solution for your application. Use lvm should accept effect for host down.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Alpha

## Compiling and Package
lvmplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-lvm.sh
```

## Usage

### Prerequisite
Same as diskplugin.csi.alibabacloud.com;


### Step 1: Create CSI Attacher
```
# kubectl create -f ./deploy/lvm/lvm-attacher.yaml
```

### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/lvm/lvm-plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/lvmplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 3: Create storageclass
```
# kubectl create -f ./deploy/lvm/storageclass.yaml
```

> vgName: define the volume group name of the storageclass;
>
> lvmSize: define the lvm size, this should be removed next release;

### Step 4: Create nginx deploy with csi
```
# kubectl create -f ./deploy/lvm/deploy.yaml
```

### Step 5: Check status of PVC/PV
```
# kubectl get pvc | grep lvm
lvm-pvc     Bound    pvc-0762fd86-5156-11e9-983f-00163e0b8d64   2Gi        RWO            csi-lvm        13h

# kubectl get pv | grep lvm
pvc-0762fd86-5156-11e9-983f-00163e0b8d64   2Gi        RWO            Delete           Bound    default/lvm-pvc     csi-lvm                 13h
```

### Step 6: check status of pod
check directory in pod

```
# kubectl get pod | grep deployment-lvm
deployment-lvm-64657b57b6-wdnwj   1/1     Running   0          12h

# kubectl exec -ti deployment-lvm-64657b57b6-wdnwj sh
# df -h | grep data
/dev/mapper/volumegroup1-pvc--0762fd86--5156--11e9--983f--00163e0b8d64  2.9G  9.0M  2.8G   1% /data

```
Check Directory in host:

```
# kubectl describe pod deployment-lvm-64657b57b6-wdnwj | grep Node:
Node:               izbp1cs36nqp29umu67zslz/192.168.1.11

# ifconfig | grep 192.168.1.11
        inet 192.168.1.11  netmask 255.255.255.0  broadcast 192.168.1.255

# lvdisplay  | grep pvc
  LV Path                /dev/volumegroup1/pvc-0762fd86-5156-11e9-983f-00163e0b8d64
  LV Name                pvc-0762fd86-5156-11e9-983f-00163e0b8d64
  
# mount | grep volumegroup
/dev/mapper/volumegroup1-pvc--0762fd86--5156--11e9--983f--00163e0b8d64 on /var/lib/kubelet/pods/f48f3715-515d-11e9-983f-00163e0b8d64/volumes/kubernetes.io~csi/pvc-0762fd86-5156-11e9-983f-00163e0b8d64/mount type ext4 (rw,relatime,data=ordered)

```
