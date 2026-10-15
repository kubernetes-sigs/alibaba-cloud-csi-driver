#!/bin/bash
# Asserts that customfuse is deployed as an independent component: its sidecars,
# volumes and driver registration must live only in its own DaemonSet and
# Deployment, never in the ones shared by the other drivers.
#
# Getting this wrong is not a rendering error — the chart still installs, but the
# customfuse provisioner and attacher run twice, in two Deployments, competing for
# the same leader-election lease. That is invisible in `helm lint` and in
# server-side validation, so it is checked here.
#
# Needs only helm; no cluster. Run: bash test/helm/customfuse-isolation.sh

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
CHART_DIR="$PROJECT_ROOT/deploy/charts/alibaba-cloud-csi-driver"

HELM="${HELM:-helm}"
command -v "$HELM" >/dev/null || { echo "helm not found; set HELM=/path/to/helm" >&2; exit 1; }

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

failures=0
fail() { echo "  FAIL: $*" >&2; failures=$((failures + 1)); }
pass() { echo "  ok: $*"; }

render() { # $1=output file, rest=helm --set args
    local out=$1; shift
    "$HELM" template test "$CHART_DIR" "$@" > "$out" 2>/dev/null
}

# resources_named prints "kind name" for every document, so a resource can be
# located without a YAML parser being available.
resources_named() {
    awk '
        /^kind: /   { kind = $2 }
        /^  name: / { if (kind != "" && name == "") { name = $2; print kind, name } }
        /^---$/     { kind = ""; name = "" }
    ' "$1"
}

# doc_of extracts the single document for "kind/name", so the checks below can
# look inside one resource rather than the whole render.
doc_of() { # $1=file $2=kind $3=name
    awk -v want_kind="$2" -v want_name="$3" '
        BEGIN { RS = "\n---\n" }
        {
            kind = ""; name = ""
            n = split($0, lines, "\n")
            for (i = 1; i <= n; i++) {
                if (lines[i] ~ /^kind: /)   { split(lines[i], a, " "); kind = a[2] }
                if (lines[i] ~ /^  name: / && name == "") { split(lines[i], b, " "); name = b[2] }
            }
            if (kind == want_kind && name == want_name) print $0
        }
    ' "$1"
}

echo "== customfuse disabled: nothing customfuse anywhere =="
render "$WORK/off.yaml"
if grep -qi customfuse "$WORK/off.yaml"; then
    fail "disabled chart still references customfuse:"
    grep -in customfuse "$WORK/off.yaml" | head -5 >&2
else
    pass "no references"
fi

echo "== customfuse enabled =="
render "$WORK/on.yaml" --set csi.customfuse.enabled=true --set csi.customfuse.controller.enabled=true

# The shared workloads must stay untouched. Their names are the ones every other
# driver shares; anything customfuse in them is the duplication this guards.
for shared in csi-plugin csi-provisioner; do
    for kind in DaemonSet Deployment; do
        doc="$(doc_of "$WORK/on.yaml" "$kind" "$shared")"
        [ -n "$doc" ] || continue
        hits="$(printf '%s\n' "$doc" | grep -i customfuse || true)"
        if [ -n "$hits" ]; then
            fail "shared $kind/$shared references customfuse:"
            printf '%s\n' "$hits" | sed 's/^/        /' >&2
        else
            pass "shared $kind/$shared is clean"
        fi
    done
done

# Its own workloads must exist, otherwise "clean shared workloads" would also be
# satisfied by customfuse not being deployed at all.
for own in "DaemonSet csi-customfuse-plugin" "Deployment csi-customfuse-provisioner"; do
    set -- $own
    if [ -n "$(doc_of "$WORK/on.yaml" "$1" "$2")" ]; then
        pass "own $1/$2 exists"
    else
        fail "own $1/$2 missing"
    fi
done

# Each sidecar may appear once. Two copies of external-customfuse-provisioner
# across two Deployments is exactly the leader-election conflict described above.
echo "== no duplicated customfuse sidecars =="
for sidecar in external-customfuse-provisioner external-customfuse-attacher; do
    count="$(grep -c -- "- name: $sidecar\$" "$WORK/on.yaml" || true)"
    if [ "$count" = "1" ]; then
        pass "$sidecar appears once"
    else
        fail "$sidecar appears $count times (expected 1)"
    fi
done

# The driver list decides which plugin serves which CSI driver. customfuse
# belonging to the shared list is what makes the shared plugin answer for it.
echo "== driver lists are separated =="
shared_drivers="$(doc_of "$WORK/on.yaml" DaemonSet csi-plugin | grep -o '\-\-driver=[^" ]*' | head -1 || true)"
own_drivers="$(doc_of "$WORK/on.yaml" DaemonSet csi-customfuse-plugin | grep -o '\-\-driver=[^" ]*' | head -1 || true)"
case "$shared_drivers" in
    *customfuse*) fail "shared csi-plugin lists customfuse: $shared_drivers" ;;
    "")           fail "could not read shared csi-plugin driver list" ;;
    *)            pass "shared csi-plugin: $shared_drivers" ;;
esac
case "$own_drivers" in
    *customfuse*) pass "own csi-customfuse-plugin: $own_drivers" ;;
    *)            fail "own csi-customfuse-plugin driver list is $own_drivers, expected customfuse" ;;
esac

# Both node plugins run on the host network, where a container port is a host port:
# whichever DaemonSet is scheduled second stays Pending with "didn't have free
# ports for the requested pod ports". The chart renders and validates either way,
# so the two port sets are compared here.
echo "== host-network ports do not collide =="
doc_of "$WORK/on.yaml" DaemonSet csi-plugin            > "$WORK/shared-ds.yaml"
doc_of "$WORK/on.yaml" DaemonSet csi-customfuse-plugin > "$WORK/own-ds.yaml"

if grep -q 'hostNetwork: true' "$WORK/own-ds.yaml"; then
    shared_ports="$(grep -o 'containerPort: [0-9]*' "$WORK/shared-ds.yaml" | awk '{print $2}' | sort -u)"
    own_ports="$(grep -o 'containerPort: [0-9]*' "$WORK/own-ds.yaml" | awk '{print $2}' | sort -u)"
    clash="$(comm -12 <(printf '%s\n' "$shared_ports") <(printf '%s\n' "$own_ports") || true)"
    if [ -n "$clash" ]; then
        fail "csi-plugin and csi-customfuse-plugin both claim host port(s): $(echo $clash)"
    else
        pass "no shared host port (csi-plugin: $(echo $shared_ports); customfuse: $(echo $own_ports))"
    fi

    # The probes reach the port by name, so the process has to bind the port the
    # container declares. SERVICE_PORT is what it binds; leaving it at the default
    # makes the probes address a port nothing listens on.
    declared="$(grep -A1 'name: healthz' "$WORK/own-ds.yaml" | grep -o 'containerPort: [0-9]*' | awk '{print $2}' | head -1)"
    bound="$(grep -A1 'name: SERVICE_PORT' "$WORK/own-ds.yaml" | grep 'value:' | tr -dc '0-9')"
    if [ -n "$declared" ] && [ "$declared" = "$bound" ]; then
        pass "customfuse binds the port it declares ($declared)"
    else
        fail "customfuse declares containerPort ${declared:-none} but SERVICE_PORT is ${bound:-unset}"
    fi
else
    pass "customfuse node plugin is off the host network; ports are not host-wide"
fi

# Every other provisioner in this chart names its one client here (ext4, nfs,
# ossfs, bmcpfs). customfuse has no single client, and the sidecar falls back to
# ext4 when the flag is absent, so the default has to be the generic marker the
# driver ignores. Any real client name lands in every dynamically provisioned PV
# as fsType and then contradicts the fuseType the volume asked for.
echo "== the provisioner defaults fsType to the generic marker =="
own_ctrl="$(doc_of "$WORK/on.yaml" Deployment csi-customfuse-provisioner)"
fstype_arg="$(printf '%s\n' "$own_ctrl" | grep -o -- '--default-fstype=[^" ]*' | head -1 || true)"
case "$fstype_arg" in
    --default-fstype=customfuse) pass "$fstype_arg" ;;
    "")                          fail "csi-customfuse-provisioner sets no --default-fstype; the sidecar then defaults it to ext4" ;;
    *)                           fail "csi-customfuse-provisioner sets $fstype_arg, a client name; only the generic marker is safe as a default" ;;
esac

# The CSIDriver object is the one customfuse resource that does belong to a shared
# template, since neither customfuse template declares it. Losing it makes every
# mount fail with "driver not found", so its presence is asserted rather than
# assumed.
echo "== CSIDriver registration =="
if [ -n "$(doc_of "$WORK/on.yaml" CSIDriver customfuseplugin.csi.alibabacloud.com)" ]; then
    pass "CSIDriver customfuseplugin.csi.alibabacloud.com present"
else
    fail "CSIDriver customfuseplugin.csi.alibabacloud.com missing"
fi

# controller.enabled=false must drop the controller side without disturbing the
# node side, so the two switches are checked to be independent.
echo "== controller disabled, node side still deployed =="
render "$WORK/nocontroller.yaml" --set csi.customfuse.enabled=true --set csi.customfuse.controller.enabled=false
if [ -n "$(doc_of "$WORK/nocontroller.yaml" Deployment csi-customfuse-provisioner)" ]; then
    fail "csi-customfuse-provisioner rendered although controller.enabled=false"
else
    pass "no customfuse controller"
fi
if [ -n "$(doc_of "$WORK/nocontroller.yaml" DaemonSet csi-customfuse-plugin)" ]; then
    pass "customfuse node plugin still rendered"
else
    fail "customfuse node plugin missing"
fi

# Enabling neighbours must not pull customfuse back into their workloads.
echo "== alongside other drivers =="
render "$WORK/multi.yaml" \
    --set csi.customfuse.enabled=true --set csi.customfuse.controller.enabled=true \
    --set csi.ens.enabled=true --set csi.bmcpfs.enabled=true
for shared in csi-plugin csi-provisioner; do
    for kind in DaemonSet Deployment; do
        doc="$(doc_of "$WORK/multi.yaml" "$kind" "$shared")"
        [ -n "$doc" ] || continue
        if printf '%s\n' "$doc" | grep -qi customfuse; then
            fail "with ens+bmcpfs enabled, shared $kind/$shared references customfuse"
        else
            pass "shared $kind/$shared still clean"
        fi
    done
done

echo
if [ "$failures" -ne 0 ]; then
    echo "FAILED: $failures check(s)" >&2
    exit 1
fi
echo "PASS: customfuse renders as an independent component"
