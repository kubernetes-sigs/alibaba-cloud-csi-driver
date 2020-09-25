#!/bin/bash

## Usage
## Append image tag which is expect.
## sh upgrade_csi-provisioner.sh v1.14.8.41-9efe2ede-aliyun

imageVersion=$1
if [ "$imageVersion" = "" ]; then
  echo "CSI provisioner version set to default... "
  imageVersion=v1.14.8.41-9efe2ede-aliyun
fi

echo `date`" Start to Upgrade CSI Provisioner to $imageVersion ..."

## set secret mounts for managed cluster
masterCount=`kubectl get node | grep master |grep -v grep | wc -l`
volumeDefineStr=""
volumeMountStr=""
if [ "$masterCount" = "0" ]; then
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
      tolerations:
      - effect: NoSchedule
        operator: Exists
        key: node-role.kubernetes.io/master
      - effect: NoSchedule
        operator: Exists
        key: node.cloudprovider.kubernetes.io/uninitialized
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
              value: /socketDir/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /socketDir
        - name: external-disk-attacher
          image: csi-image-prefix/acs/csi-attacher:v2.1.0
          args:
            - "--v=5"
            - "--csi-address=\$(ADDRESS)"
            - "--leader-election=true"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /csi
        - name: external-disk-resizer
          image: csi-image-prefix/acs/csi-resizer:v0.3.0
          args:
            - "--v=5"
            - "--csi-address=\$(ADDRESS)"
            - "--leader-election"
          env:
            - name: ADDRESS
              value: /socketDir/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /socketDir/
        - name: csi-diskprovisioner
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
            - "--driver=diskplugin.csi.alibabacloud.com"
          env:
            - name: CSI_ENDPOINT
              value: unix://socketDir/csi.sock
            - name: MAX_VOLUMES_PERNODE
              value: "15"
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
          - containerPort: 9808
            name: healthz
            protocol: TCP
          volumeMounts:
            - mountPath: /dev
              mountPropagation: "HostToContainer"
              name: host-dev
            - mountPath: /var/log/
              name: host-log
            - mountPath: /socketDir/
              name: disk-provisioner-dir
            - name: etc
              mountPath: /host/etc
volume-mount-string
          resources:
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 100m
              memory: 100Mi
        - name: disk-liveness-probe
          image: csi-image-prefix/acs/csi-livenessprobe:v2.0.0
          args:
            - --csi-address=/csi/csi.sock
            - --health-port=9808
          volumeMounts:
            - name: disk-provisioner-dir
              mountPath: /csi
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
              value: /socketDir/csi.sock
          imagePullPolicy: "Always"
          volumeMounts:
            - name: nas-provisioner-dir
              mountPath: /socketDir
        - name: csi-nasprovisioner
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
            - "--driver=nasplugin.csi.alibabacloud.com"
          env:
            - name: CSI_ENDPOINT
              value: unix://socketDir/csi.sock
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
          - containerPort: 9809
            name: healthz
            protocol: TCP
          volumeMounts:
            - mountPath: /var/log/
              name: host-log
            - mountPath: /socketDir/
              name: nas-provisioner-dir
            - name: etc
              mountPath: /host/etc
volume-mount-string
          resources:
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 100m
              memory: 100Mi
        - name: nas-liveness-probe
          image: csi-image-prefix/acs/csi-livenessprobe:v2.0.0
          args:
            - --csi-address=/csi/csi.sock
            - --health-port=9809
          volumeMounts:
            - name: nas-provisioner-dir
              mountPath: /csi
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
kubectl delete sts csi-provisioner -nkube-system

## do csi-provisioner upgrade
cat .aliyun-csi-provisioner.yaml | sed "s/csi-image-prefix/$imagePrefix/" | sed "s/csi-image-version/$imageVersion/" | sed "s/volume-define-string/$volumeDefineStr/" | sed "s/volume-mount-string/$volumeMountStr/" | kubectl apply -f -

echo "Upgrade csi provisioner to "$imageName