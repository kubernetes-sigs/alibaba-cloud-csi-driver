FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

ARG ossfsVer=1.80.6.ack.1
#COPY ossfs_1.80.6_centos7.0_x86_64.rpm /root/ossfs_1.80.6_centos7.0_x86_64.rpm
COPY nsenter /nsenter
COPY entrypoint.sh /entrypoint.sh
COPY jindofs-fuse /bin/jindofs-fuse
COPY csiplugin-connector /bin/csiplugin-connector
COPY csiplugin-connector.service /bin/csiplugin-connector.service
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh && chmod +x /bin/csiplugin-connector && chmod +x /bin/csiplugin-connector.service

ENTRYPOINT ["/entrypoint.sh"]