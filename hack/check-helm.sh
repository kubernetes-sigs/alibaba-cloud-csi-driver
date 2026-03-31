#!/bin/bash
# We need a clean Kubernetes cluster to run this script.
# This script validates the Helm chart using strict server-side validation.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
CHART_DIR="$PROJECT_ROOT/deploy/charts/alibaba-cloud-csi-driver"

# Cleanup function
cleanup() {
    echo "Cleaning up CRDs and namespaces..."
    kubectl delete -f "$CHART_DIR/crds"
    kubectl delete namespace ack-csi-fuse
}
trap cleanup EXIT

# Check cluster connectivity
echo "Checking cluster connectivity..."
kubectl cluster-info > /dev/null

cd "$CHART_DIR"

# Install CRDs first (required for validation of CRD-based resources)
echo "Installing CRDs..."
kubectl create -f ./crds

# Create namespace required for validation (ack-csi-fuse is defined in the chart)
echo "Creating namespace ack-csi-fuse..."
kubectl create namespace ack-csi-fuse

# Validate chart with default values
echo "=== Validating with default values ==="
helm lint .

# Helm's --dry-run=server doesn't actually validate against API server
# Use kubectl with strict validation for proper server-side validation
# Output template to stderr for manual review while validating
helm template alibaba-cloud-csi-driver . --namespace kube-system | \
    tee /dev/stderr | \
    kubectl apply --dry-run=server --validate=strict -f -

# Validate with each values file
for v in values-*.yaml; do
    echo "=== Validating with $v ==="
    helm lint . --values "$v"
    helm template alibaba-cloud-csi-driver . --namespace kube-system --values "$v" | \
        tee /dev/stderr | \
        kubectl apply --dry-run=server --validate=strict -f -
done

echo "=== All validations passed ==="
