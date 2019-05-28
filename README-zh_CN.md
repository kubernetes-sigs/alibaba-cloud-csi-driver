# 阿里云Kubernetes CSI插件

[![Build Status](https://travis-ci.org/AliyunContainerService/csi-plugin.svg?branch=master)](https://travis-ci.org/AliyunContainerService/csi-plugin)
[![CircleCI](https://circleci.com/gh/AliyunContainerService/csi-plugin.svg?style=svg)](https://circleci.com/gh/AliyunContainerService/csi-plugin)
[![Go Report Card](https://goreportcard.com/badge/github.com/AliyunContainerService/csi-plugin)](https://goreportcard.com/report/github.com/AliyunContainerService/csi-plugin)

[English](./README.md) | 简体中文

## 插件介绍

阿里云CSI插件实现了Kubernetes平台使用阿里云云存储卷的生命周期管理，支持动态创建、挂载、使用云数据卷。 当前的CSI实现基于K8S 1.10以上的版本；

支持的阿里云存储：***云盘、NAS、OSS、LVM***


## 版本支持

当前Kubernetes与CSI版本兼容性如下:

| Kubernetes | CSI 版本 | CSI 状态 |
| ------ | ------ | ------ |
| v1.9  | v0.1   | Alpha |
| v1.10 | v0.2   | Beta |
| v1.11 | v0.3   | Beta |
| v1.12 | v0.3   | Beta |
| v1.13 | v1.0.0 | GA |


## 使用依赖

### 1. Kubernetes 依赖配置
需要开启以下Feature gates:

	--feature-gates=VolumeSnapshotDataSource=true,CSINodeInfo=true,CSIDriverRegistry=true

启动Privileged：

	配置 kube-apiserver：--allow-privileged=true ...
	配置 kubelet：--allow-privileged=true ...

### 2. 创建CRD(csidriver、csinodeinfo):

	# kubectl create -f https://raw.githubusercontent.com/kubernetes/csi-api/ab0df28581235f5350f27ce9c27485850a3b2802/pkg/crd/testdata/csidriver.yaml --validate=false 
	# kubectl create -f https://raw.githubusercontent.com/kubernetes/csi-api/ab0df28581235f5350f27ce9c27485850a3b2802/pkg/crd/testdata/csinodeinfo.yaml --validate=false 


### 3. 创建RBAC

	# kubectl create -f ./deploy/rbac.yaml


### 4. 依赖扩展插件

CSI v1.0.0需要以下扩展插件的支持：

| 官方扩展镜像  | 阿里云加速地址 |
|-------- |---------------------|
| quay.io/k8scsi/csi-attacher:v1.0.0 | registry.cn-hangzhou.aliyuncs.com/plugins/csi-attacher:v1.0.0 |
| quay.io/k8scsi/csi-snapshotter:v1.0.0 | registry.cn-hangzhou.aliyuncs.com/plugins/csi-snapshotter:v1.0.0 |
| quay.io/k8scsi/csi-provisioner:v1.0.0 | registry.cn-hangzhou.aliyuncs.com/plugins/csi-provisioner:v1.0.0 |
| quay.io/k8scsi/csi-node-driver-registrar:v1.0.1 | registry.cn-hangzhou.aliyuncs.com/plugins/csi-node-driver-registrar:v1.0.1 |

## 云盘CSI插件

云盘CSI插件支持动态创建云盘数据卷、挂载数据卷。云盘是一种块存储类型，只能同时被一个负载使用(ReadWriteOnce)。

云盘CSI插件更多详细说明请参考[云盘](./docs/disk.md)


## NAS CSI插件

NAS CSI插件支持为应用负载挂载阿里云NAS存储卷，目前不支持动态创建NAS卷。NAS存储是一种共享存储，可以同时被多个应用负载使用(ReadWriteMany)。

NAS CSI插件更多详细说明请参考[NAS](./docs/nas.md)


## OSS CSI插件

OSS CSI插件支持为应用负载挂载阿里云OSS Bucket，目前不支持动态创建OSS Bucket。OSS存储是一种共享存储，可以同时被多个应用负载使用(ReadWriteMany)。

OSS CSI插件更多详细说明请参考[OSS](./docs/oss.md)

## LVM CSI插件

LVM CSI插件支持创建LVM卷，并挂载。LVM不是高可用存储类型，使用时应用需要注意是否容忍。

LVM CSI插件更多详细说明请参考[LVM](./docs/lvm.md)

## 问题反馈

如您使用中遇到问题，可以提交到[Issues](https://github.com/AliyunContainerService/csi-plugin/issues)


## 社区交流
加入下面钉钉群，一起讨论您遇到的问题:
![](./docs/ding_group.jpg)
