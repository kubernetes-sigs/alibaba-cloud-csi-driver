
#! /bin/sh

testRes="false"

cat << EOF | kubectl create -f -
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
 name: alicloud-disk-aware
provisioner: diskplugin.csi.alibabacloud.com
parameters:
  type: cloud_ssd
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
EOF

cat << EOF | kubectl create -f -
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-csi-aware
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: alicloud-disk-aware
  resources:
    requests:
      storage: 20Gi
EOF

for (( i=0; i<4; i++));
do
    lines=`kubectl get pvc | grep pvc-csi-aware | grep Pending | wc -l`
    if [ "$lines" = "1" ]; then
        testRes="true"
    else
        testRes="false"
    fi
    sleep 5
done


cat << EOF | kubectl create -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dynamic-aware
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
          - name: disk-pvc
            mountPath: "/data"
      volumes:
        - name: disk-pvc
          persistentVolumeClaim:
            claimName: pvc-csi-aware
EOF

for (( i=0; i<10; i++));
do
    lines=`kubectl get pod | grep dynamic-aware | grep Running | wc -l`
    if [ "$lines" = "1" ]; then
        testRes="true"
    else
        testRes="false"
    fi
    sleep 5
done

kubectl delete deploy dynamic-aware
kubectl delete pvc pvc-csi-aware
kubectl delete sc alicloud-disk-aware


if [ "$testRes" = "true" ]; then
    echo -e "****************************\n** Aware Disk Test Successful\n****************************\n"
    exit 0
else
    echo -e "****************************\n** Aware Disk Test Failed\n****************************\n"
    exit 1
fi
