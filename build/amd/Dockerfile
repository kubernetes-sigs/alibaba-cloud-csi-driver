FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.9.2009
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

ARG ossfsVer=1.80.6.ack.1
ARG update_ossfsVer=1.88.0

RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${ossfsVer}-b42b3a8_centos7.0_x86_64.rpm -o /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}_centos7.0_x86_64.rpm -o /root/ossfs_${update_ossfsVer}_centos7.0_x86_64.rpm
RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}_centos8_x86_64.rpm -o /root/ossfs_${update_ossfsVer}_centos8_x86_64.rpm
RUN curl https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-efc-1.2-3.x86_64.rpm -o /root/alinas-efc-1.2-3.x86_64.rpm
RUN curl https://aliyun-encryption.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-1.1-6.al7.noarch.rpm -o /root/aliyun-alinas-utils-1.1-6.al7.noarch.rpm

RUN mkdir /csi && \
    yum install -y ca-certificates file tzdata nfs-utils xfsprogs e4fsprogs pciutils iputils strace cloud-utils-growpart gdisk nc telnet

COPY . /csi
COPY amd64-entrypoint.sh /entrypoint.sh
COPY nsenter /nsenter
COPY nsenter /usr/bin/nsenter
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com

RUN chmod +x /csi/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh && chmod +x /csi/csiplugin-connector && chmod +x /csi/csiplugin-connector.service && \
     ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone

ENTRYPOINT ["/entrypoint.sh"]
