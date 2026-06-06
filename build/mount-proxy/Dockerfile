FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.26.4 AS builder
WORKDIR /src
ARG TARGETARCH
ARG TARGETOS
RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    export GOOS=$TARGETOS && \
    export GOARCH=$TARGETARCH && \
    export CGO_ENABLED=0 && \
    go build -o /out/csi-mount-proxy-server ./cmd/mount-proxy-server && \
    go build -o /out/csi-mount-proxy-client ./cmd/mount-proxy-client


FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update AS rpm-builder
RUN yum install -y rpm-build
RUN --mount=type=bind,source=build/mount-proxy/python3-provider.spec,target=/tmp/python3-provider.spec \
    rpmbuild -bb /tmp/python3-provider.spec


FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update AS alinux-install

#################### OSSFS 1.91 MOUNT-PROXY IMAGE ####################

FROM alinux-install AS ossfs
ARG TARGETPLATFORM
ARG OSSFS_VERSION=v1.91.10.ack.1

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

FROM alinux-install AS ossfs-1.88
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

FROM alinux-install AS ossfs2
ARG TARGETPLATFORM
ARG OSSFS_VERSION=v2.0.7.ack.1

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

FROM alinux-install AS alinas

ARG TARGETPLATFORM
ARG ALINAS_UTILS_VERSION=1.9-0.20260605142542.e1dd14.al7
ARG EFC_VERSION=1.9-20250822095129.f34347.release
# v2.1 is exclusive to a specific storage cluster version and the first to support ARM arch.
# It's temporarily unavailable for all ACK clusters,
# catering only to ARM for compatibility. An update to all CPFS storage cluster versions is required.
ARG EFC_ARM_VERSION=2.1-20260318190316.34c388.release
ARG ALINAS_UTILS_ARM_VERSION=2.1-0.20260316154701.b0a52a.generic

RUN --mount=type=bind,from=rpm-builder,source=/root/rpmbuild/RPMS/noarch,target=/tmp/rpms \
    echo 'install_weak_deps=False' >> /etc/yum.conf; \
    set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch_64" ;; \
        *) echo "unsupported platform"; exit 1 ;; \
    esac; \
    yum install -y /tmp/rpms/python3-provider-*.rpm util-linux procps tini; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.noarch.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.${ARCH}.rpm; \
    elif [ "${ARCH}" = "aarch_64" ]; then \
        yum install -y https://aliyun-alinas-eac-cn-wulanchabu.oss-cn-wulanchabu.aliyuncs.com/efc-release/2.1/aliyun-alinas-utils-${ALINAS_UTILS_ARM_VERSION}.aarch64.rpm; \
        yum install -y https://aliyun-alinas-eac-cn-wulanchabu.oss-cn-wulanchabu.aliyuncs.com/efc-release/2.1/alinas-efc-${EFC_ARM_VERSION}.aarch64.rpm; \
    fi; \
    yum clean all; \
    # python3-pyyaml is only used by nfsdclnts (NFS server diagnostic tool, irrelevant here).
    # Remove it to silence CVE scans.
    rpm -e --nodeps python3-pyyaml
# alinas-utils generates some dynamic configuration files (such as TLS keys, etc.) under /etc/aliyun/{alinas,cpfs}.
# So we will mount a hostPath to prevent the loss of these files when the container restarts.
# On the first launch, /etc/aliyun-defaults will be used to initialize the configurations under /etc/aliyun/{alinas,cpfs}.
RUN cp -r /etc/aliyun /etc/aliyun-defaults

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas" ]

#################### ALINAS MOUNT-PROXY IMAGE ####################

#################### ALL-IN-ONE MOUNT-PROXY IMAGE ####################
# This image is used in pod sidecar injected by our infra where we don't know which driver will be used.

FROM alinux-install AS aio

ARG TARGETPLATFORM
ARG OSSFS_VERSION=v1.91.10.ack.1
ARG OSSFS2_VERSION=v2.0.7.ack.1
ARG ALINAS_UTILS_VERSION=1.9-0.20260605142542.e1dd14.al7
ARG EFC_VERSION=1.9-20250822095129.f34347.release
# v2.1 is exclusive to a specific storage cluster version and the first to support ARM arch.
# It's temporarily unavailable for all ACK clusters,
# catering only to ARM for compatibility. An update to all CPFS storage cluster versions is required.
ARG EFC_ARM_VERSION=2.1-20260318190316.34c388.release
ARG ALINAS_UTILS_ARM_VERSION=2.1-0.20260316154701.b0a52a.generic

RUN --mount=type=bind,from=rpm-builder,source=/root/rpmbuild/RPMS/noarch,target=/tmp/rpms \
    echo 'install_weak_deps=False' >> /etc/yum.conf; \
    set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch_64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y /tmp/rpms/python3-provider-*.rpm fuse-devel util-linux mailcap procps tini; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${OSSFS_VERSION}_centos8.0_${ARCH}.rpm; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs2/ossfs2_${OSSFS2_VERSION}_centos8.0_${ARCH}.rpm; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.noarch.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.${ARCH}.rpm; \
    elif [ "${ARCH}" = "aarch_64" ]; then \
        yum install -y https://aliyun-alinas-eac-cn-wulanchabu.oss-cn-wulanchabu.aliyuncs.com/efc-release/2.1/aliyun-alinas-utils-${ALINAS_UTILS_ARM_VERSION}.aarch64.rpm; \
        yum install -y https://aliyun-alinas-eac-cn-wulanchabu.oss-cn-wulanchabu.aliyuncs.com/efc-release/2.1/alinas-efc-${EFC_ARM_VERSION}.aarch64.rpm; \
    fi; \
    cp -r /etc/aliyun /etc/aliyun-defaults; \
    yum clean all; \
    # python3-pyyaml is only used by nfsdclnts (NFS server diagnostic tool, irrelevant here).
    # Remove it to silence CVE scans.
    rpm -e --nodeps python3-pyyaml

COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas,ossfs,ossfs2" ]

#################### ALL-IN-ONE MOUNT-PROXY IMAGE ####################
