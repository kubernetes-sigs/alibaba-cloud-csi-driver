
## Usage

**Deprecated**

### Step 1: Create CSI plugin
```
# kubectl create -f ./deploy/cpfs/cpfs-plugin.yaml
```

### Step 2: Create CSI cpfs provisioner
```
# kubectl create -f ./deploy/cpfs/cpfs-provisioner.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/cpfsplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 3: Create StorageClass for nas csi
```
# kubectl create -f ./examples/cpfs/dynamic/sc-subpath.yaml
```

Below is a storageclass example:

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-cpfs
parameters:
  volumeAs: subpath
  server: "cpfs-09ab7b80-5bh1.cn-zhangjiakou.cpfs.nas.aliyuncs.com@tcp:cpfs-09ab7b80-nu9b.cn-zhangjiakou.cpfs.nas.aliyuncs.com@tcp:/09ab7b80"
  archiveOnDelete: "false"
provisioner: cpfsplugin.csi.alibabacloud.com
reclaimPolicy: Delete
```

Parameters:

> mountOptions: optional. specify the options for cpfs mount.
>
> volumeAs: optional. Default as 'subpath'; specify the provision type: 'subpath' means the provision volume use subpath of exist cpfs filesystem as target. And 'filesystem' means the provision volume use new created cpfs filesystem as target.
>
> server: Necessary when volumeAs is subpath. Define the cpfs server endpoints, format as: cpfsserver:/filesystem1/path1,cpfsserver2:/filesystem2/path2, the volume server target will use the cpfsserver/filesystem/path by order.
>
> archiveOnDelete: Optional. decide how to process removal path, if reclaimPolicy defined as delete. If set 'true', the removal path will be archived and not removed really, and if set 'false', the removal path will be removed when pv is deleted.

### Step 4: Check status of PVC
```
# kubectl get pvc
NAME           STATUS   VOLUME                                      CAPACITY   ACCESS MODES   STORAGECLASS    AGE
cpfs-csi-pvc   Bound    cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69   20Gi       RWX            alicloud-cpfs   2s

# kubectl get pv
NAME                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                  STORAGECLASS    REASON   AGE
cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69   20Gi       RWX            Delete           Bound    default/cpfs-csi-pvc   alicloud-cpfs            20s
```


```
# kubectl describe pv cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69
Name:            cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69
Labels:          <none>
Annotations:     pv.kubernetes.io/provisioned-by: cpfsplugin.csi.alibabacloud.com
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    alicloud-cpfs
Status:          Bound
Claim:           default/cpfs-csi-pvc
Reclaim Policy:  Delete
Access Modes:    RWX
VolumeMode:      Filesystem
Capacity:        20Gi
Node Affinity:   <none>
Message:
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            cpfsplugin.csi.alibabacloud.com
    VolumeHandle:      cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69
    ReadOnly:          false
    VolumeAttributes:      fileSystem=09ab7b80
                           server=cpfs-09ab7b80-5bh1.cn-zhangjiakou.cpfs.nas.aliyuncs.com@tcp:cpfs-09ab7b80-nu9b.cn-zhangjiakou.cpfs.nas.aliyuncs.com@tcp
                           storage.kubernetes.io/csiProvisionerIdentity=1570712973671-8081-cpfsplugin.csi.alibabacloud.com
                           subPath=/cpfs-82b1ce62-eb6f-11e9-a442-00163e07fb69
Events:                <none>
```
