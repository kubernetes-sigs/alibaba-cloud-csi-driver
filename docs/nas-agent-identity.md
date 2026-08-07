# Mount a NAS volume with agent-identity (STS credential) authentication

The NAS driver can mount `alinas`/CPFS volumes using a short-lived, scoped STS
credential obtained from a sandbox agent-identity token, instead of a
long-lived AccessKey stored in a Secret or on disk. The credential is exchanged
in-memory, injected into the mount as sensitive (masked) options, and kept
fresh for the lifetime of the mount.

## How it works

1. The node driver reads the agent-identity mount options from the volume
   context and forwards them into the `alinas` mount options.
2. The mount-proxy `alinas` driver's jwtauth interceptor exchanges the sandbox
   token for a scoped STS credential (`GetResourceCredential`) over a
   TLS-verified channel.
3. The STS triple (`access_key_id` / `access_key_secret` / `security_token`) is
   passed to `mount -t alinas` as **sensitive options**, so it is masked in
   mount logs and error messages.
4. A background refresher renews the credential before it expires and pushes the
   rotated credential to the live mount via `alinas-tls-cert-refresh` (shipped
   with `aliyun-alinas-utils`). Nothing is written to disk.
5. The refresher is stopped on unmount and on driver termination.

## Prerequisite

* A working Kubernetes cluster with the CSI plugin and NAS driver enabled.
  Please refer to the [installation guide](./install.md).
* The mount-proxy environment must be configured with the agent-identity
  endpoint and token directory (and optionally a CA file):
  * `AGENT_IDENTITY_ENDPOINT` — credential provider endpoint (required).
  * `AGENT_IDENTITY_TOKEN_DIR` — directory holding `<sandboxId>.token` files
    (required to auto-resolve the token file).
  * `AGENT_IDENTITY_CERT_FILE` — CA file for the TLS channel (optional; the
    system root pool is used when unset).
* `aliyun-alinas-utils` providing `alinas-tls-cert-refresh` on the node.

## Volume context options

Set the following options in the PV `volumeAttributes` (or StorageClass
`parameters` for dynamic provisioning):

> `authType`: set to `agent-identity` to enable the credential flow.
>
> `sandboxId`: required. Identifies the sandbox token file (`<AGENT_IDENTITY_TOKEN_DIR>/<sandboxId>.token`) and derives a unique credential scope.
>
> `sandboxCredentialProviderName`: required. The credential provider name used in the exchange request.

The interceptor forces `tls` and `ram` mount options for agent-identity mounts.

## Example

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nas-agent-identity
spec:
  capacity:
    storage: 20Gi
  accessModes:
    - ReadWriteMany
  csi:
    driver: nasplugin.csi.alibabacloud.com
    volumeHandle: nas-agent-identity
    volumeAttributes:
      server: "<file-system-id>.<region>.nas.aliyuncs.com"
      path: "/"
      authType: "agent-identity"
      sandboxId: "<sandbox-id>"
      sandboxCredentialProviderName: "<credential-provider-name>"
```

No `secretRef` / AccessKey is required: the STS credential is obtained from the
sandbox agent-identity token at mount time and refreshed automatically.
