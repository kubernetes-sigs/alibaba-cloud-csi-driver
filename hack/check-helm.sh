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
    kubectl delete namespace ack-csi-customfuse
}
trap cleanup EXIT

# Check cluster connectivity
echo "Checking cluster connectivity..."
kubectl cluster-info > /dev/null

cd "$CHART_DIR"

# Install CRDs first (required for validation of CRD-based resources)
echo "Installing CRDs..."
kubectl create -f ./crds

# Create namespaces required for validation. The chart declares them, but
# --dry-run=server never creates them, so resources placed in them fail to
# validate unless they already exist.
echo "Creating namespace ack-csi-fuse..."
kubectl create namespace ack-csi-fuse
echo "Creating namespace ack-csi-customfuse..."
kubectl create namespace ack-csi-customfuse

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

# customfuse is off in every values file, so its resources reach the API server
# only through an explicit opt-in. Both switches are exercised: the node side
# alone, then with the controller, since they render different workloads.
for cf_args in \
    "--set csi.customfuse.enabled=true" \
    "--set csi.customfuse.enabled=true --set csi.customfuse.controller.enabled=true"
do
    echo "=== Validating with $cf_args ==="
    # shellcheck disable=SC2086
    helm lint . $cf_args
    # shellcheck disable=SC2086
    helm template alibaba-cloud-csi-driver . --namespace kube-system $cf_args | \
        tee /dev/stderr | \
        kubectl apply --dry-run=server --validate=strict -f -
done

echo "=== All validations passed ==="
