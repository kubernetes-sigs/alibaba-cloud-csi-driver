#! /bin/sh

ossakid=$1
ossaksec=$2
testRes="false"

cat << EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolume
metadata:
  name: oss-csi-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: ossplugin.csi.alibabacloud.com
    volumeHandle: oss-pv
    volumeAttributes:
      bucket: "test"
      url: "oss-cn-hangzhou.aliyuncs.com"
      otherOpts: "-o max_stat_cache_size=1 -o allow_other"
      akId: "$ossakid"
      akSecret: "$ossaksec"
EOF

pvline=`kubectl describe pv oss-csi-pv | grep $ossakid | wc -l`
if [ "$pvline" != "1" ]; then
    testRes="false"
fi

cat << EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: oss-pvc
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: oss-deployment
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
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
          - name: oss-pvc
            mountPath: "/data"
      volumes:
        - name: oss-pvc
          persistentVolumeClaim:
            claimName: oss-pvc
EOF

for (( i=0; i<20; i++));
do
    lines=`kubectl get pod | grep oss-deployment | grep Running | wc -l`
    if [ "$lines" = "1" ]; then
        testRes="true"
        break
    fi
    sleep 5
done

kubectl delete deploy oss-deployment
kubectl delete pvc oss-pvc
kubectl delete pv oss-csi-pv

if [ "$testRes" = "true" ]; then
    echo -e "****************************\n** Static OSS Test Successful\n****************************\n"
else
    echo -e "****************************\n** Static OSS Test Failed\n****************************\n"
fi