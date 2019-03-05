FROM centos:7.4.1708
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

COPY ossfs_1.80.3_centos7.0_x86_64.rpm /root/ossfs_1.80.3_centos7.0_x86_64.rpm
COPY nsenter /nsenter
COPY entrypoint.sh /entrypoint.sh
COPY ossplugin.csi.alibabacloud.com /bin/ossplugin.csi.alibabacloud.com
RUN chmod +x /bin/ossplugin.csi.alibabacloud.com && chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]