# Alibaba Cloud Kubernetes CSI Plugin

[![Build Status](https://travis-ci.org/AliyunContainerService/csi-plugin.svg?branch=master)](https://travis-ci.org/AliyunContainerService/csi-plugin)
[![CircleCI](https://circleci.com/gh/AliyunContainerService/csi-plugin.svg?style=svg)](https://circleci.com/gh/AliyunContainerService/csi-plugin)
[![Go Report Card](https://goreportcard.com/badge/github.com/AliyunContainerService/csi-plugin)](https://goreportcard.com/report/github.com/AliyunContainerService/csi-plugin)

English | [简体中文](./README-zh_CN.md)

## Overview

Alibaba cloud CSI plugins implement an interface between CSI enabled Container
Orchestrator and Alibaba Cloud Storage. It allows dynamically provision Disk
volumes and attach it to workloads.
Current implementation of CSI plugins was tested in Kubernetes environment (requires Kubernetes 1.10+).

Current Support: ***Alibaba cloud Disk, OSS, NAS, LVM***;

## CSI Version Support

CSI using with Kubrentes:

| Kubernetes | CSI Version | CSI Status |
| ------ | ------ | ------ |
| v1.9  | v0.1   | Alpha |
| v1.10 | v0.2   | Beta |
| v1.11 | v0.3   | Beta |
| v1.12 | v0.3   | Beta |
| v1.13 | v1.0.0 | GA |
| v1.14 | v1.0.0 | GA |

## Requirements

### 1. Kubernetes Configurations
Enable Feature gates:

	--feature-gates=VolumeSnapshotDataSource=true,CSINodeInfo=true,CSIDriverRegistry=true

In Kubernetes 1.14, CSINodeInfo, CSIDriverRegistry are in beta, and no need to enable;

Enable Privileged：

	enable kube-apiserver with --allow-privileged=true ...
	enable kubelet with --allow-privileged=true ...

Create CRDs for csidriver、csinodeinfo(Below 1.14 version):

	# kubectl create -f https://raw.githubusercontent.com/kubernetes/csi-api/ab0df28581235f5350f27ce9c27485850a3b2802/pkg/crd/testdata/csidriver.yaml --validate=false 
	# kubectl create -f https://raw.githubusercontent.com/kubernetes/csi-api/ab0df28581235f5350f27ce9c27485850a3b2802/pkg/crd/testdata/csinodeinfo.yaml --validate=false 


### 2. RBAC

	# kubectl create -f ./deploy/rbac.yaml

### 3. Plugin Dependencies

CSI v1.0.0 should be supported by below externals:

| External Plugin Oversea Repo  | External Plugin Aliyun Repo |
|-------- |---------------------|
| quay.io/k8scsi/csi-attacher:v1.0.0 | registry.cn-hangzhou.aliyuncs.com/plugins/csi-attacher:v1.0.0 |
| quay.io/k8scsi/csi-snapshotter:v1.2.0 | registry.cn-hangzhou.aliyuncs.com/acs/csi-snapshotter:v1.2.0 |
| quay.io/k8scsi/csi-provisioner:v1.2.0 | registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v1.2.0 |
| quay.io/k8scsi/csi-node-driver-registrar:v1.1.0 | registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.1.0 |

## Disk CSI-Plugin

Disk csi-plugin support Alicloud disk provision and attachment. And alicloud disk is type of block storage, can only used as ReadWriteOnce mode. Only be attached to one node at the same time.

More detail information pls refer to [Disk](./docs/disk.md).

## NAS CSI-Plugin

Nas csi-plugin can support alicloud nas mount, and does not support provision nas volume. Nas storage is type of network storage and can be mount by multi nodes at the same time.

More detail information pls refer to [NAS](./docs/nas.md).


## OSS CSI-Plugin

OSS csi-plugin support Alicloud oss mount, and does not support provision volume.

More detail information pls refer to [OSS](./docs/oss.md).

## LVM CSI-Plugin

LVM csi-plugin support lvm create and mount. Lvm is not high availability.

More detail information pls refer to [LVM](./docs/lvm.md).

## Troubleshooting

Please submit an issue at: [Issues](https://github.com/AliyunContainerService/csi-plugin/issues)
