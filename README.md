
# Alibaba Cloud Kubernetes CSI Plugin

English | [简体中文](./README-zh_CN.md)

## Introduction
Alibaba cloud CSI plugins implement an interface between CSI enabled Container
Orchestrator and Alibaba Cloud Storage. It allows dynamically provision Disk
volumes and attach it to workloads.
Current implementation of CSI plugins was tested in Kubernetes environment (requires Kubernetes 1.14+).

Current Support: ***Alibaba cloud Disk, OSS, NAS, CPFS, LVM***;


### Disk CSI Plugin

Disk CSI Plugin support Alicloud disk provision and attachment. And alicloud disk is type of block storage, can only used as ReadWriteOnce mode. Only be attached to one node at the same time.

More detail information pls refer to [Disk](./docs/disk.md).


### NAS CSI Plugin

Nas CSI Plugin can support alicloud nas volume provision and mount. Nas storage is type of network storage and can be mount by multi nodes at the same time.

More detail information pls refer to [NAS](./docs/nas.md).


### CPFS CSI Plugin

CPFS CSI Plugin can support alicloud cpfs volume provision and mount. CPFS storage is type of network storage and can be mount by multi nodes at the same time.

More detail information pls refer to [CPFS](./docs/cpfs.md).


### OSS CSI Plugin

OSS CSI Plugin support Alicloud oss mount, and does not support provision volume. OSS storage is type of object storage and can be mount by multi nodes at the same time.

More detail information pls refer to [OSS](./docs/oss.md).


### LVM CSI Plugin

LVM CSI Plugin support lvm volume provision and mount. Lvm volume is a type of local volume and not used as high availability scenario.

More detail information pls refer to [LVM](./docs/lvm.md).


## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](https://kubernetes.io/community/).

You can reach the maintainers of this project at the [Cloud Provider SIG](https://github.com/kubernetes/community/tree/master/sig-cloud-provider).

You can join the DingDing Talking [Group](https://qr.dingtalk.com/action/joingroup?code=v1,k1,xxf5eqc7eMgILnXxj9Chab8KNZFoPtD00kaOtTKg/Rk=&_dt_no_comment=1&origin=11) to talk with us.

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

Please submit an issue at: [Issues](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/issues)
