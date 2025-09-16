FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.24.6 as builder
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
ARG OSSFS_VERSION=v1.91.8.ack.2
ARG OSSFS2_VERSION=v2.0.2beta.ack.2
ARG ALINAS_UTILS_VERSION=2.1-0.20250911190709.ea7647
ARG EFC_VERSION=2.1-20250910201826.53da93
RUN set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y fuse-devel util-linux mailcap procps tini; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${OSSFS_VERSION}_centos8.0_${ARCH}.rpm; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.generic.${ARCH}.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.release.${ARCH}.rpm; \
        yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs2/ossfs2_${OSSFS2_VERSION}_centos8.0_${ARCH}.rpm; \
    elif [ "${ARCH}" = "aarch64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.generic.${ARCH}.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.release.${ARCH}.rpm; \
    fi; \
    yum clean all

COPY --link --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=builder /out/csi-agent /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas,ossfs,ossfs2" ]

