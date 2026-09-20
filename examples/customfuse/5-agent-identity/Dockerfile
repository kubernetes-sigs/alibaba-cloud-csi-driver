# Demo 5: self-built customfuse image for an ACS sandbox, scripts baked in
#
# This image becomes the customfuse agent sidecar of your sandbox pod — the
# container the sandbox names `csi-customfuse-agent-sidecar`. It has to carry
# three things:
#
#   1. csi-mount-proxy-server as ENTRYPOINT — the sidecar *is* the mount-proxy
#   2. your FUSE client binary
#   3. /entrypoint.sh, which mount-proxy executes to perform the mount
#
# and for agent-identity, also /refresh-hook.sh. Both are fixed paths. This
# sidecar has no fuse pod, so nothing projects a ConfigMap of this volume's
# choosing into it; what is in the image is what this volume runs.
#
# Nothing about the credential is baked in, and nothing can be: an agent-identity
# credential is scoped and short-lived, exchanged per mount from the sandbox
# token. That is what makes this image reusable across sandboxes and tenants — it
# holds no secret at all.

# mount-proxy is not built here, for the same reason it is not built in
# ../1-baked-in: it is the same for every FUSE client, so it ships as its own
# image and you pass it in.
#
# This directory is the whole build context, so a copy of it anywhere builds the
# same image:
#
#   docker build --build-arg MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag> \
#       -t <image>:<tag> examples/customfuse/5-agent-identity
#
# Where that image comes from, and why building it yourself needs this
# repository's own build environment, is in ../README.md. Whichever one you name
# has to carry the mount-proxy of the CSI version running in your cluster: this
# sidecar and the plugin speak the mount-proxy protocol to each other, and the
# two ends have to match.
ARG MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag>
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

# JuiceFS here only because the mount flow needs something concrete; the driver
# neither knows nor cares which client this is.
FROM juicedata/mount:ce-v1.2.0

COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/

COPY entrypoint.sh /entrypoint.sh
COPY refresh-hook.sh /refresh-hook.sh
RUN chmod +x /entrypoint.sh /refresh-hook.sh

# The socket comes from the injected sidecar args
# (--socket=/run/cnfs/customfuse-mounter.sock), so it is not fixed here.
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
