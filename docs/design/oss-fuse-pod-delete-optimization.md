# OSS Fuse Pod Delete 操作优化设计文档

## 0. 关键发现

### ResourceVersion 覆盖机制

在实施过程中发现一个关键问题：**Reflector 会覆盖 ListFunc 外设置的 ResourceVersion**。

**错误做法**（不会生效）：
```go
// 这样设置 ResourceVersion 不会生效！
options := metav1.ListOptions{ResourceVersion: ""}
lw := &cache.ListWatch{
    ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
        // opts.ResourceVersion 已经被 Reflector 覆盖为 "0" 了
        return podClient.List(ctx, opts)
    },
}
```

**正确做法**（在 ListFunc 内部覆盖）：
```go
lw := &cache.ListWatch{
    ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
        // 必须在 ListFunc 内部覆盖！
        options.ResourceVersion = ""
        return podClient.List(ctx, options)
    },
}
```

**原因**：
1. `NewSharedIndexInformer` 内部创建 `Reflector`
2. `Reflector.list()` 方法创建新的 `options := metav1.ListOptions{ResourceVersion: r.relistResourceVersion()}`
3. 新 informer 的 `relistResourceVersion()` 返回 `"0"`（为了从 watch cache 读取）
4. Reflector 调用 `pager.ListWithAlloc(options)`，传入的是它创建的 options
5. pager 再调用我们的 ListFunc(opts)，此时 `opts.ResourceVersion` 已经是 `"0"`
6. **因此必须在 ListFunc 内部覆盖**，才能强制从 etcd 读取

**参考代码**：
- `vendor/k8s.io/client-go/tools/cache/reflector.go:504` - Reflector 创建 options
- `vendor/k8s.io/client-go/tools/cache/reflector.go:912-915` - relistResourceVersion 返回 "0"
- `vendor/k8s.io/client-go/tools/cache/reflector.go:547` - 调用 pager.ListWithAlloc

---

## 1. 背景

### 1.1 问题描述

在大规模集群测试中，OSS 挂载/卸载场景下，CSI Controller 的 `Delete` 操作频繁使用带 label 的 Pod List 请求，导致 API Server 压力过大。同时存在数据一致性风险。

### 1.2 影响范围

- **组件**: `alibaba-cloud-csi-driver` Controller
- **代码位置**: 
  - [pkg/oss/oss.go:71-78](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss/oss.go#L71-L78) - fusePodManagers 初始化
  - [pkg/mounter/fuse_pod_manager/fuse_pod_manager.go:333-410](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/fuse_pod_manager.go#L333-L410) - Delete 方法实现
  - [deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml:631-632](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml#L631-L632) - QPS/Burst 配置

### 1.3 资源约束

根据 Helm chart 配置，CSI Controller 的资源限制为：
```yaml
resources:
  limits:
    cpu: 500m
    memory: 1024Mi    # 硬限制，OOM 会导致 Pod 重启
  requests:
    cpu: 100m
    memory: 128Mi
```

**关键约束**: 组件对内存限制敏感，任何增加内存使用的方案都需要谨慎评估。

---

## 2. 问题分析

### 2.1 当前实现

在 `fuse_pod_manager.go:339-380` 的 Delete 方法中：

```go
func (fpm *FusePodManager) Delete(c *FusePodContext) error {
    // 1. 创建 informer（每次 Delete 都创建新的）
    informer := informercorev1.NewFilteredPodInformer(fpm.client, c.Namespace, 0, nil, func(options *metav1.ListOptions) {
        options.FieldSelector = listOptions.FieldSelector
        options.LabelSelector = listOptions.LabelSelector
    })
    
    // 2. 启动 informer 并等待同步
    go informer.Run(ctx.Done())
    synced := cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
    
    // 3. 遍历缓存删除 pod
    for _, obj := range informer.GetStore().List() {
        podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
    }
    
    // 4. 等待所有 pod 删除完成
    for {
        select {
        case <-deleteNotify:
            n := len(informer.GetStore().ListKeys())
            if n == 0 {
                return nil  // 返回成功
            }
        }
    }
}
```

### 2.2 核心问题

#### 问题 1: ResourceVersion 机制导致的数据不一致风险

**Kubernetes Informer 的 List 行为**（参考 `vendor/k8s.io/client-go/tools/cache/reflector.go:902-918`）：

```go
func (r *Reflector) relistResourceVersion() string {
    if r.isLastSyncResourceVersionUnavailable {
        return ""  // 从 etcd 一致性读取
    }
    if r.lastSyncResourceVersion == "" {
        return "0"  // 初始 list 使用 "0"，可能从 watch cache 读取
    }
    return r.lastSyncResourceVersion
}
```

**关键机制**:
- `ResourceVersion="0"`: 从 watch cache 读取（不保证最新，watch cache 可能有延迟）
- `ResourceVersion=""`: 从 etcd 直接读取（强一致性）
- `ResourceVersion="N"`: 从 watch cache 读取至少 RV=N 的数据

**风险场景**:
```
T1: Pod A 被创建（但还没进入 watch cache）
T2: Delete 操作创建 informer，初始 list 使用 RV="0"
T3: 从 watch cache 读取，没拿到 Pod A（watch cache 延迟）
T4: informer.GetStore().List() 返回空
T5: 直接认为已删除，返回成功
结果: Pod A 实际上还在运行，但 Delete 已经返回成功
```

#### 问题 2: 大量并发 Delete 创建独立 informer

- 每次 Delete 操作都创建新的 informer
- 每个 informer 都会触发 List + Watch 请求
- 大规模集群中并发 Delete 会导致 API Server 压力骤增

---

## 3. 解决方案

### 方案 1: 不复用 client，每次创建新的 HTTP 连接

#### 3.1.1 实现思路

在每次 Create/Delete 操作内部创建独立的 `kubernetes.Clientset`，而不是复用 `fusePodManagers` 中的 client。

```go
func (fpm *FusePodManager) Delete(c *FusePodContext) error {
    // 每次 Delete 创建新的 clientset
    cfg, _ := options.GetRestConfig()
    clientset := kubernetes.NewForConfigOrDie(cfg)
    podClient := clientset.CoreV1().Pods(c.Namespace)
    
    // 使用新 client 创建 informer
    informer := informercorev1.NewFilteredPodInformer(clientset, ...)
    // ...
}
```

#### 3.1.2 技术细节

**QPS/Burst 限制机制**（参考 `vendor/k8s.io/client-go/kubernetes/clientset.go:504-511`）：

```go
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
    configShallowCopy := *c
    if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
        // 每次 NewForConfigOrDie 都会创建独立的 RateLimiter
        configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(
            configShallowCopy.QPS, 
            configShallowCopy.Burst
        )
    }
    // ...
}
```

**HTTP Transport 行为**（参考 `vendor/k8s.io/client-go/rest/transport.go:32`）：
- `HTTPClientFor` 会创建新的 `http.Client`
- Go 的 http.Client 默认使用连接池，同域名连接可能复用
- 但每个 clientset 有独立的 Transport 实例

#### 3.1.3 风险评估

| 风险项 | 严重程度 | 说明 |
|--------|---------|------|
| **QPS/Burst 限制失效** | 🔴 高 | 每个新 clientset 有独立的 RateLimiter，无法限制总请求。例如配置 QPS=100，创建 10 个 clientset 后实际 QPS 可能达到 1000 |
| **数据一致性未解决** | 🔴 高 | 仍然使用 `NewFilteredPodInformer`，初始 List 仍使用 RV="0"，watch cache 延迟问题依然存在 |
| **API Server 连接数增加** | 🟡 中 | 多个 clientset 可能创建多个 HTTP 连接，增加 API Server 的连接维护开销 |
| **不同 API Server 副本视图不一致** | 🟡 中 | 请求可能打到不同 API Server 副本，每个副本的 watch cache 可能有短暂不一致 |
| **内存开销增加** | 🟡 中 | 每个 clientset 包含多个 typed client（admissionregistration、apps、batch 等），虽然共享 httpClient，但仍有额外的结构体开销 |

#### 3.1.4 改进方案：在单次操作内复用 client

如果只在单次 Create/Delete 操作内复用同一个 clientset（操作结束后销毁）：

**优势**:
- 该 clientset 的 RateLimiter 可以限制本次操作的请求
- 减少连接数

**劣势**:
- 多个并发操作仍然会有多个独立的 clientset，总 QPS 无法控制
- 数据一致性问题依然未解决
- clientset 创建/销毁开销

---

### 方案 2: 强制 ResourceVersion=""（已实施）

#### 3.2.1 实现思路

不使用 `NewFilteredPodInformer`，而是手动创建 `ListWatch`，在 ListFunc 中**强制覆盖** `ResourceVersion=""`。

**关键发现**: Reflector 会在调用 ListFunc 之前设置 `ResourceVersion`，因此必须在 ListFunc 内部覆盖。

**Reflector 的覆盖逻辑**（参考 `vendor/k8s.io/client-go/tools/cache/reflector.go:502-547`）：

```go
func (r *Reflector) list(stopCh <-chan struct{}) error {
    // Reflector 创建 options 并设置 ResourceVersion
    options := metav1.ListOptions{ResourceVersion: r.relistResourceVersion()}
    
    // relistResourceVersion() 对于新 informer 返回 "0"
    // 这意味着可以从 watch cache 读取
    
    // 调用 pager.ListWithAlloc(options)
    // pager 再调用我们的 ListFunc(opts)
    // 此时 opts.ResourceVersion 已经是 "0" 了
}
```

**因此我们必须在 ListFunc 内部覆盖**：

**已实施的代码**（[fuse_pod_manager.go:343-366](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/fuse_pod_manager.go#L343-L366)）：

```go
lw := &cache.ListWatch{
    ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
        // 覆盖 Reflector 设置的 ResourceVersion="0"
        // 强制从 etcd 读取，保证数据一致性
        options.ResourceVersion = ""
        options.FieldSelector = listOptions.FieldSelector
        options.LabelSelector = listOptions.LabelSelector
        return podClient.List(ctx, options)
    },
    WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
        options.FieldSelector = listOptions.FieldSelector
        options.LabelSelector = listOptions.LabelSelector
        return podClient.Watch(ctx, options)
    },
}
informer := cache.NewSharedIndexInformer(lw, &corev1.Pod{}, 0, nil)
```

#### 3.2.2 技术细节

**ResourceVersion 覆盖机制**:

1. **Reflector 初始化**（`reflector.go:504`）:
   ```go
   options := metav1.ListOptions{ResourceVersion: r.relistResourceVersion()}
   ```
   - 新 informer 的 `relistResourceVersion()` 返回 `"0"`（`reflector.go:912-915`）
   - 注释说明："For performance reasons, initial list performed by reflector uses '0' as resource version to allow it to be served from the watch cache"

2. **Pager 调用**（`reflector.go:547`）:
   ```go
   list, paginatedResult, err = pager.ListWithAlloc(context.Background(), options)
   ```
   - 将 Reflector 创建的 options 传给 pager
   - pager 再调用我们的 ListFunc(opts)

3. **ListFunc 覆盖**（我们的代码）:
   ```go
   ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
       options.ResourceVersion = ""  // 覆盖 "0" 为 ""
       return podClient.List(ctx, options)
   }
   ```

**ResourceVersion 语义**:
- `ResourceVersion="0"`: 从 watch cache 读取（**不保证最新**，watch cache 可能有延迟，可能读到过期数据）
- `ResourceVersion=""`: 一致性读（consistent read）
  - **K8s < 1.31**: 从 etcd 直接读取（quorum read，强一致）
  - **K8s >= 1.31**: 一致性读从 cache（保证 cache 同步到最新 revision，不会读到过期数据）
- `ResourceVersion="N"`: 从 watch cache 读取至少 RV=N 的数据

**K8s 1.31 一致性读解决的核心问题**：

根据 KEP-2340，K8s 1.31 的一致性读主要解决 "**Stale Read**" 问题（kubernetes/kubernetes#59848）：

**问题场景**（RV="0" 的风险）：
```
T1: 组件之前看到了 RV=1000 的数据
T2: 组件重启，创建新的 informer
T3: 新 informer 发起 List 请求，RV="0"
T4: 请求打到 watch cache，但 cache 只同步到 RV=800
T5: 返回 RV=800 的数据（比之前看到的 RV=1000 还旧！）
T6: Reflector "时光倒流"，丢失了 RV=801-1000 之间的变更
```

**一致性读的保证**（RV=""）：
1. ✅ 从 etcd 获取最新 revision（轻量级 quorum read，不返回数据）
2. ✅ 检查 watch cache 是否已同步到该 revision
3. ✅ 如果是，从 cache 读取（快速）
4. ✅ 如果否，等待 cache 同步（最多 100ms * N）
5. ✅ 如果超时，拒绝请求（返回错误）
6. ✅ **保证不会读到过期数据，不会"时光倒流"**

**重要区别**：
- K8s 1.31 的一致性读特性 **只影响 RV="" 的请求**
- **RV="0" 的行为没有改变**！仍然是从 watch cache 读取，不保证最新
- **我们使用 RV=""，因此在 K8s 1.31+ 上能获得一致性保证**

**分页行为影响**（`reflector.go:524-545`）:

Reflector 在调用 pager 之前，会根据 ResourceVersion 决定是否启用分页：

```go
switch {
case r.WatchListPageSize != 0:
    pager.PageSize = r.WatchListPageSize
case r.paginatedResult:
    // 之前已经获得过分页结果，保持默认分页
case options.ResourceVersion != "" && options.ResourceVersion != "0":
    // ResourceVersion 是具体值（不是 "" 也不是 "0"）
    // 禁用分页，强制从 watch cache 读取（如果启用了 watch cache）
    pager.PageSize = 0
default:
    // ResourceVersion 是 "" 或 "0"，使用默认分页（PageSize=500）
}
```

**关键点**：
1. Reflector 在 **调用 pager 之前** 根据 `options.ResourceVersion` 设置 `pager.PageSize`
2. 此时 `options.ResourceVersion` 是 `relistResourceVersion()` 返回的值（"0"）
3. 因此 pager 会启用分页（PageSize=500）
4. **但是**，pager 调用我们的 ListFunc 时传入的 options 仍然包含 `ResourceVersion="0"`
5. **我们在 ListFunc 内部覆盖为 `""`**，所以实际发送给 API Server 的是 `ResourceVersion=""`

**最终效果**：
- Reflector 认为会使用分页（因为它看到的是 "0"）
- 但实际 List 请求使用的是 `ResourceVersion=""`（一致性读）
- 由于 label selector 非常具体（volumeId + nodeName），匹配结果很少，分页影响不大

#### 3.2.3 K8s 1.31+ 一致性读从缓存的影响

**重要**: Kubernetes 1.31 引入了 [Consistent Reads from Cache](https://kubernetes.io/blog/2024/08/15/consistent-read-from-cache-beta/) 特性（Beta，默认启用）。

**行为变化**：

| K8s 版本 | ResourceVersion="" 的行为 | 性能特征 |
|---------|------------------------|----------|
| **< 1.31** | 从 etcd quorum read | 延迟较高，etcd 压力大 |
| **>= 1.31** | 优先从 watch cache 读取（如果 cache 足够新） | 延迟低，etcd 压力小 |

**K8s 1.31+ 的算法**：
1. API Server 先从 etcd 获取最新 revision（轻量级 quorum read，不返回数据）
2. 检查 watch cache 是否已经同步到这个 revision
3. 如果是，从 watch cache 读取（快速）
4. 如果否，等待 watch cache 同步（最多 100ms * N）
5. 如果超时，拒绝请求（返回错误）

**前提条件**：
- etcd 版本 >= 3.4.31 或 3.5.13+（修复了 progress notification bug）
- API Server 启动时检查 etcd 版本，如果版本太旧，自动回退到直接读 etcd

**性能提升**（5000 节点集群测试数据）：
- API Server CPU 使用量降低 30%
- etcd CPU 使用量降低 25%
- Pod LIST 请求 P99 延迟从 5s 降低到 1.5s（降低 3 倍）

**对我们的影响**：

| 方面 | K8s < 1.31 | K8s >= 1.31 |
|------|-----------|------------|
| **数据一致性** | ✅ 保证（etcd quorum read） | ✅ 保证（cache 足够新） |
| **etcd 压力** | 🔴 高（每次都是 quorum read） | 🟢 低（从 cache 读取） |
| **List 延迟** | 🟡 较高（etcd 读取） | 🟢 低（cache 读取） |
| **方案 2 的风险** | etcd 压力是主要风险 | ⚠️ 风险大幅降低 |

**关键发现**：
- 如果集群使用 K8s 1.31+ 且 etcd 版本足够新，**方案 2 的 etcd 压力风险几乎消除**
- 方案 2 在 K8s 1.31+ 上既能保证数据一致性，又不会显著增加 etcd 压力
- **这使得方案 2 成为更优选择**（前提是 K8s 版本支持）

**验证方法**：
```bash
# 检查 K8s 版本
kubectl version --short

# 检查 etcd 版本
kubectl -n kube-system exec etcd-<node-name> -- etcdctl version

# 检查特性是否启用（K8s 1.31+）
# 查看 API Server 日志，应该有 "ConsistentListFromCache" 相关日志
```

#### 3.2.4 风险评估

| 风险项 | 严重程度（K8s < 1.31） | 严重程度（K8s >= 1.31） | 说明 |
|--------|---------------------|---------------------|------|
| **etcd 压力增加** | 🔴 高 | 🟢 低 | K8s 1.31+ 从 cache 读取，etcd 压力大幅降低 |
| **List 延迟略增** | 🟡 中 | 🟢 低 | K8s 1.31+ 从 cache 读取，延迟更低 |
| **仍然每次创建 informer** | 🟡 中 | 🟡 中 | 没有解决"大量并发 Delete 创建独立 informer"的问题，每个 informer 仍然会触发 List + Watch |
| **内存开销** | 🟢 低 | 🟢 低 | 与原来相同，每个 informer 缓存的 pod 数量很少（label selector 很具体） |
| **Watch 连接数未减少** | 🟡 中 | 🟡 中 | 每个 informer 仍然会建立 Watch 连接，API Server 的 watch 连接数未减少 |
| **etcd 版本兼容性** | N/A | 🟡 中 | 如果 etcd 版本 < 3.4.31/3.5.13，会回退到直接读 etcd，失去性能优势 |

#### 3.2.5 适用场景

**K8s < 1.31**:
- 需要快速修复数据一致性问题
- etcd 集群性能充足，能够承受额外的 quorum read
- 并发 Delete 数量不极端（例如 < 100 并发）

**K8s >= 1.31**（推荐）:
- ✅ **强烈推荐**：既能保证数据一致性，又不会显著增加 etcd 压力
- etcd 版本 >= 3.4.31 或 3.5.13+
- 任何规模的集群都能受益

---

### 方案 3: 全局共享 Informer

#### 3.3.1 实现思路

在 `oss.go` 初始化时创建共享 informer，所有 Delete 操作复用同一个 informer 的缓存。

**初始化**（修改 [oss.go:77](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss/oss.go#L77)）：

```go
// 创建共享 informer factory
sharedInformerFactory := informers.NewSharedInformerFactoryWithOptions(
    clientset,
    0, // 不设置 resync
    informers.WithNamespace(namespace),
    informers.WithTweakListOptions(func(options *metav1.ListOptions) {
        // 只缓存 fuse pod，不限制 volumeId
        options.LabelSelector = "csi.alibabacloud.com/fuse-type"
        // 可选：按 nodeName 过滤（需要知道所有 node）
        // options.FieldSelector = fields.OneTermNotEqualSelector("spec.nodeName", "").String()
    }),
)
podInformer := sharedInformerFactory.Core().V1().Pods().Informer()

// 启动 informer
go sharedInformerFactory.Start(ctx.Done())

fusePodManagers := ossfpm.GetAllOSSFusePodManagers(csiCfg, m, clientset, podInformer)
```

**Delete 方法**（修改 [fuse_pod_manager.go:333](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/fuse_pod_manager.go#L333)）：

```go
func (fpm *FusePodManager) Delete(c *FusePodContext) error {
    // 直接从共享缓存查询，无需创建新的 informer
    allPods := fpm.sharedInformer.GetStore().List()
    
    var targetPods []*corev1.Pod
    for _, obj := range allPods {
        pod := obj.(*corev1.Pod)
        if matchesLabels(pod, c) {  // 根据 volumeId + nodeName 过滤
            targetPods = append(targetPods, pod)
        }
    }
    
    // 删除目标 pods
    for _, pod := range targetPods {
        podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
    }
    
    // 通过 watch 事件等待删除完成
    // ...
}
```

#### 3.3.2 技术细节

**Informer 缓存范围**:
- 缓存所有匹配 `csi.alibabacloud.com/fuse-type` label 的 pod
- 在本地内存中维护全量缓存
- Delete 操作通过本地过滤（根据 volumeId + nodeName）

**内存开销估算**:
- 假设集群有 N 个 OSS 挂载点（即 N 个 fuse pod）
- 每个 pod 对象在 informer 缓存中约占 10-20KB（包括 metadata、status 等）
- 总额外内存 ≈ N × 20KB
- 例如：10000 个 fuse pod ≈ 200MB

#### 3.3.3 风险评估

| 风险项 | 严重程度 | 说明 |
|--------|---------|------|
| **内存开销显著增加** | 🔴 高 | 需要缓存所有 fuse pod。大规模集群（如 10000+ 挂载点）可能额外占用 200MB+ 内存。CSI Controller 内存 limit 为 1024Mi，需评估是否会触发 OOM |
| **Informer 生命周期管理复杂** | 🟡 中 | 需要确保 informer 在组件启动时创建、停止时正确关闭。错误处理不当可能导致 goroutine 泄漏或资源泄漏 |
| **初始同步延迟** | 🟡 中 | informer 首次启动时需要 List 所有 fuse pod，如果数量很大（如 10000+），初始 List 可能耗时较长（秒级） |
| **缓存数据可能过期** | 🟢 低 | informer 通过 Watch 保持缓存更新，正常情况下延迟在毫秒级。但 Watch 断开重连期间可能有短暂不一致 |
| **过滤逻辑在应用层** | 🟢 低 | 需要在 Delete 中手动过滤 pod（根据 volumeId + nodeName），增加代码复杂度 |
| **调试复杂度增加** | 🟢 低 | 缓存问题排查比直接 List 更复杂，需要理解 informer 机制 |

#### 3.3.4 内存影响详细分析

**当前内存使用**（方案 2）:
- 每次 Delete 创建 informer，缓存少量 pod（label selector 很具体）
- informer 生命周期短（30s timeout），内存及时释放
- 峰值内存取决于并发 Delete 数量

**方案 3 内存使用**:
- 常驻内存：缓存所有 fuse pod
- 不随并发 Delete 数量变化
- 内存使用可预测，但基线更高

**决策因素**:
- 集群平均 fuse pod 数量
- CSI Controller 当前内存使用水位
- 是否有内存监控和告警
- 是否允许调整内存 limit

#### 3.3.5 适用场景

- 集群规模较大，并发 Delete 频繁（如 > 100 并发）
- API Server/etcd 压力已成为瓶颈
- 有充足的内存预算，或可以调整内存 limit
- 愿意承担 informer 生命周期管理的复杂度

---

## 4. 方案对比

### 4.1 核心指标对比

| 指标 | 方案 1: 新建 client | 方案 2: RV=""（已实施） | 方案 3: 共享 informer |
|------|-------------------|----------------------|-------------------|
| **数据一致性** | ❌ 未解决 | ✅ 解决 | ✅ 解决 |
| **QPS 控制** | ❌ 失效 | ✅ 保持 | ✅ 保持 |
| **API Server List 压力** | ❌ 增加 | 🟡 K8s < 1.31: 略增（etcd）<br>✅ K8s >= 1.31: 低（cache） | ✅ 大幅降低 |
| **API Server Watch 连接数** | ❌ 增加 | ❌ 未减少 | ✅ 大幅降低 |
| **etcd 压力** | 🟡 不变 | 🔴 K8s < 1.31: 增加（quorum read）<br>🟢 K8s >= 1.31: 低（cache） | ✅ 降低 |
| **内存开销** | 🟡 略增 | 🟢 低（与原方案相同） | 🔴 高（常驻缓存） |
| **实施复杂度** | 🟢 低 | 🟢 低（已完成） | 🟡 中 |
| **维护复杂度** | 🟢 低 | 🟢 低 | 🟡 中（生命周期管理） |
| **K8s 版本依赖** | 无 | ⚠️ K8s >= 1.31 性能最优 | 无 |

### 4.2 K8s 版本对方案选择的影响

**关键因素**：K8s 1.31 引入的 Consistent Reads from Cache 特性显著改变了方案 2 的风险评估。

| 场景 | 推荐方案 | 原因 |
|------|---------|------|
| **K8s >= 1.31 + etcd >= 3.5.13** | ✅ **方案 2** | 数据一致性 + 低 etcd 压力 + 低内存 + 实施简单 |
| **K8s < 1.31 + 小规模集群** | ✅ **方案 2** | etcd 压力可接受，实施简单 |
| **K8s < 1.31 + 大规模集群** | ⚠️ **方案 3** 或 **升级 K8s** | 方案 2 的 etcd 压力可能成为瓶颈 |
| **内存极度敏感** | ✅ **方案 2** | 方案 3 的内存开销不可接受 |
| **API Server Watch 连接数成为瓶颈** | ⚠️ **方案 3** | 方案 2 无法减少 Watch 连接数 |

### 4.3 风险评估总结

| 方案 | 主要风险 | 风险等级 | 是否可接受 |
|------|---------|---------|-----------|
| **方案 1** | QPS 限制失效、数据一致性未解决 | 🔴 高 | ❌ 不推荐 |
| **方案 2** | etcd 压力增加、Watch 连接数未减少 | 🟡 中 | ⚠️ 需监控 etcd 性能 |
| **方案 3** | 内存开销显著增加、生命周期管理复杂 | 🔴 高（内存） | ⚠️ 需评估内存预算 |

---

## 5. 决策建议

### 5.1 关键决策因素

1. **集群规模**:
   - 小规模（< 1000 挂载点）：方案 2 足够
   - 大规模（> 5000 挂载点）：需要评估方案 3

2. **内存预算**:
   - 当前 CSI Controller 内存使用水位？
   - 是否可以调整 memory limit（从 1024Mi 提升）？
   - 是否有 OOM 历史？

3. **API Server/etcd 性能**:
   - 当前 API Server QPS 水位？
   - etcd 延迟（P99）是否在可接受范围？
   - 是否有 etcd 性能告警？

4. **并发 Delete 频率**:
   - 峰值并发 Delete 数量？
   - 是否经常触发 API Server 限流？

### 4.4 推荐决策流程

```
检查 K8s 和 etcd 版本：
  ├─ K8s >= 1.31 + etcd >= 3.5.13？
  │   ├─ 是 → ✅ 使用方案 2（性能最优，风险最低）
  │   └─ 否 → 继续评估：
  │           ├─ 集群规模？
  │           │   ├─ 小规模（< 1000 挂载点）→ ✅ 方案 2（etcd 压力可接受）
  │           │   └─ 大规模（> 5000 挂载点）→ 继续评估：
  │           │           ├─ 可以升级 K8s 到 1.31+？
  │           │           │   ├─ 是 → ✅ 升级后使用方案 2
  │           │           │   └─ 否 → 继续评估：
  │           │           │           ├─ 内存预算充足？
  │           │           │           │   ├─ 是 → ⚠️ 方案 3（需评估内存风险）
  │           │           │           │   └─ 否 → ⚠️ 方案 2 + 限流/降级策略
  │           │           │           └─ API Server Watch 连接数是瓶颈？
  │           │           │               ├─ 是 → ⚠️ 方案 3
  │           │           │               └─ 否 → ✅ 方案 2
  │           └─ 中规模（1000-5000 挂载点）→ 监控 etcd 压力后决策
  └─ 是否已实施方案 2？
      ├─ 否 → 先实施方案 2（快速修复数据一致性）
      └─ 是 → 监控以下指标：
              ├─ etcd P99 延迟是否增加？
              ├─ API Server Watch 连接数是否过多？
              └─ 是否仍然出现 API Server 限流？
                    ├─ 否 → 保持方案 2
                    └─ 是 → 根据上面流程重新评估
```

### 4.5 监控指标建议

无论选择哪个方案，都建议监控以下指标：

1. **API Server 指标**:
   - `apiserver_request_total`（按 verb、resource 分组）
   - `apiserver_current_inflight_requests`
   - `apiserver_watch_events_total`

2. **etcd 指标**:
   - `etcd_server_quota_backend_bytes`（存储配额使用）
   - `etcd_disk_wal_fsync_duration_seconds`（P99）
   - `etcd_network_peer_round_trip_time_seconds`

3. **CSI Controller 指标**:
   - `container_memory_usage_bytes`（与 limit 对比）
   - `container_memory_working_set_bytes`
   - `kube_pod_container_status_last_terminated_reason`（是否 OOM）

4. **业务指标**:
   - Delete 操作 P99 延迟
   - Delete 操作失败率
   - 并发 Delete 数量

---

## 6. 实施状态

- ✅ **方案 2 已实施**: 修改 `fuse_pod_manager.go` 的 Delete 方法，强制 `ResourceVersion=""`
- ⏸️ **方案 3 待评估**: 需要收集集群规模、内存使用、API Server 压力等数据后决策

---

## 7. 参考资料

### 7.1 关键代码位置

- **Delete 方法**: [pkg/mounter/fuse_pod_manager/fuse_pod_manager.go:333-410](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/fuse_pod_manager.go#L333-L410)
- **fusePodManagers 初始化**: [pkg/oss/oss.go:77](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss/oss.go#L77)
- **QPS/Burst 配置**: [deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml:631-632](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml#L631-L632)
- **Controller 资源限制**: [deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml:684-690](file:///Users/suyashi/go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/deploy/charts/alibaba-cloud-csi-driver/templates/controller.yaml#L684-L690)

### 7.2 Kubernetes 源码参考

- **Reflector ListAndWatch**: `vendor/k8s.io/client-go/tools/cache/reflector.go:348-606`
- **relistResourceVersion 逻辑**: `vendor/k8s.io/client-go/tools/cache/reflector.go:902-918`
- **Clientset RateLimiter 创建**: `vendor/k8s.io/client-go/kubernetes/clientset.go:504-511`
- **HTTP Transport 创建**: `vendor/k8s.io/client-go/rest/transport.go:32`

### 7.3 Kubernetes 语义说明

- **ResourceVersion="0"**: 从 watch cache 读取，不保证最新数据
- **ResourceVersion=""**: 从 etcd 直接读取，强一致性（quorum read）
- **ResourceVersion="N"**: 从 watch cache 读取至少 RV=N 的数据
