#!/bin/bash
# Validates the Helm chart against a real API server using strict server-side
# validation. The cluster does not have to be clean: CRDs and namespaces that
# already exist are left alone, and only what this run creates is removed.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
CHART_DIR="$PROJECT_ROOT/deploy/charts/alibaba-cloud-csi-driver"

# Filled in below with what this run actually creates, so cleanup can tell that
# apart from what was already on the cluster.
created_crds=()
created_namespaces=()
work_dir=""

cleanup() {
    # Delete only what this run created. The trap also fires when a validation
    # step fails, and on a shared cluster these objects belong to a live install:
    # ack-csi-fuse is even marked helm.sh/resource-policy: keep by the chart,
    # because it hosts on-demand fuse pods that may be serving live mounts.
    local item
    for item in "${created_namespaces[@]}"; do
        echo "Deleting namespace $item"
        kubectl delete namespace "$item" --ignore-not-found
    done
    for item in "${created_crds[@]}"; do
        echo "Deleting CRDs from $item"
        kubectl delete -f "$item" --ignore-not-found
    done
    if [ -n "$work_dir" ]; then
        rm -rf "$work_dir"
    fi
}
trap cleanup EXIT

echo "Checking cluster connectivity..."
kubectl cluster-info > /dev/null

cd "$CHART_DIR"

# Every variant the chart is validated with: the default values, each values
# file, then the customfuse switches. customfuse is off in every values file, so
# its resources reach the API server only through an explicit opt-in, and both
# switches are exercised because they render different workloads.
variant_args=("")
for values_file in values-*.yaml; do
    variant_args+=("--values $values_file")
done
variant_args+=(
    "--set csi.customfuse.enabled=true"
    "--set csi.customfuse.enabled=true --set csi.customfuse.controller.enabled=true"
)

# Rendered once up front, because the namespaces to create are read out of the
# same manifests that get validated.
work_dir="$(mktemp -d)"
echo "Rendering ${#variant_args[@]} chart variants..."
for i in "${!variant_args[@]}"; do
    args="${variant_args[$i]}"
    echo "${args:-default values}" > "$work_dir/$i.label"
    # shellcheck disable=SC2086
    helm template alibaba-cloud-csi-driver . --namespace kube-system $args \
        > "$work_dir/$i.yaml"
done

# Install CRDs first (required for validation of CRD-based resources). Each file
# is checked on its own so that CRDs already on the cluster are left untouched
# rather than aborting the run with AlreadyExists.
echo "Installing CRDs..."
for crd_file in "$CHART_DIR"/crds/*.yaml; do
    if kubectl get -f "$crd_file" > /dev/null 2>&1; then
        echo "  $(basename "$crd_file") already installed, leaving it alone"
        continue
    fi
    kubectl create -f "$crd_file"
    created_crds+=("$crd_file")
done

# Create the namespaces those renders declare. --dry-run=server never creates
# them, so an object placed in one fails to validate unless it already exists.
echo "Preparing namespaces..."
while read -r ns; do
    if kubectl get namespace "$ns" > /dev/null 2>&1; then
        echo "  $ns already exists, leaving it alone"
        continue
    fi
    kubectl create namespace "$ns" > /dev/null
    created_namespaces+=("$ns")
done < <(awk '
    /^kind: Namespace$/ { want = 1; next }
    want && /^  name: / { print $2; want = 0 }
' "$work_dir"/*.yaml | sort -u)

for i in "${!variant_args[@]}"; do
    args="${variant_args[$i]}"
    echo "=== Validating with $(cat "$work_dir/$i.label") ==="
    # shellcheck disable=SC2086
    helm lint . $args
    # Helm's --dry-run=server doesn't actually validate against the API server,
    # so kubectl does it with strict validation. The template goes to stderr for
    # manual review while it is validated.
    tee /dev/stderr < "$work_dir/$i.yaml" | \
        kubectl apply --dry-run=server --validate=strict -f -
done

echo "=== All validations passed ==="
