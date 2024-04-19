## Expand a NAS volume
### Prerequisites
* version of csi-plugin >= 1.18.8.45
* A NAS volume is mounted to a subdirectory

## Usage
### Step 1: Create a StorageClass
provision a NAS volume that has the quota feature and the dynamic expansion feature enabled

```
# kubectl apply -f ./examples/nas/expansion/sc-expansion.yaml
```

### Step 2: Create a statefulset
```
# kubectl apply -f ./examples/nas/expansion/sts.yaml
```

### Step 3: Expand the volume

```
# kubectl patch pvc nas-pvc-0 -p '{"spec":{"resources":{"requests":{"storage":"30Gi"}}}}'
```
### Step 4: Check the expansion
```
# kubectl get pvc
NAME       STATUS   VOLUME                         CAPACITY  ACCESS MODES   STORAGECLASS            AGE
nas-pvc-0  Bound   nas-63c37cc2-b21e-4b56-b26f-****   30Gi      RWM        alicloud-nas-quota-sc   25m10s
```
