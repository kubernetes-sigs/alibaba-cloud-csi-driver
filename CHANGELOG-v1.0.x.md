
# v1.0.0
[Documentation](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/README.md)

## Action Required
* Upgrade Kubernetes Cluster to 1.14+ before deploying CSI Plugin.
* Configure RABC before deploying CSI Plugin. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/deploy/rbac.yaml)
* Deploy CSIDriver Object before deploying CSI Plugin. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/deploy/ack/csi-plugin.yaml)

## Deploy & Upgrade
* Deploy CSI Plugin with the templates. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/v1.0.0/deploy)

## Released Image:
* registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.14.5.47-5577bcb-aliyun

## Notable changes

*This is the first release of Alibaba Cloud CSI Plugin, below features are supported.*

* Support Alibaba Cloud EBS plugin for static volume；
* Support Alibaba Cloud EBS plugin for block volume;
* Support Alibaba Cloud EBS plugin for shared EBS;
* Support Alibaba Cloud EBS plugin for snapshot;
* Support Alibaba Cloud EBS provisioner for dynamic volume；
* Support Alibaba Cloud NAS plugin for static volume；
* Support Alibaba Cloud OSS plugin for static volume；
* Support LVM provisioner for dynamic volume；
* Support Alibaba Cloud AK/STS authorization;