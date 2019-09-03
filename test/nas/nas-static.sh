#! /bin/sh

nasserver=$1
testRes="false"

cat << EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-nas
  labels:
    alicloud-pvname: pv-nas
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteMany
  csi:
    driver: nasplugin.csi.alibabacloud.com
    volumeHandle: pv-nas
    volumeAttributes:
      server: "$nasserver"
      path: "/k8s"
      vers: "4.0"
      options: "noresvport,nolock"
EOF

pvline=`kubectl describe pv pv-nas | grep $nasserver | wc -l`
if [ "$pvline" != "1" ]; then
    testRes="false"
fi

cat << EOF | kubectl create -f -
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-nas
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 5Gi
  selector:
    matchLabels:
      alicloud-pvname: pv-nas
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nas-static
  labels:
    app: nginx
spec:
  replicas: 1
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
        image: nginx
        ports:
        - containerPort: 80
        volumeMounts:
          - name: pvc-nas
            mountPath: "/data"
      volumes:
        - name: pvc-nas
          persistentVolumeClaim:
            claimName: pvc-nas
EOF

for (( i=0; i<20; i++));
do
    lines=`kubectl get pod | grep nas-static | grep Running | wc -l`
    if [ "$lines" = "1" ]; then
        testRes="true"
        break
    fi
    sleep 5
done

kubectl delete deploy nas-static
kubectl delete pvc pvc-nas
kubectl delete pv pv-nas

if [ "$testRes" = "true" ]; then
    echo -e "****************************\n** Static Nas Test Successful\n****************************\n"
else
    echo -e "****************************\n** Static Nas Test Failed\n****************************\n"
fi