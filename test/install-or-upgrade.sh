#!/bin/bash
# If helm chart is changed, try install upstream version, then upgrade to local version,
# to ensure all changes are valid. e.g. we don't change some immutable fields.

set -e

UPSTREAM=${UPSTREAM:-"origin/master"}

if ! git diff --exit-code --stat "$UPSTREAM" HEAD -- deploy/chart; then
    echo "helm chart is changed, installing $UPSTREAM version first"
    git worktree add --detach --no-checkout ../upstream "$UPSTREAM"
    ( cd ../upstream && git checkout HEAD -- deploy/chart )

    helm install --namespace kube-system alibaba-cloud-csi-driver ../upstream/deploy/chart "$@"
    git worktree remove --force ../upstream
fi

echo installing local version
helm upgrade --install --namespace kube-system alibaba-cloud-csi-driver ./deploy/chart "$@"
