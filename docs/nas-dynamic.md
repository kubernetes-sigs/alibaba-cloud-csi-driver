## Mount a dynamically provisioned NAS volume

### Prerequisite
Same as diskplugin.csi.alibabacloud.com;
## Mount a dynamically provisioned NAS volume in subpath mode

#### Step 1: Create a StorageClass.
The `subpath` type nas volume require you to create a nas fileSystem and mountTarget firstly

```shell
kubectl create -f ./examples/nas/dynamic/sc-subpath.yaml
```

Parameters:

> mountOptions: optional. specify the options for nfs mount.
>
> volumeAs: optional. Default as 'subpath'; specify the provision type: 'subpath' means the provision volume use subpath of exist nfs filesystem as target. And 'filesystem' means the provision volume use new created nfs filesystem as target.
>
> server: Necessary when volumeAs is subpath. Define the nfs server endpoints, format as: nfsserver1:/path1,nfsserver2:/path2, the volume server target will use the nfsserver/path by order.
>
> mode: Optional. Define the mount option in pv spec, which will be used after mount action.
>
> modeType: Optional. Default non-recursive. Define the Mode action behavior, recursive: chmod with -R and chanage all files mode under the mounted directory. non-recursive: chmod without -R and only chanage the directory mode.
>
> archiveOnDelete: Optional. decide how to process removal path, if reclaimPolicy defined as delete. If set 'true', the removal path will be archived and not removed really, and if set 'false', the removal path will be removed when pv is deleted.

### Step 2: Run the following command to create a PVC:
Create pvc
```shell
kubectl create -f ./examples/nas/dynamic/pvc-subpath.yaml
```
Check status of pvc
```shell
kubectl get pvc | grep nas
```
Expected output:
```
nas-csi           Bound    nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2    20Gi       RWX            alicloud-nas        6s
```
Check status of pv
```shell
kubectl get pv | grep nas
```
Expected output:
```
nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2    20Gi       RWX            Delete           Bound    default/nas-csi           alicloud-nas                 21s
```
Describe pv
```shell
kubectl describe pv nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2
```
```
Name:            nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2
Labels:          csi.alibabacloud.com/nas-id=960b448c84
Annotations:     pv.kubernetes.io/provisioned-by: nasplugin.csi.alibabacloud.com
                 volume.kubernetes.io/provisioner-deletion-secret-name: 
                 volume.kubernetes.io/provisioner-deletion-secret-namespace: 
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    alicloud-nas-subpath
Status:          Bound
Claim:           default/nas-csi-pvc
Reclaim Policy:  Delete
Access Modes:    RWX
VolumeMode:      Filesystem
Capacity:        20Gi
Node Affinity:   <none>
Message:         
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            nasplugin.csi.alibabacloud.com
    FSType:            nfs
    VolumeHandle:      nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2
    ReadOnly:          false
    VolumeAttributes:      path=/nas-d42d38ba-fb5e-4e05-9eb8-f16ef4047ec2/share
                           server=960b448c84-ygd8.cn-zhangjiakou.nas.aliyuncs.com
                           storage.kubernetes.io/csiProvisionerIdentity=1713419709568-7229-nasplugin.csi.alibabacloud.com
                           volumeAs=subpath
Events:                <none>
```
## Mount a dynamically provisioned NAS volume in filesystem mode
### Step 1: Create `filesystem` type nas volume
The `filesystem` type nas volume will create/delete a nas fileSystem and mountTarget automatically.

```shell
kubectl create -f ./examples/nas/dynamic/sc-filesystem.yaml
```
>fileSystemType: The type of NAS file system. 
> Valid values: 
> standard: General-purpose NAS file system 
> extreme: Extreme NAS file system
> 
> regionId: The ID of the region to which the NAS file system belongs.
> 
> zoneId: The ID of the zone to which the NAS file system belongs.
>
> vpcId: The ID of the VPC to which the mount target of the NAS file system belongs.
>
> vSwitchId: The ID of the vSwitch to which the mount target of the NAS file system belongs.
>
> storageType: The storage type of NAS file system.
> If fileSystemType is set to standard, the valid values are Performance and Capacity. Default value: Performance.
> If fileSystemType is set to extreme, the valid values are standard and advance. Default value: standard.

### Step 2: Create a PVC and pods to mount a NAS volume.
```shell
kubectl create -f ./examples/nas/dynamic/pvc-filesystem.yaml
```
