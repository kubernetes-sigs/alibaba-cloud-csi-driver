FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.4.1708-aliyun
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN mkdir -p /acs && yum install -y make gcc libyaml-devel libtool zlib-devel glibc-headers dkms expect
COPY nsenter /acs/nsenter
COPY entrypoint.sh /acs/entrypoint.sh
COPY epel.repo /etc/yum.repos.d/
COPY cpfs-client-2.10.8-202.el7.x86_64.rpm /acs/cpfs-client-2.10.8-202.el7.x86_64.rpm
COPY cpfs-client-dkms-1.0.0-202.el7.noarch.rpm /acs/cpfs-client-dkms-1.0.0-202.el7.noarch.rpm
COPY kernel-devel-3.10.0-693.2.2.el7.x86_64.rpm /acs/kernel-devel-3.10.0-693.2.2.el7.x86_64.rpm
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com

RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /acs/entrypoint.sh && yum install -y /acs/kernel-devel-3.10.0-693.2.2.el7.x86_64.rpm && yum install -y /acs/cpfs-client-dkms-1.0.0-202.el7.noarch.rpm && yum install -y /acs/cpfs-client-2.10.8-202.el7.x86_64.rpm

ENTRYPOINT ["/acs/entrypoint.sh"]
