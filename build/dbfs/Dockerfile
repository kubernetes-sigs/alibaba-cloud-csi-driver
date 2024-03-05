FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.6.1810
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum install -y ca-certificates file tzdata nfs-utils xfsprogs e4fsprogs pciutils

COPY entrypoint.sh /entrypoint.sh
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
