# Volume Limit For Disk Plugin

Volume attach limits feature is for the maximum number of volumes that can be attached to one Node.

| Feature           | From | To   | Status |
|-------------------|------|------|--------|
| AttachVolumeLimit | 1.11 | 1.11 | Alpha  |
| AttachVolumeLimit | 1.12 | 1.16 | Beta   |
| AttachVolumeLimit | 1.17 | 1.21 | GA     |

## DiskPlugin Configuration

Config daemonset of csi-diskplugin, Define MAX_VOLUMES_PERNODE to the specified value which means the maximum number of volumes to one node.


```
            - name: MAX_VOLUMES_PERNODE
              value: 5
```
*Because Alibaba Cloud limits 16 disks to one node(ecs), so the value should be less than 16.*

## How to Use

### Introduce:

Volume limit is implemented in kubernetes scheduler and is a predicate check. CSIMaxVolumeLimitChecker is responsible for volume limit check, and decide to schedule pod to which node.

If your application spec contains "nodeName", volume limit will not work.

### Steps:

1. Taint other nodes and only schedule pods to the expected node.

```
# kubectl taint nodes *** key=value:NoSchedule
```

2. Create Deployments:

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: disk-pvc1
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 25Gi
  storageClassName: csi-disk
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment1
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 6
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
        volumeMounts:
          - name: disk-pvc
            mountPath: "/data"
      volumes:
        - name: disk-pvc
          persistentVolumeClaim:
            claimName: disk-pvc1
```

```
nginx-deployment1-7577645657-xl68z    1/1     Running   0          39m
nginx-deployment2-75697b4956-rfzln    1/1     Running   0          38m
nginx-deployment3-855b97bd7-64p6s     1/1     Running   0          38m
nginx-deployment4-85dd7bcb7d-p5gw4    1/1     Running   0          38m
nginx-deployment5-744965cb45-qgbtv    1/1     Running   0          37m
nginx-deployment6-7bc9f88544-gfdkd    0/1     Pending   0          37m
```

As we config volume limit to 5, the 6th pod is in pending status.

Describe pod as below:

```
# kubectl describe pod nginx-deployment6-7bc9f88544-gfdkd
***
Volumes:
  disk-pvc:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  disk-pvc6
    ReadOnly:   false
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason            Age                From               Message
  ----     ------            ----               ----               -------
  Warning  FailedScheduling  40s (x5 over 42s)  default-scheduler  pod has unbound immediate PersistentVolumeClaims
  Warning  FailedScheduling  40s (x2 over 40s)  default-scheduler  0/3 nodes are available: 1 node(s) exceed max volume count, 2 node(s) had taints that the pod didn't tolerate.
```

We can find event "1 node(s) exceed max volume count", which indicates volume limit works.
