# The mock FUSE client that run-runc.sh and run-sandbox.sh drive. See README.md.
#
# It stands in for a customer FUSE client: it reports the environment the driver
# delivered, then mounts a tmpfs, which is the only thing mount-proxy waits for.
# util-linux supplies mount(8) for that. ossutil is here so the mock can prove an
# agent-identity credential is usable, not merely that the files arrived.
#
# mount-proxy is built from this tree rather than copied out of a published
# image, so the mock cannot end up exercising a mount-proxy older than the driver
# under test.
#
# Build from the repository root:
#   docker build -f test/customfuse/Dockerfile -t <registry>/customfuse-mock:<tag> .

FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.26.6 AS builder
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

FROM registry-cn-hangzhou.ack.aliyuncs.com/dev/alinux:3-update

RUN yum install -y util-linux && yum clean all

COPY --link --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=builder /out/csi-mount-proxy-client /usr/local/bin/
COPY test/customfuse/mock-entrypoint.sh /entrypoint.sh
COPY test/oss/ossutils-linux /usr/local/bin/ossutil
RUN chmod +x /entrypoint.sh /usr/local/bin/ossutil
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
