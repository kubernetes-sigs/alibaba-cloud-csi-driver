#!/bin/bash
# Exercises the customfuse CSI paths on a normal (DaemonSet) deployment, using the
# mock FUSE client so that no bucket, credential or client binary is involved.
#
# What it establishes, per case, by reading what the mock reported:
#   static      volumeAttributes reach the entrypoint as environment variables
#   dynamic     a PVC's requested size reaches it as $capacity, via CreateVolume
#   override    a capacity in spec.mountOptions wins over one in volumeAttributes
#   visibility  the mount is usable from the workload container
#   teardown    deleting the workload removes the fuse pod and the mount
#
# Usage:
#   MOCK_IMAGE=<registry>/customfuse-mock:<tag> bash test/customfuse/run-runc.sh
#
# Requires: kubectl against a cluster with the customfuse driver installed, and
# the fuse image registered as "fuse-<FUSE_TYPE>" in the csi-plugin ConfigMap.

set -uo pipefail

MOCK_IMAGE="${MOCK_IMAGE:-}"
FUSE_TYPE="${FUSE_TYPE:-customfuse-mock}"
FUSE_NS="${FUSE_NS:-ack-csi-customfuse}"
NS="${NS:-default}"
PREFIX="${PREFIX:-cfmock}"
DRIVER=customfuseplugin.csi.alibabacloud.com
TIMEOUT="${TIMEOUT:-180}"

failures=0
pass() { printf '    ok   %s\n' "$*"; }
fail() { printf '    FAIL %s\n' "$*" >&2; failures=$((failures + 1)); }
step() { printf '\n== %s ==\n' "$*"; }

# apply_or_fail feeds a manifest on stdin to kubectl apply and reports the outcome.
# The exit status has to be checked: a manifest the API server rejects — an immutable
# field on an object that already exists, above all — otherwise still looks like it
# was created, and every assertion in the case then fails for a reason that is not
# there. The rejection is echoed because it is the only thing that says why.
#
# Callers hand it the manifest through a redirect rather than a pipe, so that it runs
# in this shell. Piped, it would run in a subshell and the exit would end that
# subshell alone, leaving the script to carry on without the object it failed to
# create.
apply_or_fail() { # $1=what the manifest creates
    local out
    if out="$(kubectl apply -f - 2>&1)"; then
        pass "$1 created"
    else
        fail "could not create $1"
        printf '%s\n' "$out" | sed 's/^/      /' >&2
        exit 1
    fi
}

created=()
track() { created+=("$1"); }
cleanup() {
    step "cleanup"
    for ((i = ${#created[@]} - 1; i >= 0; i--)); do
        # shellcheck disable=SC2086
        kubectl delete ${created[$i]} --ignore-not-found --wait=false >/dev/null 2>&1
    done
    kubectl -n "$NS" wait --for=delete pod -l "csitest=$PREFIX" --timeout=90s >/dev/null 2>&1
    # Dynamically provisioned PVs carry a provisioner-assigned name (customfuse-<uuid>),
    # so they are located through the claim they were bound to. Retain is enforced by
    # the driver, which means nothing else ever collects them.
    pvs=$(kubectl get pv -o jsonpath='{range .items[*]}{.metadata.name}{" "}{.spec.claimRef.name}{"\n"}{end}' 2>/dev/null \
          | awk -v p="$PREFIX" '$1 ~ p || $2 ~ "^"p {print $1}' || true)
    # Detach has to land before the PV goes away: the attacher resolves the volume
    # handle from the PV, so deleting one under a live VolumeAttachment leaves the
    # attachment stranded on its finalizer and the driver never hears about it, which
    # orphans the fuse pod. Deleting the workload only starts detach, so wait for the
    # attachments to disappear rather than for the pods.
    attachments=$(kubectl get volumeattachment -o jsonpath='{range .items[*]}{.metadata.name}{" "}{.spec.attacher}{" "}{.spec.source.persistentVolumeName}{"\n"}{end}' 2>/dev/null \
          | awk -v p="$PREFIX" -v d="$DRIVER" '$2 == d && ($3 ~ p || $3 ~ "^customfuse-") {print "volumeattachment/"$1}' || true)
    # shellcheck disable=SC2086
    [ -n "$attachments" ] && kubectl wait --for=delete $attachments --timeout=90s >/dev/null 2>&1
    for pv in $pvs; do
        kubectl patch "pv/$pv" -p '{"metadata":{"finalizers":null}}' --type=merge >/dev/null 2>&1
        kubectl delete "pv/$pv" --ignore-not-found --wait=false >/dev/null 2>&1
    done
    printf '  cleaned\n'
}
trap cleanup EXIT

[ -n "$MOCK_IMAGE" ] || { echo "MOCK_IMAGE is required; build test/customfuse/Dockerfile" >&2; exit 1; }
kubectl cluster-info >/dev/null 2>&1 || { echo "no cluster" >&2; exit 1; }

step "preflight"
if kubectl get csidriver "$DRIVER" >/dev/null 2>&1; then
    pass "CSIDriver $DRIVER present"
else
    fail "CSIDriver $DRIVER missing — is customfuse installed?"
    exit 1
fi
# The driver resolves the fuse pod image from this key; without it the fuse pod
# cannot start and every case below would fail for the same uninformative reason.
if kubectl -n kube-system get cm csi-plugin -o jsonpath="{.data.fuse-$FUSE_TYPE}" 2>/dev/null | grep -q image=; then
    pass "csi-plugin ConfigMap has fuse-$FUSE_TYPE"
else
    fail "csi-plugin ConfigMap lacks fuse-$FUSE_TYPE=image=$MOCK_IMAGE"
    echo "    kubectl -n kube-system patch cm csi-plugin --type=merge -p '{\"data\":{\"fuse-$FUSE_TYPE\":\"image=$MOCK_IMAGE\"}}'" >&2
    exit 1
fi

# ── helpers ──────────────────────────────────────────────────────────────────

# volume_label mirrors ComputeVolumeIdLabelVal: a handle that is already a valid
# label value is used as-is, anything else is hashed. Reproducing it here keeps the
# lookups working for handles that are not valid label values, such as a URI.
volume_label() { # $1=volume handle
    if printf '%s' "$1" | grep -qE '^[A-Za-z0-9]([A-Za-z0-9._-]{0,61}[A-Za-z0-9])?$'; then
        printf '%s' "$1"
    else
        printf 'h1.%s' "$(printf '%s' "$1" | sha1sum | cut -d' ' -f1)"
    fi
}

fuse_pod_of() { # $1=volume handle
    kubectl -n "$FUSE_NS" get pod \
        -l "csi.alibabacloud.com/volume-id=$(volume_label "$1")" \
        -o jsonpath='{.items[0].metadata.name}' 2>/dev/null
}

# mock_env reads a value the mock reported. Reading it back from the log rather
# than from the manifest is the point: it proves the value survived the whole
# path through the CSI protocol and mount-proxy.
mock_env() { # $1=volume handle, $2=variable
    local pod
    pod="$(fuse_pod_of "$1")"
    [ -n "$pod" ] || return 1
    kubectl -n "$FUSE_NS" logs "$pod" 2>/dev/null | sed -n "s/^CSITEST_ENV $2=//p" | tail -1
}

expect_env() { # $1=volume handle, $2=variable, $3=expected
    local got
    got="$(mock_env "$1" "$2")"
    if [ "$got" = "$3" ]; then pass "\$$2 = $got"; else fail "\$$2 = '${got:-<none>}', want '$3'"; fi
}

wait_workload() { # $1=pod name
    kubectl -n "$NS" wait --for=condition=Ready "pod/$1" --timeout="${TIMEOUT}s" >/dev/null 2>&1
}

diagnose() { # $1=pod name, $2=volume handle
    echo "    --- diagnostics ---" >&2
    kubectl -n "$NS" describe "pod/$1" 2>/dev/null | sed -n '/Events:/,$p' | head -12 >&2
    local fp; fp="$(fuse_pod_of "$2")"
    [ -n "$fp" ] && kubectl -n "$FUSE_NS" logs "$fp" --tail=25 2>/dev/null | sed 's/^/    /' >&2
}

# static_pv emits a PV whose volumeAttributes carry the values each case asserts on.
static_pv() { # $1=name, $2=extra volumeAttributes yaml, $3=mountOptions yaml
    cat <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: $1
  labels: {csitest: $PREFIX}
spec:
  capacity: {storage: 10Gi}
  accessModes: [ReadWriteMany]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
${3}
  csi:
    driver: $DRIVER
    volumeHandle: $1
    fsType: $FUSE_TYPE
    volumeAttributes:
      source: "mock://$1"
      bucket: "mock-bucket"
      url: "mock.endpoint.invalid"
      path: "sub/dir"
${2}
EOF
}

consumer_pod() { # $1=name, $2=pvc name
    cat <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: $1
  labels: {csitest: $PREFIX}
spec:
  restartPolicy: Never
  containers:
    - name: app
      image: ${BUSYBOX_IMAGE:-registry.cn-shanghai.aliyuncs.com/eci_open/busybox:1.30}
      command: ["sh", "-c", "sleep 3600"]
      volumeMounts: [{name: data, mountPath: /data}]
  volumes:
    - name: data
      persistentVolumeClaim: {claimName: $2}
EOF
}

pvc() { # $1=name, $2=volumeName ("" for dynamic), $3=size, $4=storageClass
    cat <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: $1
  namespace: $NS
  labels: {csitest: $PREFIX}
spec:
  accessModes: [ReadWriteMany]
  resources: {requests: {storage: $3}}
  storageClassName: "$4"
$( [ -n "$2" ] && echo "  volumeName: $2" )
EOF
}

# ── case 1: static PV, volumeAttributes → env ────────────────────────────────

step "static PV: volumeAttributes reach the entrypoint"
V1="$PREFIX-static"
apply_or_fail "PV $V1" < <(static_pv "$V1" '      otherOpts: "--mock-flag=1"
      capacity: "7Gi"' "")
track "pv/$V1"
apply_or_fail "PVC $V1" < <(pvc "$V1" "$V1" 10Gi "")
track "pvc/$V1 -n $NS"
apply_or_fail "consumer pod $V1" < <(consumer_pod "$V1" "$V1")
track "pod/$V1 -n $NS"

if wait_workload "$V1"; then
    pass "workload Ready"
    expect_env "$V1" source "mock://$V1"
    expect_env "$V1" bucket "mock-bucket"
    expect_env "$V1" url "mock.endpoint.invalid"
    expect_env "$V1" path "sub/dir"
    expect_env "$V1" otherOpts "--mock-flag=1"
    expect_env "$V1" capacity "7Gi"
    # fuseType is a selector, not a mount option: it picks which client image runs
    # and is deliberately not forwarded to the entrypoint, so assert what it chose.
    img="$(kubectl -n "$FUSE_NS" get pod "$(fuse_pod_of "$V1")" -o jsonpath='{.spec.containers[0].image}' 2>/dev/null)"
    if [ "$img" = "$MOCK_IMAGE" ]; then
        pass "fuseType selected image $img"
    else
        fail "fuseType selected image '${img:-<none>}', want '$MOCK_IMAGE'"
    fi
    # mountpoint is assigned by the driver, so assert its shape rather than a literal
    mp="$(mock_env "$V1" mountpoint)"
    case "$mp" in
        /*globalmount) pass "\$mountpoint = $mp" ;;
        *) fail "\$mountpoint = '${mp:-<none>}', want a globalmount path" ;;
    esac
else
    fail "workload not Ready"
    diagnose "$V1" "$V1"
fi

step "static PV: the mount is usable from the workload"
if kubectl -n "$NS" exec "$V1" -- test -f /data/.csitest-marker >/dev/null 2>&1; then
    pass "marker file visible through the mount"
    if kubectl -n "$NS" exec "$V1" -- sh -c 'echo rw > /data/.rwtest && cat /data/.rwtest' 2>/dev/null | grep -q rw; then
        pass "mount is writable"
    else
        fail "mount is not writable"
    fi
else
    fail "marker file not visible — propagation from attachPath to targetPath broke"
    diagnose "$V1" "$V1"
fi

# ── case 2: mountOptions overrides volumeAttributes ──────────────────────────

step "mountOptions capacity overrides volumeAttributes capacity"
V2="$PREFIX-override"
apply_or_fail "PV $V2" < <(static_pv "$V2" '      capacity: "7Gi"' '  mountOptions: ["capacity=99Gi"]')
track "pv/$V2"
apply_or_fail "PVC $V2" < <(pvc "$V2" "$V2" 10Gi "")
track "pvc/$V2 -n $NS"
apply_or_fail "consumer pod $V2" < <(consumer_pod "$V2" "$V2")
track "pod/$V2 -n $NS"

if wait_workload "$V2"; then
    expect_env "$V2" capacity "99Gi"
else
    fail "workload not Ready"
    diagnose "$V2" "$V2"
fi

# ── case 3: dynamic provisioning, capacity_range → $capacity ─────────────────

step "dynamic provisioning: PVC size reaches \$capacity through CreateVolume"
SC="$PREFIX-sc"
apply_or_fail "StorageClass $SC" <<EOF
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: $SC
  labels: {csitest: $PREFIX}
provisioner: $DRIVER
reclaimPolicy: Retain
volumeBindingMode: Immediate
parameters:
  fuseType: "$FUSE_TYPE"
  source: "mock://dynamic"
  url: "mock.endpoint.invalid"
EOF
track "sc/$SC"
V3="$PREFIX-dynamic"
apply_or_fail "PVC $V3" < <(pvc "$V3" "" 42Gi "$SC")
track "pvc/$V3 -n $NS"
apply_or_fail "consumer pod $V3" < <(consumer_pod "$V3" "$V3")
track "pod/$V3 -n $NS"

if wait_workload "$V3"; then
    pass "workload Ready on a dynamically provisioned volume"
    HANDLE="$(kubectl -n "$NS" get pvc "$V3" -o jsonpath='{.spec.volumeName}' 2>/dev/null)"
    if [ -n "$HANDLE" ]; then
        pass "PV $HANDLE provisioned"
        expect_env "$HANDLE" capacity "42Gi"
        expect_env "$HANDLE" source "mock://dynamic"
    else
        fail "PVC did not bind to a PV"
    fi
else
    fail "workload not Ready — CreateVolume may have failed"
    kubectl -n "$NS" describe "pvc/$V3" 2>/dev/null | sed -n '/Events:/,$p' | head -12 >&2
    diagnose "$V3" "$(kubectl -n "$NS" get pvc "$V3" -o jsonpath='{.spec.volumeName}' 2>/dev/null)"
fi

# ── case 4: a Delete reclaim policy is refused ───────────────────────────────

step "reclaimPolicy=Delete is rejected"
SCD="$PREFIX-sc-delete"
apply_or_fail "StorageClass $SCD" <<EOF
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: $SCD
  labels: {csitest: $PREFIX}
provisioner: $DRIVER
reclaimPolicy: Delete
volumeBindingMode: Immediate
parameters: {fuseType: "$FUSE_TYPE", source: "mock://del", url: "mock.endpoint.invalid"}
EOF
track "sc/$SCD"
V4="$PREFIX-delete"
apply_or_fail "PVC $V4" < <(pvc "$V4" "" 1Gi "$SCD")
track "pvc/$V4 -n $NS"

# The provisioner retries, so wait for the refusal to show up as an event rather
# than for a terminal state the PVC never reaches.
found=0
for _ in $(seq 1 24); do
    if kubectl -n "$NS" describe "pvc/$V4" 2>/dev/null | grep -qi 'retain'; then found=1; break; fi
    sleep 5
done
if [ "$found" = 1 ]; then
    pass "CreateVolume refused a non-Retain reclaim policy"
else
    fail "no refusal event mentioning Retain within 120s"
    kubectl -n "$NS" describe "pvc/$V4" 2>/dev/null | sed -n '/Events:/,$p' | head -10 >&2
fi

# ── case 5: teardown removes the fuse pod ───────────────────────────────────

step "deleting the workload tears the mount down"
FP1="$(fuse_pod_of "$V1")"
kubectl -n "$NS" delete "pod/$V1" --wait=true --timeout=90s >/dev/null 2>&1
kubectl -n "$NS" delete "pvc/$V1" --wait=true --timeout=90s >/dev/null 2>&1
gone=0
for _ in $(seq 1 24); do
    [ -z "$(fuse_pod_of "$V1")" ] && { gone=1; break; }
    sleep 5
done
if [ "$gone" = 1 ]; then
    pass "fuse pod ${FP1:-?} removed"
else
    fail "fuse pod ${FP1:-?} still present 120s after the volume went away"
fi

# ── result ──────────────────────────────────────────────────────────────────

printf '\n'
if [ "$failures" -ne 0 ]; then
    echo "FAILED: $failures check(s)" >&2
    exit 1
fi
echo "PASS: customfuse behaves correctly on a DaemonSet deployment"
