FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.23.9 as builder
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
ARG OSSFS_VERSION=v1.91.7.ack.1
ARG OSSFS2_VERSION=v2.0.2beta.ack.2
ARG ALINAS_UTILS_VERSION=1.6-1.20241101165952.ce0ef4
ARG EFC_VERSION=1.6-20241028201622.a31063
RUN set -ex; \
    case "${TARGETPLATFORM}" in \
        linux/amd64) ARCH="x86_64" ;; \
        linux/arm64) ARCH="aarch_64" ;; \
        *) echo "unknown platform"; exit 1 ;; \
    esac; \
    yum install -y fuse-devel util-linux mailcap procps tini; \
    yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${OSSFS_VERSION}_centos8.0_${ARCH}.rpm; \
    if [ "${ARCH}" = "x86_64" ]; then \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-${ALINAS_UTILS_VERSION}.al7.noarch.rpm; \
        yum install -y https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-${EFC_VERSION}.release.${ARCH}.rpm; \
        yum install -y https://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs2/ossfs2_${OSSFS2_VERSION}_centos8.0_${ARCH}.rpm; \
    fi; \
    yum clean all

COPY --link --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=builder /out/csi-agent /usr/local/bin/
ENTRYPOINT [ "tini", "--", "csi-mount-proxy-server", "--driver=alinas,ossfs,ossfs2" ]

