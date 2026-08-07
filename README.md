# Alibaba Cloud Kubernetes CSI Driver

[![GoReportCard Widget]][GoReportCardResult]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)
[![Releases](https://img.shields.io/github/v/release/kubernetes-sigs/alibaba-cloud-csi-driver?include_prereleases)](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/kubernetes-sigs/alibaba-cloud-csi-driver)](./go.mod)

English | [简体中文](./README-zh_CN.md)

The Alibaba Cloud CSI Driver implements the [Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec)
specification, allowing Container Orchestrators (COs) such as Kubernetes to
manage the full lifecycle of Alibaba Cloud storage volumes — dynamic
provisioning, attaching, mounting, resizing and snapshotting.

> [!WARNING]
> Deploying this driver to your ACK cluster manually is **not recommended**.
> Instead, users should use ACK Add-ons to automatically deploy and manage the
> [Alibaba Cloud CSI Driver](https://www.alibabacloud.com/help/en/ack/product-overview/csi-plugin).
> Manual deployment of the driver in your ACK cluster is not officially supported by Alibaba Cloud.

---

## Table of Contents

- [Overview](#overview)
- [Supported Storage](#supported-storage)
- [Feature Matrix](#feature-matrix)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Documentation](#documentation)
- [Development](#development)
  - [Repository Layout](#repository-layout)
  - [Building](#building)
  - [Testing](#testing)
- [Community, Discussion, Contribution, and Support](#community-discussion-contribution-and-support)
- [Security](#security)
- [License](#license)

## Overview

Alibaba Cloud CSI plugins implement an interface between CSI-enabled Container
Orchestrators and Alibaba Cloud Storage. They allow you to dynamically provision
storage volumes and attach them to your workloads.

The current implementation has been tested against Kubernetes 1.14+; manual
deployment via Helm requires Kubernetes 1.26+.

## Supported Storage

| Driver | CSI Name                             | Type   | Access Modes | Dynamic Provisioning | Documentation           |
|--------|--------------------------------------|--------|--------------|----------------------|-------------------------|
| Disk   | `diskplugin.csi.alibabacloud.com`    | Block  | RWO          | Yes                  | [Cloud Disk](./docs/disk.md) |
| NAS    | `nasplugin.csi.alibabacloud.com`     | File   | RWX / RWO    | Yes                  | [NAS](./docs/nas.md)    |
| OSS    | `ossplugin.csi.alibabacloud.com`     | Object | RWX / RWO    | No (bucket mount)    | [OSS](./docs/oss.md)    |
| BMCPFS | `bmcpfsplugin.csi.alibabacloud.com`  | File   | RWX / RWO    | No (static only)     | [BMCPFS](./docs/bmcpfs-helm-readme.md) |

- **Cloud Disk** is block storage that can only be used by one workload at a
  time (`ReadWriteOnce`) and attached to a single node at a time. It supports
  snapshots, expansion, topology-aware scheduling, raw block volumes, and
  [local-disk data caching](./docs/disk-datacache.md) for acceleration.
- **NAS** is a shared network file system compatible with NFS/SMB, mountable
  from many nodes simultaneously (`ReadWriteMany`). CPFS 2.0 is now served
  through the NAS CSI plugin.
- **OSS** mounts object-storage buckets into pods. It does not provision buckets
  but can be mounted from multiple nodes (`ReadWriteMany`).
- **BMCPFS** mounts a Bare Metal Cloud Parallel File System — a high-performance
  parallel file system optimized for high-I/O workloads, mountable from many
  nodes simultaneously (`ReadWriteMany`). It supports **static provisioning
  only**: the file system must be pre-created in Alibaba Cloud. Both VSC (high
  performance, Lingjun) and VPC networking are supported.

> **CPFS CSI Plugin — Removed:** use the NAS CSI plugin for CPFS 2.0.

## Feature Matrix

| Feature         | Stage | Min Kubernetes Version | Min Driver Version |
|-----------------|-------|------------------------|--------------------|
| Topology        | GA    | 1.17                   | v1.0.2             |
| Resize (Expand) | GA    | 1.16                   | v1.0.5             |
| Snapshots       | GA    | 1.20                   | v1.1.2             |

## Getting Started

### Prerequisites

- Kubernetes version >= 1.26
- `kubectl` configured to communicate with the cluster
- Helm 3
- RAM permissions for the driver to invoke Alibaba Cloud OpenAPIs on your behalf
  (see the [example policies](./docs/ram-policies))
- Individual drivers may have additional requirements; consult their
  [documentation](#documentation)

### Installation

The recommended way to run the driver is through
[Alibaba Cloud Container Service for Kubernetes (ACK)](https://www.alibabacloud.com/product/kubernetes),
where the CSI drivers are deployed and managed automatically as add-ons.

For manual deployment on self-managed clusters, install the drivers with Helm:

```shell
helm repo add alibaba-cloud-csi-driver https://kubernetes-sigs.github.io/alibaba-cloud-csi-driver
helm repo update

# Self-built cluster on ECS
helm upgrade --install alibaba-cloud-csi-driver alibaba-cloud-csi-driver/alibaba-cloud-csi-driver \
  --values values-ecs.yaml --namespace kube-system
```

See the full [Installation Guide](./docs/install.md) for RAM setup,
configuration presets, and verification steps.

## Documentation

| Topic | Description |
|-------|-------------|
| [Installation](./docs/install.md)                       | Deploy and configure the drivers |
| [Cloud Disk](./docs/disk.md)                            | Cloud disk provisioning and mounting |
| [Disk — Raw Block Volume](./docs/disk-block.md)         | Use a disk as a raw block device |
| [Disk — Volume Expansion](./docs/disk-resizer.md)       | Online disk resizing |
| [Disk — Snapshot & Restore](./docs/disk-snapshot-restore.md) | Create and restore disk snapshots |
| [Disk — Local Data Cache](./docs/disk-datacache.md)     | Accelerate a cloud disk with a local disk (dm-cache) |
| [NAS](./docs/nas.md)                                    | NAS volume provisioning and mounting |
| [NAS — Dynamic Provisioning](./docs/nas-dynamic.md)     | Subpath/filesystem provisioning |
| [NAS — Volume Expansion](./docs/nas-expansion.md)       | NAS quota expansion |
| [OSS](./docs/oss.md)                                    | Mount OSS buckets |
| [BMCPFS](./docs/bmcpfs-helm-readme.md)                  | Mount a Bare Metal Cloud Parallel File System |
| [Metrics](./docs/csi-metric.md)                         | CSI metrics |
| [FlexVolume → CSI Migration](./docs/migrate)            | Migrate from FlexVolume |

Runnable manifests for each feature live under [`examples/`](./examples).

## Development

### Repository Layout

```
.
├── cmd/          # Entry points (e.g. csi-agent)
├── pkg/          # Driver implementations
│   ├── disk/     # Cloud Disk CSI driver (incl. datacache/)
│   ├── nas/      # NAS CSI driver
│   ├── oss/      # OSS CSI driver
│   ├── bmcpfs/   # BMCPFS (Bare Metal Cloud Parallel File System) CSI driver
│   ├── cloud/    # Alibaba Cloud OpenAPI clients
│   ├── mounter/  # Mount helpers and fuse pod managers
│   └── ...       # Shared utilities, metrics, options, etc.
├── deploy/       # Helm chart and deployment manifests
├── examples/     # Example StorageClasses, PVCs and workloads
├── docs/         # User and developer documentation
├── build/        # Container image build files
└── hack/         # Development and CI scripts
```

### Building

Container images are built via:

```shell
make build
```

### Testing

```shell
make fmt      # check formatting
make vet      # go vet
make lint     # golangci-lint
make test     # unit tests

# End-to-end Helm install against a kind cluster
make check-helm-kind
```

## Community, Discussion, Contribution, and Support

Learn how to engage with the Kubernetes community on the
[community page](https://kubernetes.io/community/).

You can reach the maintainers of this project at the
[Cloud Provider SIG](https://github.com/kubernetes/community/tree/master/sig-cloud-provider).

You can join the DingTalk group (Group ID: **33936810**) to talk with us.

Please submit issues at [GitHub Issues](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/issues),
and read [CONTRIBUTING.md](./CONTRIBUTING.md) before opening a pull request.

### Code of Conduct

Participation in the Kubernetes community is governed by the
[Kubernetes Code of Conduct](code-of-conduct.md).

## Security

Please report vulnerabilities by email to kubernetes-security@service.aliyun.com.
See our [SECURITY.md](./SECURITY.md) for details.

## License

This project is licensed under the [Apache License 2.0](./LICENSE).

[GoReportCard Widget]: https://goreportcard.com/badge/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
[GoReportCardResult]: https://goreportcard.com/report/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
