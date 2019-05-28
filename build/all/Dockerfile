FROM ubuntu:18.10
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN apt upgrade -y && apt-get update && apt install -y ca-certificates file tzdata && apt install -y nfs-common jq

COPY ossfs_1.80.3_centos7.0_x86_64.rpm /root/ossfs_1.80.3_centos7.0_x86_64.rpm
COPY nsenter /nsenter
COPY entrypoint.sh /entrypoint.sh
COPY csiplugin-connector /bin/csiplugin-connector
COPY csiplugin-connector-init /bin/csiplugin-connector-init
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh && chmod +x /bin/csiplugin-connector && chmod +x /bin/csiplugin-connector-init

ENTRYPOINT ["/entrypoint.sh"]