# Demo 1: Baked-in configuration
#
# All configuration (META-URL, credentials, options) is hardcoded in the image.
# The PV only needs fuseType (optional) — no source, url, otherOpts, or Secret.
#
# Use case: internal/fixed deployments where the FUSE config never changes.

# mount-proxy is not built here. It is the same for every FUSE client, so it
# ships as its own image: build it once from the customfuse-base target of
# build/mount-proxy/Dockerfile and pass it in.
#
#   docker build --target customfuse-base -f build/mount-proxy/Dockerfile \
#       -t <registry>/csi-mount-proxy:<tag> .
#
# This directory is the whole build context, so a copy of it anywhere builds the
# same image:
#
#   docker build --build-arg MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag> \
#       -t <image>:<tag> examples/customfuse/1-baked-in
ARG MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag>
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

FROM juicedata/mount:ce-v1.2.0

COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/

# Bake in all configuration as environment variables.
# The entrypoint uses these directly — no PV volumeAttributes or Secrets needed.
#
# WARNING: values declared with ENV are persisted in the image config and are
# readable by anyone who can pull the image (e.g. `docker inspect`,
# `docker history`). Switching ENV to ARG does NOT hide them — ARGs are also
# recorded in image metadata, and any RUN that consumes them bakes the value
# into a layer. Treat this demo as illustrative only; do NOT ship real
# credentials this way.
#
# For production, keep only non-sensitive defaults here (e.g. CACHE_SIZE) and hand
# the credentials over at runtime instead. See ../2-oss-compatible for the full
# pattern. Minimal sketch:
#
#   apiVersion: v1
#   kind: Secret
#   metadata:
#     name: juicefs-creds
#     namespace: default
#   type: Opaque
#   stringData:
#     ACCESS_KEY: <AK>
#     SECRET_KEY: <SK>
#   ---
#   # On the PV that mounts this image:
#   spec:
#     csi:
#       nodePublishSecretRef: { name: juicefs-creds, namespace: default }
#
# mount-proxy sets every Secret key as an environment variable of that same name in
# the fuse container, so entrypoint.sh reads them unchanged. Two parts of this are
# easy to get wrong:
#
#   * The Secret hangs off the PV, not off a pod spec. The driver creates the fuse
#     pod and puts no environment on it, and the consumer pod's environment stays in
#     the consumer pod — no `env:` entry anywhere reaches this container.
#   * The ACCESS_KEY and SECRET_KEY lines below have to be deleted rather than left
#     as placeholders. A name already present in this container's environment is one
#     mount-proxy refuses to redefine, so the baked-in value would win and the Secret
#     entry would be dropped — an empty ENV claims the name as firmly as a real one.
ENV META_URL="redis://redis-host.default.svc:6379/1"
ENV BUCKET="my-jfs-data"
ENV URL="oss-cn-hangzhou-internal.aliyuncs.com"
ENV ACCESS_KEY="<your-access-key-id>"
ENV SECRET_KEY="<your-access-key-secret>"
ENV CACHE_SIZE="1024"
ENV CAPACITY="100"

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
