## Memory CSI Plugin

An Memory CSI plugin is available to help simplify storage management for memory type volume. You can create a pv with csi configuration, and the pvc, pod defines as usual.

Memory volume is temp storage, and similar with Empty volume. Memory storage is created with pod starting and removed with pod terminatting.

Memory volume is just like Empty volume in memory type, and difference is that Memory volume support to limit memory size.

## Configuration Requirements

* Service Accounts with required RBAC permissions

## Feature Status
Alpha

## Compiling and Package
memplugin.csi.alibabacloud.com can be compiled in a form of a container.

To build a container:
```
$ cd build && sh build-mem.sh
```

## Usage

### Prerequisite
Same as diskplugin.csi.alibabacloud.com;


### Step 1: Create CSI Provisioner
```
# kubectl create -f ./deploy/mem/mem-provisioner.yaml
```

### Step 2: Create CSI Plugin
```
# kubectl create -f ./deploy/mem/mem-plugin.yaml
```

> Note: The plugin log style can be configured by environment variable: LOG_TYPE.

> "host": logs will be printed into files which save to host(/var/log/alicloud/memplugin.csi.alibabacloud.com.log);

> "stdout": logs will be printed to stdout, can be printed by docker logs or kubectl logs.

> "both": default option, log will be printed both to stdout and host file.

### Step 3: Create storageclass
```
# kubectl create -f ./deploy/mem/storageclass.yaml
```

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: csi-mem
provisioner: memplugin.csi.alibabacloud.com
reclaimPolicy: Delete
```

### Step 4: Create nginx deploy with memory volume see usage of memory volue [examples](https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/master/examples/mem) 
```
# kubectl create -f ./examples/mem/pvc.yaml
# kubectl create -f ./examples/mem/deploy.yaml
```

### Step 5: Check status of PVC/PV
```
# kubectl get pvc | grep mem
mem-pvc        Bound    mem-ad6cb1b2-27e1-11ea-bfb5-00163e005b53    2Gi        RWO            csi-mem                         7m56s

# kubectl get pv | grep mem
mem-ad6cb1b2-27e1-11ea-bfb5-00163e005b53    2Gi        RWO            Delete           Bound    default/mem-pvc        csi-mem                                  7m49s

```

### Step 6: check status of pod
check directory in pod

```
# kubectl get pod
NAME                               READY   STATUS    RESTARTS   AGE
deployment-mem-5d6c4b6587-h68z5    1/1     Running   0          7m42s

# kubectl exec -ti deployment-mem-5d6c4b6587-h68z5 sh
# df -h | grep data
tmpfs           2.0G     0  2.0G   0% /data

```

Check Directory in host:

```
# kubectl describe pod deployment-mem-5d6c4b6587-h68z5 | grep Node
Node:               cn-beijing.172.16.1.102/172.16.1.102

# mount | grep csi | grep mem
tmpfs on /var/lib/kubelet/pods/c9fb2065-27e1-11ea-8491-00163e165b60/volumes/kubernetes.io~csi/mem-ad6cb1b2-27e1-11ea-bfb5-00163e005b53/mount type tmpfs (rw,relatime,size=2097152k)

```
