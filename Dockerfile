FROM golang:1.16.5 as builder
WORKDIR /go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver
# Perform the build
COPY . .

ARG branch=v1.0.0
ARG version=v1.14.5
ARG commitId=$(git rev-parse --short HEAD || echo "HEAD")
ARG buildTime=$(date "+%Y-%m-%d-%H:%M:%S")

RUN CGO_ENABLED=0 go build -ldflags "-X main._BRANCH_='$branch' -X main._VERSION_='$version-$commitId' -X main._BUILDTIME_='$buildTime'" -o plugin.csi.alibabacloud.com

FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI DiskPlugin"

RUN yum install -y e4fsprogs

#COPY --from=builder /go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/nsenter /nsenter
COPY --from=builder /go/src/github.com/kubernetes-sigs/alibaba-cloud-csi-driver/plugin.csi.alibabacloud.com /bin/diskplugin.csi.alibabacloud.com
RUN chmod +x /bin/diskplugin.csi.alibabacloud.com
#RUN chmod 755 /nsenter

ENTRYPOINT ["/bin/diskplugin.csi.alibabacloud.com"]
