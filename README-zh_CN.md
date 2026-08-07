# 阿里云 Kubernetes CSI 插件

[![GoReportCard Widget]][GoReportCardResult]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)
[![Releases](https://img.shields.io/github/v/release/kubernetes-sigs/alibaba-cloud-csi-driver?include_prereleases)](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/kubernetes-sigs/alibaba-cloud-csi-driver)](./go.mod)

[English](./README.md) | 简体中文

阿里云 CSI 插件实现了 [容器存储接口 (CSI)](https://github.com/container-storage-interface/spec)
规范，使 Kubernetes 等容器编排系统能够管理阿里云存储卷的完整生命周期——动态创建、
挂载、扩容与快照。

> [!WARNING]
> 不建议在 ACK 集群中手动部署该驱动程序。用户应使用 ACK 组件中心自动部署和管理
> [Alibaba Cloud CSI Driver](https://help.aliyun.com/zh/ack/product-overview/csi-plugin)。
> flexvolume 迁移场景除外，请按照 [迁移文档](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/use-csi-compatible-controller-to-migrate-from-flexvolume-to-csi) 进行迁移。
> ACK 官方不支持在集群中手动部署该驱动程序。

---

## 目录

- [插件介绍](#插件介绍)
- [支持的存储](#支持的存储)
- [版本说明](#版本说明)
- [快速开始](#快速开始)
  - [前置条件](#前置条件)
  - [安装](#安装)
- [文档](#文档)
- [开发](#开发)
  - [代码结构](#代码结构)
  - [构建](#构建)
  - [测试](#测试)
- [社区、贡献、讨论、支持](#社区贡献讨论支持)
- [安全](#安全)
- [许可证](#许可证)

## 插件介绍

阿里云 CSI 插件实现了在 Kubernetes 中对阿里云云存储卷的生命周期管理，支持动态创建、
挂载、使用云数据卷。当前的 CSI 实现基于 K8S 1.14 以上的版本；通过 Helm 手动部署要求
Kubernetes 1.26 以上。

## 支持的存储

| 驱动 | CSI 名称                             | 类型   | 访问模式    | 动态创建           | 文档                   |
|------|--------------------------------------|--------|-------------|--------------------|------------------------|
| 云盘 | `diskplugin.csi.alibabacloud.com`    | 块存储 | RWO         | 支持               | [云盘](./docs/disk.md) |
| NAS  | `nasplugin.csi.alibabacloud.com`     | 文件   | RWX / RWO   | 支持               | [NAS](./docs/nas.md)   |
| OSS  | `ossplugin.csi.alibabacloud.com`     | 对象   | RWX / RWO   | 不支持（Bucket 挂载） | [OSS](./docs/oss.md)  |
| BMCPFS | `bmcpfsplugin.csi.alibabacloud.com` | 文件  | RWX / RWO   | 不支持（仅静态）    | [BMCPFS](./docs/bmcpfs-helm-readme-zh.md) |

- **云盘** 是块存储类型，只能同时被一个负载使用（`ReadWriteOnce`），且同一时间只能
  挂载到一个节点。支持快照、扩容、拓扑感知调度、裸块设备，以及
  [本地盘为云盘加速](./docs/disk-datacache-zh.md)。
- **NAS** 是一种兼容 NFS/SMB 的共享网络文件系统，可同时被多个节点挂载
  （`ReadWriteMany`）。CPFS 2.0 现已由 NAS CSI 插件提供支持。
- **OSS** 将对象存储 Bucket 挂载到 Pod 中。不支持动态创建 Bucket，但可被多个节点
  同时挂载（`ReadWriteMany`）。
- **BMCPFS** 用于挂载并行文件系统（Bare Metal Cloud Parallel File System）——一种
  面向高 I/O 负载优化的高性能并行文件系统，可同时被多个节点挂载（`ReadWriteMany`）。
  仅支持**静态挂载**：需预先在阿里云创建好文件系统。支持 VSC（高性能，灵骏）与 VPC
  两种网络。

> **CPFS CSI 插件——已删除：** 挂载 CPFS 2.0 请使用 NAS CSI 插件。

## 版本说明

| Feature         | Stage | Min Kubernetes Version | Min Driver Version |
|-----------------|-------|------------------------|--------------------|
| Topology        | GA    | 1.17                   | v1.0.2             |
| Resize (Expand) | GA    | 1.16                   | v1.0.5             |
| Snapshots       | GA    | 1.20                   | v1.1.2             |

## 快速开始

### 前置条件

- Kubernetes 版本 >= 1.26
- 已配置 `kubectl` 可访问集群
- Helm 3
- 授予驱动调用阿里云 OpenAPI 的 RAM 权限（参考 [示例策略](./docs/ram-policies)）
- 各驱动可能有额外要求，请参考对应的 [文档](#文档)

### 安装

推荐通过 [阿里云容器服务 Kubernetes 版 (ACK)](https://www.alibabacloud.com/product/kubernetes)
运行本驱动，此时 CSI 驱动会作为组件自动部署和管理。

如需在自建集群手动部署，可使用 Helm：

```shell
helm repo add alibaba-cloud-csi-driver https://kubernetes-sigs.github.io/alibaba-cloud-csi-driver
helm repo update

# ECS 上的自建集群
helm upgrade --install alibaba-cloud-csi-driver alibaba-cloud-csi-driver/alibaba-cloud-csi-driver \
  --values values-ecs.yaml --namespace kube-system
```

完整的 RAM 配置、配置预设和验证步骤请参考 [安装文档](./docs/install.md)。

## 文档

| 主题 | 说明 |
|------|------|
| [安装](./docs/install.md)                               | 部署与配置驱动 |
| [云盘](./docs/disk.md)                                  | 云盘的创建与挂载 |
| [云盘——裸块设备](./docs/disk-block.md)                  | 将云盘作为裸块设备使用 |
| [云盘——扩容](./docs/disk-resizer.md)                    | 在线扩容云盘 |
| [云盘——快照与恢复](./docs/disk-snapshot-restore.md)     | 创建与恢复云盘快照 |
| [云盘——本地盘加速](./docs/disk-datacache-zh.md)         | 用本地盘为云盘加速（dm-cache） |
| [NAS](./docs/nas.md)                                    | NAS 卷的创建与挂载 |
| [NAS——动态创建](./docs/nas-dynamic.md)                  | 子目录/文件系统动态创建 |
| [NAS——扩容](./docs/nas-expansion.md)                    | NAS 配额扩容 |
| [OSS](./docs/oss.md)                                    | 挂载 OSS Bucket |
| [BMCPFS](./docs/bmcpfs-helm-readme-zh.md)               | 挂载并行文件系统（Bare Metal CPFS） |
| [监控指标](./docs/csi-metric.md)                        | CSI 监控指标 |
| [FlexVolume → CSI 迁移](./docs/migrate)                 | 从 FlexVolume 迁移 |

各功能可直接运行的示例清单位于 [`examples/`](./examples) 目录。

## 开发

### 代码结构

```
.
├── cmd/          # 入口程序（如 csi-agent）
├── pkg/          # 驱动实现
│   ├── disk/     # 云盘 CSI 驱动（含 datacache/）
│   ├── nas/      # NAS CSI 驱动
│   ├── oss/      # OSS CSI 驱动
│   ├── bmcpfs/   # BMCPFS（并行文件系统）CSI 驱动
│   ├── cloud/    # 阿里云 OpenAPI 客户端
│   ├── mounter/  # 挂载辅助与 fuse pod 管理
│   └── ...       # 共享工具、监控、配置项等
├── deploy/       # Helm chart 与部署清单
├── examples/     # 示例 StorageClass、PVC 与工作负载
├── docs/         # 用户与开发文档
├── build/        # 容器镜像构建文件
└── hack/         # 开发与 CI 脚本
```

### 构建

构建容器镜像：

```shell
make build
```

### 测试

```shell
make fmt      # 检查代码格式
make vet      # go vet
make lint     # golangci-lint
make test     # 单元测试

# 在 kind 集群上进行端到端 Helm 安装
make check-helm-kind
```

## 社区、贡献、讨论、支持

可以到 [Kubernetes](https://kubernetes.io/community/) 社区学习如何获取支持；

可以到 [Cloud Provider SIG](https://github.com/kubernetes/community/tree/master/sig-cloud-provider)
联系到项目管理者；

可以加入钉钉群（群 ID：**33936810**）和我们一起讨论遇到的问题；

欢迎向社区提交 [Issue](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/issues)，
提交 PR 前请阅读 [CONTRIBUTING.md](./CONTRIBUTING.md)。

### 行为准则

参与 Kubernetes 社区请参考 [Kubernetes 行为准则](code-of-conduct.md)。

## 安全

对于发现的安全漏洞，请邮件发送至 kubernetes-security@service.aliyun.com，
您可在 [SECURITY.md](./SECURITY.md) 文件中找到更多信息。

## 许可证

本项目基于 [Apache License 2.0](./LICENSE) 开源。

[GoReportCard Widget]: https://goreportcard.com/badge/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
[GoReportCardResult]: https://goreportcard.com/report/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
