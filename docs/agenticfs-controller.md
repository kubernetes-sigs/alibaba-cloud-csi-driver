# AgenticFS controller architecture

The AgenticFS mode uses the existing NAS controller factory and cloud interfaces.
The normal resource model is one PVC → one AgenticSpace → one AccessPoint. It does
not introduce a separate controller framework, background reconciler, or rollback
transaction. See the [usage and recovery guide](../examples/nas/agenticfs/README.md)
for operational behavior.

## Responsibilities

| File | Responsibility |
| --- | --- |
| `pkg/nas/controllerserver.go` | CSI dispatch, per-volume locking, PV lookup, shared volume attributes. |
| `pkg/nas/agenticfs_controller.go` | Dependency initialization and create/delete/expand orchestration; AgenticSpace API calls. |
| `pkg/nas/agenticfs_options.go` | Request validation, CNFS resolution, typed `agenticfsVolumeArgs`, and CSI volume construction. |
| `pkg/nas/agenticfs_quota.go` | Capacity arithmetic and expansion policy; build quota requests without mutating cloud responses. |
| `pkg/nas/agenticfs_accesspoint.go` | Discover/create accesspoints, verify listings, prepare deletion order, and poll state. |
| `pkg/nas/agenticfs_errors.go` | OpenAPI error classification and provisioning-failure diagnostics. |

These files remain in package `nas`, like the other NAS volume modes. Typed volume
arguments follow the NAS/Disk parameter-object pattern. The controller retains
only the dependencies it uses: `NasClientV2Interface`, `CNFSGetter`, `clock.Clock`,
region, and polling durations. Per-request resource IDs and parsed arguments are
local variables, not caches on the controller.

## Lifecycle contracts

### Creation

```text
validate name
  → resolve filesystem (direct ID takes precedence over CNFS)
  → validate placement and quota
  → CreateAgenticSpace(ClientToken = volume name)
  → discover/reuse or create AccessPoint
  → wait for Active and a non-empty domain
  → construct CSI Volume
```

Validation order is intentional and determines which error is returned when
several parameters are invalid. No billable NAS call runs before validation.
The parsed argument object is independent of the request map; volume construction
allocates a new attributes map.

`ensureAccessPoint` discovers before creating, because this API has no
ClientToken. It returns the observed ID before polling so later failures can log
it. Every page and ownership relationship is validated before a discovered
accesspoint is used. Unexpected multiple entries retain the existing Active-first
recovery behavior, not a cross-VPC placement policy.

A failed CreateVolume never deletes resources. Error classification controls the
CSI response and diagnostic severity only. A missing resource ID means the request
did not observe it; it does not prove a previous attempt created nothing.

### Deletion

```text
validate stored identifiers
  → obtain a verified accesspoint deletion list, or confirm space absence
  → delete each accesspoint and wait until absent
  → DeleteAgenticSpace(ClientToken = volume ID)
```

`accessPointsForDeletion` returns an explicit `spaceGone` flag. A failed or partial
listing is never treated as an empty space. An ambiguous listing NotFound triggers
a space-scoped Get; an explicitly missing accesspoint cannot prove space absence.

Deletion order is deterministic: the PV's accesspoint comes first, then each unique
listed ID in first-occurrence order. This preserves recovery when the listing lags
the PV. The parent filesystem is never deleted.

### Expansion

```text
validate stored identifiers and requested size
  → GetAgenticSpace
  → prepareAgenticSpaceExpansion(snapshot)
  → SetAgenticSpaceQuota
```

The quota policy helper performs no API calls and does not modify the read response.
It rejects incomplete usage data and shrinking, preserves valid file-count limits,
and retains the existing warning/omission rules for unreadable or invalid limits.
Same-size requests still issue SetAgenticSpaceQuota. The refactor does not add an
expansion ClientToken or change the SDK request shape.

## Time and cancellation

Polling uses an injected `k8s.io/utils/clock.Clock`, following Disk's waiter pattern.
Production uses `clock.RealClock`; tests advance a fake clock rather than sleeping
for the production interval or timeout.

The existing ordering is preserved:

1. Check cancellation before each Describe call.
2. Observe the response and test the target state.
3. Check whether the polling budget has elapsed.
4. Arm the next interval and wait for the timer or cancellation.

Consequently, a call returning the target state after the polling budget still
wins. The timer is stopped on every exit. The polling budget is not a guarantee on
total SDK duration: credential resolution and non-context-aware SDK calls retain
their existing cancellation behavior. Clock injection does not change that contract.

## Testing

Existing controller and deletion suites remain end-to-end regression tests over a
fake NAS client. Focused tests cover the extracted boundaries:

- `agenticfs_options_test.go`: validation order, direct-ID precedence, immutable
  request snapshots, and independent volume attribute maps.
- `agenticfs_quota_test.go`: quota request shape, unchanged cloud snapshots,
  incomplete/unsafe input rejection, and capacity fuzz invariants.
- `agenticfs_accesspoint_test.go`: deterministic polling, cancellation, timer
  cleanup, completion-versus-timeout ordering, and stable deletion order.
- `pkg/nas/cloud` and `pkg/cloud/wrap`: real vendored SDK wire contracts and error
  conversion, independent of controller policy tests.

Run on Linux:

```sh
go test -race ./pkg/nas/... ./pkg/cloud/wrap
go test ./pkg/nas -run '^$' -fuzz '^FuzzComputeAgenticSpaceSizeLimit$' -fuzztime=10s -parallel=2
go build ./cmd/...
```

When changing policy, test both the returned CSI status and whether subsequent NAS
calls were prevented. In particular, never weaken ownership validation, treat a
failed listing as empty, add compensation on creation failure, or infer space
absence from an accesspoint error merely to make a retry succeed.
