# Demo 1: Baked-in configuration
#
# All configuration (META-URL, credentials, options) is hardcoded in the image.
# The PV only needs fuseType (optional) — no source, url, otherOpts, or Secret.
#
# Use case: internal/fixed deployments where the FUSE config never changes.

# Stage 1: build the mount-proxy binaries from source.
# Run `docker build` from the repo root so the Go source tree is available as build context.
#   docker build -f examples/customfuse/1-baked-in/Dockerfile -t <image>:<tag> .
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

# Bake in all configuration as environment variables.
# The entrypoint uses these directly — no PV volumeAttributes or Secrets needed.
#
# WARNING: values declared with ENV are persisted in the image config and are
# readable by anyone who can pull the image (e.g. `docker inspect`,
# `docker history`). Switching ENV to ARG does NOT hide them — ARGs are also
# recorded in image metadata, and any RUN that consumes them bakes the value
# into a layer. Treat this demo as illustrative only; do NOT ship real
# credentials this way.
#
# For production, keep only non-sensitive defaults here (e.g. CACHE_SIZE) and
# inject credentials at runtime via Kubernetes Secrets. See ../2-oss-compatible for
# the full pattern. Minimal sketch:
#
#   apiVersion: v1
#   kind: Secret
#   metadata:
#     name: oss-credentials
#   type: Opaque
#   stringData:
#     access-key: <AK>
#     secret-key: <SK>
#   ---
#   # In the Pod / DaemonSet spec consuming this image:
#   env:
#   - name: ACCESS_KEY
#     valueFrom:
#       secretKeyRef: { name: oss-credentials, key: access-key }
#   - name: SECRET_KEY
#     valueFrom:
#       secretKeyRef: { name: oss-credentials, key: secret-key }
#
# Kubernetes env entries override Dockerfile ENV at container start, so the
# same entrypoint.sh works unchanged.
ENV META_URL="redis://:mypassword@redis-host.default.svc:6379/1"
ENV BUCKET="my-jfs-data"
ENV URL="oss-cn-hangzhou-internal.aliyuncs.com"
ENV ACCESS_KEY="<your-access-key-id>"
ENV SECRET_KEY="<your-access-key-secret>"
ENV CACHE_SIZE="1024"

COPY examples/customfuse/1-baked-in/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["csi-mount-proxy-server", "--driver=customfuse"]
