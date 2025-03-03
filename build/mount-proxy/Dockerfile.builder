FROM --platform=$BUILDPLATFORM registry-cn-hangzhou.ack.aliyuncs.com/dev/golang:1.22.9 as builder
WORKDIR /src
ARG TARGETARCH
ARG TARGETOS
RUN --mount=type=bind,target=. \
    export GOOS=$TARGETOS && \
    export GOARCH=$TARGETARCH && \
    export CGO_ENABLED=0 && \
    go build -o /out/csi-mount-proxy-server ./cmd/mount-proxy-server && \
    go build -o /out/csi-mount-proxy-client ./cmd/mount-proxy-client

FROM --platform=$TARGETPLATFORM registry.cn-hangzhou.aliyuncs.com/acs/alinux:3-update as base
RUN sudo yum install -y automake gcc-c++ git libcurl-devel libxml2-devel \
    fuse-devel make openssl-devel rpm-build rubygems ruby ruby-devel unzip
RUN ruby -v 
RUN gem -v  
RUN gem install dotenv -v 2.8.1 -N \
    && gem install clamp -v 1.0.1 -N \
    && gem install mustache -v 0.99.8 -N \
    && gem install cabin insist stud arr-pm backports rexml -N \
    && gem install pleaserun --ignore-dependencies -N \
    && gem install fpm --ignore-dependencies -N
RUN fpm -v
RUN sudo -v && curl https://gosspublic.alicdn.com/ossutil/install.sh | sudo bash
RUN ossutil -v
COPY --link --from=builder /out/csi-mount-proxy* /usr/local/bin/