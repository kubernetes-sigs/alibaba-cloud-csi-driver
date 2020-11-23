负责本地盘的创建、删除操作：通过在每个节点上启动agent，并watch PVC、PV，通过label来实现与provisioner之间的协作；

> 优点：复用K8S网络，而不依赖集群端口和网络，在边缘、混合云等场景下无缝衔接；
>
> 缺点：集群非常大时，由于每个节点都需要watch pvc、pv，会产生大量长连接；


## 创建过程：跟踪PVC

1. 添加PVC Annotations：

    provisioner 给pvc添加Annotations：pv.csi.alibabacloud.com/volume.lifecycle为creating：

2. nodeServer跟踪PVC的标签：pv.csi.alibabacloud.com/volume.lifecycle为creating时：

   查看pv.csi.alibabacloud.com/volume.spec，拿到创建lvm所需的配置信息；

3. 创建LVM卷；

4. 创建成功后更新PVC Annotations: pv.csi.alibabacloud.com/volume.lifecycle 为created；

创建流程结束；


## 删除过程：跟踪PV

1. Provisioner为PV添加Annotations：pv.csi.alibabacloud.com/volume.lifecycle为deleting：

2. nodeServer跟踪PV的Annotations：pv.csi.alibabacloud.com/volume.lifecycle 为deleting时：

   查看PV的Annotations：pv.csi.alibabacloud.com/volume.spec，拿到删除lvm所需的信息；

3. 删除LVM卷；

4. 成功后更新PV Annotations: pv.csi.alibabacloud.com/volume.lifecycle 为deleted；

删除流程结束；