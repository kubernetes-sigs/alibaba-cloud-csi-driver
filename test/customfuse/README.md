# customfuse end-to-end tests

Exercises the customfuse CSI paths with a mock FUSE client, so a change to the
driver can be checked without a bucket, a credential, or a client binary.

The mock reports the environment the driver delivered and then mounts a tmpfs.
That is enough: a tmpfs makes the target a mount point, which is the only thing
mount-proxy waits for, so every RPC from CreateVolume through NodeUnstage runs
exactly as it would with a real client. What a real client would add — talking to
storage — says nothing about whether the driver did its job.

| File | Covers |
|------|--------|
| [mock-entrypoint.sh](mock-entrypoint.sh) | The stand-in client. Reports `CSITEST_*` lines, mounts a tmpfs, stays alive |
| [run-runc.sh](run-runc.sh) | DaemonSet deployment: static PV, dynamic provisioning, mountOptions precedence, mount visibility, teardown |
| [run-sandbox.sh](run-sandbox.sh) | Sandbox injection: the sidecars the runtime injected, socket resolution by flag, mount and capacity without a controller, agent-identity credential exchange, its scope against the claim's subPath, and its rotation |

## Build the mock image

```bash
docker build -f test/customfuse/Dockerfile \
    -t <registry>/customfuse-mock:<tag> .
docker push <registry>/customfuse-mock:<tag>
```

## DaemonSet deployment

Register the image so the driver knows what to run as the fuse pod, then run the
script:

```bash
kubectl -n kube-system patch cm csi-plugin --type=merge \
    -p '{"data":{"fuse-customfuse-mock":"image=<registry>/customfuse-mock:<tag>"}}'

MOCK_IMAGE=<registry>/customfuse-mock:<tag> bash test/customfuse/run-runc.sh
```

Everything it creates is labelled `csitest=<PREFIX>` and removed on exit,
including PVs whose finalizers would otherwise hold them.

## Sandbox deployment

The platform that turns a claim into a pod is assumed installed and configured: a
runtime that injects the customfuse sidecars, and an image pull path that reaches the
mock. Neither is part of this repository and neither is configured by these scripts,
which change nothing cluster-wide — they only create and delete resources labelled
`csitest=<PREFIX>`. A failure on the platform side is reported rather than diagnosed
here: when no sandbox appears, the script prints the claim's status message and its
events, which is where the platform states why it did not bind.

Two things the script cannot assume, and therefore checks:

- the injected `agent-runtime` must be a build that understands the customfuse
  driver, since it is what forwards the volume.
- the mock must be pullable by the pool's ServiceAccount. Injection adds containers
  and volumes to a pod, not `imagePullSecrets`, so a mock in a private registry needs
  the credential on the ServiceAccount itself:

  ```bash
  kubectl -n <ns> patch sa <pool-sa> \
      -p '{"imagePullSecrets":[{"name":"<secret>"}]}'
  ```

  The script reports whether that credential is reachable before it creates anything,
  because otherwise the sidecar sits in `ErrImagePull` and the CSI logs say nothing
  about why.

`RUNTIME` names the runtime that injects customfuse, and defaults to
`csi-customfuse`; override it if your platform wires customfuse through a runtime
shared with the other drivers. What the script asserts on is the pod that came
back — its containers, their arguments, their mounts — rather than any
configuration schema, so a platform that wires it differently still gets a
meaningful answer instead of a check against a layout it never promised.

```bash
MOCK_IMAGE=<registry>/customfuse-mock:<tag> \
SANDBOXSET=<pool declaring the runtime> \
bash test/customfuse/run-sandbox.sh
```

Leave `SANDBOXSET` empty and pass `SANDBOX_IMAGE` instead to have the script create a
one-replica pool of its own, which it removes on exit like everything else.

### agent-identity

Pass `CRED_PROVIDER` to switch the volume from a Secret to a credential exchanged
per mount. It needs a bucket the exchanged credential can actually be tried
against, because the assertions that matter are the ones a credential either
passes or fails:

```bash
MOCK_IMAGE=<registry>/customfuse-mock:<tag> \
SANDBOXSET=<pool declaring the runtime> \
CRED_PROVIDER=<credential provider> \
OSS_BUCKET=<bucket> OSS_URL=<endpoint> \
OSS_PATH=<prefix in the bucket> SUB_PATH=<subdirectory of that prefix> \
bash test/customfuse/run-sandbox.sh
```

`SUB_PATH` is relative to `OSS_PATH`, not to the bucket, and the same string is
what the credential is scoped to. The script asserts the two agree: a scope
pointing elsewhere still issues a credential and the mount still succeeds, and
only then does every object operation fail, which reads as a storage problem and
is not one.

Rotation is off by default, and turning it on needs both ends asking often enough
— neither interval is this script's to set:

- the mount-proxy renews a credential this long before it expires
  (`AGENT_IDENTITY_TOKEN_REFRESH_MARGIN` on the injected sidecar). Once the margin
  reaches the credential's lifetime there is nothing left to wait for and the loop
  falls back to its 30s floor, asking as often as it can.
- the lifetime the issuer is configured to grant. An issuer that caches returns
  byte-identical material until half of that lifetime has elapsed, so asking
  sooner only collects the same credential — which the mock correctly refuses to
  report as a rotation.

With both set, add `ROTATION_WAIT=<seconds>`. The mock fingerprints what is on
disk every 15s and probes again each time it changes, so the script can check
that the credential which arrived *second* also works. That is the point: a
refresh loop can be punctual and hand over a credential that authorises nothing.

## Reading the output

The mock prefixes everything it reports, so the scripts assert on it rather than
on a human reading logs:

```
CSITEST_ENV source=mock://cfmock-static     one line per variable the driver should set
CSITEST_ENVKEY <name>                       every variable present, names only
CSITEST_CREDFILE AccessKeyId                files under $credentialDir, names only
CSITEST_MOUNTED /run/fuse.customfuse/…      the tmpfs is up
CSITEST_OSSPROBE initial=ok prefix=sub/     the delivered credential was tried against the bucket
CSITEST_CREDROTATED n=1 at=18:24:12Z        what is on disk was replaced
CSITEST_OSSPROBE rotated1=ok prefix=sub/    and the replacement was tried too — n pairs the two
```

Values are printed only for known-safe variables. Everything else is reported by
name alone, because Secret entries arrive as environment variables too and a test
image must not print a credential.

To read them by hand:

```bash
kubectl -n ack-csi-customfuse logs \
    -l csi.alibabacloud.com/volume-id=<volume-handle> | grep CSITEST_
```

## Settings

| Variable | Default | Meaning |
|----------|---------|---------|
| `MOCK_IMAGE` | — | Required. The mock image built above |
| `BUSYBOX_IMAGE` | `registry.cn-shanghai.aliyuncs.com/eci_open/busybox:1.30` | `run-runc.sh` only. The consumer container. Any image with a shell works — the pod overrides `command` |
| `FUSE_TYPE` | `customfuse-mock` | Must match the `fuse-<type>` key in the csi-plugin ConfigMap |
| `NS` | `default` | Where workloads and claims are created |
| `FUSE_NS` | `ack-csi-customfuse` | Where the driver creates fuse pods |
| `PREFIX` | `cfmock` / `cfsbx` | Name prefix and cleanup label |
| `SANDBOXSET` | — | SandboxSet to claim from; must declare `RUNTIME` in `spec.runtimes` |
| `SANDBOX_IMAGE` | — | Required only when `SANDBOXSET` is empty and the script creates the pool |
| `RUNTIME` | `csi-customfuse` | The runtime that injects the customfuse sidecars |
| `SOCKET` | `/run/cnfs/customfuse-mounter.sock` | What the injected `csi-sidecar` is asserted to pass as `--customfuse-mount-proxy-sock` |
| `NO_CLEANUP` | `0` | `1` keeps the claim, sandbox and pod so a failure can be read out of the logs |
| `TIMEOUT` | `180` / `240` | Seconds to wait for readiness |
| `CRED_PROVIDER` | — | Set to switch the volume to agent-identity; names the CredentialProvider |
| `OSS_BUCKET`, `OSS_URL`, `OSS_PATH` | —, —, `/` | The bucket the exchanged credential has to work against |
| `SUB_PATH` | — | Subdirectory of `OSS_PATH` to mount, relative to it; also what scopes the credential |
| `ROTATION_WAIT` | `0` | Seconds to watch for the delivered credential to be replaced; `0` skips the wait |
| `AGENT_IDENTITY` | first Available | AgentIdentity the claim carries, so a token is issued to the sandbox |

## When something fails

Both scripts print the failing assertion with what they got and what they
expected, then dump pod events and the relevant container logs. A failure in the
mock's own output — a variable arriving empty or wrong — points at the driver; a
failure to reach `CSITEST_MOUNTED` points at mount-proxy or the fuse pod itself.

In agent-identity mode there is a further distinction. A mount that never appears
points at the exchange, and the sidecar log says why. A mount that does appear
with a failing `CSITEST_OSSPROBE` is a credential that was issued and then
rejected, which is a question of what it was scoped to rather than of the driver:
compare the prefix on that line against the one the volume is mounted at.

When the exchange itself is refused, the status code says which of two unrelated
problems it is. A 400 means the provider's policy template could not be rendered
for this volume — a claim that carries no subPath hits this against a template
whose scoping requires one, which is a property of the template and not of the
driver. A 403 means it rendered fine but nothing authorizes this agent to ask that
provider, which is a missing role rule naming it. Neither is a mount failure, and
both are reported by the credential service rather than by anything in this
repository.

Under sandbox injection, a failure before any container starts is the platform's
to explain. The script prints the claim's status message and its recent events.
Read the events: by the time a wait has timed out the status message has usually
been overwritten with a summary of the wait itself, which can even read as a
success, while the per-attempt events still carry the reason.
