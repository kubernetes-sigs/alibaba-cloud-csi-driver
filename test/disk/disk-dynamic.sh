#! /bin/sh

testRes="false"

cat << EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: csi-pvc
spec:
  accessModes:
  - ReadWriteOnce
  storageClassName: alicloud-disk-available
  resources:
    requests:
      storage: 25Gi
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
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
          name: web
        volumeMounts:
        - name: pvc-disk
          mountPath: /data
      volumes:
        - name: pvc-disk
          persistentVolumeClaim:
            claimName: csi-pvc
EOF

for (( i=0; i<20; i++));
do
    lines=`kubectl get pod | grep web-0 | grep Running | wc -l`
    if [ "$lines" = "1" ]; then
        testRes="true"
        break
    fi
    sleep 5
done

kubectl delete sts web
kubectl delete pvc csi-pvc


if [ "$testRes" = "true" ]; then
    echo -e "****************************\n** Dynamic Disk Test Successful\n****************************\n"
else
    echo -e "****************************\n** Dynamic Disk Test Failed\n****************************\n"
fi