FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.9.2009
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

ARG ossfsVer=1.80.6.ack.1
ARG update_ossfsVer=1.86.1.ack.1
RUN mkdir /csi && \
    yum install -y ca-certificates file tzdata nfs-utils xfsprogs e4fsprogs pciutils iputils strace cloud-utils-growpart gdisk nc && \
    curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${ossfsVer}-b42b3a8_centos7.0_x86_64.rpm -o /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm && \
    curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}-ca30fc4_centos7.0_x86_64.rpm -o /root/ossfs_1.86.4_centos7.0_x86_64.rpm && \
    curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}-75ed4f1_centos8_x86_64.rpm -o /root/ossfs_1.86.4_centos8_x86_64.rpm && \
    curl https://aliyun-alinas-eac.oss-cn-beijing.aliyuncs.com/alinas-eac-1.1-1.alios7.x86_64.rpm -o /root/alinas-eac-1.1-1.alios7.x86_64.rpm && \
    curl https://aliyun-encryption.oss-cn-beijing.aliyuncs.com/aliyun-alinas-utils-1.1-2.alios7.noarch.rpm -o /root/aliyun-alinas-utils-1.1-2.alios7.noarch.rpm
COPY . /csi
COPY amd64-entrypoint.sh /entrypoint.sh

RUN chmod +x /csi/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh && chmod +x /csi/csiplugin-connector && chmod +x /csi/csiplugin-connector.service && \
    tar zvxf /csi/jindofs-fuse-3.7.3-20211207.tar.gz -C / && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

ENTRYPOINT ["/entrypoint.sh"]
