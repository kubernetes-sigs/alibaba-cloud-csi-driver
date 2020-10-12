#!/bin/bash

## Usage
## Append image tag which is expect.
## sh upgrade_csi-provisioner.sh v1.14.8.38-fe611ad1-aliyun


if [ "$1" = "" ]; then
  echo "Please input the expect csi provisioner version... "
fi
imageVersion=$1

echo `date`" Start to Upgrade CSI Provisioner to $imageVersion ..."


masterCount=`kubectl get node | grep master |grep -v grep | wc -l`
volumeDefineStr=""
volumeMountStr=""
if [ "$masterCount" -eq "0" ]; then
  volumeDefineStr="        - name: addon-token\n          secret:\n            defaultMode: 420\n            items:\n            - key: addon.token.config\n              path: token-config\n            secretName: addon.csi.token"
  volumeMountStr="            - mountPath: \/var\/addon\n              name: addon-token\n              readOnly: true"
fi


# new deploy template file
# if any changes, just update here and replace image value
cat > .aliyun-csi-provisioner.yaml << EOF
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
      priorityClassName: system-node-critical
      serviceAccount: admin
      hostNetwork: true
      containers:
        - name: snapshot-controller
          image: csi-image-prefix/plugins/snapshot-controller:v2.0.1
          args:
            - "--v=5"
            - "--leader-election=false"
          imagePullPolicy: Always
        - name: disk-snapshotter
          image: csi-image-prefix/plugins/csi-snapshotter:v2.1.1
          args:
            - "--csi-address=\$(ADDRESS)"
          env:
            - name: ADDRESS
              value: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /var/lib/kubelet/csi-provisioner/diskplugin.csi.alibabacloud.com/
        - name: external-disk-provisioner
          image: csi-image-prefix/acs/csi-provisioner:v1.4.0-aliyun
          args:
            - "--provisioner=diskplugin.csi.alibabacloud.com"
            - "--csi-address=\$(ADDRESS)"
            - "--feature-gates=Topology=True"
            - "--volume-name-prefix=disk"
            - "--strict-topology=true"
            - "--timeout=150s"
            - "--enable-leader-election=true"
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
            timeoutSeconds: 3
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
              cpu: 100m
              memory: 100Mi
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
        - name: etc
          hostPath:
            path: /etc
volume-define-string
EOF

# get registry address
imageBefore=`kubectl describe ds csi-plugin -nkube-system | grep Image: | awk 'NR==2 {print $2}'`
imagePrefix=`kubectl describe ds csi-plugin -nkube-system | grep Image: | awk 'NR==2 {print $2}' | awk -F: '{print $1}' | awk -F/ '{print $1}'`

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

## Delete old plugin
kubectl delete deployment csi-provisioner -nkube-system

## do csi-provisioner upgrade
cat .aliyun-csi-provisioner.yaml | sed "s/csi-image-prefix/$imagePrefix/" | sed "s/csi-image-version/$imageVersion/" | sed "s/volume-define-string/$volumeDefineStr/" | sed "s/volume-mount-string/$volumeMountStr/" | kubectl apply -f -

echo "Upgrade csi provisioner from $imageBefore to $imageName"