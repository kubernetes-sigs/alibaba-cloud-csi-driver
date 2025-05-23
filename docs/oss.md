## OSS CSI Plugin

An OSS CSI plugin is available to help simplify storage management.
Object Storage Service (OSS) volumes can be mounted to multiple pods. OSS is suitable for data that is frequently read but does not require high disk IOPS, such as configuration files, video files, and images.
You can mount a statically provisioned OSS volume or a dynamically provisioned OSS volume according to your needs.

## Install the CSI driver

Please refer to the [installation guide](./install.md) for detailed instructions.

## Usage

### Prerequisite

* A working Kubernetes cluster.
* Local `kubectl` configured to communicate with this cluster.
* CSI plugin with OSS driver and controller enabled.
* An OSS bucket is created. A subuser account with policies listed below is required.
  For ReadOnly access, the policy is:
  ```
  {
    "Statement": [
        {
            "Action": [
                "oss:Get*",
                "oss:List*"
            ],
            "Effect": "Allow",
            "Resource": [
                "acs:oss:*:*:<oss-bucket-name>",
                "acs:oss:*:*:<oss-bucket-name>/*"
            ]
        }
    ],
    "Version": "1"
  }
  ```
  For ReadWrite access, the policy is:
  ```
  {
    "Statement": [
        {
            "Action": "oss:*",
            "Effect": "Allow",
            "Resource": [
                "acs:oss:*:*:<oss-bucket-name>",
                "acs:oss:*:*:<oss-bucket-name>/*"
            ]
        }
    ],
    "Version": "1"
  }
  ```

### Features

* Oss Plugin support to use Secret for Access Authorization;
* Oss Plugin support to mount remote subpath under oss bucket;
* Oss Plugin support to upgrade online.

### Step 1: Create pv/pvc/deploy with csi

#### Mount a statically provisioned OSS volume 

```shell
kubectl create -f ./examples/oss/secret.yaml
kubectl create -f ./examples/oss/static/pv.yaml
kubectl create -f ./examples/oss/static/pvc.yaml
kubectl create -f ./examples/oss/deploy.yaml
```

Parameters in the `nodePublishSecretRef` section of PV:
| Parameter        | Description |
|------------------|-------------|
| name | The name of the Secret that stores the AccessKey pair information. |
| namespace | The namespace of the Secret that stores the AccessKey pair information. | 

Parameters in the `volumeAttributes` section of PV:
| Parameter        | Description |
|------------------|-------------|
| bucket | The OSS bucket that you want to mount. |
| path | The path relative to the root directory of the OSS bucket to be mounted. The default value is /. If release version is earlier than v1.6.0, you must create the path in the OSS bucket in advance. | 
| url | The endpoint of the OSS bucket you want to mount. You can retrieve the endpoint from the Overview page of the bucket in the OSS console. |
| fuseType | The type of FUSE client, default is ossfs. Set the value to `ossfs2` if you use ossfs 2.0 to mount the volume. |
| otherOpts | You can configure custom parameters in the -o *** -o *** format for the OSS volume. Example: -o umask=022 -o max_stat_cache_size=0 -o allow_other. For more information, see [Options Supported by ossfs 1.0](https://www.alibabacloud.com/help/en/oss/developer-reference/common-options) and [Options Supported by ossfs 2.0](https://www.alibabacloud.com/help/en/oss/developer-reference/description-of-mount-options) |

#### Mount a dynamically provisioned OSS volume 

```shell
kubectl create -f ./examples/oss/secret.yaml
kubectl create -f ./examples/oss/dynamic/sc.yaml
kubectl create -f ./examples/oss/dynamic/pvc.yaml
kubectl create -f ./examples/oss/deploy.yaml
```

Parameters in the `parameters` section of StorageClass:
| Parameter        | Description |
|------------------|-------------|
| name | The name of the Secret that stores the AccessKey pair information. |
| namespace | The namespace of the Secret that stores the AccessKey pair information. | 

The rest of the parameters are the same as the parameters in the `volumeAttributes` section of a static PV.

### Step 2: Check status of PVC/PV

#### Check pvc status
```shell
kubectl get pvc oss-pvc
```
Expected output:
```
oss-pvc    Bound    oss-csi-pv      5Gi        RWO        5m18s
```

#### Check pod status
```shell
kubectl get pod -l app=nginx
```
Expected output:
```
NAME                          READY   STATUS    RESTARTS   AGE
deployment-oss-*****-n****    1/1     Running   0          32s
```

#### Check data in OSS volume
```shell
kubectl exec -it deployment-oss-*****-n**** -- ls /data/
```
Expected output is the data in the specified OSS bucket.