# BMCPFS Helm 集成指南

## 概述

本指南提供了在 Kubernetes 环境中通过 Helm charts 部署和使用智算版 CPFS 的全面说明。

## 功能特性

- **高性能并行文件系统**：针对高 I/O 需求工作负载进行优化
- **双网络支持**：VSC 提供高性能，VPC 提供兼容性
- **共享访问**：支持 ReadWriteMany 多节点并发访问
- **灵骏基础设施优化**：自动节点类型检测和网络优化

## 前置条件

### 基础设施要求

1. **Kubernetes 集群**：版本 1.20 或更高
2. **Helm**：版本 3.0 或更高
3. **网络基础设施**：
   - 具有 VSC 网络的灵骏基础设施（推荐，获得最佳性能）
   - 标准 VPC 基础设施（备选支持）
4. **CPFS 文件系统**：在阿里云中预创建的云并行文件系统, 提交工单开启 bmcpfs openapi.

### 认证要求

- 访问阿里云 API 的适当 IAM 角色， 并且为角色配置可以访问 NAS 和 EFLO 服务的权限， 所需权限列表如下, 拥有该权限的秘钥需要与 csi-provisioner 相关联。 
```
{
  "Version": "1",
  "Statement": [
    {
      "Action": [
        "nas:DescribeFileSystems",
        "nas:DescribeMountTargets",
        "nas:CreateMountTarget",
        "nas:DeleteMountTarget",
        "nas:ModifyMountTarget",
        "nas:DescribeProtocolMountTarget",
        "nas:AttachVscToFilesystems",
        "nas:DetachVscFromFilesystems",
        "nas:DescribeFilesystemsVscAttachInfo",
        "nas:DescribeAccessPoints"
      ],
      "Resource": [
        "*"
      ],
      "Effect": "Allow"
    },
    {
      "Effect": "Allow",
      "Action": [
        "eflo:DescribeVsc",
        "eflo:ListVscs"
      ],
      "Resource": "*"
    }
  ]
}
```

## 安装

### 步骤 1：准备配置

为您的环境创建自定义 values 文件。对于灵骏基础设施部署：

```bash
# 基于灵骏配置文件创建 values-custom.yaml
cat > values-custom.yaml << EOF
csi:
  bmcpfs:
    enabled: true
    controller:
      enabled: true
  disk:
    enabled: false  # 禁用其他存储类型
  nas:
    enabled: false
  oss:
    enabled: false

deploy:
  regionID: "cn-hangzhou"  # 您的地域
  ramToken: v2
  ecs: false  # 用于灵骏裸金属
  accessKey:
    enabled: true
    secretName: alibaba-addon-secret # 阿里云密钥，用于访问openapi， 权限在上面已列出
    idKey: access-key-id # secret 中的access-key-id
    secretKey: access-key-secret # secret 中的access-key-secret
  newProvisionerToken: false
controller:
  enabled: true
  replicas: 2

images:
  controller:
    repo: acs/csi-plugin
    tag: "v1.33.4-576836c-aliyun"
  plugin:
    repo: acs/csi-plugin
    tag: "v1.33.4-576836c-aliyun"
EOF
```

### 步骤 2：安装 CSI 驱动

```bash
helm install bmcpfs-csi-driver \
  ./deploy/charts/alibaba-cloud-csi-driver \
  -f values-custom.yaml \
  --namespace kube-system
```

### 步骤 3：安装cnfs-nas-daemon 组件

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: cnfs-nas-daemon
  namespace: cnfs-system
spec:
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: cnfs-nas-daemon
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: cnfs-nas-daemon
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/version: 0.1.0
        helm.sh/chart: cnfs-nas-daemon-0.1.4
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: type
                    operator: NotIn
                    values:
                      - virtual-kubelet
      containers:
        - args:
            - '--socket=/run/cnfs/alinas-mounter.sock'
            - '--driver=alinas'
          image: >-
            registry-us-west-1-vpc.ack.aliyuncs.com/acs/csi-alinas:v2.0-2a209ad-aliyun # 镜像地址需要改成当前集群所在 region 地址
          imagePullPolicy: IfNotPresent
          name: mount-proxy
          resources:
            limits:
              cpu: '1'
            requests:
              cpu: 500m
              memory: 1Gi
          securityContext:
            privileged: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
            - mountPath: /dev
              mountPropagation: HostToContainer
              name: host-dev
            - mountPath: /var/lib/kubelet
              mountPropagation: Bidirectional
              name: kubelet-dir
            - mountPath: /run/cnfs
              name: cnfs-dir
            - mountPath: /etc/aliyun/cpfs
              name: etc-dir
              subPath: cpfs
            - mountPath: /etc/aliyun/alinas
              name: etc-dir
              subPath: alinas
            - mountPath: /var/log/aliyun/cpfs
              name: log-dir
              subPath: cpfs
            - mountPath: /var/log/aliyun/alinas
              name: log-dir
              subPath: alinas
            - mountPath: /var/run/alinas
              name: cnfs-dir
              subPath: alinas
            - mountPath: /var/run/cpfs
              name: cnfs-dir
              subPath: cpfs
            - mountPath: /var/run/efc
              name: cnfs-dir
              subPath: efc
            - mountPath: /etc/hosts
              name: hosts
      dnsPolicy: ClusterFirst
      hostIPC: true
      hostNetwork: true
      priorityClassName: system-node-critical
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: cnfs-nas-daemon
      serviceAccountName: cnfs-nas-daemon
      terminationGracePeriodSeconds: 30
      tolerations:
        - operator: Exists
      volumes:
        - hostPath:
            path: /var/lib/kubelet
            type: ''
          name: kubelet-dir
        - hostPath:
            path: /dev
            type: ''
          name: host-dev
        - hostPath:
            path: /run/cnfs
            type: DirectoryOrCreate
          name: cnfs-dir
        - hostPath:
            path: /etc/cnfs
            type: DirectoryOrCreate
          name: etc-dir
        - hostPath:
            path: /var/log/cnfs
            type: DirectoryOrCreate
          name: log-dir
        - hostPath:
            path: /run/cnfs/alinas.hosts
            type: FileOrCreate
          name: hosts
  updateStrategy:
    type: OnDelete
```

### 步骤 4：验证安装

```bash
# 检查 CSI 驱动注册
kubectl get csidriver bmcpfsplugin.csi.alibabacloud.com

# 验证控制器部署
kubectl get deployment -n kube-system -l app=csi-provisioner

# 检查节点插件 daemonset
kubectl get daemonset -n kube-system -l app=cnfs-nas-daemon

# 检查节点插件 cnfs-nas-daemon
kubectl get daemonset -n cnfs-system -l app=csi-plugin

# 验证 Pod 状态
kubectl get pods -n kube-system | grep csi
```

预期输出应显示：
- CSI 驱动 `bmcpfsplugin.csi.alibabacloud.com` 已注册
- 控制器 Pod 正在运行（默认 2 个副本）
- 节点插件 Pod 在所有节点上运行


## 使用方法

### 静态卷配置

对于预先存在的 CPFS 文件系统，创建静态 PersistentVolume：

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-static-pv
spec:
  capacity:
    storage: 500Gi
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "bmcpfs-xxxxxx"
    volumeAttributes:
      vscMountTarget: "cpfs-xxxxxxxxx-vsc.cn-hangzhou.cpfs.nas.aliyuncs.com"
      vpcMountTarget: "cpfs-xxxxxxxxx-vpc.cn-hangzhou.cpfs.nas.aliyuncs.com"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: bmcpfs-static-pvc
  namespace: default
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 500Gi
  volumeName: bmcpfs-static-pv
```

### 在 Pod 中使用 BMCPFS

部署使用 BMCPFS 卷的 Pod：

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: bmcpfs-test-pod
  namespace: default
spec:
  containers:
    - name: app
      image: nginx:latest
      volumeMounts:
        - name: bmcpfs-volume
          mountPath: /data
          mountPropagation: Bidirectional
      command:
        - "/bin/sh"
        - "-c"
        - "echo 'Hello BMCPFS' > /data/test.txt && tail -f /data/test.txt"
  volumes:
    - name: bmcpfs-volume
      persistentVolumeClaim:
        claimName: bmcpfs-pvc
  nodeSelector:
    # 针对灵骏节点以获得最佳性能（如果可用）
    node-type: lingjun
```

### 多 Pod 共享访问

BMCPFS 支持 ReadWriteMany，允许多个 Pod 同时访问同一卷：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bmcpfs-shared-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: shared-app
  template:
    metadata:
      labels:
        app: shared-app
    spec:
      containers:
        - name: app
          image: nginx:latest
          volumeMounts:
            - name: shared-storage
              mountPath: /shared
      volumes:
        - name: shared-storage
          persistentVolumeClaim:
            claimName: bmcpfs-pvc
```

## 网络类型检测

BMCPFS 驱动程序根据节点配置自动检测适当的网络类型：

### 灵骏节点
- **网络类型**：VSC
- **性能**：针对高吞吐量、低延迟访问进行优化

### 标准节点
- **网络类型**：VPC
- **兼容性**：标准云基础设施支持

## 挂载流程图

```
flowchart TD
    Start([卷挂载请求]) --> CheckNode["检查节点 ID 类型"]
    CheckNode --> |"以 'lingjun:' 开头"| LingjunFlow["灵骏节点流程"]
    CheckNode --> |"以 'common:' 开头"| CommonFlow["普通节点流程"]
    
    LingjunFlow --> GetVSC["获取主 VSC"]
    GetVSC --> CreateVSC{"VSC 是否存在？"}
    CreateVSC --> |否| CreateNewVSC["创建主 VSC"]
    CreateVSC --> |是| UseVSC["使用现有 VSC"]
    CreateNewVSC --> AttachCPFS["将 CPFS 附加到 VSC"]
    UseVSC --> AttachCPFS
    AttachCPFS --> VSCMount["使用 VSC 网络挂载"]
    
    CommonFlow --> VPCMount["使用 VPC 网络挂载"]
    
    VSCMount --> MountProxy["挂载代理 (Alinas)"]
    VPCMount --> MountProxy
    MountProxy --> FileSystem["已挂载文件系统"]
```

# 升级到新版本
```shell
helm upgrade bmcpfs-csi-driver \
  ./deploy/charts/alibaba-cloud-csi-driver \
  -f backup-values.yaml \
  --namespace kube-system
```

# 验证升级
```shell
kubectl rollout status deployment/csi-provisioner -n kube-system
kubectl rollout status daemonset/csi-plugin -n kube-system
```

## 最佳实践

### 性能优化

1. **使用 Lingjun 节点**: 在 Lingjun 裸金属节点上部署工作负载
2. **节点亲和性**: 配置 Pod 亲和性以调度到适当节点
3. **子路径策略**: 用于多租户场景的子路径
4. **挂载选项**: 根据工作负载模式优化挂载选项

### 高可用性

1. **控制器副本**: 运行多个控制器副本
2. **节点分布**: 确保 CSI 插件在所有节点上运行
3. **存储冗余**: 使用 CPFS 冗余功能
4. **监控**: 启用全面的监控和警报

### 资源管理

```
# 示例 CSI 组件资源限制
resources:
  controller:
    limits:
      cpu: 500m
      memory: 512Mi
    requests:
      cpu: 100m
      memory: 128Mi
  plugin:
    limits:
      cpu: 200m
      memory: 256Mi
    requests:
      cpu: 50m
      memory: 64Mi
```

## 卸载

```bash
# 首先删除所有 PVCs
kubectl delete pvc --all -A

# 卸载 Helm 发布
helm uninstall bmcpfs-csi-driver -n kube-system

# 卸载 cnfs-nas-daemon 组件
kubectl delete daemonset cnfs-nas-daemon -n cnfs-system

# 清理 CSI 驱动注册
kubectl delete csidriver bmcpfsplugin.csi.alibabacloud.com

# 验证清理
kubectl get pods -n kube-system | grep csi
```

## 共存

BMCPFS 可以与其他 CSI 驱动程序共存：

```yaml
csi:
  bmcpfs:
    enabled: true
  nas:
    enabled: true  # 保留 NAS 以保持兼容性
  disk:
    enabled: true  # 保留磁盘用于块存储
```

## 支持和社区

- **问题反馈**：在 [GitHub Issues](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/issues) 报告问题
- **文档**：参考主要[项目文档](../README.md)
- **社区**：加入钉钉群（群号：33936810）进行讨论

## 安全

向 kubernetes-security@service.aliyun.com 报告安全漏洞

## 许可证

本项目根据 Apache License 2.0 许可 - 详见 [LICENSE](../LICENSE) 文件。