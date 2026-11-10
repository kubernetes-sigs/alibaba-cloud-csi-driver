# Demo 6: agent-memory-fuse (non-OSS FUSE with agent-identity)
#
# Mounts a Tablestore Agent Memory file store using agent-memory-fuse,
# with STS credentials from the agent-identity flow.
#
# mount-proxy is not built here. It is the same for every FUSE client, so it
# ships as its own image: build it once from the customfuse-base target of
# build/mount-proxy/Dockerfile and pass it in.
#
#   docker build --target customfuse-base -f build/mount-proxy/Dockerfile \
#       -t <registry>/csi-mount-proxy:<tag> .
#
# Then build this demo:
#
#   docker build --build-arg MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag> \
#       -t <image>:<tag> examples/customfuse/6-agent-memory-fuse
ARG MOUNT_PROXY_IMAGE=<registry>/csi-mount-proxy:<tag>
FROM ${MOUNT_PROXY_IMAGE} AS mount-proxy

# agent-memory-fuse requires FUSE support (fusermount3) and /dev/fuse access.
FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update

RUN yum install -y fuse3 && yum clean all

COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=mount-proxy /usr/local/bin/csi-mount-proxy-client /usr/local/bin/

# Add the agent-memory-fuse binary.
# Download from Alibaba Cloud — see the Tablestore Agent Memory documentation.
COPY agent-memory-fuse /usr/local/bin/agent-memory-fuse
RUN chmod +x /usr/local/bin/agent-memory-fuse

COPY entrypoint.sh /entrypoint.sh
COPY refresh-hook.sh /refresh-hook.sh
RUN chmod +x /entrypoint.sh /refresh-hook.sh

ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
