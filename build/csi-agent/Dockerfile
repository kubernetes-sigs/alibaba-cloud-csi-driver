FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.22.9 as builder
WORKDIR /src
ARG TARGETARCH
ARG TARGETOS
RUN --mount=type=bind,target=. \
    export GOOS=$TARGETOS && \
    export GOARCH=$TARGETARCH && \
    export CGO_ENABLED=0 && \
    go build -trimpath \
        -ldflags "-X github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version.VERSION=${CSI_VERSION}" \
        -o /out/csi-agent ./cmd/csi-agent

FROM registry-cn-hangzhou.ack.aliyuncs.com/acs/csi-ossfs:v1.91.3-d9f3917-aliyun
COPY --link --from=builder /out/csi-agent /usr/local/bin/
