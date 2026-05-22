---
name: init-dev-env
description: Initialize the local development environment for the alibaba-cloud-csi-driver project. Use this skill when the user wants to set up their dev environment, configure test cluster access, create image pull secrets, or write test-values.yaml. Also trigger when the user mentions "init dev", "setup dev env", "configure test cluster", or asks how to push/pull test images to ACR.
---

# Initialize Development Environment

This skill walks through setting up the local development environment for building, pushing, and deploying the CSI driver to a test ACK cluster.

## Overview of steps

1. Probe the environment (kubectl, aliyun CLI)
2. Gather ACR registry info
3. Confirm all details with the user
4. Uninstall existing CSI addon (if installed via ACK component center)
5. Write `.local/test-values.yaml`
6. Create image pull secrets

## Step 1: Probe environment

Run these two commands to gather cluster and identity info:

```bash
kubectl get cm -n kube-system ack-cluster-profile -ojsonpath='{.data}'
```

This returns JSON with `clusterid`, `uid` (Alibaba Cloud account ID), and `vpcid` among other fields.

```bash
aliyun sts GetCallerIdentity
```

This returns the current RAM identity including `AccountId`. If the `AccountId` matches the `uid` from the configmap, the aliyun CLI is properly configured for this cluster's account.

### If kubectl fails (no cluster accessible)

List available clusters from aliyun CLI for the user to select:

```bash
aliyun cs GET "/clusters" | jq '[.[] | {cluster_id, name, region_id, state}]'
```

If a suitable cluster is found, after user approval, write its kubeconfig to `~/.kube/config`:

```bash
aliyun cs GET "/k8s/<cluster-id>/user_config" | jq -r .config > ~/.kube/config
chmod 600 ~/.kube/config
```

If no cluster exists, ask the user if they want to create one. If `.local/create-cluster.sh` exists, offer to reuse it. Otherwise, gather parameters and write the script first:

List existing VPCs for the user to select:
```bash
aliyun vpc DescribeVpcs --RegionId <region> | jq '[.Vpcs.Vpc[] | {VpcId, VpcName, CidrBlock}]'
```

Parameters to gather:
- `N_NODES` (default 1)
- `ACK_REGION` (default cn-beijing)
- `ACK_ZONE` (default "cn-beijing-h cn-beijing-i cn-beijing-j", use a single zone for N_NODES=1)
- `VPC_ID` (optional, set to reuse an existing VPC; leave unset to create a new one)
- Cluster name (default `ack-csi-driver-test`)

Write `.local/create-cluster.sh` with the chosen parameters:
```bash
#!/bin/bash
set -e
source "$(git rev-parse --show-toplevel)/test/cluster/cluster_setup.sh"
VPC_ID=<vpc-id> N_NODES=<n> ACK_ZONE="<zone>" create-ack-cluster <cluster-name>
```

Then run it to create the cluster.

### If aliyun CLI fails

Check memory or other loaded context for account info. If nothing found, tell the user to install and configure aliyun CLI:
- Install: https://help.aliyun.com/cli/
- Configure: `aliyun configure` with their AccessKey ID/Secret and region

## Step 2: Gather ACR info

List registries already configured in `~/.docker/config.json`:

```bash
jq -r '.auths | keys[]' ~/.docker/config.json
```

Present these for the user to select from.

Get the cluster region:

```bash
aliyun cs GET "/clusters/<cluster-id>" | jq -r .region_id
```

If the selected registry is an `aliyuncs.com` endpoint whose region matches the cluster region, derive the VPC endpoint automatically:
- Personal edition: `registry.<region>.aliyuncs.com` → VPC: `registry-vpc.<region>.aliyuncs.com`
- Enterprise edition: `<name>-registry.<region>.cr.aliyuncs.com` → VPC: `<name>-registry-vpc.<region>.cr.aliyuncs.com`

If the registry region does NOT match the cluster region (or it's a custom domain where VPC endpoint cannot be derived), set both `registry` and `registryVPC` to the same value and warn the user that in-cluster pulls will go over the public network.

Then ask the user for their **ACR namespace** (the first path segment in image URLs, e.g., `huweiwen-test` in `registry.cn-beijing.aliyuncs.com/huweiwen-test/csi-plugin`).

## Step 3: Confirm with user

Present a summary for confirmation:
- Cluster ID
- Account UID
- Region
- ACR registry endpoint (public)
- ACR registry endpoint (VPC, for in-cluster pulls)
- ACR namespace
- Image pull secret name (convention: `<namespace>-registry`)

## Step 4: Uninstall existing CSI addon

Before deploying via helm, any existing CSI installation from the ACK component center must be removed. The bundled script `scripts/uninstall-csi-addon.sh` handles this automatically:

- Skips if CSI is already managed by Helm (debug version)
- Detects installed CSI addons (`csi-plugin`, `csi-provisioner`, `managed-csiprovisioner`) via the ACK API
- Uninstalls in dependency order: `csi-plugin` first, then `csi-provisioner` or `managed-csiprovisioner` (they are mutually exclusive)
- Waits for each removal to complete before proceeding
- Cleans up leftover resources (ack-csi-fuse namespace, stale ServiceAccounts)

```bash
scripts/uninstall-csi-addon.sh
```

Note: `managed-csiprovisioner` and `csi-provisioner` cannot coexist — only one will be present.

## Step 5: Write `.local/test-values.yaml`

Template:

```yaml
images:
  registry: <registry>
  registryVPC: <registryVPC>
  controller:
    repo: <namespace>/csi-plugin
    tag: "latest-controller"
  plugin:
    repo: <namespace>/csi-plugin
    tag: "latest"
  pluginInit:
    repo: <namespace>/csi-plugin
    tag: "latest-init"

imagePullSecrets:
  - <namespace>-registry

deploy:
  clusterID: "<cluster-id>"
  ramToken: v1  # Use v1 because it exists even when ACK csi addon is not installed
```

Replace placeholders with actual values gathered in previous steps.

## Step 6: Create image pull secrets

The bundled script `scripts/create-image-pull-secret.sh` handles creating the Kubernetes secret. It extracts the docker auth entry from local `~/.docker/config.json` and passes it directly as a `dockerconfigjson` secret (same format, no parsing needed).

The script reads registry and secret name from `.local/test-values.yaml`, so it requires no configuration. It passes through all arguments to `kubectl create secret`, enabling namespace selection.

Prerequisite: user must have run `docker login <registry>` locally first. If credentials are missing, the script tells them to do so.

Run the script to create secrets in both required namespaces:

```bash
scripts/create-image-pull-secret.sh -n kube-system
scripts/create-image-pull-secret.sh -n ack-csi-fuse
```

## Post-setup notes

After running the skill, remind the user of the build & deploy workflow with `/debug-csi`.
