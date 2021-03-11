#!/bin/bash

## Usage
## Append image tag which is expect.
## bash upgrade_csi-provisioner.sh v1.18.8.45-1c5d2cd1-aliyun


if [ "$1" = "" ]; then
  echo "Please input the expect csi provisioner version... "
fi
imageVersion=$1

echo `date`" Start to Upgrade CSI Provisioner to $imageVersion ..."

masterCount=`kubectl get node | grep master |grep -v grep | wc -l`
secretCount=`kubectl get secret addon.csi.token -nkube-system | grep addon.csi.token | grep -v grep | wc -l`
volumeDefineStr=""
volumeMountStr=""
if [ "$masterCount" -eq "0" ] && [ "$secretCount" -eq "1" ]; then
  volumeDefineStr="        - name: addon-token\n          secret:\n            defaultMode: 420\n            items:\n            - key: addon.token.config\n              path: token-config\n            secretName: addon.csi.token"
  volumeMountStr="            - mountPath: \/var\/addon\n              name: addon-token\n              readOnly: true"
fi


# new deploy template file
# if any changes, just update here and replace image value

cat > .aliyun-csi-provisioner-addition.yaml << EOF
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-available
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-essd
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_essd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-ssd
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_ssd
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-efficiency
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: cloud_efficiency
reclaimPolicy: Delete
allowVolumeExpansion: true
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: alicloud-disk-topology
provisioner: diskplugin.csi.alibabacloud.com
parameters:
    type: available
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
EOF

cat > .aliyun-csi-provisioner.yaml << EOF
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-provisioner
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-provisioner
  strategy:
    rollingUpdate:
      maxSurge: 0
      maxUnavailable: 1
    type: RollingUpdate
  replicas: 2
  template:
    metadata:
      labels:
        app: csi-provisioner
    spec:
      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            preference:
              matchExpressions:
              - key: node-role.kubernetes.io/master
                operator: Exists
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - csi-provisioner
            topologyKey: kubernetes.io/hostname
      tolerations:
      - effect: NoSchedule
        operator: Exists
        key: node-role.kubernetes.io/master
      - effect: NoSchedule
        operator: Exists
        key: node.cloudprovider.kubernetes.io/uninitialized
      priorityClassName: system-node-critical
      serviceAccount: csi-admin
      hostNetwork: true
      containers:
        - name: external-disk-provisioner
          image: csi-image-prefix/acs/csi-provisioner:v1.6.0-b6f763a43-aliyun
          args:
            - "--provisioner=diskplugin.csi.alibabacloud.com"
            - "--csi-address=\$(ADDRESS)"
            - "--feature-gates=Topology=True"
            - "--volume-name-prefix=disk"
            - "--strict-topology=true"
            - "--timeout=150s"
            - "--enable-leader-election=true"
            - "--extra-create-metadata=true"
            - "--leader-election-type=leases"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-attacher
          image: csi-image-prefix/acs/csi-attacher:v2.1.0
          args:
            - "--v=5"
            - "--csi-address=\$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-disk-resizer
          image: csi-image-prefix/acs/csi-resizer:v0.3.0
          args:
            - "--v=5"
            - "--csi-address=\$(ADDRESS)"
            - "--leader-election"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
        - name: external-nas-provisioner
          image: csi-image-prefix/acs/csi-provisioner:v1.4.0-aliyun
          args:
            - "--provisioner=nasplugin.csi.alibabacloud.com"
            - "--csi-address=\$(ADDRESS)"
            - "--volume-name-prefix=nas"
            - "--timeout=150s"
            - "--enable-leader-election=true"
            - "--leader-election-type=leases"
            - "--retry-interval-start=500ms"
            - "--v=5"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
        - name: external-csi-snapshotter
          image: csi-image-prefix/acs/csi-snapshotter:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--csi-address=\$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          imagePullPolicy: Always
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /csi
        - name: external-snapshot-controller
          image: csi-image-prefix/acs/snapshot-controller:v3.0.2-1038b92d8-aliyun
          args:
            - "--v=5"
            - "--leader-election=true"
          imagePullPolicy: Always
        - name: csi-provisioner
          securityContext:
            privileged: true
            capabilities:
              add: ["SYS_ADMIN"]
            allowPrivilegeEscalation: true
          image: csi-image-prefix/acs/csi-plugin:csi-image-version
          imagePullPolicy: "Always"
          args:
            - "--endpoint=\$(CSI_ENDPOINT)"
            - "--v=2"
            - "--driver=nas,disk"
          env:
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-provisioner/driverplugin.csi.alibabacloud.com-replace/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
            - name: SERVICE_TYPE
              value: "provisioner"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            timeoutSeconds: 5
            failureThreshold: 5
          ports:
            - name: healthz
              containerPort: 11270
              protocol: TCP
          volumeMounts:
            - name: host-dev
              mountPath: /dev
              mountPropagation: "HostToContainer"
            - name: host-log
              mountPath: /var/log/
            - name: etc
              mountPath: /host/etc
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com
            - name: nas-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/nasplugin.csi.alibabacloud.com
volume-mount-string
          resources:
            limits:
              cpu: 1000m
              memory: 1000Mi
            requests:
              cpu: 100m
              memory: 100Mi
      volumes:
        - name: disk-provisioner-dir
          emptyDir: {}
        - name: nas-provisioner-dir
          emptyDir: {}
        - name: host-log
          hostPath:
            path: /var/log/
        - name: host-dev
          hostPath:
            path: /dev
volume-define-string
        - name: etc
          hostPath:
            path: /etc
EOF

provisionerExist=`kubectl get sts -nkube-system | grep csi-provisioner | awk '{print $1}'`
if [ "$provisionerExist" = "csi-provisioner" ]; then
    kubectl delete sts csi-provisioner -nkube-system
fi

# get registry address
imageBefore=`kubectl describe deploy csi-provisioner -nkube-system | grep Image: | grep acs/csi-plugin | awk 'NR==1 {print $2}'`
imagePrefix=`kubectl describe ds csi-plugin -nkube-system | grep Image: | grep acs/csi-plugin | awk 'NR==1 {print $2}' | awk -F: '{print $1}' | awk -F/ '{print $1}'`

# New image name
imageName=${imagePrefix}/acs/csi-plugin:$imageVersion

## Check current version
if [[ $imagePrefix != registry* ]] || [[ $imagePrefix != *aliyuncs.com ]]; then
    echo "Setting image is not expect: "$imagePrefix
    exit 0
fi

if [[ $imageVersion != *aliyun ]]; then
    echo "current image version is not expect: "$imageVersion
    exit 0
fi

echo "Apply StorageClass..."
kubectl apply -f .aliyun-csi-provisioner-addition.yaml


## Delete old plugin
# kubectl delete deployment csi-provisioner -nkube-system
canApply=`kubectl get deploy csi-provisioner -nkube-system -oyaml |grep kubectl.kubernetes.io/last-applied-configuration | wc -l`
if [ "$canApply" != "0" ]; then
  echo "Upgrade CSI provisioner with kubectl apply..."
  cat .aliyun-csi-provisioner.yaml | sed "s/csi-image-prefix/$imagePrefix/" | sed "s/csi-image-version/$imageVersion/" | sed "s/volume-define-string/$volumeDefineStr/" | sed "s/volume-mount-string/$volumeMountStr/" | kubectl apply -f -
else
  echo "Upgrade CSI provisioner with kubectl replace..."
  cat .aliyun-csi-provisioner.yaml | sed "s/csi-image-prefix/$imagePrefix/" | sed "s/csi-image-version/$imageVersion/" | sed "s/volume-define-string/$volumeDefineStr/" | sed "s/volume-mount-string/$volumeMountStr/" | kubectl replace -f -
fi

echo "Upgrade csi provisioner from $imageBefore to $imageName"
