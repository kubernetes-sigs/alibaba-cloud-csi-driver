#!/bin/bash

## Usage
## Append image tag which is expect.
## sh upgrade_csi-plugin.sh v1.14.8.41-9efe2ede-aliyun


if [ "$1" = "" ]; then
  echo "Please input the expect csi plugin version... "
fi
imageVersion=$1
echo `date`" Start to Upgrade CSI Provisioner to $imageVersion ..."

# new deploy template file
# if any changes, just update here and replace image value
cat > .aliyun-csi-plugin.yaml << EOF
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: diskplugin.csi.alibabacloud.com
spec:
  attachRequired: true
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: nasplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: ossplugin.csi.alibabacloud.com
spec:
  attachRequired: false
  podInfoOnMount: true
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-plugin
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: csi-plugin
  template:
    metadata:
      labels:
        app: csi-plugin
    spec:
      tolerations:
        - operator: Exists
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
      nodeSelector:
        beta.kubernetes.io/os: linux
      priorityClassName: system-node-critical
      serviceAccount: admin
      hostNetwork: true
      hostPID: true
      containers:
        - name: disk-driver-registrar
          image: csi-image-prefix/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
            - "--v=5"
            - "--csi-address=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
            - name: registration-dir
              mountPath: /registration
        - name: csi-diskplugin
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
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
            - name: CSI_ENDPOINT
              value: unix://var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock
            - name: ACCESS_KEY_ID
              value: ""
            - name: ACCESS_KEY_SECRET
              value: ""
            - name: MAX_VOLUMES_PERNODE
              value: "15"
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 3
            failureThreshold: 5
          ports:
          - containerPort: 9810
            name: healthz
            protocol: TCP
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
              mountPropagation: "Bidirectional"
            - name: container-dir
              mountPath: /var/lib/container
              mountPropagation: "Bidirectional"
            - mountPath: /dev
              mountPropagation: "HostToContainer"
              name: host-dev
            - mountPath: /var/log/
              name: host-log
            - name: etc
              mountPath: /host/etc
        - name: disk-liveness-probe
          image: csi-image-prefix/acs/csi-livenessprobe:v2.0.0
          args:
            - --csi-address=/var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com/csi.sock
            - --health-port=9810
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet

        - name: nas-driver-registrar
          image: csi-image-prefix/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
          - "--v=5"
          - "--csi-address=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
          - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
          - name: kubelet-dir
            mountPath: /var/lib/kubelet/
          - name: registration-dir
            mountPath: /registration

        - name: csi-nasplugin
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
          - name: KUBE_NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: CSI_ENDPOINT
            value: unix://var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 3
            failureThreshold: 5
          ports:
          - containerPort: 9811
            name: healthz
            protocol: TCP
          volumeMounts:
          - name: kubelet-dir
            mountPath: /var/lib/kubelet/
            mountPropagation: "Bidirectional"
          - mountPath: /var/log/
            name: host-log
        - name: nas-liveness-probe
          image: csi-image-prefix/acs/csi-livenessprobe:v2.0.0
          args:
            - --csi-address=/var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com/csi.sock
            - --health-port=9811
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet

        - name: oss-driver-registrar
          image: csi-image-prefix/acs/csi-node-driver-registrar:v1.2.0
          imagePullPolicy: Always
          args:
          - "--v=5"
          - "--csi-address=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
          - "--kubelet-registration-path=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock"
          volumeMounts:
          - name: kubelet-dir
            mountPath: /var/lib/kubelet/
          - name: registration-dir
            mountPath: /registration

        - name: csi-ossplugin
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
          - "--driver=ossplugin.csi.alibabacloud.com"
          env:
          - name: CSI_ENDPOINT
            value: unix://var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 3
            failureThreshold: 5
          ports:
          - containerPort: 9812
            name: healthz
            protocol: TCP
          volumeMounts:
          - name: kubelet-dir
            mountPath: /var/lib/kubelet/
            mountPropagation: "Bidirectional"
          - name: etc
            mountPath: /host/etc
          - mountPath: /var/log/
            name: host-log
          - mountPath: /host/usr/
            name: ossconnectordir
        - name: oss-liveness-probe
          image: csi-image-prefix/acs/csi-livenessprobe:v2.0.0
          args:
            - --csi-address=/var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com/csi.sock
            - --health-port=9812
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
      volumes:
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry
            type: DirectoryOrCreate
        - name: container-dir
          hostPath:
            path: /var/lib/container
            type: DirectoryOrCreate
        - name: kubelet-dir
          hostPath:
            path: /var/lib/kubelet
            type: Directory
        - name: host-dev
          hostPath:
            path: /dev
        - name: host-log
          hostPath:
            path: /var/log/
        - name: etc
          hostPath:
            path: /etc
        - name: ossconnectordir
          hostPath:
            path: /usr/
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 10%
    type: RollingUpdate
EOF

## Set target csi-plugin version
imageBefore=`kubectl describe ds csi-plugin -nkube-system | grep Image: | awk 'NR==2 {print $2}'`
imagePrefix=`kubectl describe ds csi-plugin -nkube-system | grep Image: | awk 'NR==2 {print $2}' | awk -F: '{print $1}' | awk -F/ '{print $1}'`
imageName=${imagePrefix}/acs/csi-plugin:$imageVersion

## Check current version
if [[ $imagePrefix != registry* ]] || [[ $imagePrefix != *aliyuncs.com ]]; then
    echo "Current image is not expect: "$imagePrefix
    exit 0
fi

## Delete old plugin
kubectl delete csidriver diskplugin.csi.alibabacloud.com nasplugin.csi.alibabacloud.com ossplugin.csi.alibabacloud.com
kubectl delete ds csi-plugin -nkube-system

## do csi-plugin upgrade
cat .aliyun-csi-plugin.yaml | sed "s/csi-image-prefix/$imagePrefix/" | sed "s/csi-image-version/$imageVersion/" | kubectl apply -f -

echo "Upgrade csi From "$imageBefore" to "$imageName
