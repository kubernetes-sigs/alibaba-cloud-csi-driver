## flexvolume 迁移至 csi -- 云盘 sts - volumeClaimTemplates version

> 声明：
> 1. 此转换方案为Alpha版本，请务必在业务低峰时进行；
> 2. 进行flexvolume到csi的转换，尽量在短时间内完成，并尽快切换到csi插件体系。尽量减少两种插件并行存在的时间。
> 3. 大规模生产集群暂不适合使用此方案，如果期望进行插件切换，请开Issue留言联系我们；
> 4. 此版本文档只适用于 statefulset 版本的应用

### 背景介绍
csi, flexvolume 简介： https://help.aliyun.com/document_detail/157025.html?spm=5176.2020520152.0.dexternal.52a116ddjh5Afv

csi 插件体系为 csi-provisioner deployment + csi-plugin daemonset 组合而成

flexvolume 插件体系为 disk-controller deployment + flexvolume daemonset 组合而成

本文是将应用从使用 flexvolume 插件体系创建的存储卷迁移到使用 csi 插件体系创建的存储卷的方案, csi 插件体系较 flexvolume 更加稳定高效

### 环境准备
我们以一个例子来说明升级流程, flexvolume 应用 yaml 如下。下面讲述如何将这个 flexvolume 的应用升级为 csi 的应用
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web-multi
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: disk-ssd-test
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: disk-ssd-test
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "alicloud-disk-efficiency"
      resources:
        requests:
          storage: 20Gi
```

1. 首先使用 ``` kubectl apply -f sts.yaml``` 命令创建上述 statefulset,
2. 使用 ``` kubectl get pvc -oyaml```, ``` kubectl get pv -oyaml```查看 statefulset 创建的 PVC/PV 

> 本例是有三个副本的 statefulset, 会有三对 PVC/PV, 为了缩短文章的长度，以下均已首个副本的 PVC/PV 为例.

```yaml
apiVersion: v1
items:
- apiVersion: v1
  kind: PersistentVolumeClaim
  metadata:
    name: disk-ssd-test-web-multi-0
    namespace: default
  spec:
    accessModes:
    - ReadWriteOnce
    resources:
      requests:
        storage: 20Gi
    storageClassName: alicloud-disk-efficiency
    volumeMode: Filesystem
    volumeName: d-8vb61ll1g3ykxojuvdg4
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
---
apiVersion: v1
items:
- apiVersion: v1
  kind: PersistentVolume
  metadata:
    name: d-8vb61ll1g3ykxojuvdg4
  spec:
    accessModes:
    - ReadWriteOnce
    capacity:
      storage: 20Gi
    claimRef:
      apiVersion: v1
      kind: PersistentVolumeClaim
      name: disk-ssd-test-web-multi-0
      namespace: default
      resourceVersion: "726972"
      uid: a4ee1509-9258-48e4-afcd-884dc1d70131
    flexVolume:
      driver: alicloud/disk
      fsType: ext4
      options:
        VolumeId: d-8vb61ll1g3ykxojuvdg4
    persistentVolumeReclaimPolicy: Delete
    storageClassName: alicloud-disk-efficiency
    volumeMode: Filesystem
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""

```
3. 记录下 pv 使用的云盘id
```id: d-8vb61ll1g3ykxojuvdg4```


### 部署 csi 插件
#### ACK专有版
在 Kubernetes 集群中部署CSI插件, 和默认部署版本的主要区别为：CSIDriver中attachRequired为false；

```yaml
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: diskplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: nasplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: ossplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-plugin
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-plugin
  template:
    metadata:
      labels:
        app: csi-plugin
    spec:
      tolerations:
        - operator: Exists
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
      nodeSelector:
        beta.kubernetes.io/os: linux
      serviceAccount: admin
      priorityClassName: system-node-critical
      hostNetwork: true
      hostPID: true
      containers:
        - name: disk-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
            - name: registration-dir
              mountPath: /registration
        - name: nas-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
            - name: registration-dir
              mountPath: /registration
        - name: oss-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
            - name: registration-dir
              mountPath: /registration
        - name: csi-plugin
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          args:
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--v=2"
            - "--driver=oss,nas,disk"
          env:
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-plugins/driverplugin.csi.alibabacloud.com-replace/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
            - name: SERVICE_TYPE
              value: "plugin"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            timeoutSeconds: 5
            failureThreshold: 5
          ports:
            - name: healthz
              containerPort: 11260
              protocol: TCP
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
              mountPropagation: "Bidirectional"
            - name: etc
              mountPath: /host/etc
            - name: host-log
              mountPath: /var/log/
            - name: ossconnectordir
              mountPath: /host/usr/
            - name: container-dir
              mountPath: /var/lib/container
              mountPropagation: "Bidirectional"
            - name: host-dev
              mountPath: /dev
              mountPropagation: "HostToContainer"
      volumes:
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry
            type: DirectoryOrCreate
        - name: container-dir
          hostPath:
            path: /var/lib/container
            type: DirectoryOrCreate
        - name: kubelet-dir
          hostPath:
            path: /var/lib/kubelet
            type: Directory
        - name: host-dev
          hostPath:
            path: /dev
        - name: host-log
          hostPath:
            path: /var/log/
        - name: etc
          hostPath:
            path: /etc
        - name: ossconnectordir
          hostPath:
            path: /usr/
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 10%
    type: RollingUpdate
```

部署csi-provisioner：为了和alicloud-disk-controller区分，这里的storageclass都在名字后面加上-csi字样；
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-available-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-essd-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_essd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-ssd-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_ssd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-efficiency-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_efficiency
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-topology-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-provisioner
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-provisioner
  replicas: 2
  template:
    metadata:
      labels:
        app: csi-provisioner
    spec:
      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            preference:
              matchExpressions:
              - key: node-role.kubernetes.io/master
                operator: Exists
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - csi-provisioner
              topologyKey: "kubernetes.io/hostname"
      tolerations:
      - effect: NoSchedule
        operator: Exists
        key: node-role.kubernetes.io/master
      - effect: NoSchedule
        operator: Exists
        key: node.cloudprovider.kubernetes.io/uninitialized
      priorityClassName: system-node-critical
      serviceAccount: admin
      hostNetwork: true
      containers:
        - name: external-disk-provisioner
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v3.0.0-3f86569-aliyun
          args:
            - "--csi-address=$(ADDRESS)"
            - "--feature-gates=Topology=True"
            - "--volume-name-prefix=disk"
            - "--strict-topology=true"
            - "--timeout=150s"
            - "--leader-election=true"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-attacher
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-attacher:v2.1.0
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-resizer
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-resizer:v0.3.0
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-nas-provisioner
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v1.4.0-aliyun
          args:
            - "--provisioner=nasplugin.csi.alibabacloud.com"
            - "--csi-address=$(ADDRESS)"
            - "--volume-name-prefix=nas"
            - "--timeout=150s"
            - "--enable-leader-election=true"
            - "--leader-election-type=leases"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
        - name: external-csi-snapshotter
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-snapshotter:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          imagePullPolicy: Always
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /csi
        - name: external-snapshot-controller
          image: registry.cn-hangzhou.aliyuncs.com/acs/snapshot-controller:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--leader-election=true"
          imagePullPolicy: Always
        - name: csi-provisioner
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          args:
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--v=2"
            - "--driver=nas,disk"
          env:
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-provisioner/driverplugin.csi.alibabacloud.com-replace/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
            - name: SERVICE_TYPE
              value: "provisioner"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            timeoutSeconds: 5
            failureThreshold: 5
          ports:
            - name: healthz
              containerPort: 11270
              protocol: TCP
          volumeMounts:
            - name: host-dev
              mountPath: /dev
              mountPropagation: "HostToContainer"
            - name: host-log
              mountPath: /var/log/
            - name: etc
              mountPath: /host/etc
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
          resources:
            limits:
              cpu: 1000m
              memory: 1000Mi
            requests:
              cpu: 100m
              memory: 100Mi
      volumes:
        - name: disk-provisioner-dir
          emptyDir: {}
        - name: nas-provisioner-dir
          emptyDir: {}
        - name: host-log
          hostPath:
            path: /var/log/
        - name: host-dev
          hostPath:
            path: /dev
        - name: etc
          hostPath:
            path: /etc
```
#### ACK托管版
csi-plugin 插件部署
```yaml
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: diskplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: nasplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: ossplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-plugin
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-plugin
  template:
    metadata:
      labels:
        app: csi-plugin
    spec:
      tolerations:
        - operator: Exists
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
      nodeSelector:
        beta.kubernetes.io/os: linux
      serviceAccount: admin
      priorityClassName: system-node-critical
      hostNetwork: true
      hostPID: true
      containers:
        - name: disk-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
            - name: registration-dir
              mountPath: /registration
        - name: nas-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
            - name: registration-dir
              mountPath: /registration
        - name: oss-driver-registrar
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
            - name: registration-dir
              mountPath: /registration
        - name: csi-plugin
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          args:
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--v=2"
            - "--driver=oss,nas,disk"
          env:
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-plugins/driverplugin.csi.alibabacloud.com-replace/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
            - name: SERVICE_TYPE
              value: "plugin"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            timeoutSeconds: 5
            failureThreshold: 5
          ports:
            - name: healthz
              containerPort: 11260
              protocol: TCP
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet/
              mountPropagation: "Bidirectional"
            - name: etc
              mountPath: /host/etc
            - name: host-log
              mountPath: /var/log/
            - name: ossconnectordir
              mountPath: /host/usr/
            - name: container-dir
              mountPath: /var/lib/container
              mountPropagation: "Bidirectional"
            - name: host-dev
              mountPath: /dev
              mountPropagation: "HostToContainer"
            - mountPath: /var/addon
              name: addon-token
              readOnly: true
      volumes:
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry
            type: DirectoryOrCreate
        - name: container-dir
          hostPath:
            path: /var/lib/container
            type: DirectoryOrCreate
        - name: kubelet-dir
          hostPath:
            path: /var/lib/kubelet
            type: Directory
        - name: host-dev
          hostPath:
            path: /dev
        - name: host-log
          hostPath:
            path: /var/log/
        - name: etc
          hostPath:
            path: /etc
        - name: ossconnectordir
          hostPath:
            path: /usr/
        - name: addon-token
          secret:
            defaultMode: 420
            items:
            - key: addon.token.config
              path: token-config
            secretName: addon.csi.token
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 10%
    type: RollingUpdate
```

csi-provisioenr 插件部署: 为了和alicloud-disk-controller区分，这里的storageclass都在名字后面加上-csi字样；
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-available-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-essd-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_essd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-ssd-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_ssd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-efficiency-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_efficiency
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-topology-csi
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-provisioner
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-provisioner
  replicas: 2
  template:
    metadata:
      labels:
        app: csi-provisioner
    spec:
      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            preference:
              matchExpressions:
              - key: node-role.kubernetes.io/master
                operator: Exists
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - csi-provisioner
              topologyKey: "kubernetes.io/hostname"
      tolerations:
      - effect: NoSchedule
        operator: Exists
        key: node-role.kubernetes.io/master
      - effect: NoSchedule
        operator: Exists
        key: node.cloudprovider.kubernetes.io/uninitialized
      priorityClassName: system-node-critical
      serviceAccount: admin
      hostNetwork: true
      containers:
        - name: external-disk-provisioner
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v3.0.0-3f86569-aliyun
          args:
            - "--csi-address=$(ADDRESS)"
            - "--feature-gates=Topology=True"
            - "--volume-name-prefix=disk"
            - "--strict-topology=true"
            - "--timeout=150s"
            - "--leader-election=true"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-attacher
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-attacher:v2.1.0
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-resizer
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-resizer:v0.3.0
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-nas-provisioner
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v1.4.0-aliyun
          args:
            - "--provisioner=nasplugin.csi.alibabacloud.com"
            - "--csi-address=$(ADDRESS)"
            - "--volume-name-prefix=nas"
            - "--timeout=150s"
            - "--enable-leader-election=true"
            - "--leader-election-type=leases"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
        - name: external-csi-snapshotter
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-snapshotter:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          imagePullPolicy: Always
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /csi
        - name: external-snapshot-controller
          image: registry.cn-hangzhou.aliyuncs.com/acs/snapshot-controller:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--leader-election=true"
          imagePullPolicy: Always
        - name: csi-provisioner
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          args:
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--v=2"
            - "--driver=nas,disk"
          env:
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-provisioner/driverplugin.csi.alibabacloud.com-replace/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
            - name: SERVICE_TYPE
              value: "provisioner"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            timeoutSeconds: 5
            failureThreshold: 5
          ports:
            - name: healthz
              containerPort: 11270
              protocol: TCP
          volumeMounts:
            - name: host-dev
              mountPath: /dev
              mountPropagation: "HostToContainer"
            - name: host-log
              mountPath: /var/log/
            - name: etc
              mountPath: /host/etc
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
            - mountPath: /var/addon
              name: addon-token
              readOnly: true
          resources:
            limits:
              cpu: 1000m
              memory: 1000Mi
            requests:
              cpu: 100m
              memory: 100Mi
      volumes:
        - name: disk-provisioner-dir
          emptyDir: {}
        - name: nas-provisioner-dir
          emptyDir: {}
        - name: host-log
          hostPath:
            path: /var/log/
        - name: host-dev
          hostPath:
            path: /dev
        - name: addon-token
          secret:
            defaultMode: 420
            items:
            - key: addon.token.config
              path: token-config
            secretName: addon.csi.token
        - name: etc
          hostPath:
            path: /etc
```

### 使用 csi 插件部署应用

使用上面模板部署了 csi 插件后，使用 csi storageclass 创建测试用 statefulset, 验证 csi 插件是否部署成功；

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web-csi-tr
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: disk-csi-tr
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: disk-csi-tr
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "alicloud-disk-available-csi"
      resources:
        requests:
          storage: 20Gi
```


### 变更 flexvolume 的 PV persistentVolumeReclaimPolicy 为 Retain

```kubectl edit pv d-8vb61ll1g3ykxojuvdg4```
```yaml
    persistentVolumeReclaimPolicy: Retain
    storageClassName: alicloud-disk-efficiency
    volumeMode: Filesystem
```

### 删除 flexvolume 版本的 Statefulset

```kubectl delete statefulset web-multi```

### 为 云盘 ```d-8vb61ll1g3ykxojuvdg4``` 打快照

为了保证数据安全，此时需要在 ecs 上为云盘打快照。

### 删除 flexvolume 版本的 PVC/PV

```kubectl delete pvc disk-ssd-test-web-multi-0```
```kubectl delete pv d-8vb61ll1g3ykxojuvdg4```

### 创建 CSI 版本的 PV, PVC

手动创建 csi 版本的同名 PVC/PV yaml

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: d-8vb61ll1g3ykxojuvdg4
  labels:
    alicloud-pvname: disk-ssd-test-web-multi-0
spec:
  storageClassName: "alicloud-disk-available-csi"
  capacity:
    storage: 20Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  csi:
    driver: diskplugin.csi.alibabacloud.com
    volumeHandle: d-8vb61ll1g3ykxojuvdg4
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: disk-ssd-test-web-multi-0
spec:
  storageClassName: "alicloud-disk-available-csi"
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi
  volumeName: d-8vb61ll1g3ykxojuvdg4
```

```kubectl apply -f -```

### 变更 flexvolume 版本的 statefulset yaml

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web-multi
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: disk-ssd-test
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: disk-ssd-test
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "alicloud-disk-efficiency-csi"
      resources:
        requests:
          storage: 20Gi
```
```kubectl apply -f - ```

### 删除 flexvolume & alicloud-disk-controller 资源

确认所有 flexvolume 相关资源都已经删除之后，删除 flexvolume 系列组件, 以及 相关sc

> 如果您的集群有多个使用flexvolume插件部署的应用，需要将所有应用迁移后再删除!


```kubectl delete deploy alicloud-disk-controller -nkube-system```

```kubectl delete ds flexvolume -nkube-system```

```kubectl delete sc alicloud-disk-available alicloud-disk-efficiency alicloud-disk-essd alicloud-disk-ssd```
### 修改 kubelet 参数
#### 使用daemonSet
```yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: kubelet-set
spec:
  selector:
    matchLabels:
      app: kubelet-set
  template:
    metadata:
      labels:
        app: kubelet-set
    spec:
      tolerations:
        - operator: "Exists"
      hostNetwork: true
      hostPID: true
      containers:
        - name: kubelet-set
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          env:
          - name: enableADController
            value: "true"
          command: ["sh", "-c"]
          args:
          - echo "Starting kubelet flag set to $enableADController";
            ifFlagTrueNum=`cat /host/etc/systemd/system/kubelet.service.d/10-kubeadm.conf | grep enable-controller-attach-detach=$enableADController | grep -v grep | wc -l`;
            echo "ifFlagTrueNum is $ifFlagTrueNum";
            if [ "$ifFlagTrueNum" = "0" ]; then
                curValue="true";
                if [ "$enableADController" = "true" ]; then
                    curValue="false";
                fi;
                sed -i "s/enable-controller-attach-detach=$curValue/enable-controller-attach-detach=$enableADController/" /host/etc/systemd/system/kubelet.service.d/10-kubeadm.conf;
                restartKubelet="true";
                echo "current value is $curValue, change to expect "$enableADController;
            fi;
            if [ "$restartKubelet" = "true" ]; then
                /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload;
                /nsenter --mount=/proc/1/ns/mnt service kubelet restart;
                echo "restart kubelet";
            fi;
            while true;
            do
                sleep 5;
            done;
          volumeMounts:
          - name: etc
            mountPath: /host/etc
      volumes:
        - name: etc
          hostPath:
            path: /etc
```
> 确认kubelet 更改成功后即可删除 ds， 但是如果做了集群扩容，则需要重新跑一遍 daemonset， 如果一直保留不删除ds，当扩容时，会自动到新增节点上去执行
### 修改 csidriver 
- 先删除当前 csidriver
```kubectl delete csidriver diskplugin.csi.alibabacloud.com```
- 使用以下 yaml 部署新的csidriver
```yaml
apiVersion: v1
items:
- apiVersion: storage.k8s.io/v1
  kind: CSIDriver
  metadata:
    name: diskplugin.csi.alibabacloud.com
  spec:
    attachRequired: true
    podInfoOnMount: true
    volumeLifecycleModes:
    - Persistent
kind: List
```
### 验证csi是否升级成功
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web-csi
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: disk-csi
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: disk-csi
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "alicloud-disk-available-csi"
      resources:
        requests:
          storage: 20Gi
```
> 看到这个 pod 正常启动，则代表 flexvolume 迁移成功