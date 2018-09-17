FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI DiskPlugin"

RUN yum install -y e4fsprogs

COPY nsenter /
COPY csi-diskplugin /bin/csi-diskplugin
RUN chmod +x /bin/csi-diskplugin
RUN chmod 755 /nsenter

ENTRYPOINT ["/bin/csi-diskplugin"]