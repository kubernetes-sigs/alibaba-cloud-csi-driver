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
# Default to alinas image; pass "aio" to use mount-proxy (all-in-one) image
IMAGE_TYPE="${2:-alinas}"

case "$IMAGE_TYPE" in
    alinas)
        IMAGE="$IMAGE_REPOSITORY_PREFIX/csi-alinas:$IMAGE_TAG"
        ARGS='["--socket=/run/cnfs/alinas-mounter.sock", "--driver=alinas"]'
        ;;
    aio)
        IMAGE="$IMAGE_REPOSITORY_PREFIX/mount-proxy:$IMAGE_TAG"
        ARGS='["--socket=/run/cnfs/alinas-mounter.sock", "--driver=alinas,ossfs,ossfs2"]'
        ;;
    *)
        echo "Unknown image type: $IMAGE_TYPE (use 'alinas' or 'aio')" >&2
        exit 1
        ;;
esac

echo "Registry prefix: $IMAGE_REPOSITORY_PREFIX"
echo "Image tag:       $IMAGE_TAG"
echo "Image type:      $IMAGE_TYPE"
echo "Image:           $IMAGE"
echo "Pull secret:     $SECRET_NAME"

echo "--- Enabling AlinasMountProxy feature gate on csi-plugin ---"
CURRENT_ARGS=$(kubectl -n kube-system get ds csi-plugin -o jsonpath='{.spec.template.spec.containers[0].args}')
if echo "$CURRENT_ARGS" | grep -q "AlinasMountProxy"; then
    echo "Feature gate already enabled."
else
    kubectl -n kube-system patch ds csi-plugin --type=json -p '[
      {"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--feature-gates=AlinasMountProxy=true"}
    ]'
    echo "Waiting for csi-plugin to roll out..."
    kubectl -n kube-system rollout status ds/csi-plugin --timeout=60s
fi

echo "--- Patching cnfs-nas-daemon ---"
kubectl -n cnfs-system patch ds cnfs-nas-daemon --type=strategic -p "
spec:
  updateStrategy:
    type: RollingUpdate
  template:
    spec:
      containers:
      - name: mount-proxy
        image: $IMAGE
        args: $ARGS
      imagePullSecrets:
      - name: $SECRET_NAME
"
kubectl -n cnfs-system rollout status ds/cnfs-nas-daemon --timeout=120s

echo ""
echo "Done. cnfs-nas-daemon is now running $IMAGE"
echo "Test with a PV using 'useClient: efcclient' in volumeAttributes."
