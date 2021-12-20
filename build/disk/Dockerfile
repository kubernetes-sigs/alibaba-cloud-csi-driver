FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum install -y ca-certificates file tzdata nfs-utils xfsprogs e4fsprogs nc pciutils

COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com

ENTRYPOINT ["/bin/plugin.csi.alibabacloud.com"]
