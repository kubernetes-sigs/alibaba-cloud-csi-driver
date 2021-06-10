
## RBAC 安装：

下载RBAC配置文件到操作机，并部署：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/rbac.yaml](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/rbac.yaml)

执行：
>$ kubectl apply -f rbac.yaml


## CSI-Plugin 安装：

### 1. 下载模板：
下载最新版本的CSI Plugin部署模板：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/ack/csi-plugin.yaml](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/ack/csi-plugin.yaml)

将部署模板下载到您的操作机，并保存（csi-plugin.yaml）。


### 2. 适配模板并部署：
根据集群所在的Region修改模板中的镜像地址。例如：如果是cn-beijing的集群

则将 registry.cn-hangzhou.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0 中的：
> registry 改为 registry-vpc
>
> cn-hangzhou 改为 cn-beijing

即：registry-vpc.cn-beijing.aliyuncs.com/acs/csi-node-driver-registrar:v1.2.0

模板中的其他镜像也是如此更新；

执行部署：
>$ kubectl apply -f csi-plugin.yaml

### 3. 检查安装情况：

> $ kubectl get pod -nkube-system | grep csi-plugin
>
> $ kubectl describe ds csi-plugin -nkube-system | grep image


## CSI-Provisioner 安装：

### 1. 下载模板：
下载最新版本的CSI Provisioner部署模板：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/ack/csi-provisioner.yaml](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/ack/csi-provisioner.yaml)

将部署模板下载到您的操作机，并保存（csi-provisioner.yaml）。


### 2. 适配模板并部署：
根据集群所在的Region修改模板中的镜像地址。例如：如果是cn-beijing的集群

则将 registry.cn-hangzhou.aliyuncs.com/acs/csi-provisioner:v1.6.0-e360c7e43-aliyun 中的：
> registry 改为 registry-vpc
>
> cn-hangzhou 改为 cn-beijing

即：registry-vpc.cn-beijing.aliyuncs.com/acs/csi-provisioner:v1.6.0-e360c7e43-aliyun

**模板中的其他镜像也是如此更新；**

执行部署：
>$ kubectl apply -f csi-provisioner.yaml

### 3. 检查安装情况：

> $ kubectl get pod -nkube-system | grep csi-provisioner
>
> $ kubectl describe deploy csi-provisioner -nkube-system | grep image

