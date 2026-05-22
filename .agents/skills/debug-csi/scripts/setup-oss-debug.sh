#!/bin/bash
set -e

REPO_ROOT="$(git rev-parse --show-toplevel)"
VALUES_FILE="$REPO_ROOT/.local/test-values.yaml"

if [ ! -f "$VALUES_FILE" ]; then
    echo "Error: $VALUES_FILE not found. Run /init-dev-env first." >&2
    exit 1
fi

REGISTRY=$(yq '.images.registry' "$VALUES_FILE")
NAMESPACE=$(yq '.images.controller.repo | split("/") | .[0]' "$VALUES_FILE")
SECRET_NAME=$(yq '.imagePullSecrets[0]' "$VALUES_FILE")

IMAGE_REPOSITORY_PREFIX="$REGISTRY/$NAMESPACE"
IMAGE_TAG="${1:-latest}"

echo "Registry prefix: $IMAGE_REPOSITORY_PREFIX"
echo "Image tag:       $IMAGE_TAG"
echo "Pull secret:     $SECRET_NAME"

echo "--- Setting OSSFS image tags via configmap ---"
kubectl apply --server-side --field-manager=debug-ossfs -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: csi-plugin
  namespace: kube-system
data:
  fuse-ossfs: |
    image-tag=$IMAGE_TAG
  fuse-ossfs2: |
    image-tag=$IMAGE_TAG
EOF

echo "--- Patching csi-provisioner IMAGE_REPOSITORY_PREFIX ---"
kubectl -n kube-system patch deploy/csi-provisioner -p "
spec:
  template:
    spec:
      containers:
      - name: csi-provisioner
        env:
        - name: IMAGE_REPOSITORY_PREFIX
          value: $IMAGE_REPOSITORY_PREFIX
"

echo "--- Patching ack-csi-fuse service accounts ---"
kubectl -n ack-csi-fuse patch sa default -p "{\"imagePullSecrets\": [{\"name\": \"$SECRET_NAME\"}]}"
kubectl -n ack-csi-fuse patch sa csi-fuse-ossfs -p "{\"imagePullSecrets\": [{\"name\": \"$SECRET_NAME\"}]}"

echo "Done."
