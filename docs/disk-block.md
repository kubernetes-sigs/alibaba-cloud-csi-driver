
## Block Volume
Block Volume is type of kubernetes volume type, which enable you use the raw block storage in your pod.
In Block Volume type, the driver does not try to format disk and just bind the block device to target path.

Block Volume: [https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/raw-block-pv.md](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/raw-block-pv.md)

CSI Block Volume: [https://kubernetes-csi.github.io/docs/raw-block.html](https://kubernetes-csi.github.io/docs/raw-block.html)

## Feature Supports

| Feature        | Default | From | To   | Stage |
|----------------|---------|------|------|-------|
| CSIBlockVolume | false   | 1.11 | 1.13 | Alpha |
| CSIBlockVolume | true    | 1.14 | 1.17 | Beta  |
| CSIBlockVolume | true    | 1.18 | 1.21 | GA    |
| BlockVolume    | false   | 1.9  | 1.12 | Alpha |
| BlockVolume    | true    | 1.13 | 1.17 | Beta  |
| BlockVolume    | true    | 1.18 | 1.21 | GA    |


## How to Deploy

### Requirements

Enable Feature Gates:

    --feature-gates=VolumeSnapshotDataSource=true,CSINodeInfo=true,CSIDriverRegistry=true,BlockVolume=true,CSIBlockVolume=true

Other Requirements same as common disk.

### Deploy Plugin

The different between Block and FileSystem for disk-plugin: block type need volumeDevices volume mount;

```shell
kubectl create -f ./deploy/disk/block/disk-plugin.yaml
```

Other parts are same as common disk.

## How to Use

### 1. Template

StorageClass: No changed.

PVC: volumeMode is configured "Block"

```yaml
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

Use volumeDevices to define the target file.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-block
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
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

### 2. Create PV, PVC

Check that: "VolumeMode: Block" for PV and PVC.

Create PVC, PV
```shell
kubectl create -f pvc.yaml
```
Check pvc status
```shell
kubectl get pvc | grep csi-disk
```
Expected output:
```
block-pvc         Bound    pvc-6017e55c-4168-11e9-983f-00163e0b8d64   20Gi       RWO            csi-disk       8s
```
Describe pvc
```shell
kubectl describe pvc block-pvc
```
Expected output:
```
Name:          block-pvc
Namespace:     default
StorageClass:  csi-disk
Status:        Bound
Volume:        pvc-6017e55c-4168-11e9-983f-00163e0b8d64
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: diskplugin.csi.alibabacloud.com
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      20Gi
Access Modes:  RWO
VolumeMode:    Block
Events:
```
Check pv status
```shell
kubectl get pv
```
Expected output:
```
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                     STORAGECLASS   REASON   AGE
pvc-6017e55c-4168-11e9-983f-00163e0b8d64   20Gi       RWO            Delete           Bound    default/block-pvc         csi-disk                17m
```
Describe pv
```shell
kubectl describe pv pvc-6017e55c-4168-11e9-983f-00163e0b8d64
```
Expected output:
```
Name:            pvc-6017e55c-4168-11e9-983f-00163e0b8d64
Labels:          <none>
Annotations:     pv.kubernetes.io/provisioned-by: diskplugin.csi.alibabacloud.com
Finalizers:      [kubernetes.io/pv-protection external-attacher/diskplugin-csi-alibabacloud-com]
StorageClass:    csi-disk
Status:          Bound
Claim:           default/block-pvc
Reclaim Policy:  Delete
Access Modes:    RWO
VolumeMode:      Block
Capacity:        20Gi
Node Affinity:   <none>
Message:
Source:
    Type:              CSI (a Container Storage Interface (CSI) volume source)
    Driver:            diskplugin.csi.alibabacloud.com
    VolumeHandle:      d-bp1bwmga7u3utyo3twbb
    ReadOnly:          false
    VolumeAttributes:      fsType=ext4
                           readOnly=false
                           regionId=cn-hangzhou
                           storage.kubernetes.io/csiProvisionerIdentity=1551951411379-8081-diskplugin.csi.alibabacloud.com
                           type=cloud_ssd
                           zoneId=cn-hangzhou-b
Events:                <none>
```

### 3. Create Workload

The disk is attached in Pod and be writable with dd command.

```shell
kubectl create -f deploy.yaml
```
```shell
kubectl get pod | grep nginx-block
```
Expected output:
```
nginx-block-86ddb74484-sz2hm   1/1     Running   0          24s
```

Run the following command to list the block device
```shell
kubectl exec -it nginx-block-86ddb74484-sz2hm sh
```
```shell
ls -l /dev/sdc
```
Expected output:
```
brw-rw---- 1 root disk 253, 16 Mar  8 06:07 /dev/sdc
```
Run the following command to check the disk is writable
```shell
dd if=/dev/zero of=/dev/sdc bs=1024k count=100
```
Expected output:
```
100+0 records in
100+0 records out
104857600 bytes (105 MB) copied, 0.0584359 s, 1.8 GB/s
```