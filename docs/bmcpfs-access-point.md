# BMCPFS Access Point（AP）挂载使用文档

本文档介绍如何使用 bmcpfs 驱动（`bmcpfsplugin.csi.alibabacloud.com`）通过 CPFS Access Point 挂载文件系统，包括无鉴权挂载、RAM 鉴权（静态 AK / 可轮转 STS）两类场景。

## 功能概述

| 能力 | 说明 |
| --- | --- |
| AP 挂载 | 通过 `accessPointId` 指定接入点，支持 tcp 与 vsc 两种网络类型（驱动按节点类型自动选择） |
| 无鉴权挂载 | 不配置 `nodePublishSecretRef`，仅通过 AP 挂载 |
| RAM 鉴权：静态 AK | Secret 含 AK/SK 两键，写入 `g_unas_AKFile`；不支持热更新 |
| RAM 鉴权：STS 轮转 | Secret 含 STS 三元组，支持外部轮转，写入 `g_unas_STSFile`；EFC 客户端每 10 分钟自动重读并刷新签名 |
| 鉴权模式自动识别 | 无需 `authType` 参数，驱动按 `nodePublishSecretRef` 的 Secret 键集合自动识别 AK / STS 模式 |
| 单 Pod 多 AP | 同一 Pod 可同时挂载同一文件系统的多个不同 AP（每个 AP 一个 PV） |

STS 轮转基于 Kubernetes 社区的 `CSIDriver requiresRepublish` 机制：kubelet 周期性重新调用 NodePublishVolume 并携带最新 Secret 内容，驱动检测到变化后原子更新 STS 配置文件，EFC 客户端在下一个 10 分钟周期内加载新凭证。全程无需重启 Pod、不中断挂载。

## 前提条件

1. **CSI 组件版本**：安装包含 bmcpfs AP 支持的 csi-plugin / csi-provisioner 版本，helm values 中启用：

   ```yaml
   csi:
     bmcpfs:
       enabled: true
   ```

   bmcpfs 的 CSIDriver 对象默认渲染 `requiresRepublish: true`（STS 轮转依赖），无需额外开关。

2. **EFC 客户端版本**：节点上的 EFC 客户端需支持 `g_unas_Accesspoint`、`g_unas_AKFile`、`g_unas_STSFile` 挂载参数。
3. **fileserver 集群配置**：CPFS fileserver 集群需设置 flag `efc_pov_UmmSigningRegion=<当前 region>`（RAM 鉴权签名校验的服务端前置条件，请联系 CPFS 侧配置）。
4. **AP 已创建**：通过 NAS 控制台/OpenAPI 预先创建 Access Point，获得 `ap-` 开头的 AP ID。

## 卷配置规范

### volumeHandle 唯一性要求（重要）

同一 Pod 引用同一文件系统的多个 AP 时，kubelet 按 `volumeHandle` 对卷去重。因此**每个 AP PV 的 volumeHandle 必须唯一**，格式为：

```
volumeHandle: "<bmcpfsId>+<唯一后缀>"
# 推荐用 AP ID 作后缀: cpfs-0123456789+ap-aaaaaaaa
# 存量 fileset 卷的 <bmcpfsId>+<filesetId> 格式继续兼容，无需变更
```

第一段必须为文件系统 ID（驱动据此执行 attach）；后缀仅用于保证唯一性，内容不限（推荐 AP ID，便于运维对应）。实际挂载的 AP 以 `volumeAttributes.accessPointId` 为准，与后缀无强制关联。

### volumeAttributes

| 键 | 必填 | 说明 |
| --- | --- | --- |
| `vpcMountTarget` | 与 vsc 二选一或同时配置 | VPC 挂载点域名（tcp 链路使用） |
| `vscMountTarget` | 与 vpc 二选一或同时配置 | VSC 挂载点域名（vsc 链路使用） |
| `accessPointId` | AP 场景必填 | Access Point ID，非空即启用 AP 挂载 |

### 鉴权模式识别与校验规则

驱动不使用 `authType` 参数，而是根据 `nodePublishSecretRef` 引用的 Secret 键集合自动识别模式：

| Secret 内容 | 识别结果 |
| --- | --- |
| 未配置 `nodePublishSecretRef` | 无鉴权挂载 |
| 恰好含 `accessKeyId`、`accessKeySecret` | AK 模式 |
| 含 `accessKeyId`、`accessKeySecret`、`securityToken`（可选 `expiration`） | STS 模式 |
| 其它形状（缺键、空值、未知键） | 拒绝挂载（`InvalidArgument`） |

注意：

- 键集合为严格白名单匹配，拼写错误（如 `security_token`）会被直接拒绝，不会静默降级为 AK 模式或匿名挂载。
- 配置了 Secret 时 `accessPointId` 必须非空（RAM 鉴权仅对 AP 挂载生效）。
- 鉴权模式在首次挂载时定格：挂载后若 Secret 形状变化（如 AK 改为三元组），驱动告警并忽略，需重建 Pod 重新挂载生效。
- 不支持在 PV `mountOptions` 中手工填写 `g_unas_AKFile` / `g_unas_STSFile`，此类选项会被驱动剥离并告警；凭证文件路径由驱动统一管理。

## 使用方式

### 用例 1：无鉴权 AP 挂载

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: bmcpfs-ap-pvc
  namespace: default
spec:
  accessModes: ["ReadWriteMany"]
  storageClassName: ""
  resources:
    requests:
      storage: 500Gi
  volumeName: bmcpfs-ap-pv
```

驱动生成的挂载命令等价于（tcp 链路）：

```bash
mount -t alinas -o efc,protocol=efc,net=tcp,fstype=cpfs \
  -o g_unas_Accesspoint=ap-aaaaaaaa \
  cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com:/ <targetPath>
```

> 测试阶段 AP 选项名为 `g_unas_Accesspoint`；GA 后将切换为 `accesspoint`。切换由节点插件环境变量 `BMCPFS_AP_OPTION_STYLE`（`legacy` / `ga`，默认 `legacy`）控制，业务 PV 配置无需变化。

### 用例 2：RAM 鉴权——静态 AK

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: bmcpfs-ak-secret
  namespace: kube-system
type: Opaque
stringData:
  accessKeyId: "LTAI5t********"
  accessKeySecret: "********"
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-ak-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
    nodePublishSecretRef:
      name: bmcpfs-ak-secret
      namespace: kube-system
```

驱动检测到 Secret 仅含 AK/SK 两键，自动按 AK 模式挂载。

**注意**：AK 配置不支持热更新。修改 Secret 中的 AK 后，需要重建使用该卷的 Pod 触发重新挂载才能生效。

### 用例 3：RAM 鉴权——STS 轮转

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: bmcpfs-sts-secret
  namespace: kube-system
type: Opaque
stringData:               # 由外部系统周期性轮转更新
  accessKeyId: "STS.xxx"
  accessKeySecret: "********"
  securityToken: "********"
  expiration: "2026-08-10T12:00:00Z"   # 可选，仅用于告警
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: bmcpfs-ap-sts-pv
spec:
  capacity:
    storage: 500Gi
  accessModes: ["ReadWriteMany"]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: bmcpfsplugin.csi.alibabacloud.com
    volumeHandle: "cpfs-0123456789+ap-aaaaaaaa"
    volumeAttributes:
      vpcMountTarget: "cpfs-0123456789-vpc.cn-hangzhou.cpfs.aliyuncs.com"
      vscMountTarget: "cpfs-0123456789-vsc.cn-hangzhou.cpfs.aliyuncs.com"
      accessPointId: "ap-aaaaaaaa"
    nodePublishSecretRef:
      name: bmcpfs-sts-secret
      namespace: kube-system
```

驱动检测到 Secret 含 `securityToken`，自动按 STS 模式挂载。

**STS 轮转说明**：

- 外部系统只需更新 Secret 中的三元组，无需操作节点或 Pod。
- Secret 中**不需要** md5 字段；驱动写入 STS 配置文件时自动计算 `md5(accessKeyId + accessKeySecret + securityToken)`。
- 新凭证生效链路：Secret 更新 → kubelet 周期性 republish 带入 → 驱动原子重写 STS 文件 → EFC 下一个 10 分钟周期加载。**端到端生效延迟上限约为 republish 间隔 + 10 分钟**，请保证轮转时新 STS 的剩余有效期充分大于该窗口（建议 ≥ 30 分钟）。
- 同一 PV 在同一节点上的多个 Pod 共享同一份凭证文件（凭证为 PV 级配置）。

### 用例 4：单 Pod 挂载同一文件系统的多个 AP

为每个 AP 各建一个 PV/PVC（volumeHandle 后缀不同即可），Pod 同时引用：

```yaml
# PV/PVC 定义同用例 1~3，两个 PV 的 volumeHandle 分别为：
#   cpfs-0123456789+ap-aaaaaaaa
#   cpfs-0123456789+ap-bbbbbbbb
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bmcpfs-multi-ap-app
spec:
  replicas: 2
  selector:
    matchLabels: {app: bmcpfs-multi-ap-app}
  template:
    metadata:
      labels: {app: bmcpfs-multi-ap-app}
    spec:
      containers:
        - name: app
          image: registry.cn-hangzhou.aliyuncs.com/acs/busybox:latest
          command: ["sleep", "infinity"]
          volumeMounts:
            - {name: data-a, mountPath: /data-a}
            - {name: data-b, mountPath: /data-b}
      volumes:
        - name: data-a
          persistentVolumeClaim: {claimName: bmcpfs-ap-a-pvc}
        - name: data-b
          persistentVolumeClaim: {claimName: bmcpfs-ap-b-pvc}
```

两个 AP 各自独立挂载、独立鉴权（可引用不同 Secret），互不影响。

## 行为说明与约束

### 卸载与 attach 保留行为

同一文件系统的多个 AP/fileset PV 在同一节点上共享一条 CPFS↔VSC attach。为保证多 AP 场景卸载安全，当前版本对 volumeHandle 带 `+` 后缀的卷（AP 卷、fileset 卷）在卸载时**不执行 detach**：

- 卷卸载（Pod 删除、PVC 删除）正常完成，不会误断同节点其它 AP 的挂载。
- 文件系统与节点 VSC 之间的 attach 会保留，随 VSC / 节点回收自动清理。
- 如需提前回收（如触及 VSC attach 数量上限），可在确认该节点无该文件系统任何活跃挂载后，通过 NAS OpenAPI `DetachVscFromFilesystems` 手工 detach；也可使用 `SKIP_BMCPFS_DETACH` 环境变量做全局控制。
- 后续版本计划通过增强的 external-attacher 方案恢复标准的按引用 detach 行为，届时该约束将移除，PV 配置无需变化。

### 其它约束

| 约束 | 说明 |
| --- | --- |
| accessModes | AP 卷推荐 `ReadWriteMany` |
| AK 热更新 | 不支持；修改 AK 需重建 Pod |
| 凭证文件 | 由驱动管理于节点 `/run/cnfs/efc-credentials/<volumeId>/`（目录 0700、文件 0600，位于 tmpfs），卷卸载后自动清理，请勿手工修改 |
| Secret 内容 | 驱动不校验凭证与 AP 权限的匹配性；权限不足表现为挂载失败或 IO 报错，请核对 AP 的 RAM 策略 |

## 故障排查

| 现象 | 排查方向 |
| --- | --- |
| 挂载失败，事件含 `InvalidArgument` | 核对 Secret 键集合是否符合 AK / STS 白名单（拼写、多余键、空值）、`accessPointId` 是否配置 |
| 多 AP 只挂上一个 / 挂错 AP | 检查各 PV 的 volumeHandle 是否唯一（kubelet 按 handle 去重，重复时只会挂其中一个） |
| 挂载失败，EFC 报鉴权错误 | 确认 fileserver 已配置 `efc_pov_UmmSigningRegion`；核对 AK/STS 是否有效、AP RAM 策略是否授权 |
| 轮转后新 STS 未生效 | 依次确认：Secret 已更新 → 节点上 `/run/cnfs/efc-credentials/<volumeId>/sts.json` 内容已更新（驱动侧）→ 等待 EFC 10 分钟读取周期（客户端侧）。若文件未更新，检查 CSIDriver 对象是否含 `requiresRepublish: true` |
| STS 过期导致 IO 失败 | 检查外部轮转系统是否停止更新 Secret；确认轮转周期满足"剩余有效期 ≥ republish 间隔 + 10 分钟"的要求 |
| 驱动日志 | 节点侧：csi-plugin DaemonSet Pod（bmcpfs 相关日志）；挂载执行侧：alinas mount-proxy 日志 |

验证 STS 文件已更新的快捷方式（节点上执行）：

```bash
ls -l --time-style=full-iso /run/cnfs/efc-credentials/<volumeId>/sts.json   # 看修改时间
```
