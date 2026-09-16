# Demo 4: ConfigMap-based entrypoint
#
# The entrypoint script is NOT baked into this image — it comes from a
# ConfigMap (see configmap-production.yaml / configmap-development.yaml).
# This image only provides the FUSE client binary and mount-proxy.
#
# Complex format/mount options are hardcoded in ConfigMaps.
# PV provides only instance-specific values (source, bucket, url)
# and points to a ConfigMap via entrypointConfig.

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
#       -t <image>:<tag> examples/customfuse/4-configmap
ARG MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag>
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

FROM juicedata/mount:ce-v1.2.0

COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/

# No COPY entrypoint.sh — it is provided via ConfigMap at runtime.
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
