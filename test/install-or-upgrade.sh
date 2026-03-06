#!/bin/bash
# If helm chart is changed, try install upstream version, then upgrade to local version,
# to ensure all changes are valid. e.g. we don't change some immutable fields.

set -e

UPSTREAM=${UPSTREAM:-"origin/master"}
CHART=deploy/charts/alibaba-cloud-csi-driver

install() {
    base=$1
    shift

    echo installing CRDs  # helm does not upgrade CRDs if already exists
    kubectl apply -f "$base/$CHART/crds"

    echo installing helm chart
    helm upgrade --install --namespace kube-system alibaba-cloud-csi-driver "$base/$CHART" "$@"
}

if ! git diff --exit-code --stat "$UPSTREAM" HEAD -- $CHART; then
    echo "helm chart is changed, installing $UPSTREAM version first"
    git worktree add --detach --no-checkout ../upstream "$UPSTREAM"
    ( cd ../upstream && git checkout HEAD -- $CHART )  # only checkout chart for install
    install ../upstream "$@"
    git worktree remove --force ../upstream
fi

echo installing local version
install . "$@"
