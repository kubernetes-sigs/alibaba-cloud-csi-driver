#!/bin/bash
set -e

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && git rev-parse --show-toplevel)"
VALUES_FILE="$ROOT/.local/test-values.yaml"

CLUSTER_ID=$(yq '.deploy.clusterID' "$VALUES_FILE")

if [ -z "$CLUSTER_ID" ] || [ "$CLUSTER_ID" = "null" ]; then
    echo "Error: clusterID not found in $VALUES_FILE"
    exit 1
fi

echo "Cluster: $CLUSTER_ID"

# If CSI is already deployed via helm, it's likely a debug version — skip
if kubectl get deploy -n kube-system csi-provisioner -o jsonpath='{.metadata.labels.app\.kubernetes\.io/managed-by}' 2>/dev/null | grep -q Helm; then
    echo "CSI is deployed via Helm (debug version), nothing to uninstall"
    exit 0
fi

uninstall_addon() {
    local addon=$1
    echo "Uninstalling $addon..."
    aliyun cs POST "/clusters/${CLUSTER_ID}/components/uninstall" \
        --header "Content-Type=application/json" \
        --body "[{\"name\":\"${addon}\"}]"

    echo "Waiting for $addon to be removed..."
    while aliyun cs GET "/clusters/${CLUSTER_ID}/addon_instances" | \
        jq -e ".addons[] | select(.name == \"${addon}\")" > /dev/null 2>&1; do
        sleep 5
    done
    echo "$addon removed"
}

# Check which CSI addons are installed
INSTALLED=$(aliyun cs GET "/clusters/${CLUSTER_ID}/addon_instances" | \
    jq -r '.addons[] | select(.name == "csi-plugin" or .name == "csi-provisioner" or .name == "managed-csiprovisioner") | .name')

if [ -z "$INSTALLED" ]; then
    echo "No CSI addons installed"
else
    echo "Found installed CSI addons: $INSTALLED"

    # Uninstall csi-plugin first (it depends on csi-provisioner/managed-csiprovisioner)
    if echo "$INSTALLED" | grep -q "^csi-plugin$"; then
        uninstall_addon csi-plugin
    fi

    # Uninstall csi-provisioner or managed-csiprovisioner (mutually exclusive)
    for addon in csi-provisioner managed-csiprovisioner; do
        if echo "$INSTALLED" | grep -q "^${addon}$"; then
            uninstall_addon "$addon"
        fi
    done
fi

# Clean up leftover resources that addon uninstall doesn't remove
if kubectl get ns ack-csi-fuse &> /dev/null; then
    if [ -z "$(kubectl get pods -n ack-csi-fuse -o name 2>/dev/null)" ]; then
        echo "Deleting leftover ack-csi-fuse namespace..."
        kubectl delete ns ack-csi-fuse
    else
        echo "Warning: ack-csi-fuse namespace has running pods, skipping deletion"
    fi
fi

# Remove leftover SA in kube-system that conflicts with helm install
kubectl delete sa csi-fuse-ossfs -n kube-system 2>/dev/null && echo "Deleted leftover SA csi-fuse-ossfs in kube-system" || true

echo "All CSI addons uninstalled"
