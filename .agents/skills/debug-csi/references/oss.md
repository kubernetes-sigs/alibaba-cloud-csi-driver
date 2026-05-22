# OSS / OSSFS Debugging

OSSFS mounts are served by fuse pods managed by the csi-provisioner.

## Set up OSSFS debug environment

Run the setup script to configure everything in one shot (reads registry, namespace, and secret name from `.local/test-values.yaml`):

```bash
scripts/setup-oss-debug.sh [image-tag]
```

This does three things:
1. Sets the OSSFS and OSSFS2 image tags in the `csi-plugin` configmap (defaults to `latest`)
2. Patches `IMAGE_REPOSITORY_PREFIX` on csi-provisioner
3. Sets image pull secrets on both service accounts (`default` and `csi-fuse-ossfs`) in `ack-csi-fuse`

Rollout restart csi-provisioner if the patches report no change.

## Debugging fuse pods

List fuse pods (one per volume):

```bash
kubectl -n ack-csi-fuse get pods
kubectl -n ack-csi-fuse logs <pod-name>
```

## Restarting fuse pods

The fuse pod lifecycle is tied to ControllerPublish/ControllerUnpublish. To recreate a fuse pod with a new image:

1. Delete the consumer pod
2. Wait for ControllerUnpublish to clean up the fuse pod
3. Recreate the consumer pod

Simply deleting the fuse pod is not enough — it will not be recreated automatically.
