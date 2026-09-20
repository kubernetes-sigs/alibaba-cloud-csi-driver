#!/bin/bash
# Exercises the customfuse CSI paths on a sandbox deployment, where the driver is
# injected into the sandbox pod rather than running as a DaemonSet, using the mock
# FUSE client so that no bucket, credential or client binary is involved.
#
# What this covers that run-runc.sh cannot:
#   * the injected csi-customfuse-agent-sidecar serves the mount
#   * there is no ControllerPublish, so the socket has to come from the
#     --customfuse-mount-proxy-sock flag rather than from PublishContext, whose
#     value is supplied by a caller that does not distinguish drivers
#   * capacity can only come from the PV, since no controller fills it in
#   * a volume arrives through SandboxClaim.dynamicVolumesMount, not a PVC
#   * under agent-identity the credential is exchanged per mount, scoped by the
#     claim's subPath, delivered as files, and rotated while the mount lives on
#
# Which runtime injects the customfuse sidecars, and how, is the sandbox
# platform's business rather than this repository's. RUNTIME names the one to
# use; nothing here configures it. What the script asserts on is the pod that
# came back — the containers present, their arguments, their mounts — so a
# platform that wires customfuse differently still gets a meaningful answer
# instead of a check against a config schema it never promised.
#
# Passing CRED_PROVIDER=<credential provider> switches the volume to
# agent-identity, where the credential is exchanged per mount instead of read from
# a Secret. It needs OSS_BUCKET and OSS_URL, since the point of the extra
# assertions is that the exchanged credential works. The claim then also names an
# AgentIdentity — AGENT_IDENTITY, or the first Available one in NS — because the
# credential is issued per sandbox and a sandbox claimed without an agent name is
# never issued one.
#
# SUB_PATH=<path> asks for a subdirectory of the volume instead of its root. It is
# relative to the volume's own path rather than to the bucket, and under
# agent-identity it is also what scopes the credential, so the two are asserted to
# agree: a credential scoped elsewhere still renders a policy, still mounts, and
# only fails once objects are read, which presents as a storage problem and is not
# one. Leaving it empty mounts at the volume's root, the one case where the two
# coincide.
#
# ROTATION_WAIT=<seconds> additionally waits for the delivered credential to be
# replaced and checks that the replacement still works. It is off by default
# because observing one needs both ends asking often enough — the mount-proxy's
# refresh margin raised and the credential's lifetime lowered — and neither is a
# property of this repository. See the rotation step for the detail.
#
# Usage:
#   MOCK_IMAGE=<registry>/customfuse-mock:<tag> \
#   SANDBOXSET=<pool declaring the runtime> \
#   bash test/customfuse/run-sandbox.sh
#
# The sandbox platform is assumed installed and configured: a runtime that
# injects the customfuse sidecars, and an image pull path that reaches the mock.
# This script changes nothing cluster-wide — it only creates and deletes the
# resources it names, all labelled csitest=<PREFIX>.

set -uo pipefail

MOCK_IMAGE="${MOCK_IMAGE:-}"
SANDBOXSET="${SANDBOXSET:-}"
NS="${NS:-default}"
PREFIX="${PREFIX:-cfsbx}"
DRIVER=customfuseplugin.csi.alibabacloud.com
# The runtime the platform injects customfuse through. Matches the name the
# examples declare; override it on a cluster whose platform wires customfuse
# into a runtime shared with the other drivers.
RUNTIME="${RUNTIME:-csi-customfuse}"
SOCKET="${SOCKET:-/run/cnfs/customfuse-mounter.sock}"
NO_CLEANUP="${NO_CLEANUP:-0}"
TIMEOUT="${TIMEOUT:-240}"
# Only needed when this script creates the pool; a pool that already exists brings
# its own image. Left without a default because which image a sandbox runs is a
# property of the cluster, not of this repository.
SANDBOX_IMAGE="${SANDBOX_IMAGE:-}"
# Setting CRED_PROVIDER switches the volume to agent-identity: the credential is
# exchanged per mount instead of read from a Secret. It needs a real bucket, since
# the point of the extra assertions is that the exchanged credential works.
CRED_PROVIDER="${CRED_PROVIDER:-}"
OSS_BUCKET="${OSS_BUCKET:-}"
OSS_URL="${OSS_URL:-}"
OSS_PATH="${OSS_PATH:-/}"
# The claim's subPath, relative to the volume's own path rather than to the bucket.
# Empty mounts the volume at its root, which is the one case where the two coincide.
SUB_PATH="${SUB_PATH:-}"
# How long to watch for the delivered credential to be replaced, in seconds. Zero
# skips the wait. Rotation is only observable when both ends are asked for it often
# enough; see the rotation step for the two knobs that govern it.
ROTATION_WAIT="${ROTATION_WAIT:-0}"
# The credential is issued per sandbox, and a sandbox only asks for one when it is
# claimed with an agent name. Without it no token file is ever written, the sandbox
# layer has no sandboxId to inject, and the mount is refused. Auto-detected when empty.
AGENT_IDENTITY="${AGENT_IDENTITY:-}"

failures=0
pass() { printf '    ok   %s\n' "$*"; }
fail() { printf '    FAIL %s\n' "$*" >&2; failures=$((failures + 1)); }
skip() { printf '    skip %s\n' "$*"; }
step() { printf '\n== %s ==\n' "$*"; }

# apply_or_fail feeds a manifest on stdin to kubectl apply and reports the outcome.
# The exit status has to be checked: a manifest the API server rejects — an
# immutable field on an object that already exists, above all — otherwise still
# prints as created, and every later step then fails for a reason that is not there.
#
# Callers hand it the manifest as a heredoc redirect, not through a pipe, so that it
# runs in this shell. Piped, it runs in a subshell and the exit ends that subshell
# alone, leaving the script to carry on without the object it failed to create.
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
    # NO_CLEANUP leaves the claim, the sandbox and its pod standing so that a
    # failed mount can be read out of the container logs; the trap would
    # otherwise delete the only evidence of it.
    if [ "$NO_CLEANUP" = "1" ]; then
        step "cleanup skipped (NO_CLEANUP=1)"
        printf '  resources kept for inspection; delete them by hand when done\n'
        return
    fi
    step "cleanup"
    for ((i = ${#created[@]} - 1; i >= 0; i--)); do
        # shellcheck disable=SC2086
        kubectl delete ${created[$i]} --ignore-not-found --wait=false >/dev/null 2>&1
    done
    # A claimed sandbox has its ownerReferences cleared so the pool scales up, which
    # also means deleting the SandboxSet leaves it behind. Failed attempts can leave
    # several, so they are removed by pool label rather than by owner. A pool that was
    # borrowed rather than created names its sandboxes after itself and not after us,
    # and by now the claim is gone, so the one this run took is named directly.
    for sbx in $( { [ -n "${POD:-}" ] && printf 'sandbox/%s\n' "$POD"
                    kubectl -n "$NS" get sandbox -o name 2>/dev/null | grep "$PREFIX"
                  } | sort -u ); do
        kubectl -n "$NS" delete "$sbx" --ignore-not-found --wait=false >/dev/null 2>&1
    done
    for pv in $(kubectl get pv -o name 2>/dev/null | grep "$PREFIX" || true); do
        kubectl patch "$pv" -p '{"metadata":{"finalizers":null}}' --type=merge >/dev/null 2>&1
        kubectl delete "$pv" --ignore-not-found --wait=false >/dev/null 2>&1
    done
    printf '  cleaned\n'
}
trap cleanup EXIT

[ -n "$MOCK_IMAGE" ] || { echo "MOCK_IMAGE is required" >&2; exit 1; }
kubectl cluster-info >/dev/null 2>&1 || { echo "no cluster" >&2; exit 1; }

step "preflight"
for crd in sandboxsets.agents.kruise.io sandboxclaims.agents.kruise.io; do
    if kubectl get crd "$crd" >/dev/null 2>&1; then pass "$crd present"; else
        fail "$crd missing — this cluster has no sandbox support"; exit 1; fi
done
kubectl get csidriver "$DRIVER" >/dev/null 2>&1 && pass "CSIDriver present" || { fail "CSIDriver $DRIVER missing"; exit 1; }

# Everything the sandbox platform itself needs — the components that turn a claim
# into a pod, their allow-lists, their images — is assumed correct: none of it is
# part of this repository, and a failure there is reported by the claim's status
# message below rather than diagnosed here.

# The mock is injected as a sidecar of the sandbox pod. Injection adds containers
# and volumes to that pod, not imagePullSecrets, so a mock in a private registry has
# to be reachable through the pool's ServiceAccount — otherwise the sidecar sits in
# ErrImagePull with nothing in the CSI logs to explain it.
MOCK_HOST="${MOCK_IMAGE%%/*}"
case "$MOCK_HOST" in
    *.*) ;;                 # a host only counts as one if it looks like a domain
    *)   MOCK_HOST="" ;;    # docker.io shorthand, which needs no credential here
esac
if [ -n "$MOCK_HOST" ]; then
    POOL_SA=default
    if [ -n "$SANDBOXSET" ]; then
        POOL_SA="$(kubectl -n "$NS" get "sandboxset/$SANDBOXSET" -o jsonpath='{.spec.template.spec.serviceAccountName}' 2>/dev/null)"
        POOL_SA="${POOL_SA:-default}"
    fi
    PULL_SECRETS="$(kubectl -n "$NS" get sa "$POOL_SA" -o jsonpath='{range .imagePullSecrets[*]}{.name}{"\n"}{end}' 2>/dev/null)"
    covered=0
    for s in $PULL_SECRETS; do
        if kubectl -n "$NS" get secret "$s" -o jsonpath='{.data.\.dockerconfigjson}' 2>/dev/null \
            | base64 -d 2>/dev/null | grep -q "$MOCK_HOST"; then
            covered=1
            pass "$MOCK_HOST covered by secret $s (via sa/$POOL_SA)"
            break
        fi
    done
    if [ "$covered" = 0 ]; then
        skip "no pull credential for $MOCK_HOST reachable from sa/$POOL_SA in $NS"
        echo "    if that registry needs one, the injected sidecar will not start:" >&2
        echo "      kubectl -n $NS patch sa $POOL_SA -p '{\"imagePullSecrets\":[{\"name\":\"<secret>\"}]}'" >&2
    fi
fi

# ── the pool ─────────────────────────────────────────────────────────────────

# A sandbox is built with whatever the platform's configuration said at the time,
# so one already sitting in a pool has no customfuse sidecar and claiming it would
# not add one. Reusing a given pool is allowed, but then its sandboxes have to
# postdate the platform being configured for customfuse.
step "sandbox pool"
if [ -n "$SANDBOXSET" ]; then
    if kubectl -n "$NS" get "sandboxset/$SANDBOXSET" -o jsonpath='{.spec.runtimes}' 2>/dev/null | grep -q "\"$RUNTIME\""; then
        pass "reusing SandboxSet $SANDBOXSET (its sandboxes must postdate the platform's customfuse configuration)"
    else
        fail "SandboxSet $SANDBOXSET does not declare runtime '$RUNTIME'"
        exit 1
    fi
else
    [ -n "$SANDBOX_IMAGE" ] || {
        fail "SANDBOX_IMAGE is required to create a pool; alternatively pass SANDBOXSET=<existing pool>"
        exit 1
    }
    SANDBOXSET="$PREFIX-pool"
    apply_or_fail "SandboxSet $SANDBOXSET" <<EOF
apiVersion: agents.kruise.io/v1alpha1
kind: SandboxSet
metadata:
  name: $SANDBOXSET
  namespace: $NS
  labels: {csitest: $PREFIX}
spec:
  replicas: 1
  runtimes:
    - name: $RUNTIME
    - name: agent-runtime
  template:
    metadata:
      labels:
        alibabacloud.com/acs: "true"
        alibabacloud.com/compute-class: agent-sandbox
        alibabacloud.com/compute-qos: default
    spec:
      automountServiceAccountToken: false
      terminationGracePeriodSeconds: 30
      containers:
        - name: sandbox
          image: $SANDBOX_IMAGE
          imagePullPolicy: IfNotPresent
          resources:
            limits: {cpu: "1", memory: 1Gi}
            requests: {cpu: "1", memory: 1Gi}
EOF
    track "sandboxset/$SANDBOXSET -n $NS"
    printf '    note: this pool is built with the platform configuration as it stands now\n'
    printf '    waiting for the pool to fill (up to %ss)\n' "$TIMEOUT"
    for _ in $(seq 1 $((TIMEOUT / 5))); do
        READY="$(kubectl -n "$NS" get "sandboxset/$SANDBOXSET" -o jsonpath='{.status.availableReplicas}' 2>/dev/null)"
        [ "${READY:-0}" -ge 1 ] 2>/dev/null && break
        sleep 5
    done
    if [ "${READY:-0}" -ge 1 ] 2>/dev/null; then
        pass "pool has $READY available sandbox(es)"
    else
        fail "pool did not fill within ${TIMEOUT}s"
        kubectl -n "$NS" get "sandboxset/$SANDBOXSET" -o yaml 2>/dev/null | sed -n '/^status:/,$p' | head -10 >&2
        # Which container is unhappy, and why: a sandbox carries five of them once
        # customfuse is injected, and "not ready" alone does not say which.
        for p in $(kubectl -n "$NS" get pod -l "agents.kruise.io/sandbox-pool=$SANDBOXSET" -o name 2>/dev/null); do
            printf '    --- %s ---\n' "$p" >&2
            kubectl -n "$NS" get "$p" \
                -o jsonpath='{range .status.initContainerStatuses[*]}      {.name}{"  ready="}{.ready}{"  restarts="}{.restartCount}{"  "}{.state}{"\n"}{end}{range .status.containerStatuses[*]}      {.name}{"  ready="}{.ready}{"  restarts="}{.restartCount}{"  "}{.state}{"\n"}{end}' \
                2>/dev/null >&2
            kubectl -n "$NS" describe "$p" 2>/dev/null | sed -n '/^Events:/,$p' | tail -6 >&2
            for c in csi-sidecar csi-customfuse-agent-sidecar init; do
                LOG="$(kubectl -n "$NS" logs "$p" -c "$c" --tail=12 2>/dev/null)"
                [ -n "$LOG" ] && { printf '      == %s ==\n' "$c" >&2; printf '%s\n' "$LOG" | sed 's/^/        /' >&2; }
            done
        done
        exit 1
    fi
fi

# ── the volume, carrying no credential ───────────────────────────────────────

step "volume"
V="$PREFIX-vol"
if [ -n "$CRED_PROVIDER" ]; then
    [ -n "$OSS_BUCKET" ] && [ -n "$OSS_URL" ] || { fail "CRED_PROVIDER needs OSS_BUCKET and OSS_URL"; exit 1; }
    if [ -z "$AGENT_IDENTITY" ]; then
        AGENT_IDENTITY="$(kubectl -n "$NS" get agentidentities -o jsonpath='{.items[0].metadata.name}' 2>/dev/null)"
    fi
    [ -n "$AGENT_IDENTITY" ] || { fail "CRED_PROVIDER needs an AgentIdentity in $NS; set AGENT_IDENTITY=<name>"; exit 1; }
    if ! kubectl -n "$NS" get "agentidentity/$AGENT_IDENTITY" \
            -o jsonpath='{.status.conditions[?(@.type=="Available")].status}' 2>/dev/null | grep -q True; then
        fail "agentidentity/$AGENT_IDENTITY in $NS is not Available"; exit 1
    fi
    # No nodePublishSecretRef: with agent-identity there is no static credential to
    # point at. Nor does the PV name the provider or the sandbox: the sandbox side
    # supplies both, the claim below carrying the provider name. Writing the provider
    # here as well would let the mount succeed on the PV's copy and leave the path the
    # examples document — authType alone on the PV — the one thing never exercised.
    AUTH_ATTRS="      authType: \"agent-identity\"
      source: \"$OSS_BUCKET\"
      bucket: \"$OSS_BUCKET\"
      url: \"$OSS_URL\"
      path: \"$OSS_PATH\""
    pass "agent-identity mode: provider $CRED_PROVIDER, agent $AGENT_IDENTITY, bucket $OSS_BUCKET"
else
    AUTH_ATTRS="      source: \"mock://sandbox\"
      bucket: \"mock-bucket\"
      url: \"mock.endpoint.invalid\""
fi
apply_or_fail "PV $V" <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: $V
  labels: {csitest: $PREFIX}
spec:
  capacity: {storage: 10Gi}
  accessModes: [ReadWriteMany]
  persistentVolumeReclaimPolicy: Retain
  storageClassName: ""
  csi:
    driver: $DRIVER
    volumeHandle: $V
    fsType: customfuse-mock
    volumeAttributes:
      # fuseType is spelled out even though fsType above says the same thing: the
      # sandbox side validates volumeAttributes before the request reaches this
      # driver, and that check reads fuseType only. Both are set here so the
      # driver's "if both are given they must agree" path is exercised too.
      fuseType: customfuse-mock
      capacity: "33Gi"
$AUTH_ATTRS
EOF
track "pv/$V"
# spec.mountOptions is deliberately absent. It is read by the CSI flow around
# kubelet, which a claim does not go through: the volume reaches NodePublishVolume
# as the volumeAttributes on this object and nothing else. Capacity can therefore
# only come from them, and run-runc.sh is where mountOptions precedence is asserted.
pass "PV $V carries capacity in volumeAttributes only"

# ── sandbox ──────────────────────────────────────────────────────────────────

step "sandbox"
SC_="$PREFIX-claim"
apply_or_fail "SandboxClaim $SC_" <<EOF
apiVersion: agents.kruise.io/v1alpha1
kind: SandboxClaim
metadata:
  name: $SC_
  namespace: $NS
  labels: {csitest: $PREFIX}
spec:
  templateName: $SANDBOXSET
  replicas: 1
$([ -n "$AGENT_IDENTITY" ] && cat <<ANNOT
  # Carried onto the claimed sandbox, where it is what opts the sandbox into token
  # issuance. The annotation travels with the claim rather than the pool so that one
  # pool can serve both this and the Secret-passthrough case.
  annotations:
    security.agents.kruise.io/agent-name: "$AGENT_IDENTITY"
ANNOT
)
  dynamicVolumesMount:
    - pvName: $V
      mountPath: /mnt/data
      readOnly: false
$([ -n "$SUB_PATH" ] && printf '      subPath: %s\n' "$SUB_PATH")
$([ -n "$CRED_PROVIDER" ] && cat <<ATTR
      attributes:
        credentialProviderName: "$CRED_PROVIDER"
ATTR
)
  claimTimeout: "${TIMEOUT}s"
EOF
track "sandboxclaim/$SC_ -n $NS"

printf '    waiting for a sandbox pod (up to %ss)\n' "$TIMEOUT"
# The claim's status carries only counters, so the sandbox it took is found through
# the label the claim stamps on it. A Sandbox and its pod share a name.
POD=""
for _ in $(seq 1 $((TIMEOUT / 5))); do
    SBX="$(kubectl -n "$NS" get sandbox -l "agents.kruise.io/claim-name=$SC_" \
        -o jsonpath='{.items[0].metadata.name}' 2>/dev/null)"
    if [ -n "$SBX" ]; then
        POD="$(kubectl -n "$NS" get pod "$SBX" -o jsonpath='{.metadata.name}' 2>/dev/null)"
        [ -n "$POD" ] && break
    fi
    sleep 5
done

if [ -z "$POD" ]; then
    fail "no sandbox pod appeared within ${TIMEOUT}s"
    # The claim's status message is the one place the platform states why it did not
    # bind, so it goes first: what follows is context rather than the answer.
    MSG="$(kubectl -n "$NS" get "sandboxclaim/$SC_" -o jsonpath='{.status.message}' 2>/dev/null)"
    printf '    claim status.message: %s\n' "${MSG:-<empty>}" >&2
    kubectl -n "$NS" describe "sandboxclaim/$SC_" 2>/dev/null | sed -n '/Events:/,$p' | head -12 >&2
    exit 1
fi
pass "sandbox pod $POD"

step "injection took effect"
CONTAINERS="$(kubectl -n "$NS" get "pod/$POD" -o jsonpath='{range .spec.initContainers[*]}{.name}{"\n"}{end}{range .spec.containers[*]}{.name}{"\n"}{end}' 2>/dev/null)"
# Only what customfuse itself needs is asserted. Whatever else the runtime injects
# for its other drivers is the platform's business, and this script neither
# configures it nor has a basis for expecting it.
for c in csi-sidecar csi-customfuse-agent-sidecar; do
    if printf '%s' "$CONTAINERS" | grep -qx "$c"; then pass "$c injected"; else fail "$c not injected"; fi
done

# The flag is what makes the mount reach this driver's mount-proxy instead of the
# socket named in PublishContext, which a driver-agnostic caller may have filled
# with the neighbouring mounter's path.
SIDECAR_ARGS="$(kubectl -n "$NS" get "pod/$POD" -o jsonpath='{range .spec.initContainers[*]}{.args}{end}' 2>/dev/null)"
if printf '%s' "$SIDECAR_ARGS" | grep -q -- "customfuse-mount-proxy-sock=$SOCKET"; then
    pass "csi-sidecar carries --customfuse-mount-proxy-sock=$SOCKET"
else
    fail "csi-sidecar lacks --customfuse-mount-proxy-sock; the mount would follow PublishContext instead"
fi
if printf '%s' "$SIDECAR_ARGS" | grep -q -- '--driver=[^]"]*customfuse'; then
    pass "csi-sidecar serves customfuse"
else
    fail "csi-sidecar does not list customfuse in --driver"
fi

if [ -n "$CRED_PROVIDER" ]; then
    step "credential scope"
    # The credential this mount is handed is scoped by an annotation the platform
    # builds from the claim's volumes, so it is the one artifact showing what the
    # provider was actually asked for. It has to describe the same prefix the client
    # is pointed at: a scope that disagrees still renders a policy and still issues
    # a credential, the mount still succeeds, and only then does every object
    # operation fail — which reads as a storage problem and is not one.
    #
    # The sub-path is relative to the volume's own path, not to the bucket, so the
    # expected value is the two joined. They coincide only when the volume is rooted
    # at the bucket.
    scope_base="$(printf '%s' "$OSS_PATH" | sed 's|^/*||; s|/*$||')"
    scope_sub="$(printf '%s' "$SUB_PATH" | sed 's|^/*||; s|/*$||')"
    if [ -n "$scope_base" ] && [ -n "$scope_sub" ]; then WANT_SCOPE="$scope_base/$scope_sub"
    else WANT_SCOPE="$scope_sub"; fi

    SCOPE="$(kubectl -n "$NS" get "sandbox/$POD" -o jsonpath='{.metadata.annotations}' 2>/dev/null \
        | python3 -c 'import json,sys; print(json.load(sys.stdin).get("security.agents.kruise.io/storage-auth",""))' 2>/dev/null)"
    if [ -z "$SCOPE" ]; then
        fail "sandbox/$POD carries no storage-auth annotation; nothing scoped the credential"
    else
        GOT_PROVIDER="$(printf '%s' "$SCOPE" | python3 -c 'import json,sys; i=json.load(sys.stdin)[0]; print(i.get("credentialProviderName",""))' 2>/dev/null)"
        GOT_BUCKET="$(printf '%s' "$SCOPE" | python3 -c 'import json,sys; i=json.load(sys.stdin)[0]; print(i.get("attributes",{}).get("bucket-name",""))' 2>/dev/null)"
        GOT_SCOPE="$(printf '%s' "$SCOPE" | python3 -c 'import json,sys; i=json.load(sys.stdin)[0]; print(i.get("attributes",{}).get("sub-path",""))' 2>/dev/null)"
        if [ "$GOT_PROVIDER" = "$CRED_PROVIDER" ]; then pass "scope names provider $GOT_PROVIDER"
        else fail "scope names provider '$GOT_PROVIDER', want '$CRED_PROVIDER'"; fi
        if [ "$GOT_BUCKET" = "$OSS_BUCKET" ]; then pass "scope bucket-name = $GOT_BUCKET"
        else fail "scope bucket-name = '$GOT_BUCKET', want '$OSS_BUCKET'"; fi
        if [ "$GOT_SCOPE" = "$WANT_SCOPE" ]; then
            if [ -n "$WANT_SCOPE" ]; then pass "scope sub-path = $GOT_SCOPE (volume path $OSS_PATH + subPath $SUB_PATH)"
            else pass "scope carries no sub-path, matching a claim without one"; fi
        else
            fail "scope sub-path = '${GOT_SCOPE:-<none>}', want '${WANT_SCOPE:-<none>}' — the credential would not cover what the client reads"
        fi
    fi
fi

step "mount"
mounted=0
LOG=""
# A mount attempt that fails consumes the sandbox it was made on: the platform recycles
# it and the claim retries against a fresh one from the pool. Polling the pod discovered
# above therefore polls something already gone — it reports failure for a claim that goes
# on to succeed, and at cleanup the sandbox that did get claimed is left standing, because
# it is named for the pool rather than for this run. Re-resolving from the claim's label
# each round fixes both; oldest first, so POD ends up on the newest.
for _ in $(seq 1 $((TIMEOUT / 5))); do
    for sbx in $(kubectl -n "$NS" get sandbox -l "agents.kruise.io/claim-name=$SC_" \
                    --sort-by=.metadata.creationTimestamp \
                    -o jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}' 2>/dev/null); do
        POD="$sbx"
        if kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar 2>/dev/null \
                | grep -q '^CSITEST_MOUNTED '; then
            mounted=1; break
        fi
    done
    [ "$mounted" = 1 ] && break
    sleep 5
done

if [ "$mounted" = 1 ]; then
    pass "mock reported the tmpfs mount"
    LOG="$(kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar 2>/dev/null)"
    expect() { # $1=var $2=want
        local got; got="$(printf '%s' "$LOG" | sed -n "s/^CSITEST_ENV $1=//p" | tail -1)"
        if [ "$got" = "$2" ]; then pass "\$$1 = $got"; else fail "\$$1 = '${got:-<none>}', want '$2'"; fi
    }
    if [ -n "$CRED_PROVIDER" ]; then
        expect source "$OSS_BUCKET"
        expect bucket "$OSS_BUCKET"
        expect url "$OSS_URL"
    else
        expect source "mock://sandbox"
        expect bucket "mock-bucket"
        expect url "mock.endpoint.invalid"
    fi
    # No controller runs here to fill capacity in, so it can only have come from
    # volumeAttributes. Precedence over spec.mountOptions is a runc case: that is
    # resolved in the CSI flow around kubelet, which a claim never traverses.
    expect capacity "33Gi"

    if [ -n "$CRED_PROVIDER" ]; then
        expect authType "agent-identity"
        cdir="$(printf '%s' "$LOG" | sed -n 's/^CSITEST_ENV credentialDir=//p' | tail -1)"
        case "$cdir" in
            /*) pass "\$credentialDir = $cdir" ;;
            *)  fail "\$credentialDir = '${cdir:-<none>}', want an absolute path" ;;
        esac

        for field in AccessKeyId AccessKeySecret SecurityToken Expiration; do
            if printf '%s' "$LOG" | grep -qx "CSITEST_CREDFILE $field"; then
                pass "credential field $field delivered"
            else
                fail "credential field $field missing from $cdir"
            fi
        done

        # The exchange options configure this driver's machinery and would reach the
        # client as environment variables if they were not stripped, handing a FUSE
        # binary the means to mint credentials for itself.
        for leaked in sandboxCredProviderName jwtauth_endpoint jwtauth_tokenfile; do
            if printf '%s' "$LOG" | grep -qx "CSITEST_ENVKEY $leaked"; then
                fail "$leaked reached the entrypoint; infra options were not stripped"
            else
                pass "$leaked stripped before the entrypoint"
            fi
        done

        # The assertion that matters: the credential works. Files can arrive complete
        # and still authorise nothing, which is what a scope that does not cover the
        # mounted prefix looks like. The mock reports the prefix it listed, so a
        # denial can be told apart from having asked about the wrong one.
        probe="$(printf '%s' "$LOG" | sed -n 's/^CSITEST_OSSPROBE initial=//p' | tail -1)"
        case "$probe" in
            ok*)    pass "exchanged credential lists oss://$OSS_BUCKET/${probe#ok prefix=}" ;;
            skip:*) fail "probe skipped (${probe#skip:}) — the credential was never exercised" ;;
            "")     fail "no probe result; is ossutil in the mock image?" ;;
            *)      fail "exchanged credential rejected by OSS: ${probe#fail:}" ;;
        esac

        # Values must never be printed, and the fields are read from files the test
        # cannot see, so the check is that nothing resembling them appears at all.
        if printf '%s' "$LOG" | grep -qE 'STS\.|LTAI[0-9A-Za-z]{8,}'; then
            fail "something shaped like a credential appears in the mock output"
        else
            pass "no credential material in the mock output"
        fi
    fi
else
    fail "no mount within ${TIMEOUT}s"
    # A mount that never happens was usually refused before it reached the client, and
    # the refusal is reported to the claim rather than to a container. By now the pod
    # may also have been recycled, which leaves its logs simply gone, so the claim goes
    # first. Its status.message is the obvious place to look but by this point the
    # controller has usually replaced the reason with a bare timeout summary; the
    # per-attempt events keep the full text, so both are read.
    MSG="$(kubectl -n "$NS" get "sandboxclaim/$SC_" -o jsonpath='{.status.message}' 2>/dev/null)"
    [ -n "$MSG" ] && printf '    claim status.message: %s\n' "$MSG" >&2
    kubectl -n "$NS" get events --field-selector "involvedObject.name=$SC_" \
        --sort-by=.lastTimestamp -o jsonpath='{range .items[*]}{.message}{"\n"}{end}' 2>/dev/null \
        | grep -v '^[[:space:]]*$' | tail -3 | sed 's/^/    claim event: /' >&2
    kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar --tail=30 2>/dev/null | sed 's/^/    /' >&2
    kubectl -n "$NS" logs "$POD" -c csi-sidecar --tail=30 2>/dev/null | sed 's/^/    /' >&2
fi

if [ -n "$CRED_PROVIDER" ]; then
    if [ "${ROTATION_WAIT:-0}" -gt 0 ]; then
        step "credential rotation"
        # Two intervals decide whether a replacement can be seen at all, and neither
        # is set here. On this side it is the mount-proxy's refresh margin: once that
        # reaches the credential's lifetime the loop is floored at its minimum sleep
        # and asks again as often as it can. On the issuing side it is the credential's
        # lifetime, because an issuer that caches hands back byte-identical material
        # until half of it has elapsed — asking sooner returns the same credential,
        # which the mock rightly refuses to report as a rotation. Raise the margin and
        # lower the lifetime to make this a short test.
        rotated=0
        for _ in $(seq 1 $((ROTATION_WAIT / 15))); do
            if kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar 2>/dev/null \
                    | grep -q '^CSITEST_CREDROTATED '; then
                rotated=1; break
            fi
            sleep 15
        done
        if [ "$rotated" = 1 ]; then
            RLOG="$(kubectl -n "$NS" logs "$POD" -c csi-customfuse-agent-sidecar 2>/dev/null)"
            rline="$(printf '%s' "$RLOG" | grep '^CSITEST_CREDROTATED ' | tail -1)"
            pass "${rline#CSITEST_}"
            n="$(printf '%s' "$rline" | sed -n 's/.*n=\([0-9]*\).*/\1/p')"
            # Punctuality is not the point. A refresh loop can replace the files on
            # schedule and hand over a credential that no longer authorises anything,
            # which is exactly the failure this step exists to catch, so the material
            # that arrived second is probed like the first one was.
            rprobe="$(printf '%s' "$RLOG" | sed -n "s/^CSITEST_OSSPROBE rotated$n=//p" | tail -1)"
            case "$rprobe" in
                ok*)    pass "rotated credential lists oss://$OSS_BUCKET/${rprobe#ok prefix=}" ;;
                skip:*) fail "rotated credential not exercised (${rprobe#skip:})" ;;
                "")     fail "the credential that replaced the first was never probed" ;;
                *)      fail "rotated credential rejected by OSS: ${rprobe#fail:}" ;;
            esac
        else
            fail "the delivered credential was not replaced within ${ROTATION_WAIT}s"
            printf '    the mock reports a replacement only when the content differs.\n' >&2
            printf '    an issuer that caches returns the same bytes until half the\n' >&2
            printf '    credential lifetime has elapsed, so either raise ROTATION_WAIT\n' >&2
            printf '    or lower that lifetime.\n' >&2
        fi
    else
        skip "rotation not exercised (ROTATION_WAIT=0)"
    fi
fi

step "visibility from the workload container"
# What makes the mount visible there is propagation, not the mount: the sidecar
# mounts under a volume shared with the workload, and a mount created after the
# workload started reaches it only through a mode that forwards it. Both ends are
# asserted, because either one on its own reads as correctly configured and still
# leaves the workload looking at an empty directory.
VIS="$(kubectl -n "$NS" get "pod/$POD" -o json 2>/dev/null | python3 -c '
import json, sys

# A sandbox whose mount never succeeded is recycled within seconds, so by the time
# this runs the pod is often simply gone. That is an absence of evidence rather than
# a misconfigured pod, and the mount step has already reported the real failure, so
# it is named here instead of falling out of the parser as an empty string.
raw = sys.stdin.read()
if not raw.strip():
    print("pod-gone"); sys.exit(0)

spec = json.loads(raw)["spec"]
sidecars = spec.get("initContainers") or []
workload = (spec.get("containers") or [None])[0]
agent = next((c for c in sidecars
              if c.get("name") == "csi-customfuse-agent-sidecar"), None)
if agent is None or workload is None:
    print("missing-container"); sys.exit(0)

agent_root = next((m for m in agent.get("volumeMounts") or []
                   if m.get("mountPath") == "/run/csi/mount-root"), None)
if agent_root is None:
    print("agent-no-mount-root"); sys.exit(0)
if agent_root.get("mountPropagation") != "Bidirectional":
    print("agent-not-bidirectional"); sys.exit(0)

# Matched on the volume rather than the path: the workload is free to mount the
# shared root wherever it likes, and a path comparison would fail on a config that
# is wired correctly.
app_root = next((m for m in workload.get("volumeMounts") or []
                 if m.get("name") == agent_root.get("name")), None)
if app_root is None:
    print("workload-no-mount-root"); sys.exit(0)
if app_root.get("mountPropagation") not in ("HostToContainer", "Bidirectional"):
    print("workload-no-propagation"); sys.exit(0)
print("ok %s %s" % (workload.get("name"), app_root["mountPath"]))
' 2>/dev/null)"

case "$VIS" in
    ok*)
        APP="${VIS#ok }"; APP="${APP%% *}"
        ROOT="${VIS##* }"
        pass "$APP shares the sidecar's mount-root volume at $ROOT"
        pass "propagation on both ends forwards a mount made after the workload started"
        # Reading the mock's marker through the workload's own root is the same claim
        # stated from inside it. It needs exec, which a sandbox pod may not offer, so
        # an inability to run it is reported rather than counted against the driver.
        MOUNTED="$(printf '%s' "$LOG" | sed -n 's/^CSITEST_MOUNTED //p' | tail -1)"
        REL="${MOUNTED#/run/csi/mount-root/}"
        if [ -n "$MOUNTED" ] && [ "$REL" != "$MOUNTED" ] && [ "${REL#/}" = "$REL" ] \
            && kubectl -n "$NS" exec "$POD" -c "$APP" -- test -f "$ROOT/$REL/.csitest-marker" >/dev/null 2>&1; then
            pass "marker visible in $APP at $ROOT/$REL"
        else
            skip "could not exec into $APP to read the marker through the mount"
        fi
        ;;
    agent-no-mount-root)     fail "the customfuse sidecar has no /run/csi/mount-root mount" ;;
    agent-not-bidirectional) fail "the sidecar's mount-root is not Bidirectional, so the mount stays inside it" ;;
    workload-no-mount-root)  fail "the workload container does not mount the sidecar's mount-root volume" ;;
    workload-no-propagation) fail "the workload's mount-root forwards no propagation, so later mounts stay invisible to it" ;;
    missing-container)       fail "could not read the customfuse sidecar or the workload container off $POD" ;;
    pod-gone)                skip "$POD was recycled before its spec could be read, so propagation is unverified" ;;
    *)                       fail "unexpected result from the pod spec: ${VIS:-<empty>}" ;;
esac

printf '\n'
if [ "$failures" -ne 0 ]; then
    echo "FAILED: $failures check(s)" >&2
    exit 1
fi
echo "PASS: customfuse behaves correctly under sandbox injection"
