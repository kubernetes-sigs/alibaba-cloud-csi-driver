# Mount NAS with RRSA (per-Pod STS credentials)

With `authType: rrsa`, the NAS driver mounts an access point using a short-lived
STS credential obtained with the Pod's own ServiceAccount token, instead of an
AccessKey stored in a Secret. The credential is renewed while the volume stays
mounted, and is never written to a Secret or to the PV.

Only [access points](https://www.alibabacloud.com/help/en/nas/user-guide/access-point-overview/)
with the RAM policy enabled check the credential, including AgenticFS access
points. Ordinary NAS mount targets are controlled by permission groups and ignore
it.

## How it works

1. kubelet requests a ServiceAccount token for the Pod, as configured by the
   NAS CSIDriver's `tokenRequests`, and passes it to `NodePublishVolume`.
2. csi-plugin calls `AssumeRoleWithOIDC` with that token and receives an STS
   credential for the configured RAM role.
3. The mount broker mounts the access point with that credential.
4. kubelet calls `NodePublishVolume` again about once a minute
   (`requiresRepublish`). From 20 minutes before the credential expires,
   csi-plugin exchanges a new one and the broker installs it on the live mount.

## Prerequisites

* [RRSA is enabled](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/user-guide/use-rrsa-to-authorize-pods-to-access-different-cloud-services)
  on the ACK cluster, which creates the cluster's OIDC identity provider
  `ack-rrsa-<cluster-id>` in RAM. The ACK webhook and SDK setup on that page are
  not needed: csi-plugin performs the exchange itself.
* NAS mounts go through the mount broker: the `AlinasMountProxy` feature gate
  is enabled on csi-plugin, and the broker (cnfs-nas-daemon on ACK) supports
  credential refresh. Otherwise an `rrsa` volume fails with
  `FailedPrecondition` before anything is mounted.
* A NAS access point with the
  [RAM policy enabled](https://www.alibabacloud.com/help/en/nas/user-guide/management-access-point),
  reachable from the cluster VPC.

## Step 1: Create the RAM role

Create a RAM role for the cluster's OIDC identity provider, as described in
[Use RRSA for pod-level access control](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/user-guide/use-rrsa-to-authorize-pods-to-access-different-cloud-services),
and trust the ServiceAccount of the Pods that mount the volume:

```json
{
  "Version": "1",
  "Statement": [{
    "Effect": "Allow",
    "Action": "sts:AssumeRole",
    "Principal": {"Federated": ["acs:ram::<account-id>:oidc-provider/ack-rrsa-<cluster-id>"]},
    "Condition": {"StringEquals": {
      "oidc:aud": "sts.aliyuncs.com",
      "oidc:sub": "system:serviceaccount:<namespace>:<serviceaccount>"
    }}
  }]
}
```

Grant it the access point client permissions described in
[Configure an access point policy](https://www.alibabacloud.com/help/en/nas/user-guide/management-access-point),
scoped to the access point:

```json
{
  "Version": "1",
  "Statement": [{
    "Effect": "Allow",
    "Action": ["nas:ClientMount", "nas:ClientWrite", "nas:ClientRootAccess"],
    "Resource": "acs:nas:<region>:<account-id>:filesystem/<file-system-id>",
    "Condition": {"StringEquals": {
      "nas:AccessPointArn": "acs:nas:<region>:<account-id>:accesspoint/<access-point-id>"
    }}
  }]
}
```

* Without `nas:ClientWrite`, the server rejects writes.
* Without `nas:ClientRootAccess`, the server maps root to nobody.

## Step 2: Configure the driver

Helm values:

```yaml
deploy:
  featureGates: "AlinasMountProxy=true"
csi:
  nas:
    tokenRequests:
      - audience: sts.aliyuncs.com
    # Optional. Credential lifetime; empty uses the STS default (1h).
    rrsaDuration: 1h
```

`tokenRequests` also enables `requiresRepublish` on the NAS CSIDriver. On
Kubernetes 1.35+ the chart also sets `serviceAccountTokenInSecrets`, so the
token is delivered as a secret and stripped from request logs. When rendering
with `helm template` for an older cluster, pass `--kube-version`.

`rrsaDuration` must be within what
[AssumeRoleWithOIDC](https://www.alibabacloud.com/help/en/ram/developer-reference/api-sts-2015-04-01-assumerolewithoidc)
accepts: at least 15 minutes, and no longer than the role's
[maximum session duration](https://www.alibabacloud.com/help/en/ram/user-guide/specify-the-maximum-session-duration-for-a-ram-role). A lifetime of 20 minutes or
less is renewed on every republish (about once a minute), and each renewal
costs one `AssumeRoleWithOIDC` call.

## Step 3: Create the PV, PVC and Pod

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nas-rrsa
spec:
  capacity:
    storage: 10Gi
  accessModes: [ReadWriteMany]
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: nasplugin.csi.alibabacloud.com
    volumeHandle: nas-rrsa
    volumeAttributes:
      server: "<access-point-id>.<file-system-id>-<suffix>.<region>.nas.aliyuncs.com"
      path: "/"
      mountProtocol: alinas
      options: "tls,vers=3,ram"
      authType: rrsa
      roleName: "<role-name>"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nas-rrsa
  namespace: <namespace>
spec:
  accessModes: [ReadWriteMany]
  storageClassName: ""
  volumeName: nas-rrsa
  resources:
    requests:
      storage: 10Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: nas-rrsa
  namespace: <namespace>
spec:
  serviceAccountName: <serviceaccount>   # must match oidc:sub in the trust policy
  containers:
  - name: app
    image: busybox
    command: [sleep, infinity]
    volumeMounts:
    - name: data
      mountPath: /data
  volumes:
  - name: data
    persistentVolumeClaim:
      claimName: nas-rrsa
```

Volume attributes:

| Attribute | Description |
|-----------|-------------|
| `authType` | `rrsa`. |
| `roleName` | A RAM role in the cluster's account. The role ARN and OIDC provider ARN are derived from the account ID and cluster ID. |
| `roleArn`, `oidcProviderArn` | Instead of `roleName`, e.g. for a role in another account. Both are required. |
| `rrsaEndpoint`, `rrsaAudience` | Only for a credential broker, see below. |

`options` must include `tls` and `ram`, which RAM authentication requires.
`authType` is not supported for losetup volumes.

### Dynamic provisioning

The same attributes can be set as StorageClass `parameters`. When `authType` is
`rrsa`, they are copied into every PV the class provisions, and are checked
before any cloud resource is created, so a misconfigured class fails at
provisioning instead of at mount time.

This applies to AgenticFS (`volumeAs: Agentic`), and to access points created
with `volumeAs: accesspoint` and `accesspointEnableRam: "true"`. For the
latter, add `ram` to `options`; the driver adds `tls` to access point mounts
itself.

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: agenticfs-rrsa
provisioner: nasplugin.csi.alibabacloud.com
parameters:
  volumeAs: Agentic
  fileSystemId: "<file-system-id>"
  zoneId: "<zone-id>"
  vpcId: "<vpc-id>"
  vSwitchId: "<vswitch-id>"
  authType: rrsa
  roleName: "<role-name>"
reclaimPolicy: Delete
```

Each provisioned volume gets a new access point, which a policy scoped to fixed
access point ARNs, as in Step 1, cannot name in advance. With plain RRSA, scope
the role to the filesystem instead by dropping the `nas:AccessPointArn`
condition, so any Pod the role trusts can mount any access point of that
filesystem. A credential broker can narrow this per Pod.

## Using a credential broker

With plain RRSA, every Pod using the role gets the role's full permissions. To
narrow them per Pod, for example to only the access points the Pod mounts and
write access only where it mounts read-write, send the token to a credential
broker instead of STS. A broker is an HTTPS service in the cluster that:

1. Accepts the same request as STS
   [AssumeRoleWithOIDC](https://www.alibabacloud.com/help/en/ram/developer-reference/api-sts-2015-04-01-assumerolewithoidc):
   a form-encoded `POST` to
   the endpoint URL with `Action`, `Version`, `Timestamp`, `Format`, `RoleArn`,
   `OIDCProviderArn`, `OIDCToken`, `RoleSessionName`, and `DurationSeconds` if
   `rrsaDuration` is set.
2. Verifies `OIDCToken` against the cluster's issuer and its own audience. The
   token identifies the Pod (name and UID), so the broker can look up the Pod's
   volumes.
3. Obtains the credential under its own identity, typically `AssumeRoleWithOIDC`
   with its own ServiceAccount token plus a session `Policy` derived from the
   Pod's volumes. The role's trust policy then names the broker's ServiceAccount,
   not the Pod's.
4. Responds like STS: HTTP 200 with
   `{"Credentials": {"AccessKeyId", "AccessKeySecret", "SecurityToken", "Expiration"}}`,
   or a non-200 status with `{"Code", "Message"}`, which fails the mount with
   that message.

Configuration for a broker:

* Volume attributes:
  * `rrsaEndpoint`: the broker URL, which must use `https`. A Service name such
    as `https://<service>.<namespace>.svc` works even though csi-plugin uses the
    host's DNS: csi-plugin reads the ClusterIP from the API server and still
    verifies the certificate against the Service name.
  * `rrsaAudience`: the audience the broker accepts. Use one other than
    `sts.aliyuncs.com`, so the Pod's token cannot be used with STS directly.
* Helm: list that audience in `csi.nas.tokenRequests` instead of, or in addition
  to, `sts.aliyuncs.com`.
* CA: create a Secret `oidc-proxy-ca` with key `ca.crt` in the csi-plugin
  namespace. The chart mounts it, and points `NAS_RRSA_CA_FILE` at it, whenever
  `tokenRequests` contains an audience other than `sts.aliyuncs.com`; until the
  Secret exists, `rrsa` mounts fail naming the missing file. The CA is added to
  the system roots, and a rotated CA takes effect within about 5 minutes of kubelet
  updating the mounted file.

```shell
kubectl -n kube-system create secret generic oidc-proxy-ca --from-file=ca.crt=<ca.pem>
```
