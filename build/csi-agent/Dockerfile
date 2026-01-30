FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.25.5 as builder
WORKDIR /src
ARG TARGETARCH
ARG TARGETOS
RUN --mount=type=bind,target=. \
    export GOOS=$TARGETOS && \
    export GOARCH=$TARGETARCH && \
    export CGO_ENABLED=0 && \
    go build -trimpath \
        -ldflags "-X github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version.VERSION=${CSI_VERSION}" \
        -o /out/csi-agent ./cmd/csi-agent && \
    go build -trimpath -o /out/csi-mount-proxy-server ./cmd/mount-proxy-server


FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update

ARG TARGETPLATFORM
# The versions specified here are for RPM packages, which may not necessarily align with the image version
ARG OSSFS_VERSION=v1.91.9.ack.1
ARG OSSFS2_VERSION=v2.0.5.ack.1
ARG ALINAS_UTILS_VERSION=1.9-0.20260130140828.ea13c0.al7
ARG EFC_VERSION=1.9-20250822095129.f34347.release
# v2.1 is exclusive to a specific storage cluster version and the first to support ARM arch.
# It's temporarily unavailable for all ACK clusters,
# catering only to ARM for compatibility. An update to all CPFS storage cluster versions is required.
ARG EFC_ARM_VERSION=2.1-20251105164215.8422bf.release
ARG ALINAS_UTILS_ARM_VERSION=2.1-0.20260114190446.d3e4f6.generic
RUN set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch_64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y fuse-devel util-linux mailcap procps python38 tini; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${OSSFS_VERSION}_centos8.0_${ARCH}.rpm; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs2/ossfs2_${OSSFS2_VERSION}_centos8.0_${ARCH}.rpm; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.noarch.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.${ARCH}.rpm; \
    elif [ "${ARCH}" = "aarch_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_ARM_VERSION}.aarch64.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_ARM_VERSION}.aarch64.rpm; \
    fi; \
    cp -r /etc/aliyun /etc/aliyun-defaults; \
    yum clean all; \
    update-alternatives --set python3 /usr/bin/python3.8


COPY --link --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=builder /out/csi-agent /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas,ossfs,ossfs2" ]