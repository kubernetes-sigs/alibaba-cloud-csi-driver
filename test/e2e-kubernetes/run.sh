#!/bin/bash

ROOT=$(realpath "$(dirname "${BASH_SOURCE[0]}")/../..")
OUTPUT=${OUTPUT:-/workspace/output/reports}
KUBECONFIG=${KUBECONFIG:-$HOME/.kube/config}
E2E_BIN=${E2E_BIN:-/usr/local/bin/e2e.test}

E2E_ARGS=(
    -provider skeleton
    -kubeconfig "$KUBECONFIG"
    -repo-root "$ROOT"
)
GINKGO_ARGS=(
    --no-color
    --output-dir "$OUTPUT"
)
EXIT=0
set -x

# remove skipping volumeLimits after https://github.com/kubernetes/kubernetes/pull/126978 merged
ginkgo "${GINKGO_ARGS[@]}" \
    --skip='\[Disruptive\]|\[Serial\]' \
    --skip='eed.csi.+\(immediate binding\)' \
    --skip='eed.csi.+volumeLimits' \
    --focus='External.Storage' \
    --junit-report=parallel-junit.xml \
    --procs=20 "$E2E_BIN" -- "${E2E_ARGS[@]}" \
        -storage.testdriver="$ROOT/test/e2e-kubernetes/disk-manifest.yaml" \
        -storage.testdriver="$ROOT/test/e2e-kubernetes/eed-manifest.yaml" || EXIT=$?

ginkgo "${GINKGO_ARGS[@]}" \
    --focus='External.Storage.+(\[Disruptive\]|\[Serial\])' \
    --junit-report=serial-junit.xml \
    "$E2E_BIN" -- "${E2E_ARGS[@]}" \
        -storage.testdriver="$ROOT/test/e2e-kubernetes/disk-manifest.yaml" || EXIT=$?

exit $EXIT
