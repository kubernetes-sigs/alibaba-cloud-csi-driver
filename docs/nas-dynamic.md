
## Usage

### Step 1: Create CSI plugin
```
# kubectl create -f ./deploy/nas/nas-plugin.yaml
```

### Step 2: Create CSI provisioner
```
# kubectl create -f ./deploy/nas/nas-provisioner.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/nasplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 3: Create StorageClass for nas csi

#### 3.1 Create `subpath` type nas volume
The `subpath` type nas volume require you to create a nas fileSystem and mountTarget firstly

```
# kubectl create -f ./examples/nas/dynamic/sc-subpath.yaml
```

Below is a storageclass example:

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-nas
mountOptions:
- nolock,tcp,noresvport
- vers=3
parameters:
  volumeAs: subpath
  server: "2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com:/nasroot1/"
  archiveOnDelete: "false"
provisioner: nasplugin.csi.alibabacloud.com
reclaimPolicy: Delete
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

#### 3.2 Create `filesystem` type nas volume
The `filesystem` type nas volume will create/delete a nas fileSystem and mountTarget automatically.

```
# kubectl create -f ./examples/nas/dynamic/sc-filesystem.yaml
```

Below is a storageclass example:

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-nas-fs
mountOptions:
- nolock,tcp,noresvport
- vers=4
parameters:
  volumeAs: filesystem
  #storageType: Performance
  #zoneId: cn-hangzhou-f
  vpcId: "vpc-xxxxxxxx"
  vSwitchId: "vsw-xxxxxxxx"
  #accessGroupName: "DEFAULT_VPC_GROUP_NAME"
  #deleteVolume: "true"
provisioner: nasplugin.csi.alibabacloud.com
reclaimPolicy: Delete
```

Parameters:

> mountOptions: optional. specify the options for nfs mount.
>
> volumeAs: optional. Default as 'subpath'; specify the provision type: 'subpath' means the provision volume use subpath of exist nfs filesystem as target. And 'filesystem' means the provision volume use new created nfs filesystem as target.
>
> storageType: storageType is optional, should be 'Performance' or 'Capacity', default value is 'Performance'.
>
> zoneId: zoneId is optional, default value is the kubernetes cluster's worker's zoneId
>
> vpcId: vpcId is required, it is the node's vpc id you want to mount to
>
> vSwitchId: vSwitchId is required, it is the node's vSwitchId id you want to mount to
>
> accessGroupName: accessGroupName is optional, default value is 'DEFAULT_VPC_GROUP_NAME'
>
> deleteVolume: deleteVolume is optional, default value is 'true'


### Step 4: Check status of PVC
```
# kubectl get pvc | grep nas
nas-csi           Bound    nas-872f797c-ca1c-11e9-b9da-00163e063505    20Gi       RWX            alicloud-nas        6s

# kubectl get pv | grep nas
nas-872f797c-ca1c-11e9-b9da-00163e063505    20Gi       RWX            Delete           Bound    default/nas-csi           alicloud-nas                 21s
```


```
# kubectl describe pv nas-csi-pv
Name:            nas-872f797c-ca1c-11e9-b9da-00163e063505
Labels:          <none>
Annotations:     pv.kubernetes.io/provisioned-by: nasplugin.csi.alibabacloud.com
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    alicloud-nas
Status:          Bound
Claim:           default/nas-csi
Reclaim Policy:  Delete
Access Modes:    RWX
VolumeMode:      Filesystem
Capacity:        20Gi
Node Affinity:   <none>
Message:
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            nasplugin.csi.alibabacloud.com
    VolumeHandle:      nas-872f797c-ca1c-11e9-b9da-00163e063505
    ReadOnly:          false
    VolumeAttributes:      mode=755
                           path=/nasroot1/nas-872f797c-ca1c-11e9-b9da-00163e063505
                           server=2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com
                           storage.kubernetes.io/csiProvisionerIdentity=1567055848005-8081-nasplugin.csi.alibabacloud.com
                           vers=3
Events:                <none>
```
