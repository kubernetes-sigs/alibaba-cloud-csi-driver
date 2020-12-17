# flexvolume 迁移至 csi -- 集群中还没有使用ACK数据卷服务

说明：

> 本文档为您提供一个方案，实现将已有集群从Flexvolume插件转换到CSI插件；
> 
> 但如果您集群没有关键业务，并可以删除重建，建议您重新申请阿里云ACK最新集群，默认部署CSI最新版本；

## 背景介绍
CSI/Flexvolume 简介：[https://help.aliyun.com/document_detail/157025.html](https://help.aliyun.com/document_detail/157025.html)

> csi插件体系包含：csi-provisioner(deployment) + csi-plugin(daemonset)；
>
> flexvolume插件体系包含： disk-controller(deployment) + flexvolume(daemonset)；

两者区别： 
> 目前CSI和Flexvolume 插件运行依赖的kubelet参数不同；
>
> Flexvolume 插件需要配置kubelet 的 enable-controller-attach-detach 为 false；
>
> CSI 插件需要配置kubelet 的 enable-controller-attach-detach 为 true；

CSI 插件体系较 flexvolume 更加稳定高效，本文是将集群从flexvolume 插件体系迁移到csi 插件体系的方案，需要您的集群还没有开始使用容器存储服务；即没有使用Flexvolume进行挂载数据卷；

如果您的集群已经使用了Flexvolume挂载相关存储，请参考本系列其他文章；

## 删除Flexvolume插件：

```
kubectl delete deploy alicloud-disk-controller -nkube-system
kubectl delete sc alicloud-disk-available alicloud-disk-efficiency alicloud-disk-essd alicloud-disk-ssd
kubectl delete ds flexvolume -nkube-system
```

删除成功的话，下列命令都应该返回为空：

```
kubectl get sc 
kubectl get pod -nkube-system | grep flexvolume
kubectl get pod -nkube-system | grep alicloud-disk-controller
```

## 修改 kubelet 参数：

部署下面模板，将集群中所有kubelet的参数改成CSI匹配的参数：

执行下面脚本会重启kubelet，请评估对应用影响；

```yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: kubelet-set
spec:
  selector:
    matchLabels:
      app: kubelet-set
  template:
    metadata:
      labels:
        app: kubelet-set
    spec:
      tolerations:
        - operator: "Exists"
      hostNetwork: true
      hostPID: true
      containers:
        - name: kubelet-set
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: registry.cn-hangzhou.aliyuncs.com/acs/csi-plugin:v1.16.9.43-f36bb540-aliyun
          imagePullPolicy: "Always"
          env:
          - name: enableADController
            value: "true"
          command: ["sh", "-c"]
          args:
          - echo "Starting kubelet flag set to $enableADController";
            ifFlagTrueNum=`cat /host/etc/systemd/system/kubelet.service.d/10-kubeadm.conf | grep enable-controller-attach-detach=$enableADController | grep -v grep | wc -l`;
            echo "ifFlagTrueNum is $ifFlagTrueNum";
            if [ "$ifFlagTrueNum" = "0" ]; then
                curValue="true";
                if [ "$enableADController" = "true" ]; then
                    curValue="false";
                fi;
                sed -i "s/enable-controller-attach-detach=$curValue/enable-controller-attach-detach=$enableADController/" /host/etc/systemd/system/kubelet.service.d/10-kubeadm.conf;
                restartKubelet="true";
                echo "current value is $curValue, change to expect "$enableADController;
            fi;
            if [ "$restartKubelet" = "true" ]; then
                /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload;
                /nsenter --mount=/proc/1/ns/mnt service kubelet restart;
                echo "restart kubelet";
            fi;
            while true;
            do
                sleep 5;
            done;
          volumeMounts:
          - name: etc
            mountPath: /host/etc
      volumes:
        - name: etc
          hostPath:
            path: /etc
```

## 部署 csi 插件：

登陆下面地址查看最新CSI发行版本：
[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/releases)

部署模板地址：
[https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/release-v1.1.0/deploy/ack](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/release-v1.1.0/deploy/ack)

其中 {release-v1.1.0} 为版本号，需要根据最新版本进行调整；

部署模板分为下面4个文件，您只需要部署其中两个即可，说明如下：

```
csi-plugin-dedicate.yaml：阿里云专有集群，部署此模板；
csi-plugin-managed.yaml：阿里云托管集群，部署此模板；
csi-provisioner-dedicate.yaml：阿里云专有集群，部署此模板；
csi-provisioner-managed.yaml：阿里云托管集群，部署此模板；
```

执行CSI部署命令：

```
kubectl apply -f csi-plugin-***.yaml
kubectl apply -f csi-provisioner-***.yaml
```
