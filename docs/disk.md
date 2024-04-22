# Disk CSI Driver

Disk CSI driver is available to help simplify storage management.
Once user creates PVC with the reference to a Disk storage class, disk and
corresponding PV object get dynamically created and become ready to be used by
workloads.

## Features Support

**Disk Snapshot:** [disk-snapshot](./disk-snapshot-restore.md)

**Block Volumes:** [disk-block](./disk-block.md)

**Resize Volume:** [disk-shared](./disk-resizer.md)

## Configuration Requirements

* Authorizations to access related cloud resources
* StorageClass with `diskplugin.csi.alibabacloud.com` as provisioner name.
* Service Accounts with required RBAC permissions

## Demo

[![](demo.png)](http://cloud.video.taobao.com/play/u/1962692024/p/1/e/6/t/1/50224108448.mp4)


## How to Use

We recommend [Alibaba Cloud Container Service for Kubernetes (ACK)](https://www.alibabacloud.com/product/kubernetes), which will deploy CSI drivers for you automatically.

Alternatively, you can try to deploy it manually on [ECS](https://www.alibabacloud.com/product/ecs), which is covered by this document. We do not provide commercial support for such deployment, however.

### Requirements

* A working Kubernetes cluster deployed on ECS.
* Local `kubectl` configured to communicate with this cluster.

### Step 1: Create RAM User with Required Permissions

CSI will create disks, attach disk to ECS, etc., on behalf of you. To do so, we need authorizations.

1. Create a RAM user, enable OpenAPI access. Once the user is created, record the AccessKey ID and AccessKey Secret.
2. Create a policy, paste in the following script
   ```json
   {
       "Version": "1",
       "Statement": [
           {
               "Action": [
                   "ecs:AttachDisk",
                   "ecs:DetachDisk",
                   "ecs:DescribeDisks",
                   "ecs:CreateDisk",
                   "ecs:ResizeDisk",
                   "ecs:CreateSnapshot",
                   "ecs:DeleteSnapshot",
                   "ecs:CreateAutoSnapshotPolicy",
                   "ecs:ApplyAutoSnapshotPolicy",
                   "ecs:CancelAutoSnapshotPolicy",
                   "ecs:DeleteAutoSnapshotPolicy",
                   "ecs:DescribeAutoSnapshotPolicyEX",
                   "ecs:ModifyAutoSnapshotPolicyEx",
                   "ecs:AddTags",
                   "ecs:RemoveTags",
                   "ecs:DescribeTags",
                   "ecs:DescribeSnapshots",
                   "ecs:ListTagResources",
                   "ecs:TagResources",
                   "ecs:UntagResources",
                   "ecs:ModifyDiskSpec",
                   "ecs:CreateSnapshot",
                   "ecs:DescribeSnapshotGroups",
                   "ecs:CreateSnapshotGroup",
                   "ecs:DeleteSnapshotGroup",
                   "ecs:CopySnapshot",
                   "ecs:DeleteDisk",
                   "ecs:DescribeInstanceAttribute",
                   "ecs:DescribeInstanceHistoryEvents",
                   "ecs:DescribeTaskAttribute",
                   "ecs:DescribeInstances"
               ],
               "Resource": [
                   "*"
               ],
               "Effect": "Allow"
           },
           {
               "Action": [
                   "kms:ListAliases"
               ],
               "Resource": [
                   "*"
               ],
               "Effect": "Allow"
           }
       ]
   }
   ```
3. Authorize the new policy for the new RAM user.

### Step 2: Config the AccessKey to Cluster

Store the access key of just created RAM user in the cluster.
```shell
kubectl create secret -n kube-system generic csi-access-key \
    --from-literal=id='LTA******************GWN' \
    --from-literal=secret='***********'
```

### Step 3: Apply Necessary Resources

Create the service account, RBAC role, deploy the CSI provisioner and CSI plugin.
```shell
helm install csi-driver ./deploy/chart --values deploy/chart/values-ecs.yaml --namespace kube-system
```

Note: this will also deploy OSS and NAS CSI driver. Edit the manifests before applying if you want to disable them.

### Step 4: Create StorageClass
Storage class is necessary for dynamic volume provisioning.
We already provided some predefined storage classes in the previous step. For more advanced features, please refer to [Aliyun docs](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/use-dynamically-provisioned-disk-volumes?#6d16e8a415nie).

### Step 5: Check the Status of CSI driver

Checks that all pods are running and ready.
```shell
kubectl get pods -n kube-system -l app=csi-plugin
```
Expected output:
```
NAME               READY   STATUS    RESTARTS   AGE
csi-plugin-jmxz8   4/4     Running   0          170m
```
```shell
kubectl get pods -n kube-system -l app=csi-provisioner
```
Expected output:
```
NAME                               READY   STATUS    RESTARTS   AGE
csi-provisioner-76fcb8b894-5gmc2   9/9     Running   0          7d8h
csi-provisioner-76fcb8b894-mlgj5   9/9     Running   0          7d8h
```

### Step 6: Test the Deployment

To make sure your CSI plugin is working, create a simple workload to test it out:
```shell
kubectl apply -f examples/disk/workload.yaml
```

Now check that pod `test-disk` is running, a new disk is created, attached, formatted, and mounted into the new pod.
After you are done, remove the test workload:
```shell
kubectl delete -f examples/disk/workload.yaml
```

Enjoy!
