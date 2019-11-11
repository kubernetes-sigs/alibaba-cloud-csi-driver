FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.4.1708-aliyun
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN mkdir -p /acs
COPY nsenter /acs/nsenter
COPY kmod-cpfs-client-2.10.8-202-3.10.0-957.5.1.el7.x86_64.rpm /acs/kmod-cpfs-client-2.10.8-202-3.10.0-957.5.1.el7.x86_64.rpm
COPY cpfs-client-2.10.8-202-3.10.0-957.5.1.el7.x86_64.rpm /acs/cpfs-client-2.10.8-202-3.10.0-957.5.1.el7.x86_64.rpm
COPY kmod-cpfs-client-2.10.8-202-3.10.0-957.21.3.el7.x86_64.rpm /acs/kmod-cpfs-client-2.10.8-202-3.10.0-957.21.3.el7.x86_64.rpm
COPY cpfs-client-2.10.8-202-3.10.0-957.21.3.el7.x86_64.rpm /acs/cpfs-client-2.10.8-202-3.10.0-957.21.3.el7.x86_64.rpm
COPY mount.lustre /usr/sbin/mount.lustre
COPY lctl /usr/sbin/lctl
COPY entrypoint.sh /acs/entrypoint.sh
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com

RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /acs/entrypoint.sh && chmod +x /usr/sbin/mount.lustre && chmod +x /usr/sbin/lctl

ENTRYPOINT ["/acs/entrypoint.sh"]
