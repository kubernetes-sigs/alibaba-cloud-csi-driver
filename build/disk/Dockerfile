FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum install -y e4fsprogs

COPY csi-diskplugin /bin/csi-diskplugin
RUN chmod +x /bin/csi-diskplugin

ENTRYPOINT ["/bin/csi-diskplugin"]