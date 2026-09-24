# Mount a customfuse volume with agent-identity (STS credential) authentication

The customfuse driver can exchange a sandbox agent-identity token for a
short-lived, scoped STS credential instead of taking a long-lived AccessKey from
a Secret. The credential is delivered to the fuse container as files and kept
fresh for the lifetime of the mount, so no AccessKey is stored in the cluster.

Unlike the other drivers, customfuse runs a FUSE client chosen by the volume, so
it cannot know how that client wants a credential handed to it — `alinas` can
assume its client takes the STS triple as mount options, `ossfs` can assume its
client performs the exchange itself. customfuse can assume neither. It therefore
performs the exchange itself and delivers the result as files, leaving your
entrypoint to adapt them to your client.

## How it works

You supply the FUSE client and the script that mounts it. The driver supplies the
credential:

1. The volume sets `authType: agent-identity` and carries no Secret.
2. **Before your entrypoint starts**, the driver has exchanged the sandbox token
   for a scoped STS credential and written it to `$credentialDir`, one value per
   file.
3. Your entrypoint reads those files and mounts.
4. **Before the credential expires**, the driver writes a new one to the same
   directory. A client that re-reads the files picks it up on its own; one that
   does not can be told through a *refresh hook* — a script of yours that the
   driver runs after each rotation. See [Rotation](#rotation).
5. **When your entrypoint exits** — unmounted or crashed — the credential is
   removed.

A failed exchange fails the mount: there is no static credential to fall back to,
and a client started without one fails later and less clearly.

## What the driver guarantees, and what it does not

**The driver guarantees:**

* The credential is on disk before your entrypoint starts.
* What is on disk is always one complete credential — a rotation is never observed
  half-written. Reading it correctly is still on you; see
  [Reading the credential](#reading-the-credential).
* The credential is renewed before it expires, for as long as the mount lives.
* The credential is removed once the mount is gone.
* No credential value is ever passed to your entrypoint or hook as an argument or
  an environment variable — only the directory to read it from. Nothing lands in
  `/proc/<pid>/cmdline` or in another process's environment.

**The driver does not:**

* Assume anything about how your client consumes a credential, or prescribe how
  it should. Delivery is files; everything above that is yours.
* Guarantee your client notices a rotation. A client that re-reads the files does;
  one that does not must be told, which is what the refresh hook is for.
* Work around a client that cannot accept a new credential at all while running.
  Such a client needs a credential outliving the mount; no delivery mechanism
  helps.

**Your entrypoint is responsible for:**

* Composing the credential into whatever form your client takes — a JSON
  document, an ini profile, a properties file, command line arguments,
  environment variables.
* Reading all fields from a **single** resolved directory. See
  [Reading the credential](#reading-the-credential) — getting this wrong produces
  a credential that never existed, with no error raised.
* Telling your client about a rotation, via the refresh hook, if it does not
  re-read the files itself.

## Prerequisite

* A working Kubernetes cluster with the CSI plugin and customfuse driver enabled.
  Please refer to the [installation guide](./install.md) and
  [customfuse](./customfuse.md).
* An [ACS Agent Sandbox](https://help.aliyun.com/en/cs/user-guide/agent-sandbox).
  The token exchanged here is the sandbox's own, so where one is not running there
  is nothing to exchange. Under injection the driver is the sidecar itself and
  creates no fuse pod; what else that changes is the table in
  [Where the driver runs](./customfuse.md#where-the-driver-runs).
* The agent-identity endpoint and token directory are configured by the sandbox
  platform — not per volume, and not from this repository. They are the variables
  [nas-agent-identity.md](./nas-agent-identity.md#prerequisite) lists, and the
  only one worth tuning is `AGENT_IDENTITY_TOKEN_REFRESH_MARGIN`: how long before
  expiry a credential is renewed, as a Go duration such as `20m`, defaulting to
  `20m`.

## Volume context options

Only `authType` is yours to set:

> `authType`: set to `agent-identity` to enable the credential flow.

Two more are required, and neither is yours to write: the sandbox controller
supplies both.

> `sandboxId`: required. Names the internal credential to exchange
> (`<sandboxId>.token`) and derives the credential's scope.
>
> `sandboxCredProviderName` (or `credentialProviderName`, as in OSS): required.
> The provider that issues the STS credential, and therefore what the credential
> is scoped to.

Two are optional, and only about where your scripts and the credential live:

> `credentialDir`: pins the absolute path the credential files are written to.
> Unset, each mount gets its own directory under `/run/customfuse-credentials`
> and your entrypoint reads the path from `$credentialDir`, which is the normal
> case. Set it only when your client's own configuration hardcodes where it reads
> a credential from, and so cannot be told at mount time. That path is a symlink
> to the directory actually holding the files, so it cannot be one your client
> already created: rotating onto an existing directory fails the mount rather
> than overwriting it.
>
> `credentialRefreshHookKey`: names an `entrypointConfig` key projected as the
> refresh hook into a fuse pod. A sandbox has no fuse pod, so there the hook is
> whichever of two fixed paths the sidecar finds — see [Rotation](#rotation).

## What the entrypoint receives

| Variable | Value |
|----------|-------|
| `$credentialDir` | Directory holding the credential files. Set for the entrypoint **and** for the refresh hook. |
| `$authType` | `agent-identity`, so one entrypoint can serve both this and the default Secret-passthrough flow. |

The directory contains one file per field:

```
$credentialDir/AccessKeyId
$credentialDir/AccessKeySecret
$credentialDir/SecurityToken
$credentialDir/Expiration
```

One value per file rather than a finished format, because no format is neutral —
even "JSON" is not one thing: the same three values are `AccessKeySecret` and
`SecurityToken` in one ecosystem, `SecretAccessKey` and `SessionToken` in
another. Composing from separate values needs no parsing; a finished format would
force every entrypoint wanting a different one to first take it apart.

## Reading the credential

`$credentialDir` is a symlink that a rotation replaces. Resolve it **once** and
read every field from the result:

```bash
read_credential() {
    local dir
    dir=$(readlink -f "$credentialDir") || return 1
    ak=$(cat "$dir/AccessKeyId") || return 1
    sk=$(cat "$dir/AccessKeySecret") || return 1
    tk=$(cat "$dir/SecurityToken") || return 1
}

for attempt in 1 2 3; do
    read_credential && break
    [ "$attempt" = 3 ] && { echo "cannot read credential" >&2; exit 1; }
    sleep 1
done
```

Reading `$credentialDir/AccessKeyId`, `$credentialDir/AccessKeySecret` and so on
directly resolves the symlink once per field, so a rotation landing in between
yields fields from two different credentials — a combination that never existed,
with no error raised. The retry covers the opposite case: the directory a
rotation replaced is deleted shortly afterwards, so a resolve landing just before
one may find its files already gone.

## Rotation

The credential is renewed `AGENT_IDENTITY_TOKEN_REFRESH_MARGIN` before it
expires. Whether a renewal reaches your client depends on the client, and there
are only three behaviours it can have:

| If your client… | Then… |
|---|---|
| re-reads the credential files | nothing to do — ship no hook |
| must be told (its credential is already parsed into memory, or lives in its own store) | supply a hook that tells it |
| cannot accept a new credential while running | no hook helps; it needs a credential outliving the mount |

The hook is an arbitrary executable, so "tell it" can be whatever your client
supports — invoking its CLI, sending it a signal, calling its admin endpoint,
rewriting its own config. It receives `$credentialDir` plus the volume's own
attributes (`$mountpoint`, `$source`, and whatever the PersistentVolume set), so
it can name the volume it is being asked about. No credential value reaches it,
and neither does any Secret.

A hook that fails is reported: the files rotated but the client was not told,
which is a half-applied rotation and must not pass silently.

Where the hook comes from:

* **Fuse pod** — a volume that set `credentialRefreshHookKey` gets that key
  projected as `/etc/fuse-config/refresh-hook.sh`. Nothing is projected unless
  the key is named.
* **Sandbox sidecar** — `/etc/fuse-config/refresh-hook.sh` first, then
  `/refresh-hook.sh` in your image; whichever exists first runs, and if neither
  does, no hook runs, which is what a client that re-reads the files wants. One
  that exists but is not executable is reported and skipped rather than failing
  the mount.

## Example

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: customfuse-agent-identity
spec:
  capacity:
    storage: 100Gi
  accessModes:
    - ReadWriteMany
  csi:
    driver: customfuseplugin.csi.alibabacloud.com
    volumeHandle: customfuse-agent-identity
    volumeAttributes:
      fuseType: "myfuse"
      authType: "agent-identity"
      source: "<mount-source>"
      url: "<endpoint>"
```

No `nodePublishSecretRef` is required, and no attribute here names your scripts.

Runnable end to end:
[`examples/customfuse/5-agent-identity`](../examples/customfuse/5-agent-identity/)
is the smallest working sandbox setup, with both scripts baked into the image.
Its "Which case is your client?" section adapts JindoFS, JuiceFS and s3fs to the
three rotation behaviours above.
