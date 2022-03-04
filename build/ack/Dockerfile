FROM registry.cn-hangzhou.aliyuncs.com/plugins/centos:7.9.2009
LABEL maintainers="Alibaba Cloud Authors"
LABEL description="Alibaba Cloud CSI Plugin"

RUN yum install -y ca-certificates file tzdata nfs-utils xfsprogs e4fsprogs pciutils iputils strace cloud-utils-growpart gdisk nc

ARG ossfsVer=1.80.6.ack.1
ARG update_ossfsVer=1.86.1.ack.1
RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/ossfs/ossfs_${ossfsVer}-b42b3a8_centos7.0_x86_64.rpm -o /root/ossfs_${ossfsVer}_centos7.0_x86_64.rpm
RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}-ca30fc4_centos7.0_x86_64.rpm -o /root/ossfs_1.86.4_centos7.0_x86_64.rpm
RUN curl http://ack-csiplugin.oss-cn-hangzhou.aliyuncs.com/pre/ossfs/ossfs_${update_ossfsVer}-75ed4f1_centos8_x86_64.rpm -o /root/ossfs_1.86.4_centos8_x86_64.rpm
RUN curl https://aliyun-unas.oss-cn-beijing.aliyuncs.com/unas-4.19.91-24-for-ack-poc-1.3.rpm -o /root/unas-4.19.91-24-for-ack-poc-1.3.rpm
RUN curl https://aliyun-unas.oss-cn-beijing.aliyuncs.com/unas-kernel-3.10.0-1160.15.2-for-ack.rpm -o /root/unas-kernel-3.10.0-1160.15.2-for-ack.rpm
RUN curl https://aliyun-unas.oss-cn-beijing.aliyuncs.com/unas-kernel-5.10.23-5-for-ack.rpm -o /root/unas-kernel-5.10.23-5-for-ack.rpm

COPY nsenter /nsenter
COPY jindofs-fuse-3.7.3-20211207.tar.gz /jindofs-fuse-3.7.3-20211207.tar.gz
COPY entrypoint.sh /entrypoint.sh
COPY csiplugin-connector /bin/csiplugin-connector
COPY csiplugin-connector.service /bin/csiplugin-connector.service
COPY plugin.csi.alibabacloud.com /bin/plugin.csi.alibabacloud.com
RUN chmod +x /bin/plugin.csi.alibabacloud.com && chmod +x /entrypoint.sh && chmod +x /bin/csiplugin-connector && chmod +x /bin/csiplugin-connector.service && tar zvxf /jindofs-fuse-3.7.3-20211207.tar.gz -C /

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ENTRYPOINT ["/entrypoint.sh"]
