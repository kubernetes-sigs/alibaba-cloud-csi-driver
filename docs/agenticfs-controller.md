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
| `pkg/nas/agenticfs_quota.go` | Capacity arithmetic and quota validation. |
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
  → SetAgenticSpaceQuota(SizeLimit)
```

Expansion is a single call that sets the absolute size, so a retry or a same-size
request is harmless and needs no ClientToken. FileCountLimit is omitted, which
keeps the space's current limit. NAS also accepts a smaller SizeLimit; this
happens only when a user recovers from a failed expansion by lowering the PVC
request, and is applied as requested.

## Time and cancellation

Polling uses an injected `k8s.io/utils/clock.Clock`, following Disk's waiter pattern.
Production uses `clock.RealClock`; tests advance a fake clock rather than sleeping
for the production interval or timeout.

The polling loop runs in this order:

1. Check cancellation before each Describe call.
2. Observe the response and test the target state.
3. Check whether the polling budget has elapsed.
4. Arm the next interval and wait for the timer or cancellation.

Consequently, a call returning the target state after the polling budget still
wins. The timer is stopped on every exit. The polling budget is not a guarantee on
total SDK duration: credential refresh in progress cannot be interrupted.
Clock injection does not change that contract.

## Testing

The controller and deletion suites are end-to-end tests over a fake NAS client.
Focused tests cover the individual components:

- `agenticfs_options_test.go`: validation order, direct-ID precedence, immutable
  request snapshots, and independent volume attribute maps.
- `agenticfs_quota_test.go`: capacity fuzz invariants.
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

## OpenAPI rate limits

NAS meters the control plane per (region, action, account), so the driver keeps
one adaptive token bucket per action instead of sharing one budget:

| Variable | Meaning | Default |
|---|---|---|
| `NAS_LIMIT_PERSECOND` | aggregate ceiling across all actions | 200 |
| `NAS_ACTION_LIMIT_PERSECOND` | rate an action starts at | 40, the measured operating point |
| `NAS_ACTION_MAX_PERSECOND` | ceiling an action may grow to | the aggregate ceiling |

`CreateDir` starts at 20 QPS and `DescribeFileSystems` at 100 QPS from the
published quota table. Every other action starts at the per-action rate.

Each action discovers its own quota instead of being told it:

- Slow start doubles the rate every second while calls succeed: from the 40 QPS
  start, 40 - 80 - 160 - 200 takes about three seconds.
- The first `Throttling.*` response ends slow start and multiplies the rate by
  0.75. A burst of concurrent throttles counts as one event (5s window).
- From then on the rate grows by 1 QPS every 5 seconds, so a raised account
  quota is picked up without a restart, and a quota shared with another
  consumer is tracked without hammering it.

A throttled call is retried up to five times with jittered exponential backoff,
bounded by the request context, and the token is acquired before every attempt
so a retry still spends the action's budget.

Discovery costs one overshoot per action: at most one slow-start step beyond
the real quota (64 QPS against a 40 QPS quota is one second's worth of rejected
requests), after which every probe adds a single QPS. On the reference account,
300 concurrent AgenticFS volumes at 40 QPS per write action completed in 19.0s
with zero throttling, while an effective 200 took 149s with 100 throttling
responses.

`NAS_LIMIT_PERSECOND` used to be the only bucket. It now caps the aggregate
only, so a deployment that set it to a small value to protect the account
(the old 40) will keep that aggregate cap and see no per-action parallelism
until it is raised.
