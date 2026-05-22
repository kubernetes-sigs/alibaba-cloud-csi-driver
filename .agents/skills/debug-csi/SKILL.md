---
name: debug-csi
description: Build, deploy, and debug the alibaba-cloud-csi-driver on a test cluster. Use this skill when the user wants to build images, deploy via helm, check logs, restart components, or troubleshoot the CSI plugin/provisioner. Also trigger for "deploy csi", "check csi logs", "restart csi", "debug csi", "build and push", or any question about how to iterate on this driver in a dev cluster.
---

# Debug CSI Driver

This skill covers the build-deploy-debug loop for the CSI driver on an ACK test cluster.

Prerequisite: `.local/test-values.yaml` must exist. If it doesn't (or the cluster it refers to is unreachable), run `/init-dev-env` first.

## Build and deploy

### Uninstall component-center version (if present)

Before deploying via helm, check whether the CSI driver was installed via the ACK component center:

```bash
helm list -n kube-system
```

If there's a `csi` release, it's already a debug version — skip uninstall. Otherwise run:

```bash
../../init-dev-env/scripts/uninstall-csi-addon.sh
```

### Build and push images

```bash
PLATFORM=linux/amd64 IMAGE_REPO=<registry>/<namespace> IMAGE_TAG=latest ./build/build-all-multi.sh
```
Refer to `.local/test-values.yaml` for the registry and namespace.
`IMAGE_TAG` defaults to `latest`; set it to a custom tag (e.g. `v1.35.3-dev`) if needed.
The script builds and pushes all images: `csi-plugin`, `csi-plugin:init`, `csi-plugin:controller`, `csi-ossfs`, `csi-ossfs2`, `csi-alinas`, and their variants.

### Deploy via helm

```bash
helm upgrade -f .local/test-values.yaml -n kube-system ack-csi-plugin ./deploy/charts/alibaba-cloud-csi-driver
# One-time: set imagePullPolicy to Always
scripts/set-image-pull-policy.sh
```

If only images changed (no chart changes), rolling-restart instead of helm upgrade:
```bash
kubectl -n kube-system rollout restart deploy/csi-provisioner
kubectl -n kube-system rollout restart daemonset/csi-plugin
```

## Debugging and logs

**csi-plugin** (daemonset, runs on each node, handles NodePublish/NodeUnpublish):
```bash
kubectl -n kube-system logs -l app=csi-plugin -c csi-plugin --tail=100
```

**csi-provisioner** (deployment, handles ControllerPublish/ControllerUnpublish):
```bash
kubectl -n kube-system logs -l app=csi-provisioner -c csi-provisioner --tail=100
```

## Restarting components

Patching the pod template (env, imagePullPolicy, etc.) automatically triggers a rolling update.
`rollout restart` is only needed when the patch reports "no change" (e.g. image tag is the same).

## Test pod images

Use `mirrors-ssl.aliyuncs.com/busybox` or `mirrors-ssl.aliyuncs.com/debian` as test pod images.
Also use these for ephemeral containers and `kubectl debug` commands.
Docker Hub and other registries outside China are unreliable due to network issues.
`mirrors-ssl.aliyuncs.com` mirrors all popular official Docker Hub images.

## OSS / OSSFS

For OSSFS-related debugging (fuse pod image config, fuse pod logs, fuse pod lifecycle), read `references/oss.md`.
