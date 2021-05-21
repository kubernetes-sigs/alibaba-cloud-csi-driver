## 注意事项：

> 1. 升级脚本需要在Linux环境执行（Mac环境下由于shell语法不一致的问题，可能导致升级失败）；
>
> 2. 下载脚本的时候，可以在浏览器打开升级脚本，然后将脚本拷贝到升级节点；直接执行wget可能出现将html字符段一起下载；
>
> 3. 脚本执行命令为：bash **.sh {image-version}

## CSI-Plugin升级：

### 1. CSI-Plugin镜像版本信息：

版本信息详见：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

**说明：CSI 发布的版本为：v1.0.7, v1.1.1,,,,这样的版本号，所对应的主镜像为：acs/csi-plugin，升级时需要参考此镜像的版本号；**

### 2. 升级CSI Plugin版本到最新：

选择最新release版本号（例如: v1.1.3 对应acs/csi-plugin版本为 v1.18.8.47-906bd535-aliyun），下载并执行脚本（注意链接地址，请查询最新版本并链接）：

[https://raw.githubusercontent.com/kubernetes-sigs/alibaba-cloud-csi-driver/release-v1.1.3/deploy/ack/upgrade/upgrade_csi-plugin.sh](https://raw.githubusercontent.com/kubernetes-sigs/alibaba-cloud-csi-driver/release-v1.1.3/deploy/ack/upgrade/upgrade_csi-plugin.sh)

执行升级命令：
>$ bash upgrade_csi-plugin.sh v1.18.8.47-906bd535-aliyun

### 3. 检查升级情况：

查看升级后的csi-plugin 的镜像版本，是否镜像都已更新，并Running；

```
$ kubectl get pod -nkube-system | grep csi-plugin
csi-plugin-lbjpk                                     4/4     Running     0          7d23h
csi-plugin-pzcbb                                     4/4     Running     0          7d23h
csi-plugin-xq5ql                                     4/4     Running     0          7d23h

$ kubectl describe ds csi-plugin -nkube-system | grep Image
```

## CSI-Provisioner插件升级(先升级CSI-Plugin)：

### 1. CSI-Provisioner镜像版本信息(同CSI-Plugin插件)：
版本信息详见：[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

### 2. 升级CSI Provisioner版本到最新：
选择最新release版本（例如：v1.1.3 对应acs/csi-plugin的版本为 v1.18.8.47-906bd535-aliyun），下载脚本并执行：

> 注意：CSI Provisioner所指的镜像版本为：acs/csi-plugin的镜像版本，而acs/csi-provisioner这个镜像本质上是一个sidercar，不是升级需要的参数；

[https://raw.githubusercontent.com/kubernetes-sigs/alibaba-cloud-csi-driver/release-v1.1.3/deploy/ack/upgrade/upgrade_csi-provisioner.sh](https://raw.githubusercontent.com/kubernetes-sigs/alibaba-cloud-csi-driver/release-v1.1.3/deploy/ack/upgrade/upgrade_csi-provisioner.sh)

> $ bash upgrade_csi-provisioner.sh v1.18.8.47-906bd535-aliyun

**（这里填写acs/csi-plugin的镜像版本号）**

### 3. 检查升级情况：
查看升级后的csi-provisioner 镜像版本，是否都已更新，并Running；

```
$ kubectl get pod -nkube-system | grep csi-provisioner
csi-provisioner-5c85c7f5c9-4ndth                     7/7     Running     1          7d23h
csi-provisioner-5c85c7f5c9-c5bjj                     7/7     Running     0          7d23h

$ kubectl describe deploy csi-provisioner -nkube-system | grep Image
```

## 插件说明：

CSI-Plugin、CSI-Provisioner两个插件的阿里云发布镜像都是：acs/csi-plugin:{version}，部署模板中的其他镜像说明如下：

> acs/csi-node-driver-registrar：负责csi-plugin插件注册的sidecar容器；[Refer](https://kubernetes-csi.github.io/docs/node-driver-registrar.html)
>
> acs/csi-provisioner：数据卷动态创建的sidecar，负责调用CreateVolume等接口；[Refer](https://kubernetes-csi.github.io/docs/external-provisioner.html)
>
> acs/csi-attacher：数据卷挂载的sidecar，负责调用ControllerPublishVolume接口；[Refer](https://kubernetes-csi.github.io/docs/external-attacher.html)
>
> acs/csi-resizer：数据卷扩容的sidecar，负责调用ControllerExpandVolume接口；[Refer](https://kubernetes-csi.github.io/docs/external-resizer.html)
> 
> acs/csi-snapshotter：数据卷快照sidecar，负责创建volumeSnapshotContent对象；[Refer](https://kubernetes-csi.github.io/docs/external-snapshotter.html)
> 
> acs/snapshot-controller：数据卷快照sidecar，负责调用CreateSnapshot接口；[Refer](https://kubernetes-csi.github.io/docs/external-snapshotter.html)

