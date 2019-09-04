# Alibaba Cloud Kubernetes CSI Plugin

English | [简体中文](./README-zh_CN.md)

## Introduction
Alibaba cloud CSI plugins implement an interface between CSI enabled Container
Orchestrator and Alibaba Cloud Storage. It allows dynamically provision Disk
volumes and attach it to workloads.
Current implementation of CSI plugins was tested in Kubernetes environment (requires Kubernetes 1.10+).

Current Support: ***Alibaba cloud Disk, OSS, NAS, LVM***;


### Disk CSI-Plugin

Disk csi-plugin support Alicloud disk provision and attachment. And alicloud disk is type of block storage, can only used as ReadWriteOnce mode. Only be attached to one node at the same time.

More detail information pls refer to [Disk](./docs/disk.md).

### NAS CSI-Plugin

Nas csi-plugin can support alicloud nas mount, and does not support provision nas volume. Nas storage is type of network storage and can be mount by multi nodes at the same time.

More detail information pls refer to [NAS](./docs/nas.md).


### OSS CSI-Plugin

OSS csi-plugin support Alicloud oss mount, and does not support provision volume.

More detail information pls refer to [OSS](./docs/oss.md).

### LVM CSI-Plugin

LVM csi-plugin support lvm create and mount. Lvm is not high availability.

More detail information pls refer to [LVM](./docs/lvm.md).


## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](https://kubernetes.io/community/).

You can reach the maintainers of this project at the [Cloud Provider SIG](https://github.com/kubernetes/community/tree/master/sig-cloud-provider).


### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

Please submit an issue at: [Issues](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/issues)
