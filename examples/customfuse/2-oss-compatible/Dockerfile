# Demo 2: OSS-compatible configuration
#
# The image contains only the FUSE client and mount-proxy.
# Configuration comes from PV volumeAttributes and Kubernetes Secret.

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
#       -t <image>:<tag> examples/customfuse/2-oss-compatible
ARG MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag>
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

FROM juicedata/mount:ce-v1.2.0

COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
