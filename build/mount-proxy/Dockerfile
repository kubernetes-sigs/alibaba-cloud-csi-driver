FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.24.6 as builder
WORKDIR /src
ARG TARGETARCH
ARG TARGETOS
RUN --mount=type=bind,target=. \
    export GOOS=$TARGETOS && \
    export GOARCH=$TARGETARCH && \
    export CGO_ENABLED=0 && \
    go build -o /out/csi-mount-proxy-server ./cmd/mount-proxy-server && \
    go build -o /out/csi-mount-proxy-client ./cmd/mount-proxy-client

#################### OSSFS 1.91 MOUNT-PROXY IMAGE ####################

FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update as ossfs
ARG TARGETPLATFORM
ARG OSSFS_VERSION

# install ossfs
RUN set -ex; \
    case "$TARGETPLATFORM" in \
        linux/amd64) OSSFS_ARCH="x86_64" ;; \
        linux/arm64) OSSFS_ARCH="aarch_64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${OSSFS_VERSION}_centos8.0_${OSSFS_ARCH}.rpm; \
    yum install -y fuse-devel util-linux mailcap procps; \
    yum clean all

RUN mkdir /csi && mv /etc/mime.types /csi/mime.types
RUN ln -s /host/etc/mime.types /etc/mime.types

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT ["csi-mount-proxy-server", "--driver=ossfs"]

#################### OSSFS 1.91 MOUNT-PROXY IMAGE ####################

#################### OSSFS 1.88 MOUNT-PROXY IMAGE ####################

FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update as ossfs-1.88
ARG TARGETPLATFORM

# install ossfs
RUN set -ex; \
    case "$TARGETPLATFORM" in \
        linux/amd64) \
            yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_1.88.4_centos8.0_x86_64.rpm; \
            ;; \
        linux/arm64) \
            yum install -y \
                https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/fuse-libs-2.9.2-11.el7.aarch64.rpm \
                https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/fuse-2.9.2-11.el7.aarch64.rpm; \
            curl -sSL -o /usr/lib64/libfuse.so.2 https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/libfuse.so.2; \
            curl -sSL -o /usr/local/bin/ossfs https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/multi-private/arm64-ossfs-v1.80.6/ossfs-8u; \
            chmod +x /usr/local/bin/ossfs; \
            ;; \
        *) \
            echo "unknown platform"; \
            exit 1; \
            ;; \
    esac; \
    yum install -y util-linux mailcap procps && yum clean all

RUN mkdir /csi && mv /etc/mime.types /csi/mime.types
RUN ln -s /host/etc/mime.types /etc/mime.types

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT ["csi-mount-proxy-server", "--driver=ossfs"]

#################### OSSFS 1.88 MOUNT-PROXY IMAGE ####################

#################### OSSFS2 MOUNT-PROXY IMAGE ####################

FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update as ossfs2
ARG TARGETPLATFORM
ARG OSSFS_VERSION

# install ossfs
# only supports x86_64
RUN set -ex; \
    case "$TARGETPLATFORM" in \
        linux/amd64) OSSFS_ARCH="x86_64" ;; \
        linux/arm64) OSSFS_ARCH="aarch_64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs2/ossfs2_${OSSFS_VERSION}_centos8.0_${OSSFS_ARCH}.rpm; \
    yum install -y fuse3-devel util-linux mailcap procps; \
    yum clean all

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT ["csi-mount-proxy-server", "--driver=ossfs2"]

#################### OSSFS2 MOUNT-PROXY IMAGE ####################

#################### ALINAS MOUNT-PROXY IMAGE ####################

FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update as alinas

ARG TARGETPLATFORM
ARG ALINAS_UTILS_VERSION=2.1-0.20250911190709.ea7647
ARG EFC_VERSION=2.1-20250910201826.53da93
RUN set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch64" ;; \
        *) echo "unsupported platform"; exit 1 ;; \
    esac; \
    yum install -y util-linux procps tini; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.generic.x86_64.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.release.${ARCH}.rpm; \
    elif [ "${ARCH}" = "aarch64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.generic.aarch64.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.release.${ARCH}.rpm; \
    fi; \
    yum clean all

# alinas-utils generates some dynamic configuration files (such as TLS keys, etc.) under /etc/aliyun/{alinas,cpfs}.
# So we will mount a hostPath to prevent the loss of these files when the container restarts.
# On the first launch, /etc/aliyun-defaults will be used to initialize the configurations under /etc/aliyun/{alinas,cpfs}.
RUN cp -r /etc/aliyun /etc/aliyun-defaults

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas" ]

#################### ALINAS MOUNT-PROXY IMAGE ####################
