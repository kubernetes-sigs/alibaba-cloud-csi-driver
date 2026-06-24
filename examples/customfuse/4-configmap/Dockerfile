# Demo 4: ConfigMap-based entrypoint
#
# The entrypoint script is NOT baked into this image — it comes from a
# ConfigMap (see configmap-production.yaml / configmap-development.yaml).
# This image only provides the FUSE client binary and mount-proxy.
#
# Complex format/mount options are hardcoded in ConfigMaps.
# PV provides only instance-specific values (source, bucket, url)
# and points to a ConfigMap via entrypointConfig.

# Stage 1: build the mount-proxy binaries from source.
# Run `docker build` from the repo root so the Go source tree is available as build context.
#   docker build -f examples/customfuse/4-configmap/Dockerfile -t <image>:<tag> .
FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.26.3 AS builder
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

# Stage 2: assemble the final image on top of the FUSE client.
# TODO: replace with actual JuiceFS image tag
FROM juicedata/mount:ce-v1.2.0

COPY --link --from=builder /out/csi-mount-proxy-server /usr/local/bin/
COPY --link --from=builder /out/csi-mount-proxy-client /usr/local/bin/

# No COPY entrypoint.sh — it is provided via ConfigMap at runtime.
ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
