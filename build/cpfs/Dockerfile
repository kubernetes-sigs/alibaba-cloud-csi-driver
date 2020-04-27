FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.4.1708-aliyun
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum install wget -y
RUN mkdir -p /acs
COPY nsenter /nsenter
COPY cpfs-client-dkms-1.0.0-204.el7.noarch.rpm /acs/cpfs-client-dkms-1.0.0-204.el7.noarch.rpm
COPY cpfs-client-2.10.8-202.el7.x86_64.rpm /acs/cpfs-client-2.10.8-202.el7.x86_64.rpm
COPY mount.lustre /usr/sbin/mount.lustre
COPY lctl /usr/sbin/lctl
COPY entrypoint.sh /acs/entrypoint.sh
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com

RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /acs/entrypoint.sh && chmod +x /usr/sbin/mount.lustre && chmod +x /usr/sbin/lctl

ENTRYPOINT ["/acs/entrypoint.sh"]
