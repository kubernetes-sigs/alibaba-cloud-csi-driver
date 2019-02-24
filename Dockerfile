FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI DiskPlugin"

RUN yum install -y e4fsprogs

COPY nsenter /
COPY diskplugin.csi.aliyun.com /bin/diskplugin.csi.aliyun.com
RUN chmod +x /bin/diskplugin.csi.aliyun.com
RUN chmod 755 /nsenter

ENTRYPOINT ["/bin/diskplugin.csi.aliyun.com"]
