
# v1.0.0
[Documentation](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/README.md)

## Action Required
* Need Kubernetes Cluster to 1.14+ before deploying csi plugin.
* RABC requirement before deploying csi plugin. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/deploy/rbac.yaml)
* CSIDriver requirement before deploying csi plugin. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/v1.0.0/deploy/ack/csi-plugin.yaml)

## Deploy & Upgrade
* Deploy csi with the templates. [Reference](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/v1.0.0/deploy)

## Released Image:
* registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.14.5.47-5577bcb-aliyun

## Notable changes

*This is the first release of alibabacloud csi plugin, below features are supported.*

* Support AlibabaCloud EBS plugin, support static volume；
* Support AlibabaCloud EBS Dynamic Provisioner, support dynamic volume；
* Support AlibabaCloud AK/STS authorization;
* Support AlibabaCloud NAS plugin, support static volume；
* Support AlibabaCloud OSS plugin, support static volume；
* Support LVM dynamic provisioner, support dnamic volume；
* Support Block Volume, Shared Ebs, Snapshot in EBS Plugin;