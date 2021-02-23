
# 将CSI-Provisioner 从选主模式变成非选主模式

## 检查您的集群是为托管集群还是专有集群？

您可以到ACK控制台看集群属性，也可以通过下面命令来判断：

> 托管集群：kubectl get node | grep master 为空，即不存在master节点； 适用本文档；
>
> 专有集群：kubectl get node | grep master 非空，即有多个master节点； 不适用本文档；


## 配置RBAC：

如果集群中已经存在csi-admin这个ServiceAccount，可以不执行这步操作；

```
$ kubectl get sa -nkube-system | grep csi-admin
```

配置RBAC命令：

```
$ kubectl apply -f rbac.yaml
```

RBAC配置参考：https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/rbac.yaml


## 部署非选主CSI-Provsisioner：

下面模板为非选主模式：

```
$ kubectl apply -f csi-provisioner.yaml
```

csi-provisioner非选主版配置参考：https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/blob/master/deploy/ack/noleader/csi-provisioner-managed.yaml
