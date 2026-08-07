# 用本地盘为云盘加速（Data Cache）

云盘 CSI 驱动可以将高速的**本地盘**置于较慢的**云盘**之前，以加速 I/O。该能力基于
Linux 设备映射（device-mapper）的 `dm-cache` 目标实现：云盘作为持久化的 *origin*
（源设备），节点本地存储上的一个文件作为缓存。Pod 的 I/O 会尽可能由本地缓存提供服务，
而数据最终持久化在云盘上。

该功能是标准云盘 CSI 驱动（`diskplugin.csi.alibabacloud.com`）的一项能力，**并非**
独立的插件。

## 适用场景

在以下情况下使用本地盘缓存：

- 工作负载以读为主，或存在可放入本地盘的热点数据集；
- 运行节点具备高速本地存储（如本地 NVMe SSD）；
- 可以接受缓存是节点本地的（缓存卷迁移到其他节点后需要重新建立缓存）。

对于延迟不敏感或一次性写入的负载，缓存带来的收益有限。

## 工作原理

1. **创建卷** — `StorageClass` 携带 `dataCacheSize`（以及可选的 `dataCacheMode`）。
   这些参数在 `CreateVolume` 阶段校验，并保存到 PV 的 `VolumeContext` 中，从而传递到
   节点。
2. **节点挂载（stage/publish）** — 卷在节点上挂载时，驱动会：
   - 在节点本地存储的 `/var/alibaba-cloud-csi/data-cache/` 下 `fallocate` 两个后备
     文件：`<volumeID>.data`（缓存数据，大小由 `dataCacheSize` 决定）和
     `<volumeID>.meta`（dm-cache 元数据，大小自动计算）；
   - 将每个文件绑定到一个 loop 设备；
   - 创建名为 `csi-datacache-<volumeID>` 的 device-mapper `cache` 设备，
     暴露为 `/dev/mapper/csi-datacache-<volumeID>`，以云盘作为源设备；
   - 格式化/挂载 **dm-cache 设备**（而非原始云盘），使所有 Pod I/O 流经本地缓存。
3. **持久化与崩溃安全** — 只要后备文件可能仍持有未回写（writeback）的数据，就不会被
   截断或删除。节点重启后，dm-cache 超级块会被重新打开（而非重新格式化），因此脏块会
   在下一次卸载时回写到云盘。
4. **扩容** — 在 `NodeExpandVolume` 时，dm-cache 目标会先扩展以覆盖扩容后的云盘，
   然后再扩容文件系统。
5. **卸载（teardown）** — 卷卸载时，驱动将缓存切换到 `cleaner` 策略，把所有脏块回写
   到云盘，随后移除 dm-cache 设备并删除后备文件。

## 前置条件

- **节点**必须为 Linux，且具备 device-mapper（`/dev/mapper/control`）、loop 设备，
  以及 `dm-cache` 内核目标。
- 节点插件需要将宿主机 `/var/alibaba-cloud-csi` 挂载进容器（Helm chart 默认已配置）。
  缓存的 `.data`/`.meta` 文件存放于此。
- 节点需要有足够的本地磁盘剩余空间，以容纳每个调度到该节点的缓存卷的 `dataCacheSize`。

如果节点上 device-mapper 不可用，或缓存文件无法创建，驱动会**回退（fallback）** 到
直接使用原始云盘，并在 Pod 上产生 `DataCacheFallback` 警告事件。此时卷仍可正常工作，
只是没有加速效果。

## StorageClass 参数

| 参数             | 是否必填         | 取值                          | 默认值         | 说明 |
|------------------|------------------|-------------------------------|----------------|------|
| `dataCacheSize`  | 是（用于开启）   | Kubernetes 数量，如 `10Gi`    | —              | 本地缓存大小。非零值即**开启**本地盘缓存。 |
| `dataCacheMode`  | 否               | `writethrough`、`writeback`   | `writethrough` | 缓存写模式，见下文。 |

校验规则：

- `dataCacheSize` 必须能解析为数量；无法识别的 `dataCacheMode` 会被拒绝。
- 设置了 `dataCacheMode` 但 `dataCacheSize` 为零，会报错。
- 设置了 `dataCacheSize` 但未指定 `dataCacheMode` 时，模式默认为 `writethrough`。

### 写模式

- **`writethrough`（默认，推荐）** — 写操作在返回成功前会同时写入缓存与云盘。加速读；
  由于云盘始终是最新的，对节点故障是安全的。
- **`writeback`** — 写操作在数据落入本地缓存后即返回成功，随后再异步回写到云盘。可同时
  加速写，但在回写之前，脏数据仅存在于节点本地盘上。**如果在回写前节点本地盘丢失，
  这部分数据将丢失。** 仅在本地盘可靠性满足需求且确需写加速时使用。

## 使用步骤

### 1. 创建 StorageClass

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: alicloud-disk-datacache
provisioner: diskplugin.csi.alibabacloud.com
parameters:
  type: cloud_essd
  dataCacheSize: "10Gi"
  dataCacheMode: "writethrough"   # 或 "writeback"
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
```

```shell
kubectl apply -f storageclass.yaml
```

### 2. 创建 PVC

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: datacache-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: alicloud-disk-datacache
  resources:
    requests:
      storage: 100Gi
```

```shell
kubectl apply -f pvc.yaml
```

### 3. 在工作负载中挂载

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: datacache-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: datacache-demo
  template:
    metadata:
      labels:
        app: datacache-demo
    spec:
      containers:
        - name: app
          image: nginx:1.25
          volumeMounts:
            - name: data
              mountPath: /data
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: datacache-pvc
```

```shell
kubectl apply -f deploy.yaml
```

## 验证缓存是否生效

在 Pod 所在的节点上执行：

```shell
# 该卷对应的 dm-cache 设备应存在
ls -l /dev/mapper/csi-datacache-<volumeID>

# 查看缓存表与运行状态
dmsetup status csi-datacache-<volumeID>

# 后备文件
ls -lh /var/alibaba-cloud-csi/data-cache/
```

`dmsetup status` 会报告缓存块使用量、命中/未命中计数、脏块数量、io 模式，以及当前策略
（正常运行时为 `mq`）。

如果缓存**未**生效，检查是否存在 `DataCacheFallback` 警告事件：

```shell
kubectl describe pod <pod-name>
```

并确认节点满足[前置条件](#前置条件)。

## 扩容

支持卷扩容。请确保 `StorageClass` 设置了 `allowVolumeExpansion: true`，然后修改 PVC
请求的容量。驱动会扩容云盘、同步扩展 dm-cache 目标，并扩容文件系统。

## 注意事项与限制

- 缓存是**节点本地**的。当 Pod（及其 `ReadWriteOnce` 卷）迁移到新节点时，会在新节点
  重新建立缓存；原节点的缓存文件会在卸载时清理。
- 使用 `writeback` 时，未回写的数据在回写前仅存在于节点本地盘上。除非确需写加速并接受
  相应的可靠性权衡，否则优先使用 `writethrough`。
- 节点上需为每个缓存卷预留足够的本地磁盘空间。
- 需要具备 `dm-cache` 内核目标的 Linux 节点；否则驱动会透明回退到直接使用原始云盘。

## 相关文档

- [云盘](./disk.md)
- [云盘——扩容](./disk-resizer.md)
- [安装](./install.md)
