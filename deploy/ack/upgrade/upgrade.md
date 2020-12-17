
## CSI-Plugin升级：

### 1. CSI-Plugin镜像版本信息：

版本信息详见：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

**说明：CSI 发布的版本为：v1.0.1, v1.0.2,,,,这样的版本号，所对应的主镜像为：acs/csi-plugin，升级时需要参考此镜像的版本号；**

### 2. 升级CSI Plugin版本到最新：

选择最新release版本号（例如v：1.1.0 对应acs/csi-plugin版本为 v1.16.9.43-f36bb540-aliyun），下载并执行脚本（注意链接地址）：

[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/release-v1.1.0/deploy/ack/upgrade/upgrade_csi-plugin.sh](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/release-v1.1.0/deploy/ack/upgrade/upgrade_csi-plugin.sh)

执行：
>$ sh upgrade_csi-plugin.sh v1.16.9.43-f36bb540-aliyun

### 3. 检查升级情况：

查看升级后的csi-plugin 的镜像版本，是否镜像都已更新，并Running；
> $ kubectl get pod -nkube-system | grep csi-plugin
>
> $ kubectl describe ds csi-plugin -nkube-system | grep Image


## CSI-Provisioner升级：

### 1. CSI-Provisioner镜像版本信息：
版本信息详见：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

### 2. 升级CSI Provisioner版本到最新：
选择最新release版本（例如：v1.1.0 对应acs/csi-plugin的版本为 v1.16.9.43-f36bb540-aliyun），下载脚本并执行：

> 注意：CSI Provisioner所指的镜像版本为：acs/csi-plugin的镜像版本，例如：v1.16.9.43-f36bb540-aliyun，而acs/csi-provisioner这个镜像本质上是一个sidercar，不是升级需要的参数；

[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/release-v1.1.0/deploy/ack/upgrade/upgrade_csi-plugin.sh](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/release-v1.1.0/deploy/ack/upgrade/upgrade_csi-plugin.sh)

> $ sh upgrade_csi-provisioner.sh v1.16.9.43-f36bb540-aliyun

（这里填写acs/csi-plugin的镜像版本号）

### 3. 检查升级情况：
查看升级后的csi-provisioner 镜像版本，是否都已更新，并Running；

> $ kubectl get pod -nkube-system | grep csi-provisioner
>
> $ kubectl describe deploy csi-provisioner -nkube-system | grep Image

