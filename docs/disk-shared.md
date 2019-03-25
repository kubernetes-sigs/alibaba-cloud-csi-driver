
## Shared Disk

Alibaba Cloud provides shared-disk for customers, which can be attached to multi-nodes(no more than 8).

Shared Disk always attached to multi-nodes, and should be mounted as block volume.

Shared Disk Refer: [https://help.aliyun.com/document_detail/108902.html](https://help.aliyun.com/document_detail/108902.html)

## How to Deploy

Same as common disk.

## How to Use

### 1. Template

Storageclass:

	Shared disk has different types: san_ssd, san_efficiency

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: csi-disk
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    zoneId: cn-hangzhou-b
    regionId: cn-hangzhou
    fsType: ext4
    type: san_ssd
    readOnly: "false"
reclaimPolicy: Delete
```

PVC: Same as block volume;

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: block-pvc
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Block
  storageClassName: csi-disk
  resources:
    requests:
      storage: 20Gi
```

Use daemonset to test mount to multi node.

```
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: ds
  labels:
    k8s-volume: nginx
spec:
  selector:
    matchLabels:
      name: nginx
  template:
    metadata:
      labels:
        name: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
        volumeDevices:
          - name: disk-pvc
            devicePath: /dev/sdc
      volumes:
        - name: disk-pvc
          persistentVolumeClaim:
            claimName: block-pvc
```

### 2. Create StorageClass

```
# kubectl create -f storageclass.yaml

# kubectl get sc
NAME       PROVISIONER                       AGE
csi-disk   diskplugin.csi.alibabacloud.com   18m

# kubectl describe sc csi-disk
Name:                  csi-disk
IsDefaultClass:        No
Annotations:           <none>
Provisioner:           diskplugin.csi.alibabacloud.com
Parameters:            fsType=ext4,readOnly=false,regionId=cn-hangzhou,type=san_ssd,zoneId=cn-hangzhou-b
AllowVolumeExpansion:  <unset>
MountOptions:          <none>
ReclaimPolicy:         Delete
VolumeBindingMode:     Immediate
Events:                <none>
```

### 3. Create PV, PVC

Set "VolumeMode:    Block" for PVC.

```
Create PVC
# kubectl create -f pvc.yaml

# kubectl get pvc
NAME        STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
block-pvc   Bound    pvc-a113496d-4198-11e9-983f-00163e0b8d64   20Gi       RWX            csi-disk       19m

# # kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM               STORAGECLASS   REASON   AGE
pvc-a113496d-4198-11e9-983f-00163e0b8d64   20Gi       RWX            Delete           Bound    default/block-pvc   csi-disk                19m
# kubectl describe pv pvc-a113496d-4198-11e9-983f-00163e0b8d64
Name:            pvc-a113496d-4198-11e9-983f-00163e0b8d64
Labels:          <none>
Annotations:     pv.kubernetes.io/provisioned-by: diskplugin.csi.alibabacloud.com
Finalizers:      [kubernetes.io/pv-protection external-attacher/diskplugin-csi-alibabacloud-com]
StorageClass:    csi-disk
Status:          Bound
Claim:           default/block-pvc
Reclaim Policy:  Delete
Access Modes:    RWX
VolumeMode:      Block
Capacity:        20Gi
Node Affinity:   <none>
Message:
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            diskplugin.csi.alibabacloud.com
    VolumeHandle:      v-bp1bwmga7u3uyehtt86t
    ReadOnly:          false
    VolumeAttributes:      fsType=ext4
                           readOnly=false
                           regionId=cn-hangzhou
                           shared=enable
                           storage.kubernetes.io/csiProvisionerIdentity=1552045668948-8081-diskplugin.csi.alibabacloud.com
                           type=san_ssd
                           zoneId=cn-hangzhou-b
Events:                <none>
```

### 4. Create DaemonSet

The disk is attached in 2 Pods and be writable with dd command.

```
# kubectl create -f daemonset.yaml

# kubectl get pod | grep ds
ds-jd7h4                     1/1     Running   0          20m
ds-zdlhz                     1/1     Running   0          20m

# kubectl exec -ti ds-jd7h4 sh
# ls -l /dev/sdc
brw-rw---- 1 root disk 253, 32 Mar  8 11:53 /dev/sdc
# dd if=/dev/zero of=/dev/sdc bs=1024k count=100
100+0 records in
100+0 records out
104857600 bytes (105 MB) copied, 1.36235 s, 77.0 MB/s

# kubectl exec -ti ds-zdlhz sh
# ls -l /dev/sdc
brw-rw---- 1 root disk 253, 32 Mar  8 11:52 /dev/sdc
# dd if=/dev/zero of=/dev/sdc bs=1024k count=100
100+0 records in
100+0 records out
104857600 bytes (105 MB) copied, 0.0616086 s, 1.7 GB/s
```