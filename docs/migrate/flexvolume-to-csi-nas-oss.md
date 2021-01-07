## 1. 说明：
Flexvolume是Kubernetes早期支持的存储插件类型，目前社区依然支持Flexvolume作为存储插件使用。但由于Flexvolume为非容器化方案，社区提出了CSI存储接入规范，并会对CSI规范长期维护和更新。

阿里云容器服务从早期只支持Flexvolume插件，到目前即支持Flexvolume、也支持CSI。随着CSI的不断演进，Flexvolume在各种编排能力上和CSI逐步形成功能差异，所以从长期规划上会逐渐从Flexvolume演进到CSI的支持。

本文介绍如何将Flexvolume插件挂载的NAS、OSS类型存储转换到CSI插件类型；

## 2. 使用限制：

> 执行此转换，请将集群升级到1.16或者以上；
> 
> 如果您的集群中有使用Flexvolume挂载云盘，暂不适合使用此方案；
> 
> 实现插件类型转换，需要重启运行中的pod，重新定义pv、pvc；


## 3. CSI插件准备：
在集群中部署CSI 插件，模板见链接（您可以使用最新版本）：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

部署命令：

```
# kubectl apply -f csi-plugin.yaml
# kubectl apply -f csi-provisioner.yaml
```

部署时，需要将镜像地址换成您的集群所在的region地址，这样可以提高下载速度；例如：

> 把：registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
> 
> 改成：registry-vpc.cn-beijing.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun

部署完成后，检查是否部署成功：

```
# kubectl get pod -nkube-system |grep csi
csi-plugin-5jqf6                                    9/9     Running   0          4m15s
csi-plugin-dt8fd                                    9/9     Running   0          4m15s
csi-plugin-nlnpn                                    9/9     Running   0          4m15s
csi-plugin-p5lv6                                    9/9     Running   0          4m15s
csi-plugin-s29wp                                    9/9     Running   0          4m15s
csi-provisioner-58cb59cf6c-5c9qj                    8/8     Running   0          4m9s
csi-provisioner-58cb59cf6c-k96v2                    8/8     Running   0          4m9s
```

注意：部署csi-provisioner的时候，如果遇到StorageClass相关的error信息。<br />原因是：Flexvolume、Disk-Controller插件也有相应的StorageClass，并和csi插件的StorageClass重名。

您可以删除老的storageclass，再重新apply csi-provisioner；如下：

```
# kubectl get sc
NAME                       PROVISIONER                       AGE
alicloud-disk-available    alicloud/disk                     5h11m
alicloud-disk-efficiency   alicloud/disk                     5h11m
alicloud-disk-essd         alicloud/disk                     5h11m
alicloud-disk-ssd          alicloud/disk                     5h11m

# kubectl delete sc alicloud-disk-available alicloud-disk-efficiency alicloud-disk-essd alicloud-disk-ssd
storageclass.storage.k8s.io "alicloud-disk-available" deleted
storageclass.storage.k8s.io "alicloud-disk-efficiency" deleted
storageclass.storage.k8s.io "alicloud-disk-essd" deleted
storageclass.storage.k8s.io "alicloud-disk-ssd" deleted

# kubectl apply -f csi-provisioner.yaml

# kubectl get sc
NAME                       PROVISIONER                       AGE
alicloud-disk-available    diskplugin.csi.alibabacloud.com   84s
alicloud-disk-efficiency   diskplugin.csi.alibabacloud.com   84s
alicloud-disk-essd         diskplugin.csi.alibabacloud.com   84s
alicloud-disk-ssd          diskplugin.csi.alibabacloud.com   84s
alicloud-disk-topology     diskplugin.csi.alibabacloud.com   5m18s
```

## 4. 静态NAS卷迁移：

步骤：
> 1) 保存Flexvolume版本的pvc、pv模板；
> 
> 2) 创建csi类型的pvc、pv对象；
> 
> 3) 更新应用，使用新创建的csi类型的pvc；

下面为deployment使用nas pvc的示例：

```
# kubectl get deploy
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
nas-static   2/2     2            2           6m21s

# kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
nas-static-7f9446758b-2qfxj   1/1     Running   0          11m
nas-static-7f9446758b-wcwc5   1/1     Running   0          11m

# kubectl describe pod nas-static-7f9446758b-2qfxj |grep ClaimName
    ClaimName:  nas-pvc

# kubectl get pvc
NAME      STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
nas-pvc   Bound    nas-pv   512Gi      RWX            nas            12m

## 登陆pod所在节点，查看挂载信息：
2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com:/aliyun on /var/lib/kubelet/pods/b9f14e39-86ed-4c39-b4b7-baba37b15bac/volumes/alicloud~nas/nas-pv type nfs4 (rw,relatime,vers=3,rsize=1048576,wsize=1048576,namlen=255,hard,noresvport,proto=tcp,timeo=600,retrans=2,sec=sys,clientaddr=192.168.7.208,local_lock=none,addr=192.168.1.152)
```

步骤1：保存原始pvc(nas-pvc)、pv(nas-pv)

```
# kubectl get pvc nas-pvc -oyaml > pvc.nas-pvc.yaml
# cat pvc.nas-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nas-pvc
  namespace: default
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 512Gi
  selector:
    matchLabels:
      alicloud-pvname: nas-pv
  storageClassName: nas

# kubectl get pv nas-pv -oyaml > pv.nas-pv.yaml
# cat pv.nas-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  labels:
    alicloud-pvname: nas-pv
  name: nas-pv
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 512Gi
  flexVolume:
    driver: alicloud/nas
    options:
      path: /aliyun
      server: 2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com
      vers: "3"
  persistentVolumeReclaimPolicy: Retain
  storageClassName: nas
```

步骤2：创建CSI类型pvc(nas-pvc-new)、pv(nas-pv-new)

```
# cat pvc.nas-pvc-new.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nas-pvc-new
  namespace: default
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 512Gi
  selector:
    matchLabels:
      alicloud-pvname: nas-pv-new
  storageClassName: nas

## flexvolume配置改为csi配置；
## options配置可以写到：mountOptions中
# cat pv.nas-pv-new.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  labels:
    alicloud-pvname: nas-pv-new
  name: nas-pv-new
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 512Gi
  csi:
    driver: nasplugin.csi.alibabacloud.com
    volumeHandle: nas-pv-new
    volumeAttributes:
      server: "2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com"
      path: "/aliyun"
  mountOptions:
  - nolock,tcp,noresvport
  - vers=3
  persistentVolumeReclaimPolicy: Retain
  storageClassName: nas

# kubectl apply -f pvc.nas-pvc-new.yaml
# kubectl apply -f pv.nas-pv-new.yaml

# kubectl get pvc
NAME          STATUS   VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
nas-pvc       Bound    nas-pv       512Gi      RWX            nas            30m
nas-pvc-new   Bound    nas-pv-new   512Gi      RWX            nas            2s
```

步骤3：更新应用，使用新的csi类型pvc；

```
# kubectl edit deploy nas-static

把数据卷配置改为新的pvc；
      volumes:
      - name: pvc-nas
        persistentVolumeClaim:
          claimName: nas-pvc-new

## 查看Pod已经重启：
# kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
nas-static-6f7d965689-7sm5b   1/1     Running   0          70s
nas-static-6f7d965689-8npxw   1/1     Running   0          53s

## 查看挂载信息：
2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com:/aliyun on /var/lib/kubelet/pods/ac02ea3f-125f-4b38-9bcf-9b117f62eaf0/volumes/kubernetes.io~csi/nas-pv-new/mount type nfs (rw,relatime,vers=3,rsize=1048576,wsize=1048576,namlen=255,hard,nolock,noresvport,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=192.168.1.152,mountvers=3,mountport=2049,mountproto=tcp,local_lock=all,addr=192.168.1.152)
```

通过上面3个步骤，即完成了deployment从使用flexvolume挂载nas转换成使用csi插件挂载nas；

## 5. 动态NAS场景：

步骤：
> 1) 保存Flexvolume版本的storageclass、pvc、pv模板；
> 
> 2) 创建csi类型的storageclass、pvc、pv对象；
> 
> 3) 更新应用，使用新创建的csi类型的pvc；
> 
> 4) 后续新创建pvc时，都使用csi类型的storageclass；


下面为statefulset使用动态nas pv的示例：

```
# kubectl get sts
NAME   READY   AGE
web    2/2     3m11s

# kubectl get pod
NAME    READY   STATUS    RESTARTS   AGE
web-0   1/1     Running   0          3m14s
web-1   1/1     Running   0          118s

# kubectl get pvc
NAME            STATUS   VOLUME                                                           CAPACITY   ACCESS MODES   STORAGECLASS   AGE
nas-pvc-web-0   Bound    default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3   20Gi       RWX            alicloud-nas   3m32s
nas-pvc-web-1   Bound    default-nas-pvc-web-1-pvc-95eee8f9-54f4-42dc-a34f-b554423220eb   20Gi       RWX            alicloud-nas   2m16s

# kubectl get pv
NAME                                                             CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3   20Gi       RWX            Delete           Bound    default/nas-pvc-web-0   alicloud-nas            2m34s
default-nas-pvc-web-1-pvc-95eee8f9-54f4-42dc-a34f-b554423220eb   20Gi       RWX            Delete           Bound    default/nas-pvc-web-1   alicloud-nas            2m14s

# kubectl get sc alicloud-nas
NAME           PROVISIONER    AGE
alicloud-nas   alicloud/nas   106m
```

步骤1：保存原始pvc、pv、storageclass

```
## 保存pvc模板：
# kubectl get pvc nas-pvc-web-0 -oyaml > origin/nas-pvc-web-0.yaml
# cat origin/nas-pvc-web-0.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  annotations:
    volume.beta.kubernetes.io/storage-provisioner: alicloud/nas
  finalizers:
  - kubernetes.io/pvc-protection
  labels:
    app: nginx
  name: nas-pvc-web-0
  namespace: default
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 20Gi
  storageClassName: alicloud-nas
# kubectl get pvc nas-pvc-web-1 -oyaml > origin/nas-pvc-web-1.yaml

## 保存pv模板：
# kubectl get pv default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3 -oyaml > origin/default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3.yaml
# cat origin/default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  annotations:
    pv.kubernetes.io/provisioned-by: alicloud/nas
  labels:
    createdby.aliyun.com: alicloud-nas-controller
    version.controller.aliyun.com: v1.14.3-58bf821
  name: default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 20Gi
  flexVolume:
    driver: alicloud/nas
    options:
      options: nolock,tcp,noresvport,vers=3
      path: /root/default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3
      server: 2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com
      vers: "3"
  persistentVolumeReclaimPolicy: Delete
  storageClassName: alicloud-nas
# kubectl get pv default-nas-pvc-web-1-pvc-95eee8f9-54f4-42dc-a34f-b554423220eb -oyaml > origin/default-nas-pvc-web-1-pvc-95eee8f9-54f4-42dc-a34f-b554423220eb.yaml

## 保存stroageclass模板：
# kubectl get sc alicloud-nas -oyaml > origin/alicloud-nas.yaml
# cat origin/alicloud-nas.yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-nas
mountOptions:
- nolock,tcp,noresvport
- vers=3
parameters:
  archiveOnDelete: "true"
  driver: flexvolume
  server: 2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com:/root/
provisioner: alicloud/nas
reclaimPolicy: Delete
volumeBindingMode: Immediate
```

步骤2：创建csi类型的storageclass、pvc、pv

```
## 老pvc：nas-pvc-web-0 对应新pvc: nas-pvc-new-web-0，这里nas-pvc字符是statefulset中配置的volume名字，更新为nas-pvc-new；
## provisioner改为：nasplugin.csi.alibabacloud.com
## name改为：nas-pvc-new-web-0
## storageClassName改为：alicloud-nas-new
## alicloud-pvname改为：nas-pvc-new-web-0-pv （这个名字是任意的，但为了可读性，可以在pvc名字后面加上pv或者namespace等，保证全局唯一）

# cat nas-pvc-new-web-0.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  annotations:
    volume.beta.kubernetes.io/storage-provisioner: nasplugin.csi.alibabacloud.com
  labels:
    app: nginx
  name: nas-pvc-new-web-0
  namespace: default
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 20Gi
  storageClassName: alicloud-nas-new
  selector:
    matchLabels:
      alicloud-pvname: nas-pvc-new-web-0-pv

## 老pv：default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3 更新为下面新pv；
## provisioned-by改为：nasplugin.csi.alibabacloud.com
## labels改为：nas-pvc-new-web-0-pv
## name改为：nas-pvc-new-web-0-pv
## flexvolume配置改为csi配置：volumeHandle改为pv的名字；
## storageClassName改为：alicloud-nas-new

# cat nas-pvc-new-web-0-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  annotations:
    pv.kubernetes.io/provisioned-by: nasplugin.csi.alibabacloud.com
  finalizers:
  - kubernetes.io/pv-protection
  labels:
    alicloud-pvname: nas-pvc-new-web-0-pv
  name: nas-pvc-new-web-0-pv
spec:
  accessModes:
  - ReadWriteMany
  capacity:
    storage: 20Gi
  csi:
    driver: nasplugin.csi.alibabacloud.com
    volumeHandle: nas-pvc-new-web-0-pv
    volumeAttributes:
      server: "2564f49129-ysu87.cn-shenzhen.nas.aliyuncs.com"
      path: "/root/default-nas-pvc-web-0-pvc-951911bd-b7f3-4551-8647-22f1f5d576b3"
  mountOptions:
  - nolock,tcp,noresvport
  - vers=3
  persistentVolumeReclaimPolicy: Delete
  storageClassName: alicloud-nas-new

同理创建另外的pvc、pv：
老pvc：nas-pvc-web-1更新为: nas-pvc-new-web-1
老pv：default-nas-pvc-web-1-pvc-95eee8f9-54f4-42dc-a34f-b554423220eb更新为: nas-pvc-new-web-1-pv

# kubectl get pvc
NAME                STATUS   VOLUME                                                           CAPACITY   ACCESS MODES   STORAGECLASS       AGE
nas-pvc-new-web-0   Bound    nas-pvc-new-web-0-pv                                             20Gi       RWX            alicloud-nas-new   13m
nas-pvc-new-web-1   Bound    nas-pvc-new-web-1-pv                                             20Gi       RWX            alicloud-nas-new   10m
```

步骤3：更新应用，volume改为nas-pvc-new

```
## statefulset模板中的volume，从nas-pvc更新为nas-pvc-new：
  volumeClaimTemplates:
  - metadata:
      name: nas-pvc-new
    spec:

## 删除老的statefulset（statefulset的volumeClaimTemplates不能编辑），并创建新的；
# kubectl delete sts web
# kubectl apply -f sts.yaml
```

检查新建应用：

```
# kubectl get pod
NAME    READY   STATUS    RESTARTS   AGE
web-0   1/1     Running   0          9m37s
web-1   1/1     Running   0          9m20s

# kubectl describe pod web-0 | grep ClaimName
    ClaimName:  nas-pvc-new-web-0

# kubectl get pvc nas-pvc-new-web-0
NAME                STATUS   VOLUME                 CAPACITY   ACCESS MODES   STORAGECLASS       AGE
nas-pvc-new-web-0   Bound    nas-pvc-new-web-0-pv   20Gi       RWX            alicloud-nas-new   17m
```

## 6. 静态OSS转换：
步骤：

> 1) 保存Flexvolume版本的pvc、pv模板；
> 
> 2) 创建csi类型的pvc、pv对象；
> 
> 3) 更新应用，使用新创建的csi类型的pvc；

下面为deployment 使用 oss pvc/pv 的示例

```
# kubectl get deploy
NAME              READY   UP-TO-DATE   AVAILABLE   AGE
nginx-oss-deploy   1/1     1            1           5m37s

# kubectl get pod
NAME                              READY   STATUS    RESTARTS   AGE
nginx-oss-deploy-f49d8d6c4-df892   1/1     Running   0          6m1s

# kubectl describe pod nas-static-7f9446758b-2qfxj |grep ClaimName
    ClaimName:  pvc-oss

# kubectl get pvc
NAME      STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
pvc-oss   Bound    pv-oss   5Gi        RWX                           7m23s

# kubectl get pv pv-oss -o yaml

    flexVolume:
      driver: alicloud/oss
      options:
        akId: xxxx
        akSecret: xxx
        bucket: oss-mount-test
        otherOpts: -o max_stat_cache_size=0 -o allow_other
        url: oss-cn-shanghai.aliyuncs.com

```

步骤一：保存原始的 pvc(pvc-oss), pv(pv-oss)

```
# kubectl get pvc pvc-oss -oyaml > flex.oss-pvc.yaml
# cat flex.oss-pvc.yaml
apiVersion: v1
items:
- apiVersion: v1
  kind: PersistentVolumeClaim
    name: pvc-oss
    namespace: default
  spec:
    accessModes:
    - ReadWriteMany
    resources:
      requests:
        storage: 5Gi
    volumeMode: Filesystem
    volumeName: pv-oss
  status:
    accessModes:
    - ReadWriteMany
    capacity:
      storage: 5Gi
    phase: Bound
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""

# kubectl get pv pv-oss -oyaml > flex.oss-pv.yaml
# cat flex.oss-pv.yaml
apiVersion: v1
items:
- apiVersion: v1
  kind: PersistentVolume
  metadata:
    name: pv-oss
  spec:
    accessModes:
    - ReadWriteMany
    capacity:
      storage: 5Gi
    claimRef:
      apiVersion: v1
      kind: PersistentVolumeClaim
      name: pvc-oss
      namespace: default
    flexVolume:
      driver: alicloud/oss
      options:
        akId: xxx
        akSecret: xxx
        bucket: oss-mount-test
        otherOpts: -o max_stat_cache_size=0 -o allow_other
        url: oss-cn-shanghai.aliyuncs.com
    persistentVolumeReclaimPolicy: Retain
    volumeMode: Filesystem
  status:
    phase: Bound
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

步骤二：创建CSI类型pvc(pvc-oss-new)、pv(pv-oss-new)

```
# cat csi.oss-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: oss-csi-pvc
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  selector:
    matchLabels:
      alicloud-pvname: oss-csi-pv

# cat csi.oss-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: oss-csi-pv
  labels:
    alicloud-pvname: oss-csi-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: ossplugin.csi.alibabacloud.com
    # set volumeHandle same value pv name
    volumeHandle: oss-csi-pv
    volumeAttributes:
      bucket: "oss-mount-test"
      url: "oss-cn-shanghai.aliyuncs.com"
      otherOpts: "-o max_stat_cache_size=0 -o allow_other"
      akId: "***"
      akSecret: "***"
      path: "/"

# kubectl apply -f csi.oss-pv.yaml
# kubectl apply -f csi.oss-pvc.yaml

# kubectl get pvc
NAME         STATUS   VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS   AGE
oss-csi-pvc   Bound    oss-csi-pv   5Gi        RWO                           7m15s
pvc-oss       Bound    pv-oss       5Gi        RWX                           52m

```

步骤3：更新应用，使用新的csi类型pvc；

```
# kubectl edit deploy nginx-oss-deploy
        volumes:
        - name: oss1
          persistentVolumeClaim:
            claimName: oss-csi-pvc
// 查看pod已重启
# kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
nginx-oss-deploy-d84b9bb86-nxbf5   1/1     Running   0          21s
```

通过上面3个步骤，即完成了deployment从使用flexvolume挂载oss转换成使用csi插件挂载oss；

对应oss数据卷使用Secret的场景，步骤:

> 1）获取 flexvolume deployment 使用的secret
> 
> 2）使用 secret 创建静态卷 oss pvc/pv,
> 
> 3）编辑 flexvolume deployment 引用新创建的 pvc/pv

步骤一：获取 flexvolume deployment 使用的secret

```
# kubectl get secret osssecret -o yaml
apiVersion: v1
data:
  akId: xxxx
  akSecret: xxxx
kind: Secret
metadata:
  name: oss-secret
  namespace: default
type: alicloud/oss

```

步骤二： 使用 secret 创建静态卷 oss-csi pvc/pv,

```
# kubectl apply -f secret-pvc-pv.yaml

apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: oss-csi-pvc-secret
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 5Gi
  selector:
    matchLabels:
      alicloud-pvname: oss-csi-pv-secret
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: oss-csi-pv-secret
  labels:
    alicloud-pvname: oss-csi-pv-secret
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: ossplugin.csi.alibabacloud.com
    volumeHandle: oss-csi-pv-secret // 需要和PV名字一致。
    nodePublishSecretRef:
      name: oss-secret
      namespace: default
    volumeAttributes:
      bucket: "oss"
      url: "oss-cn-hangzhou.aliyuncs.com"
      otherOpts: "-o max_stat_cache_size=0 -o allow_other"

# kubectl get pvc oss-csi-pvc-secret
NAME                 STATUS   VOLUME              CAPACITY   ACCESS MODES   STORAGECLASS   AGE
oss-csi-pvc-secret   Bound    oss-csi-pv-secret   5Gi        RWX                           3m35s
```

步骤三： 更改deployment 引用，使用新创建的pvc

```
# kubectl edit deploy nginx-oss-deploy
        volumes:
        - name: oss1
          persistentVolumeClaim:
            claimName: oss-csi-pvc-secret

// 查看pod已重启
# kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
nginx-oss-deploy-d84bxxrs6-nxbf5   1/1     Running   0          2s
```
