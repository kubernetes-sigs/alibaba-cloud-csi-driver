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
  [本地盘为云盘加速](./docs/disk-datacache.md)。
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

- 已配置 `kubectl` 可访问集群
- Helm 3
- 授予驱动调用阿里云 OpenAPI 的 RAM 权限（参考 [示例策略](./docs/ram-policies)）
- 各驱动可能有额外要求，请参考对应的 [文档](./docs/)

### 安装

部署方式、RAM 配置、配置预设和验证步骤请参考 [安装文档](./docs/install.md)。

## 开发

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
