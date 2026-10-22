# Mount a customfuse volume with agent-identity (STS credential) authentication

The customfuse driver can exchange a sandbox agent-identity token for a
short-lived, scoped STS credential instead of taking a long-lived AccessKey from
a Secret. The credential is delivered via stdin and kept fresh for the lifetime
of the mount, so no AccessKey is stored in the cluster and no credential touches
the filesystem.

Unlike the other drivers, customfuse runs a FUSE client chosen by the volume, so
it cannot know how that client wants a credential handed to it — `alinas` can
assume its client takes the STS triple as mount options, `ossfs` can assume its
client performs the exchange itself. customfuse can assume neither. It therefore
performs the exchange itself and delivers the result via stdin, leaving your
entrypoint to adapt it to your client.

## How it works

You supply the FUSE client and the script that mounts it. The driver supplies the
credential:

1. The volume sets `authType: agent-identity` and carries no Secret.
2. **Your entrypoint receives** the credential on stdin as four lines
   (AccessKeyId, AccessKeySecret, SecurityToken, Expiration).
3. Your entrypoint reads stdin and mounts.
4. **Before the credential expires**, the driver runs the *refresh hook* with
   the new credential on its stdin in the same format. See [Rotation](#rotation).
5. **When your entrypoint exits** — unmounted or crashed — the refresh loop
   stops.

A failed exchange fails the mount: there is no static credential to fall back to,
and a client started without one fails later and less clearly.

## What the driver guarantees, and what it does not

**The driver guarantees:**

* The initial credential is on your entrypoint's stdin before it runs.
* Each rotation delivers a complete credential to the hook's stdin — all four
  fields arrive atomically.
* The credential is renewed before it expires, for as long as the mount lives.
* The refresh loop stops when the entrypoint exits. The driver does not write or
  remove any files — if your scripts write credential files, cleanup is yours.
* No credential value is ever passed as an argument or an environment variable.
  Nothing lands in `/proc/<pid>/cmdline` or `/proc/<pid>/environ`.

**The driver does not:**

* Assume anything about how your client consumes a credential, or prescribe how
  it should. Delivery is stdin; everything above that is yours.
* Guarantee your client notices a rotation. A refresh hook is run after each
  rotation; what it does is up to you.
* Work around a client that cannot accept a new credential at all while running.
  Such a client needs a credential outliving the mount; no delivery mechanism
  helps.

**Your entrypoint is responsible for:**

* Reading the four lines from stdin and composing them into whatever form your
  client takes — command line arguments, a config file, environment variables,
  an ini profile.
* Telling your client about a rotation, via the refresh hook, if it cannot
  pick up the new credential on its own.

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

One is optional:

> `credentialRefreshHookKey`: names an `entrypointConfig` key projected as the
> refresh hook into a fuse pod. A sandbox has no fuse pod, so there the hook is
> whichever of two fixed paths the sidecar finds — see [Rotation](#rotation).

## What the entrypoint receives

The credential arrives on **stdin** as four lines, one field per line:

```
AccessKeyId
AccessKeySecret
SecurityToken
Expiration
```

The entrypoint also receives `$authType` (`agent-identity`) as an environment
variable, so one script can serve both this and the default Secret-passthrough
flow.

## Reading the credential

```bash
read -r ak
read -r sk
read -r token
read -r expiration
```

No files, no symlinks, no jq — just `read`. The credential never touches the
filesystem, argv or `/proc/<pid>/environ`.

## Rotation

The credential is renewed `AGENT_IDENTITY_TOKEN_REFRESH_MARGIN` before it
expires. Whether a renewal reaches your client depends on the client, and there
are only three behaviours it can have:

| If your client… | Then… |
|---|---|
| re-reads credential files written by your entrypoint/hook | nothing to do — ship a hook that writes files, no further notification needed |
| must be told (its credential is already parsed into memory, or lives in its own store) | supply a hook that tells it |
| cannot accept a new credential while running | no hook helps; it needs a credential outliving the mount |

The hook is an arbitrary executable, so "tell it" can be whatever your client
supports — invoking its CLI, sending it a signal, calling its admin endpoint,
rewriting its own config. It receives the new credential on **stdin** (same
four-line format as the entrypoint) plus the volume's own attributes
(`$mountpoint`, `$source`, and whatever the PersistentVolume set) as environment
variables, so it can name the volume it is being asked about.

A hook that fails is retried with exponential backoff. The credential was already
fetched, so retries do not cause redundant STS exchanges.

Where the hook comes from:

* **Fuse pod** — a volume that set `credentialRefreshHookKey` gets that key
  projected as `/etc/fuse-config/refresh-hook.sh`. Nothing is projected unless
  the key is named.
* **Sandbox sidecar** — `/etc/fuse-config/refresh-hook.sh` first, then
  `/refresh-hook.sh` in your image; whichever exists first runs, and if neither
  does, no hook runs. One
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
